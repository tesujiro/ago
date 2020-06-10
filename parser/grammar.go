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
const UNARY = 57390

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
	"'='",
	"'?'",
	"':'",
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
	"','",
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
	67, 12,
	-2, 124,
	-1, 75,
	62, 106,
	63, 106,
	-2, 112,
	-1, 101,
	11, 0,
	12, 0,
	-2, 53,
	-1, 102,
	11, 0,
	12, 0,
	-2, 54,
	-1, 103,
	13, 0,
	14, 0,
	52, 0,
	53, 0,
	-2, 55,
	-1, 104,
	13, 0,
	14, 0,
	52, 0,
	53, 0,
	-2, 56,
	-1, 105,
	13, 0,
	14, 0,
	52, 0,
	53, 0,
	-2, 57,
	-1, 106,
	13, 0,
	14, 0,
	52, 0,
	53, 0,
	-2, 58,
	-1, 117,
	67, 126,
	-2, 124,
	-1, 133,
	65, 109,
	-2, 106,
	-1, 169,
	67, 12,
	-2, 124,
	-1, 174,
	67, 12,
	-2, 124,
	-1, 176,
	67, 12,
	-2, 124,
	-1, 192,
	67, 12,
	-2, 124,
	-1, 197,
	67, 12,
	-2, 124,
	-1, 201,
	67, 12,
	-2, 124,
	-1, 207,
	67, 12,
	-2, 124,
	-1, 211,
	67, 12,
	-2, 124,
	-1, 228,
	62, 106,
	63, 106,
	-2, 110,
	-1, 232,
	67, 12,
	-2, 124,
	-1, 256,
	67, 12,
	-2, 124,
	-1, 258,
	67, 12,
	-2, 124,
	-1, 270,
	67, 12,
	-2, 124,
}

const yyPrivate = 57344

const yyLast = 1204

