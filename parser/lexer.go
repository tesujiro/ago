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
	"BEGIN":    BEGIN,
	"END":      END,
	"delete":   DELETE,
	"print":    PRINT,
	"if":       IF,
	"else":     ELSE,
	"for":      FOR,
	"break":    BREAK,
	"continue": CONTINUE,
	"in":       IN,
	"true":     TRUE,
	"false":    FALSE,
	"nil":      NIL,
	"function": FUNC,
	"func":     FUNC,
	"return":   RETURN,
	"while":    WHILE,
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
	/*
		if IN_REGEXP {
			fmt.Println("IN_REGEXP")
			regexp := ""
			for {
				_, tok, lit := l.Scan()
				if tok == token.QUO || tok == token.EOF || lit == "\n" {
					break
				}
				fmt.Printf("lit=[%v] tok=[%v]\n", lit, tok)
				if lit == "" {
					regexp += tok.String()
					//TODO; '\\'
				} else {
					regexp += lit
				}
			}
			token_id = REGEXP
			lval.token = ast.Token{Token: token_id, Literal: regexp}
			//IN_REGEXP = false
			return token_id
		}
	*/
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
		if len(lit) > 3 && lit[:2] == "\"/" && lit[len(lit)-2:] == "/\"" {
			token_id = REGEXP
			lit = lit[2 : len(lit)-2]
		} else {
			token_id = STRING
			if len(lit) > 1 {
				lit = lit[1 : len(lit)-1]
			}
		}
	case token.EOF:
		token_id = 0
	case token.ILLEGAL:
		switch lit {
		case "$", "~":
			token_id = int([]rune(lit)[0])
		}
	default:
		if symbol, ok := compSymbols[tok.String()]; ok {
			token_id = symbol
		} else {
			if len(tok.String()) == 1 {
				token_id = int(tok.String()[0])
			} else {
				token_id = IDENT
			}
		}
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
