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
	github.com/tesujiro/ago/ast v0.0.0-00010101000000-000000000000 // indirect
	github.com/tesujiro/ago/debug v0.0.0-00010101000000-000000000000
	github.com/tesujiro/ago/lib v0.0.0-00010101000000-000000000000
	github.com/tesujiro/ago/parser v0.0.0-00010101000000-000000000000
	github.com/tesujiro/ago/vm v0.0.0-00010101000000-000000000000
	//github.com/tesujiro/ago/ast v0.0.0-20190921213547-14c7a56cfab4 // indirect
	//github.com/tesujiro/ago/debug v0.0.0-20190921213547-14c7a56cfab4
	//github.com/tesujiro/ago/lib v0.0.0-20190921213547-14c7a56cfab4
	//github.com/tesujiro/ago/parser v0.0.0-20190921213547-14c7a56cfab4
	//github.com/tesujiro/ago/vm v0.0.0-20190921213547-14c7a56cfab4
	golang.org/x/lint v0.0.0-20190909230951-414d861bb4ac // indirect
)
