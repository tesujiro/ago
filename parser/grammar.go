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
const vars = 57390
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
	68, 12,
	-2, 119,
	-1, 67,
	64, 101,
	65, 101,
	-2, 107,
	-1, 108,
	68, 121,
	-2, 119,
	-1, 124,
	51, 104,
	-2, 101,
	-1, 171,
	68, 12,
	-2, 119,
	-1, 176,
	68, 12,
	-2, 119,
	-1, 179,
	68, 12,
	-2, 119,
	-1, 194,
	68, 12,
	-2, 119,
	-1, 199,
	68, 12,
	-2, 119,
	-1, 201,
	68, 12,
	-2, 119,
	-1, 207,
	68, 12,
	-2, 119,
	-1, 212,
	68, 12,
	-2, 119,
	-1, 227,
	64, 101,
	65, 101,
	-2, 105,
	-1, 230,
	68, 12,
	-2, 119,
	-1, 246,
	68, 12,
	-2, 119,
	-1, 248,
	68, 12,
	-2, 119,
	-1, 255,
	68, 12,
	-2, 119,
}

const yyPrivate = 57344

const yyLast = 840

var yyAct = [...]int{

	14, 66, 47, 181, 23, 173, 59, 100, 175, 16,
	60, 21, 12, 136, 13, 5, 5, 65, 67, 242,
	65, 134, 69, 233, 69, 65, 221, 65, 65, 68,
	202, 70, 259, 191, 65, 65, 92, 57, 257, 254,
	245, 239, 236, 232, 231, 229, 226, 223, 218, 106,
	124, 46, 45, 248, 102, 103, 104, 105, 109, 246,
	110, 127, 128, 129, 130, 131, 132, 133, 198, 135,
	46, 45, 65, 65, 65, 65, 65, 65, 65, 65,
	65, 65, 65, 65, 44, 65, 65, 194, 179, 154,
	156, 49, 2, 165, 155, 157, 159, 251, 46, 45,
	40, 199, 230, 44, 243, 192, 222, 163, 193, 190,
	158, 186, 160, 165, 244, 76, 124, 138, 80, 82,
	169, 212, 42, 168, 86, 87, 172, 177, 164, 99,
	98, 44, 89, 88, 182, 41, 90, 185, 67, 188,
	107, 65, 69, 90, 180, 140, 136, 137, 253, 68,
	46, 45, 238, 162, 46, 45, 65, 170, 65, 79,
	81, 71, 72, 73, 74, 75, 183, 46, 124, 184,
	46, 45, 86, 87, 200, 124, 197, 196, 110, 206,
	203, 235, 209, 44, 215, 205, 101, 44, 153, 208,
	210, 46, 45, 97, 46, 45, 42, 214, 213, 35,
	167, 207, 217, 44, 219, 201, 220, 46, 45, 216,
	224, 3, 227, 182, 4, 228, 69, 125, 37, 39,
	225, 76, 139, 68, 44, 234, 124, 44, 187, 43,
	86, 87, 182, 240, 237, 8, 110, 241, 76, 20,
	44, 161, 18, 38, 182, 111, 247, 86, 87, 250,
	114, 252, 124, 182, 37, 249, 63, 195, 258, 48,
	256, 7, 110, 108, 91, 6, 93, 94, 1, 73,
	74, 75, 0, 95, 96, 0, 211, 51, 52, 53,
	54, 55, 56, 0, 0, 0, 0, 76, 77, 78,
	80, 82, 85, 0, 0, 0, 86, 87, 0, 0,
	0, 0, 0, 0, 50, 83, 0, 0, 46, 45,
	0, 141, 142, 143, 144, 145, 146, 147, 148, 149,
	150, 151, 152, 0, 22, 28, 32, 29, 30, 31,
	84, 79, 81, 71, 72, 73, 74, 75, 0, 26,
	27, 44, 166, 0, 0, 0, 0, 112, 0, 0,
	0, 113, 15, 0, 126, 0, 115, 116, 117, 118,
	119, 61, 122, 123, 120, 121, 0, 17, 0, 0,
	0, 0, 0, 0, 0, 0, 33, 34, 0, 36,
	189, 24, 0, 19, 0, 25, 0, 255, 22, 28,
	32, 29, 30, 31, 0, 63, 0, 63, 0, 0,
	0, 0, 0, 26, 27, 0, 0, 0, 0, 0,
	0, 112, 0, 0, 0, 113, 15, 0, 126, 0,
	115, 116, 117, 118, 119, 61, 122, 123, 120, 121,
	0, 17, 0, 0, 0, 0, 0, 0, 0, 0,
	33, 34, 0, 36, 0, 24, 0, 19, 0, 174,
	76, 171, 204, 28, 32, 29, 30, 31, 0, 86,
	87, 0, 0, 0, 0, 0, 0, 26, 27, 0,
	0, 0, 0, 0, 0, 112, 0, 0, 0, 113,
	15, 0, 126, 0, 115, 116, 117, 118, 119, 61,
	122, 123, 120, 121, 0, 17, 71, 72, 73, 74,
	75, 0, 0, 0, 33, 34, 0, 36, 0, 24,
	0, 19, 0, 25, 22, 28, 32, 29, 30, 31,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 26,
	27, 0, 0, 0, 0, 0, 0, 112, 0, 0,
	0, 113, 15, 0, 126, 0, 115, 116, 117, 118,
	119, 61, 122, 123, 120, 121, 0, 17, 0, 0,
	0, 0, 0, 0, 0, 0, 33, 34, 0, 36,
	0, 24, 0, 19, 0, 25, 22, 28, 32, 29,
	30, 31, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 26, 27, 22, 28, 32, 29, 30, 31, 0,
	0, 0, 0, 0, 64, 0, 0, 0, 26, 27,
	0, 0, 0, 61, 0, 0, 0, 0, 10, 11,
	0, 15, 0, 0, 0, 0, 0, 0, 33, 34,
	9, 36, 0, 24, 0, 19, 17, 25, 0, 0,
	0, 0, 62, 0, 0, 33, 34, 0, 36, 0,
	24, 0, 19, 0, 25, 0, 13, 22, 28, 32,
	29, 30, 31, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 26, 27, 22, 28, 32, 29, 30, 31,
	0, 0, 0, 0, 0, 15, 0, 0, 0, 26,
	27, 0, 0, 0, 61, 0, 0, 0, 0, 0,
	17, 0, 15, 0, 0, 0, 0, 0, 0, 33,
	34, 61, 36, 0, 24, 0, 19, 17, 178, 0,
	176, 0, 0, 0, 0, 0, 33, 34, 0, 36,
	0, 24, 0, 19, 0, 25, 22, 28, 32, 29,
	30, 31, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 26, 27, 22, 28, 32, 29, 30, 31, 0,
	0, 0, 0, 0, 15, 0, 0, 0, 26, 27,
	0, 0, 0, 61, 0, 0, 0, 0, 0, 17,
	0, 64, 0, 0, 0, 0, 0, 0, 33, 34,
	61, 36, 0, 24, 0, 19, 22, 58, 0, 29,
	30, 31, 0, 0, 0, 33, 34, 0, 36, 0,
	24, 0, 19, 0, 25, 0, 0, 0, 0, 112,
	0, 0, 0, 113, 15, 0, 126, 0, 115, 116,
	117, 118, 119, 61, 122, 123, 120, 121, 0, 17,
}
var yyPact = [...]int{

	-54, 589, -1000, -54, -1000, -1000, -1000, -53, -54, 131,
	-1000, -1000, 178, -54, 256, 732, 572, 749, -1000, 749,
	277, 68, 71, -1000, 749, 670, 749, 749, -1000, -1000,
	-1000, -1000, -1000, 749, 749, -1000, -1000, -1000, -54, -54,
	-1000, 64, 182, 670, 670, 670, 670, -19, -54, 510,
	670, 670, 670, 670, 670, 670, 670, -1000, 670, 95,
	191, 57, 100, 440, 52, -1000, 90, -1000, 68, -1000,
	68, 749, 749, 749, 749, 749, 749, 749, 749, 749,
	749, 749, 749, 184, 749, 749, -1000, -1000, 670, 670,
	670, 228, 175, -1000, -1000, 228, 228, 120, -1000, 182,
	62, -1000, 191, 292, 151, -1000, -1000, -1000, -54, -1000,
	191, -1000, 670, 670, 122, 384, 653, 21, -1000, -1000,
	-1000, -1000, 670, 670, 256, 118, 670, 191, 191, 191,
	191, 191, 191, 191, 45, 175, -54, 749, 670, -1000,
	749, 211, 211, 228, 228, 228, 153, 105, 105, 440,
	440, 440, 440, -1000, -1000, 749, -1000, 749, 43, -38,
	39, -1000, -1000, 42, 20, -54, 670, 510, 191, 95,
	34, -54, 138, -39, 448, -1000, -54, 134, 670, -54,
	-1000, -1000, 191, 670, -54, 54, -1000, 670, -1000, 277,
	-1000, -1000, -1000, -1000, -54, 180, 191, -1000, 670, -54,
	-20, -54, 670, -43, 78, 175, -21, -54, 154, -22,
	95, 749, -54, -1000, -23, -1000, 35, -24, -1000, -25,
	-46, 670, 177, -1000, -26, 792, 115, -1000, -27, -1000,
	-54, -1000, -1000, 670, -50, 38, -1000, -1000, 49, -1000,
	-28, -8, 670, -14, 670, -1000, -54, 31, -54, 82,
	-29, 320, -30, -1000, -1000, -54, -1000, -1000, -36, -1000,
}
var yyPgo = [...]int{

	0, 268, 265, 261, 8, 5, 235, 259, 2, 250,
	10, 9, 245, 239, 11, 1, 0, 242, 3, 4,
	222, 6, 217, 21, 7, 91, 211, 214, 199, 193,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 2, 3, 3, 3, 3,
	3, 6, 8, 8, 7, 7, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 9, 9,
	9, 12, 23, 23, 21, 21, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	11, 11, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 29, 19, 20, 20, 14, 14, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	14, 14, 14, 14, 22, 22, 15, 15, 16, 16,
	17, 17, 24, 24, 24, 5, 5, 18, 18, 25,
	25, 26, 26, 27, 28,
}
var yyR2 = [...]int{

	0, 1, 2, 3, 2, 2, 5, 1, 1, 1,
	3, 3, 0, 2, 2, 4, 1, 1, 2, 1,
	2, 1, 4, 5, 9, 11, 9, 4, 5, 5,
	8, 1, 1, 1, 1, 9, 2, 2, 5, 7,
	5, 3, 0, 1, 1, 4, 3, 3, 3, 3,
	3, 3, 3, 5, 3, 3, 2, 1, 4, 3,
	1, 2, 1, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	1, 2, 2, 0, 3, 0, 2, 2, 1, 4,
	4, 4, 7, 3, 2, 2, 1, 1, 1, 1,
	1, 1, 2, 2, 1, 4, 0, 1, 1, 2,
	4, 1, 0, 1, 4, 0, 1, 0, 1, 0,
	1, 1, 2, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -25, -26, -27, 69, -2, -3, -6, 41,
	29, 30, -10, 67, -16, 32, -11, 47, -17, 63,
	-13, -14, 4, -19, 61, 65, 19, 20, 5, 7,
	8, 9, 6, 56, 57, -28, 59, -27, -6, -26,
	-25, 4, 65, 51, 49, 17, 16, -8, -7, -25,
	48, 21, 22, 23, 24, 25, 26, -23, 65, -21,
	-10, 41, 70, -13, 32, -16, -15, -16, -14, -19,
	-14, 56, 57, 58, 59, 60, 10, 11, 12, 54,
	13, 55, 14, 28, 53, 15, 19, 20, 65, 64,
	65, -13, -10, -13, -13, -13, -13, -29, -25, 65,
	-24, 4, -10, -10, -10, -10, 68, -25, -27, -4,
	-10, -12, 27, 31, -9, 36, 37, 38, 39, 40,
	44, 45, 42, 43, -16, -22, 34, -10, -10, -10,
	-10, -10, -10, -10, -23, -10, 51, 47, 65, -20,
	55, -13, -13, -13, -13, -13, -13, -13, -13, -13,
	-13, -13, -13, 4, -19, -11, -19, -11, -23, -21,
	-23, 66, 33, -24, 66, 51, 50, -25, -10, -21,
	35, 67, -10, -5, 65, -4, 67, -10, 65, 67,
	-23, -18, -10, 48, 51, -10, 66, -25, -15, -13,
	66, 71, 66, 66, 67, -25, -10, -4, 34, 67,
	-8, 67, 69, -5, 4, -10, -8, 67, -10, -8,
	-21, -25, 67, -10, -8, 4, -10, -8, 68, -8,
	-18, 69, 28, 68, -8, 66, 68, -16, -8, 68,
	67, 68, 68, 69, -18, 4, 68, -4, 37, 68,
	-8, -18, 69, 66, 65, 68, 67, -18, 67, -10,
	-8, 66, -8, 66, 68, 67, -4, 68, -8, 68,
}
var yyDef = [...]int{

	119, -2, 1, 120, 121, 123, 2, 0, 119, 0,
	7, 8, 9, -2, 101, 42, 57, 106, 108, 0,
	60, 62, 111, 80, 0, 0, 0, 0, 96, 97,
	98, 99, 100, 0, 0, 83, 124, 122, 119, 4,
	5, 0, 112, 0, 0, 0, 0, 0, 119, 0,
	0, 0, 0, 0, 0, 0, 0, 56, 42, 43,
	44, 0, 0, 61, 0, 101, 85, -2, 0, 88,
	109, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 81, 82, 42, 0,
	42, 87, 0, 94, 95, 102, 103, 0, 3, 112,
	0, 113, 10, 0, 54, 55, 11, 13, -2, 14,
	16, 17, 0, 19, 21, 115, 0, 0, 31, 32,
	33, 34, 42, 117, -2, 0, 0, 46, 47, 48,
	49, 50, 51, 52, 0, 44, 119, 106, 42, 59,
	0, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73, 74, 75, 76, 77, 78, 79, 0, 0,
	0, 93, 84, 0, 0, 119, 0, 0, 18, 20,
	0, -2, 16, 0, 115, 116, -2, 0, 0, -2,
	36, 37, 118, 0, 119, 0, 90, 0, 58, 86,
	91, 110, 89, 6, -2, 0, 53, 15, 0, -2,
	0, -2, 117, 0, 111, 16, 0, -2, 0, 0,
	41, 0, -2, 45, 0, 114, 0, 0, 22, 0,
	0, 117, 0, 27, 0, 93, 0, -2, 0, 92,
	-2, 40, 23, 117, 0, 0, 28, 29, 0, 38,
	0, 0, 117, 0, 0, 39, -2, 0, -2, 0,
	0, 0, 0, 30, 24, -2, 26, 35, 0, 25,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 61, 3, 3, 63, 60, 3, 3,
	65, 66, 58, 56, 51, 57, 3, 59, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 50, 69,
	55, 48, 54, 49, 3, 3, 3, 3, 3, 3,
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
	42, 43, 44, 45, 46, 47, 52, 62,
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
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyDollar[3].stmts}
		}
	case 28:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: yyDollar[4].stmts, Expr: yyDollar[2].expr}
		}
	case 29:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.stmt = &ast.LoopStmt{Stmts: []ast.Stmt{yyDollar[5].stmt}, Expr: yyDollar[3].expr}
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
			yyVAL.stmt = &ast.IfStmt{If: yyDollar[2].expr, Then: yyDollar[4].stmts, Else: nil}
		}
	case 39:
		yyDollar = yyS[yypt-7 : yypt+1]
		{
			yyVAL.stmt.(*ast.IfStmt).ElseIf = append(yyVAL.stmt.(*ast.IfStmt).ElseIf, &ast.IfStmt{If: yyDollar[4].expr, Then: yyDollar[6].stmts})
		}
	case 40:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			if yyVAL.stmt.(*ast.IfStmt).Else != nil {
				yylex.Error("multiple else statement")
			} else {
				//$$.(*ast.IfStmt).Else = append($$.(*ast.IfStmt).Else, $4...)
				yyVAL.stmt.(*ast.IfStmt).Else = yyDollar[4].stmts
			}
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.AssExpr{Left: yyDollar[1].exprs, Right: yyDollar[3].exprs}
		}
	case 42:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.exprs = []ast.Expr{}
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exprs = yyDollar[1].exprs
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 45:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.AssExpr{Left: []ast.Expr{yyDollar[1].expr}, Right: []ast.Expr{yyDollar[3].expr}}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "+=", Right: yyDollar[3].expr}
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "-=", Right: yyDollar[3].expr}
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "*=", Right: yyDollar[3].expr}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "/=", Right: yyDollar[3].expr}
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "%=", Right: yyDollar[3].expr}
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "^=", Right: yyDollar[3].expr}
		}
	case 53:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.expr = &ast.TriOpExpr{Cond: yyDollar[1].expr, Then: yyDollar[3].expr, Else: yyDollar[5].expr}
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "||", Right: yyDollar[3].expr}
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "&&", Right: yyDollar[3].expr}
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.CallExpr{Name: "printf", SubExprs: yyDollar[2].exprs}
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 58:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &ast.GetlineExpr{Command: yyDollar[1].expr, Var: yyDollar[4].expr}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.GetlineExpr{Var: yyDollar[2].expr, Redir: yyDollar[3].expr}
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "CAT", Right: yyDollar[2].expr}
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "+", Right: yyDollar[3].expr}
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "-", Right: yyDollar[3].expr}
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "*", Right: yyDollar[3].expr}
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "/", Right: yyDollar[3].expr}
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "%", Right: yyDollar[3].expr}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "^", Right: yyDollar[3].expr}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "==", Right: yyDollar[3].expr}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "!=", Right: yyDollar[3].expr}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: ">", Right: yyDollar[3].expr}
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: ">=", Right: yyDollar[3].expr}
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "<", Right: yyDollar[3].expr}
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.BinOpExpr{Left: yyDollar[1].expr, Operator: "<=", Right: yyDollar[3].expr}
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.ContainKeyExpr{KeyExpr: yyDollar[1].expr, MapID: yyDollar[3].token.Literal}
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.MatchExpr{Expr: yyDollar[1].expr, RegExpr: yyDollar[3].expr}
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.MatchExpr{Expr: yyDollar[1].expr, RegExpr: yyDollar[3].expr}
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: &ast.MatchExpr{Expr: yyDollar[1].expr, RegExpr: yyDollar[3].expr}}
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: &ast.MatchExpr{Expr: yyDollar[1].expr, RegExpr: yyDollar[3].expr}}
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &ast.MatchExpr{Expr: &defaultExpr, RegExpr: yyDollar[1].expr}
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "++", After: true}
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[1].expr, Operator: "--", After: true}
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			//fmt.Println("YACC: want regexp!!")
			inRegExp = true
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			//fmt.Println("FINISH")
			yyVAL.expr = &ast.StringExpr{Literal: yyDollar[3].token.Literal}
		}
	case 85:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.expr = nil
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "!", Expr: yyDollar[2].expr}
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 89:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].token.Literal, SubExprs: yyDollar[3].exprs}
		}
	case 90:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &ast.CallExpr{Name: yyDollar[1].token.Literal, SubExprs: yyDollar[3].exprs}
		}
	case 91:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &ast.AnonymousCallExpr{Expr: yyDollar[1].expr, SubExprs: yyDollar[3].exprs}
		}
	case 92:
		yyDollar = yyS[yypt-7 : yypt+1]
		{
			yyVAL.expr = &ast.FuncExpr{Args: yyDollar[3].identArgs, Stmts: yyDollar[6].stmts}
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.expr = &ast.ParentExpr{SubExpr: yyDollar[2].expr}
		}
	case 94:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[2].expr, Operator: "++"}
		}
	case 95:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.CompExpr{Left: yyDollar[2].expr, Operator: "--"}
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &ast.NumExpr{Literal: yyDollar[1].token.Literal}
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &ast.ConstExpr{Literal: yyDollar[1].token.Literal}
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &ast.ConstExpr{Literal: yyDollar[1].token.Literal}
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &ast.ConstExpr{Literal: yyDollar[1].token.Literal}
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &ast.StringExpr{Literal: yyDollar[1].token.Literal}
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 102:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "+", Expr: yyDollar[2].expr}
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.UnaryExpr{Operator: "-", Expr: yyDollar[2].expr}
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.exprs = []ast.Expr{yyDollar[1].expr}
		}
	case 105:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[4].expr)
		}
	case 106:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.expr = nil
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 109:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.expr = &ast.FieldExpr{Expr: yyDollar[2].expr}
		}
	case 110:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.expr = &ast.ItemExpr{Expr: yyDollar[1].expr, Index: yyDollar[3].exprs}
		}
	case 111:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = &ast.IdentExpr{Literal: yyDollar[1].token.Literal}
		}
	case 112:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.identArgs = []string{}
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.identArgs = []string{yyDollar[1].token.Literal}
		}
	case 114:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.identArgs = append(yyDollar[1].identArgs, yyDollar[4].token.Literal)
		}
	case 115:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.stmt = nil
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 117:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.expr = nil
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.expr = yyDollar[1].expr
		}
	}
	goto yystack /* stack new state and value */
}
