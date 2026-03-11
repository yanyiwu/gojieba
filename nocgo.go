//go:build !cgo

package gojieba

// Force a package-level compile error with a message that explains how to
// build gojieba correctly when cgo is unavailable.
const _ = "gojieba requires cgo; cross-compilation needs CGO_ENABLED=1 and a target C/C++ toolchain configured via CC/CXX" - 1
