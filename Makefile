all : ago

ago : ago.go ./parser/*.y ./parser/*.go ./ast/*.go ./vm/*.go ./lib/*.go ./parser/grammar.go
	#go get -u
	go build -o ago .
	cp ago $$GOPATH/bin

./parser/grammar.go : ./parser/grammar.go.y ./ast/*.go
	goyacc -l -o ./parser/grammar.go ./parser/grammar.go.y
	gofmt -s -w ./parser

.PHONY: lint
lint: 
	golint . ./...

# make case=xxx test  ## test specific case
.PHONY: test
test: ./*_test.go ./vm/*_test.go ./parser/grammar.go
	go vet ./...
	#go get -u
ifdef case
	TESTCASE=${case} go test -v -count=1 .
else
	go test -v -count=1 github.com/tesujiro/ago/parser github.com/tesujiro/ago/vm . -coverpkg ./...
endif

.PHONY: cover
cover:
	go test -v github.com/tesujiro/ago/parser github.com/tesujiro/ago/vm . -coverpkg ./... -coverprofile=cover.out && go tool cover -html=cover.out -o cover.html

.PHONY: testcase
testcase: ./*_test.go ./vm/*_test.go ./parser/grammar.go
	TESTCASE=0 go test -v -count=1 .

.PHONY: prof_cpu
prof_cpu:
	go tool pprof --pdf ./ago ./cpu.pprof > ./cpu.pdf

.PHONY: prof_mem
prof_mem:
	go tool pprof --pdf ./ago ./mem.pprof > ./mem.pdf

.PHONY: bench
bench:
	go test -bench . -benchmem

.PHONY: merge
merge:
	git checkout master && git merge develop && git push && git checkout develop
