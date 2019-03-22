# ago
If you were a gopher, you wrote a awk script for some task.
```
> awk '{if $1>10 {print}}' some.txt
awk: syntax error at source line 1
 context is
	{if >>>  $ <<< 1>10 {print}}
>
```
Ago is an alternative awk for gophers.

# Installation
```
$ go get -u github.com/tesujiro/ago
```

# Examples
```
> echo Long Time Ago | ago '{print "Hello",$NF,"!"}'
Hello Ago !
> printf "AAA\nBBB\nAAA\DDD" | ago '/A+/{++N};END{print N+0}'
2
> printf "AAA\nAAA\nDDD\nDDD\n" | ago '!A[$0]++'
AAA
DDD
> echo AAA BBB CCC AAA ZZZ AAA CCC | ago '{for i=1; i<=NF; i++{ COUNT[$i]++ }} END{for (key in COUNT){ print key,COUNT[key] }}'
AAA 3
BBB 1
CCC 2
ZZZ 1
> printf 'aaa,bbb,"ccc,ddd\neeee"\n' | ago -F , '{ret=split($0,a,"\"");if(ret%2==0){pLine=$0;pNF=NF;getline;NF=pNF+NF-1;$0=pLine $0};print NF;print}'
4
aaa,bbb,"ccc,dddeeee"
>

```

# Version
under development

Please note that the language and API may change at any time.


# Still not implemented
* nextfile
* var: FNR, OFMT, ENVIRON, CONVFMT
* print error with line number
* asort(),asort()
* import & call go library
* goroutine
* define variadic functions
* test command option

# Difference from AWK
* {} block has a local scope. variables name beginning with uppercase are global, others are local.
* with "-g" option all variables have global scopes.
* multiple value assignment (ex. BEGIN{a,b=1,2})
* define func in action (ex. { swap=func(a,b){return b,a}; print swap("a","b"); } // b a )
* anonymous func (ex. { print func(a,b){return b,a}("a","b"); } // b a )
* "if", "for", "while" conditions do not need parentheses, statements need curly braces. (ex. for a<10 {a=a+1}) 
* atan2(0,-0)=0 (not Pi)
* number format: exponent format (ex.2e+3) and hexadecimal (ex.0x10) is supported, octal format (ex.0123) is NOT supported.
* A command can be piped to getline (ex. "date" | getline DATE), but the command is not invoked from shell, cannot use shell functions in the command.
* can set multi chars to RS. (same as GAWK)
* Changing RS variable is valid only before scanning files. (Can change RS only in "BEGIN{}" rule.)

# To be fixed
* no bool literal (true/false->0/1)
