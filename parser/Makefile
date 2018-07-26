all : grammar.go

grammar.go : grammar.go.y ../ast/*.go
	goyacc -o $@ grammar.go.y
	gofmt -s -w .
