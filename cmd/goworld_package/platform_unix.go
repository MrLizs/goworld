//go:build !windows
// +build !windows

package goworld_package

import "syscall"

const (
	// BinaryExtension extension used on unix
	BinaryExtension = ""
	// StopSignal syscall used to stop server
	StopSignal = syscall.SIGTERM
)
