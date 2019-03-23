package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func BenchmarkLoop(b *testing.B) {
	script := fmt.Sprintf("BEGIN{for i=0;i<%v;i++{}}", b.N)
	os.Args = []string{"ago", "-c"} // -c : for cpu profiler
	os.Args = append(os.Args, script)
	rc := _main()
	if rc != 0 {
		log.Printf("return error code:%v\n", rc)
		os.Exit(1)
	}
	return
}
