# ago
[![GoDoc Reference](https://godoc.org/github.com/tesujiro/ago?status.svg)](http://godoc.org/github.com/tesujiro/ago)
[![Build Status](https://travis-ci.org/tesujiro/ago.svg?branch=master)](https://travis-ci.org/tesujiro/ago)
[![Coverage](https://codecov.io/gh/tesujiro/ago/branch/master/graph/badge.svg)](https://codecov.io/gh/tesujiro/ago)
[![Go Report Card](https://goreportcard.com/badge/github.com/tesujiro/ago)](https://goreportcard.com/report/github.com/tesujiro/ago)

If you are a gopher, you wrote a AWK script for some task.
```
$ awk '{if $1>10 {print}}' some.txt
awk: syntax error at source line 1
 context is
	{if >>>  $ <<< 1>10 {print}}
$
```
Ago is an alternative AWK for gophers.
```
$ ago '{if $1>10 {print}}' some.txt
15 You are a gopher!!
$
```
# Installation
```
$ go get -u github.com/tesujiro/ago
```

# Examples
```
$ echo A Long Time Ago | ago '{print "Hello, ",$NF,"!"}'
Hello, Ago !
$ printf "AAA\nBBB\nAAA\DDD" | ago '/A+/{++N};END{print N}'
2
$ printf "AAA\nAAA\nDDD\nDDD\n" | ago '!A[$0]++'
AAA
DDD
$ echo AAA BBB CCC AAA ZZZ AAA CCC | ago '{ for i=1; i<=NF; i++{ COUNT[$i]++ }} END{for (key in COUNT){ print key, COUNT[key] }}'
AAA 3
BBB 1
CCC 2
ZZZ 1
$ printf 'aaa,bbb,"ccc,ddd\neeee"\n' | ago -F , '{ ret=split($0,a,"\""); if(ret%2==0){ pLine=$0; pNF=NF; getline; NF=pNF+NF-1; $0=pLine $0}; print NF; print}'
4
aaa,bbb,"ccc,dddeeee"
$ echo 20 | ago 'BEGIN{ Factorial=func(x){ if x==1 {1} else { x*Factorial(x-1) }}} {print Factorial($1)}'
2432902008176640000
$ echo 12 34| ago '{ print func(a, b){return b, a}($1, $2) }'
34 12
$

```

# Version
under development

Please note that the language and command args may change at any time.


# Still not implemented
* var: CONVFMT
* print error with line number
* asort(),asort()
* import & call go library
* goroutine
* define variadic functions
* test command option

# Difference from AWK
* {} block has a local scope. variables name beginning with uppercase are global, others are local.
* with "-g" option all variables have global scopes (same as AWK).
* multiple value assignment (ex. BEGIN{a,b=1,2})
* define func in action (ex. { swap=func(a,b){return b,a}; print swap("a","b"); } // b a )
* anonymous func (ex. { print func(a,b){return b,a}("a","b"); } // b a )
* "if", "for", "while" conditions do not need parentheses, statements need curly braces. (ex. for a<10 {a=a+1}) 
* atan2(0,-0)=0 (not Pi)
* number format: exponent format (ex.2e+3) and hexadecimal (ex.0x10) is supported, octal format (ex.0123) is NOT supported.
* A command can be piped to getline (ex. "date" | getline DATE), but the command is not invoked from shell, cannot use shell functions in the command.
* can set multiple chars to RS. (same as GAWK)
* Changing RS variable is valid only before scanning files. (Can change RS only in "BEGIN{}" rule.)
* All unicode character can be used for variable names(ex.  ビール="beer!";print ビール)

# To be fixed
* no bool literal (true/false->0/1)
