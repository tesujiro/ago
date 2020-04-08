ago -g '
	{ nf = nf + NF }
END	{ print nf }
' emp.data
