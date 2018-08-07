package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

const scriptPath = "./goa_test.json"

type test struct {
	script string
	in     string
	ok     string
}

func TestGoa(t *testing.T) {
	tests := []test{
		//BASIC EXPRESSION
		{script: "BEGIN{print 1+1}", ok: "2\n"},
		{script: "BEGIN{print 1+2}", ok: "3\n"},
		{script: "BEGIN{print nil}", ok: "<nil>\n"},
		{script: "BEGIN{print 1}", ok: "1\n"},
		//{script: "BEGIN{print 9223372036854775807}", ok: "9223372036854775807\n"},
		{script: "BEGIN{print 1.1}", ok: "1.1\n"},
		{script: "BEGIN{print 123}", ok: "123\n"},
		{script: "BEGIN{print 123.456}", ok: "123.456\n"},
		{script: "BEGIN{print \"abc\"}", ok: "abc\n"},
		{script: "BEGIN{print +1+4}", ok: "5\n"},
		{script: "BEGIN{print -1+3}", ok: "2\n"},
		{script: "BEGIN{print 1+1}", ok: "2\n"},
		{script: "BEGIN{print 1+1.1}", ok: "2.1\n"},
		{script: "BEGIN{print 1.1+4}", ok: "5.1\n"},
		{script: "BEGIN{print 1.1+1.1}", ok: "2.2\n"},
		{script: "BEGIN{print 3-1.1}", ok: "1.9\n"},
		{script: "BEGIN{print 2.2-1.1}", ok: "1.1\n"},
		{script: "BEGIN{print 3-1-1}", ok: "1\n"},
		{script: "BEGIN{print 3-(1-1)}", ok: "3\n"},
		{script: "BEGIN{print 3*5}", ok: "15\n"},
		{script: "BEGIN{print 1.5*2}", ok: "3\n"},
		{script: "BEGIN{print 5*1.2}", ok: "6\n"},
		{script: "BEGIN{print 15/5}", ok: "3\n"},
		{script: "BEGIN{print 16/5}", ok: "3\n"},
		{script: "BEGIN{print 3/1.5}", ok: "2\n"},
		{script: "BEGIN{print 3/0}", ok: "error:devision by zero\n"},
		{script: "BEGIN{print 15%5}", ok: "0\n"},
		{script: "BEGIN{print 16%5}", ok: "1\n"},
		{script: "BEGIN{print 15%4.1}", ok: "3\n"},
		{script: "BEGIN{print \"a b c\"+\" d e f\"}", ok: "a b c d e f\n"},
		{script: "BEGIN{print \"a b c\"-\" d e f\"}", ok: "0\n"},
		{script: "BEGIN{print 15.2%7.1}", ok: "1\n"},

		// variable and scope
		// builtin
		{script: "BEGIN{NF++;print NF}", ok: "1\n"},
		{script: "BEGIN{NF++}END{print NF}", ok: "1\n"},
		{script: "BEGIN{NF=1}END{print NF}", ok: "1\n"},
		{script: "BEGIN{NF=1.1}END{print NF}", ok: "error:type of NF must be int ,not float64.\n"},
		{script: "BEGIN{NF=\"aaa\"}", ok: "error:type of NF must be int ,not string.\n"},
		{script: "BEGIN{$0=\"aaa\";print}", ok: "error:Field Index Out of Range:0\n"},
		{script: "BEGIN{$1=\"aaa\";print}", ok: "aaa\n"},
		{script: "BEGIN{print FS}", ok: "\n"},
		{script: "BEGIN{FS=\"X\"}END{print FS}", ok: "X\n"},
		{script: "BEGIN{FS=123}END{print FS}", ok: "error:type of FS must be string ,not int.\n"},
		// global
		{script: "BEGIN{A=1}END{print A}", ok: "1\n"},
		{script: "BEGIN{Abc=1}END{print Abc}", ok: "1\n"},
		{script: "BEGIN{ABC=1}END{print ABC}", ok: "1\n"},
		{script: "BEGIN{A++;print A}", ok: "1\n"},
		{script: "BEGIN{A++}END{print A}", ok: "1\n"},
		{script: "BEGIN{A=1.1;print A}", ok: "1.1\n"},
		{script: "BEGIN{A=1.1}END{print A}", ok: "1.1\n"},
		{script: "BEGIN{A=\"AAA\";print A}", ok: "AAA\n"},
		{script: "BEGIN{A=\"AAA\"}END{print A}", ok: "AAA\n"},
		// local
		{script: "BEGIN{l=1}END{print l}", ok: "\n"},
		{script: "BEGIN{lmn=1}END{print lmn}", ok: "\n"},
		{script: "BEGIN{lMN=1}END{print lMN}", ok: "\n"},
		{script: "BEGIN{l++;print l}", ok: "1\n"},
		{script: "BEGIN{l++}END{print l}", ok: "\n"},
		{script: "BEGIN{l=1;print l}", ok: "1\n"},
		{script: "BEGIN{l=1.1;print l}", ok: "1.1\n"},
		{script: "BEGIN{l=1.1}END{print l}", ok: "\n"},
		{script: "BEGIN{l=\"AAA\";print l}", ok: "AAA\n"},
		{script: "BEGIN{l=\"AAA\"}END{print l}", ok: "\n"},

		// bool expression
		{script: "BEGIN{print true}", ok: "true\n"},
		{script: "BEGIN{print false}", ok: "false\n"},
		{script: "BEGIN{print !true}", ok: "false\n"},
		{script: "BEGIN{print !false}", ok: "true\n"},
		{script: "BEGIN{print !0}", ok: "true\n"},
		{script: "BEGIN{print !1}", ok: "false\n"},
		{script: "BEGIN{print !11}", ok: "false\n"},
		{script: "BEGIN{print !1.1}", ok: "false\n"},
		{script: "BEGIN{print !\"\"}", ok: "true\n"},
		{script: "BEGIN{print !\"aa\"}", ok: "false\n"},
		{script: "BEGIN{a=1;b=2;print a+1==b}", ok: "true\n"},
		{script: "BEGIN{a=1;b=2;print a+1!=b}", ok: "false\n"},
		{script: "BEGIN{a=1;b=2;print a!=b}", ok: "true\n"},
		{script: "BEGIN{a=1;b=2;print !(a+1==b)}", ok: "false\n"},
		{script: "BEGIN{a=1;b=2;print !(a==b)}", ok: "true\n"},
		{script: "BEGIN{a=1;b=2;print a<b}", ok: "true\n"},
		{script: "BEGIN{a=1;b=1;print a<=b}", ok: "true\n"},
		{script: "BEGIN{a=1;b=2;print a>b}", ok: "false\n"},
		{script: "BEGIN{a=1;b=1;print a>=b}", ok: "true\n"},
		{script: "BEGIN{a=1;b=0.1;print a>b}", ok: "true\n"},
		{script: "BEGIN{a=1;print a==1}", ok: "true\n"},
		{script: "BEGIN{a=1;b=2;print a==1&&b==2}", ok: "true\n"},
		{script: "BEGIN{a=1;b=2;print a==2&&b==2}", ok: "false\n"},
		{script: "BEGIN{a=1;b=2;print a==1||b==2}", ok: "true\n"},
		{script: "BEGIN{a=1;b=2;print a==2||b==2}", ok: "true\n"},
		{script: "BEGIN{print 12&&34}", ok: "error:cannot convert to bool\n"},

		// assignment
		{script: "BEGIN{a=1;b=2;print a+b}", ok: "3\n"},
		{script: "BEGIN{a=1;b=2;print a+1==b}", ok: "true\n"},
		{script: "BEGIN{a=1;b=2;print a+1!=b}", ok: "false\n"},
		{script: "BEGIN{a=1;b=2;print a<b}", ok: "true\n"},
		{script: "BEGIN{a=1;b=1;print a<=b}", ok: "true\n"},
		{script: "BEGIN{a=1;b=2;print a>b}", ok: "false\n"},
		{script: "BEGIN{a=1;b=1;print a>=b}", ok: "true\n"},
		{script: "BEGIN{a=1;b=0.1;print a>b}", ok: "true\n"},
		{script: "BEGIN{a=1;b=0.1;c=15;print (a+b)*c}", ok: "16.5\n"},
		{script: "BEGIN{a=1;b=0.1;c=15;print (a+b)*c/0.5}", ok: "33\n"},

		// composite expression
		{script: "BEGIN{a=10;print a++}", ok: "11\n"},
		{script: "BEGIN{a=1.9;print a++}", ok: "2.9\n"},
		{script: "BEGIN{print 10++}", ok: "error:Invalid operation\n"},
		{script: "BEGIN{a=\"a\";print a++}", ok: "1\n"},
		{script: "BEGIN{a=10;print a--}", ok: "9\n"},
		{script: "BEGIN{a=2.9;print a--}", ok: "1.9\n"},
		{script: "BEGIN{print 10--}", ok: "error:Invalid operation\n"},
		{script: "BEGIN{a=\"a\";print a--}", ok: "-1\n"},
		{script: "BEGIN{a=10;a+=2;print a}", ok: "12\n"},
		{script: "BEGIN{a=10;a+=2.5;print a}", ok: "12.5\n"},
		{script: "BEGIN{a=10;a+=2;print a}", ok: "12\n"},
		{script: "BEGIN{a=10;a-=2;print a}", ok: "8\n"},
		{script: "BEGIN{a=10;a*=2;print a}", ok: "20\n"},
		{script: "BEGIN{a=10;a/=2;print a}", ok: "5\n"},

		// multi expressions
		{script: "BEGIN{a,b=1,2;print a}", ok: "1\n"},
		{script: "BEGIN{a,b=1,2;print b}", ok: "2\n"},
		{script: "BEGIN{a,b=1,2,3;print b}", ok: "2\n"},
		{script: "BEGIN{a,b,c=1,2;print b}", ok: "2\n"},

		// if statement
		{script: "BEGIN{a=1;if a==1 { a=2 ;print a;}}", ok: "2\n"},
		{script: "BEGIN{a=1;if a==1 { a=2 };print a}", ok: "2\n"},
		{script: "BEGIN{a=1;if a==1 { env_test=2 };print env_test}", ok: "\n"},
		{script: "BEGIN{a=2;if a==1 { a=2 } else { a=3;b=4;print b }}", ok: "4\n"},
		{script: "BEGIN{a=1;b=1;if a==1 { b=2 };print b}", ok: "2\n"},
		{script: "BEGIN{a=1;if a==1 { a=11 } else if a==2 { a=12 } else if a==3 { a=13 };print a}", ok: "11\n"},
		{script: "BEGIN{a=2;if a==1 { a=11 } else if a==2 { a=12 } else if a==3 { a=13 };print a}", ok: "12\n"},
		{script: "BEGIN{a=3;if a==1 { a=11 } else if a==2 { a=12 } else if a==3 { a=13 };print a}", ok: "13\n"},
		{script: "BEGIN{a=1;if a==1 { env_test=11 } else if a==2 { env_test=12 } else { env_test=13 };print env_test}", ok: "\n"},
		{script: "BEGIN{a=2;if a==1 { env_test=11 } else if a==2 { env_test=12 } else { env_test=13 };print env_test}", ok: "\n"},
		{script: "BEGIN{a=3;if a==1 { env_test=11 } else if a==2 { env_test=12 } else { env_test=13 };print env_test}", ok: "\n"},

		// for statement
		{script: "BEGIN{a=0;for{ if a==10 { break }; a= a+1 };print a}", ok: "10\n"},
		{script: "BEGIN{a=0;b=0;for{ a=a+1;if a==10 { break }; if b==5 {continue};b= b+1 };print b}", ok: "5\n"},
		{script: "BEGIN{a=0;for a<=10 { a= a+1 };print a}", ok: "11\n"},
		{script: "BEGIN{a=0;for b { a= a+1 };print a}", ok: "error:for condition type string cannot convert to bool\n"},
		{script: "BEGIN{a=0;for a { a= a+1 };print a}", ok: "error:for condition type int cannot convert to bool\n"},
		//{script: "BEGIN{a=0;for{ a=a+1;if a==10 { return a; };};print a}", ok: "10\n"},
		//{script: "BEGIN{a=0;for{ a=10;return a };print a}", ok: "10\n"},

		// map: awk-array (associated array = map)
		{script: "BEGIN{a[1]=1;print a[1]}", ok: "1\n"},
		{script: "BEGIN{a[1]=1;print a[2]}", ok: "\n"},
		{script: "BEGIN{a[\"a\"]=1;print a[\"a\"]}", ok: "1\n"},
		{script: "BEGIN{a[1,2]=1;print a[1,2]}", ok: "1\n"},
		//{script: "BEGIN{a[1]=1;a=2;print a[1]}", ok: "some error\n"},
		{script: "BEGIN{a[\"a\"]=1;print a[\"a\"]}", ok: "1\n"},
		{script: "BEGIN{a[1]=1;print a[2]}", ok: "\n"},
		//{script: "BEGIN{a[1]=1;print a}", ok: "error\n"},
		//{script: "BEGIN{a[1]=1;a[2]=2;for (i in a) {print i,a[i]}}", ok: "1 1\n2 2\n"},
		//{script: "BEGIN{a[\"1\"]=1;a[\"2\"]=2;for (i in a) {print i,a[i]}}", ok: "1 1\n2 2\n"},

		// map
		//{script: "BEGIN{a[1]=1;print a[1]}", ok: "1\n"},
		//{script: "BEGIN{a[1]=1;print a[2]}", ok: "\n"},
		//{script: "BEGIN{a[1]=1;print a}", ok: "error\n"},
		//{script: "BEGIN{a[\"a\"]=1;print a[\"a\"]}", ok: "1\n"},
		//{script: "BEGIN{a[\"a\"]=1;print a[\"b\"]}", ok: "\n"},
		//{script: "BEGIN{a[1]=\"a\";print a[1]}", ok: "a\n"},
		//{script: "BEGIN{a[1]=\"a\";print a[2]}", ok: "\n"},
		//{script: "BEGIN{a[\"a\"]=\"a\";print a[\"a\"]}", ok: "a\n"},
		//{script: "BEGIN{a[\"a\"]=\"a\";print a[\"b\"]}", ok: "\n"},

		// command parameter

		// field
		{script: "{print $1}", in: "Hello World!\n", ok: "Hello\n"},
		{script: "$1==\"AAA\"{print;COUNT++} END{print COUNT}", in: "AAA BBB CCC\nAAA BBB CCC\n", ok: "AAA BBB CCC\nAAA BBB CCC\n2\n"},
		{script: "NR==1{$2=$1 ;print $0,NF} NR==2{$5=$1; print $0,NF}", in: "AAA BBB CCC\nAAA BBB CCC\n", ok: "AAA AAA CCC 3\nAAA BBB CCC  AAA 5\n"},
	}

	//fmt.Println("tests:", tests)

	realStdin := os.Stdin
	realStdout := os.Stdout
	realStderr := os.Stderr

	for _, test := range tests {
		//t.Logf("script:%v\n", test.script)

		// IN PIPE
		readFromIn, writeToIn, err := os.Pipe()
		if err != nil {
			t.Fatal("Pipe error:", err)
		}
		//os.Stdin.Sync()
		os.Stdin = readFromIn
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
			script_reader := strings.NewReader(test.script)
			runScript(script_reader, os.Stdin)
			//close(chanDone) //NG
			writeToOut.Close()
		}()

		// Write to Stdin goroutine
		go func() {
			scanner := bufio.NewScanner(strings.NewReader(test.in))
			for scanner.Scan() {
				_, err = writeToIn.WriteString(scanner.Text() + "\n")
				if err != nil {
					t.Fatal("Stdin WriteString error:", err)
				}
			}
			//readFromIn.Close() //NG
			writeToIn.Close()
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
		//fmt.Fprintf(realStdout, "result:[%v]\ttest.ok:[%v]\n", resultOut, test.ok)
		if resultOut != strings.Replace(test.ok, "\r", "", -1) { //replace for Windows
			t.Errorf("Stdout - received: %v - expected: %v - runSource: %v", resultOut, test.ok, test.script)
		}
	}

	os.Stdin = realStdin
	os.Stderr = realStderr
	os.Stdout = realStdout
}
