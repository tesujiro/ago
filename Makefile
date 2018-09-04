all : goa

goa : goa.go ./parser/*.y ./parser/*.go ./ast/*.go ./vm/*.go ./lib/*.go ./parser/grammar.go
	go build -o goa .

./parser/grammar.go : ./parser/grammar.go.y ./ast/*.go
	goyacc -o ./parser/grammar.go ./parser/grammar.go.y
	gofmt -s -w ./parser

# make case=xxx test  ## test specific case
.PHONY: test
test: ./*_test.go ./vm/*_test.go ./parser/grammar.go
ifdef case
	TESTCASE=${case} go test -v -count=1 .
else
	go test -v -count=1 . ./vm
endif

.PHONY: testcase
testcase: ./*_test.go ./vm/*_test.go ./parser/grammar.go
	TESTCASE=0 go test -v -count=1 .

.PHONY: prof
prof:
	# make prof ARG=[PPROF PATH]
	go tool pprof --pdf ./goa ${ARG} > ./prof.pdf
