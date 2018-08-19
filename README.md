# goa
go on awk (just like Unagi on Katsudon)


# difference from AWK
* function call is call by value.
* function args cannot be omitted (like sub(r, t[, s]))
* ++A : invalid operation
* A++ : incremented before evaluation
* FS=="\n" : error

# To be fixed
* REGEXP "/.../" -> /.../
