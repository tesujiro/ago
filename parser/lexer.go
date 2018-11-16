// Package parser implements parser for anko.
package parser

import (
	"errors"
	"fmt"
	"unicode"

	"github.com/tesujiro/ago/ast"
)

const (
	// EOF is short for End of file.
	EOF = -1
	// EOL is short for End of line.
	EOL = '\n'
)

// Error provides a convenient interface for handling runtime error.
// It can be Error interface with type cast which can call Pos().
type Error struct {
	Message  string
	Pos      ast.Position
	Filename string
	Fatal    bool
}

var EOF_FLAG bool
var traceLexer bool
var maybe_regexp int
var regexp_str string

// Error returns the error message.
func (e *Error) Error() string {
	return e.Message
}

// Scanner stores informations for lexer.
type Scanner struct {
	src      []rune
	offset   int
	lineHead int
	line     int
}

// opName is correction of operation names.
var opName = map[string]int{
	"BEGIN":    BEGIN,
	"END":      END,
	"delete":   DELETE,
	"print":    PRINT,
	"printf":   PRINTF,
	"if":       IF,
	"else":     ELSE,
	"for":      FOR,
	"break":    BREAK,
	"continue": CONTINUE,
	"next":     NEXT,
	"in":       IN,
	"true":     TRUE,
	"false":    FALSE,
	"nil":      NIL,
	"function": FUNC,
	"func":     FUNC,
	"return":   RETURN,
	"exit":     EXIT,
	"while":    WHILE,
	"do":       DO,
}

