ago -g '
/Beth/	{ nlines = nlines + 1 }
END	{ print nlines }
' emp.data
