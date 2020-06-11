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
const POW = 57352
const EQEQ = 57353
const NEQ = 57354
const GE = 57355
const LE = 57356
const NOTTILDE = 57357
const ANDAND = 57358
const OROR = 57359
const LEN = 57360
const PLUSPLUS = 57361
const MINUSMINUS = 57362
const PLUSEQ = 57363
const MINUSEQ = 57364
const MULEQ = 57365
const DIVEQ = 57366
const MODEQ = 57367
const POWEQ = 57368
const DELETE = 57369
const IN = 57370
const BEGIN = 57371
const END = 57372
const PRINT = 57373
const PRINTF = 57374
const REGEXP = 57375
const IF = 57376
const ELSE = 57377
const FOR = 57378
const WHILE = 57379
const DO = 57380
const BREAK = 57381
const CONTINUE = 57382
const FUNC = 57383
const RETURN = 57384
const EXIT = 57385
const NEXT = 57386
const NEXTFILE = 57387
const CONCAT_OP = 57388
const GETLINE = 57389
const LOWEST = 57390
const UNARY = 57391

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
	"POW",
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
	"POWEQ",
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
	"NEXTFILE",
	"CONCAT_OP",
	"GETLINE",
	"LOWEST",
	"'='",
	"'?'",
	"':'",
	"','",
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
	68, 12,
	-2, 121,
	-1, 75,
	64, 103,
	65, 103,
	-2, 109,
	-1, 101,
	11, 0,
	12, 0,
	-2, 50,
	-1, 102,
	11, 0,
	12, 0,
	-2, 51,
	-1, 103,
	13, 0,
	14, 0,
	54, 0,
	55, 0,
	-2, 52,
	-1, 104,
	13, 0,
	14, 0,
	54, 0,
	55, 0,
	-2, 53,
	-1, 105,
	13, 0,
	14, 0,
	54, 0,
	55, 0,
	-2, 54,
	-1, 106,
	13, 0,
	14, 0,
	54, 0,
	55, 0,
	-2, 55,
	-1, 117,
	68, 123,
	-2, 121,
	-1, 133,
	52, 106,
	-2, 103,
	-1, 169,
	68, 12,
	-2, 121,
	-1, 174,
	68, 12,
	-2, 121,
	-1, 176,
	68, 12,
	-2, 121,
	-1, 191,
	68, 12,
	-2, 121,
	-1, 196,
	68, 12,
	-2, 121,
	-1, 200,
	68, 12,
	-2, 121,
	-1, 206,
	68, 12,
	-2, 121,
	-1, 210,
	68, 12,
	-2, 121,
	-1, 226,
	64, 103,
	65, 103,
	-2, 107,
	-1, 230,
	68, 12,
	-2, 121,
	-1, 251,
	68, 12,
	-2, 121,
	-1, 253,
	68, 12,
	-2, 121,
	-1, 261,
	68, 12,
	-2, 121,
}

const yyPrivate = 57344

const yyLast = 1026

