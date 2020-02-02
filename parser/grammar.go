package parser

import __yyfmt__ "fmt"
import (
	//"fmt"
	"github.com/tesujiro/ago/ast"
)

var defaultExpr = ast.FieldExpr{Expr: &ast.NumExpr{Literal: "0"}}
var defaultExprs = []ast.Expr{&defaultExpr}
var inRegExp bool

type yySymType struct {
	yys       int
	token     ast.Token
	rule      ast.Rule
	rules     []ast.Rule
	pattern   ast.Pattern
	stmt      ast.Stmt
	stmts     []ast.Stmt
	expr      ast.Expr
	exprs     []ast.Expr
	identArgs []string
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

var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 13,
	65, 12,
	-2, 113,
	-1, 66,
	61, 95,
	62, 95,
	-2, 101,
	-1, 106,
	65, 115,
	-2, 113,
	-1, 121,
	48, 98,
	-2, 95,
	-1, 166,
	65, 12,
	-2, 113,
	-1, 171,
	65, 12,
	-2, 113,
	-1, 173,
	65, 12,
	-2, 113,
	-1, 188,
	65, 12,
	-2, 113,
	-1, 193,
	65, 12,
	-2, 113,
	-1, 195,
	65, 12,
	-2, 113,
	-1, 199,
	65, 12,
	-2, 113,
	-1, 203,
	65, 12,
	-2, 113,
	-1, 216,
	61, 95,
	62, 95,
	-2, 99,
	-1, 219,
	65, 12,
	-2, 113,
	-1, 232,
	65, 12,
	-2, 113,
	-1, 233,
	65, 12,
	-2, 113,
}

const yyPrivate = 57344

const yyLast = 730

