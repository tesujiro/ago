//line ./parser/grammar.go.y:2
package parser

import __yyfmt__ "fmt"

//line ./parser/grammar.go.y:2
import (
	//"fmt"
	"github.com/tesujiro/goa/ast"
)

var defaultExpr = ast.FieldExpr{Expr: &ast.NumExpr{Literal: "0"}}
var defaultExprs = []ast.Expr{&defaultExpr}

//var IN_REGEXP bool

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
const ANDAND = 57356
const OROR = 57357
const LEN = 57358
const PLUSPLUS = 57359
const MINUSMINUS = 57360
const PLUSEQ = 57361
const MINUSEQ = 57362
const MULEQ = 57363
const DIVEQ = 57364
const MODEQ = 57365
const DELETE = 57366
const IN = 57367
const BEGIN = 57368
const END = 57369
const PRINT = 57370
const PRINTF = 57371
const REGEXP = 57372
const IF = 57373
const ELSE = 57374
const FOR = 57375
const WHILE = 57376
const DO = 57377
const BREAK = 57378
const CONTINUE = 57379
const FUNC = 57380
const RETURN = 57381
const EXIT = 57382
const NEXT = 57383
const CONCAT_OP = 57384
const vars = 57385
const UNARY = 57386

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

//line ./parser/grammar.go.y:574

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 14,
	63, 12,
	-2, 101,
	-1, 97,
	63, 103,
	-2, 101,
	-1, 112,
	46, 88,
	-2, 85,
	-1, 150,
	63, 12,
	-2, 101,
	-1, 155,
	63, 12,
	-2, 101,
	-1, 157,
	63, 12,
	-2, 101,
	-1, 170,
	63, 12,
	-2, 101,
	-1, 175,
	63, 12,
	-2, 101,
	-1, 177,
	63, 12,
	-2, 101,
	-1, 181,
	63, 12,
	-2, 101,
	-1, 185,
	63, 12,
	-2, 101,
	-1, 198,
	59, 85,
	60, 85,
	-2, 89,
	-1, 202,
	63, 12,
	-2, 101,
	-1, 215,
	63, 12,
	-2, 101,
	-1, 216,
	63, 12,
	-2, 101,
}

const yyPrivate = 57344

const yyLast = 645

