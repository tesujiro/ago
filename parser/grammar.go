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
const GETLINE = 57386
const vars = 57387
const UNARY = 57388

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
	"GETLINE",
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

//line ./parser/grammar.go.y:626

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 13,
	61, 79,
	-2, 70,
	-1, 14,
	65, 12,
	-2, 110,
	-1, 88,
	61, 92,
	62, 92,
	-2, 100,
	-1, 106,
	65, 112,
	-2, 110,
	-1, 121,
	48, 95,
	-2, 92,
	-1, 163,
	65, 12,
	-2, 110,
	-1, 168,
	65, 12,
	-2, 110,
	-1, 170,
	65, 12,
	-2, 110,
	-1, 184,
	65, 12,
	-2, 110,
	-1, 189,
	65, 12,
	-2, 110,
	-1, 191,
	65, 12,
	-2, 110,
	-1, 195,
	65, 12,
	-2, 110,
	-1, 199,
	65, 12,
	-2, 110,
	-1, 212,
	61, 92,
	62, 92,
	-2, 96,
	-1, 215,
	65, 12,
	-2, 110,
	-1, 228,
	65, 12,
	-2, 110,
	-1, 229,
	65, 12,
	-2, 110,
}

const yyPrivate = 57344

const yyLast = 709

var yyAct = [...]int{

	59, 172, 12, 107, 23, 47, 98, 132, 5, 218,
	60, 15, 13, 14, 192, 5, 235, 58, 45, 44,
	234, 227, 222, 220, 229, 66, 180, 217, 91, 64,
	89, 67, 64, 216, 214, 211, 67, 88, 64, 209,
	64, 64, 205, 104, 100, 101, 102, 64, 64, 43,
	108, 124, 125, 126, 127, 128, 129, 103, 131, 228,
	130, 121, 184, 170, 45, 44, 188, 215, 225, 45,
	44, 208, 181, 179, 45, 44, 177, 56, 45, 44,
	64, 64, 64, 64, 64, 64, 64, 64, 64, 64,
	64, 147, 148, 49, 2, 43, 45, 44, 189, 41,
	43, 157, 40, 150, 155, 43, 226, 86, 133, 43,
	42, 160, 97, 199, 164, 169, 183, 167, 195, 157,
	86, 173, 22, 191, 176, 121, 233, 43, 153, 161,
	85, 84, 96, 4, 156, 174, 132, 37, 175, 45,
	44, 46, 105, 221, 154, 149, 162, 151, 134, 45,
	51, 52, 53, 54, 55, 82, 83, 42, 3, 186,
	108, 219, 202, 187, 8, 64, 39, 91, 99, 190,
	43, 121, 38, 37, 194, 50, 196, 146, 65, 200,
	171, 18, 106, 122, 152, 19, 21, 87, 109, 203,
	201, 17, 197, 173, 207, 204, 112, 206, 48, 165,
	159, 210, 7, 89, 62, 213, 6, 1, 0, 67,
	212, 82, 83, 90, 0, 92, 93, 0, 0, 173,
	224, 223, 94, 95, 0, 0, 178, 230, 0, 82,
	83, 0, 0, 0, 231, 232, 73, 74, 76, 78,
	81, 0, 0, 0, 82, 83, 68, 69, 70, 71,
	72, 185, 79, 0, 0, 135, 136, 137, 138, 139,
	140, 141, 142, 143, 144, 145, 70, 71, 72, 198,
	24, 30, 34, 31, 32, 33, 80, 75, 77, 68,
	69, 70, 71, 72, 28, 29, 45, 44, 0, 0,
	0, 110, 0, 0, 0, 111, 16, 0, 123, 0,
	113, 114, 115, 116, 117, 61, 119, 120, 118, 0,
	25, 0, 0, 0, 0, 0, 0, 43, 158, 35,
	36, 0, 22, 0, 26, 0, 20, 0, 166, 0,
	163, 24, 30, 34, 31, 32, 33, 0, 0, 0,
	182, 0, 0, 0, 0, 28, 29, 0, 0, 0,
	0, 0, 110, 0, 0, 0, 111, 16, 0, 123,
	0, 113, 114, 115, 116, 117, 61, 119, 120, 118,
	0, 25, 0, 0, 0, 0, 0, 0, 0, 0,
	35, 36, 0, 22, 0, 26, 0, 20, 0, 27,
	24, 30, 34, 31, 32, 33, 0, 0, 0, 0,
	0, 0, 0, 0, 28, 29, 0, 0, 0, 0,
	0, 0, 0, 10, 11, 0, 16, 24, 30, 34,
	31, 32, 33, 0, 0, 9, 0, 0, 0, 0,
	25, 28, 29, 0, 0, 0, 0, 0, 0, 35,
	36, 0, 22, 16, 26, 0, 20, 0, 27, 0,
	14, 0, 61, 0, 0, 76, 78, 25, 0, 0,
	0, 82, 83, 0, 0, 0, 35, 36, 0, 22,
	0, 26, 0, 20, 0, 27, 0, 168, 24, 30,
	34, 31, 32, 33, 0, 0, 0, 0, 0, 0,
	0, 0, 28, 29, 75, 77, 68, 69, 70, 71,
	72, 0, 0, 0, 16, 193, 30, 34, 31, 32,
	33, 0, 0, 61, 0, 0, 0, 0, 25, 28,
	29, 0, 0, 0, 0, 0, 0, 35, 36, 0,
	22, 16, 26, 0, 20, 0, 27, 0, 0, 0,
	61, 0, 0, 0, 0, 25, 0, 0, 0, 0,
	0, 0, 0, 0, 35, 36, 0, 22, 0, 26,
	0, 20, 0, 27, 24, 30, 34, 31, 32, 33,
	0, 0, 0, 0, 0, 0, 0, 0, 28, 29,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	63, 24, 30, 34, 31, 32, 33, 0, 0, 61,
	0, 0, 0, 0, 25, 28, 29, 0, 0, 0,
	0, 0, 0, 35, 36, 0, 22, 16, 26, 0,
	20, 0, 27, 0, 0, 0, 61, 0, 0, 0,
	0, 25, 0, 0, 0, 0, 0, 0, 0, 0,
	35, 36, 0, 22, 0, 26, 0, 20, 0, 57,
	24, 30, 34, 31, 32, 33, 0, 0, 0, 0,
	0, 0, 0, 0, 28, 29, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 63, 0, 0, 0,
	0, 0, 0, 0, 0, 61, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 35,
	36, 0, 22, 0, 26, 0, 20, 0, 27,
}
var yyPact = [...]int{

	-58, 386, -1000, -58, -1000, -1000, -1000, -51, -58, 95,
	-1000, -1000, 124, 93, -58, 130, 587, 560, -1000, -1000,
	646, 226, -1000, 69, 58, 646, 560, 474, 560, 560,
	-1000, -1000, -1000, -1000, -1000, 560, 560, -1000, -58, -58,
	-1000, 50, 164, 474, 474, 474, 66, -22, -58, 327,
	474, 474, 474, 474, 474, 474, -1000, 474, 88, 124,
	-1000, 48, 226, 46, -1000, 117, 69, -1000, 560, 560,
	560, 560, 560, 560, 560, 560, 560, 560, 560, 173,
	66, 66, -1000, -1000, 474, 474, 474, 76, -1000, 69,
	137, 81, -1000, -1000, 137, 137, -1000, 164, 71, -1000,
	271, 134, -1000, -1000, -1000, -1000, -58, -1000, 124, -1000,
	474, 474, 113, 266, 413, -1, -1000, -1000, -1000, 474,
	474, 130, 90, 474, 124, 124, 124, 124, 124, 124,
	13, 81, -58, 474, -1000, 211, 211, 137, 137, 137,
	443, 443, 193, 193, 193, 193, -1000, -1000, -1000, 10,
	-41, 9, -1000, 560, -1000, 53, -2, -58, 474, 327,
	124, 88, 34, -58, 59, -52, 501, -1000, -58, 54,
	-58, -1000, -1000, 124, 474, -58, 49, -1000, 474, -1000,
	-1000, -1000, 193, -1000, -58, 158, 124, -1000, 474, -58,
	-23, -58, 474, 45, -26, -58, -30, 88, 646, -58,
	-1000, -31, -1000, 3, -32, -1000, -38, -57, 157, -1000,
	-42, 108, -1000, -43, -1000, -58, -1000, -1000, 474, 5,
	-1000, 44, -1000, -44, -5, -40, 474, -1000, -58, -58,
	63, -45, -49, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 207, 206, 202, 3, 199, 164, 198, 5, 196,
	0, 191, 188, 186, 4, 187, 11, 185, 1, 10,
	184, 17, 183, 60, 6, 93, 158, 133, 181, 178,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 2, 3, 3, 3, 3,
	3, 6, 8, 8, 7, 7, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 9, 9, 9, 12, 23, 23,
	21, 21, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 11, 11, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 29, 19, 20, 20, 14, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 14, 14, 14, 22, 22, 17, 17, 15,
	15, 16, 16, 24, 24, 24, 5, 5, 18, 18,
	25, 25, 26, 26, 27, 28,
}
var yyR2 = [...]int{

	0, 1, 2, 3, 2, 2, 5, 1, 1, 1,
	3, 3, 0, 2, 2, 4, 1, 1, 2, 1,
	2, 1, 4, 5, 9, 4, 5, 8, 1, 1,
	1, 9, 2, 2, 5, 7, 5, 3, 0, 1,
	1, 4, 3, 3, 3, 3, 3, 3, 5, 3,
	3, 2, 1, 1, 2, 1, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	1, 2, 2, 3, 0, 3, 0, 2, 2, 1,
	4, 4, 4, 7, 3, 2, 2, 1, 1, 1,
	1, 1, 1, 2, 2, 1, 4, 4, 1, 0,
	1, 1, 2, 0, 1, 4, 0, 1, 0, 1,
	0, 1, 1, 2, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -25, -26, -27, 66, -2, -3, -6, 39,
	27, 28, -10, -19, 64, -16, 30, -11, -28, -17,
	60, -13, 56, -14, 4, 44, 58, 62, 18, 19,
	5, 7, 8, 9, 6, 53, 54, -27, -6, -26,
	-25, 4, 62, 46, 16, 15, 48, -8, -7, -25,
	45, 20, 21, 22, 23, 24, -23, 62, -21, -10,
	-19, 39, -13, 30, -16, -29, -14, -19, 53, 54,
	55, 56, 57, 10, 11, 51, 12, 52, 13, 26,
	50, 14, 18, 19, 62, 61, 62, -15, -16, -14,
	-13, -10, -13, -13, -13, -13, -25, 62, -24, 4,
	-10, -10, -10, -19, 65, -25, -27, -4, -10, -12,
	25, 29, -9, 34, 35, 36, 37, 38, 42, 40,
	41, -16, -22, 32, -10, -10, -10, -10, -10, -10,
	-23, -10, 48, 62, 31, -13, -13, -13, -13, -13,
	-13, -13, -13, -13, -13, -13, 4, -19, -19, -23,
	-21, -23, -20, 52, 63, -24, 63, 48, 47, -25,
	-10, -21, 33, 64, -10, -5, 62, -4, 64, -10,
	64, -23, -18, -10, 45, 48, -10, 63, -25, 63,
	67, 63, -13, 63, 64, -25, -10, -4, 32, 64,
	-8, 64, 66, 4, -8, 64, -8, -21, -25, 64,
	-10, -8, 4, -10, -8, 65, -8, -18, 26, 65,
	-8, 65, -16, -8, 65, 64, 65, 65, 66, 4,
	65, 35, 65, -8, -18, 63, 62, 65, 64, 64,
	-10, -8, -8, 63, 65, 65,
}
var yyDef = [...]int{

	110, -2, 1, 111, 112, 114, 2, 0, 110, 0,
	7, 8, 9, -2, -2, 92, 38, 52, 74, 101,
	0, 53, 115, 55, 98, 99, 0, 0, 0, 0,
	87, 88, 89, 90, 91, 0, 0, 113, 110, 4,
	5, 0, 103, 0, 0, 0, 0, 0, 110, 0,
	0, 0, 0, 0, 0, 0, 51, 38, 39, 40,
	70, 0, 54, 0, 92, 0, 102, 79, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 71, 72, 38, 0, 38, 76, -2, 0,
	78, 0, 85, 86, 93, 94, 3, 103, 0, 104,
	0, 49, 50, 10, 11, 13, -2, 14, 16, 17,
	0, 19, 21, 106, 0, 0, 28, 29, 30, 38,
	108, -2, 0, 0, 42, 43, 44, 45, 46, 47,
	0, 40, 110, 38, 75, 56, 57, 58, 59, 60,
	61, 62, 63, 64, 65, 66, 67, 68, 69, 0,
	0, 0, 73, 0, 84, 0, 0, 110, 0, 0,
	18, 20, 0, -2, 16, 0, 0, 107, -2, 0,
	-2, 32, 33, 109, 0, 110, 0, 81, 0, 82,
	97, 80, 77, 6, -2, 0, 48, 15, 0, -2,
	0, -2, 108, 98, 0, -2, 0, 37, 0, -2,
	41, 0, 105, 0, 0, 22, 0, 0, 0, 25,
	0, 0, -2, 0, 83, -2, 36, 23, 108, 0,
	26, 0, 34, 0, 0, 0, 0, 35, -2, -2,
	0, 0, 0, 27, 24, 31,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 58, 3, 3, 60, 57, 3, 3,
	62, 63, 55, 53, 48, 54, 3, 56, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 47, 66,
	52, 45, 51, 46, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 61, 3, 67, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 64, 3, 65, 50,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 49, 59,
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
		//line ./parser/grammar.go.y:85
		{
			yyVAL.rules = []ast.Rule{}
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:89
		{
			yyVAL.rules = append(yyDollar[1].rules, yyDollar[2].rule)
			yylex.(*Lexer).result = yyVAL.rules
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:97
		{
			yyVAL.rule = ast.Rule{Pattern: yyDollar[1].pattern, Action: yyDollar[2].stmts}
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:101
		{
			yyVAL.rule = ast.Rule{Pattern: yyDollar[1].pattern, Action: []ast.Stmt{&ast.PrintStmt{Exprs: defaultExprs}}}
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:105
		{
			yyVAL.rule = ast.Rule{Pattern: &ast.ExprPattern{}, Action: yyDollar[1].stmts}
		}
	case 6:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/grammar.go.y:117
		{
			//fmt.Println("FUNC RULE")
			yyVAL.pattern = &ast.FuncPattern{Name: yyDollar[2].token.Literal, Args: yyDollar[4].ident_args}
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:122
		{
			yyVAL.pattern = &ast.BeginPattern{}
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:126
		{
			yyVAL.pattern = &ast.EndPattern{}
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:130
		{
			yyVAL.pattern = &ast.ExprPattern{Expr: yyDollar[1].expr}
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:134
		{
			yyVAL.pattern = &ast.StartStopPattern{
				Start: &ast.MatchExpr{Expr: &defaultExpr, RegExpr: yyDollar[1].expr},
				Stop:  &ast.MatchExpr{Expr: &defaultExpr, RegExpr: yyDollar[3].expr},
			}
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:143
		{
			yyVAL.stmts = yyDollar[2].stmts
		}
	case 12:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/grammar.go.y:149
		{
			yyVAL.stmts = []ast.Stmt{}
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:153
		{
			yyVAL.stmts = yyDollar[1].stmts
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:159
		{
			yyVAL.stmts = []ast.Stmt{yyDollar[2].stmt}
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:163
		{
			yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[4].stmt)
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:169
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyDollar[1].expr}
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:173
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyDollar[1].expr}
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:177
		{
			yyVAL.stmt = &ast.DelStmt{Expr: yyDollar[2].expr}
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:181
		{
			yyVAL.stmt = &ast.PrintStmt{Exprs: defaultExprs}
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:185
		{
			yyVAL.stmt = &ast.PrintStmt{Exprs: yyDollar[2].exprs}
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:189
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 22:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:193
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyDollar[3].stmts}
		}
	case 23:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/grammar.go.y:197
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyDollar[4].stmts, Expr: yyDollar[2].expr}
		}
	case 24:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line ./parser/grammar.go.y:201
		{
			yyVAL.stmt = &ast.CForLoopStmt{Stmt1: yyDollar[2].stmt, Expr2: yyDollar[4].expr, Expr3: yyDollar[6].expr, Stmts: yyDollar[8].stmts}
		}
	case 25:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:205
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyDollar[3].stmts}
		}
	case 26:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/grammar.go.y:209
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyDollar[4].stmts, Expr: yyDollar[2].expr}
		}
	case 27:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./parser/grammar.go.y:213
		{
			yyVAL.stmt = &ast.DoLoopStmt{Stmts: yyDollar[3].stmts, Expr: yyDollar[7].expr}
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:217
		{
			yyVAL.stmt = &ast.BreakStmt{}
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:221
		{
			yyVAL.stmt = &ast.ContinueStmt{}
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:225
		{
			yyVAL.stmt = &ast.NextStmt{}
		}
	case 31:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line ./parser/grammar.go.y:229
		{
			yyVAL.stmt = &ast.MapLoopStmt{KeyId: yyDollar[3].token.Literal, MapId: yyDollar[5].token.Literal, Stmts: yyDollar[8].stmts}
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:233
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyDollar[2].exprs}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:237
		{
			yyVAL.stmt = &ast.ExitStmt{Expr: yyDollar[2].expr}
		}
	case 34:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/grammar.go.y:243
		{
			yyVAL.stmt = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].stmts, Else: nil}
		}
	case 35:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./parser/grammar.go.y:247
		{
			yyVAL.stmt.(*ast.IfStmt).ElseIf = append(yyVAL.stmt.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyDollar[4].expr, Then: yyDollar[6].stmts})
		}
	case 36:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/grammar.go.y:251
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
		//line ./parser/grammar.go.y:262
		{
			yyVAL.expr = &ast.AssExpr{Left: yyDollar[1].exprs, Right: yyDollar[3].exprs}
		}
	case 38:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/grammar.go.y:268
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:272
		{
			yyVAL.exprs = yyDollar[1].exprs
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:278
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:282
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:288
		{
			yyVAL.expr = &ast.AssExpr{Left: []ast.Expr{yyDollar[1].expr}, Right: []ast.Expr{yyDollar[3].expr}}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:293
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "+=", Right: yyDollar[3].expr}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:297
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "-=", Right: yyDollar[3].expr}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:301
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "*=", Right: yyDollar[3].expr}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:305
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "/=", Right: yyDollar[3].expr}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:309
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "%=", Right: yyDollar[3].expr}
		}
	case 48:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser/grammar.go.y:314
		{
			yyVAL.expr = &ast.TriOpExpr{Cond: yyDollar[1].expr, Then: yyDollar[3].expr, Else: yyDollar[5].expr}
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:319
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "||", Right: yyDollar[3].expr}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:323
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "&&", Right: yyDollar[3].expr}
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:327
		{
			yyVAL.expr = &ast.CallExpr{Name: "printf", SubExprs: yyDollar[2].exprs}
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:331
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:337
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:342
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "CAT", Right: yyDollar[2].expr}
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:348
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:353
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "+", Right: yyDollar[3].expr}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:357
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "-", Right: yyDollar[3].expr}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:361
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "*", Right: yyDollar[3].expr}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:365
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "/", Right: yyDollar[3].expr}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:369
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "%", Right: yyDollar[3].expr}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:374
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "==", Right: yyDollar[3].expr}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:378
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "!=", Right: yyDollar[3].expr}
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:382
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: ">", Right: yyDollar[3].expr}
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:386
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: ">=", Right: yyDollar[3].expr}
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:390
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "<", Right: yyDollar[3].expr}
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:394
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "<=", Right: yyDollar[3].expr}
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:398
		{
			yyVAL.expr = &ast.ContainKeyExpr{KeyExpr: yyDollar[1].expr, MapId: yyDollar[3].token.Literal}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:403
		{
			yyVAL.expr = &ast.MatchExpr{Expr: yyDollar[1].expr, RegExpr: yyDollar[3].expr}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:407
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: &ast.MatchExpr{Expr: yyDollar[1].expr, RegExpr: yyDollar[3].expr}}
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:411
		{
			yyVAL.expr = &ast.MatchExpr{Expr: &defaultExpr, RegExpr: yyDollar[1].expr}
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:416
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "++", After: true}
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:420
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "--", After: true}
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:425
		{
			yyVAL.expr = &ast.GetlineExpr{Var: yyDollar[2].expr, Redir: yyDollar[3].expr}
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:431
		{
			//fmt.Println("YACC: want regexp!!")
			IN_REGEXP = true
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:436
		{
			//fmt.Println("FINISH")
			yyVAL.expr = &ast.RegExpr{Literal: yyDollar[3].token.Literal}
		}
	case 76:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/grammar.go.y:443
		{
			yyVAL.expr = nil
		}
	case 77:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:447
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:453
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:457
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 80:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:462
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].token.Literal, SubExprs: yyDollar[3].exprs}
		}
	case 81:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:466
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].token.Literal, SubExprs: yyDollar[3].exprs}
		}
	case 82:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:470
		{
			yyVAL.expr = &ast.AnonymousCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
		}
	case 83:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./parser/grammar.go.y:475
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyDollar[3].ident_args, Stmts: yyDollar[6].stmts}
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:480
		{
			yyVAL.expr = &ast.ParentExpr{SubExpr: yyDollar[2].expr}
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:485
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[2].expr, Operator: "++"}
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:489
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[2].expr, Operator: "--"}
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:494
		{
			yyVAL.expr = &ast.NumExpr{Literal: yyDollar[1].token.Literal}
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:498
		{
			yyVAL.expr = &ast.ConstExpr{Literal: yyDollar[1].token.Literal}
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:502
		{
			yyVAL.expr = &ast.ConstExpr{Literal: yyDollar[1].token.Literal}
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:506
		{
			yyVAL.expr = &ast.ConstExpr{Literal: yyDollar[1].token.Literal}
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:510
		{
			yyVAL.expr = &ast.StringExpr{Literal: yyDollar[1].token.Literal}
		}
	case 92:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:515
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 93:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:520
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "+", Expr: yyDollar[2].expr}
		}
	case 94:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:524
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:530
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 96:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:534
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 97:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:540
		{
			yyVAL.expr = &ast.ItemExpr{Expr: yyDollar[1].expr, Index: yyDollar[3].exprs}
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:544
		{
			yyVAL.expr = &ast.IdentExpr{Literal: yyDollar[1].token.Literal}
		}
	case 99:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/grammar.go.y:550
		{
			yyVAL.expr = nil
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:554
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:561
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 102:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:565
		{
			yyVAL.expr = &ast.FieldExpr{Expr: yyDollar[2].expr}
		}
	case 103:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/grammar.go.y:571
		{
			yyVAL.ident_args = []string{}
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:575
		{
			yyVAL.ident_args = []string{yyDollar[1].token.Literal}
		}
	case 105:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:579
		{
			yyVAL.ident_args = append(yyDollar[1].ident_args, yyDollar[4].token.Literal)
		}
	case 106:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/grammar.go.y:585
		{
			yyVAL.stmt = nil
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:589
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 108:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/grammar.go.y:595
		{
			yyVAL.expr = nil
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:599
		{
			yyVAL.expr = yyDollar[1].expr
		}
	}
	goto yystack /* stack new state and value */
}
