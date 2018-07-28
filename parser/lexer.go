package parser

import (
	"errors"
	"go/scanner"
	"go/token"

	"github.com/tesujiro/goa/ast"
)

type Lexer struct {
	scanner.Scanner
	result []ast.Rule
	err    error
}

// tokenTab is a correction of operation names.
var tokenTab = map[string]int{
	"BEGIN": LEX_BEGIN,
	"END":   LEX_END,
	"print": LEX_PRINT,
}

var compSymbols = map[string]int{
	"==": EQEQ,
	"!=": NEQ,
	">=": GE,
	"<=": LE,
	"&&": ANDAND,
	"||": OROR,
	"++": PLUSPLUS,
	"--": MINUSMINUS,
	"+=": PLUSEQ,
	"-=": MINUSEQ,
	"*=": MULEQ,
	"/=": DIVEQ,
}

func (l *Lexer) Lex(lval *yySymType) (token_id int) {
	//TODO: Position
	_, tok, lit := l.Scan()
	if name, ok := tokenTab[lit]; ok {
		token_id = name
		lval.token = ast.Token{Token: token_id, Literal: lit}
		//fmt.Printf("tok=%v\ttoken=%#v\n", tok.String(), lval.token)
		return token_id
	}
	switch tok {
	case token.IDENT:
		token_id = IDENT
	case token.INT:
		token_id = NUMBER
	case token.FLOAT:
		token_id = NUMBER
	case token.STRING:
		token_id = STRING
		if len(lit) > 1 {
			lit = lit[1 : len(lit)-1]
		}
	default:
		if symbol, ok := compSymbols[tok.String()]; ok {
			token_id = symbol
		} else {
			if len(tok.String()) == 1 {
				token_id = int(tok.String()[0])
			}
		}
	}
	if lit == "$" {
		token_id = int('$')
	}
	lval.token = ast.Token{Token: token_id, Literal: lit}
	//fmt.Printf("tok=%v\ttoken=%#v\n", tok.String(), lval.token)
	return token_id
}

func (l *Lexer) Error(e string) {
	l.err = errors.New(e)
	//l.position = l.Position //TODO
}

func Parse(yylex yyLexer) ([]ast.Rule, error) {
	l := yylex.(*Lexer)
	if yyParse(yylex) != 0 {
		return nil, l.err
	}
	return l.result, nil
}
