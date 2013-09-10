
//line n1ql.y:2
package goyacc
import __yyfmt__ "fmt"
//line n1ql.y:2
		import "github.com/couchbaselabs/clog"
import "github.com/couchbaselabs/tuqtng/parser"
import "github.com/couchbaselabs/tuqtng/ast"

func logDebugGrammar(format string, v ...interface{}) {
    clog.To(parser.PARSER_CHANNEL, format, v...)
}

//line n1ql.y:12
type yySymType struct {
	yys int
s string
n int
f float64}

const ALTER = 57346
const BETWEEN = 57347
const BUCKET = 57348
const CAST = 57349
const COLLATE = 57350
const DATABASE = 57351
const DELETE = 57352
const EACH = 57353
const EXCEPT = 57354
const EXISTS = 57355
const IF = 57356
const INLINE = 57357
const INSERT = 57358
const INTERSECT = 57359
const INTO = 57360
const JOIN = 57361
const PATH = 57362
const UNION = 57363
const UPDATE = 57364
const EXPLAIN = 57365
const CREATE = 57366
const DROP = 57367
const PRIMARY = 57368
const VIEW = 57369
const INDEX = 57370
const ON = 57371
const USING = 57372
const DISTINCT = 57373
const UNIQUE = 57374
const SELECT = 57375
const AS = 57376
const FROM = 57377
const WHERE = 57378
const ORDER = 57379
const BY = 57380
const ASC = 57381
const DESC = 57382
const LIMIT = 57383
const OFFSET = 57384
const GROUP = 57385
const HAVING = 57386
const LBRACE = 57387
const RBRACE = 57388
const LBRACKET = 57389
const RBRACKET = 57390
const COMMA = 57391
const COLON = 57392
const TRUE = 57393
const FALSE = 57394
const NULL = 57395
const INT = 57396
const NUMBER = 57397
const IDENTIFIER = 57398
const STRING = 57399
const PLUS = 57400
const MINUS = 57401
const MULT = 57402
const DIV = 57403
const CONCAT = 57404
const AND = 57405
const OR = 57406
const NOT = 57407
const EQ = 57408
const NE = 57409
const GT = 57410
const GTE = 57411
const LT = 57412
const LTE = 57413
const LPAREN = 57414
const RPAREN = 57415
const LIKE = 57416
const IS = 57417
const VALUED = 57418
const MISSING = 57419
const DOT = 57420
const CASE = 57421
const WHEN = 57422
const THEN = 57423
const ELSE = 57424
const END = 57425
const ANY = 57426
const ALL = 57427
const OVER = 57428
const FIRST = 57429
const ARRAY = 57430
const IN = 57431
const MOD = 57432

