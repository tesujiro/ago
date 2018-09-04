# goa
go on awk (just like Unagi on Katsudon)

# version
under development

# Still not implemented
* regular expressin not match "!~"
* printf, sprintf
* comment
* nextfile
* getline
* func: rand, srand, int, split, system
* var: FNR, OFMT, RS, ENVIRON, CONVFMT, RSTART, RLENGTH
* print error with line number
* asort(),asorti()
* import & call go library
* goroutine
* -g global varibale option
* define variadice functions

# Difference from AWK
* {} block has a local scope. variables name beginning with uppercase are global, others are local
* with -g option all variables have global scopes.
* multiple value (ex. BEGIN{a,b=1,2})
* define func in action (ex. { swap=func(a,b){return b,a}; print swap("a","b"); } // b a )
* anonymous func (ex. { print func(a,b){return b,a}("a","b"); } // b a )
* if, for, while conditions do not need parentheses, statements need curly braces. (ex. for a<10 {a=a+1}) 

# To be fixed
* no function option args (ex. sub(r, t[, s]))
* function call is call by value.
* sub,gsub returns replaced string
* REGEXP "/.../" -> /.../
* function match() does not set RSTART, RLENGTH
* error: S=="\n"
