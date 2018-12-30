package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

type test struct {
	script  string
	in      string
	ok      string
	prepare func()
	cleanup func()
	rc      int
}

func TestGoa(t *testing.T) {
	var tempScriptPath string
	var tempDataPath string
	tests := []test{
		//BASIC EXPRESSION
		{script: "BEGIN{print 1+1}", ok: "2\n"},
		{script: "BEGIN{print 1+1}#comment", ok: "2\n"},
		{script: "BEGIN{print 1+2}", ok: "3\n"},
		{script: "BEGIN{print 1+'2'}", ok: "3\n"},
		{script: "BEGIN{print 1e3}", ok: "1000\n"},
		{script: "BEGIN{print 1e10}", ok: "1e+10\n"},
		{script: "BEGIN{print 0x0a}", ok: "10\n"},
		{script: "BEGIN{print 0x10}", ok: "16\n"},
		{script: "BEGIN{print nil}", ok: "\n"},
		{script: "BEGIN{print 1}", ok: "1\n"},
		{script: "BEGIN{print 123abc}", ok: "Syntax error: identifier starts immediately after numeric literal\n"},
		//{script: "BEGIN{print 9223372036854775807}", ok: "9223372036854775807\n"},
		{script: "BEGIN{print 1.1}", ok: "1.1\n"},
		{script: "BEGIN{print 123}", ok: "123\n"},
		{script: "BEGIN{print 123.456}", ok: "123.456\n"},
		{script: "BEGIN{print \"abc\"}", ok: "abc\n"},
		{script: "BEGIN{print 'abc'}", ok: "abc\n"},
		{script: "BEGIN{print ''}", ok: "\n"},
		{script: "BEGIN{print '\\b'}", ok: "\b\n"},
		{script: "BEGIN{print '\\r'}", ok: "\r\n"},
		{script: "BEGIN{print '[\\f]'}", ok: "[\f]\n"},
		{script: "BEGIN{print '\\c'}", ok: "c\n"},
		{script: "BEGIN{print +1+4}", ok: "5\n"},
		{script: "BEGIN{print +1.0+4}", ok: "5\n"},
		{script: "BEGIN{print -1+3}", ok: "2\n"},
		{script: "BEGIN{print -1.0+3}", ok: "2\n"},
		{script: "BEGIN{print +(1/0)}", ok: "error:devision by zero\n"},
		{script: "BEGIN{print 1+1}", ok: "2\n"},
		{script: "BEGIN{print 1+1.1}", ok: "2.1\n"},
		{script: "BEGIN{print 1.1+4}", ok: "5.1\n"},
		{script: "BEGIN{print 1.1+1.1}", ok: "2.2\n"},
		//{script: "BEGIN{print 1e2}", ok: "100\n"},
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
		{script: "BEGIN{}END{3/0}", ok: "error:devision by zero\n"},
		{script: "BEGIN{print 15%5}", ok: "0\n"},
		{script: "BEGIN{print 16%5}", ok: "1\n"},
		{script: "BEGIN{print 15%4.1}", ok: "3\n"},
		{script: "BEGIN{3%0}", ok: "error:devision by zero\n"},
		{script: "BEGIN{a[1]=1;print 3%a}", ok: "error:can't read value of array\n"},
		{script: "BEGIN{print \"a b c\"+1234}", ok: "1234\n"},
		{script: "BEGIN{print \"a b c\"+\" d e f\"}", ok: "a b c d e f\n"},
		{script: "BEGIN{print \"a b c\"-\" d e f\"}", ok: "0\n"},
		{script: "BEGIN{print \"a b c\" \" d e f\"}", ok: "a b c d e f\n"},
		{script: "BEGIN{print \"a\" \"b\"}", ok: "ab\n"},
		{script: "BEGIN{print 1 \"b\"}", ok: "1b\n"},
		{script: "BEGIN{print 1 1}", ok: "11\n"},
		{script: "BEGIN{a=1;print ++a 1}", ok: "21\n"},
		{script: "BEGIN{a=1/0;print a}", ok: "error:devision by zero\n"},
		{script: "BEGIN{print \"a\" \"b\" \"c\"}", ok: "abc\n"},
		{script: "BEGIN{print 15.2%7.1}", ok: "1\n"},
		{script: "BEGIN{a=123;print a}", ok: "123\n"},
		{script: "BEGIN{a=b=123;print a,b}", ok: "123 123\n"},
		{script: "BEGIN{a,b=123,123;print a,b}", ok: "123 123\n"},
		{script: "BEGIN{map=123;print map}", ok: "123\n"},
		{script: "BEGIN{print \"123\" \"45\"}", ok: "12345\n"},
		{script: "BEGIN{print \"123\" 45}", ok: "12345\n"},
		{script: "BEGIN{print \"123\" 4+5}", ok: "1239\n"},
		{script: "BEGIN{print '1' 0.2}", ok: "10.2\n"},
		{script: "BEGIN{print 123 45}", ok: "12345\n"},
		{script: "BEGIN{print 1.23 4.5}", ok: "5.73\n"},
		{script: "BEGIN{a[1]=1;a[2]=10;print a[1]+a[2]}", ok: "11\n"},
		{script: "BEGIN{a[1]=1;a[2]=10;print a[1]-a[2]}", ok: "-9\n"},
		{script: "BEGIN{a[1]=10;a[2]=5;print a[1]*a[2]}", ok: "50\n"},
		{script: "BEGIN{a[1]=10;a[2]=5;print a[1]/a[2]}", ok: "2\n"},
		{script: "BEGIN{a[1]=1;a[2]=10;print a+a[2]}", ok: "error:can't read value of array\n"},
		{script: "BEGIN{a[1]=1;a[2]=10;print a-a[2]}", ok: "error:can't read value of array\n"},
		{script: "BEGIN{a[1]=10;a[2]=5;print a*a[2]}", ok: "error:can't read value of array\n"},
		{script: "BEGIN{a[1]=10;a[2]=5;print a/a[2]}", ok: "error:can't read value of array\n"},
		{script: "BEGIN{a[1]=1;a[2]=10;print a[1]+a}", ok: "error:can't read value of array\n"},
		{script: "BEGIN{a[1]=1;a[2]=10;print a[1]-a}", ok: "error:can't read value of array\n"},
		{script: "BEGIN{a[1]=10;a[2]=5;print a[1]*a}", ok: "error:can't read value of array\n"},
		{script: "BEGIN{a[1]=10;a[2]=5;print a[1]/a}", ok: "error:can't read value of array\n"},
		{script: "BEGIN{a[1]=1;a[2]=10;print 1+length(a)}", ok: "3\n"},
		{script: "BEGIN{a[1]=1;a[2]=10;print 1+len(a)}", ok: "3\n"},
		{script: "BEGIN{a[1]=1;a[2]=10;print 1-length(a)}", ok: "-1\n"},
		{script: "BEGIN{a[1]=1;a[2]=10;print 1*length(a)}", ok: "2\n"},
		{script: "BEGIN{a[1]=1;a[2]=10;print 1/length(a)}", ok: "0\n"},
		{script: "BEGIN{a[1]=1;print b[3/0]}", ok: "error:devision by zero\n"},
		{script: "BEGIN{a,b=123,123,123;print a,b}", ok: "123 123\n"},
		{script: "BEGIN{m[1]=123;m[2]=456;a,b=m[1],m[2];print a,b}", ok: "123 456\n"},
		{script: "BEGIN{a,b=123;print a,b}", ok: "error:single value assign to multi values\n"},
		{script: "BEGIN{a,b=1/0,2;print a,b}", ok: "error:devision by zero\n"},
		{script: "BEGIN{NF[0],b=1,2;print a,b}", ok: "error:type int does not support index operation\n"},
		// basic error
		{script: "BEGIN{a", ok: "Syntax error: syntax error\n"},
		{script: "BEGIN{a='", ok: "Syntax error: syntax error\n"},
		{script: "BEGIN{a='\n", ok: "Syntax error: syntax error\n"},
		{script: "BEGIN{a=\"\n", ok: "Syntax error: syntax error\n"},
		{script: "BEGIN{a\n=1;print a;", ok: "Syntax error: syntax error\n"},
		// printf
		{script: "BEGIN{a=1;printf \"%d\",a}", ok: "1\n"},
		{script: "BEGIN{printf \"%.2d\",1}", ok: "01\n"},
		{script: "BEGIN{printf \"%.2f\",1.34}", ok: "1.34\n"},
		{script: "BEGIN{printf \"%s\",\"abc\"}", ok: "abc\n"},
		// ternary operator
		{script: "BEGIN{a=1;print a==1?a+1:a+2}", ok: "2\n"},
		{script: "BEGIN{a=2;print a==1?a+1:a+2}", ok: "4\n"},
		{script: "BEGIN{a=\"a\";print a==\"a\"?a+\"1\":a+\"2\"}", ok: "a1\n"},
		{script: "BEGIN{a=\"b\";print a==\"a\"?a+\"1\":a+\"2\"}", ok: "b2\n"},
		{script: "BEGIN{print 1/0?a+1:a+2}", ok: "error:devision by zero\n"},
		{script: "BEGIN{a[1]=1;print a?a+1:a+2}", ok: "error:convert to bool failed in ternary operator\n"},
		// composite expression
		{script: "BEGIN{a=123;print ++a;print a}", ok: "124\n124\n"},
		{script: "BEGIN{print ++1}", ok: "error:invalid operation\n"},
		{script: "BEGIN{a=123;print a++;print a}", ok: "123\n124\n"},
		{script: "BEGIN{a=123;print --a;print a}", ok: "122\n122\n"},
		{script: "BEGIN{a=123;print a--;print a}", ok: "123\n122\n"},
		{script: "BEGIN{a[1]=123;print ++a[1];print a[1]}", ok: "124\n124\n"},
		{script: "BEGIN{a[1]=123;print a[1]++;print a[1]}", ok: "123\n124\n"},
		{script: "BEGIN{a[1]=123;print --a[1];print a[1]}", ok: "122\n122\n"},
		{script: "BEGIN{a[1]=123;print a[1]--;print a[1]}", ok: "123\n122\n"},
		{script: "BEGIN{a=123.4;++a;print a}", ok: "124.4\n"},
		{script: "BEGIN{a=123;--a;print a}", ok: "122\n"},
		{script: "BEGIN{a=123.4;--a;print a}", ok: "122.4\n"},
		{script: "BEGIN{a=123;a+=4;print a}", ok: "127\n"},
		{script: "BEGIN{a+=4;print a}", ok: "4\n"},
		{script: "BEGIN{a=123;a-=4;print a}", ok: "119\n"},
		{script: "BEGIN{a=123;a*=4;print a}", ok: "492\n"},
		{script: "BEGIN{a=123;a/=4;print a}", ok: "30\n"},
		{script: "BEGIN{a=123;a%=4;print a}", ok: "3\n"},
		{script: "BEGIN{a+=1/0}", ok: "error:devision by zero\n"},

		// Comment
		{script: `BEGIN{ /*a=100;*/ a= a+100;print a; }`, ok: "100\n"},
		{script: `BEGIN{ /*a=100;a= a+100;print a; 
		*}`, ok: "Syntax error: syntax error\n"},
		{script: `BEGIN{ #a=100;
		a= a+100;print a; }`, ok: "100\n"},
		{script: `BEGIN{ #a=100; }`, ok: "Syntax error: syntax error\n"},

		// JAPANESE
		{script: "BEGIN{print \"あいう\"}", ok: "あいう\n"},
		{script: "BEGIN{a=\"あいう\";b=\"えお\";print a+b}", ok: "あいうえお\n"},
		{script: "BEGIN{a[\"あ\"]=1;a[\"い\"]=2;a[\"う\"]=3;for(key in a){print key,a[key]}}", ok: "あ 1\nい 2\nう 3\n"},

		// variable and scope
		// builtin
		{script: "BEGIN{++NF;print NF}", ok: "1\n"},
		{script: "BEGIN{++NF}END{print NF}", ok: "1\n"},
		{script: "BEGIN{NF=1}END{print NF}", ok: "1\n"},
		{script: "BEGIN{NF=1.1}END{print NF}", ok: "error:type of NF must be int ,not float64.\n"},
		{script: "BEGIN{NF=\"aaa\"}", ok: "error:type of NF must be int ,not string.\n"},
		{script: "BEGIN{$0=\"aaa\";print}", ok: "aaa\n"},
		{script: "BEGIN{$1=\"aaa\";print}", ok: "aaa\n"},
		{script: "BEGIN{print FS}", ok: "\n"},
		{script: "BEGIN{FS=\"X\"}END{print FS}", ok: "X\n"},
		{script: "BEGIN{FS=123}END{print FS}", ok: "error:type of FS must be string ,not int.\n"},
		// global
		{script: "BEGIN{A=1}END{print A}", ok: "1\n"},
		{script: "BEGIN{Abc=1}END{print Abc}", ok: "1\n"},
		{script: "BEGIN{ABC=1}END{print ABC}", ok: "1\n"},
		{script: "BEGIN{++A;print A}", ok: "1\n"},
		{script: "BEGIN{++A}END{print A}", ok: "1\n"},
		{script: "BEGIN{A=1.1;print A}", ok: "1.1\n"},
		{script: "BEGIN{A=1.1}END{print A}", ok: "1.1\n"},
		{script: "BEGIN{A=\"AAA\";print A}", ok: "AAA\n"},
		{script: "BEGIN{A=\"AAA\"}END{print A}", ok: "AAA\n"},
		// local
		{script: "BEGIN{l=1}END{print l}", ok: "\n"},
		{script: "BEGIN{lmn=1}END{print lmn}", ok: "\n"},
		{script: "BEGIN{lMN=1}END{print lMN}", ok: "\n"},
		{script: "BEGIN{++l;print l}", ok: "1\n"},
		{script: "BEGIN{++l}END{print l}", ok: "\n"},
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
		{script: "BEGIN{a=1;b=2;print a&&b}", ok: "true\n"},
		{script: "BEGIN{a=1;b=2;print a&b}", ok: "Syntax error: syntax error\n"},
		{script: "BEGIN{a=1;b=2;print a==1||b==2}", ok: "true\n"},
		{script: "BEGIN{a=1;b=2;print a==2||b==2}", ok: "true\n"},
		{script: "BEGIN{a=1;b=2;print a||b}", ok: "true\n"},
		{script: "BEGIN{a=1;b=2;print a|b}", ok: "Syntax error: syntax error\n"},
		{script: "BEGIN{print 1||1}", ok: "true\n"},
		{script: "BEGIN{print 0||0}", ok: "false\n"},
		{script: "BEGIN{print 1/0||1}", ok: "error:devision by zero\n"},
		{script: "BEGIN{print 1||1/0}", ok: "error:devision by zero\n"},
		{script: "BEGIN{a[1]=1;print a||0}", ok: "error:convert to bool failed in left expression of OR operator\n"},
		{script: "BEGIN{a[1]=1;print 0||a}", ok: "error:convert to bool failed in right expression of OR operator\n"},
		{script: "BEGIN{print 0&&1}", ok: "false\n"},
		{script: "BEGIN{print 1&&0}", ok: "false\n"},
		{script: "BEGIN{print 1/0&&1}", ok: "error:devision by zero\n"},
		{script: "BEGIN{print 1&&1/0}", ok: "error:devision by zero\n"},
		{script: "BEGIN{a[1]=1;print a&&1}", ok: "error:convert to bool failed in left expression of AND operator\n"},
		{script: "BEGIN{a[1]=1;print 1&&a}", ok: "error:convert to bool failed in right expression of AND operator\n"},

		// regular expression
		{script: "BEGIN{print \"aaa\"~/aaa/}", ok: "true\n"},
		{script: "BEGIN{print \"aaa\"!~/aaa/}", ok: "false\n"},
		{script: "BEGIN{print \"aaa\"~/abc/}", ok: "false\n"},
		{script: "BEGIN{print \"aaa\"~/a+/}", ok: "true\n"},
		{script: "BEGIN{print \"aaa\"~/^a+$/}", ok: "true\n"},
		{script: "BEGIN{print \"abc\"~/^a+$/}", ok: "false\n"},
		{script: "\"AAA\"~/AAA/{print}", in: "AAA", ok: "AAA\n"},
		{script: "$0~/AAA/{print}", in: "AAA", ok: "AAA\n"},
		{script: "/AAA/{print}", in: "AAA", ok: "AAA\n"},
		{script: "BEGIN{S=\"abcaaa\";gsub(/a+/,\"A\",S);print S}", ok: "AbcA\n"},
		{script: "BEGIN{S=\"abcaaa\";print gsub(/a+/,\"A\",S);print S}", ok: "2\nAbcA\n"},
		{script: "BEGIN{print 1/0~/aaa/}", ok: "error:devision by zero\n"},

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
		{script: "BEGIN{a=10;print ++a}", ok: "11\n"},
		{script: "BEGIN{a=1.9;print ++a}", ok: "2.9\n"},
		{script: "BEGIN{print ++10}", ok: "error:invalid operation\n"},
		{script: "BEGIN{a=\"a\";print ++a}", ok: "1\n"},
		{script: "BEGIN{a=10;print --a}", ok: "9\n"},
		{script: "BEGIN{a=2.9;print --a}", ok: "1.9\n"},
		{script: "BEGIN{print --10}", ok: "error:invalid operation\n"},
		{script: "BEGIN{a=\"a\";print --a}", ok: "-1\n"},
		{script: "BEGIN{a=10;a+=2;print a}", ok: "12\n"},
		{script: "BEGIN{a=10;a+=2.5;print a}", ok: "12.5\n"},
		{script: "BEGIN{a=10;a+=2;print a}", ok: "12\n"},
		{script: "BEGIN{a=10;a-=2;print a}", ok: "8\n"},
		{script: "BEGIN{a=10;a*=2;print a}", ok: "20\n"},
		{script: "BEGIN{a=10;a/=2;print a}", ok: "5\n"},
		{script: "BEGIN{a=\"abc\";a+=\"xyz\";print a}", ok: "abcxyz\n"},
		// TODO & TOFIX
		//{script: "BEGIN{a=\"abc\";a-=\"xyz\";print a}", ok: "abcxyz\n"},
		//{script: "BEGIN{a=\"1\";a+=\"2\";print a}", ok: "3\n"}, //TOFIX
		//{script: "BEGIN{a=\"abc\";a-=\"xyz\";print a}", ok: "abcxyz\n"},

		// multi expressions
		{script: "BEGIN{a,b=1,2;print a}", ok: "1\n"},
		{script: "BEGIN{a,b=1,2;print b}", ok: "2\n"},
		{script: "BEGIN{a,b=1,2,3;print b}", ok: "2\n"},
		{script: "BEGIN{a,b,c=1,2;print b}", ok: "2\n"},

		// {script: "BEGIN{t=func(){return 1,2}();print t}", ok: "1 2\n"}, //TODO: SHOULD BE ERROR??

		// if statement
		{script: "BEGIN{a=1;if a==1 { a=2 ;print a;}}", ok: "2\n"},
		{script: "BEGIN{a=1;if 100 { a=2 ;print a;}}", ok: "2\n"},
		{script: "BEGIN{a=1;if 1.23 { a=2 ;print a;}}", ok: "2\n"},
		{script: "BEGIN{a=1;if 0 { a=2 ;print a;}}", ok: ""},
		{script: "BEGIN{a=1;if \"a\" { a=2 ;print a;}}", ok: "2\n"},
		{script: "BEGIN{a=1;if \"\" { a=2 ;print a;}}", ok: ""},
		{script: "BEGIN{a=1;if a==1 { a=2 };print a}", ok: "2\n"},
		{script: "BEGIN{a=1;if a==1 { env_test=2 };print env_test}", ok: "\n"},
		{script: "BEGIN{a=2;if a==1 { a=2 } else { a=3;b=4;print b }}", ok: "4\n"},
		{script: "BEGIN{a=1;b=1;if a==1 { b=2 };print b}", ok: "2\n"},
		{script: "BEGIN{a=2;if a==1 { a=11 } else if a==2 { a=12 } ;print a}", ok: "12\n"},
		{script: "BEGIN{a=2;if a==1 { a=11 } else if 1 { a=12 } ;print a}", ok: "12\n"},
		{script: "BEGIN{a=2;if a==1 { a=11 } else if 0 { a=12 } ;print a}", ok: "2\n"},
		{script: "BEGIN{a=1;if a==1 { a=11 } else if a==2 { a=12 } else if a==3 { a=13 };print a}", ok: "11\n"},
		{script: "BEGIN{a=2;if a==1 { a=11 } else if a==2 { a=12 } else if a==3 { a=13 };print a}", ok: "12\n"},
		{script: "BEGIN{a=3;if a==1 { a=11 } else if a==2 { a=12 } else if a==3 { a=13 };print a}", ok: "13\n"},
		{script: "BEGIN{a=1;if a==1 { env_test=11 } else if a==2 { env_test=12 } else { env_test=13 };print env_test}", ok: "\n"},
		{script: "BEGIN{a=2;if a==1 { env_test=11 } else if a==2 { env_test=12 } else { env_test=13 };print env_test}", ok: "\n"},
		{script: "BEGIN{a=3;if a==1 { env_test=11 } else if a==2 { env_test=12 } else { env_test=13 };print env_test}", ok: "\n"},
		{script: "BEGIN{if a==1/0 { print a;}}", ok: "error:devision by zero\n"},
		{script: "BEGIN{a[1]=1;if a { print a;}}", ok: "error:convert to bool failed in if condition\n"},
		{script: "BEGIN{a=1;if a==0 { print a}else if a/0 { print a}}", ok: "error:devision by zero\n"},
		{script: "BEGIN{a[1]=1;if a[1]==0 { print a}else if a { print a}}", ok: "error:convert to bool failed in else if condition\n"},
		{script: "BEGIN{a=1;if a==0 { print a}else if a==1 { print a/0}}", ok: "error:devision by zero\n"},

		// for statement
		{script: "BEGIN{a=0;for{ if a==10 { break }; a= a+1 };print a}", ok: "10\n"},
		{script: "BEGIN{a=0;b=0;for{ a=a+1;if a==10 { break }; if b==5 {continue};b= b+1 };print b}", ok: "5\n"},
		{script: "BEGIN{a=0;for a<=10 { a= a+1 };print a}", ok: "11\n"},
		{script: "BEGIN{a=0;for a { a= a+1 };print a}", ok: "0\n"},
		{script: "BEGIN{a=1;for a { a= a-1 };print a}", ok: "0\n"},
		{script: "BEGIN{s[1]=1;for s { a= a-1 };print a}", ok: "error:convert to bool failed in while condition\n"},
		{script: "BEGIN{s=\"\";for s { s= s+1 };print s}", ok: "\n"},
		{script: "BEGIN{s=\"str\";for s { s= \"\" };print s}", ok: "\n"},
		// while statement == for statement
		{script: "BEGIN{a=0;b=0;while{ a=a+1;if a==10 { break }; if b==5 {continue};b= b+1 };print b}", ok: "5\n"},
		{script: "BEGIN{a=0;while a<=10 { a= a+1 };print a}", ok: "11\n"},
		{script: "BEGIN{a=0;while a { a= a+1 };print a}", ok: "0\n"},
		{script: "BEGIN{a=1;while a { a= a-1 };print a}", ok: "0\n"},
		{script: "BEGIN{s=\"\";while s { s= s+1 };print s}", ok: "\n"},
		{script: "BEGIN{s=\"str\";while s { s= \"\" };print s}", ok: "\n"},
		{script: "BEGIN{a=1;while a/0 { a= a+1 };print a}", ok: "error:devision by zero\n"},
		{script: "BEGIN{fnc=func(){a=1;while a { return a };};print fnc()}", ok: "1\n"},
		// for;;{}
		{script: "BEGIN{for i=1;i<=3;++i{print i}}", ok: "1\n2\n3\n"},
		{script: "BEGIN{for 1;i<=3;++i{print i}}", ok: "\n1\n2\n3\n"},
		{script: "BEGIN{for 1;1;++i{print i;if i==3{break}}}", ok: "\n1\n2\n3\n"},
		{script: "BEGIN{for ;1;++i{print i;if i==3{break}}}", ok: "\n1\n2\n3\n"},
		{script: "BEGIN{for 1;;++i{print i;if i==3{break}}}", ok: "\n1\n2\n3\n"},
		{script: "BEGIN{for ;;++i{print i;if i==3{break}}}", ok: "\n1\n2\n3\n"},
		{script: "BEGIN{for ;;{print i;if i==3{break};i++}}", ok: "\n1\n2\n3\n"},
		{script: "BEGIN{for i=1;i<=3;++i{if i<3{continue};print i}}", ok: "3\n"},
		{script: "BEGIN{fnc=func(){for i=1;i<=3;++i{return i}};print fnc()}", ok: "1\n"},
		{script: "BEGIN{for i=1/0;i<=3;++i{print i}}", ok: "error:devision by zero\n"},
		{script: "BEGIN{for i=1;i<=3/0;++i{print i}}", ok: "error:devision by zero\n"},
		{script: "BEGIN{for i=1;i<=3;i/0{i++}}", ok: "error:devision by zero\n"},
		{script: "BEGIN{for i=1;i<=3;++i{i/0}}", ok: "error:devision by zero\n"},
		{script: "BEGIN{a[1]=1;for i=1;a;++i{print i}}", ok: "error:convert to bool failed in for loop condition\n"},
		// do while statement
		{script: "BEGIN{a=0;do{a=a+1} while(a<10);print a}", ok: "10\n"},
		{script: "BEGIN{a=0;do{a=a+1;if a==5{break}} while(a<10);print a}", ok: "5\n"},
		{script: "BEGIN{a=-10;do{a=a+1} while(a);print a}", ok: "0\n"},
		{script: "BEGIN{a[1]=1;do{a[1]=a[1]+1} while(a);print a[1]}", ok: "error:convert to bool failed in do loop condition\n"},
		{script: "BEGIN{a=0;do{a=a+1;if a<10 {continue}else{break}} while(1);print a}", ok: "10\n"},
		{script: "BEGIN{a=0;do{a=a/0} while(a<10);print a}", ok: "error:devision by zero\n"},
		{script: "BEGIN{a=0;do{a=a+1} while(a<10/0);print a}", ok: "error:devision by zero\n"},
		{script: "BEGIN{fnc=func(){a=0;do{return a} while(a<10)};print fnc()}", ok: "0\n"},

		// map: awk-array (associated array = map)
		{script: "BEGIN{print a[1]}", ok: "\n"},
		{script: "BEGIN{print a[1/0]}", ok: "error:devision by zero\n"},
		{script: "BEGIN{print NF[1]}", ok: "error:type int does not support index operation\n"},
		{script: "BEGIN{a[1/0]=1}", ok: "error:devision by zero\n"},
		{script: "BEGIN{NF[1]=1}", ok: "error:type int does not support index operation\n"},
		{script: "BEGIN{a[1]=1;print a[1]}", ok: "1\n"},
		{script: "BEGIN{a[1]=1;print a[2]}", ok: "\n"},
		{script: "BEGIN{a[1]=1;a[2]=2;print a[1]+a[2]}", ok: "3\n"},
		{script: "BEGIN{a[1]=1;a[2]=2;print a+a}", ok: "error:can't read value of array\n"},
		{script: "BEGIN{a[1]=1;print 1 in a}", ok: "true\n"},
		{script: "BEGIN{a[1]=1;print 2 in a}", ok: "false\n"},
		{script: "BEGIN{print 1 in b}", ok: "false\n"},
		{script: "BEGIN{print 1/0 in b}", ok: "error:devision by zero\n"},
		{script: "BEGIN{a[1]=1;print \"1\" in a}", ok: "true\n"},
		{script: "BEGIN{a[1]=1;print \"2\" in a}", ok: "false\n"},
		{script: "BEGIN{a[\"1\"]=1;print 1 in a}", ok: "true\n"},
		{script: "BEGIN{a[\"1\"]=1;print 2 in a}", ok: "false\n"},
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
		{script: "BEGIN{a[1]=1;++a[1];print a[1]}", ok: "2\n"},
		{script: "BEGIN{a[1]=1;a[1]--;print a[1]}", ok: "0\n"},
		{script: "BEGIN{a[1]=1;a[2]=2;a[3]=3;b=a;b[1]=10;print a[1]}", ok: "10\n"}, // THIS SPEC OK?
		{script: "BEGIN{a[1]=1;delete a[1];print a[1]}", ok: "\n"},
		{script: "BEGIN{a[1]=1;delete a[1];print a[2]}", ok: "\n"},
		{script: "BEGIN{a[1]=1;delete a[2];print a[1]}", ok: "1\n"},
		{script: "BEGIN{a[1]=1;delete a;print a[1]}", ok: "\n"},
		{script: "BEGIN{a[1]=1;delete a[1/0]}", ok: "error:devision by zero\n"},
		{script: "BEGIN{a=1;delete a}", ok: "error:type int does not support delete operation\n"},
		{script: "BEGIN{delete 1}", ok: "error:type *ast.NumExpr does not support delete operation\n"},
		{script: "BEGIN{a[1]=1;delete a;a=2}", ok: "error:can't assign to a; it's an associated array name.\n"},
		{script: "BEGIN{delete a;a=2}", ok: "error:can't assign to a; it's an associated array name.\n"},
		{script: "BEGIN{list=func(){a[1]=1;a[2]=2;a[3]=3;return a};delete list()[1]}", ok: "error:non variable does not support delete operation\n"},
		{script: "BEGIN{list=func(){a[1]=1;a[2]=2;a[3]=3;return a};list()[1]=3}", ok: "error:invalid assignment\n"},
		// for ( key in map )
		{script: "BEGIN{for (i in a) {print i,a[i]}}", ok: ""},
		{script: "BEGIN{a[1]=1;a[2]=2;for (i in a) {print i,a[i]}}", ok: "1 1\n2 2\n"},
		{script: "BEGIN{a[1]=1;a[2]=2;for (i in a) {};print i}", ok: "\n"},
		{script: "BEGIN{a[1]=1;a[2]=2;i=0;for (i in a) {};print i}", ok: "2\n"},
		{script: "BEGIN{a=1;for (i in a) {print i}}", ok: "error:for key loop not in associated array,int\n"},
		{script: "BEGIN{a[\"1\"]=1;a[\"2\"]=2;for (i in a) {print i,a[i]}}", ok: "1 1\n2 2\n"},
		{script: "BEGIN{a[1]=1;a[2]=2;for (i in a) {print i,a[i]}}", ok: "1 1\n2 2\n"},
		{script: "BEGIN{a[1]++;a[2]=2;for (i in a) {print i,a[i]}}", ok: "1 1\n2 2\n"},
		{script: "BEGIN{a[1]=1;a[2]=2;a[3]=3;for (i in a) {print i;if i==\"2\" { break }}}", ok: "1\n2\n"},
		{script: "{A[$0]++} END{for(key in A){print key}}", in: "AAA", ok: "AAA\n"},
		{script: "BEGIN{fnc=func(){a[1]=1;a[2]=2;for (i in a) {if i=='2'{return i}else {continue}}};print fnc()}", ok: "2\n"},
		{script: "BEGIN{a[1]=1;for (i in a) {print i/0}}", ok: "error:devision by zero\n"},

		// function
		{script: "BEGIN{add=func(a,b){return a+b}; x=add(10,5);print x}", ok: "15\n"},
		{script: "BEGIN{x=add(10,5);print x}", ok: "error:unknown symbol\n"},
		{script: "BEGIN{add=123;x=add(10,5);print x}", ok: "error:cannot call type int\n"},
		{script: "BEGIN{add=func(a,b){return a+b}; x=add(10/0,5);print x}", ok: "error:devision by zero\n"},
		{script: "BEGIN{A[1]=10;add=func(a,b){return a+b}; x=add(A,5);print x}", ok: "error:can't read value of array\n"},
		{script: "BEGIN{add=func(a,b){return a+b}; print add(10,5)}", ok: "15\n"},
		{script: "BEGIN{add=func(a,b){return a+b}; print add(1.1,2.1)}", ok: "3.2\n"},
		{script: "BEGIN{add=func(a,b){return a+b}; print add(\"あ\",\"いう\")}", ok: "あいう\n"},
		{script: "BEGIN{print func(a,b){return a+b}(10,5)}", ok: "15\n"},
		{script: "BEGIN{a=123;add=func(a,b){return a+b}; add(10,5);print a}", ok: "123\n"},
		{script: "BEGIN{c=100;add=func(a,b){return a+b+c}; print add(10,5)}", ok: "115\n"},
		{script: "BEGIN{one=func(){return 1}; print one()}", ok: "1\n"},
		{script: "BEGIN{a=10;plusone=func(){a++};plusone();print a}", ok: "11\n"},
		{script: "BEGIN{print func(){return 1}()}", ok: "1\n"},
		{script: "BEGIN{a=10;plusone=func(){a++;return};plusone();print a}", ok: "11\n"},
		{script: "BEGIN{hash=func(){m[1]=1;m[2]=2;m[3]=3;return m}; m=hash();print m[1]}", ok: "1\n"},
		{script: "BEGIN{map=func(){m[1]=1;m[2]=2;m[3]=3;return m}; print map()[1]}", ok: "1\n"},
		{script: "BEGIN{print func(){m[1]=1;m[2]=2;m[3]=3;return m}()[1]}", ok: "1\n"},
		{script: "BEGIN{err=func(a){return a/0}; x=err(10);print x}", ok: "error:devision by zero\n"},
		// call go func with variadic args
		{script: "BEGIN{println(\"abc\",\"def\")}", ok: "abc def\n"},
		{script: "BEGIN{println()}", ok: "\n"},
		{script: "BEGIN{printf(\"hello,\\t%s\\n\",\"world!\")}", ok: "hello,\tworld!\n"}, //TOFIX
		{script: "BEGIN{print sum(1,2,3)}", ok: "6\n"},
		{script: "BEGIN{print sum()}", ok: "0\n"},
		{script: "BEGIN{print cat(\"abc\",\"def\")}", ok: "abcdef\n"},
		{script: "BEGIN{print cat()}", ok: "\n"},

		// multi result function
		{script: "BEGIN{Cross=func(a1,a2){return a2,a1;};print Cross(1,5)}", ok: "5 1\n"},
		{script: "BEGIN{Cross=func(a1,a2){return a2,a1;};x,y=Cross(1,5);print x}", ok: "5\n"},
		{script: "BEGIN{Cross=func(a1,a2){return a2,a1;};x,y=Cross(1,5);print y}", ok: "1\n"},
		{script: "BEGIN{First=func(a1,a2){return a1;};x,y=First(1,5);print x}", ok: "error:single value assign to multi values\n"},
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
		{script: "function one(){return 1}BEGIN{print one()}", ok: "1\n"},
		{script: "func one(){return 1}BEGIN{print one()}", ok: "1\n"},
		{script: "func printOne(){print 1}BEGIN{printOne()}", ok: "1\n"},

		// command parameter

		// built in variables
		{script: "BEGIN{FS=\":\"}{print $2}", in: "AAA:BBB:CCC\nAAA:BBB:CCC\n", ok: "BBB\nBBB\n"},
		{script: "{print}", in: "AAA BBB  CCC\n", ok: "AAA BBB  CCC\n"}, // AWK is AWK
		{script: "BEGIN{OFS=\":\"}{$1=$1;print}", in: "AAA BBB  CCC\n", ok: "AAA:BBB:CCC\n"},
		//{script: "BEGIN{OFS=\"\n\"}{$1=$1;print}", in: "AAA BBB CCC\nAAA BBB CCC\n", ok: "AAA\nBBB\nCCC\nAAA\nBBB\nCCC\n"}, //TODO
		{script: "BEGIN{ORS=\":\"}{$1=$1;print}", in: "AAA BBB CCC\nCCC DDD EEE\n", ok: "AAA BBB CCC:CCC DDD EEE:\n"},

		// Numeric Built-in Functions
		// lib: int
		{script: "BEGIN{print int(123)}", ok: "123\n"},
		{script: "BEGIN{print int(\"123\")}", ok: "123\n"},
		{script: "BEGIN{print int(123.45)}", ok: "123\n"},
		// lib: sqrt,exp,log
		{script: "BEGIN{print sqrt(4)}", ok: "2\n"},
		{script: "BEGIN{print sqrt(3)}", ok: "1.7320508075688772\n"},
		{script: "BEGIN{print exp(2)}", ok: "7.38905609893065\n"},
		{script: "BEGIN{print exp(0)}", ok: "1\n"},
		{script: "BEGIN{print exp(1)}", ok: "2.718281828459045\n"},
		// lib: sin,cos,atan2
		{script: "BEGIN{print sin(0)}", ok: "0\n"},
		{script: "BEGIN{print sin(3.141592650358979/2)}", ok: "1\n"},
		{script: "BEGIN{print cos(0)}", ok: "1\n"},
		{script: "BEGIN{print int(cos(3.141592650358979/3)*200)}", ok: "100\n"},
		{script: "BEGIN{print atan2(0,0)}", ok: "0\n"},
		{script: "BEGIN{print atan2(1,0)}", ok: "1.5707963267948966\n"},
		// lib: rand,arand,srand
		{script: "BEGIN{r1,r2=rand(),rand(); print r1==r2}", ok: "false\n"},
		{script: "BEGIN{srand();r1,r2=rand(),rand(); print r1==r2}", ok: "false\n"},

		// Built-in Functions for String Manipulation
		// lib:len
		{script: "BEGIN{print length(\"Hello World!\")}", ok: "12\n"},
		{script: "BEGIN{print len(\"Hello World!\")}", ok: "12\n"},
		{script: "BEGIN{a[1]=1;a[2]=2;print len(a)}", ok: "2\n"},
		{script: "BEGIN{f=func(){return 1,2};print len(f())}", ok: "2\n"},
		{script: "BEGIN{print length(123)}", ok: "3\n"},
		{script: "{print length()}", in: "Hello!", ok: "6\n"},
		// lib:sprintf
		{script: "BEGIN{print sprintf(\"%d:%s\",100,\"abc\")}", ok: "100:abc\n"},
		// lib:substr
		{script: "BEGIN{print substr(\"abcde\",1,3)}", ok: "abc\n"},
		{script: "BEGIN{print substr(\"abcde\",0,3)}", ok: "abc\n"},
		{script: "BEGIN{print substr(\"abcde\",-1,3)}", ok: "abc\n"},
		{script: "BEGIN{print substr(\"abcde\",1,5)}", ok: "abcde\n"},
		{script: "BEGIN{print substr(\"abcde\",1,6)}", ok: "abcde\n"},
		{script: "BEGIN{print substr(\"abcde\",3,2)}", ok: "cd\n"},
		{script: "BEGIN{print substr(\"abcde\",3)}", ok: "cde\n"},
		{script: "BEGIN{print substr(\"abcde\",2,0)}", ok: "\n"},
		{script: "BEGIN{print substr(\"abcde\",2,-1)}", ok: "\n"},
		{script: "BEGIN{print substr(12345,1,3)}", ok: "123\n"},
		{script: "BEGIN{print substr(12.345,1,4)}", ok: "12.3\n"},
		{script: "BEGIN{print substr(\"\",1,3)}", ok: "\n"},
		// lib:match
		{script: "BEGIN{print match(\"abcdde\",/cd+/);print RSTART,RLENGTH}", ok: "3\n3 3\n"},
		{script: "BEGIN{print match(\"abcde\",/dc+/);print RSTART,RLENGTH}", ok: "0\n0 -1\n"},
		{script: "BEGIN{print match(\"abcde\",/cd+/)}", ok: "3\n"},
		{script: "BEGIN{print match(\"abcde\",/dc+/)}", ok: "0\n"},
		// lib:split
		{script: "BEGIN{ar[1]=\"\";print split(\"a:b:c\",ar,\":\");print ar[1]}", ok: "3\na\n"},
		{script: "BEGIN{ar[1]=\"\";print split(\"a:b:c\",ar);print ar[1]}", ok: "1\na:b:c\n"},
		{script: "BEGIN{split(\"a:b:c\",ar,\":\");print ar[3]}", ok: "c\n"},
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
		// lib: sub,gsub
		{script: "BEGIN{print gsub(/a/,\"A\",S)}", ok: "0\n"},
		{script: "BEGIN{S=\"aabbaacc\";print sub(/a/,\"A\",S);print S}", ok: "1\nAabbaacc\n"},
		{script: "BEGIN{S=\"aabbaacc\";print gsub(/a/,\"A\",S);print S}", ok: "4\nAAbbAAcc\n"},
		{script: "BEGIN{S=\"aabbaacc\";print sub(/a+/,\"A\",S);print S}", ok: "1\nAbbaacc\n"},
		{script: "BEGIN{S=\"aabbaacc\";print gsub(/^a+/,\"A\",S);print S}", ok: "1\nAbbaacc\n"},
		{script: "BEGIN{S=\"aabbaacc\";print sub(/^a+/,\"\",S);print S}", ok: "1\nbbaacc\n"},
		{script: "BEGIN{S=\"aabbaacc\";print sub(/c+$/,\"\",S);print S}", ok: "1\naabbaa\n"},
		{script: "{print sub(/a+/,\"\");print }", in: "aabbaacc", ok: "1\nbbaacc\n"},
		{script: "{print gsub(/a+/,\"\");print }", in: "aabbaacc", ok: "2\nbbcc\n"},
		// lib: mktime,strftime
		//{script: "BEGIN{print mktime(\"2018 09 01 13 07 50\")}", ok: "1535774870\n"},  //JAPAN
		{script: "BEGIN{print mktime(\"2018 99 01 13 07 50\")}", ok: "0\n"}, //JAPAN
		{script: "BEGIN{print strftime(\"%Y/%m/%d %H:%M:%S\",mktime(\"2018 09 01 13 07 50\"))}", ok: "2018/09/01 13:07:50\n"},
		{script: "BEGIN{print strftime(\"%y-%m-%d %H:%M:%S\",mktime(\"2018 09 01 13 07 50\"))}", ok: "18-09-01 13:07:50\n"},
		// lib: systime
		{script: "BEGIN{systime()}", ok: ""},
		// lib: system
		{script: "BEGIN{print system(\"echo aaa\")}", ok: "0\n"},

		// field
		{script: "{print $1}", in: "Hello World!\n", ok: "Hello\n"},
		{script: "{print $(1/1)}", in: "Hello World!\n", ok: "Hello\n"},
		{script: "{print $(1/0)}", in: "Hello World!\n", ok: "error:devision by zero\n"},
		{script: "{b=$1;print b}", in: "Hello World!\n", ok: "Hello\n"},
		{script: "{$1=2;b=$1;print b}", in: "Hello World!\n", ok: "2\n"},
		{script: "{print $'a'}", in: "Hello World!\n", ok: "error:field index not int :string\n"},
		{script: "{print $''}", in: "Hello World!\n", ok: "error:field index not int :string\n"},
		//{script: "{print $'1'}", in: "Hello World!\n", ok: "Hello\n"}, //TODO
		//{script: "{print $'1.1'}", in: "Hello World!\n", ok: "Hello\n"}, //TODO
		//{script: "{print $'1.xx'}", in: "Hello World!\n", ok: "Hello\n"}, //TODO
		{script: "{a=1.1;$a=1;print $a}", in: "Hello World!\n", ok: "error:field index not int :float64\n"},
		{script: "{$(1/0)=1;print $a}", in: "Hello World!\n", ok: "error:devision by zero\n"},
		{script: "{$(-1)='xx';}", in: "Hello World!\n", ok: "error:Field Index Out of Range:-1\n"},
		{script: "{a[1]=2;$1=a;print $1}", in: "Hello World!\n", ok: "error:field value is not string :map[interface {}]interface {}\n"},
		{script: "{print NF}", in: "\n \n\t\naaa\n", ok: "0\n0\n0\n1\n"},
		{script: "BEGIN{FS=\":\"}{print NF}", in: "\n:\naaa:bbb\n", ok: "0\n2\n2\n"},
		{script: "BEGIN{FS=\"\"}{print NF}", in: "aaa\n", ok: "3\n"},
		{script: "BEGIN{FS=\"\"}{print NF}", in: "あああ\n", ok: "9\n"}, //TODO: Diff from awk?
		{script: "{print length($1)*1}", in: "Hello World!\n", ok: "5\n"},
		{script: "$1==\"AAA\"{print;COUNT++} END{print COUNT}", in: "AAA BBB CCC\nAAA BBB CCC\n", ok: "AAA BBB CCC\nAAA BBB CCC\n2\n"},
		{script: "NR==1{$2=$1 ;print $0,NF} NR==2{$5=$1; print $0,NF}", in: "AAA BBB CCC\nAAA BBB CCC\n", ok: "AAA AAA CCC 3\nAAA BBB CCC  AAA 5\n"},

		//patterns
		{script: "BEGIN{print 1}BEGIN{print 2}", ok: "1\n2\n"},
		{script: "BEGIN{print 1}END{print 2}", ok: "1\n2\n"},
		{script: "1/0{print 1.5}", in: "AAA\n", ok: "error:devision by zero\n"},
		// /start/./stop/
		{script: "/AAA/,/CCC/", in: "AAA\nBBB\nCCC\nDDD\n", ok: "AAA\nBBB\nCCC\n"},
		{script: "/AAA/,/CCC/{print}", in: "AAA\nBBB\nCCC\nDDD\n", ok: "AAA\nBBB\nCCC\n"},
		{script: "/BBB/,/CCC/{print}", in: "AAA\nBBB\nCCC\nDDD\n", ok: "BBB\nCCC\n"},
		{script: "/BBB/,/BBB/{print}", in: "AAA\nBBB\nCCC\nDDD\n", ok: "BBB\n"},
		{script: "/ZZZ/,/BBB/{print}", in: "AAA\nBBB\nCCC\nDDD\n", ok: ""},
		{script: "/A+/,/A+/{print}", in: "AAX\nBBB\nAAZ\nDDD\n", ok: "AAX\nAAZ\n"},
		{script: "//,/A+/{print}", in: "AAA\nBBB\nAAA\nDDD\n", ok: "AAA\nBBB\nAAA\nDDD\n"},
		{script: "/A+/,//{print}", in: "AAA\nBBB\nAAA\nDDD\n", ok: "AAA\nAAA\n"},

		// next
		{script: "/BBB/,/CCC/{next}1", in: "AAA\nBBB\nCCC\nDDD\n", ok: "AAA\nDDD\n"},
		{script: "{if NR%2==0{next}}1", in: "AAA\nBBB\nCCC\nDDD\n", ok: "AAA\nCCC\n"},
		{script: "{for{if NR%2==0 {next};break;}}1", in: "AAA\nBBB\nCCC\nDDD\n", ok: "AAA\nCCC\n"},
		{script: "function skipper(){if NR%2==0 {next};return}{skipper()}1", in: "AAA\nBBB\nCCC\nDDD\n", ok: "AAA\nCCC\n"},

		// exit (no check return code)
		{script: "BEGIN{exit 0}1", in: "AAA\nBBB\nCCC\nDDD\n", ok: ""},
		{script: "BEGIN{exit 0/0}1", in: "\n", ok: "error:devision by zero\n"},
		{script: "NR==3{exit 0}1", in: "AAA\nBBB\nCCC\nDDD\n", ok: "AAA\nBBB\n"},
		{script: "NR==3{exit 1+1}1", in: "AAA\nBBB\nCCC\nDDD\n", ok: "AAA\nBBB\n", rc: 2},
		{script: "{if $0==\"BBB\" {exit 1}}1", in: "AAA\nBBB\nCCC\nDDD\n", ok: "AAA\n", rc: 1},

		// getline
		{script: `
		BEGIN{
			getline
			print "BEGIN",$0
		}
		{
			print "MAIN",$0
		}
		END{
			getline
			print "END",$0
		}`, in: "AAA\nBBB\nCCC\nDDD\n", ok: "BEGIN AAA\nMAIN BBB\nMAIN CCC\nMAIN DDD\nEND DDD\n"},

		// One Liner
		{script: "1", in: "AAA\n", ok: "AAA\n"},
		{script: "1;{print \"\"}", in: "AAA\nBBB\n", ok: "AAA\n\nBBB\n\n"},
		{script: "BEGIN{A[1]=1}A", in: "AAA\n", ok: "error:convert to bool failed in rule expression\n"},
		{script: "END{print}", in: "AAA\nBBB\nAAA\nDDD\n", ok: "DDD\n"},
		{script: "NF", in: "\n\nAAA\nBBB\n\n\nAAA\nDDD\n", ok: "AAA\nBBB\nAAA\nDDD\n"},
		{script: "$0", in: "\n\nAAA\nBBB\n\n\nAAA\nDDD\n", ok: "AAA\nBBB\nAAA\nDDD\n"},
		{script: "/./", in: "\n\nAAA\nBBB\n\n\nAAA\nDDD\n", ok: "AAA\nBBB\nAAA\nDDD\n"},
		{script: "NR==1", in: "AAA\nBBB\nAAA\nDDD\n", ok: "AAA\n"},
		{script: "NR%2", in: "AAA\nBBB\nAAA\nDDD\n", ok: "AAA\nAAA\n"},
		{script: "NR%2==0", in: "AAA\nBBB\nAAA\nDDD\n", ok: "BBB\nDDD\n"},
		{script: "{N+=length($0) } END{print N}", in: "AAA\nBBB\n", ok: "6\n"},
		{script: "{N+=NF} END{print N}", in: "AAA\nBBB\n", ok: "2\n"},
		{script: "END{print NR}", in: "AAA\nBBB\nAAA\nDDD\n", ok: "4\n"},
		{script: "{gsub(/[ \t]+/, \"\")}1", in: "AAA \tBBB\n", ok: "AAABBB\n"},
		{script: "{sub(/[ \t]+/, \"\")}1", in: "AAA \tBBB\n", ok: "AAABBB\n"},
		//{script: "A !~ $0 {A=$0}", in: "AAA\nAAA\nAAA\nDDD\n\nAAA\n", ok: "AAA\nDDD\nAAA\n"},
		{script: "!A[$0]++", in: "AAA\nAAA\nAAA\nDDD\nAAA\n", ok: "AAA\nDDD\n"},
		{script: "!($0 in A){A[$0];print}", in: "AAA\nAAA\nAAA\nDDD\nAAA\n", ok: "AAA\nDDD\n"}, //TODO
		{script: "{A[++C]=$0}END{for i=C;i>0;--i{print A[i]}}", in: "AAA\nBBB\nAAA\nDDD\n", ok: "DDD\nAAA\nBBB\nAAA\n"},
		{script: "/A+/{++N};END{print N+0}", in: "AAA\nBBB\nAAA\nDDD\n", ok: "2\n"},
		{script: "NF{$0=++A \" :\" $0};1", in: "AAA\n\nBBB\n", ok: "1 :AAA\n\n2 :BBB\n"},
		{script: "{print (NF? ++A \" :\" : \"\") $0}", in: "AAA\n\nBBB\n", ok: "1 :AAA\n\n2 :BBB\n"},
		{script: "$1 > Max {Max=$1; Maxline=$0}; END{ print Max, Maxline}", in: "10 AAA\n30 BBB\n20 CCC\n10 DDD\n", ok: "30 30 BBB\n"},

		// MAP
		{script: `{
						COUNT[$1]++
					}
					END{
						for (key in COUNT){
							print key,COUNT[key]
						}
						exit 0
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
		// MAP
		{script: `{
									for i=1; i<=NF; i++{
										COUNT[$i]++
									}
								}
								END{
									for (key in COUNT){
										print key,COUNT[key]
									}
									exit 0
								}`, in: `AAA BBB CCC AAA ZZZ AAA CCC
			`, ok: `AAA 3
BBB 1
CCC 2
ZZZ 1
`},

		// Command argment test
		{prepare: func() {}, cleanup: func() {}, rc: 0},
		{prepare: func() { os.Args = []string{os.Args[0], "-version"} }, cleanup: func() { *ver = false }, rc: 0, ok: "Version: 0.0.0\n"},
		{prepare: func() { *ver = true }, cleanup: func() { *ver = false }, rc: 0, ok: "Version: 0.0.0\n"},
		{prepare: func() { *dbg = true }, cleanup: func() { *dbg = false }, script: "{}", in: "aaa\n", rc: 0},
		{prepare: func() { *ast_dump = true }, cleanup: func() { *ast_dump = false }, script: "BEGIN{}{print 1}END{}", rc: 0},
		{prepare: func() { *globalVar = true }, cleanup: func() { *globalVar = false }, rc: 0},
		//{prepare: func() { *cpu_prof = true }, cleanup: func() { *cpu_prof = false }, rc: 0},
		//{prepare: func() { *mem_prof = true }, cleanup: func() { *mem_prof = false }, rc: 0},
		{prepare: func() { *dbglexer = true }, cleanup: func() { *dbglexer = false }, rc: 0},
		//{prepare: func() { os.Args = []string{os.Args[0], "-v", "XX"} }, cleanup: func() { variables = hash{} }, rc: 0, script: "BEGIN{print XX}", ok: "xx\n"},
		{prepare: func() { variables.Set("XX=xx") }, cleanup: func() { variables = hash{} }, rc: 0, script: "BEGIN{print XX}", ok: "xx\n"},
		{prepare: func() { variables.Set("XX") }, cleanup: func() { variables = hash{} }, rc: 0, script: "BEGIN{print XX}"},

		// test for script file
		{
			prepare: func() {
				scriptfile, err := ioutil.TempFile("", "example.*.ago")
				if err != nil {
					log.Fatal(err)
				}
				tempScriptPath = scriptfile.Name()
				fmt.Fprintf(scriptfile, "BEGIN{print 'Hello, World!';}")
				os.Args = []string{os.Args[0], "-f", scriptfile.Name()}
			},
			cleanup: func() {
				os.Remove(tempScriptPath)
				*program_file = ""
			},
			rc: 0,
			ok: "Hello, World!\n",
		},
		{
			prepare: func() { os.Args = []string{os.Args[0], "-f", "./xxaabbyyccccdd"} },
			cleanup: func() { *program_file = "" },
			rc:      1,
			ok:      "script file open error: open ./xxaabbyyccccdd: no such file or directory\n",
		},
		// test for data file
		{
			prepare: func() {
				datafile, err := ioutil.TempFile("", "example.*.data.ago")
				if err != nil {
					log.Fatal(err)
				}
				tempDataPath = datafile.Name()
				fmt.Fprintf(datafile, "AAA BBB CCC\nDDD EEE FFF\n")
				os.Args = []string{os.Args[0], "{print $1}", datafile.Name()}
				//fmt.Printf("os.Args=%#v\n", os.Args)
			},
			cleanup: func() {
				os.Remove(tempDataPath)
			},
			rc: 0,
			ok: "AAA\nDDD\n",
		},
		{
			prepare: func() { os.Args = []string{os.Args[0], "{print $1}", "./xxaabbyyccccdd"} },
			rc:      1,
			ok:      "input file open error: open ./xxaabbyyccccdd: no such file or directory\n",
		},
	}

	realStdin := os.Stdin
	realStdout := os.Stdout
	realStderr := os.Stderr
	case_number := 0

	for _, test := range tests {
		case_number++
		//t.Logf("script:%v\n", test.script)
		switch os.Getenv("TESTCASE") {
		case "":
		case "0":
			{
				fmt.Fprintf(realStdout, "case:%v script:%v\n", case_number, test.script)
			}
		default:
			{
				c, err := strconv.Atoi(os.Getenv("TESTCASE"))
				if err != nil {
					t.Fatal("Atoi error:", err)
				}
				if case_number != c {
					continue
				}
				fmt.Fprintf(realStdout, "case:%v script:%v\n", case_number, test.script)
			}
		}

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

			os.Args = []string{"ago"}
			if test.prepare != nil {
				test.prepare()
			}
			if test.script != "" {
				os.Args = append(os.Args, test.script)
			}
			rc := _main()
			if rc != test.rc && !strings.Contains(test.ok, "error") {
				t.Errorf("return code want:%v get:%v case:%v\n", test.rc, rc, test)
			}
			if test.cleanup != nil {
				test.cleanup()
			}

			/*
				rc := runScript(script_reader, os.Stdin)
				if rc != 0 {
					t.Fatal("runscript return code:", rc)
				}
			*/
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
		if test.ok != "" && resultOut != strings.Replace(test.ok, "\r", "", -1) { //replace for Windows
			t.Errorf("Case:[%v] received: %v - expected: %v - runSource: %v", case_number, resultOut, test.ok, test.script)
		}

		readFromIn.Close()
		writeToIn.Close()
		readFromOut.Close()
		writeToOut.Close()
	}

	os.Stdin = realStdin
	os.Stderr = realStderr
	os.Stdout = realStdout
}
