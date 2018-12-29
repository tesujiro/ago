//line ./parser/grammar.go.y:2
package parser

import __yyfmt__ "fmt"

//line ./parser/grammar.go.y:2
import (
	//"fmt"
	"github.com/tesujiro/ago/ast"
)

var defaultExpr = ast.FieldExpr{Expr: &ast.NumExpr{Literal: "0"}}
var defaultExprs = []ast.Expr{&defaultExpr}
var IN_REGEXP bool

//line ./parser/grammar.go.y:13
type yySymType struct {
	yys        int
	token      ast.Token
	rule       ast.Rule
	rules      []ast.Rule
	pattern    ast.Pattern
	stmt       ast.Stmt
	stmts      []ast.Stmt
	expr       ast.Expr
	exprs      []ast.Expr
	ident_args []string
}

const IDENT = 57346
const NUMBER = 57347
const STRING = 57348
const TRUE = 57349
const FALSE = 57350
const NIL = 57351
const EQEQ = 57352
const NEQ = 57353
const GE = 57354
const LE = 57355
const NOTTILDE = 57356
const ANDAND = 57357
const OROR = 57358
const LEN = 57359
const PLUSPLUS = 57360
const MINUSMINUS = 57361
const PLUSEQ = 57362
const MINUSEQ = 57363
const MULEQ = 57364
const DIVEQ = 57365
const MODEQ = 57366
const DELETE = 57367
const IN = 57368
const BEGIN = 57369
const END = 57370
const PRINT = 57371
const PRINTF = 57372
const REGEXP = 57373
const IF = 57374
const ELSE = 57375
const FOR = 57376
const WHILE = 57377
const DO = 57378
const BREAK = 57379
const CONTINUE = 57380
const FUNC = 57381
const RETURN = 57382
const EXIT = 57383
const NEXT = 57384
const CONCAT_OP = 57385
const vars = 57386
const UNARY = 57387

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENT",
	"NUMBER",
	"STRING",
	"TRUE",
	"FALSE",
	"NIL",
	"EQEQ",
	"NEQ",
	"GE",
	"LE",
	"NOTTILDE",
	"ANDAND",
	"OROR",
	"LEN",
	"PLUSPLUS",
	"MINUSMINUS",
	"PLUSEQ",
	"MINUSEQ",
	"MULEQ",
	"DIVEQ",
	"MODEQ",
	"DELETE",
	"IN",
	"BEGIN",
	"END",
	"PRINT",
	"PRINTF",
	"REGEXP",
	"IF",
	"ELSE",
	"FOR",
	"WHILE",
	"DO",
	"BREAK",
	"CONTINUE",
	"FUNC",
	"RETURN",
	"EXIT",
	"NEXT",
	"CONCAT_OP",
	"'='",
	"'?'",
	"':'",
	"','",
	"vars",
	"'~'",
	"'>'",
	"'<'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'!'",
	"UNARY",
	"'$'",
	"'['",
	"'('",
	"')'",
	"'{'",
	"'}'",
	"';'",
	"']'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line ./parser/grammar.go.y:597

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 13,
	60, 76,
	-2, 70,
	-1, 14,
	64, 12,
	-2, 105,
	-1, 102,
	64, 107,
	-2, 105,
	-1, 117,
	47, 92,
	-2, 89,
	-1, 157,
	64, 12,
	-2, 105,
	-1, 162,
	64, 12,
	-2, 105,
	-1, 164,
	64, 12,
	-2, 105,
	-1, 177,
	64, 12,
	-2, 105,
	-1, 182,
	64, 12,
	-2, 105,
	-1, 184,
	64, 12,
	-2, 105,
	-1, 188,
	64, 12,
	-2, 105,
	-1, 192,
	64, 12,
	-2, 105,
	-1, 205,
	60, 89,
	61, 89,
	-2, 93,
	-1, 209,
	64, 12,
	-2, 105,
	-1, 222,
	64, 12,
	-2, 105,
	-1, 223,
	64, 12,
	-2, 105,
}

const yyPrivate = 57344

const yyLast = 605

