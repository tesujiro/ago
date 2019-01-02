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
	"'|'",
	"']'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line ./parser/grammar.go.y:630

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 13,
	61, 80,
	-2, 72,
	-1, 14,
	65, 12,
	-2, 111,
	-1, 67,
	61, 93,
	62, 93,
	-2, 99,
	-1, 107,
	65, 113,
	-2, 111,
	-1, 122,
	48, 96,
	-2, 93,
	-1, 165,
	65, 12,
	-2, 111,
	-1, 170,
	65, 12,
	-2, 111,
	-1, 172,
	65, 12,
	-2, 111,
	-1, 187,
	65, 12,
	-2, 111,
	-1, 192,
	65, 12,
	-2, 111,
	-1, 194,
	65, 12,
	-2, 111,
	-1, 198,
	65, 12,
	-2, 111,
	-1, 202,
	65, 12,
	-2, 111,
	-1, 215,
	61, 93,
	62, 93,
	-2, 97,
	-1, 218,
	65, 12,
	-2, 111,
	-1, 231,
	65, 12,
	-2, 111,
	-1, 232,
	65, 12,
	-2, 111,
}

const yyPrivate = 57344

const yyLast = 707

var yyAct = [...]int{

	59, 174, 12, 24, 108, 47, 66, 99, 45, 44,
	58, 15, 60, 14, 13, 5, 133, 5, 131, 221,
	195, 238, 68, 49, 2, 71, 237, 230, 92, 65,
	67, 69, 40, 65, 69, 56, 184, 225, 65, 43,
	65, 65, 223, 220, 101, 102, 103, 65, 65, 219,
	109, 125, 126, 127, 128, 129, 130, 218, 132, 104,
	217, 122, 97, 214, 45, 44, 212, 208, 45, 44,
	45, 44, 106, 45, 44, 105, 191, 232, 231, 45,
	44, 187, 172, 228, 65, 65, 65, 65, 65, 65,
	65, 65, 65, 65, 65, 43, 159, 151, 152, 43,
	154, 43, 185, 183, 43, 159, 157, 153, 192, 155,
	43, 186, 162, 202, 211, 166, 171, 198, 236, 169,
	158, 156, 175, 163, 179, 178, 122, 41, 194, 89,
	88, 161, 229, 135, 42, 98, 90, 23, 68, 173,
	133, 181, 137, 45, 44, 176, 67, 69, 177, 65,
	90, 45, 44, 46, 4, 134, 224, 180, 37, 164,
	138, 189, 109, 86, 87, 45, 190, 222, 205, 92,
	100, 193, 3, 122, 43, 160, 197, 150, 199, 70,
	39, 203, 43, 188, 19, 42, 123, 200, 136, 22,
	86, 87, 206, 204, 37, 20, 175, 210, 207, 110,
	209, 201, 8, 107, 213, 68, 17, 63, 216, 113,
	38, 48, 167, 215, 69, 7, 91, 6, 93, 94,
	1, 0, 175, 227, 226, 95, 96, 74, 75, 76,
	233, 0, 51, 52, 53, 54, 55, 234, 235, 0,
	77, 78, 80, 82, 85, 0, 0, 0, 86, 87,
	0, 0, 0, 0, 0, 0, 83, 50, 0, 0,
	0, 0, 139, 140, 141, 142, 143, 144, 145, 146,
	147, 148, 149, 0, 25, 30, 34, 31, 32, 33,
	84, 79, 81, 72, 73, 74, 75, 76, 28, 29,
	0, 0, 0, 0, 0, 111, 0, 0, 0, 112,
	16, 0, 124, 0, 114, 115, 116, 117, 118, 61,
	120, 121, 119, 0, 18, 0, 86, 87, 0, 0,
	0, 0, 0, 35, 36, 0, 23, 182, 26, 0,
	21, 0, 168, 0, 165, 25, 30, 34, 31, 32,
	33, 0, 0, 0, 0, 0, 0, 0, 0, 28,
	29, 72, 73, 74, 75, 76, 111, 0, 0, 0,
	112, 16, 0, 124, 0, 114, 115, 116, 117, 118,
	61, 120, 121, 119, 0, 18, 0, 0, 0, 0,
	0, 0, 0, 0, 35, 36, 0, 23, 0, 26,
	0, 21, 0, 27, 25, 30, 34, 31, 32, 33,
	0, 0, 0, 0, 0, 0, 0, 0, 28, 29,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	64, 0, 0, 0, 0, 80, 82, 0, 0, 61,
	0, 86, 87, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 35, 36, 0, 23, 0, 26, 0,
	21, 0, 27, 0, 0, 0, 0, 62, 25, 30,
	34, 31, 32, 33, 79, 81, 72, 73, 74, 75,
	76, 0, 28, 29, 0, 0, 0, 0, 0, 0,
	0, 10, 11, 0, 16, 25, 30, 34, 31, 32,
	33, 0, 0, 9, 0, 0, 0, 0, 18, 28,
	29, 0, 0, 0, 0, 0, 0, 35, 36, 0,
	23, 16, 26, 0, 21, 0, 27, 0, 14, 0,
	61, 0, 0, 0, 0, 18, 0, 0, 0, 0,
	0, 0, 0, 0, 35, 36, 0, 23, 0, 26,
	0, 21, 0, 27, 0, 170, 25, 30, 34, 31,
	32, 33, 0, 0, 0, 0, 0, 0, 0, 0,
	28, 29, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 16, 196, 30, 34, 31, 32, 33, 0,
	0, 61, 0, 0, 0, 0, 18, 28, 29, 0,
	0, 0, 0, 0, 0, 35, 36, 0, 23, 16,
	26, 0, 21, 0, 27, 0, 0, 0, 61, 0,
	0, 0, 0, 18, 0, 0, 0, 0, 0, 0,
	0, 0, 35, 36, 0, 23, 0, 26, 0, 21,
	0, 27, 25, 30, 34, 31, 32, 33, 0, 0,
	0, 0, 0, 0, 0, 0, 28, 29, 25, 30,
	34, 31, 32, 33, 0, 0, 0, 0, 16, 0,
	0, 0, 28, 29, 0, 0, 0, 61, 0, 0,
	0, 0, 18, 0, 64, 0, 0, 0, 0, 0,
	0, 35, 36, 61, 23, 0, 26, 0, 21, 0,
	57, 0, 0, 0, 0, 0, 0, 35, 36, 0,
	23, 0, 26, 0, 21, 0, 27,
}
var yyPact = [...]int{

	-49, 454, -1000, -49, -1000, -1000, -1000, -51, -49, 123,
	-1000, -1000, 136, 105, -49, 212, 628, 390, 644, -1000,
	-1000, 644, 230, -1000, 68, 74, 644, 542, 644, 644,
	-1000, -1000, -1000, -1000, -1000, 644, 644, -1000, -49, -49,
	-1000, 73, 166, 542, 542, 542, 81, 10, -49, 331,
	542, 542, 542, 542, 542, 542, -1000, 542, 92, 136,
	-1000, 72, 111, 230, 71, -1000, 90, -1000, 68, -1000,
	129, 68, 644, 644, 644, 644, 644, 644, 644, 644,
	644, 644, 644, 173, 81, 81, -1000, -1000, 542, 542,
	542, 145, 58, -1000, -1000, 145, 145, -1000, 166, 57,
	-1000, 128, 150, -1000, -1000, -1000, -1000, -49, -1000, 136,
	-1000, 542, 542, 126, 270, 481, 18, -1000, -1000, -1000,
	542, 542, 212, 100, 542, 136, 136, 136, 136, 136,
	136, 61, 58, -49, 644, 542, -1000, 644, -1000, 172,
	172, 145, 145, 145, 413, 413, 298, 298, 298, 298,
	-1000, -1000, -1000, 40, -32, 39, -1000, 48, 17, -49,
	542, 331, 136, 92, 44, -49, 64, -46, 569, -1000,
	-49, 53, -49, -1000, -1000, 136, 542, -49, 49, -1000,
	542, -1000, 230, -1000, -1000, -1000, -1000, -49, 164, 136,
	-1000, 542, -49, 2, -49, 542, 88, 1, -49, -2,
	92, 644, -49, -1000, -5, -1000, -7, -16, -1000, -22,
	-47, 163, -1000, -23, 121, -1000, -28, -1000, -49, -1000,
	-1000, 542, 20, -1000, 70, -1000, -38, 14, 13, 542,
	-1000, -49, -49, 55, -39, -44, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 220, 217, 215, 4, 212, 202, 211, 5, 209,
	0, 206, 199, 189, 3, 6, 11, 195, 1, 12,
	188, 10, 186, 18, 7, 23, 172, 154, 184, 179,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 2, 3, 3, 3, 3,
	3, 6, 8, 8, 7, 7, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 9, 9, 9, 12, 23, 23,
	21, 21, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 11, 11, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 29, 19, 20, 20, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 14, 14, 14, 14, 22, 22, 15, 15,
	16, 16, 17, 17, 24, 24, 24, 5, 5, 18,
	18, 25, 25, 26, 26, 27, 28,
}
var yyR2 = [...]int{

	0, 1, 2, 3, 2, 2, 5, 1, 1, 1,
	3, 3, 0, 2, 2, 4, 1, 1, 2, 1,
	2, 1, 4, 5, 9, 4, 5, 8, 1, 1,
	1, 9, 2, 2, 5, 7, 5, 3, 0, 1,
	1, 4, 3, 3, 3, 3, 3, 3, 5, 3,
	3, 2, 1, 4, 3, 1, 2, 1, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 1, 2, 2, 0, 3, 0, 2, 2,
	1, 4, 4, 4, 7, 3, 2, 2, 1, 1,
	1, 1, 1, 1, 2, 2, 1, 4, 0, 1,
	1, 2, 4, 1, 0, 1, 4, 0, 1, 0,
	1, 0, 1, 1, 2, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -25, -26, -27, 66, -2, -3, -6, 39,
	27, 28, -10, -19, 64, -16, 30, -11, 44, -28,
	-17, 60, -13, 56, -14, 4, 58, 62, 18, 19,
	5, 7, 8, 9, 6, 53, 54, -27, -6, -26,
	-25, 4, 62, 46, 16, 15, 48, -8, -7, -25,
	45, 20, 21, 22, 23, 24, -23, 62, -21, -10,
	-19, 39, 67, -13, 30, -16, -15, -16, -14, -19,
	-29, -14, 53, 54, 55, 56, 57, 10, 11, 51,
	12, 52, 13, 26, 50, 14, 18, 19, 62, 61,
	62, -13, -10, -13, -13, -13, -13, -25, 62, -24,
	4, -10, -10, -10, -19, 65, -25, -27, -4, -10,
	-12, 25, 29, -9, 34, 35, 36, 37, 38, 42,
	40, 41, -16, -22, 32, -10, -10, -10, -10, -10,
	-10, -23, -10, 48, 44, 62, -20, 52, 31, -13,
	-13, -13, -13, -13, -13, -13, -13, -13, -13, -13,
	4, -19, -19, -23, -21, -23, 63, -24, 63, 48,
	47, -25, -10, -21, 33, 64, -10, -5, 62, -4,
	64, -10, 64, -23, -18, -10, 45, 48, -10, 63,
	-25, -15, -13, 63, 68, 63, 63, 64, -25, -10,
	-4, 32, 64, -8, 64, 66, 4, -8, 64, -8,
	-21, -25, 64, -10, -8, 4, -10, -8, 65, -8,
	-18, 26, 65, -8, 65, -16, -8, 65, 64, 65,
	65, 66, 4, 65, 35, 65, -8, -18, 63, 62,
	65, 64, 64, -10, -8, -8, 63, 65, 65,
}
var yyDef = [...]int{

	111, -2, 1, 112, 113, 115, 2, 0, 111, 0,
	7, 8, 9, -2, -2, 93, 38, 52, 98, 75,
	100, 0, 55, 116, 57, 103, 0, 0, 0, 0,
	88, 89, 90, 91, 92, 0, 0, 114, 111, 4,
	5, 0, 104, 0, 0, 0, 0, 0, 111, 0,
	0, 0, 0, 0, 0, 0, 51, 38, 39, 40,
	72, 0, 0, 56, 0, 93, 77, -2, 0, 80,
	0, 101, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 73, 74, 38, 0,
	38, 79, 0, 86, 87, 94, 95, 3, 104, 0,
	105, 0, 49, 50, 10, 11, 13, -2, 14, 16,
	17, 0, 19, 21, 107, 0, 0, 28, 29, 30,
	38, 109, -2, 0, 0, 42, 43, 44, 45, 46,
	47, 0, 40, 111, 98, 38, 54, 0, 76, 58,
	59, 60, 61, 62, 63, 64, 65, 66, 67, 68,
	69, 70, 71, 0, 0, 0, 85, 0, 0, 111,
	0, 0, 18, 20, 0, -2, 16, 0, 0, 108,
	-2, 0, -2, 32, 33, 110, 0, 111, 0, 82,
	0, 53, 78, 83, 102, 81, 6, -2, 0, 48,
	15, 0, -2, 0, -2, 109, 103, 0, -2, 0,
	37, 0, -2, 41, 0, 106, 0, 0, 22, 0,
	0, 0, 25, 0, 0, -2, 0, 84, -2, 36,
	23, 109, 0, 26, 0, 34, 0, 0, 0, 0,
	35, -2, -2, 0, 0, 0, 27, 24, 31,
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
	3, 61, 3, 68, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 64, 67, 65, 50,
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
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:336
		{
			yyVAL.expr = &ast.GetlineExpr{Command: yyDollar[1].expr, Var: yyDollar[4].expr}
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:340
		{
			yyVAL.expr = &ast.GetlineExpr{Var: yyDollar[2].expr, Redir: yyDollar[3].expr}
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:346
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:351
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "CAT", Right: yyDollar[2].expr}
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:357
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:362
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "+", Right: yyDollar[3].expr}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:366
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "-", Right: yyDollar[3].expr}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:370
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "*", Right: yyDollar[3].expr}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:374
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "/", Right: yyDollar[3].expr}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:378
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "%", Right: yyDollar[3].expr}
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:383
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "==", Right: yyDollar[3].expr}
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:387
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "!=", Right: yyDollar[3].expr}
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:391
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: ">", Right: yyDollar[3].expr}
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:395
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: ">=", Right: yyDollar[3].expr}
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:399
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "<", Right: yyDollar[3].expr}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:403
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "<=", Right: yyDollar[3].expr}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:407
		{
			yyVAL.expr = &ast.ContainKeyExpr{KeyExpr: yyDollar[1].expr, MapId: yyDollar[3].token.Literal}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:412
		{
			yyVAL.expr = &ast.MatchExpr{Expr: yyDollar[1].expr, RegExpr: yyDollar[3].expr}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:416
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: &ast.MatchExpr{Expr: yyDollar[1].expr, RegExpr: yyDollar[3].expr}}
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:420
		{
			yyVAL.expr = &ast.MatchExpr{Expr: &defaultExpr, RegExpr: yyDollar[1].expr}
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:425
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "++", After: true}
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:429
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "--", After: true}
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:435
		{
			//fmt.Println("YACC: want regexp!!")
			IN_REGEXP = true
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:440
		{
			//fmt.Println("FINISH")
			yyVAL.expr = &ast.RegExpr{Literal: yyDollar[3].token.Literal}
		}
	case 77:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/grammar.go.y:447
		{
			yyVAL.expr = nil
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:451
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:457
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:461
		{
			yyVAL.expr = yyDollar[1].expr
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
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].token.Literal, SubExprs: yyDollar[3].exprs}
		}
	case 83:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:474
		{
			yyVAL.expr = &ast.AnonymousCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
		}
	case 84:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./parser/grammar.go.y:479
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyDollar[3].ident_args, Stmts: yyDollar[6].stmts}
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser/grammar.go.y:484
		{
			yyVAL.expr = &ast.ParentExpr{SubExpr: yyDollar[2].expr}
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:489
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[2].expr, Operator: "++"}
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:493
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[2].expr, Operator: "--"}
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:498
		{
			yyVAL.expr = &ast.NumExpr{Literal: yyDollar[1].token.Literal}
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
			yyVAL.expr = &ast.ConstExpr{Literal: yyDollar[1].token.Literal}
		}
	case 92:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:514
		{
			yyVAL.expr = &ast.StringExpr{Literal: yyDollar[1].token.Literal}
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:519
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 94:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:524
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "+", Expr: yyDollar[2].expr}
		}
	case 95:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:528
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:534
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 97:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:538
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 98:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/grammar.go.y:544
		{
			yyVAL.expr = nil
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:548
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:555
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 101:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser/grammar.go.y:559
		{
			yyVAL.expr = &ast.FieldExpr{Expr: yyDollar[2].expr}
		}
	case 102:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:565
		{
			yyVAL.expr = &ast.ItemExpr{Expr: yyDollar[1].expr, Index: yyDollar[3].exprs}
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:569
		{
			yyVAL.expr = &ast.IdentExpr{Literal: yyDollar[1].token.Literal}
		}
	case 104:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/grammar.go.y:575
		{
			yyVAL.ident_args = []string{}
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:579
		{
			yyVAL.ident_args = []string{yyDollar[1].token.Literal}
		}
	case 106:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser/grammar.go.y:583
		{
			yyVAL.ident_args = append(yyDollar[1].ident_args, yyDollar[4].token.Literal)
		}
	case 107:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/grammar.go.y:589
		{
			yyVAL.stmt = nil
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:593
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 109:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser/grammar.go.y:599
		{
			yyVAL.expr = nil
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser/grammar.go.y:603
		{
			yyVAL.expr = yyDollar[1].expr
		}
	}
	goto yystack /* stack new state and value */
}
