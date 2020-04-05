ago -g '
	{ last = $0 }
END	{ print last }
' emp.data
