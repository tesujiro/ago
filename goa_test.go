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
		{script: "BEGIN{print nil}", ok: "\n"},
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
		{script: "BEGIN{print \"a b c\"+1234}", ok: "1234\n"},
		{script: "BEGIN{print \"a b c\"+\" d e f\"}", ok: "a b c d e f\n"},
		{script: "BEGIN{print \"a b c\"-\" d e f\"}", ok: "0\n"},
		{script: "BEGIN{print 15.2%7.1}", ok: "1\n"},
		{script: "BEGIN{a=123;print a}", ok: "123\n"},
		{script: "BEGIN{map=123;print map}", ok: "123\n"},
		// composite expression
		{script: "BEGIN{a=123;a++;print a}", ok: "124\n"},
		//{script: "BEGIN{a=123;print a++}", ok: "123\n"}, //TODO:AWK
		//{script: "BEGIN{a=123;print ++a}", ok: "123\n"}, //TODO:AWK
		{script: "BEGIN{a=123.4;a++;print a}", ok: "124.4\n"},
		{script: "BEGIN{a=123;a--;print a}", ok: "122\n"},
		{script: "BEGIN{a=123.4;a--;print a}", ok: "122.4\n"},
		{script: "BEGIN{a=123;a+=4;print a}", ok: "127\n"},
		{script: "BEGIN{a=123;a-=4;print a}", ok: "119\n"},
		{script: "BEGIN{a=123;a*=4;print a}", ok: "492\n"},
		{script: "BEGIN{a=123;a/=4;print a}", ok: "30\n"},
		//{script: "BEGIN{a=123;a%=4;print a}", ok: "30\n"}, //TODO

		// JAPANESE
		{script: "BEGIN{print \"あいう\"}", ok: "あいう\n"},
		{script: "BEGIN{a=\"あいう\";b=\"えお\";print a+b}", ok: "あいうえお\n"},
		{script: "BEGIN{a[\"あ\"]=1;a[\"い\"]=2;a[\"う\"]=3;for(key in a){print key,a[key]}}", ok: "あ 1\nい 2\nう 3\n"},

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

		// regular expression
		{script: "BEGIN{print \"aaa\"~\"/aaa/\"}", ok: "true\n"},
		{script: "BEGIN{print \"aaa\"~\"/abc/\"}", ok: "false\n"},
		{script: "BEGIN{print \"aaa\"~\"/a+/\"}", ok: "true\n"},
		{script: "BEGIN{print \"aaa\"~\"/^a+$/\"}", ok: "true\n"},
		{script: "BEGIN{print \"abc\"~\"/^a+$/\"}", ok: "false\n"},
		{script: "\"AAA\"~\"/AAA/\"{print}", in: "AAA", ok: "AAA\n"},
		{script: "$0~\"/AAA/\"{print}", in: "AAA", ok: "AAA\n"},
		{script: "\"/AAA/\"{print}", in: "AAA", ok: "AAA\n"},

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
		{script: "BEGIN{print 10++}", ok: "error:invalid operation\n"},
		{script: "BEGIN{a=\"a\";print a++}", ok: "1\n"},
		{script: "BEGIN{a=10;print a--}", ok: "9\n"},
		{script: "BEGIN{a=2.9;print a--}", ok: "1.9\n"},
		{script: "BEGIN{print 10--}", ok: "error:invalid operation\n"},
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
		{script: "BEGIN{print a[1]}", ok: "\n"},
		{script: "BEGIN{a[1]=1;print a[1]}", ok: "1\n"},
		{script: "BEGIN{a[1]=1;print a[2]}", ok: "\n"},
		{script: "BEGIN{a[1]=1;a[2]=2;print a[1],a[2]}", ok: "1 2\n"},
		{script: "BEGIN{a[1]=1;a[1]=2;print a[1]}", ok: "2\n"},
		{script: "BEGIN{a[1]=\"a\";print a[1]}", ok: "a\n"},
		{script: "BEGIN{a[1]=\"a\";print a[1]+a[2]}", ok: "a\n"},
		{script: "BEGIN{a[\"a\"]=1;print a[\"a\"]}", ok: "1\n"},
		{script: "BEGIN{a[1,2]=1;print a[1,2]}", ok: "1\n"},
		{script: "BEGIN{a[1]=1;a=2}", ok: "error:can't assign to a; it's an associated array name.\n"},
		{script: "BEGIN{a[\"a\"]=1;print a[\"a\"]}", ok: "1\n"},
		{script: "BEGIN{a[1]=1;print a[2]}", ok: "\n"},
		{script: "BEGIN{a[1]=1;print a}", ok: "error:type map does not support print operation\n"},
		{script: "BEGIN{a[1]=1;a[1]++;print a[1]}", ok: "2\n"},
		{script: "BEGIN{a[1]=1;a[1]--;print a[1]}", ok: "0\n"},
		{script: "BEGIN{a[1]=1;a[2]=2;a[3]=3;b=a;b[1]=10;print a[1]}", ok: "10\n"}, // THIS SPEC OK?
		{script: "BEGIN{a[1]=1;delete a[1];print a[1]}", ok: "\n"},
		{script: "BEGIN{a[1]=1;delete a[1];print a[2]}", ok: "\n"},
		{script: "BEGIN{a[1]=1;delete a[2];print a[1]}", ok: "1\n"},
		{script: "BEGIN{a[1]=1;delete a;print a[1]}", ok: "\n"},
		{script: "BEGIN{a=1;delete a}", ok: "error:type int does not support delete operation\n"},
		{script: "BEGIN{a[1]=1;delete a;a=2}", ok: "error:can't assign to a; it's an associated array name.\n"},
		{script: "BEGIN{delete a;a=2}", ok: "error:can't assign to a; it's an associated array name.\n"},
		{script: "BEGIN{list=func(){a[1]=1;a[2]=2;a[3]=3;return a};delete list()[1]}", ok: "error:non variable does not support delete operation\n"},
		{script: "BEGIN{list=func(){a[1]=1;a[2]=2;a[3]=3;return a};list()[1]=3}", ok: "error:invalid assignment\n"},
		{script: "BEGIN{a[1]=1;a[2]=2;for (i in a) {print i,a[i]}}", ok: "1 1\n2 2\n"},
		{script: "BEGIN{a[1]=1;a[2]=2;for (i in a) {};print i}", ok: "\n"},
		{script: "BEGIN{a[1]=1;a[2]=2;i=0;for (i in a) {};print i}", ok: "2\n"},
		{script: "BEGIN{for (i in a) {print i}}", ok: "error:unknown symbol\n"},
		{script: "BEGIN{a=1;for (i in a) {print i}}", ok: "error:for key loop not in associated array,int\n"},
		{script: "BEGIN{a[\"1\"]=1;a[\"2\"]=2;for (i in a) {print i,a[i]}}", ok: "1 1\n2 2\n"},
		{script: "BEGIN{a[1]=1;a[2]=2;for (i in a) {print i,a[i]}}", ok: "1 1\n2 2\n"},
		{script: "BEGIN{a[1]++;a[2]=2;for (i in a) {print i,a[i]}}", ok: "1 1\n2 2\n"},
		{script: "BEGIN{a[1]=1;a[2]=2;a[3]=3;for (i in a) {print i;if i==\"2\" { break }}}", ok: "1\n2\n"},
		{script: "{A[$0]++} END{for(key in A){print key}}", in: "AAA", ok: "AAA\n"},

		// function
		{script: "BEGIN{add=func(a,b){return a+b}; print add(10,5)}", ok: "15\n"},
		{script: "BEGIN{add=func(a,b){return a+b}; print add(1.1,2.1)}", ok: "3.2\n"},
		{script: "BEGIN{add=func(a,b){return a+b}; print add(\"あ\",\"いう\")}", ok: "あいう\n"},
		{script: "BEGIN{print func(a,b){return a+b}(10,5)}", ok: "15\n"},
		{script: "BEGIN{a=123;add=func(a,b){return a+b}; add(10,5);print a}", ok: "123\n"},
		{script: "BEGIN{c=100;add=func(a,b){return a+b+c}; print add(10,5)}", ok: "115\n"},
		{script: "BEGIN{one=func(){return 1}; print one()}", ok: "1\n"},
		{script: "BEGIN{a=10;plusone=func(){a++};plusone();print a}", ok: "11\n"},
		{script: "BEGIN{print func(){return 1}()}", ok: "1\n"},
		//{script: "BEGIN{a=10;func plusone(){a++;return};plusone();print a}", ok: "11\n"}, //TODO: panic
		{script: "BEGIN{hash=func(){m[1]=1;m[2]=2;m[3]=3;return m}; m=hash();print m[1]}", ok: "1\n"},
		{script: "BEGIN{map=func(){m[1]=1;m[2]=2;m[3]=3;return m}; print map()[1]}", ok: "1\n"},
		{script: "BEGIN{print func(){m[1]=1;m[2]=2;m[3]=3;return m}()[1]}", ok: "1\n"},
		// multi result function
		{script: "BEGIN{Cross=func(a1,a2){return a2,a1;};print Cross(1,5)}", ok: "5 1\n"},
		{script: "BEGIN{Cross=func(a1,a2){return a2,a1;};x,y=Cross(1,5);print x}", ok: "5\n"},
		{script: "BEGIN{Cross=func(a1,a2){return a2,a1;};x,y=Cross(1,5);print y}", ok: "1\n"},
		{script: "BEGIN{Cross=func(a1,a2){return a2,a1;};print Cross(\"a\",\"b\")}", ok: "b a\n"},
		{script: "BEGIN{a=1;Fn=func(){a=100;};Fn();print a}", ok: "100\n"},
		// anonymous func
		{script: "BEGIN{print func (x){return x+100;}(10)}", ok: "110\n"},
		{script: "BEGIN{print func (x){return x+100;}()}", ok: "error:function wants 1 arguments but received 0\n"},
		{script: "BEGIN{print (1+1)(10)}", ok: "error:cannot call type int\n"},
		{script: "BEGIN{Fn=func (x){return func(y) {return x*10+y};};Fn2=Fn(10);print Fn2(2)}", ok: "102\n"},
		// recursive call
		{script: "BEGIN{Factorial=func(x){if x==1 {1} else { x*Factorial(x-1)}};print Factorial(3)}", ok: "6\n"},
		{script: "BEGIN{Factorial=func(x){if x==1 {return 1} else { return x*Factorial(x-1)}};print Factorial(3)}", ok: "6\n"},
		// higher order function
		{script: "BEGIN{func (x){return func(y) {return x*10+y};}()(2)}", ok: "error:function wants 1 arguments but received 0\n"},
		{script: "BEGIN{func (x){return func(y) {return x*10+y};}(10)()}", ok: "error:function wants 1 arguments but received 0\n"},
		{script: "BEGIN{print func (x){return func(y) {return x*10+y};}(10)(2)}", ok: "102\n"},
		{script: "BEGIN{Fibo=func(){x,y=0,1;return func(){x,y=y,x+y;return y}};f=Fibo();f();f();f();print f();}", ok: "5\n"},
		// higher order & recursive
		{script: "BEGIN{mod=func(x){f=func(y){ if y<x {return y} else { return f(y-x) }};return f};mod3=mod(3);print mod3(11);}", ok: "2\n"},
		{script: "BEGIN{f=func(x){if x==1 {return 1} else {return x*f(x-1)}};print f(1)}", ok: "1\n"},
		{script: "BEGIN{f=func(x){if x==1 {return 1} else {return x*f(x-1)}};print f(3)}", ok: "6\n"},
		{script: "BEGIN{f=func(x){if x==1 {return func(){ return 1}} else { return func(){ return x*f(x-1)() }}};print f(1)()}", ok: "1\n"},
		{script: "BEGIN{f=func(x){if x==1 {return func(){ return 1}} else { return func(){ return x*f(x-1)() }}};print f(3)()}", ok: "6\n"}, //ERROR
		{script: "BEGIN{f=func(x){if x==1 {return func(){ return x}} else { return func(){ return x*(x-1) }}};print f(1)()}", ok: "1\n"},
		{script: "BEGIN{f=func(x){return func(){ return x}};print f(1)()}", ok: "1\n"},
		// PROBLEM
		{script: "BEGIN{print func(x){return func(){return x}}(1)()}", ok: "1\n"},
		{script: "BEGIN{print func(x){if true {return func(){return x}}}(1)()}", ok: "1\n"},
		// func rule
		{script: "BEGIN{print 1}BEGIN{print 2}", ok: "1\n2\n"},
		{script: "BEGIN{print 1}END{print 2}", ok: "1\n2\n"},
		{script: "function one(){return 1}BEGIN{print one()}", ok: "1\n"},
		{script: "func one(){return 1}BEGIN{print one()}", ok: "1\n"},

		// command parameter

		// built in variables
		{script: "BEGIN{FS=\":\"}{print $2}", in: "AAA:BBB:CCC\nAAA:BBB:CCC\n", ok: "BBB\nBBB\n"},
		{script: "{print}", in: "AAA BBB  CCC\n", ok: "AAA BBB  CCC\n"}, // AWK is AWK
		{script: "BEGIN{OFS=\":\"}{$1=$1;print}", in: "AAA BBB  CCC\n", ok: "AAA:BBB:CCC\n"},
		//{script: "BEGIN{OFS=\"\n\"}{$1=$1;print}", in: "AAA BBB CCC\nAAA BBB CCC\n", ok: "AAA\nBBB\nCCC\nAAA\nBBB\nCCC\n"}, //TODO
		{script: "BEGIN{ORS=\":\"}{$1=$1;print}", in: "AAA BBB CCC\nCCC DDD EEE\n", ok: "AAA BBB CCC:CCC DDD EEE:\n"},

		// Numeric Built-in Functions
		// lib: int
		// lib: sqrt,exp,log
		// lib: sin,cos,atan2
		// lib: rand,arand,srand

		// Built-in Functions for String Manipulation
		// lib:len
		{script: "BEGIN{print length(\"Hello World!\")}", ok: "12\n"},
		{script: "BEGIN{print len(\"Hello World!\")}", ok: "12\n"},
		{script: "BEGIN{a[1]=1;a[2]=2;print len(a)}", ok: "2\n"},
		{script: "BEGIN{f=func(){return 1,2};print len(f())}", ok: "2\n"},
		//{script: "BEGIN{print len(123)}", ok: "invalid argument 123 (type int) for len\n"},//TODO:
		// lib:substr
		{script: "BEGIN{print substr(\"abcde\",1,3)}", ok: "abc\n"},
		{script: "BEGIN{print substr(\"abcde\",0,3)}", ok: "abc\n"},
		{script: "BEGIN{print substr(\"abcde\",-1,3)}", ok: "abc\n"},
		{script: "BEGIN{print substr(\"abcde\",1,5)}", ok: "abcde\n"},
		{script: "BEGIN{print substr(\"abcde\",1,6)}", ok: "abcde\n"},
		{script: "BEGIN{print substr(\"abcde\",3,2)}", ok: "cd\n"},
		//{script: "BEGIN{print substr(\"abcde\",3)}", ok: "cde\n"}, //TODO: Default parameter
		{script: "BEGIN{print substr(\"abcde\",2,0)}", ok: "\n"},
		{script: "BEGIN{print substr(\"abcde\",2,-1)}", ok: "\n"},
		{script: "BEGIN{print substr(12345,1,3)}", ok: "123\n"},
		{script: "BEGIN{print substr(12.345,1,4)}", ok: "12.3\n"},
		{script: "BEGIN{print substr(\"\",1,3)}", ok: "\n"},
		// lib:split
		//{script: "BEGIN{split(\"a:b:c\",ar,\":\");print ar[3]}", ok: "c\n"},
		// lib:index
		{script: "BEGIN{print index(\"abc\",\"bc\")}", ok: "2\n"},
		{script: "BEGIN{print index(\"abc\",\"yz\")}", ok: "0\n"},
		{script: "BEGIN{print index(\"\",\"yz\")}", ok: "0\n"},
		{script: "BEGIN{print index(\"abc\",\"\")}", ok: "1\n"},
		{script: "BEGIN{print index(\"\",\"\")}", ok: "0\n"},
		// lib: tolower,toupper
		{script: "BEGIN{print tolower(\"\")}", ok: "\n"},
		{script: "BEGIN{print tolower(\"Hello, World!\")}", ok: "hello, world!\n"},
		{script: "BEGIN{print toupper(\"\")}", ok: "\n"},
		{script: "BEGIN{print toupper(\"Hello, World!\")}", ok: "HELLO, WORLD!\n"},

		// field
		{script: "{print $1}", in: "Hello World!\n", ok: "Hello\n"},
		{script: "$1==\"AAA\"{print;COUNT++} END{print COUNT}", in: "AAA BBB CCC\nAAA BBB CCC\n", ok: "AAA BBB CCC\nAAA BBB CCC\n2\n"},
		{script: "NR==1{$2=$1 ;print $0,NF} NR==2{$5=$1; print $0,NF}", in: "AAA BBB CCC\nAAA BBB CCC\n", ok: "AAA AAA CCC 3\nAAA BBB CCC  AAA 5\n"},
		// MAP
		{script: `{
						COUNT[$1]++
					}
					END{
						for (key in COUNT){
							print key,COUNT[key]
						}
					}`, in: `AAA
BBB
CCC
AAA
ZZZ
AAA
CCC
`, ok: `AAA 3
BBB 1
CCC 2
ZZZ 1
`},
	}

	//fmt.Println("tests:", tests)

	realStdin := os.Stdin
	realStdout := os.Stdout
	realStderr := os.Stderr

	for _, test := range tests {
		//t.Logf("script:%v\n", test.script)
		//fmt.Fprintf(realStdout, "script:%v\n", test.script)

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
