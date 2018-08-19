# goa
go on awk (just like Unagi on Katsudon)

# still not implemented
* /start/,/stop/
* printf, sprintf
* do while, while
* for(;;)
* next, nextfile
* getline
* exit
* func: rand, srand, int, split, system
* var: FNR, OFMT, RS, ENVIRON, CONVFMT, RSTART, RLENGTH

# difference from AWK
* variables name beginning with uppercase are global, others are local
* function call is call by value.
* sub,gsub returns replaced string

# To be fixed
* REGEXP "/.../" -> /.../
* function args cannot be omitted (like sub(r, t[, s]))
* function match() does not set RSTART, RLENGTH
* FS=="\n" : error
* ++A : invalid operation
* A++ : incremented before evaluation
