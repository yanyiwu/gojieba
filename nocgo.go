//go:build !cgo

package gojieba

const _ = "gojieba requires cgo; cross-compilation needs CGO_ENABLED=1 and a target C/C++ toolchain configured via CC/CXX" - 1