var yyAct = [...]int{

	14, 23, 74, 67, 98, 145, 178, 143, 277, 171,
	189, 13, 196, 5, 198, 5, 248, 55, 73, 75,
	76, 73, 77, 237, 65, 223, 73, 202, 73, 73,
	273, 269, 253, 242, 240, 73, 73, 43, 44, 46,
	48, 51, 54, 53, 197, 258, 5, 173, 236, 234,
	231, 73, 73, 227, 49, 225, 220, 115, 133, 256,
	192, 191, 163, 145, 4, 176, 162, 163, 36, 180,
	264, 249, 224, 190, 188, 52, 184, 50, 45, 47,
	73, 73, 73, 73, 73, 73, 181, 88, 87, 57,
	2, 42, 158, 40, 250, 157, 147, 159, 39, 41,
	97, 89, 161, 36, 149, 118, 15, 89, 146, 241,
	73, 168, 73, 150, 46, 48, 119, 239, 12, 84,
	215, 117, 85, 86, 3, 133, 167, 96, 85, 86,
	8, 99, 38, 68, 17, 107, 78, 134, 37, 177,
	148, 19, 120, 91, 123, 56, 116, 75, 76, 186,
	73, 7, 41, 45, 47, 6, 1, 108, 110, 100,
	101, 102, 103, 104, 105, 106, 133, 0, 0, 112,
	113, 114, 0, 133, 0, 136, 137, 138, 139, 140,
	141, 142, 203, 144, 209, 109, 111, 200, 0, 0,
	0, 0, 206, 0, 208, 0, 0, 0, 0, 133,
	0, 0, 0, 0, 68, 68, 68, 165, 0, 222,
	214, 228, 76, 195, 84, 218, 0, 0, 0, 221,
	0, 0, 0, 85, 86, 226, 0, 0, 0, 229,
	238, 133, 0, 199, 0, 185, 0, 0, 166, 68,
	0, 170, 175, 0, 247, 133, 219, 133, 68, 179,
	245, 0, 182, 193, 0, 257, 133, 0, 251, 0,
	81, 82, 83, 0, 68, 133, 0, 0, 0, 0,
	0, 210, 0, 133, 263, 267, 265, 268, 243, 0,
	0, 194, 0, 274, 0, 0, 0, 0, 275, 205,
	0, 0, 252, 0, 254, 244, 0, 68, 0, 0,
	212, 0, 213, 262, 0, 84, 0, 0, 199, 235,
	22, 255, 271, 216, 85, 86, 0, 260, 0, 179,
	276, 0, 0, 0, 0, 199, 0, 199, 71, 272,
	0, 0, 0, 199, 233, 0, 90, 0, 92, 93,
	179, 0, 0, 0, 261, 94, 95, 0, 0, 79,
	80, 81, 82, 83, 179, 0, 0, 43, 44, 46,
	48, 51, 54, 0, 0, 179, 0, 259, 0, 0,
	0, 0, 0, 24, 29, 33, 30, 31, 32, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 27, 28,
	151, 152, 153, 154, 155, 156, 121, 50, 45, 47,
	122, 16, 0, 135, 0, 124, 125, 126, 127, 128,
	69, 131, 132, 129, 130, 0, 18, 0, 0, 0,
	71, 0, 71, 34, 35, 0, 21, 0, 25, 0,
	20, 0, 26, 0, 0, 270, 0, 5, 24, 29,
	33, 30, 31, 32, 43, 44, 46, 48, 51, 0,
	0, 0, 0, 27, 28, 0, 0, 0, 0, 0,
	187, 121, 0, 0, 0, 122, 16, 0, 135, 0,
	124, 125, 126, 127, 128, 69, 131, 132, 129, 130,
	0, 18, 0, 0, 50, 45, 47, 0, 34, 35,
	0, 21, 0, 25, 0, 20, 0, 172, 0, 0,
	169, 24, 29, 33, 30, 31, 32, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 27, 28, 0, 0,
	0, 0, 0, 0, 121, 0, 0, 0, 122, 16,
	0, 135, 0, 124, 125, 126, 127, 128, 69, 131,
	132, 129, 130, 0, 18, 59, 60, 61, 62, 63,
	64, 34, 35, 0, 21, 0, 25, 0, 20, 0,
	26, 204, 29, 33, 30, 31, 32, 0, 0, 0,
	0, 0, 58, 0, 0, 0, 27, 28, 0, 0,
	0, 0, 0, 0, 121, 0, 0, 0, 122, 16,
	0, 135, 0, 124, 125, 126, 127, 128, 69, 131,
	132, 129, 130, 24, 18, 0, 30, 31, 32, 0,
	0, 34, 35, 0, 21, 0, 25, 0, 20, 0,
	26, 0, 0, 0, 0, 0, 121, 0, 0, 0,
	122, 16, 0, 135, 0, 124, 125, 126, 127, 128,
	69, 131, 132, 129, 130, 0, 18, 24, 29, 33,
	30, 31, 32, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 27, 28, 0, 0, 0, 5, 0, 0,
	0, 0, 10, 11, 0, 16, 0, 24, 29, 33,
	30, 31, 32, 0, 9, 0, 0, 0, 0, 0,
	18, 0, 27, 28, 0, 0, 0, 34, 35, 0,
	21, 0, 25, 0, 20, 72, 26, 0, 0, 13,
	0, 0, 0, 0, 69, 24, 29, 33, 30, 31,
	32, 0, 0, 0, 0, 0, 0, 34, 35, 0,
	27, 28, 25, 0, 20, 0, 26, 0, 0, 0,
	0, 0, 70, 16, 0, 24, 29, 33, 30, 31,
	32, 0, 69, 0, 0, 0, 0, 0, 18, 0,
	27, 28, 0, 0, 0, 34, 35, 0, 21, 0,
	25, 0, 20, 16, 26, 0, 0, 174, 0, 0,
	0, 0, 69, 0, 0, 0, 0, 0, 18, 0,
	0, 0, 0, 0, 0, 34, 35, 0, 21, 0,
	25, 0, 20, 0, 26, 24, 29, 33, 30, 31,
	32, 43, 44, 46, 48, 51, 54, 53, 0, 0,
	27, 28, 24, 29, 33, 30, 31, 32, 49, 0,
	0, 0, 0, 16, 0, 0, 0, 27, 28, 0,
	0, 0, 69, 0, 0, 0, 0, 0, 18, 52,
	16, 50, 45, 47, 0, 34, 35, 0, 21, 69,
	25, 0, 20, 0, 217, 18, 232, 0, 0, 0,
	0, 0, 34, 35, 0, 21, 0, 25, 0, 20,
	0, 183, 24, 29, 33, 30, 31, 32, 43, 44,
	46, 48, 51, 54, 53, 0, 0, 27, 28, 0,
	0, 0, 0, 0, 0, 49, 0, 0, 0, 0,
	16, 0, 0, 0, 0, 0, 0, 0, 0, 69,
	0, 0, 0, 0, 0, 18, 52, 0, 50, 45,
	47, 0, 34, 35, 0, 21, 0, 25, 0, 20,
	0, 66, 0, 211, 43, 44, 46, 48, 51, 54,
	53, 43, 44, 46, 48, 51, 54, 53, 0, 0,
	0, 49, 0, 0, 0, 0, 0, 0, 49, 0,
	0, 0, 24, 29, 33, 30, 31, 32, 0, 0,
	0, 0, 52, 0, 50, 45, 47, 27, 28, 52,
	0, 50, 45, 47, 0, 0, 0, 0, 0, 207,
	72, 0, 0, 0, 0, 0, 201, 0, 0, 69,
	0, 43, 44, 46, 48, 51, 54, 53, 0, 0,
	0, 0, 34, 35, 0, 21, 0, 25, 49, 20,
	0, 26, 24, 29, 33, 30, 31, 32, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 27, 28, 52,
	0, 50, 45, 47, 0, 0, 0, 0, 0, 0,
	72, 0, 0, 0, 266, 0, 0, 0, 0, 69,
	0, 43, 44, 46, 48, 51, 54, 53, 0, 0,
	0, 0, 34, 35, 0, 0, 0, 25, 49, 20,
	0, 26, 0, 43, 44, 46, 48, 51, 54, 53,
	43, 44, 46, 48, 51, 54, 53, 0, 0, 52,
	49, 50, 45, 47, 0, 0, 0, 49, 0, 0,
	0, 0, 0, 0, 246, 0, 0, 0, 0, 0,
	0, 52, 0, 50, 45, 47, 0, 0, 52, 0,
	50, 45, 47, 0, 0, 0, 230, 0, 0, 0,
	0, 0, 0, 160, 43, 44, 46, 48, 51, 54,
	53, 43, 44, 46, 48, 51, 54, 53, 0, 0,
	0, 49, 0, 0, 0, 0, 0, 0, 49, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 52, 164, 50, 45, 47, 0, 0, 52,
	0, 50, 45, 47,
}
var yyPact = [...]int{

	-53, 643, -1000, -53, -1000, -1000, -1000, -55, -53, 89,
	-1000, -1000, 26, -53, 524, -1000, 878, 673, 1028, -1000,
	1028, -1000, 295, 25, 38, 1028, 741, 1028, 1028, -1000,
	-1000, -1000, -1000, -1000, 1028, 1028, -1000, -53, -53, -1000,
	37, 127, 741, 741, 741, 741, 741, 741, 741, 131,
	968, 968, 741, 741, 741, -10, -53, 497, 741, 741,
	741, 741, 741, 741, 741, -1000, 741, -2, 1150, 36,
	61, 295, 33, -1000, 51, -1000, 25, 25, 80, 1028,
	1028, 1028, 1028, 1028, 1028, -1000, -1000, 741, 741, 741,
	109, 1089, -1000, -1000, 109, 109, -1000, 127, 2, -1000,
	1150, 101, 101, -1000, -1000, -1000, -1000, -1000, -1000, 1028,
	-1000, 1028, 1143, 346, 433, -1000, -1000, -53, -1000, 1150,
	-1000, 741, 741, 76, 434, 711, -1, -1000, -1000, -1000,
	-1000, 741, 741, 524, 21, 818, 1150, 1150, 1150, 1150,
	1150, 1150, 1150, 12, 1089, -53, 1028, 741, -1000, 1028,
	-1000, 204, 204, 109, 109, 109, 103, 10, -60, 9,
	-1000, -3, -6, -53, 741, 497, 1150, -2, -22, -53,
	940, -41, 557, -1000, -53, 933, -53, -1000, -1000, 1150,
	741, -53, 877, 741, -1000, 741, -1000, 295, -1000, -1000,
	-1000, -1000, -53, 116, 1150, -1000, 801, -53, 497, -1000,
	-11, -53, 741, -43, 44, 1089, -12, -53, -14, -2,
	1028, -53, 1082, 1150, -17, -1000, 800, 741, -18, -53,
	-1000, -19, -45, 741, 113, -1000, -33, 72, -1000, -34,
	599, -1000, -53, 1060, -1000, -1000, -1000, 741, -52, 7,
	-1000, 31, -1000, -53, 497, -35, 599, -7, 741, -21,
	741, -1000, -53, -1000, -53, 497, -53, 6, -53, 1000,
	-53, -1000, -53, -36, 369, -37, -1000, -1000, -53, -1000,
	-53, -1000, 497, -1000, -1000, -59, -1000, -1000,
}
var yyPgo = [...]int{

	0, 156, 155, 151, 47, 9, 130, 145, 17, 144,
	116, 134, 142, 310, 1, 2, 0, 141, 6, 106,
	140, 3, 137, 7, 4, 89, 124, 64, 14, 136,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 2, 3, 3, 3, 3,
	3, 6, 8, 8, 7, 7, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 12, 23, 23,
	21, 21, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 29, 19,
	11, 11, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 20, 20, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 22,
	22, 15, 15, 16, 16, 17, 17, 24, 24, 24,
	5, 5, 18, 18, 25, 25, 26, 26, 28, 28,
	27,
}
var yyR2 = [...]int{

	0, 1, 2, 3, 2, 2, 5, 1, 1, 1,
	3, 3, 0, 2, 2, 4, 1, 1, 2, 1,
	2, 1, 4, 5, 9, 11, 9, 10, 4, 5,
	8, 1, 1, 1, 1, 9, 2, 2, 5, 6,
	6, 8, 7, 8, 10, 5, 5, 3, 0, 1,
	1, 4, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 1, 3, 3, 3, 3, 3,
	3, 5, 3, 3, 2, 1, 4, 3, 0, 3,
	1, 2, 1, 3, 3, 3, 3, 3, 3, 2,
	2, 0, 2, 2, 4, 4, 4, 7, 3, 2,
	2, 1, 1, 1, 1, 1, 1, 2, 2, 1,
	4, 0, 1, 1, 2, 4, 1, 0, 1, 4,
	0, 1, 0, 1, 0, 1, 1, 2, 0, 1,
	1,
}
var yyChk = [...]int{

	-1000, -1, -25, -26, -27, 68, -2, -3, -6, 41,
	29, 30, -10, 66, -16, -19, 32, -11, 47, -17,
	61, 57, -13, -14, 4, 59, 63, 19, 20, 5,
	7, 8, 9, 6, 54, 55, -27, -6, -26, -25,
	4, 63, 65, 11, 12, 52, 13, 53, 14, 28,
	51, 15, 49, 17, 16, -8, -7, -25, 48, 21,
	22, 23, 24, 25, 26, -23, 63, -21, -10, 41,
	69, -13, 32, -16, -15, -16, -14, -14, -29, 54,
	55, 56, 57, 58, 10, 19, 20, 63, 62, 63,
	-13, -10, -13, -13, -13, -13, -25, 63, -24, 4,
	-10, -10, -10, -10, -10, -10, -10, 4, -19, -11,
	-19, -11, -10, -10, -10, 67, -25, -27, -4, -10,
	-12, 27, 31, -9, 36, 37, 38, 39, 40, 44,
	45, 42, 43, -16, -22, 34, -10, -10, -10, -10,
	-10, -10, -10, -23, -10, 65, 47, 63, -20, 53,
	33, -13, -13, -13, -13, -13, -13, -23, -21, -23,
	64, -24, 64, 65, 50, -25, -10, -21, 35, 66,
	-10, -5, 63, -4, 66, -10, 66, -23, -18, -10,
	48, 65, -10, 63, 64, -25, -15, -13, 64, 70,
	64, 64, 66, -25, -10, -4, 34, 66, -28, -27,
	-8, 66, 68, -5, 4, -10, -8, 66, -8, -21,
	-25, 66, -10, -10, -8, 4, -10, 63, -8, -4,
	67, -8, -18, 68, 28, 67, -8, 67, -16, -8,
	64, 67, 66, -10, 67, -25, 67, 68, -18, 4,
	67, 37, 67, -4, -27, -8, 64, -18, 68, 64,
	63, -28, -4, 67, -4, -27, 66, -18, 66, -10,
	-27, -25, -4, -8, 64, -8, 64, -28, -28, 67,
	66, -4, -27, 67, -28, -8, -4, 67,
}
var yyDef = [...]int{

	124, -2, 1, 125, 126, 130, 2, 0, 124, 0,
	7, 8, 9, -2, 106, 64, 48, 75, 111, 113,
	0, 78, 80, 82, 116, 0, 0, 0, 0, 101,
	102, 103, 104, 105, 0, 0, 127, 124, 4, 5,
	0, 117, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 124, 0, 0, 0,
	0, 0, 0, 0, 0, 74, 48, 49, 50, 0,
	0, 81, 0, 106, 91, -2, 0, 114, 0, 0,
	0, 0, 0, 0, 0, 89, 90, 48, 0, 48,
	93, 0, 99, 100, 107, 108, 3, 117, 0, 118,
	10, -2, -2, -2, -2, -2, -2, 59, 60, 61,
	62, 63, 0, 72, 73, 11, 13, -2, 14, 16,
	17, 0, 19, 21, 120, 0, 0, 31, 32, 33,
	34, 48, 122, -2, 0, 0, 52, 65, 66, 67,
	68, 69, 70, 0, 50, 124, 111, 48, 77, 0,
	79, 83, 84, 85, 86, 87, 88, 0, 0, 0,
	98, 0, 0, 124, 0, 0, 18, 20, 128, -2,
	16, 0, 120, 121, -2, 0, -2, 36, 37, 123,
	0, 124, 0, 0, 95, 0, 76, 92, 96, 115,
	94, 6, -2, 0, 71, 15, 0, -2, 0, 129,
	0, -2, 122, 0, 116, 16, 0, -2, 0, 47,
	0, -2, 0, 51, 0, 119, 0, 0, 0, 124,
	22, 0, 0, 122, 0, 28, 0, 0, -2, 0,
	98, 97, -2, 0, 45, 46, 23, 122, 0, 0,
	29, 0, 38, 128, 0, 0, 98, 0, 122, 0,
	0, 39, 40, 42, 124, 0, -2, 0, -2, 0,
	128, 43, 128, 0, 0, 0, 30, 41, 128, 24,
	-2, 26, 0, 35, 44, 0, 27, 25,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 59, 3, 3, 61, 58, 3, 3,
	63, 64, 56, 54, 65, 55, 3, 57, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 50, 68,
	53, 48, 52, 49, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 62, 3, 70, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 66, 69, 67, 51,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 60,
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
			yyVAL.stmt = &ast.IfStmt{If: yyDollar[3].expr, Then: []ast.Stmt{yyDollar[5].stmt}, Else: nil}
		}
	case 40:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.stmt = &ast.IfStmt{If: yyDollar[3].expr, Then: []ast.Stmt{yyDollar[6].stmt}, Else: nil}
		}
	case 41:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.stmt = &ast.IfStmt{If: yyDollar[3].expr, Then: []ast.Stmt{yyDollar[6].stmt}, Else: nil}
		}
	case 42:
		yyDollar = yyS[yypt-7 : yypt+1]
		{
			yyVAL.stmt.(*ast.IfStmt).ElseIf = append(yyVAL.stmt.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyDollar[4].expr, Then: yyDollar[6].stmts})
		}
	case 43:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.stmt.(*ast.IfStmt).ElseIf = append(yyVAL.stmt.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyDollar[5].expr, Then: []ast.Stmt{yyDollar[7].stmt}})
		}
	case 44:
		yyDollar = yyS[yypt-10 : yypt+1]
		{
			yyVAL.stmt.(*ast.IfStmt).ElseIf = append(yyVAL.stmt.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyDollar[5].expr, Then: []ast.Stmt{yyDollar[8].stmt}})
		}
	case 45:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			if yyVAL.stmt.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				//$$.(*ast.IfStmt).Else = append($$.(*ast.IfStmt).Else, $4...)
				yyVAL.stmt.(*ast.IfStmt).Else = yyDollar[4].stmts
			}
		}
	case 46:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			if yyVAL.stmt.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				yyVAL.stmt.(*ast.IfStmt).Else = []ast.Stmt{yyDollar[4].stmt}
			}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.AssExpr{Left: yyDollar[1].exprs, Right: yyDollar[3].exprs}
		}
	case 48:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exprs = yyDollar[1].exprs
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 51:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.AssExpr{Left: []ast.Expr{yyDollar[1].expr}, Right: []ast.Expr{yyDollar[3].expr}}
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "==", Right: yyDollar[3].expr}
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "!=", Right: yyDollar[3].expr}
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: ">", Right: yyDollar[3].expr}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: ">=", Right: yyDollar[3].expr}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "<", Right: yyDollar[3].expr}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "<=", Right: yyDollar[3].expr}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.ContainKeyExpr{KeyExpr: yyDollar[1].expr, MapID: yyDollar[3].token.Literal}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.MatchExpr{Expr: yyDollar[1].expr, RegExpr: yyDollar[3].expr}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.MatchExpr{Expr: yyDollar[1].expr, RegExpr: yyDollar[3].expr}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: &ast.MatchExpr{Expr: yyDollar[1].expr, RegExpr: yyDollar[3].expr}}
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: &ast.MatchExpr{Expr: yyDollar[1].expr, RegExpr: yyDollar[3].expr}}
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &ast.MatchExpr{Expr: &defaultExpr, RegExpr: yyDollar[1].expr}
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "+=", Right: yyDollar[3].expr}
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "-=", Right: yyDollar[3].expr}
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "*=", Right: yyDollar[3].expr}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "/=", Right: yyDollar[3].expr}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "%=", Right: yyDollar[3].expr}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "^=", Right: yyDollar[3].expr}
		}
	case 71:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.expr = &ast.TriOpExpr{Cond: yyDollar[1].expr, Then: yyDollar[3].expr, Else: yyDollar[5].expr}
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "||", Right: yyDollar[3].expr}
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "&&", Right: yyDollar[3].expr}
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.CallExpr{Name: "printf", SubExprs: yyDollar[2].exprs}
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 76:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &ast.GetlineExpr{Command: yyDollar[1].expr, Var: yyDollar[4].expr}
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.GetlineExpr{Var: yyDollar[2].expr, Redir: yyDollar[3].expr}
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			//fmt.Println("YACC: want regexp!!")
			inRegExp = true
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			//fmt.Println("FINISH")
			yyVAL.expr = &ast.StringExpr{Literal: yyDollar[3].token.Literal}
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "CAT", Right: yyDollar[2].expr}
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "+", Right: yyDollar[3].expr}
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "-", Right: yyDollar[3].expr}
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "*", Right: yyDollar[3].expr}
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "/", Right: yyDollar[3].expr}
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "%", Right: yyDollar[3].expr}
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "^", Right: yyDollar[3].expr}
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "++", After: true}
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "--", After: true}
		}
	case 91:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.expr = nil
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 93:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
		}
	case 94:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].token.Literal, SubExprs: yyDollar[3].exprs}
		}
	case 95:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].token.Literal, SubExprs: yyDollar[3].exprs}
		}
	case 96:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &ast.AnonymousCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
		}
	case 97:
		yyDollar = yyS[yypt-7 : yypt+1]
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyDollar[3].identArgs, Stmts: yyDollar[6].stmts}
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.ParentExpr{SubExpr: yyDollar[2].expr}
		}
	case 99:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[2].expr, Operator: "++"}
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[2].expr, Operator: "--"}
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &ast.NumExpr{Literal: yyDollar[1].token.Literal}
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &ast.ConstExpr{Literal: yyDollar[1].token.Literal}
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &ast.ConstExpr{Literal: yyDollar[1].token.Literal}
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &ast.ConstExpr{Literal: yyDollar[1].token.Literal}
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &ast.StringExpr{Literal: yyDollar[1].token.Literal}
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 107:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "+", Expr: yyDollar[2].expr}
		}
	case 108:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 110:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
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
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 114:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.FieldExpr{Expr: yyDollar[2].expr}
		}
	case 115:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &ast.ItemExpr{Expr: yyDollar[1].expr, Index: yyDollar[3].exprs}
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &ast.IdentExpr{Literal: yyDollar[1].token.Literal}
		}
	case 117:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.identArgs = []string{}
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.identArgs = []string{yyDollar[1].token.Literal}
		}
	case 119:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.identArgs = append(yyDollar[1].identArgs, yyDollar[4].token.Literal)
		}
	case 120:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.stmt = nil
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 122:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.expr = nil
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = yyDollar[1].expr
		}
	}
	goto yystack /* stack new state and value */
}
