ago -g '
	{ for  i = 1; i <= NF; i = i + 1 {sum = sum + $i}}
END	{ print sum }
' emp.data