var yyAct = [...]int{

	58, 166, 12, 23, 103, 46, 94, 57, 5, 126,
	59, 15, 13, 14, 128, 5, 212, 185, 44, 43,
	229, 228, 221, 216, 65, 214, 55, 87, 211, 63,
	210, 66, 63, 174, 208, 204, 202, 63, 198, 63,
	63, 100, 223, 96, 97, 98, 63, 63, 42, 104,
	120, 121, 122, 123, 124, 125, 99, 127, 222, 177,
	117, 164, 219, 181, 48, 2, 209, 175, 44, 43,
	201, 44, 43, 39, 44, 43, 44, 43, 173, 63,
	63, 63, 63, 63, 63, 63, 63, 63, 63, 63,
	143, 144, 146, 145, 182, 147, 44, 43, 42, 171,
	149, 42, 92, 40, 42, 85, 42, 154, 220, 151,
	158, 163, 101, 151, 161, 155, 192, 167, 129, 188,
	170, 117, 184, 227, 176, 165, 42, 128, 150, 84,
	83, 41, 93, 85, 22, 50, 51, 52, 53, 54,
	45, 81, 82, 148, 168, 215, 156, 169, 130, 44,
	44, 43, 213, 179, 104, 44, 43, 195, 180, 49,
	41, 87, 95, 183, 3, 117, 142, 153, 187, 64,
	189, 21, 38, 193, 81, 82, 190, 69, 70, 71,
	42, 152, 196, 194, 18, 42, 167, 200, 197, 61,
	199, 81, 82, 172, 203, 206, 8, 86, 207, 88,
	89, 118, 66, 205, 37, 19, 90, 91, 67, 68,
	69, 70, 71, 167, 218, 217, 178, 105, 75, 77,
	17, 224, 108, 47, 81, 82, 159, 7, 225, 226,
	6, 4, 1, 0, 191, 36, 0, 0, 0, 131,
	132, 133, 134, 135, 136, 137, 138, 139, 140, 141,
	24, 29, 33, 30, 31, 32, 74, 76, 67, 68,
	69, 70, 71, 0, 27, 28, 0, 0, 0, 0,
	36, 106, 0, 0, 0, 107, 16, 0, 119, 102,
	109, 110, 111, 112, 113, 60, 115, 116, 114, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 34, 35,
	0, 22, 0, 25, 0, 20, 0, 160, 0, 157,
	24, 29, 33, 30, 31, 32, 0, 0, 0, 0,
	0, 0, 0, 0, 27, 28, 0, 0, 0, 0,
	0, 106, 0, 0, 0, 107, 16, 0, 119, 0,
	109, 110, 111, 112, 113, 60, 115, 116, 114, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 34, 35,
	0, 22, 0, 25, 0, 20, 0, 26, 24, 29,
	33, 30, 31, 32, 0, 0, 0, 0, 0, 0,
	0, 0, 27, 28, 0, 24, 29, 33, 30, 31,
	32, 10, 11, 0, 16, 0, 0, 0, 0, 27,
	28, 0, 0, 9, 0, 0, 0, 0, 0, 0,
	0, 16, 0, 0, 0, 0, 34, 35, 0, 22,
	60, 25, 0, 20, 0, 26, 0, 14, 0, 0,
	0, 0, 0, 34, 35, 0, 22, 0, 25, 0,
	20, 0, 26, 0, 162, 24, 29, 33, 30, 31,
	32, 0, 0, 0, 0, 0, 0, 0, 0, 27,
	28, 0, 24, 29, 33, 30, 31, 32, 0, 0,
	0, 16, 0, 0, 0, 0, 27, 28, 0, 0,
	60, 0, 0, 0, 0, 0, 0, 0, 62, 0,
	0, 0, 0, 34, 35, 0, 22, 60, 25, 0,
	20, 0, 26, 0, 0, 0, 0, 0, 0, 0,
	34, 35, 0, 22, 0, 25, 0, 20, 0, 26,
	186, 29, 33, 30, 31, 32, 0, 0, 0, 0,
	0, 0, 0, 0, 27, 28, 0, 24, 29, 33,
	30, 31, 32, 0, 0, 0, 16, 0, 0, 0,
	0, 27, 28, 0, 0, 60, 0, 0, 72, 73,
	75, 77, 80, 16, 0, 0, 81, 82, 34, 35,
	0, 22, 60, 25, 78, 20, 0, 26, 0, 0,
	0, 0, 0, 0, 0, 34, 35, 0, 22, 0,
	25, 0, 20, 0, 56, 0, 0, 79, 74, 76,
	67, 68, 69, 70, 71,
}
var yyPact = [...]int{

	-57, 364, -1000, -57, -1000, -1000, -1000, -50, -57, 99,
	-1000, -1000, 140, 93, -57, 115, 533, 458, -1000, -1000,
	458, 548, -1000, 69, 72, 458, 441, 458, 458, -1000,
	-1000, -1000, -1000, -1000, 458, 458, -1000, -57, -57, -1000,
	71, 158, 441, 441, 441, 79, -23, -57, 306, 441,
	441, 441, 441, 441, 441, -1000, 441, 80, 140, -1000,
	70, 548, 57, -1000, 117, 69, -1000, 458, 458, 458,
	458, 458, 458, 458, 458, 458, 458, 458, 162, 79,
	79, -1000, -1000, 441, 441, 441, 173, 81, -1000, -1000,
	173, 173, -1000, 158, 66, -1000, 135, 134, -1000, -1000,
	-1000, -1000, -57, -1000, 140, -1000, 441, 441, 113, 246,
	381, -2, -1000, -1000, -1000, 441, 441, 115, 100, 441,
	140, 140, 140, 140, 140, 140, 37, 81, -57, 441,
	-1000, 123, 123, 173, 173, 173, 206, 206, 156, 156,
	156, 156, -1000, -1000, -1000, 16, -33, 5, -1000, 62,
	-4, -57, 441, 306, 140, 80, 31, -57, 59, -48,
	516, -1000, -57, 56, -57, -1000, -1000, 140, 441, -57,
	53, -1000, 441, -1000, -1000, -1000, -1000, -57, 153, 140,
	-1000, 441, -57, -26, -57, 441, 44, -28, -57, -29,
	80, 458, -57, -1000, -30, -1000, 3, -34, -1000, -36,
	-49, 148, -1000, -39, 110, -1000, 69, -41, -1000, -57,
	-1000, -1000, 441, 0, -1000, 47, -1000, -42, -5, -21,
	441, -1000, -57, -57, 61, -43, -44, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 232, 230, 227, 4, 226, 196, 223, 5, 222,
	0, 220, 217, 171, 3, 11, 205, 1, 10, 7,
	201, 9, 6, 64, 164, 231, 184, 169,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 2, 3, 3, 3, 3,
	3, 6, 8, 8, 7, 7, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 9, 9, 9, 12, 21, 21,
	19, 19, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 11, 11, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 27, 18, 14, 14, 14, 14, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 20, 20, 15, 15, 16, 16, 22, 22,
	22, 5, 5, 17, 17, 23, 23, 24, 24, 25,
	26,
}
var yyR2 = [...]int{

	0, 1, 2, 3, 2, 2, 5, 1, 1, 1,
	3, 3, 0, 2, 2, 4, 1, 1, 2, 1,
	2, 1, 4, 5, 9, 4, 5, 8, 1, 1,
	1, 9, 2, 2, 5, 7, 5, 3, 0, 1,
	1, 4, 3, 3, 3, 3, 3, 3, 5, 3,
	3, 2, 1, 1, 2, 1, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	1, 2, 2, 0, 3, 2, 1, 4, 4, 4,
	7, 3, 2, 2, 1, 1, 1, 1, 1, 1,
	2, 2, 1, 4, 1, 2, 4, 1, 0, 1,
	4, 0, 1, 0, 1, 0, 1, 1, 2, 1,
	1,
}
var yyChk = [...]int{

	-1000, -1, -23, -24, -25, 65, -2, -3, -6, 39,
	27, 28, -10, -18, 63, -15, 30, -11, -26, -16,
	59, -13, 55, -14, 4, 57, 61, 18, 19, 5,
	7, 8, 9, 6, 52, 53, -25, -6, -24, -23,
	4, 61, 45, 16, 15, 47, -8, -7, -23, 44,
	20, 21, 22, 23, 24, -21, 61, -19, -10, -18,
	39, -13, 30, -15, -27, -14, -18, 52, 53, 54,
	55, 56, 10, 11, 50, 12, 51, 13, 26, 49,
	14, 18, 19, 61, 60, 61, -13, -10, -13, -13,
	-13, -13, -23, 61, -22, 4, -10, -10, -10, -18,
	64, -23, -25, -4, -10, -12, 25, 29, -9, 34,
	35, 36, 37, 38, 42, 40, 41, -15, -20, 32,
	-10, -10, -10, -10, -10, -10, -21, -10, 47, 61,
	31, -13, -13, -13, -13, -13, -13, -13, -13, -13,
	-13, -13, 4, -18, -18, -21, -19, -21, 62, -22,
	62, 47, 46, -23, -10, -19, 33, 63, -10, -5,
	61, -4, 63, -10, 63, -21, -17, -10, 44, 47,
	-10, 62, -23, 62, 66, 62, 62, 63, -23, -10,
	-4, 32, 63, -8, 63, 65, 4, -8, 63, -8,
	-19, -23, 63, -10, -8, 4, -10, -8, 64, -8,
	-17, 26, 64, -8, 64, -15, -14, -8, 64, 63,
	64, 64, 65, 4, 64, 35, 64, -8, -17, 62,
	61, 64, 63, 63, -10, -8, -8, 62, 64, 64,
}
var yyDef = [...]int{

	105, -2, 1, 106, 107, 109, 2, 0, 105, 0,
	7, 8, 9, -2, -2, 89, 38, 52, 73, 94,
	0, 53, 110, 55, 97, 0, 0, 0, 0, 84,
	85, 86, 87, 88, 0, 0, 108, 105, 4, 5,
	0, 98, 0, 0, 0, 0, 0, 105, 0, 0,
	0, 0, 0, 0, 0, 51, 38, 39, 40, 70,
	0, 54, 0, 89, 0, 95, 76, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 71, 72, 38, 0, 38, 75, 0, 82, 83,
	90, 91, 3, 98, 0, 99, 0, 49, 50, 10,
	11, 13, -2, 14, 16, 17, 0, 19, 21, 101,
	0, 0, 28, 29, 30, 38, 103, -2, 0, 0,
	42, 43, 44, 45, 46, 47, 0, 40, 105, 38,
	74, 56, 57, 58, 59, 60, 61, 62, 63, 64,
	65, 66, 67, 68, 69, 0, 0, 0, 81, 0,
	0, 105, 0, 0, 18, 20, 0, -2, 16, 0,
	0, 102, -2, 0, -2, 32, 33, 104, 0, 105,
	0, 78, 0, 79, 96, 77, 6, -2, 0, 48,
	15, 0, -2, 0, -2, 103, 97, 0, -2, 0,
	37, 0, -2, 41, 0, 100, 0, 0, 22, 0,
	0, 0, 25, 0, 0, -2, 0, 0, 80, -2,
	36, 23, 103, 0, 26, 0, 34, 0, 0, 0,
	0, 35, -2, -2, 0, 0, 0, 27, 24, 31,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 57, 3, 3, 59, 56, 3, 3,
	61, 62, 54, 52, 47, 53, 3, 55, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 46, 65,
	51, 44, 50, 45, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 60, 3, 66, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 63, 3, 64, 49,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 48, 58,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:82
		{
			yyVAL.rules = []ast.Rule{}
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:86
		{
			yyVAL.rules = append(yyDollar[1].rules, yyDollar[2].rule)
			yylex.(*Lexer).result = yyVAL.rules
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:94
		{
			yyVAL.rule = ast.Rule{Pattern: yyDollar[1].pattern, Action: yyDollar[2].stmts}
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:98
		{
			yyVAL.rule = ast.Rule{Pattern: yyDollar[1].pattern, Action: []ast.Stmt{&ast.PrintStmt{Exprs: defaultExprs}}}
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:102
		{
			yyVAL.rule = ast.Rule{Pattern: &ast.ExprPattern{}, Action: yyDollar[1].stmts}
		}
	case 6:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/grammar.go.y:114
		{
			//fmt.Println("FUNC RULE")
			yyVAL.pattern = &ast.FuncPattern{Name: yyDollar[2].token.Literal, Args: yyDollar[4].ident_args}
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:119
		{
			yyVAL.pattern = &ast.BeginPattern{}
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:123
		{
			yyVAL.pattern = &ast.EndPattern{}
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:127
		{
			yyVAL.pattern = &ast.ExprPattern{Expr: yyDollar[1].expr}
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:131
		{
			yyVAL.pattern = &ast.StartStopPattern{
				Start: &ast.MatchExpr{Expr: &defaultExpr, RegExpr: yyDollar[1].expr},
				Stop:  &ast.MatchExpr{Expr: &defaultExpr, RegExpr: yyDollar[3].expr},
			}
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:140
		{
			yyVAL.stmts = yyDollar[2].stmts
		}
	case 12:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/grammar.go.y:146
		{
			yyVAL.stmts = []ast.Stmt{}
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:150
		{
			yyVAL.stmts = yyDollar[1].stmts
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:156
		{
			yyVAL.stmts = []ast.Stmt{yyDollar[2].stmt}
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:160
		{
			yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[4].stmt)
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:166
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyDollar[1].expr}
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:170
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyDollar[1].expr}
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:174
		{
			yyVAL.stmt = &ast.DelStmt{Expr: yyDollar[2].expr}
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:178
		{
			yyVAL.stmt = &ast.PrintStmt{Exprs: defaultExprs}
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:182
		{
			yyVAL.stmt = &ast.PrintStmt{Exprs: yyDollar[2].exprs}
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:186
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 22:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:190
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyDollar[3].stmts}
		}
	case 23:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/grammar.go.y:194
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyDollar[4].stmts, Expr: yyDollar[2].expr}
		}
	case 24:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line ./parser/grammar.go.y:198
		{
			yyVAL.stmt = &ast.CForLoopStmt{Stmt1: yyDollar[2].stmt, Expr2: yyDollar[4].expr, Expr3: yyDollar[6].expr, Stmts: yyDollar[8].stmts}
		}
	case 25:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:202
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyDollar[3].stmts}
		}
	case 26:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/grammar.go.y:206
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyDollar[4].stmts, Expr: yyDollar[2].expr}
		}
	case 27:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./parser/grammar.go.y:210
		{
			yyVAL.stmt = &ast.DoLoopStmt{Stmts: yyDollar[3].stmts, Expr: yyDollar[7].expr}
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:214
		{
			yyVAL.stmt = &ast.BreakStmt{}
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:218
		{
			yyVAL.stmt = &ast.ContinueStmt{}
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:222
		{
			yyVAL.stmt = &ast.NextStmt{}
		}
	case 31:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line ./parser/grammar.go.y:226
		{
			yyVAL.stmt = &ast.MapLoopStmt{KeyId: yyDollar[3].token.Literal, MapId: yyDollar[5].token.Literal, Stmts: yyDollar[8].stmts}
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:230
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyDollar[2].exprs}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:234
		{
			yyVAL.stmt = &ast.ExitStmt{Expr: yyDollar[2].expr}
		}
	case 34:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/grammar.go.y:240
		{
			yyVAL.stmt = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].stmts, Else: nil}
		}
	case 35:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./parser/grammar.go.y:244
		{
			yyVAL.stmt.(*ast.IfStmt).ElseIf = append(yyVAL.stmt.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyDollar[4].expr, Then: yyDollar[6].stmts})
		}
	case 36:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/grammar.go.y:248
		{
			if yyVAL.stmt.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				//$$.(*ast.IfStmt).Else = append($$.(*ast.IfStmt).Else, $4...)
				yyVAL.stmt.(*ast.IfStmt).Else = yyDollar[4].stmts
			}
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:259
		{
			yyVAL.expr = &ast.AssExpr{Left: yyDollar[1].exprs, Right: yyDollar[3].exprs}
		}
	case 38:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/grammar.go.y:265
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:269
		{
			yyVAL.exprs = yyDollar[1].exprs
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:275
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:279
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:285
		{
			yyVAL.expr = &ast.AssExpr{Left: []ast.Expr{yyDollar[1].expr}, Right: []ast.Expr{yyDollar[3].expr}}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:290
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "+=", Right: yyDollar[3].expr}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:294
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "-=", Right: yyDollar[3].expr}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:298
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "*=", Right: yyDollar[3].expr}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:302
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "/=", Right: yyDollar[3].expr}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:306
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "%=", Right: yyDollar[3].expr}
		}
	case 48:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/grammar.go.y:311
		{
			yyVAL.expr = &ast.TriOpExpr{Cond: yyDollar[1].expr, Then: yyDollar[3].expr, Else: yyDollar[5].expr}
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:316
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "||", Right: yyDollar[3].expr}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:320
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "&&", Right: yyDollar[3].expr}
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:324
		{
			yyVAL.expr = &ast.CallExpr{Name: "printf", SubExprs: yyDollar[2].exprs}
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:328
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:334
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:339
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "CAT", Right: yyDollar[2].expr}
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:345
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:350
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "+", Right: yyDollar[3].expr}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:354
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "-", Right: yyDollar[3].expr}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:358
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "*", Right: yyDollar[3].expr}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:362
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "/", Right: yyDollar[3].expr}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:366
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "%", Right: yyDollar[3].expr}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:371
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "==", Right: yyDollar[3].expr}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:375
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "!=", Right: yyDollar[3].expr}
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:379
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: ">", Right: yyDollar[3].expr}
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:383
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: ">=", Right: yyDollar[3].expr}
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:387
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "<", Right: yyDollar[3].expr}
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:391
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "<=", Right: yyDollar[3].expr}
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:395
		{
			yyVAL.expr = &ast.ContainKeyExpr{KeyExpr: yyDollar[1].expr, MapId: yyDollar[3].token.Literal}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:400
		{
			yyVAL.expr = &ast.MatchExpr{Expr: yyDollar[1].expr, RegExpr: yyDollar[3].expr}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:404
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: &ast.MatchExpr{Expr: yyDollar[1].expr, RegExpr: yyDollar[3].expr}}
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:408
		{
			yyVAL.expr = &ast.MatchExpr{Expr: &defaultExpr, RegExpr: yyDollar[1].expr}
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:413
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "++", After: true}
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:417
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "--", After: true}
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:423
		{
			//fmt.Println("YACC: want regexp!!")
			IN_REGEXP = true
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:428
		{
			//fmt.Println("FINISH")
			yyVAL.expr = &ast.RegExpr{Literal: yyDollar[3].token.Literal}
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:435
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:439
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 77:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:444
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].token.Literal, SubExprs: yyDollar[3].exprs}
		}
	case 78:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:448
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].token.Literal, SubExprs: yyDollar[3].exprs}
		}
	case 79:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:452
		{
			yyVAL.expr = &ast.AnonymousCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
		}
	case 80:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./parser/grammar.go.y:457
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyDollar[3].ident_args, Stmts: yyDollar[6].stmts}
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:462
		{
			yyVAL.expr = &ast.ParentExpr{SubExpr: yyDollar[2].expr}
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:467
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[2].expr, Operator: "++"}
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:471
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[2].expr, Operator: "--"}
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:476
		{
			yyVAL.expr = &ast.NumExpr{Literal: yyDollar[1].token.Literal}
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:480
		{
			yyVAL.expr = &ast.ConstExpr{Literal: yyDollar[1].token.Literal}
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:484
		{
			yyVAL.expr = &ast.ConstExpr{Literal: yyDollar[1].token.Literal}
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:488
		{
			yyVAL.expr = &ast.ConstExpr{Literal: yyDollar[1].token.Literal}
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:492
		{
			yyVAL.expr = &ast.StringExpr{Literal: yyDollar[1].token.Literal}
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:497
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:502
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "+", Expr: yyDollar[2].expr}
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:506
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
		}
	case 92:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:512
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 93:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:516
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:522
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 95:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:526
		{
			yyVAL.expr = &ast.FieldExpr{Expr: yyDollar[2].expr}
		}
	case 96:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:532
		{
			yyVAL.expr = &ast.ItemExpr{Expr: yyDollar[1].expr, Index: yyDollar[3].exprs}
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:536
		{
			yyVAL.expr = &ast.IdentExpr{Literal: yyDollar[1].token.Literal}
		}
	case 98:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/grammar.go.y:542
		{
			yyVAL.ident_args = []string{}
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:546
		{
			yyVAL.ident_args = []string{yyDollar[1].token.Literal}
		}
	case 100:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:550
		{
			yyVAL.ident_args = append(yyDollar[1].ident_args, yyDollar[4].token.Literal)
		}
	case 101:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/grammar.go.y:556
		{
			yyVAL.stmt = nil
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:560
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 103:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/grammar.go.y:566
		{
			yyVAL.expr = nil
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:570
		{
			yyVAL.expr = yyDollar[1].expr
		}
	}
	goto yystack /* stack new state and value */
}
