//go:build linux

package main

// Package fipsonly restricts all TLS configuration to FIPS-approved settings.
// This package only exists when using Go compiled with GOEXPERIMENT=boringcrypto.
// See https://go.dev/src/crypto/tls/fipsonly/fipsonly.go
import _ "crypto/tls/fipsonly"
