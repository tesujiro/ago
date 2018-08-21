# goa
go on awk (just like Unagi on Katsudon)

# version
under development

# Still not implemented
* printf, sprintf
* for(;;)
* next, nextfile
* getline
* exit
* func: rand, srand, int, split, system
* var: FNR, OFMT, RS, ENVIRON, CONVFMT, RSTART, RLENGTH

# Difference from AWK
* {} block has localscope. variables name beginning with uppercase are global, others are local
* define func in action (ex. { swap=func(a,b){return b,a}; print swap("a","b"); } // b a )
* anonymous func (ex. { print func(a,b){return b,a}("a","b"); } // b a )
* function call is call by value.
* sub,gsub returns replaced string
* "while (expression){statements}" can be "for expression {statements}"

# To be fixed
* 1==true, 0==false
* NR,$n is not set in END{} 
* REGEXP "/.../" -> /.../
* function args cannot be omitted (like sub(r, t[, s]))
* function match() does not set RSTART, RLENGTH
* FS=="\n" : error
* ++A : invalid operation
* A++ : incremented before evaluation
