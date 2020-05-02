	{ s = s substr($1, 1 ,3) " "}
END	{ print substr(s, 1, length(s)-1) }
