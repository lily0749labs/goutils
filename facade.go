// Package goutils provides facade-style access to the utility subpackages.
package goutils

import (
	anyutil "github.com/lily0749labs/goutils/anyto"
	cryptoutil "github.com/lily0749labs/goutils/crypto"
	idutil "github.com/lily0749labs/goutils/id"
	randutil "github.com/lily0749labs/goutils/rand"
	strutil "github.com/lily0749labs/goutils/strto"
	timeutil "github.com/lily0749labs/goutils/time"
	validutil "github.com/lily0749labs/goutils/valid"
)

var (
	// AnyTo provides facade-style access to the any-value conversion utilities.
	AnyTo = anyutil.AnyTo
	// Crypto provides facade-style access to the cryptography utilities.
	Crypto = cryptoutil.Crypto
	// ID provides facade-style access to the ID generation utilities.
	ID = idutil.ID
	// Rand provides facade-style access to the random utilities.
	Rand = randutil.Rand
	// StrTo provides facade-style access to the string conversion utilities.
	StrTo = strutil.StrTo
	// Time provides facade-style access to the time utilities.
	Time = timeutil.Time
	// Valid provides facade-style access to the validation utilities.
	Valid = validutil.Valid
)
