package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

const scriptPath = "./goa_test.json"

type Test struct {
	Script string `json:script`
	In     string `json:in`
	Ok     string `json:ok`
}

func TestGoaJson(t *testing.T) {
	tests := []Test{}

	bytes, err := ioutil.ReadFile(scriptPath)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(bytes, &tests); err != nil {
		panic(err)
	}
	//fmt.Println("tests:", tests)

	realStdin := os.Stdin
	realStdout := os.Stdout
	realStderr := os.Stderr

	for _, test := range tests {
		fmt.Fprintf(realStdin, "script:%v\n", test.Script)

		// IN PIPE
		readFromIn, writeToIn, err := os.Pipe()
		if err != nil {
			t.Fatal("Pipe error:", err)
		}
		os.Stdin = readFromIn
		_ = writeToIn
		//logger.Print("pipe in created")

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

		// Run Script goroutine
		go func() {
			runScript(string(test.Script), test.In)
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

		// Result Check
		//fmt.Fprintf(realStdout, "result:[%v]\ttest.Ok:[%v]\n", resultOut, test.Ok)
		if resultOut != strings.Replace(test.Ok, "\r", "", -1) { //replace for Windows
			t.Fatalf("Stdout - received: %v - expected: %v - runSource: %v", resultOut, test.Ok, test.Script)
		}
	}

	os.Stdin = realStdin
	os.Stderr = realStderr
	os.Stdout = realStdout
}
