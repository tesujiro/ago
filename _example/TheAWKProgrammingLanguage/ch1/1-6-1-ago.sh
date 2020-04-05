ago -g '
$2 > 6	{ n = n + 1; pay = pay + $2 * $3 }
END	{ if (n > 0) {
		print n, "emloyees, total pay is", pay,
			 "average pay is", pay/n
	  } else {
		print" no employees are paid more tahan $6/hour"
	  }
	}
' emp.data
