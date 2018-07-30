all : goa

goa : goa.go ./parser/*.y ./parser/*.go ./ast/*.go ./vm/*.go
	go build -o goa .

./parser/grammar.go : ./parser/grammar.go.y ./ast/*.go
	goyacc -o ./parser/grammar.go ./parser/grammar.go.y
	gofmt -s -w ./parser

.PHONY: test
test: *_test.go test/*.goa test/*.in test/*.ok
	#go test -v . ./vm ./lib
	go test -v -count=1 . 

.PHONY: prof
prof:
	# make prof ARG=[PPROF PATH]
	go tool pprof --pdf ./goa ${ARG} > ./prof.pdf