// Scan analyses token, and decide identify or literals.
func (s *Scanner) Scan() (tok int, lit string, pos ast.Position, err error) {
retry:
	//s.skipBlank()
	if maybe_regexp == 0 {
		s.skipBlank()
		pos = s.pos()
	} else {
		//fmt.Println("maybe_regexp:", maybe_regexp, " IN_REGEXP:", IN_REGEXP)
		maybe_regexp++
		blank := s.skipBlank()
		regexp_str = blank
		pos = s.pos()
		if IN_REGEXP {
			for ch := s.peek(); !isEOL(ch) && ch != '/'; ch = s.peek() {
				regexp_str = fmt.Sprintf("%s%c", regexp_str, ch)
				s.next()
			}
			tok = REGEXP
			lit = regexp_str
			regexp_str = ""
			IN_REGEXP = false
			s.next()
			return
		}
		maybe_regexp = 0
		regexp_str = ""
	}
	s.peek()
	/*
		ch := s.peek()
			if maybe_regexp != 0 {
				regexp_str = fmt.Sprintf("%s%c", regexp_str, ch)
			}
	*/
	switch ch := s.peek(); {
	case isLetter(ch):
		lit, err = s.scanIdentifier()
		if err != nil {
			return
		}
		if name, ok := opName[lit]; ok {
			tok = name
		} else {
			tok = IDENT
		}
	case isDigit(ch):
		tok = NUMBER
		lit, err = s.scanNumber()
		if err != nil {
			return
		}
	case ch == '"':
		tok = STRING
		lit, err = s.scanString('"')
		if err != nil {
			return
		}

	case ch == '\'':
		tok = STRING
		lit, err = s.scanString('\'')
		if err != nil {
			return
		}
	/*
		case ch == '`':
			tok = STRING
			lit, err = s.scanRawString('`')
			if err != nil {
				return
			}
	*/
	default:
		switch ch {
		case EOF:
			if !EOF_FLAG {
				tok = int(';')
				lit = string(';')
				EOF_FLAG = true
			} else {
				tok = EOF
				EOF_FLAG = true
			}
		case EOL:
			tok = int(';')
			lit = string(';')
		case '#': // COMMENT
			for !isEOL(s.peek()) {
				s.next()
			}
			goto retry
		case '!':
			s.next()
			switch s.peek() {
			case '=':
				tok = NEQ
				lit = "!="
			case '~':
				tok = NOTTILDE
				lit = "!~"
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '=':
			s.next()
			switch s.peek() {
			case '=':
				tok = EQEQ
				lit = "=="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '+':
			s.next()
			switch s.peek() {
			case '+':
				tok = PLUSPLUS
				lit = "++"
			case '=':
				tok = PLUSEQ
				lit = "+="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '-':
			s.next()
			switch s.peek() {
			case '-':
				tok = MINUSMINUS
				lit = "--"
			case '=':
				tok = MINUSEQ
				lit = "-="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '*':
			s.next()
			switch s.peek() {
			case '=':
				tok = MULEQ
				lit = "*="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '/':
			//fmt.Println("in lexer: QUO")
			maybe_regexp = 1
			s.next()
			switch s.peek() {
			case '=': //TODO:  ??
				tok = DIVEQ
				lit = "/="
				/*  in AWK , '//' is not a comment, but a regexp
				case '/':
					for !isEOL(s.peek()) {
						s.next()
					}
					goto retry
				*/
			case '*':
				for {
					_, err = s.scanRawString('*')
					if err != nil {
						return
					}

					if s.peek() == '/' {
						s.next()
						goto retry
					}

					s.back()
				}
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '%':
			s.next()
			switch s.peek() {
			case '=':
				tok = MODEQ
				lit = "%="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '>':
			s.next()
			switch s.peek() {
			case '=':
				tok = GE
				lit = ">="
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '<':
			s.next()
			switch s.peek() {
			case '=':
				tok = LE
				lit = "<="
			/*
				case '-':
					tok = OPCHAN
					lit = "<-"
			*/
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '|':
			s.next()
			switch s.peek() {
			case '|':
				tok = OROR
				lit = "||"
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		case '&':
			s.next()
			switch s.peek() {
			case '&':
				tok = ANDAND
				lit = "&&"
			default:
				s.back()
				tok = int(ch)
				lit = string(ch)
			}
		/*
			case '.':
				s.next()
				if s.peek() == '.' {
					s.next()
					if s.peek() == '.' {
						tok = VARARG
					} else {
						err = fmt.Errorf("syntax error on '%v' at %v:%v", string(ch), pos.Line, pos.Column)
						return
					}
				} else {
					s.back()
					tok = int(ch)
					lit = string(ch)
				}
		*/
		case '(', ')', ':', ';', '{', '}', '[', ']', ',', '^', '?', '$', '~', '.':
			tok = int(ch)
			lit = string(ch)
		default:
			err = fmt.Errorf("syntax error on '%v' at %v:%v", string(ch), pos.Line, pos.Column)
			tok = int(ch)
			lit = string(ch)
			return
		}
		s.next()
	}
	return
}

// isLetter returns true if the rune is a letter for identity.
func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

// isDigit returns true if the rune is a number.
func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

// isHex returns true if the rune is a hex digits.
func isHex(ch rune) bool {
	return ('0' <= ch && ch <= '9') || ('a' <= ch && ch <= 'f') || ('A' <= ch && ch <= 'F')
}

// isEOL returns true if the rune is at end-of-line or end-of-file.
func isEOL(ch rune) bool {
	return ch == '\n' || ch == -1
}

// isBlank returns true if the rune is empty character..
func isBlank(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\r'
}

func appendNumberAndPoint(s *Scanner, ret []rune) []rune {
	for isDigit(s.peek()) || s.peek() == '.' {
		ret = append(ret, s.peek())
		s.next()
	}
	return ret
}

// peek returns current rune in the code.
func (s *Scanner) peek() rune {
	if s.reachEOF() {
		return EOF
	}
	return s.src[s.offset]
}

// next moves offset to next.
func (s *Scanner) next() {
	if !s.reachEOF() {
		if s.peek() == '\n' {
			s.lineHead = s.offset + 1
			s.line++
		}
		s.offset++
	}
}

// current returns the current offset.
func (s *Scanner) current() int {
	return s.offset
}

// offset sets the offset value.
func (s *Scanner) set(o int) {
	s.offset = o
}

// back moves back offset once to top.
func (s *Scanner) back() {
	s.offset--
}

// reachEOF returns true if offset is at end-of-file.
func (s *Scanner) reachEOF() bool {
	return len(s.src) <= s.offset
}

// pos returns the position of current.
func (s *Scanner) pos() ast.Position {
	return ast.Position{Line: s.line + 1, Column: s.offset - s.lineHead + 1}
}

// skipBlank moves position into non-black character.
/*
func (s *Scanner) skipBlank() {
	for isBlank(s.peek()) {
		s.next()
	}
}
*/
func (s *Scanner) skipBlank() string {
	str := ""
	for ch := s.peek(); isBlank(ch); ch = s.peek() {
		str = fmt.Sprintf("%s%c", str, ch)
		s.next()
	}
	return str
}

// scanIdentifier returns identifier beginning at current position.
func (s *Scanner) scanIdentifier() (string, error) {
	var ret []rune
	for {
		if !isLetter(s.peek()) && !isDigit(s.peek()) {
			break
		}
		ret = append(ret, s.peek())
		s.next()
	}
	return string(ret), nil
}

// scanNumber returns number beginning at current position.
func (s *Scanner) scanNumber() (string, error) {
	var ret []rune
	ch := s.peek()
	ret = append(ret, ch)
	s.next()
	if ch == '0' && s.peek() == 'x' {
		ret = append(ret, s.peek())
		s.next()
		for isHex(s.peek()) {
			ret = append(ret, s.peek())
			s.next()
		}
	} else {
		ret = appendNumberAndPoint(s, ret)
		if s.peek() == 'e' {
			ret = append(ret, s.peek())
			s.next()
			if isDigit(s.peek()) || s.peek() == '+' || s.peek() == '-' {
				ret = append(ret, s.peek())
				s.next()
				ret = appendNumberAndPoint(s, ret)
			}
			ret = appendNumberAndPoint(s, ret)
		}
		if isLetter(s.peek()) {
			return "", errors.New("identifier starts immediately after numeric literal")
		}
	}
	return string(ret), nil
}

// scanRawString returns raw-string starting at current position.
func (s *Scanner) scanRawString(l rune) (string, error) {
	var ret []rune
	for {
		s.next()
		if s.peek() == EOF {
			return "", errors.New("unexpected EOF")
		}
		if s.peek() == l {
			s.next()
			break
		}
		ret = append(ret, s.peek())
	}
	return string(ret), nil
}

// scanString returns string starting at current position.
// This handles backslash escaping.
func (s *Scanner) scanString(l rune) (string, error) {
	var ret []rune
eos:
	for {
		s.next()
		switch s.peek() {
		case EOL:
			return "", errors.New("unexpected EOL")
		case EOF:
			return "", errors.New("unexpected EOF")
		case l:
			s.next()
			break eos
		case '\\':
			s.next()
			switch s.peek() {
			case 'b':
				ret = append(ret, '\b')
				continue
			case 'f':
				ret = append(ret, '\f')
				continue
			case 'r':
				ret = append(ret, '\r')
				continue
			case 'n':
				ret = append(ret, '\n')
				continue
			case 't':
				ret = append(ret, '\t')
				continue
			}
			ret = append(ret, s.peek())
			continue
		default:
			ret = append(ret, s.peek())
		}
	}
	return string(ret), nil
}

// Lexer provides interface to parse codes.
type Lexer struct {
	s      *Scanner
	lit    string
	pos    ast.Position
	e      error
	result []ast.Rule
}

// Lex scans the token and literals.
func (l *Lexer) Lex(lval *yySymType) int {
	tok, lit, pos, err := l.s.Scan()
	if traceLexer {
		fmt.Printf("tok:%v\tlit:%v\tpos:%v\terr:%v\n", tok, lit, pos, err)
	}
	if err != nil {
		l.e = &Error{Message: err.Error(), Pos: pos, Fatal: true}
	}
	lval.token = ast.Token{Token: tok, Literal: lit}
	lval.token.SetPosition(pos)
	l.lit = lit
	l.pos = pos
	return tok
}

// Error sets parse error.
func (l *Lexer) Error(msg string) {
	l.e = &Error{Message: msg, Pos: l.pos, Fatal: false}
}

// Parse provides way to parse the code using Scanner.
func Parse(s *Scanner) ([]ast.Rule, error) {
	l := Lexer{s: s}
	if yyParse(&l) != 0 {
		return nil, l.e
	}
	return l.result, l.e
}

/*
// EnableErrorVerbose enabled verbose errors from the parser
func EnableErrorVerbose() {
	yyErrorVerbose = true
}
*/

func TraceLexer() {
	traceLexer = true
}

func TraceOffLexer() {
	traceLexer = false
}

func initialize() {
	EOF_FLAG = false
}

// ParseSrc provides way to parse the code from source.
func ParseSrc(src string) ([]ast.Rule, error) {
	initialize()
	scanner := &Scanner{
		src: []rune(src),
	}
	return Parse(scanner)
}
