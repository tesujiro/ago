package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const scriptDir = "./test"
const scriptExt = ".goa"
const scriptOk = ".ok"
const scriptInput = ".in"

type test struct {
	script string
	in     string
	ok     string
}

func TestGoa(t *testing.T) {
	files, err := ioutil.ReadDir(scriptDir)
	if err != nil {
		panic(err)
	}

	exists := func(name string) bool {
		_, err := os.Stat(name)
		return !os.IsNotExist(err)
	}
	tests := []test{}
	for _, file := range files {
		if filepath.Ext(file.Name()) == scriptExt {
			var script, in, ok string
			script = filepath.Join(scriptDir, file.Name())
			basename := file.Name()[:len(file.Name())-len(filepath.Ext(file.Name()))]
			in_file := filepath.Join(scriptDir, basename+scriptInput)
			if exists(in_file) {
				in = in_file
			}
			ok_file := filepath.Join(scriptDir, basename+scriptOk)
			if exists(ok_file) {
				ok = ok_file
			}
			tests = append(tests, test{script: script, in: in, ok: ok})
		}
	}

	fmt.Println("tests:", tests)

	realStdin := os.Stdin
	realStdout := os.Stdout
	realStderr := os.Stderr

	// IN PIPE
	//readFromIn, writeToIn, err := os.Pipe()
	readFromIn, _, err := os.Pipe()
	if err != nil {
		t.Fatal("Pipe error:", err)
	}
	os.Stdin = readFromIn
	//logger.Print("pipe in created")

	for _, test := range tests {

		// OUT PIPE
		readFromOut, writeToOut, err := os.Pipe()
		if err != nil {
			os.Stdin = realStdin
			os.Stderr = realStderr
			t.Fatal("Pipe error:", err)
		}
		os.Stdout = writeToOut
		//logger.Print("pipe out created")

		// Read Stdout goroutine
		readerOut := bufio.NewScanner(readFromOut)
		chanOut := make(chan string)
		go func() {
			for readerOut.Scan() {
				chanOut <- readerOut.Text()
			}
			close(chanOut)
			return
		}()

		// Read Script
		fp_script, err := os.Open(test.script)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		test_script, _ := ioutil.ReadAll(fp_script)

		// Run Script goroutine
		go func() {
			runScript(string(test_script), test.in)
			//close(chanDone) //NG
			writeToOut.Close()
		}()

		// Get Result
		var resultOut string
	LOOP:
		for {
			select {
			case dataOut, ok := <-chanOut:
				if !ok {
					break LOOP
				}
				dataOut = strings.TrimSpace(dataOut)
				resultOut = fmt.Sprintf("%s%s%s", resultOut, dataOut, "\n")
			}
		}

		// Read OK iile
		fp_ok, err := os.Open(test.ok)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		test_ok, _ := ioutil.ReadAll(fp_ok)

		// Result Check
		if resultOut != string(test_ok) {
			t.Fatalf("Stdout - received: %v - expected: %v - runSource: %v", resultOut, string(test_ok), string(test_script))
		}
	}

	os.Stdin = realStdin
	os.Stderr = realStderr
	os.Stdout = realStdout
}
