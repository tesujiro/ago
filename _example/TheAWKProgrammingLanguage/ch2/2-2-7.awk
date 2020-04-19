$3 > maxpop	{ maxpop = $3 + 0; country = $1 }
END		{ print "country with greatest population:",
    			country, maxpop
		}
