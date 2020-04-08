awk '
	{ field = $NF }
END	{ print field }
' emp.data
