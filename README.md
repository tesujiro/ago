# goa
go on awk (just like Unagi on Katsudon)

# version
under development

# Still not implemented
* printf, sprintf
* next, nextfile
* getline
* exit
* func: rand, srand, int, split, system
* var: FNR, OFMT, RS, ENVIRON, CONVFMT, RSTART, RLENGTH
* print error with line number
* datetime funcs
* asort(),asorti()

# Difference from AWK
* {} block has localscope. variables name beginning with uppercase are global, others are local
* define func in action (ex. { swap=func(a,b){return b,a}; print swap("a","b"); } // b a )
* anonymous func (ex. { print func(a,b){return b,a}("a","b"); } // b a )
* function call is call by value.
* sub,gsub returns replaced string
* if, for, while conditions do not need parentheses, statements need curly braces. (ex. for a<10 {a=a+1}) 

# To be fixed
* set NF=1 when empty line
* REGEXP "/.../" -> /.../
* function args cannot be omitted (like sub(r, t[, s]))
* function match() does not set RSTART, RLENGTH
* FS=="\n" : error
* ++A : invalid operation
* A++ : incremented before evaluation
