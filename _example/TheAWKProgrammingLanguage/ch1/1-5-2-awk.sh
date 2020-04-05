awk '
END	{ print NR, "employees" }
' emp.data
