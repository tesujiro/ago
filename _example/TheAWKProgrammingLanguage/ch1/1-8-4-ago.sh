ago -g '
	{ field = $NF }
END	{ print field }
' emp.data
