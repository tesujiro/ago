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
	github.com/ortuman/jackal v0.8.2 // indirect
	github.com/pelletier/go-toml v1.6.0 // indirect
	github.com/pkg/profile v1.3.0
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.5.0 // indirect
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
	golang.org/x/sys v0.0.0-20191115151921-52ab43148777 // indirect
	golang.org/x/text v0.3.2 // indirect
)