var yyToknames = []string{
	"ALTER",
	"BETWEEN",
	"BUCKET",
	"CAST",
	"COLLATE",
	"DATABASE",
	"DELETE",
	"EACH",
	"EXCEPT",
	"EXISTS",
	"IF",
	"INLINE",
	"INSERT",
	"INTERSECT",
	"INTO",
	"JOIN",
	"PATH",
	"UNION",
	"UPDATE",
	"EXPLAIN",
	"CREATE",
	"DROP",
	"PRIMARY",
	"VIEW",
	"INDEX",
	"ON",
	"USING",
	"DISTINCT",
	"UNIQUE",
	"SELECT",
	"AS",
	"FROM",
	"WHERE",
	"ORDER",
	"BY",
	"ASC",
	"DESC",
	"LIMIT",
	"OFFSET",
	"GROUP",
	"HAVING",
	"LBRACE",
	"RBRACE",
	"LBRACKET",
	"RBRACKET",
	"COMMA",
	"COLON",
	"TRUE",
	"FALSE",
	"NULL",
	"INT",
	"NUMBER",
	"IDENTIFIER",
	"STRING",
	"PLUS",
	"MINUS",
	"MULT",
	"DIV",
	"CONCAT",
	"AND",
	"OR",
	"NOT",
	"EQ",
	"NE",
	"GT",
	"GTE",
	"LT",
	"LTE",
	"LPAREN",
	"RPAREN",
	"LIKE",
	"IS",
	"VALUED",
	"MISSING",
	"DOT",
	"CASE",
	"WHEN",
	"THEN",
	"ELSE",
	"END",
	"ANY",
	"ALL",
	"OVER",
	"FIRST",
	"ARRAY",
	"IN",
	"MOD",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 141
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1001

var yyAct = []int{

	50, 237, 189, 83, 34, 180, 131, 136, 31, 93,
	76, 102, 103, 104, 105, 107, 254, 253, 117, 135,
	252, 251, 104, 105, 107, 204, 86, 117, 120, 81,
	84, 248, 226, 242, 45, 49, 79, 120, 125, 240,
	117, 85, 236, 106, 155, 146, 140, 88, 172, 95,
	120, 280, 106, 258, 86, 245, 119, 87, 217, 244,
	220, 126, 127, 128, 129, 84, 124, 102, 103, 104,
	105, 107, 108, 109, 117, 110, 115, 113, 114, 111,
	112, 219, 218, 116, 120, 87, 188, 118, 267, 273,
	210, 123, 274, 173, 173, 154, 238, 247, 202, 106,
	133, 153, 259, 157, 158, 159, 160, 161, 162, 163,
	164, 165, 166, 167, 168, 169, 170, 171, 175, 257,
	174, 156, 130, 152, 187, 239, 190, 256, 46, 151,
	176, 185, 51, 133, 35, 35, 231, 37, 81, 230,
	32, 178, 177, 36, 229, 79, 35, 228, 200, 203,
	211, 209, 206, 201, 173, 145, 95, 144, 207, 142,
	141, 99, 89, 82, 212, 43, 147, 143, 92, 197,
	223, 199, 196, 215, 148, 137, 205, 198, 195, 214,
	47, 48, 187, 187, 13, 121, 122, 221, 222, 185,
	185, 91, 181, 182, 149, 150, 101, 40, 41, 21,
	27, 25, 138, 232, 17, 233, 73, 283, 74, 235,
	29, 30, 68, 69, 70, 71, 72, 56, 64, 234,
	53, 186, 243, 266, 187, 98, 52, 249, 250, 246,
	241, 185, 100, 58, 179, 265, 208, 97, 96, 255,
	59, 22, 42, 23, 19, 60, 61, 132, 62, 63,
	67, 26, 261, 262, 263, 264, 2, 66, 65, 184,
	18, 190, 183, 268, 12, 10, 225, 275, 276, 3,
	12, 10, 278, 17, 279, 16, 44, 119, 57, 17,
	55, 16, 54, 90, 39, 284, 94, 277, 102, 103,
	104, 105, 107, 108, 109, 117, 110, 115, 113, 114,
	111, 112, 33, 78, 116, 120, 77, 75, 118, 28,
	271, 119, 15, 272, 213, 14, 24, 38, 20, 11,
	106, 7, 102, 103, 104, 105, 107, 108, 109, 117,
	110, 115, 113, 114, 111, 112, 9, 8, 116, 120,
	6, 5, 118, 4, 119, 1, 0, 282, 0, 0,
	0, 0, 0, 0, 106, 102, 103, 104, 105, 107,
	108, 109, 117, 110, 115, 113, 114, 111, 112, 0,
	0, 116, 120, 0, 0, 118, 0, 119, 0, 0,
	281, 0, 0, 0, 0, 0, 0, 106, 102, 103,
	104, 105, 107, 108, 109, 117, 110, 115, 113, 114,
	111, 112, 0, 0, 116, 120, 0, 0, 118, 0,
	119, 0, 0, 270, 0, 0, 0, 0, 0, 0,
	106, 102, 103, 104, 105, 107, 108, 109, 117, 110,
	115, 113, 114, 111, 112, 0, 0, 116, 120, 0,
	0, 118, 0, 119, 0, 0, 269, 0, 0, 0,
	0, 0, 0, 106, 102, 103, 104, 105, 107, 108,
	109, 117, 110, 115, 113, 114, 111, 112, 0, 0,
	116, 120, 0, 0, 118, 0, 260, 119, 0, 0,
	0, 0, 0, 0, 0, 0, 106, 0, 102, 103,
	104, 105, 107, 108, 109, 117, 110, 115, 113, 114,
	111, 112, 0, 0, 116, 120, 0, 0, 118, 0,
	0, 227, 119, 216, 0, 0, 0, 0, 0, 0,
	106, 0, 0, 102, 103, 104, 105, 107, 108, 109,
	117, 110, 115, 113, 114, 111, 112, 0, 0, 116,
	120, 0, 0, 118, 0, 119, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 106, 102, 103, 104, 105,
	107, 108, 109, 117, 110, 115, 113, 114, 111, 112,
	0, 0, 116, 120, 0, 0, 118, 0, 119, 0,
	0, 0, 0, 0, 194, 0, 0, 0, 106, 102,
	103, 104, 105, 107, 108, 109, 117, 110, 115, 113,
	114, 111, 112, 0, 0, 116, 120, 0, 0, 118,
	0, 119, 0, 0, 0, 0, 0, 193, 0, 0,
	0, 106, 102, 103, 104, 105, 107, 108, 109, 117,
	110, 115, 113, 114, 111, 112, 0, 0, 116, 120,
	0, 0, 118, 0, 119, 0, 0, 0, 0, 0,
	192, 0, 0, 0, 106, 102, 103, 104, 105, 107,
	108, 109, 117, 110, 115, 113, 114, 111, 112, 0,
	0, 116, 120, 0, 0, 118, 0, 119, 0, 0,
	0, 0, 0, 191, 0, 0, 0, 106, 102, 103,
	104, 105, 107, 108, 109, 117, 110, 115, 113, 114,
	111, 112, 0, 0, 116, 120, 0, 0, 118, 0,
	119, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	106, 102, 103, 104, 105, 107, 108, 109, 117, 110,
	115, 113, 114, 111, 112, 0, 0, 116, 120, 0,
	0, 224, 0, 119, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 106, 102, 103, 104, 105, 107, 108,
	109, 117, 110, 115, 113, 114, 111, 112, 0, 119,
	116, 120, 0, 0, 139, 0, 0, 0, 0, 0,
	102, 103, 104, 105, 107, 108, 106, 117, 110, 115,
	113, 114, 111, 112, 0, 119, 116, 120, 0, 0,
	118, 0, 0, 0, 0, 0, 102, 103, 104, 105,
	107, 0, 106, 117, 110, 115, 113, 114, 111, 112,
	0, 0, 116, 120, 0, 73, 118, 74, 0, 0,
	0, 68, 69, 70, 71, 72, 56, 64, 106, 53,
	186, 0, 0, 0, 0, 52, 0, 0, 0, 0,
	0, 0, 58, 0, 0, 0, 0, 0, 0, 59,
	0, 0, 0, 0, 60, 61, 0, 62, 63, 73,
	0, 74, 0, 0, 0, 68, 69, 70, 71, 72,
	56, 64, 0, 53, 80, 0, 0, 0, 0, 52,
	0, 0, 0, 0, 0, 0, 58, 0, 0, 0,
	0, 0, 0, 59, 0, 0, 0, 0, 60, 61,
	0, 62, 63, 73, 0, 74, 134, 0, 0, 68,
	69, 70, 71, 72, 56, 64, 0, 53, 0, 0,
	0, 0, 0, 52, 0, 0, 0, 0, 0, 0,
	58, 0, 0, 0, 0, 0, 0, 59, 0, 0,
	0, 0, 60, 61, 0, 62, 63, 73, 0, 74,
	0, 0, 0, 68, 69, 70, 71, 72, 56, 64,
	0, 53, 0, 0, 0, 0, 0, 52, 0, 0,
	0, 0, 0, 0, 58, 0, 0, 0, 0, 0,
	0, 59, 0, 0, 0, 0, 60, 61, 0, 62,
	63,
}
var yyPact = []int{

	246, -1000, -1000, 240, -1000, -1000, -1000, -1000, -1000, -1000,
	216, 162, 215, 166, 164, 179, 90, -1000, -1000, 87,
	156, 160, 214, 109, 164, 78, 138, 912, 824, -1000,
	-1000, -1000, 107, -56, 7, -1000, -31, 106, -1000, 149,
	114, 912, 209, 208, 138, -1000, 105, 171, 158, -1000,
	630, -1000, 912, 912, -1000, -1000, 19, -1000, 912, -42,
	912, 912, 912, 912, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 76, 868, -1000, -1000, 126, -1000, 168,
	-1000, 696, -32, -1000, 104, 103, 113, 101, 99, -33,
	-1000, 112, -1000, -1000, 125, 155, 73, 45, -1000, -34,
	-1000, 912, 912, 912, 912, 912, 912, 912, 912, 912,
	912, 912, 912, 912, 912, 912, 912, -26, 98, 912,
	65, -1000, -1000, 161, 13, 912, 597, 564, 531, 498,
	-1000, 132, 123, 119, -1000, 129, 122, 824, 97, 38,
	79, -64, -1000, 128, -1000, -1000, 96, -1000, 912, -1000,
	-1000, 206, 95, 18, 94, 79, 135, -38, -38, -25,
	-25, -25, -25, 748, 722, -47, -47, -47, -47, -47,
	-47, -47, 912, -1000, 465, -1000, 5, -1000, -1000, -1000,
	-13, 780, 780, 121, -1000, -1000, -1000, 663, -1000, -50,
	430, 91, 88, 83, 80, -1000, 43, 912, -1000, 912,
	-1000, -1000, -1000, -1000, 79, -1000, -36, -1000, 69, -39,
	912, -45, -1000, -1000, 912, -47, -1000, -1000, -1000, -1000,
	-1000, -14, -18, 780, 37, -52, 912, 912, -68, -69,
	-72, -73, -1000, -1000, -1000, -21, 71, -1000, -1000, -1000,
	63, -20, 46, -1000, -1000, -1000, -1000, -1000, -1000, 630,
	396, 912, 912, 912, 912, -1000, -1000, 205, 193, 16,
	912, 363, 330, 230, 9, 69, 69, 912, -1000, -1000,
	-1000, 912, -1000, 912, -1000, -1000, -1000, -22, 297, 264,
	177, -1000, -1000, 69, -1000,
}
var yyPgo = []int{

	0, 345, 256, 343, 341, 340, 337, 336, 1, 19,
	321, 319, 318, 317, 184, 316, 251, 180, 315, 314,
	7, 312, 309, 307, 10, 306, 303, 0, 8, 302,
	3, 4, 9, 286, 284, 283, 132, 282, 280, 278,
	2, 266, 5, 262, 259, 258, 257, 250, 6, 247,
}
var yyR1 = []int{

	0, 1, 1, 2, 2, 2, 4, 4, 6, 6,
	6, 6, 7, 7, 7, 7, 8, 8, 5, 5,
	3, 10, 11, 11, 17, 17, 19, 19, 14, 21,
	22, 22, 22, 23, 24, 24, 25, 25, 25, 26,
	26, 15, 15, 15, 18, 18, 28, 28, 30, 30,
	29, 29, 16, 16, 12, 12, 32, 32, 33, 33,
	33, 13, 13, 13, 34, 35, 20, 27, 27, 27,
	27, 27, 27, 27, 27, 27, 27, 27, 27, 27,
	27, 27, 27, 27, 27, 27, 27, 27, 27, 27,
	27, 27, 36, 36, 36, 37, 38, 38, 38, 38,
	38, 38, 38, 38, 38, 38, 38, 38, 38, 38,
	40, 40, 41, 41, 31, 31, 31, 42, 42, 43,
	43, 44, 44, 39, 39, 39, 39, 39, 39, 39,
	45, 45, 46, 46, 48, 48, 49, 47, 47, 9,
	9,
}
var yyR2 = []int{

	0, 1, 2, 1, 1, 1, 1, 1, 5, 8,
	7, 10, 8, 11, 10, 13, 1, 1, 5, 8,
	1, 3, 4, 4, 0, 4, 0, 2, 3, 1,
	0, 1, 1, 1, 1, 3, 1, 1, 3, 1,
	3, 0, 2, 5, 2, 5, 1, 2, 4, 5,
	1, 3, 0, 2, 0, 3, 1, 3, 1, 2,
	2, 0, 1, 2, 2, 2, 1, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 1, 2, 2, 1, 1, 1, 1, 3, 5,
	7, 7, 9, 7, 9, 7, 3, 4, 5, 5,
	3, 5, 0, 2, 1, 4, 3, 1, 3, 1,
	1, 1, 3, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 2, 3, 1, 3, 3, 2, 3, 1,
	3,
}
var yyChk = []int{

	-1000, -1, -2, 23, -3, -4, -5, -10, -6, -7,
	25, -11, 24, -14, -18, -21, 35, 33, -2, 28,
	-12, 37, 26, 28, -15, 35, -16, 36, -22, 31,
	32, -28, 50, -29, -31, 56, 56, 50, -13, -34,
	41, 38, 28, 56, -16, -28, 50, -17, 43, -20,
	-27, -36, 65, 59, -37, -38, 56, -39, 72, 79,
	84, 85, 87, 88, 57, -45, -46, -47, 51, 52,
	53, 54, 55, 45, 47, -23, -24, -25, -26, -20,
	60, -27, 56, -30, 86, 34, 47, 78, 78, 56,
	-35, 42, 54, -32, -33, -20, 29, 29, -17, 56,
	-14, 38, 58, 59, 60, 61, 90, 62, 63, 64,
	66, 70, 71, 68, 69, 67, 74, 65, 78, 47,
	75, -36, -36, 72, -20, 80, -27, -27, -27, -27,
	46, -48, -49, 57, 48, -9, -20, 49, 34, 78,
	78, 56, 56, 54, 56, 56, 78, 54, 49, 39,
	40, 56, 50, 56, 50, 78, -9, -27, -27, -27,
	-27, -27, -27, -27, -27, -27, -27, -27, -27, -27,
	-27, -27, 74, 56, -27, 53, 65, 77, 76, 73,
	-42, 31, 32, -43, -44, -20, 60, -27, 73, -40,
	-27, 86, 86, 86, 86, 46, 49, 50, 48, 49,
	-24, 56, 60, -28, 89, 48, 56, -32, 30, 56,
	72, 56, -28, -19, 44, -27, 48, 53, 77, 76,
	73, -42, -42, 49, 78, -41, 82, 81, 56, 56,
	56, 56, -48, -20, -9, -31, 78, -8, 27, 56,
	78, -9, 78, -20, 73, 73, -42, 60, 83, -27,
	-27, 89, 89, 89, 89, -30, 56, 56, 73, 56,
	80, -27, -27, -27, -27, 30, 30, 72, -40, 83,
	83, 80, 83, 80, 83, -8, -8, -9, -27, -27,
	73, 83, 83, 30, -8,
}
var yyDef = []int{

	0, -2, 1, 0, 3, 4, 5, 20, 6, 7,
	0, 54, 0, 41, 52, 30, 0, 29, 2, 0,
	61, 0, 0, 0, 52, 0, 24, 0, 0, 31,
	32, 44, 0, 46, 50, 114, 0, 0, 21, 62,
	0, 0, 0, 0, 24, 42, 0, 0, 0, 53,
	66, 91, 0, 0, 94, 95, 96, 97, 0, 0,
	0, 0, 0, 0, 123, 124, 125, 126, 127, 128,
	129, 130, 131, 0, 0, 28, 33, 34, 36, 37,
	39, 66, 0, 47, 0, 0, 0, 0, 0, 0,
	63, 0, 64, 55, 56, 58, 0, 0, 22, 0,
	23, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 92, 93, 0, 0, 0, 0, 0, 0, 0,
	132, 0, 134, 0, 137, 0, 139, 0, 0, 0,
	0, 0, 51, 0, 116, 18, 0, 65, 0, 59,
	60, 8, 0, 0, 0, 0, 26, 67, 68, 69,
	70, 71, 72, 73, 74, 75, 76, 77, 78, 79,
	80, 81, 0, 83, 0, 85, 0, 87, 89, 106,
	0, 0, 0, 117, 119, 120, 121, 66, 98, 112,
	0, 0, 0, 0, 0, 133, 0, 0, 138, 0,
	35, 38, 40, 45, 0, 115, 0, 57, 0, 0,
	0, 0, 43, 25, 0, 82, 84, 86, 88, 90,
	107, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 135, 136, 140, 48, 0, 10, 16, 17,
	0, 0, 0, 27, 108, 109, 118, 122, 99, 113,
	110, 0, 0, 0, 0, 49, 19, 9, 12, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 111, 100,
	101, 0, 103, 0, 105, 11, 14, 0, 0, 0,
	13, 102, 104, 0, 15,
}
var yyTok1 = []int{

	1,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73, 74, 75, 76, 77, 78, 79, 80, 81,
	82, 83, 84, 85, 86, 87, 88, 89, 90,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
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

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %U %s\n", uint(char), yyTokname(c))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
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
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
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
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
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
			if yyn < 0 || yyn == yychar {
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
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf("saw %s\n", yyTokname(yychar))
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
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
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
		//line n1ql.y:53
		{
		logDebugGrammar("INPUT")
	}
	case 2:
		//line n1ql.y:57
		{
		logDebugGrammar("INPUT - EXPLAIN")
		parsingStatement.SetExplainOnly(true)
	}
	case 3:
		//line n1ql.y:63
		{
		logDebugGrammar("STMT - SELECT")
	}
	case 4:
		//line n1ql.y:67
		{
	}
	case 5:
		//line n1ql.y:70
		{
		logDebugGrammar("STMT - DROP INDEX")
	}
	case 6:
		//line n1ql.y:77
		{
		logDebugGrammar("STMT - CREATE PRIMARY INDEX")
	}
	case 7:
		//line n1ql.y:81
		{
		logDebugGrammar("STMT - CREATE SECONDARY INDEX")
	}
	case 8:
		//line n1ql.y:87
		{
		bucket := yyS[yypt-0].s
		createIndexStmt := ast.NewCreateIndexStatement()
		createIndexStmt.Bucket = bucket
		createIndexStmt.Primary = true
		parsingStatement = createIndexStmt
	}
	case 9:
		//line n1ql.y:95
		{
		pool := yyS[yypt-2].s
		bucket := yyS[yypt-0].s
		createIndexStmt := ast.NewCreateIndexStatement()
		createIndexStmt.Pool = pool
		createIndexStmt.Bucket = bucket
		createIndexStmt.Primary = true
		parsingStatement = createIndexStmt
	}
	case 10:
		//line n1ql.y:105
		{
		method := parsingStack.Pop().(string)
		bucket := yyS[yypt-2].s
		createIndexStmt := ast.NewCreateIndexStatement()
		createIndexStmt.Bucket = bucket
		createIndexStmt.Method = method
		createIndexStmt.Primary = true
		parsingStatement = createIndexStmt
	}
	case 11:
		//line n1ql.y:115
		{
		method := parsingStack.Pop().(string)
		bucket := yyS[yypt-2].s
		pool := yyS[yypt-4].s
		createIndexStmt := ast.NewCreateIndexStatement()
		createIndexStmt.Pool = pool
		createIndexStmt.Bucket = bucket
		createIndexStmt.Method = method
		createIndexStmt.Primary = true
		parsingStatement = createIndexStmt
	}
	case 12:
		//line n1ql.y:129
		{
		on := parsingStack.Pop().(ast.ExpressionList)
		bucket := yyS[yypt-3].s
		name := yyS[yypt-5].s
		createIndexStmt := ast.NewCreateIndexStatement()
		createIndexStmt.On = on
		createIndexStmt.Bucket = bucket
		createIndexStmt.Name = name
		createIndexStmt.Primary = false
		parsingStatement = createIndexStmt
	}
	case 13:
		//line n1ql.y:141
		{
		on := parsingStack.Pop().(ast.ExpressionList)
		bucket := yyS[yypt-3].s
		pool := yyS[yypt-5].s
		name := yyS[yypt-8].s
		createIndexStmt := ast.NewCreateIndexStatement()
		createIndexStmt.On = on
		createIndexStmt.Pool = pool
		createIndexStmt.Bucket = bucket
		createIndexStmt.Name = name
		createIndexStmt.Primary = false
		parsingStatement = createIndexStmt
	}
	case 14:
		//line n1ql.y:155
		{
		method := parsingStack.Pop().(string)
		on := parsingStack.Pop().(ast.ExpressionList)
		bucket := yyS[yypt-5].s
		name := yyS[yypt-7].s
		createIndexStmt := ast.NewCreateIndexStatement()
		createIndexStmt.On = on
		createIndexStmt.Bucket = bucket
		createIndexStmt.Name = name
		createIndexStmt.Method = method
		createIndexStmt.Primary = false
		parsingStatement = createIndexStmt
	}
	case 15:
		//line n1ql.y:169
		{
		method := parsingStack.Pop().(string)
		on := parsingStack.Pop().(ast.ExpressionList)
		bucket := yyS[yypt-5].s
		pool := yyS[yypt-7].s
		name := yyS[yypt-10].s
		createIndexStmt := ast.NewCreateIndexStatement()
		createIndexStmt.On = on
		createIndexStmt.Pool = pool
		createIndexStmt.Bucket = bucket
		createIndexStmt.Name = name
		createIndexStmt.Method = method
		createIndexStmt.Primary = false
		parsingStatement = createIndexStmt
	}
	case 16:
		//line n1ql.y:187
		{
		parsingStack.Push("view")
	}
	case 17:
		//line n1ql.y:191
		{
		parsingStack.Push(yyS[yypt-0].s)
	}
	case 18:
		//line n1ql.y:197
		{
		bucket := yyS[yypt-2].s
		name := yyS[yypt-0].s
		dropIndexStmt := ast.NewDropIndexStatement()
		dropIndexStmt.Bucket = bucket
		dropIndexStmt.Name = name
		parsingStatement = dropIndexStmt
	}
	case 19:
		//line n1ql.y:206
		{
		bucket := yyS[yypt-2].s
		pool := yyS[yypt-4].s
		name := yyS[yypt-0].s
		dropIndexStmt := ast.NewDropIndexStatement()
		dropIndexStmt.Pool = pool
		dropIndexStmt.Bucket = bucket
		dropIndexStmt.Name = name
		parsingStatement = dropIndexStmt
	}
	case 20:
		//line n1ql.y:220
		{
		logDebugGrammar("SELECT_STMT")
	}
	case 21:
		//line n1ql.y:226
		{
		// future extensibility for comining queries with UNION, etc
	logDebugGrammar("SELECT_COMPOUND")
	}
	case 22:
		//line n1ql.y:233
		{
		logDebugGrammar("SELECT_CORE")
	}
	case 23:
		//line n1ql.y:237
		{
		logDebugGrammar("SELECT_CORE")
	}
	case 24:
		//line n1ql.y:243
		{
	}
	case 25:
		//line n1ql.y:246
		{
		group_by := parsingStack.Pop().(ast.ExpressionList)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.GroupBy = group_by
		default:
			logDebugGrammar("This statement does not support GROUP BY")
		}
	}
	case 26:
		//line n1ql.y:258
		{
	}
	case 27:
		//line n1ql.y:261
		{
		logDebugGrammar("SELECT HAVING - EXPR")
		having_part := parsingStack.Pop().(ast.Expression)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Having = having_part
		default:
			logDebugGrammar("This statement does not support HAVING")
		}
	}
	case 28:
		//line n1ql.y:274
		{
		logDebugGrammar("SELECT_SELECT")
	}
	case 29:
		//line n1ql.y:280
		{
		logDebugGrammar("SELECT_SELECT_HEAD")
	}
	case 30:
		//line n1ql.y:286
		{
	}
	case 31:
		//line n1ql.y:289
		{
		logDebugGrammar("SELECT_SELECT_QUALIFIER DISTINCT")
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Distinct = true
		default:
			logDebugGrammar("This statement does not support WHERE")
		}
	}
	case 32:
		//line n1ql.y:299
		{
		logDebugGrammar("SELECT_SELECT_QUALIFIER UNIQUE")
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Distinct = true
		default:
			logDebugGrammar("This statement does not support WHERE")
		}
	}
	case 33:
		//line n1ql.y:311
		{
		logDebugGrammar("SELECT SELECT TAIL - EXPR")
		result_expr_list := parsingStack.Pop().(ast.ResultExpressionList)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Select = result_expr_list
		default:
			logDebugGrammar("This statement does not support WHERE")
		}
	
	}
	case 34:
		//line n1ql.y:325
		{
		result_expr := parsingStack.Pop().(*ast.ResultExpression)
		parsingStack.Push(ast.ResultExpressionList{result_expr})
	}
	case 35:
		//line n1ql.y:330
		{
		result_expr_list := parsingStack.Pop().(ast.ResultExpressionList)
		result_expr := parsingStack.Pop().(*ast.ResultExpression)
		// list items pushed onto the stack end up in reverse order
	// this prepends items in the list to restore order
	new_list := ast.ResultExpressionList{result_expr}
		for _, v := range result_expr_list {
			new_list = append(new_list, v)
		}
		parsingStack.Push(new_list)
	}
	case 36:
		//line n1ql.y:343
		{
		logDebugGrammar("RESULT STAR")
	}
	case 37:
		//line n1ql.y:347
		{
		logDebugGrammar("RESULT EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpression(expr_part)
		parsingStack.Push(result_expr)
	}
	case 38:
		//line n1ql.y:354
		{
		logDebugGrammar("SORT EXPR ASC")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewResultExpressionWithAlias(expr_part, yyS[yypt-0].s)
		parsingStack.Push(result_expr)
	}
	case 39:
		//line n1ql.y:363
		{
		logDebugGrammar("STAR")
		result_expr := ast.NewStarResultExpression()
		parsingStack.Push(result_expr)
	}
	case 40:
		//line n1ql.y:369
		{
		logDebugGrammar("PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		result_expr := ast.NewDotStarResultExpression(expr_part)
		parsingStack.Push(result_expr)
	}
	case 41:
		//line n1ql.y:378
		{
		logDebugGrammar("SELECT FROM - EMPTY")
	}
	case 42:
		//line n1ql.y:382
		{
		logDebugGrammar("SELECT FROM - DATASOURCE")
		from := parsingStack.Pop().(*ast.From)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.From = from
		default:
			logDebugGrammar("This statement does not support FROM")
		}
	}
	case 43:
		//line n1ql.y:393
		{
		logDebugGrammar("SELECT FROM - DATASOURCE WITH POOL")
		from := parsingStack.Pop().(*ast.From)
		from.Pool = yyS[yypt-2].s
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.From = from
		default:
			logDebugGrammar("This statement does not support FROM")
		}
	}
	case 44:
		//line n1ql.y:407
		{
		logDebugGrammar("SELECT FROM - DATASOURCE")
		from := parsingStack.Pop().(*ast.From)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.From = from
		default:
			logDebugGrammar("This statement does not support FROM")
		}
	}
	case 45:
		//line n1ql.y:418
		{
		logDebugGrammar("SELECT FROM - DATASOURCE WITH POOL")
		from := parsingStack.Pop().(*ast.From)
		from.Pool = yyS[yypt-2].s
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.From = from
		default:
			logDebugGrammar("This statement does not support FROM")
		}
	}
	case 46:
		//line n1ql.y:432
		{
		logDebugGrammar("FROM DATASOURCE WITHOUT OVER")
	}
	case 47:
		//line n1ql.y:436
		{
		logDebugGrammar("FROM DATASOURCE WITH OVER")
		rest := parsingStack.Pop().(*ast.From)
		last := parsingStack.Pop().(*ast.From)
		last.Over = rest
		parsingStack.Push(last)
	}
	case 48:
		//line n1ql.y:446
		{
		logDebugGrammar("OVER IN")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-2].s})
	}
	case 49:
		//line n1ql.y:452
		{
		logDebugGrammar("OVER IN nested")
		rest := parsingStack.Pop().(*ast.From)
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-3].s, Over:rest})
	}
	case 50:
		//line n1ql.y:461
		{
		logDebugGrammar("FROM DATASOURCE")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj})
	}
	case 51:
		//line n1ql.y:467
		{
	    // fixme support over as
	logDebugGrammar("FROM DATASOURCE AS")
		proj := parsingStack.Pop().(ast.Expression)
		parsingStack.Push(&ast.From{Projection: proj, As: yyS[yypt-0].s})
	}
	case 52:
		//line n1ql.y:475
		{
		logDebugGrammar("SELECT WHERE - EMPTY")
	}
	case 53:
		//line n1ql.y:479
		{
		logDebugGrammar("SELECT WHERE - EXPR")
		where_part := parsingStack.Pop().(ast.Expression)
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Where = where_part
		default:
			logDebugGrammar("This statement does not support WHERE")
		}
	}
	case 55:
		//line n1ql.y:493
		{
	
	}
	case 56:
		//line n1ql.y:499
		{
	
	}
	case 57:
		//line n1ql.y:503
		{
	
	}
	case 58:
		//line n1ql.y:508
		{
		logDebugGrammar("SORT EXPR")
		expr := parsingStack.Pop()
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.OrderBy = append(parsingStatement.OrderBy, ast.NewSortExpression(expr.(ast.Expression), true))
		default:
			logDebugGrammar("This statement does not support ORDER BY")
		}
	}
	case 59:
		//line n1ql.y:519
		{
		logDebugGrammar("SORT EXPR ASC")
		expr := parsingStack.Pop()
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.OrderBy = append(parsingStatement.OrderBy, ast.NewSortExpression(expr.(ast.Expression), true))
		default:
			logDebugGrammar("This statement does not support ORDER BY")
		}
	}
	case 60:
		//line n1ql.y:530
		{
		logDebugGrammar("SORT EXPR DESC")
		expr := parsingStack.Pop()
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.OrderBy = append(parsingStatement.OrderBy, ast.NewSortExpression(expr.(ast.Expression), false))
		default:
			logDebugGrammar("This statement does not support ORDER BY")
		}
	}
	case 61:
		//line n1ql.y:542
		{
	
	}
	case 62:
		//line n1ql.y:546
		{
	
	}
	case 63:
		//line n1ql.y:550
		{
	
	}
	case 64:
		//line n1ql.y:556
		{
		logDebugGrammar("LIMIT %d", yyS[yypt-0].n)
		if yyS[yypt-0].n < 0 {
			panic("LIMIT cannot be negative")
		}
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Limit = yyS[yypt-0].n
		default:
			logDebugGrammar("This statement does not support LIMIT")
		}
	}
	case 65:
		//line n1ql.y:570
		{
		logDebugGrammar("OFFSET %d", yyS[yypt-0].n)
		if yyS[yypt-0].n < 0 {
			panic("OFFSET cannot be negative")
		}
		switch parsingStatement := parsingStatement.(type) {
		case *ast.SelectStatement:
			parsingStatement.Offset = yyS[yypt-0].n
		default:
			logDebugGrammar("This statement does not support OFFSET")
		}
	}
	case 66:
		//line n1ql.y:586
		{
		logDebugGrammar("EXPRESSION")
	}
	case 67:
		//line n1ql.y:591
		{
		logDebugGrammar("EXPR - PLUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewPlusOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 68:
		//line n1ql.y:599
		{
		logDebugGrammar("EXPR - MINUS")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewSubtractOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 69:
		//line n1ql.y:607
		{
		logDebugGrammar("EXPR - MULT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewMultiplyOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 70:
		//line n1ql.y:615
		{
		logDebugGrammar("EXPR - DIV")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewDivideOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 71:
		//line n1ql.y:623
		{
		logDebugGrammar("EXPR - MOD")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewModuloOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 72:
		//line n1ql.y:631
		{
		logDebugGrammar("EXPR - CONCAT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewStringConcatenateOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 73:
		//line n1ql.y:639
		{
		logDebugGrammar("EXPR - AND")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewAndOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
		parsingStack.Push(thisExpression)
	}
	case 74:
		//line n1ql.y:647
		{
		logDebugGrammar("EXPR - OR")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewOrOperator(ast.ExpressionList{left.(ast.Expression), right.(ast.Expression)})
		parsingStack.Push(thisExpression)
	}
	case 75:
		//line n1ql.y:655
		{
		logDebugGrammar("EXPR - EQ")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewEqualToOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 76:
		//line n1ql.y:663
		{
		logDebugGrammar("EXPR - LT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 77:
		//line n1ql.y:671
		{
		logDebugGrammar("EXPR - LTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLessThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 78:
		//line n1ql.y:679
		{
		logDebugGrammar("EXPR - GT")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 79:
		//line n1ql.y:687
		{
		logDebugGrammar("EXPR - GTE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewGreaterThanOrEqualOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 80:
		//line n1ql.y:695
		{
		logDebugGrammar("EXPR - NE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotEqualToOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 81:
		//line n1ql.y:703
		{
		logDebugGrammar("EXPR - LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewLikeOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 82:
		//line n1ql.y:711
		{
		logDebugGrammar("EXPR - NOT LIKE")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewNotLikeOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 83:
		//line n1ql.y:719
		{
		logDebugGrammar("EXPR DOT MEMBER")
		right := ast.NewProperty(yyS[yypt-0].s)
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right)
		parsingStack.Push(thisExpression)
	}
	case 84:
		//line n1ql.y:727
		{
		logDebugGrammar("EXPR BRACKET MEMBER")
		right := parsingStack.Pop()
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), right.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 85:
		//line n1ql.y:735
		{
		logDebugGrammar("SUFFIX_EXPR IS NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNullOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 86:
		//line n1ql.y:742
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT NULL")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotNullOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 87:
		//line n1ql.y:749
		{
		logDebugGrammar("SUFFIX_EXPR IS MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsMissingOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 88:
		//line n1ql.y:756
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT MISSING")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotMissingOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 89:
		//line n1ql.y:763
		{
		logDebugGrammar("SUFFIX_EXPR IS VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsValuedOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 90:
		//line n1ql.y:770
		{
		logDebugGrammar("SUFFIX_EXPR IS NOT VALUED")
		operand := parsingStack.Pop()
		thisExpression := ast.NewIsNotValuedOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 91:
		//line n1ql.y:777
		{
	
	}
	case 92:
		//line n1ql.y:783
		{
		logDebugGrammar("EXPR - NOT")
		operand := parsingStack.Pop()
		thisExpression := ast.NewNotOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 93:
		//line n1ql.y:790
		{
		logDebugGrammar("EXPR - CHANGE SIGN")
		operand := parsingStack.Pop()
		thisExpression := ast.NewChangeSignOperator(operand.(ast.Expression))
		parsingStack.Push(thisExpression)
	}
	case 94:
		//line n1ql.y:797
		{
	
	}
	case 95:
		//line n1ql.y:802
		{
		logDebugGrammar("SUFFIX_EXPR")
	}
	case 96:
		//line n1ql.y:808
		{
		logDebugGrammar("IDENTIFIER - %s", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s)
		parsingStack.Push(thisExpression)
	}
	case 97:
		//line n1ql.y:814
		{
		logDebugGrammar("LITERAL")
	}
	case 98:
		//line n1ql.y:818
		{
		logDebugGrammar("NESTED EXPR")
	}
	case 99:
		//line n1ql.y:822
		{
		logDebugGrammar("CASE WHEN THEN ELSE END")
		cwtee := ast.NewCaseOperator()
		topStack := parsingStack.Pop()
		switch topStack := topStack.(type) {
		case ast.Expression:
			cwtee.Else = topStack
			// now look for whenthens
		nextStack := parsingStack.Pop().([]*ast.WhenThen)
			cwtee.WhenThens = nextStack
		case []*ast.WhenThen:
			// no else
		cwtee.WhenThens = topStack
		}
		parsingStack.Push(cwtee)
	}
	case 100:
		//line n1ql.y:839
		{
		logDebugGrammar("ANY OVER")
		sub := parsingStack.Pop().(ast.Expression)
		condition := parsingStack.Pop().(ast.Expression)
		collectionAny := ast.NewCollectionAnyOperator(condition, sub, yyS[yypt-3].s)
		parsingStack.Push(collectionAny)
	}
	case 101:
		//line n1ql.y:847
		{
		logDebugGrammar("ALL OVER")
		sub := parsingStack.Pop().(ast.Expression)
		condition := parsingStack.Pop().(ast.Expression)
		collectionAny := ast.NewCollectionAllOperator(condition, sub, yyS[yypt-3].s)
		parsingStack.Push(collectionAny)
	}
	case 102:
		//line n1ql.y:855
		{
		logDebugGrammar("FIRST OVER")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(condition, sub, yyS[yypt-5].s, output)
		parsingStack.Push(collectionFirst)
	}
	case 103:
		//line n1ql.y:864
		{
		logDebugGrammar("FIRST OVER")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionFirst := ast.NewCollectionFirstOperator(nil, sub, yyS[yypt-3].s, output)
		parsingStack.Push(collectionFirst)
	}
	case 104:
		//line n1ql.y:872
		{
		logDebugGrammar("ARRAY OVER WHEN")
		condition := parsingStack.Pop().(ast.Expression)
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(condition, sub, yyS[yypt-5].s, output)
		parsingStack.Push(collectionArray)
	}
	case 105:
		//line n1ql.y:881
		{
		logDebugGrammar("ARRAY OVER")
		sub := parsingStack.Pop().(ast.Expression)
		output := parsingStack.Pop().(ast.Expression)
		collectionArray := ast.NewCollectionArrayOperator(nil, sub, yyS[yypt-3].s, output)
		parsingStack.Push(collectionArray)
	}
	case 106:
		//line n1ql.y:889
		{
		logDebugGrammar("FUNCTION EXPR NOPARAM")
		thisExpression := ast.NewFunctionCall(yyS[yypt-2].s, ast.FunctionArgExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 107:
		//line n1ql.y:895
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-3].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 108:
		//line n1ql.y:902
		{
		logDebugGrammar("FUNCTION DISTINCT EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		function := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		function.SetDistinct(true)
		parsingStack.Push(function)
	}
	case 109:
		//line n1ql.y:910
		{
		logDebugGrammar("FUNCTION EXPR PARAM")
		funarg_exp_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		thisExpression := ast.NewFunctionCall(yyS[yypt-4].s, funarg_exp_list)
		parsingStack.Push(thisExpression)
	}
	case 110:
		//line n1ql.y:919
		{
		logDebugGrammar("THEN_LIST - SINGLE")
		when_then_list := make([]*ast.WhenThen, 0)
		when_then := ast.WhenThen{Then: parsingStack.Pop().(ast.Expression), When: parsingStack.Pop().(ast.Expression)}
		when_then_list = append(when_then_list, &when_then)
		parsingStack.Push(when_then_list)
	}
	case 111:
		//line n1ql.y:927
		{
		logDebugGrammar("THEN_LIST - COMPOUND")
		rest := parsingStack.Pop().([]*ast.WhenThen)
		last := ast.WhenThen{Then: parsingStack.Pop().(ast.Expression), When: parsingStack.Pop().(ast.Expression)}
		new_list := make([]*ast.WhenThen, 0, len(rest) + 1)
		new_list = append(new_list, &last)
		for _, v := range rest {
			new_list = append(new_list, v)
		}
		parsingStack.Push(new_list)
	}
	case 112:
		//line n1ql.y:941
		{
		logDebugGrammar("ELSE - EMPTY")
	}
	case 113:
		//line n1ql.y:945
		{
		logDebugGrammar("ELSE - EXPR")
	}
	case 114:
		//line n1ql.y:951
		{
		logDebugGrammar("PATH - %v", yyS[yypt-0].s)
		thisExpression := ast.NewProperty(yyS[yypt-0].s)
		parsingStack.Push(thisExpression)
	}
	case 115:
		//line n1ql.y:957
		{
		logDebugGrammar("PATH BRACKET - %v[%v]", yyS[yypt-3].s, yyS[yypt-1].n)
		left := parsingStack.Pop()
		thisExpression := ast.NewBracketMemberOperator(left.(ast.Expression), ast.NewLiteralNumber(float64(yyS[yypt-1].n)))
		parsingStack.Push(thisExpression)
	}
	case 116:
		//line n1ql.y:964
		{
		logDebugGrammar("PATH DOT PATH - $1.s")
		right := ast.NewProperty(yyS[yypt-0].s)
		left := parsingStack.Pop()
		thisExpression := ast.NewDotMemberOperator(left.(ast.Expression), right)
		parsingStack.Push(thisExpression)
	}
	case 117:
		//line n1ql.y:975
		{
		funarg_expr := parsingStack.Pop().(*ast.FunctionArgExpression)
		parsingStack.Push(ast.FunctionArgExpressionList{funarg_expr})
	}
	case 118:
		//line n1ql.y:980
		{
		funarg_expr_list := parsingStack.Pop().(ast.FunctionArgExpressionList)
		funarg_expr := parsingStack.Pop().(*ast.FunctionArgExpression)
		// list items pushed onto the stack end up in reverse order
	// this prepends items in the list to restore order
	new_list := ast.FunctionArgExpressionList{funarg_expr}
		for _, v := range funarg_expr_list {
			new_list = append(new_list, v)
		}
		parsingStack.Push(new_list)
	}
	case 119:
		//line n1ql.y:994
		{
		logDebugGrammar("FUNARG STAR")
	}
	case 120:
		//line n1ql.y:998
		{
		logDebugGrammar("FUNARG EXPR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 121:
		//line n1ql.y:1007
		{
		logDebugGrammar("FUNSTAR")
		funarg_expr := ast.NewStarFunctionArgExpression()
		parsingStack.Push(funarg_expr)
	}
	case 122:
		//line n1ql.y:1013
		{
		logDebugGrammar("FUN PATH DOT STAR")
		expr_part := parsingStack.Pop().(ast.Expression)
		funarg_expr := ast.NewDotStarFunctionArgExpression(expr_part)
		parsingStack.Push(funarg_expr)
	}
	case 123:
		//line n1ql.y:1023
		{
		logDebugGrammar("STRING %s", yyS[yypt-0].s)
		thisExpression := ast.NewLiteralString(yyS[yypt-0].s)
		parsingStack.Push(thisExpression)
	}
	case 124:
		//line n1ql.y:1029
		{
		logDebugGrammar("NUMBER")
	}
	case 125:
		//line n1ql.y:1033
		{
		logDebugGrammar("OBJECT")
	}
	case 126:
		//line n1ql.y:1037
		{
		logDebugGrammar("ARRAY")
	}
	case 127:
		//line n1ql.y:1041
		{
		logDebugGrammar("TRUE")
		thisExpression := ast.NewLiteralBool(true)
		parsingStack.Push(thisExpression)
	}
	case 128:
		//line n1ql.y:1047
		{
		logDebugGrammar("FALSE")
		thisExpression := ast.NewLiteralBool(false)
		parsingStack.Push(thisExpression)
	}
	case 129:
		//line n1ql.y:1053
		{
		logDebugGrammar("NULL")
		thisExpression := ast.NewLiteralNull()
		parsingStack.Push(thisExpression)
	}
	case 130:
		//line n1ql.y:1061
		{
		logDebugGrammar("NUMBER %d", yyS[yypt-0].n)
		thisExpression := ast.NewLiteralNumber(float64(yyS[yypt-0].n))
		parsingStack.Push(thisExpression)
	}
	case 131:
		//line n1ql.y:1067
		{
		logDebugGrammar("NUMBER %f", yyS[yypt-0].f)
		thisExpression := ast.NewLiteralNumber(yyS[yypt-0].f)
		parsingStack.Push(thisExpression)
	}
	case 132:
		//line n1ql.y:1075
		{
		logDebugGrammar("EMPTY OBJECT")
		emptyObject := ast.NewLiteralObject(map[string]ast.Expression{})
		parsingStack.Push(emptyObject)
	}
	case 133:
		//line n1ql.y:1081
		{
		logDebugGrammar("OBJECT")
	}
	case 134:
		//line n1ql.y:1087
		{
		logDebugGrammar("NAMED EXPR LIST SINGLE")
	}
	case 135:
		//line n1ql.y:1091
		{
		logDebugGrammar("NAMED EXPR LIST COMPOUND")
		last := parsingStack.Pop().(*ast.LiteralObject)
		rest := parsingStack.Pop().(*ast.LiteralObject)
		for k,v := range last.Val {
			rest.Val[k] = v
		}
		parsingStack.Push(rest)
	}
	case 136:
		//line n1ql.y:1103
		{
		logDebugGrammar("NAMED EXPR SINGLE")
		thisKey := yyS[yypt-2].s
		thisValue := parsingStack.Pop().(ast.Expression)
		thisExpression := ast.NewLiteralObject(map[string]ast.Expression{thisKey: thisValue})
		parsingStack.Push(thisExpression)
	}
	case 137:
		//line n1ql.y:1113
		{
		logDebugGrammar("EMPTY ARRAY")
		thisExpression := ast.NewLiteralArray(ast.ExpressionList{})
		parsingStack.Push(thisExpression)
	}
	case 138:
		//line n1ql.y:1119
		{
		logDebugGrammar("ARRAY")
		exp_list := parsingStack.Pop().(ast.ExpressionList)
		thisExpression := ast.NewLiteralArray(exp_list)
		parsingStack.Push(thisExpression)
	}
	case 139:
		//line n1ql.y:1128
		{
		logDebugGrammar("EXPRESSION LIST SINGLE")
		exp_list := make(ast.ExpressionList, 0)
		exp_list = append(exp_list, parsingStack.Pop().(ast.Expression))
		parsingStack.Push(exp_list)
	}
	case 140:
		//line n1ql.y:1135
		{
		logDebugGrammar("EXPRESSION LIST COMPOUND")
		rest := parsingStack.Pop().(ast.ExpressionList)
		last := parsingStack.Pop()
		new_list := make(ast.ExpressionList, 0, len(rest) + 1)
		new_list = append(new_list, last.(ast.Expression))
		for _, v := range rest {
			new_list = append(new_list, v)
		}
		parsingStack.Push(new_list)
	}
	}
	goto yystack /* stack new state and value */
}
