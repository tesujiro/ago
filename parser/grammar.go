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
	-1, 83,
	62, 106,
	63, 106,
	-2, 112,
	-1, 106,
	67, 126,
	-2, 124,
	-1, 122,
	65, 109,
	-2, 106,
	-1, 170,
	67, 12,
	-2, 124,
	-1, 175,
	67, 12,
	-2, 124,
	-1, 177,
	67, 12,
	-2, 124,
	-1, 193,
	67, 12,
	-2, 124,
	-1, 198,
	67, 12,
	-2, 124,
	-1, 202,
	67, 12,
	-2, 124,
	-1, 208,
	67, 12,
	-2, 124,
	-1, 212,
	67, 12,
	-2, 124,
	-1, 229,
	62, 106,
	63, 106,
	-2, 110,
	-1, 233,
	67, 12,
	-2, 124,
	-1, 257,
	67, 12,
	-2, 124,
	-1, 259,
	67, 12,
	-2, 124,
	-1, 271,
	67, 12,
	-2, 124,
}

const yyPrivate = 57344

const yyLast = 1035

var yyAct = [...]int{

	14, 22, 172, 98, 174, 82, 5, 197, 75, 152,
	249, 199, 150, 13, 190, 5, 238, 224, 203, 81,
	83, 84, 81, 85, 278, 274, 81, 270, 81, 81,
	73, 46, 254, 243, 241, 81, 81, 179, 237, 198,
	4, 5, 235, 232, 36, 45, 44, 45, 44, 122,
	228, 226, 221, 107, 45, 44, 104, 81, 81, 81,
	81, 81, 81, 259, 81, 81, 81, 81, 81, 81,
	81, 81, 257, 108, 193, 12, 192, 164, 43, 36,
	43, 181, 45, 44, 48, 2, 177, 43, 106, 163,
	164, 76, 152, 39, 265, 233, 158, 212, 182, 157,
	91, 162, 160, 250, 208, 16, 45, 44, 18, 191,
	45, 44, 225, 189, 122, 43, 100, 101, 102, 103,
	168, 185, 96, 125, 126, 127, 128, 129, 130, 131,
	40, 42, 105, 178, 251, 45, 44, 87, 86, 43,
	154, 81, 41, 43, 81, 45, 44, 89, 151, 45,
	44, 97, 89, 156, 83, 84, 202, 81, 267, 187,
	76, 76, 153, 76, 242, 169, 45, 122, 43, 139,
	142, 196, 140, 143, 122, 159, 204, 240, 43, 71,
	72, 216, 43, 247, 167, 76, 99, 171, 176, 41,
	210, 166, 138, 231, 76, 180, 88, 161, 183, 123,
	122, 155, 201, 3, 220, 45, 44, 207, 20, 209,
	200, 38, 229, 84, 70, 50, 51, 52, 53, 54,
	55, 8, 109, 71, 72, 215, 112, 47, 76, 37,
	219, 7, 122, 6, 222, 1, 244, 186, 43, 195,
	227, 223, 49, 0, 230, 0, 122, 206, 122, 194,
	253, 0, 255, 0, 0, 76, 252, 122, 213, 0,
	214, 263, 239, 0, 0, 246, 122, 211, 0, 0,
	272, 217, 245, 268, 122, 269, 248, 180, 277, 0,
	0, 275, 45, 44, 0, 200, 70, 258, 256, 264,
	0, 266, 234, 0, 261, 71, 72, 0, 180, 70,
	15, 0, 200, 276, 200, 236, 273, 0, 71, 72,
	200, 0, 180, 0, 0, 43, 165, 0, 0, 79,
	0, 0, 0, 180, 0, 260, 90, 0, 92, 93,
	65, 66, 67, 68, 69, 94, 95, 0, 0, 0,
	262, 0, 0, 0, 0, 67, 68, 69, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 132, 133, 134,
	135, 136, 137, 0, 141, 141, 144, 145, 146, 147,
	148, 149, 24, 29, 33, 30, 31, 32, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 27, 28, 0,
	0, 0, 0, 0, 0, 110, 0, 0, 0, 111,
	17, 0, 124, 0, 113, 114, 115, 116, 117, 77,
	120, 121, 118, 119, 0, 19, 0, 0, 0, 0,
	0, 0, 34, 35, 0, 23, 0, 25, 0, 21,
	0, 26, 0, 0, 271, 0, 5, 0, 0, 0,
	0, 79, 0, 0, 79, 24, 29, 33, 30, 31,
	32, 0, 0, 0, 0, 0, 0, 188, 0, 0,
	27, 28, 0, 0, 0, 0, 0, 0, 110, 0,
	0, 0, 111, 17, 0, 124, 0, 113, 114, 115,
	116, 117, 77, 120, 121, 118, 119, 0, 19, 0,
	0, 0, 0, 0, 0, 34, 35, 0, 23, 0,
	25, 0, 21, 0, 173, 0, 0, 170, 24, 29,
	33, 30, 31, 32, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 27, 28, 0, 0, 0, 0, 0,
	0, 110, 0, 0, 0, 111, 17, 0, 124, 0,
	113, 114, 115, 116, 117, 77, 120, 121, 118, 119,
	0, 19, 0, 0, 0, 0, 0, 0, 34, 35,
	0, 23, 0, 25, 0, 21, 0, 26, 205, 29,
	33, 30, 31, 32, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 27, 28, 0, 0, 0, 0, 0,
	0, 110, 0, 0, 0, 111, 17, 0, 124, 0,
	113, 114, 115, 116, 117, 77, 120, 121, 118, 119,
	24, 19, 0, 30, 31, 32, 0, 0, 34, 35,
	0, 23, 0, 25, 0, 21, 0, 26, 0, 0,
	0, 0, 0, 110, 0, 0, 0, 111, 17, 0,
	124, 0, 113, 114, 115, 116, 117, 77, 120, 121,
	118, 119, 0, 19, 24, 29, 33, 30, 31, 32,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 27,
	28, 0, 0, 0, 5, 0, 0, 0, 0, 10,
	11, 0, 17, 0, 24, 29, 33, 30, 31, 32,
	0, 9, 0, 0, 0, 0, 0, 19, 0, 27,
	28, 0, 0, 0, 34, 35, 0, 23, 0, 25,
	0, 21, 80, 26, 0, 0, 13, 0, 0, 0,
	0, 77, 24, 29, 33, 30, 31, 32, 0, 0,
	0, 0, 0, 0, 34, 35, 0, 27, 28, 25,
	0, 21, 0, 26, 0, 0, 0, 0, 0, 78,
	17, 0, 24, 29, 33, 30, 31, 32, 0, 77,
	0, 0, 0, 0, 0, 19, 0, 27, 28, 0,
	0, 0, 34, 35, 0, 23, 0, 25, 0, 21,
	17, 26, 0, 0, 175, 0, 0, 0, 0, 77,
	0, 0, 0, 0, 0, 19, 0, 0, 0, 0,
	0, 0, 34, 35, 0, 23, 0, 25, 0, 21,
	0, 26, 24, 29, 33, 30, 31, 32, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 27, 28, 24,
	29, 33, 30, 31, 32, 0, 0, 0, 0, 0,
	17, 0, 0, 0, 27, 28, 0, 0, 0, 77,
	0, 0, 0, 0, 0, 19, 0, 17, 0, 0,
	0, 0, 34, 35, 0, 23, 77, 25, 0, 21,
	0, 218, 19, 0, 0, 0, 0, 0, 0, 34,
	35, 0, 23, 0, 25, 0, 21, 0, 184, 24,
	29, 33, 30, 31, 32, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 27, 28, 24, 29, 33, 30,
	31, 32, 0, 0, 0, 0, 0, 17, 0, 0,
	0, 27, 28, 0, 0, 0, 77, 0, 0, 0,
	0, 0, 19, 0, 80, 0, 0, 0, 0, 34,
	35, 0, 23, 77, 25, 0, 21, 0, 74, 0,
	0, 0, 0, 0, 0, 0, 34, 35, 0, 23,
	0, 25, 0, 21, 0, 26, 24, 29, 33, 30,
	31, 32, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 27, 28, 0, 0, 0, 70, 56, 57, 59,
	61, 64, 0, 0, 80, 71, 72, 0, 0, 0,
	0, 0, 0, 77, 62, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 34, 35, 0, 0,
	0, 25, 0, 21, 0, 26, 0, 63, 58, 60,
	65, 66, 67, 68, 69,
}
var yyPact = [...]int{

	-62, 650, -1000, -62, -1000, -1000, -1000, -53, -62, 126,
	-1000, -1000, 66, -62, 194, 976, -1000, 885, 680, 962,
	-1000, 962, 75, -1000, 89, 962, 748, 962, 962, -1000,
	-1000, -1000, -1000, -1000, 962, 962, -1000, -62, -62, -1000,
	88, 182, 748, 748, 748, 748, -11, -62, 504, 748,
	748, 748, 748, 748, 748, 748, 962, 962, 962, 962,
	962, 962, 188, 902, 902, 962, 962, 962, 962, 962,
	962, -1000, -1000, -1000, 748, 27, 189, 79, 115, 276,
	77, -1000, 100, -1000, 75, 75, 748, 748, 142, 748,
	204, 133, -1000, -1000, 204, 204, -1000, 182, 25, -1000,
	189, 266, 150, -1000, -1000, -1000, -62, -1000, 189, -1000,
	748, 748, 130, 441, 718, 20, -1000, -1000, -1000, -1000,
	748, 748, 194, 33, 825, 189, 189, 189, 189, 189,
	189, 189, 276, 276, 276, 276, 276, 276, -1000, -1000,
	962, 276, -1000, 962, 289, 289, 204, 204, 204, 160,
	57, 133, -62, 962, 748, -1000, 962, 49, -56, -1000,
	45, -1000, 12, 8, -62, 748, 504, 189, 27, -27,
	-62, 90, -50, 564, -1000, -62, 38, -62, -1000, -1000,
	189, 748, -62, 31, 748, -1000, 748, -1000, 276, -1000,
	-1000, -1000, -1000, -62, 177, 189, -1000, 808, -62, 504,
	-1000, -15, -62, 748, -51, 84, 133, -16, -62, -17,
	27, 962, -62, 129, 189, -24, -1000, 29, 748, -25,
	-62, -1000, -29, -52, 748, 173, -1000, -33, 127, -1000,
	-34, 606, -1000, -62, 119, -1000, -1000, -1000, 748, -58,
	39, -1000, 71, -1000, -62, 504, -35, 606, 6, 748,
	-3, 748, -1000, -62, -1000, -62, 504, -62, 30, -62,
	94, -62, -1000, -62, -40, 368, -42, -1000, -1000, -62,
	-1000, -62, -1000, 504, -1000, -1000, -43, -1000, -1000,
}
var yyPgo = [...]int{

	0, 235, 233, 231, 4, 2, 221, 227, 31, 226,
	73, 108, 222, 300, 1, 5, 0, 208, 37, 105,
	201, 8, 199, 12, 3, 84, 203, 40, 11, 196,
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
	29, 30, -10, 66, -16, -13, -19, 32, -11, 47,
	-17, 61, -14, 57, 4, 59, 63, 19, 20, 5,
	7, 8, 9, 6, 54, 55, -27, -6, -26, -25,
	4, 63, 65, 49, 17, 16, -8, -7, -25, 48,
	21, 22, 23, 24, 25, 26, 11, 12, 52, 13,
	53, 14, 28, 51, 15, 54, 55, 56, 57, 58,
	10, 19, 20, -23, 63, -21, -10, 41, 69, -13,
	32, -16, -15, -16, -14, -14, 63, 62, -29, 63,
	-13, -10, -13, -13, -13, -13, -25, 63, -24, 4,
	-10, -10, -10, -10, 67, -25, -27, -4, -10, -12,
	27, 31, -9, 36, 37, 38, 39, 40, 44, 45,
	42, 43, -16, -22, 34, -10, -10, -10, -10, -10,
	-10, -10, -13, -13, -13, -13, -13, -13, 4, -19,
	-11, -13, -19, -11, -13, -13, -13, -13, -13, -13,
	-23, -10, 65, 47, 63, -20, 53, -23, -21, 33,
	-23, 64, -24, 64, 65, 50, -25, -10, -21, 35,
	66, -10, -5, 63, -4, 66, -10, 66, -23, -18,
	-10, 48, 65, -10, 63, 64, -25, -15, -13, 64,
	70, 64, 64, 66, -25, -10, -4, 34, 66, -28,
	-27, -8, 66, 68, -5, 4, -10, -8, 66, -8,
	-21, -25, 66, -10, -10, -8, 4, -10, 63, -8,
	-4, 67, -8, -18, 68, 28, 67, -8, 67, -16,
	-8, 64, 67, 66, -10, 67, -25, 67, 68, -18,
	4, 67, 37, 67, -4, -27, -8, 64, -18, 68,
	64, 63, -28, -4, 67, -4, -27, 66, -18, 66,
	-10, -27, -25, -4, -8, 64, -8, 64, -28, -28,
	67, 66, -4, -27, 67, -28, -8, -4, 67,
}
var yyDef = [...]int{

	124, -2, 1, 125, 126, 130, 2, 0, 124, 0,
	7, 8, 9, -2, 106, 80, 64, 48, 75, 111,
	113, 0, 82, 78, 116, 0, 0, 0, 0, 101,
	102, 103, 104, 105, 0, 0, 127, 124, 4, 5,
	0, 117, 0, 0, 0, 0, 0, 124, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 89, 90, 74, 48, 49, 50, 0, 0, 81,
	0, 106, 91, -2, 0, 114, 48, 0, 0, 48,
	93, 0, 99, 100, 107, 108, 3, 117, 0, 118,
	10, 0, 72, 73, 11, 13, -2, 14, 16, 17,
	0, 19, 21, 120, 0, 0, 31, 32, 33, 34,
	48, 122, -2, 0, 0, 52, 65, 66, 67, 68,
	69, 70, 53, 54, 55, 56, 57, 58, 59, 60,
	61, 80, 62, 63, 83, 84, 85, 86, 87, 88,
	0, 50, 124, 111, 48, 77, 0, 0, 0, 79,
	0, 98, 0, 0, 124, 0, 0, 18, 20, 128,
	-2, 16, 0, 120, 121, -2, 0, -2, 36, 37,
	123, 0, 124, 0, 0, 95, 0, 76, 92, 96,
	115, 94, 6, -2, 0, 71, 15, 0, -2, 0,
	129, 0, -2, 122, 0, 116, 16, 0, -2, 0,
	47, 0, -2, 0, 51, 0, 119, 0, 0, 0,
	124, 22, 0, 0, 122, 0, 28, 0, 0, -2,
	0, 98, 97, -2, 0, 45, 46, 23, 122, 0,
	0, 29, 0, 38, 128, 0, 0, 98, 0, 122,
	0, 0, 39, 40, 42, 124, 0, -2, 0, -2,
	0, 128, 43, 128, 0, 0, 0, 30, 41, 128,
	24, -2, 26, 0, 35, 44, 0, 27, 25,
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