var yyAct = [...]int{

	56, 159, 12, 21, 89, 44, 55, 5, 98, 121,
	14, 15, 5, 123, 205, 178, 42, 41, 222, 221,
	214, 209, 207, 62, 204, 82, 53, 46, 2, 61,
	216, 61, 167, 203, 201, 61, 37, 61, 61, 42,
	41, 91, 92, 93, 61, 61, 40, 99, 115, 116,
	117, 118, 119, 120, 197, 122, 195, 191, 112, 42,
	41, 42, 41, 87, 202, 212, 95, 215, 170, 40,
	42, 41, 157, 96, 168, 61, 61, 61, 61, 61,
	61, 61, 61, 61, 61, 61, 139, 185, 138, 40,
	140, 40, 194, 142, 166, 174, 68, 69, 71, 73,
	40, 164, 147, 76, 77, 151, 156, 181, 220, 148,
	213, 74, 160, 154, 144, 163, 112, 144, 177, 124,
	158, 38, 42, 41, 123, 146, 175, 80, 161, 169,
	39, 162, 143, 88, 75, 70, 72, 63, 64, 65,
	66, 67, 79, 78, 80, 43, 172, 99, 76, 77,
	208, 165, 40, 149, 82, 173, 176, 137, 112, 94,
	42, 180, 3, 182, 42, 41, 186, 206, 183, 141,
	36, 20, 171, 188, 4, 189, 187, 39, 34, 160,
	193, 190, 90, 192, 65, 66, 67, 196, 199, 59,
	184, 200, 76, 77, 40, 81, 198, 83, 84, 48,
	49, 50, 51, 52, 85, 86, 160, 211, 210, 76,
	77, 34, 136, 113, 217, 18, 100, 17, 103, 45,
	97, 218, 219, 47, 42, 41, 63, 64, 65, 66,
	67, 152, 7, 6, 1, 125, 126, 127, 128, 129,
	130, 131, 132, 133, 134, 135, 22, 27, 31, 28,
	29, 30, 0, 8, 40, 145, 0, 0, 0, 25,
	26, 35, 0, 0, 0, 0, 101, 0, 0, 0,
	102, 16, 57, 114, 0, 104, 105, 106, 107, 108,
	58, 110, 111, 109, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 32, 33, 0, 0, 0, 23, 0,
	19, 0, 153, 0, 150, 22, 27, 31, 28, 29,
	30, 0, 0, 0, 0, 0, 0, 0, 25, 26,
	0, 0, 0, 0, 0, 101, 0, 0, 0, 102,
	16, 57, 114, 0, 104, 105, 106, 107, 108, 58,
	110, 111, 109, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 32, 33, 0, 0, 0, 23, 0, 19,
	0, 24, 22, 27, 31, 28, 29, 30, 0, 0,
	0, 0, 0, 0, 0, 25, 26, 22, 27, 31,
	28, 29, 30, 0, 10, 11, 0, 16, 13, 0,
	25, 26, 0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 16, 57, 0, 0, 0, 0, 0, 32,
	33, 58, 0, 0, 23, 0, 19, 0, 24, 0,
	14, 0, 0, 0, 32, 33, 0, 0, 0, 23,
	0, 19, 0, 24, 0, 155, 22, 27, 31, 28,
	29, 30, 0, 0, 0, 0, 0, 0, 0, 25,
	26, 179, 27, 31, 28, 29, 30, 0, 0, 0,
	0, 16, 57, 0, 25, 26, 0, 0, 0, 0,
	58, 0, 0, 0, 0, 0, 16, 57, 0, 0,
	0, 0, 0, 32, 33, 58, 0, 0, 23, 0,
	19, 0, 24, 0, 0, 0, 0, 0, 32, 33,
	0, 0, 0, 23, 0, 19, 0, 24, 22, 27,
	31, 28, 29, 30, 0, 0, 0, 0, 0, 0,
	0, 25, 26, 22, 27, 31, 28, 29, 30, 0,
	0, 0, 0, 60, 57, 0, 25, 26, 0, 0,
	0, 0, 58, 0, 0, 0, 0, 0, 16, 57,
	0, 0, 0, 0, 0, 32, 33, 58, 0, 0,
	23, 0, 19, 0, 24, 0, 0, 0, 0, 0,
	32, 33, 0, 0, 0, 23, 0, 19, 0, 54,
	22, 27, 31, 28, 29, 30, 0, 0, 0, 0,
	0, 0, 0, 25, 26, 0, 0, 0, 0, 0,
	0, 71, 73, 0, 0, 60, 76, 77, 0, 0,
	0, 0, 0, 0, 58, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 32, 33, 0,
	0, 0, 23, 0, 19, 0, 24, 0, 70, 72,
	63, 64, 65, 66, 67,
}
var yyPact = [...]int{

	-57, 358, -1000, -57, -1000, -1000, -1000, -52, -57, 117,
	-1000, -1000, 150, 99, -57, 180, 519, 504, -1000, 576,
	86, 83, 84, 504, 432, 504, 504, -1000, -1000, -1000,
	-1000, -1000, 504, 504, -1000, -57, -57, -1000, 73, 178,
	432, 432, 432, 129, 3, -57, 301, 432, 432, 432,
	432, 432, 432, -1000, 432, 78, 150, -1000, 70, 86,
	59, -1000, 83, 504, 504, 504, 504, 504, 504, 504,
	504, 504, 504, 504, 208, 127, -1000, -1000, 432, 432,
	432, 192, 108, -1000, -1000, 192, 192, -1000, 178, 71,
	-1000, 210, 146, -1000, -1000, -1000, -1000, -57, -1000, 150,
	-1000, 432, 432, 121, 242, 373, 10, -1000, -1000, -1000,
	432, 432, 180, 85, 432, 150, 150, 150, 150, 150,
	150, 40, 108, -57, 432, 131, 131, 192, 192, 192,
	589, 589, 175, 175, 175, 175, -1000, -1000, 33, -33,
	13, -1000, 68, 6, -57, 432, 301, 150, 78, 64,
	-57, 56, -49, 447, -1000, -57, 45, -57, -1000, -1000,
	150, 432, -57, 25, -1000, 432, -1000, -1000, -1000, -1000,
	-57, 169, 150, -1000, 432, -57, -6, -57, 432, 67,
	-7, -57, -9, 78, 576, -57, -1000, -29, -1000, 2,
	-30, -1000, -39, -50, 163, -1000, -41, 116, -1000, 83,
	-42, -1000, -57, -1000, -1000, 432, 4, -1000, 50, -1000,
	-43, 5, -32, 432, -1000, -57, -57, 47, -44, -45,
	-1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 234, 233, 232, 8, 231, 253, 219, 5, 218,
	0, 217, 216, 171, 3, 11, 215, 1, 6, 213,
	9, 4, 27, 162, 174,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 2, 3, 3, 3, 3,
	3, 6, 8, 8, 7, 7, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 9, 9, 9, 12, 20, 20,
	18, 18, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 11, 11, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 19, 19,
	16, 16, 15, 15, 21, 21, 21, 5, 5, 17,
	17, 22, 22, 23, 23, 24,
}
var yyR2 = [...]int{

	0, 1, 2, 3, 2, 2, 5, 1, 1, 1,
	3, 3, 0, 2, 2, 4, 1, 1, 2, 1,
	2, 1, 4, 5, 9, 4, 5, 8, 1, 1,
	1, 9, 2, 2, 5, 7, 5, 3, 0, 1,
	1, 4, 3, 3, 3, 3, 3, 3, 5, 3,
	3, 2, 1, 1, 2, 1, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 1,
	2, 2, 2, 4, 4, 4, 7, 3, 2, 2,
	1, 1, 1, 1, 1, 1, 2, 2, 1, 4,
	4, 1, 1, 2, 0, 1, 4, 0, 1, 0,
	1, 0, 1, 1, 2, 1,
}
var yyChk = [...]int{

	-1000, -1, -22, -23, -24, 64, -2, -3, -6, 38,
	26, 27, -10, 30, 62, -15, 29, -11, -16, 58,
	-13, -14, 4, 56, 60, 17, 18, 5, 7, 8,
	9, 6, 51, 52, -24, -6, -23, -22, 4, 60,
	44, 15, 14, 46, -8, -7, -22, 43, 19, 20,
	21, 22, 23, -20, 60, -18, -10, 30, 38, -13,
	29, -15, -14, 51, 52, 53, 54, 55, 10, 11,
	49, 12, 50, 13, 25, 48, 17, 18, 60, 59,
	60, -13, -10, -13, -13, -13, -13, -22, 60, -21,
	4, -10, -10, -10, 30, 63, -22, -24, -4, -10,
	-12, 24, 28, -9, 33, 34, 35, 36, 37, 41,
	39, 40, -15, -19, 31, -10, -10, -10, -10, -10,
	-10, -20, -10, 46, 60, -13, -13, -13, -13, -13,
	-13, -13, -13, -13, -13, -13, 4, 30, -20, -18,
	-20, 61, -21, 61, 46, 45, -22, -10, -18, 32,
	62, -10, -5, 60, -4, 62, -10, 62, -20, -17,
	-10, 43, 46, -10, 61, -22, 61, 65, 61, 61,
	62, -22, -10, -4, 31, 62, -8, 62, 64, 4,
	-8, 62, -8, -18, -22, 62, -10, -8, 4, -10,
	-8, 63, -8, -17, 25, 63, -8, 63, -15, -14,
	-8, 63, 62, 63, 63, 64, 4, 63, 34, 63,
	-8, -17, 61, 60, 63, 62, 62, -10, -8, -8,
	61, 63, 63,
}
var yyDef = [...]int{

	101, -2, 1, 102, 103, 105, 2, 0, 101, 0,
	7, 8, 9, 69, -2, 85, 38, 52, 92, 0,
	53, 55, 91, 0, 0, 0, 0, 80, 81, 82,
	83, 84, 0, 0, 104, 101, 4, 5, 0, 94,
	0, 0, 0, 0, 0, 101, 0, 0, 0, 0,
	0, 0, 0, 51, 38, 39, 40, 69, 0, 54,
	0, 85, 93, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 70, 71, 38, 0,
	38, 72, 0, 78, 79, 86, 87, 3, 94, 0,
	95, 0, 49, 50, 10, 11, 13, -2, 14, 16,
	17, 0, 19, 21, 97, 0, 0, 28, 29, 30,
	38, 99, -2, 0, 0, 42, 43, 44, 45, 46,
	47, 0, 40, 101, 38, 56, 57, 58, 59, 60,
	61, 62, 63, 64, 65, 66, 67, 68, 0, 0,
	0, 77, 0, 0, 101, 0, 0, 18, 20, 0,
	-2, 16, 0, 0, 98, -2, 0, -2, 32, 33,
	100, 0, 101, 0, 74, 0, 75, 90, 73, 6,
	-2, 0, 48, 15, 0, -2, 0, -2, 99, 91,
	0, -2, 0, 37, 0, -2, 41, 0, 96, 0,
	0, 22, 0, 0, 0, 25, 0, 0, -2, 0,
	0, 76, -2, 36, 23, 99, 0, 26, 0, 34,
	0, 0, 0, 0, 35, -2, -2, 0, 0, 0,
	27, 24, 31,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 56, 3, 3, 58, 55, 3, 3,
	60, 61, 53, 51, 46, 52, 3, 54, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 45, 64,
	50, 43, 49, 44, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 59, 3, 65, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 62, 3, 63, 48,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 47, 57,
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
		//line ./parser/grammar.go.y:81
		{
			yyVAL.rules = []ast.Rule{}
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:85
		{
			yyVAL.rules = append(yyDollar[1].rules, yyDollar[2].rule)
			yylex.(*Lexer).result = yyVAL.rules
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:93
		{
			yyVAL.rule = ast.Rule{Pattern: yyDollar[1].pattern, Action: yyDollar[2].stmts}
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:97
		{
			yyVAL.rule = ast.Rule{Pattern: yyDollar[1].pattern, Action: []ast.Stmt{&ast.PrintStmt{Exprs: defaultExprs}}}
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:101
		{
			yyVAL.rule = ast.Rule{Pattern: &ast.ExprPattern{}, Action: yyDollar[1].stmts}
		}
	case 6:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/grammar.go.y:113
		{
			//fmt.Println("FUNC RULE")
			yyVAL.pattern = &ast.FuncPattern{Name: yyDollar[2].token.Literal, Args: yyDollar[4].ident_args}
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:118
		{
			yyVAL.pattern = &ast.BeginPattern{}
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:122
		{
			yyVAL.pattern = &ast.EndPattern{}
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:126
		{
			yyVAL.pattern = &ast.ExprPattern{Expr: yyDollar[1].expr}
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:130
		{
			yyVAL.pattern = &ast.StartStopPattern{
				Start: &ast.MatchExpr{Expr: &defaultExpr, RegExpr: yyDollar[1].token.Literal},
				Stop:  &ast.MatchExpr{Expr: &defaultExpr, RegExpr: yyDollar[3].token.Literal},
			}
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:139
		{
			yyVAL.stmts = yyDollar[2].stmts
		}
	case 12:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/grammar.go.y:145
		{
			yyVAL.stmts = []ast.Stmt{}
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:149
		{
			yyVAL.stmts = yyDollar[1].stmts
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:155
		{
			yyVAL.stmts = []ast.Stmt{yyDollar[2].stmt}
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:159
		{
			yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[4].stmt)
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:165
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyDollar[1].expr}
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:169
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyDollar[1].expr}
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:173
		{
			yyVAL.stmt = &ast.DelStmt{Expr: yyDollar[2].expr}
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:177
		{
			yyVAL.stmt = &ast.PrintStmt{Exprs: defaultExprs}
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:181
		{
			yyVAL.stmt = &ast.PrintStmt{Exprs: yyDollar[2].exprs}
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:185
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 22:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:189
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyDollar[3].stmts}
		}
	case 23:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/grammar.go.y:193
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyDollar[4].stmts, Expr: yyDollar[2].expr}
		}
	case 24:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line ./parser/grammar.go.y:197
		{
			yyVAL.stmt = &ast.CForLoopStmt{Stmt1: yyDollar[2].stmt, Expr2: yyDollar[4].expr, Expr3: yyDollar[6].expr, Stmts: yyDollar[8].stmts}
		}
	case 25:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:201
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyDollar[3].stmts}
		}
	case 26:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/grammar.go.y:205
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyDollar[4].stmts, Expr: yyDollar[2].expr}
		}
	case 27:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./parser/grammar.go.y:209
		{
			yyVAL.stmt = &ast.DoLoopStmt{Stmts: yyDollar[3].stmts, Expr: yyDollar[7].expr}
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:213
		{
			yyVAL.stmt = &ast.BreakStmt{}
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:217
		{
			yyVAL.stmt = &ast.ContinueStmt{}
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:221
		{
			yyVAL.stmt = &ast.NextStmt{}
		}
	case 31:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line ./parser/grammar.go.y:225
		{
			yyVAL.stmt = &ast.MapLoopStmt{KeyId: yyDollar[3].token.Literal, MapId: yyDollar[5].token.Literal, Stmts: yyDollar[8].stmts}
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:229
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyDollar[2].exprs}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:233
		{
			yyVAL.stmt = &ast.ExitStmt{Expr: yyDollar[2].expr}
		}
	case 34:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/grammar.go.y:239
		{
			yyVAL.stmt = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].stmts, Else: nil}
		}
	case 35:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./parser/grammar.go.y:243
		{
			yyVAL.stmt.(*ast.IfStmt).ElseIf = append(yyVAL.stmt.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyDollar[4].expr, Then: yyDollar[6].stmts})
		}
	case 36:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/grammar.go.y:247
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
		//line ./parser/grammar.go.y:258
		{
			yyVAL.expr = &ast.AssExpr{Left: yyDollar[1].exprs, Right: yyDollar[3].exprs}
		}
	case 38:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/grammar.go.y:264
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:268
		{
			yyVAL.exprs = yyDollar[1].exprs
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:274
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:278
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:284
		{
			yyVAL.expr = &ast.AssExpr{Left: []ast.Expr{yyDollar[1].expr}, Right: []ast.Expr{yyDollar[3].expr}}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:289
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "+=", Right: yyDollar[3].expr}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:293
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "-=", Right: yyDollar[3].expr}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:297
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "*=", Right: yyDollar[3].expr}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:301
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "/=", Right: yyDollar[3].expr}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:305
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "%=", Right: yyDollar[3].expr}
		}
	case 48:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/grammar.go.y:310
		{
			yyVAL.expr = &ast.TriOpExpr{Cond: yyDollar[1].expr, Then: yyDollar[3].expr, Else: yyDollar[5].expr}
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:315
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "||", Right: yyDollar[3].expr}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:319
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "&&", Right: yyDollar[3].expr}
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:323
		{
			yyVAL.expr = &ast.CallExpr{Name: "printf", SubExprs: yyDollar[2].exprs}
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:327
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:333
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:338
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "CAT", Right: yyDollar[2].expr}
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:344
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:349
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "+", Right: yyDollar[3].expr}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:353
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "-", Right: yyDollar[3].expr}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:357
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "*", Right: yyDollar[3].expr}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:361
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "/", Right: yyDollar[3].expr}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:365
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "%", Right: yyDollar[3].expr}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:370
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "==", Right: yyDollar[3].expr}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:374
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "!=", Right: yyDollar[3].expr}
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:378
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: ">", Right: yyDollar[3].expr}
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:382
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: ">=", Right: yyDollar[3].expr}
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:386
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "<", Right: yyDollar[3].expr}
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:390
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "<=", Right: yyDollar[3].expr}
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:394
		{
			yyVAL.expr = &ast.ContainKeyExpr{KeyExpr: yyDollar[1].expr, MapId: yyDollar[3].token.Literal}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:399
		{
			yyVAL.expr = &ast.MatchExpr{Expr: yyDollar[1].expr, RegExpr: yyDollar[3].token.Literal}
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:403
		{
			yyVAL.expr = &ast.MatchExpr{Expr: &defaultExpr, RegExpr: yyDollar[1].token.Literal}
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:408
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "++", After: true}
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:412
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "--", After: true}
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:418
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
		}
	case 73:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:423
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].token.Literal, SubExprs: yyDollar[3].exprs}
		}
	case 74:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:427
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].token.Literal, SubExprs: yyDollar[3].exprs}
		}
	case 75:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:431
		{
			yyVAL.expr = &ast.AnonymousCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
		}
	case 76:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./parser/grammar.go.y:436
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyDollar[3].ident_args, Stmts: yyDollar[6].stmts}
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:441
		{
			yyVAL.expr = &ast.ParentExpr{SubExpr: yyDollar[2].expr}
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:446
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[2].expr, Operator: "++"}
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:450
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[2].expr, Operator: "--"}
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:455
		{
			yyVAL.expr = &ast.NumExpr{Literal: yyDollar[1].token.Literal}
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:459
		{
			yyVAL.expr = &ast.ConstExpr{Literal: yyDollar[1].token.Literal}
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:463
		{
			yyVAL.expr = &ast.ConstExpr{Literal: yyDollar[1].token.Literal}
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:467
		{
			yyVAL.expr = &ast.ConstExpr{Literal: yyDollar[1].token.Literal}
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:471
		{
			yyVAL.expr = &ast.StringExpr{Literal: yyDollar[1].token.Literal}
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:476
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:481
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "+", Expr: yyDollar[2].expr}
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:485
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:491
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 89:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:495
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 90:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:501
		{
			yyVAL.expr = &ast.ItemExpr{Expr: yyDollar[1].expr, Index: yyDollar[3].exprs}
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:505
		{
			yyVAL.expr = &ast.IdentExpr{Literal: yyDollar[1].token.Literal}
		}
	case 92:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:511
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 93:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:515
		{
			yyVAL.expr = &ast.FieldExpr{Expr: yyDollar[2].expr}
		}
	case 94:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/grammar.go.y:521
		{
			yyVAL.ident_args = []string{}
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:525
		{
			yyVAL.ident_args = []string{yyDollar[1].token.Literal}
		}
	case 96:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:529
		{
			yyVAL.ident_args = append(yyDollar[1].ident_args, yyDollar[4].token.Literal)
		}
	case 97:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/grammar.go.y:535
		{
			yyVAL.stmt = nil
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:539
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 99:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/grammar.go.y:545
		{
			yyVAL.expr = nil
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:549
		{
			yyVAL.expr = yyDollar[1].expr
		}
	}
	goto yystack /* stack new state and value */
}
