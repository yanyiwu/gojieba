//go:build !cgo

package gojieba

// Force a package-level compile error with a message that explains how to
// build gojieba correctly when cgo is unavailable.
var _ int = "gojieba requires cgo (CGO_ENABLED=1) and a C/C++ toolchain for cross-compilation"
