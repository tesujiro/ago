awk '
	{ nf = nf + NF }
END	{ print nf }
' emp.data
