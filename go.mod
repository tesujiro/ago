module github.com/tesujiro/ago

go 1.13

replace (
	github.com/tesujiro/ago/ast => ./ast
	github.com/tesujiro/ago/debug => ./debug
	github.com/tesujiro/ago/lib => ./lib
	github.com/tesujiro/ago/parser => ./parser
	github.com/tesujiro/ago/vm => ./vm
)

require (
	github.com/pkg/profile v1.3.0
	github.com/tesujiro/ago/debug v0.0.0-20200224044000-d07d2c158e90
	github.com/tesujiro/ago/lib v0.0.0-00010101000000-000000000000
	github.com/tesujiro/ago/parser v0.0.0-00010101000000-000000000000
	github.com/tesujiro/ago/vm v0.0.0-20200224044000-d07d2c158e90
)
