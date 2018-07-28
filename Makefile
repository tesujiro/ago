all : goa

goa : goa.go ./parser/*.y ./parser/*.go ./ast/*.go ./vm/*.go
	go build -o goa .

./parser/grammar.go : ./parser/grammar.go.y ./ast/*.go
	goyacc -o ./parser/grammar.go ./parser/grammar.go.y
	gofmt -s -w ./parser

test: *_test.go
	go test -v . ./vm ./lib

prof:
	# make prof ARG=[PPROF PATH]
	go tool pprof --pdf ./goa ${ARG} > ./prof.pdf