var yyAct = [...]int{

	59, 175, 12, 107, 16, 47, 98, 58, 132, 130,
	46, 45, 46, 45, 65, 13, 5, 5, 21, 23,
	222, 196, 239, 49, 2, 56, 90, 238, 185, 46,
	45, 231, 40, 46, 45, 226, 67, 68, 69, 68,
	224, 44, 221, 44, 100, 101, 102, 103, 220, 218,
	108, 124, 125, 126, 127, 128, 129, 215, 131, 219,
	44, 203, 96, 213, 44, 209, 14, 192, 104, 233,
	232, 188, 105, 173, 46, 45, 229, 160, 199, 46,
	45, 230, 195, 64, 66, 186, 64, 150, 152, 212,
	184, 64, 187, 64, 64, 154, 153, 160, 155, 193,
	64, 64, 149, 151, 158, 44, 180, 87, 86, 41,
	44, 163, 159, 134, 167, 172, 121, 170, 42, 164,
	97, 176, 237, 88, 179, 88, 136, 156, 177, 174,
	162, 178, 51, 52, 53, 54, 55, 64, 64, 64,
	64, 64, 64, 64, 64, 64, 64, 64, 182, 64,
	64, 132, 67, 68, 133, 225, 181, 50, 165, 157,
	46, 45, 190, 108, 46, 45, 191, 42, 84, 85,
	90, 46, 194, 46, 45, 84, 85, 198, 3, 200,
	121, 223, 204, 206, 189, 201, 39, 99, 8, 84,
	85, 44, 161, 207, 205, 44, 38, 176, 211, 208,
	66, 210, 202, 64, 44, 214, 43, 148, 95, 217,
	70, 71, 72, 73, 74, 20, 35, 64, 122, 64,
	135, 67, 68, 176, 228, 227, 72, 73, 74, 121,
	18, 234, 62, 4, 109, 112, 48, 37, 235, 236,
	89, 168, 91, 92, 7, 6, 1, 0, 0, 93,
	94, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 75, 76, 78, 80, 83, 0, 216,
	0, 84, 85, 37, 0, 0, 0, 0, 0, 81,
	0, 0, 106, 0, 0, 0, 137, 138, 139, 140,
	141, 142, 143, 144, 145, 146, 147, 22, 28, 32,
	29, 30, 31, 82, 77, 79, 70, 71, 72, 73,
	74, 26, 27, 0, 0, 0, 0, 0, 110, 0,
	0, 0, 111, 15, 0, 123, 0, 113, 114, 115,
	116, 117, 60, 119, 120, 118, 0, 17, 0, 0,
	0, 0, 0, 0, 0, 0, 33, 34, 0, 36,
	0, 24, 183, 19, 0, 169, 0, 166, 22, 28,
	32, 29, 30, 31, 0, 0, 62, 0, 62, 0,
	0, 0, 26, 27, 0, 0, 0, 0, 0, 110,
	0, 0, 0, 111, 15, 0, 123, 0, 113, 114,
	115, 116, 117, 60, 119, 120, 118, 0, 17, 0,
	0, 0, 0, 0, 0, 0, 0, 33, 34, 0,
	36, 0, 24, 0, 19, 0, 25, 22, 28, 32,
	29, 30, 31, 0, 0, 0, 0, 0, 0, 0,
	0, 26, 27, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 63, 0, 0, 0, 0, 78, 80,
	0, 0, 60, 0, 84, 85, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 33, 34, 0, 36,
	0, 24, 0, 19, 0, 25, 0, 0, 0, 0,
	61, 22, 28, 32, 29, 30, 31, 77, 79, 70,
	71, 72, 73, 74, 0, 26, 27, 0, 0, 0,
	0, 0, 0, 0, 10, 11, 0, 15, 22, 28,
	32, 29, 30, 31, 0, 0, 9, 0, 0, 0,
	0, 17, 26, 27, 0, 0, 0, 0, 0, 0,
	33, 34, 0, 36, 15, 24, 0, 19, 0, 25,
	0, 13, 0, 60, 0, 0, 0, 0, 17, 0,
	0, 0, 0, 0, 0, 0, 0, 33, 34, 0,
	36, 0, 24, 0, 19, 0, 25, 0, 171, 22,
	28, 32, 29, 30, 31, 0, 0, 0, 0, 0,
	0, 0, 0, 26, 27, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 15, 197, 28, 32, 29,
	30, 31, 0, 0, 60, 0, 0, 0, 0, 17,
	26, 27, 0, 0, 0, 0, 0, 0, 33, 34,
	0, 36, 15, 24, 0, 19, 0, 25, 0, 0,
	0, 60, 0, 0, 0, 0, 17, 0, 0, 0,
	0, 0, 0, 0, 0, 33, 34, 0, 36, 0,
	24, 0, 19, 0, 25, 22, 28, 32, 29, 30,
	31, 0, 0, 0, 0, 0, 0, 0, 0, 26,
	27, 22, 28, 32, 29, 30, 31, 0, 0, 0,
	0, 15, 0, 0, 0, 26, 27, 0, 0, 0,
	60, 0, 0, 0, 0, 17, 0, 63, 0, 0,
	0, 0, 0, 0, 33, 34, 60, 36, 0, 24,
	0, 19, 0, 57, 0, 0, 0, 0, 0, 0,
	33, 34, 0, 36, 0, 24, 0, 19, 0, 25,
}
var yyPact = [...]int{

	-50, 477, -1000, -50, -1000, -1000, -1000, -49, -50, 105,
	-1000, -1000, 158, -50, 112, 651, 413, 667, -1000, 667,
	253, 46, 61, -1000, 667, 565, 667, 667, -1000, -1000,
	-1000, -1000, -1000, 667, 667, -1000, -1000, -1000, -50, -50,
	-1000, 58, 183, 565, 565, 565, 565, 3, -50, 354,
	565, 565, 565, 565, 565, 565, -1000, 565, 103, 149,
	56, 110, 157, 51, -1000, 74, -1000, 46, -1000, 46,
	667, 667, 667, 667, 667, 667, 667, 667, 667, 667,
	667, 203, 667, 667, -1000, -1000, 565, 565, 565, 150,
	64, -1000, -1000, 150, 150, 128, -1000, 183, 49, -1000,
	149, 145, 156, -1000, -1000, -1000, -50, -1000, 149, -1000,
	565, 565, 125, 293, 504, 9, -1000, -1000, -1000, 565,
	565, 112, 83, 565, 149, 149, 149, 149, 149, 149,
	43, 64, -50, 667, 565, -1000, 667, 171, 171, 150,
	150, 150, 436, 436, 157, 157, 157, 157, -1000, -1000,
	667, -1000, 667, 27, -40, 22, -1000, -1000, 29, 7,
	-50, 565, 354, 149, 103, 35, -50, 18, -45, 592,
	-1000, -50, 14, -50, -1000, -1000, 149, 565, -50, -3,
	-1000, 565, -1000, 253, -1000, -1000, -1000, -1000, -50, 179,
	149, -1000, 565, -50, 0, -50, 565, 63, -2, -50,
	-8, 103, 667, -50, -1000, -16, -1000, -5, -17, -1000,
	-23, -46, 177, -1000, -25, 120, -1000, -30, -1000, -50,
	-1000, -1000, 565, 13, -1000, 19, -1000, -34, 6, 5,
	565, -1000, -50, -50, 59, -38, -43, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 246, 245, 244, 3, 241, 188, 236, 5, 235,
	0, 4, 234, 215, 18, 14, 66, 230, 1, 19,
	220, 7, 218, 9, 6, 23, 178, 233, 216, 208,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 2, 3, 3, 3, 3,
	3, 6, 8, 8, 7, 7, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 9, 9, 9, 12, 23, 23,
	21, 21, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 11, 11, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 29, 19, 20,
	20, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 22, 22,
	15, 15, 16, 16, 17, 17, 24, 24, 24, 5,
	5, 18, 18, 25, 25, 26, 26, 27, 28,
}
var yyR2 = [...]int{

	0, 1, 2, 3, 2, 2, 5, 1, 1, 1,
	3, 3, 0, 2, 2, 4, 1, 1, 2, 1,
	2, 1, 4, 5, 9, 4, 5, 8, 1, 1,
	1, 9, 2, 2, 5, 7, 5, 3, 0, 1,
	1, 4, 3, 3, 3, 3, 3, 3, 5, 3,
	3, 2, 1, 4, 3, 1, 2, 1, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 1, 2, 2, 0, 3, 0,
	2, 2, 1, 4, 4, 4, 7, 3, 2, 2,
	1, 1, 1, 1, 1, 1, 2, 2, 1, 4,
	0, 1, 1, 2, 4, 1, 0, 1, 4, 0,
	1, 0, 1, 0, 1, 1, 2, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -25, -26, -27, 66, -2, -3, -6, 39,
	27, 28, -10, 64, -16, 30, -11, 44, -17, 60,
	-13, -14, 4, -19, 58, 62, 18, 19, 5, 7,
	8, 9, 6, 53, 54, -28, 56, -27, -6, -26,
	-25, 4, 62, 48, 46, 16, 15, -8, -7, -25,
	45, 20, 21, 22, 23, 24, -23, 62, -21, -10,
	39, 67, -13, 30, -16, -15, -16, -14, -19, -14,
	53, 54, 55, 56, 57, 10, 11, 51, 12, 52,
	13, 26, 50, 14, 18, 19, 62, 61, 62, -13,
	-10, -13, -13, -13, -13, -29, -25, 62, -24, 4,
	-10, -10, -10, -10, 65, -25, -27, -4, -10, -12,
	25, 29, -9, 34, 35, 36, 37, 38, 42, 40,
	41, -16, -22, 32, -10, -10, -10, -10, -10, -10,
	-23, -10, 48, 44, 62, -20, 52, -13, -13, -13,
	-13, -13, -13, -13, -13, -13, -13, -13, 4, -19,
	-11, -19, -11, -23, -21, -23, 63, 31, -24, 63,
	48, 47, -25, -10, -21, 33, 64, -10, -5, 62,
	-4, 64, -10, 64, -23, -18, -10, 45, 48, -10,
	63, -25, -15, -13, 63, 68, 63, 63, 64, -25,
	-10, -4, 32, 64, -8, 64, 66, 4, -8, 64,
	-8, -21, -25, 64, -10, -8, 4, -10, -8, 65,
	-8, -18, 26, 65, -8, 65, -16, -8, 65, 64,
	65, 65, 66, 4, 65, 35, 65, -8, -18, 63,
	62, 65, 64, 64, -10, -8, -8, 63, 65, 65,
}
var yyDef = [...]int{

	113, -2, 1, 114, 115, 117, 2, 0, 113, 0,
	7, 8, 9, -2, 95, 38, 52, 100, 102, 0,
	55, 57, 105, 74, 0, 0, 0, 0, 90, 91,
	92, 93, 94, 0, 0, 77, 118, 116, 113, 4,
	5, 0, 106, 0, 0, 0, 0, 0, 113, 0,
	0, 0, 0, 0, 0, 0, 51, 38, 39, 40,
	0, 0, 56, 0, 95, 79, -2, 0, 82, 103,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 75, 76, 38, 0, 38, 81,
	0, 88, 89, 96, 97, 0, 3, 106, 0, 107,
	10, 0, 49, 50, 11, 13, -2, 14, 16, 17,
	0, 19, 21, 109, 0, 0, 28, 29, 30, 38,
	111, -2, 0, 0, 42, 43, 44, 45, 46, 47,
	0, 40, 113, 100, 38, 54, 0, 58, 59, 60,
	61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
	71, 72, 73, 0, 0, 0, 87, 78, 0, 0,
	113, 0, 0, 18, 20, 0, -2, 16, 0, 0,
	110, -2, 0, -2, 32, 33, 112, 0, 113, 0,
	84, 0, 53, 80, 85, 104, 83, 6, -2, 0,
	48, 15, 0, -2, 0, -2, 111, 105, 0, -2,
	0, 37, 0, -2, 41, 0, 108, 0, 0, 22,
	0, 0, 0, 25, 0, 0, -2, 0, 86, -2,
	36, 23, 111, 0, 26, 0, 34, 0, 0, 0,
	0, 35, -2, -2, 0, 0, 0, 27, 24, 31,
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
		{
			yyVAL.rules = []ast.Rule{}
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.rules = append(yyDollar[1].rules, yyDollar[2].rule)
			yylex.(*Lexer).result = yyVAL.rules
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.rule = ast.Rule{Pattern: yyDollar[1].pattern, Action: yyDollar[2].stmts}
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.rule = ast.Rule{Pattern: yyDollar[1].pattern, Action: []ast.Stmt{&ast.PrintStmt{Exprs: defaultExprs}}}
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.rule = ast.Rule{Pattern: &ast.ExprPattern{}, Action: yyDollar[1].stmts}
		}
	case 6:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			//fmt.Println("FUNC RULE")
			yyVAL.pattern = &ast.FuncPattern{Name: yyDollar[2].token.Literal, Args: yyDollar[4].identArgs}
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.pattern = &ast.BeginPattern{}
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.pattern = &ast.EndPattern{}
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.pattern = &ast.ExprPattern{Expr: yyDollar[1].expr}
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.pattern = &ast.StartStopPattern{
				Start: yyDollar[1].expr,
				Stop:  yyDollar[3].expr,
			}
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.stmts = yyDollar[2].stmts
		}
	case 12:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.stmts = []ast.Stmt{}
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmts = yyDollar[1].stmts
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmts = []ast.Stmt{yyDollar[2].stmt}
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[4].stmt)
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyDollar[1].expr}
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmt = &ast.ExprStmt{Expr: yyDollar[1].expr}
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmt = &ast.DelStmt{Expr: yyDollar[2].expr}
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmt = &ast.PrintStmt{Exprs: defaultExprs}
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmt = &ast.PrintStmt{Exprs: yyDollar[2].exprs}
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 22:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyDollar[3].stmts}
		}
	case 23:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyDollar[4].stmts, Expr: yyDollar[2].expr}
		}
	case 24:
		yyDollar = yyS[yypt-9 : yypt+1]
		{
			yyVAL.stmt = &ast.CForLoopStmt{Stmt1: yyDollar[2].stmt, Expr2: yyDollar[4].expr, Expr3: yyDollar[6].expr, Stmts: yyDollar[8].stmts}
		}
	case 25:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyDollar[3].stmts}
		}
	case 26:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyDollar[4].stmts, Expr: yyDollar[2].expr}
		}
	case 27:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.stmt = &ast.DoLoopStmt{Stmts: yyDollar[3].stmts, Expr: yyDollar[7].expr}
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmt = &ast.BreakStmt{}
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmt = &ast.ContinueStmt{}
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmt = &ast.NextStmt{}
		}
	case 31:
		yyDollar = yyS[yypt-9 : yypt+1]
		{
			yyVAL.stmt = &ast.MapLoopStmt{KeyID: yyDollar[3].token.Literal, MapID: yyDollar[5].token.Literal, Stmts: yyDollar[8].stmts}
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyDollar[2].exprs}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmt = &ast.ExitStmt{Expr: yyDollar[2].expr}
		}
	case 34:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.stmt = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].stmts, Else: nil}
		}
	case 35:
		yyDollar = yyS[yypt-7 : yypt+1]
		{
			yyVAL.stmt.(*ast.IfStmt).ElseIf = append(yyVAL.stmt.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyDollar[4].expr, Then: yyDollar[6].stmts})
		}
	case 36:
		yyDollar = yyS[yypt-5 : yypt+1]
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
		{
			yyVAL.expr = &ast.AssExpr{Left: yyDollar[1].exprs, Right: yyDollar[3].exprs}
		}
	case 38:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exprs = yyDollar[1].exprs
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.AssExpr{Left: []ast.Expr{yyDollar[1].expr}, Right: []ast.Expr{yyDollar[3].expr}}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "+=", Right: yyDollar[3].expr}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "-=", Right: yyDollar[3].expr}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "*=", Right: yyDollar[3].expr}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "/=", Right: yyDollar[3].expr}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "%=", Right: yyDollar[3].expr}
		}
	case 48:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.expr = &ast.TriOpExpr{Cond: yyDollar[1].expr, Then: yyDollar[3].expr, Else: yyDollar[5].expr}
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "||", Right: yyDollar[3].expr}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "&&", Right: yyDollar[3].expr}
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.CallExpr{Name: "printf", SubExprs: yyDollar[2].exprs}
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 53:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &ast.GetlineExpr{Command: yyDollar[1].expr, Var: yyDollar[4].expr}
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.GetlineExpr{Var: yyDollar[2].expr, Redir: yyDollar[3].expr}
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "CAT", Right: yyDollar[2].expr}
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "+", Right: yyDollar[3].expr}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "-", Right: yyDollar[3].expr}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "*", Right: yyDollar[3].expr}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "/", Right: yyDollar[3].expr}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "%", Right: yyDollar[3].expr}
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "==", Right: yyDollar[3].expr}
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "!=", Right: yyDollar[3].expr}
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: ">", Right: yyDollar[3].expr}
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: ">=", Right: yyDollar[3].expr}
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "<", Right: yyDollar[3].expr}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "<=", Right: yyDollar[3].expr}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.ContainKeyExpr{KeyExpr: yyDollar[1].expr, MapID: yyDollar[3].token.Literal}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.MatchExpr{Expr: yyDollar[1].expr, RegExpr: yyDollar[3].expr}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.MatchExpr{Expr: yyDollar[1].expr, RegExpr: yyDollar[3].expr}
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: &ast.MatchExpr{Expr: yyDollar[1].expr, RegExpr: yyDollar[3].expr}}
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: &ast.MatchExpr{Expr: yyDollar[1].expr, RegExpr: yyDollar[3].expr}}
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &ast.MatchExpr{Expr: &defaultExpr, RegExpr: yyDollar[1].expr}
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "++", After: true}
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "--", After: true}
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			//fmt.Println("YACC: want regexp!!")
			inRegExp = true
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			//fmt.Println("FINISH")
			yyVAL.expr = &ast.StringExpr{Literal: yyDollar[3].token.Literal}
		}
	case 79:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.expr = nil
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 83:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].token.Literal, SubExprs: yyDollar[3].exprs}
		}
	case 84:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].token.Literal, SubExprs: yyDollar[3].exprs}
		}
	case 85:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &ast.AnonymousCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
		}
	case 86:
		yyDollar = yyS[yypt-7 : yypt+1]
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyDollar[3].identArgs, Stmts: yyDollar[6].stmts}
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.ParentExpr{SubExpr: yyDollar[2].expr}
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[2].expr, Operator: "++"}
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[2].expr, Operator: "--"}
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &ast.NumExpr{Literal: yyDollar[1].token.Literal}
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &ast.ConstExpr{Literal: yyDollar[1].token.Literal}
		}
	case 92:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &ast.ConstExpr{Literal: yyDollar[1].token.Literal}
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &ast.ConstExpr{Literal: yyDollar[1].token.Literal}
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &ast.StringExpr{Literal: yyDollar[1].token.Literal}
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "+", Expr: yyDollar[2].expr}
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 99:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 100:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.expr = nil
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.FieldExpr{Expr: yyDollar[2].expr}
		}
	case 104:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &ast.ItemExpr{Expr: yyDollar[1].expr, Index: yyDollar[3].exprs}
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &ast.IdentExpr{Literal: yyDollar[1].token.Literal}
		}
	case 106:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.identArgs = []string{}
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.identArgs = []string{yyDollar[1].token.Literal}
		}
	case 108:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.identArgs = append(yyDollar[1].identArgs, yyDollar[4].token.Literal)
		}
	case 109:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.stmt = nil
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 111:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.expr = nil
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = yyDollar[1].expr
		}
	}
	goto yystack /* stack new state and value */
}
