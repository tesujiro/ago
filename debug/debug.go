// Package debug is a debug library for ago.
package debug

import "fmt"

var debug bool

// On sets debug mode on.
func On() {
	debug = true
}

// Off sets debug mode off.
func Off() {
	debug = false
}

// Println formats using the default formats for its operands and writes to standard output.
func Println(a ...interface{}) {
	if debug {
		fmt.Println(a...)
	}
}

// Printf formats according to a format specifier and writes to standard output.
func Printf(format string, a ...interface{}) {
	if debug {
		fmt.Printf(format, a...)
	}
}
