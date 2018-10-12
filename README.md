# ago
awk goes on
- awk implemented on go

# Installation
```
$ go get -u github.com/tesujiro/ago
```

# Examples
```
# echo Once Time Ago | ago '{print "Hello",$NF,"!"}'
Hello Ago !

```

# Version
under development

Please note that the language and API may change at any time.


# Still not implemented
* nextfile
* getline
* var: FNR, OFMT, RS, ENVIRON, CONVFMT
* print error with line number
* asort(),asorti()
* import & call go library
* goroutine
* define variadic functions
* test command option

# Difference from AWK
* {} block has a local scope. variables name beginning with uppercase are global, others are local
* with -g option all variables have global scopes.
* multiple value (ex. BEGIN{a,b=1,2})
* define func in action (ex. { swap=func(a,b){return b,a}; print swap("a","b"); } // b a )
* anonymous func (ex. { print func(a,b){return b,a}("a","b"); } // b a )
* if, for, while conditions do not need parentheses, statements need curly braces. (ex. for a<10 {a=a+1}) 
* atan2(0,-0)=0 (not Pi)

# To be fixed
* split(s,g,f): arg g must be initialized before calling the func
* sub(), gsub(): arg string must be set to a global variable
* error: S=="\n"
* no bool literal (true/false->0/1)