var yyAct = [...]int{

	68, 23, 12, 171, 74, 145, 98, 178, 197, 55,
	13, 67, 5, 143, 5, 173, 245, 195, 235, 221,
	76, 201, 77, 267, 188, 264, 260, 91, 249, 240,
	65, 238, 234, 232, 229, 225, 223, 218, 115, 253,
	251, 191, 163, 100, 101, 102, 103, 104, 105, 106,
	196, 176, 5, 112, 113, 114, 190, 14, 119, 136,
	137, 138, 139, 140, 141, 142, 4, 144, 163, 257,
	36, 246, 189, 118, 222, 73, 75, 187, 73, 183,
	247, 146, 162, 73, 147, 73, 73, 88, 87, 40,
	41, 97, 73, 73, 89, 149, 180, 145, 239, 181,
	158, 157, 168, 159, 161, 36, 150, 84, 73, 73,
	15, 89, 17, 85, 86, 133, 85, 86, 46, 48,
	57, 2, 166, 117, 3, 170, 175, 8, 237, 39,
	214, 99, 38, 179, 167, 37, 182, 73, 73, 73,
	73, 73, 73, 107, 78, 177, 134, 148, 76, 19,
	41, 185, 120, 79, 80, 81, 82, 83, 96, 45,
	47, 108, 110, 109, 111, 193, 119, 73, 123, 73,
	56, 7, 6, 204, 1, 0, 202, 116, 84, 199,
	0, 194, 133, 0, 205, 212, 207, 85, 86, 0,
	0, 211, 208, 0, 0, 0, 215, 0, 119, 0,
	0, 213, 179, 0, 75, 0, 216, 73, 0, 220,
	219, 76, 119, 217, 0, 0, 224, 0, 0, 0,
	227, 0, 179, 133, 231, 0, 0, 228, 0, 236,
	133, 0, 119, 0, 0, 198, 179, 241, 165, 0,
	242, 0, 0, 244, 0, 0, 179, 243, 254, 198,
	248, 0, 250, 252, 0, 133, 0, 0, 119, 255,
	0, 256, 22, 258, 119, 0, 184, 226, 0, 133,
	0, 265, 0, 262, 0, 0, 0, 0, 0, 266,
	71, 0, 198, 84, 192, 0, 0, 0, 90, 133,
	92, 93, 85, 86, 0, 198, 0, 94, 95, 0,
	0, 0, 209, 0, 0, 0, 0, 0, 198, 0,
	198, 0, 0, 0, 0, 133, 0, 198, 0, 0,
	0, 133, 0, 0, 263, 24, 29, 33, 30, 31,
	32, 81, 82, 83, 0, 0, 0, 0, 233, 0,
	27, 28, 151, 152, 153, 154, 155, 156, 121, 0,
	0, 0, 122, 16, 0, 135, 0, 124, 125, 126,
	127, 128, 69, 131, 132, 129, 130, 0, 18, 0,
	0, 0, 71, 0, 71, 0, 0, 34, 35, 0,
	21, 0, 25, 0, 20, 0, 26, 0, 261, 0,
	5, 24, 29, 33, 30, 31, 32, 43, 44, 46,
	48, 51, 54, 0, 0, 0, 27, 28, 0, 0,
	0, 0, 186, 0, 121, 0, 0, 0, 122, 16,
	0, 135, 0, 124, 125, 126, 127, 128, 69, 131,
	132, 129, 130, 0, 18, 0, 0, 0, 0, 50,
	45, 47, 0, 34, 35, 0, 21, 0, 25, 0,
	20, 0, 172, 0, 169, 24, 29, 33, 30, 31,
	32, 43, 44, 46, 48, 51, 0, 0, 0, 0,
	27, 28, 0, 0, 0, 0, 0, 0, 121, 0,
	0, 0, 122, 16, 0, 135, 0, 124, 125, 126,
	127, 128, 69, 131, 132, 129, 130, 0, 18, 0,
	0, 0, 0, 50, 45, 47, 0, 34, 35, 0,
	21, 0, 25, 0, 20, 0, 26, 203, 29, 33,
	30, 31, 32, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 27, 28, 0, 0, 0, 0, 0, 0,
	121, 0, 0, 0, 122, 16, 0, 135, 0, 124,
	125, 126, 127, 128, 69, 131, 132, 129, 130, 0,
	18, 59, 60, 61, 62, 63, 64, 0, 0, 34,
	35, 0, 21, 0, 25, 0, 20, 0, 26, 24,
	29, 33, 30, 31, 32, 0, 0, 0, 0, 58,
	0, 0, 0, 0, 27, 28, 24, 29, 33, 30,
	31, 32, 0, 0, 10, 11, 0, 16, 0, 0,
	0, 27, 28, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 18, 0, 72, 0, 0, 0, 0, 0,
	0, 34, 35, 69, 21, 0, 25, 0, 20, 0,
	26, 0, 13, 0, 0, 0, 0, 0, 34, 35,
	0, 0, 0, 25, 0, 20, 0, 26, 0, 0,
	0, 0, 70, 24, 29, 33, 30, 31, 32, 43,
	44, 46, 48, 51, 54, 53, 0, 0, 27, 28,
	0, 0, 0, 0, 0, 0, 49, 0, 0, 0,
	0, 16, 43, 44, 46, 48, 51, 54, 53, 0,
	69, 0, 0, 0, 0, 0, 18, 0, 52, 49,
	0, 50, 45, 47, 0, 34, 35, 0, 21, 0,
	25, 0, 20, 0, 26, 230, 174, 5, 0, 0,
	0, 52, 0, 0, 50, 45, 47, 24, 29, 33,
	30, 31, 32, 0, 0, 0, 0, 0, 210, 0,
	5, 0, 27, 28, 24, 29, 33, 30, 31, 32,
	0, 0, 0, 0, 0, 16, 0, 0, 0, 27,
	28, 0, 0, 0, 69, 0, 0, 0, 0, 0,
	18, 0, 16, 0, 0, 0, 0, 0, 0, 34,
	35, 69, 21, 0, 25, 0, 20, 18, 26, 0,
	0, 0, 0, 0, 0, 0, 34, 35, 0, 21,
	0, 25, 0, 20, 0, 66, 24, 29, 33, 30,
	31, 32, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 27, 28, 24, 29, 33, 30, 31, 32, 0,
	0, 0, 0, 0, 72, 0, 0, 0, 27, 28,
	0, 0, 0, 69, 0, 0, 0, 0, 0, 0,
	0, 72, 0, 0, 0, 0, 0, 0, 34, 35,
	69, 21, 0, 25, 0, 20, 0, 26, 43, 44,
	46, 48, 51, 54, 53, 34, 35, 0, 0, 0,
	25, 0, 20, 0, 26, 49, 43, 44, 46, 48,
	51, 54, 53, 0, 0, 0, 43, 44, 46, 48,
	51, 54, 53, 49, 0, 0, 0, 52, 0, 0,
	50, 45, 47, 49, 43, 44, 46, 48, 51, 54,
	53, 0, 0, 0, 206, 52, 0, 0, 50, 45,
	47, 49, 0, 0, 0, 52, 0, 0, 50, 45,
	47, 0, 200, 43, 44, 46, 48, 51, 54, 53,
	0, 259, 0, 52, 0, 0, 50, 45, 47, 0,
	49, 43, 44, 46, 48, 51, 54, 53, 0, 160,
	0, 43, 44, 46, 48, 51, 54, 53, 49, 0,
	0, 0, 52, 164, 0, 50, 45, 47, 49, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	52, 0, 42, 50, 45, 47, 0, 0, 0, 0,
	52, 0, 0, 50, 45, 47,
}
var yyPact = [...]int{

	-55, 575, -1000, -55, -1000, -1000, -1000, -57, -55, 85,
	-1000, -1000, 960, -55, 540, -1000, 750, 592, 829, -1000,
	829, -1000, 97, 23, 29, 829, 733, 829, 829, -1000,
	-1000, -1000, -1000, -1000, 829, 829, -1000, -55, -55, -1000,
	26, 127, 733, 733, 733, 733, 733, 733, 733, 139,
	812, 812, 733, 733, 733, -30, -55, 451, 733, 733,
	733, 733, 733, 733, 733, -1000, 733, 45, 970, 25,
	34, 97, 19, -1000, 40, -1000, 23, 23, 73, 829,
	829, 829, 829, 829, 829, -1000, -1000, 733, 733, 733,
	168, 913, -1000, -1000, 168, 168, -1000, 127, 16, -1000,
	970, 105, 105, -1000, -1000, -1000, -1000, -1000, -1000, 829,
	-1000, 829, 942, 386, 450, -1000, -1000, -55, -1000, 970,
	-1000, 733, 733, 67, 387, 659, -16, -1000, -1000, -1000,
	-1000, 733, 733, 540, 47, 733, 970, 970, 970, 970,
	970, 970, 970, 13, 913, -55, 829, 733, -1000, 829,
	-1000, 273, 273, 168, 168, 168, 94, 11, -47, 6,
	-1000, -10, -26, -55, 733, 451, 970, 45, -17, -55,
	885, -48, 513, -1000, -55, 867, -55, -1000, -1000, 970,
	733, -55, 681, -1000, 733, -1000, 97, -1000, -1000, -1000,
	-1000, -55, 126, 970, -1000, 733, -55, 451, -1000, -31,
	-55, 733, -50, 46, 913, -32, -55, -33, 45, 829,
	-55, 451, 450, -34, -1000, 658, -35, -55, -1000, -36,
	-51, 733, 124, -1000, -37, 61, -1000, -39, -55, -1000,
	-55, 451, -1000, -1000, -1000, 733, -53, 5, -1000, 15,
	-1000, -55, -40, -55, -27, 733, -28, 733, -1000, -1000,
	-55, -55, 3, -55, 895, -1000, -42, 321, -43, -1000,
	-1000, -55, -1000, 451, -1000, -45, -1000, -1000,
}
var yyPgo = [...]int{

	0, 174, 172, 171, 15, 3, 127, 170, 9, 168,
	0, 112, 152, 262, 1, 4, 57, 149, 7, 110,
	147, 11, 146, 13, 6, 120, 124, 66, 8, 144,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 2, 3, 3, 3, 3,
	3, 6, 8, 8, 7, 7, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 9, 9,
	9, 9, 9, 9, 12, 23, 23, 21, 21, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 29, 19, 11, 11, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 20, 20,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 14, 14, 14, 14, 22, 22, 15, 15,
	16, 16, 17, 17, 24, 24, 24, 5, 5, 18,
	18, 25, 25, 26, 26, 28, 28, 27,
}
var yyR2 = [...]int{

	0, 1, 2, 3, 2, 2, 5, 1, 1, 1,
	3, 3, 0, 2, 2, 4, 1, 1, 2, 1,
	2, 1, 4, 5, 9, 11, 9, 10, 4, 5,
	8, 1, 1, 1, 1, 9, 2, 2, 5, 6,
	7, 8, 5, 5, 3, 0, 1, 1, 4, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 1, 3, 3, 3, 3, 3, 3, 5, 3,
	3, 2, 1, 4, 3, 0, 3, 1, 2, 1,
	3, 3, 3, 3, 3, 3, 2, 2, 0, 2,
	2, 4, 4, 4, 7, 3, 2, 2, 1, 1,
	1, 1, 1, 1, 2, 2, 1, 4, 0, 1,
	1, 2, 4, 1, 0, 1, 4, 0, 1, 0,
	1, 0, 1, 1, 2, 0, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -25, -26, -27, 69, -2, -3, -6, 41,
	29, 30, -10, 67, -16, -19, 32, -11, 47, -17,
	63, 59, -13, -14, 4, 61, 65, 19, 20, 5,
	7, 8, 9, 6, 56, 57, -27, -6, -26, -25,
	4, 65, 52, 11, 12, 54, 13, 55, 14, 28,
	53, 15, 50, 17, 16, -8, -7, -25, 49, 21,
	22, 23, 24, 25, 26, -23, 65, -21, -10, 41,
	70, -13, 32, -16, -15, -16, -14, -14, -29, 56,
	57, 58, 59, 60, 10, 19, 20, 65, 64, 65,
	-13, -10, -13, -13, -13, -13, -25, 65, -24, 4,
	-10, -10, -10, -10, -10, -10, -10, 4, -19, -11,
	-19, -11, -10, -10, -10, 68, -25, -27, -4, -10,
	-12, 27, 31, -9, 36, 37, 38, 39, 40, 44,
	45, 42, 43, -16, -22, 34, -10, -10, -10, -10,
	-10, -10, -10, -23, -10, 52, 47, 65, -20, 55,
	33, -13, -13, -13, -13, -13, -13, -23, -21, -23,
	66, -24, 66, 52, 51, -25, -10, -21, 35, 67,
	-10, -5, 65, -4, 67, -10, 67, -23, -18, -10,
	49, 52, -10, 66, -25, -15, -13, 66, 71, 66,
	66, 67, -25, -10, -4, 34, 67, -28, -27, -8,
	67, 69, -5, 4, -10, -8, 67, -8, -21, -25,
	67, -28, -10, -8, 4, -10, -8, -4, 68, -8,
	-18, 69, 28, 68, -8, 68, -16, -8, -4, 68,
	67, -28, 68, -25, 68, 69, -18, 4, 68, 37,
	68, -28, -8, -4, -18, 69, 66, 65, -28, 68,
	-28, 67, -18, 67, -10, -28, -8, 66, -8, 66,
	68, 67, -4, -27, 68, -8, -4, 68,
}
var yyDef = [...]int{

	121, -2, 1, 122, 123, 127, 2, 0, 121, 0,
	7, 8, 9, -2, 103, 61, 45, 72, 108, 110,
	0, 75, 77, 79, 113, 0, 0, 0, 0, 98,
	99, 100, 101, 102, 0, 0, 124, 121, 4, 5,
	0, 114, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 121, 0, 0, 0,
	0, 0, 0, 0, 0, 71, 45, 46, 47, 0,
	0, 78, 0, 103, 88, -2, 0, 111, 0, 0,
	0, 0, 0, 0, 0, 86, 87, 45, 0, 45,
	90, 0, 96, 97, 104, 105, 3, 114, 0, 115,
	10, -2, -2, -2, -2, -2, -2, 56, 57, 58,
	59, 60, 0, 69, 70, 11, 13, -2, 14, 16,
	17, 0, 19, 21, 117, 0, 0, 31, 32, 33,
	34, 45, 119, -2, 0, 0, 49, 62, 63, 64,
	65, 66, 67, 0, 47, 121, 108, 45, 74, 0,
	76, 80, 81, 82, 83, 84, 85, 0, 0, 0,
	95, 0, 0, 121, 0, 0, 18, 20, 125, -2,
	16, 0, 117, 118, -2, 0, -2, 36, 37, 120,
	0, 121, 125, 92, 0, 73, 89, 93, 112, 91,
	6, -2, 0, 68, 15, 0, -2, 0, 126, 0,
	-2, 119, 0, 113, 16, 0, -2, 0, 44, 0,
	-2, 0, 48, 0, 116, 125, 0, 121, 22, 0,
	0, 119, 0, 28, 0, 0, -2, 0, 125, 94,
	-2, 0, 42, 43, 23, 119, 0, 0, 29, 0,
	38, 125, 0, 125, 0, 119, 0, 0, 39, 40,
	125, -2, 0, -2, 0, 41, 0, 0, 0, 30,
	24, -2, 26, 0, 35, 0, 27, 25,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 61, 3, 3, 63, 60, 3, 3,
	65, 66, 58, 56, 52, 57, 3, 59, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 51, 69,
	55, 49, 54, 50, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 64, 3, 71, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 67, 70, 68, 53,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 62,
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
		yyDollar = yyS[yypt-11 : yypt+1]
		{
			yyVAL.stmt = &ast.CForLoopStmt{Stmt1: yyDollar[3].stmt, Expr2: yyDollar[5].expr, Expr3: yyDollar[7].expr, Stmts: yyDollar[10].stmts}
		}
	case 26:
		yyDollar = yyS[yypt-9 : yypt+1]
		{
			yyVAL.stmt = &ast.CForLoopStmt{Stmt1: yyDollar[3].stmt, Expr2: yyDollar[5].expr, Expr3: yyDollar[7].expr, Stmts: []ast.Stmt{yyDollar[9].stmt}}
		}
	case 27:
		yyDollar = yyS[yypt-10 : yypt+1]
		{
			yyVAL.stmt = &ast.CForLoopStmt{Stmt1: yyDollar[3].stmt, Expr2: yyDollar[5].expr, Expr3: yyDollar[7].expr, Stmts: []ast.Stmt{yyDollar[10].stmt}}
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyDollar[3].stmts}
		}
	case 29:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyDollar[4].stmts, Expr: yyDollar[2].expr}
		}
	case 30:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.stmt = &ast.DoLoopStmt{Stmts: yyDollar[3].stmts, Expr: yyDollar[7].expr}
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmt = &ast.BreakStmt{}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmt = &ast.ContinueStmt{}
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmt = &ast.NextStmt{}
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmt = &ast.NextfileStmt{}
		}
	case 35:
		yyDollar = yyS[yypt-9 : yypt+1]
		{
			yyVAL.stmt = &ast.MapLoopStmt{KeyID: yyDollar[3].token.Literal, MapID: yyDollar[5].token.Literal, Stmts: yyDollar[8].stmts}
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmt = &ast.ReturnStmt{Exprs: yyDollar[2].exprs}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmt = &ast.ExitStmt{Expr: yyDollar[2].expr}
		}
	case 38:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			//fmt.Println("stmt_if:1")
			yyVAL.stmt = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].stmts, Else: nil}
		}
	case 39:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.stmt = &ast.IfStmt{If: yyDollar[2].expr, Then: []ast.Stmt{yyDollar[4].stmt}, Else: nil}
		}
	case 40:
		yyDollar = yyS[yypt-7 : yypt+1]
		{
			yyVAL.stmt.(*ast.IfStmt).ElseIf = append(yyVAL.stmt.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyDollar[4].expr, Then: yyDollar[6].stmts})
		}
	case 41:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.stmt.(*ast.IfStmt).ElseIf = append(yyVAL.stmt.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyDollar[4].expr, Then: []ast.Stmt{yyDollar[6].stmt}})
		}
	case 42:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			if yyVAL.stmt.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				//$$.(*ast.IfStmt).Else = append($$.(*ast.IfStmt).Else, $4...)
				yyVAL.stmt.(*ast.IfStmt).Else = yyDollar[4].stmts
			}
		}
	case 43:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			if yyVAL.stmt.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				yyVAL.stmt.(*ast.IfStmt).Else = []ast.Stmt{yyDollar[4].stmt}
			}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.AssExpr{Left: yyDollar[1].exprs, Right: yyDollar[3].exprs}
		}
	case 45:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exprs = yyDollar[1].exprs
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 48:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.AssExpr{Left: []ast.Expr{yyDollar[1].expr}, Right: []ast.Expr{yyDollar[3].expr}}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "==", Right: yyDollar[3].expr}
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "!=", Right: yyDollar[3].expr}
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: ">", Right: yyDollar[3].expr}
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: ">=", Right: yyDollar[3].expr}
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "<", Right: yyDollar[3].expr}
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "<=", Right: yyDollar[3].expr}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.ContainKeyExpr{KeyExpr: yyDollar[1].expr, MapID: yyDollar[3].token.Literal}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.MatchExpr{Expr: yyDollar[1].expr, RegExpr: yyDollar[3].expr}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.MatchExpr{Expr: yyDollar[1].expr, RegExpr: yyDollar[3].expr}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: &ast.MatchExpr{Expr: yyDollar[1].expr, RegExpr: yyDollar[3].expr}}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: &ast.MatchExpr{Expr: yyDollar[1].expr, RegExpr: yyDollar[3].expr}}
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &ast.MatchExpr{Expr: &defaultExpr, RegExpr: yyDollar[1].expr}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "+=", Right: yyDollar[3].expr}
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "-=", Right: yyDollar[3].expr}
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "*=", Right: yyDollar[3].expr}
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "/=", Right: yyDollar[3].expr}
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "%=", Right: yyDollar[3].expr}
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "^=", Right: yyDollar[3].expr}
		}
	case 68:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.expr = &ast.TriOpExpr{Cond: yyDollar[1].expr, Then: yyDollar[3].expr, Else: yyDollar[5].expr}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "||", Right: yyDollar[3].expr}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "&&", Right: yyDollar[3].expr}
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.CallExpr{Name: "printf", SubExprs: yyDollar[2].exprs}
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 73:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &ast.GetlineExpr{Command: yyDollar[1].expr, Var: yyDollar[4].expr}
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.GetlineExpr{Var: yyDollar[2].expr, Redir: yyDollar[3].expr}
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			//fmt.Println("YACC: want regexp!!")
			inRegExp = true
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			//fmt.Println("FINISH")
			yyVAL.expr = &ast.StringExpr{Literal: yyDollar[3].token.Literal}
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "CAT", Right: yyDollar[2].expr}
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "+", Right: yyDollar[3].expr}
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "-", Right: yyDollar[3].expr}
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "*", Right: yyDollar[3].expr}
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "/", Right: yyDollar[3].expr}
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "%", Right: yyDollar[3].expr}
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "^", Right: yyDollar[3].expr}
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "++", After: true}
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "--", After: true}
		}
	case 88:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.expr = nil
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
		}
	case 91:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].token.Literal, SubExprs: yyDollar[3].exprs}
		}
	case 92:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].token.Literal, SubExprs: yyDollar[3].exprs}
		}
	case 93:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &ast.AnonymousCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
		}
	case 94:
		yyDollar = yyS[yypt-7 : yypt+1]
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyDollar[3].identArgs, Stmts: yyDollar[6].stmts}
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.ParentExpr{SubExpr: yyDollar[2].expr}
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[2].expr, Operator: "++"}
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[2].expr, Operator: "--"}
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &ast.NumExpr{Literal: yyDollar[1].token.Literal}
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &ast.ConstExpr{Literal: yyDollar[1].token.Literal}
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &ast.ConstExpr{Literal: yyDollar[1].token.Literal}
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &ast.ConstExpr{Literal: yyDollar[1].token.Literal}
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &ast.StringExpr{Literal: yyDollar[1].token.Literal}
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 104:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "+", Expr: yyDollar[2].expr}
		}
	case 105:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 107:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 108:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.expr = nil
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.FieldExpr{Expr: yyDollar[2].expr}
		}
	case 112:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &ast.ItemExpr{Expr: yyDollar[1].expr, Index: yyDollar[3].exprs}
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &ast.IdentExpr{Literal: yyDollar[1].token.Literal}
		}
	case 114:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.identArgs = []string{}
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.identArgs = []string{yyDollar[1].token.Literal}
		}
	case 116:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.identArgs = append(yyDollar[1].identArgs, yyDollar[4].token.Literal)
		}
	case 117:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.stmt = nil
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 119:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.expr = nil
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = yyDollar[1].expr
		}
	}
	goto yystack /* stack new state and value */
}
