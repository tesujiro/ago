package debug

import "fmt"

var debug bool

func On() {
	debug = true
}

func Off() {
	debug = false
}

func Println(a ...interface{}) {
	if debug {
		fmt.Println(a...)
	}
}

func Printf(format string, a ...interface{}) {
	if debug {
		fmt.Printf(format, a...)
	}
}
