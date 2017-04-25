//line cc.y:34
package cc

import __yyfmt__ "fmt"

//line cc.y:34
type typeClass struct {
	c Storage
	q TypeQual
	t *Type
}

type idecor struct {
	d func(*Type) (*Type, string)
	i *Init
}

//line cc.y:49
type yySymType struct {
	yys      int
	abdecor  func(*Type) *Type
	decl     *Decl
	decls    []*Decl
	decor    func(*Type) (*Type, string)
	decors   []func(*Type) (*Type, string)
	expr     *Expr
	exprs    []*Expr
	idec     idecor
	idecs    []idecor
	init     *Init
	inits    []*Init
	label    *Label
	labels   []*Label
	span     Span
	prefix   *Prefix
	prefixes []*Prefix
	stmt     *Stmt
	stmts    []*Stmt
	str      string
	strs     []string
	tc       typeClass
	tk       TypeKind
	typ      *Type
}

const tokAuto = 57346
const tokBreak = 57347
const tokCase = 57348
const tokChar = 57349
const tokConst = 57350
const tokContinue = 57351
const tokDefault = 57352
const tokDo = 57353
const tokDotDotDot = 57354
const tokDouble = 57355
const tokEnum = 57356
const tokError = 57357
const tokExtern = 57358
const tokFloat = 57359
const tokFor = 57360
const tokGoto = 57361
const tokIf = 57362
const tokInline = 57363
const tokInt = 57364
const tokLitChar = 57365
const tokLong = 57366
const tokName = 57367
const tokNumber = 57368
const tokOffsetof = 57369
const tokRegister = 57370
const tokReturn = 57371
const tokShort = 57372
const tokSigned = 57373
const tokStatic = 57374
const tokStruct = 57375
const tokSwitch = 57376
const tokTypeName = 57377
const tokTypedef = 57378
const tokUnion = 57379
const tokUnsigned = 57380
const tokVaArg = 57381
const tokVoid = 57382
const tokVolatile = 57383
const tokWhile = 57384
const tokString = 57385
const tokDefine = 57386
const tokShift = 57387
const tokElse = 57388
const tokAddEq = 57389
const tokSubEq = 57390
const tokMulEq = 57391
const tokDivEq = 57392
const tokModEq = 57393
const tokLshEq = 57394
const tokRshEq = 57395
const tokAndEq = 57396
const tokXorEq = 57397
const tokOrEq = 57398
const tokOrOr = 57399
const tokAndAnd = 57400
const tokEqEq = 57401
const tokNotEq = 57402
const tokLtEq = 57403
const tokGtEq = 57404
const tokLsh = 57405
const tokRsh = 57406
const tokCast = 57407
const tokSizeof = 57408
const tokUnary = 57409
const tokDec = 57410
const tokInc = 57411
const tokArrow = 57412
const startExpr = 57413
const startProg = 57414
const tokEOF = 57415

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"tokAuto",
	"tokBreak",
	"tokCase",
	"tokChar",
	"tokConst",
	"tokContinue",
	"tokDefault",
	"tokDo",
	"tokDotDotDot",
	"tokDouble",
	"tokEnum",
	"tokError",
	"tokExtern",
	"tokFloat",
	"tokFor",
	"tokGoto",
	"tokIf",
	"tokInline",
	"tokInt",
	"tokLitChar",
	"tokLong",
	"tokName",
	"tokNumber",
	"tokOffsetof",
	"tokRegister",
	"tokReturn",
	"tokShort",
	"tokSigned",
	"tokStatic",
	"tokStruct",
	"tokSwitch",
	"tokTypeName",
	"tokTypedef",
	"tokUnion",
	"tokUnsigned",
	"tokVaArg",
	"tokVoid",
	"tokVolatile",
	"tokWhile",
	"tokString",
	"tokDefine",
	"tokShift",
	"tokElse",
	"'{'",
	"','",
	"'='",
	"tokAddEq",
	"tokSubEq",
	"tokMulEq",
	"tokDivEq",
	"tokModEq",
	"tokLshEq",
	"tokRshEq",
	"tokAndEq",
	"tokXorEq",
	"tokOrEq",
	"'?'",
	"':'",
	"tokOrOr",
	"tokAndAnd",
	"'|'",
	"'^'",
	"'&'",
	"tokEqEq",
	"tokNotEq",
	"'<'",
	"'>'",
	"tokLtEq",
	"tokGtEq",
	"tokLsh",
	"tokRsh",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"tokCast",
	"'!'",
	"'~'",
	"tokSizeof",
	"tokUnary",
	"'.'",
	"'['",
	"']'",
	"'('",
	"')'",
	"tokDec",
	"tokInc",
	"tokArrow",
	"startExpr",
	"startProg",
	"tokEOF",
	"'}'",
	"';'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 118,
	48, 96,
	97, 96,
	-2, 178,
	-1, 137,
	47, 169,
	-2, 143,
	-1, 139,
	47, 169,
	-2, 148,
	-1, 238,
	97, 204,
	-2, 168,
	-1, 269,
	61, 169,
	-2, 87,
}

const yyNprod = 214
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1338

var yyAct = [...]int{

	7, 309, 120, 259, 113, 336, 229, 225, 196, 31,
	266, 279, 240, 99, 100, 101, 102, 103, 104, 105,
	106, 107, 50, 337, 112, 237, 219, 5, 193, 214,
	217, 223, 117, 4, 123, 134, 110, 132, 227, 130,
	118, 137, 139, 64, 65, 66, 111, 364, 362, 352,
	347, 97, 93, 345, 92, 331, 95, 94, 96, 35,
	36, 330, 141, 142, 143, 144, 145, 146, 147, 148,
	149, 150, 151, 152, 153, 154, 155, 156, 157, 158,
	159, 329, 161, 162, 163, 164, 165, 166, 167, 168,
	169, 170, 171, 128, 129, 136, 135, 289, 288, 175,
	176, 33, 187, 305, 160, 292, 67, 68, 62, 63,
	64, 65, 66, 244, 60, 367, 185, 174, 97, 93,
	296, 92, 6, 95, 94, 96, 361, 197, 97, 93,
	355, 92, 111, 95, 94, 96, 131, 3, 2, 354,
	211, 353, 181, 249, 195, 177, 178, 284, 78, 77,
	76, 75, 73, 74, 69, 70, 71, 72, 67, 68,
	62, 63, 64, 65, 66, 283, 199, 198, 295, 200,
	97, 93, 124, 92, 351, 95, 94, 96, 252, 206,
	208, 189, 125, 188, 235, 124, 190, 182, 136, 135,
	224, 226, 129, 136, 135, 125, 230, 189, 184, 188,
	339, 232, 233, 204, 242, 195, 180, 208, 243, 179,
	338, 212, 224, 221, 335, 173, 333, 234, 209, 213,
	216, 31, 115, 189, 121, 188, 238, 231, 109, 131,
	255, 282, 258, 205, 211, 122, 299, 256, 221, 232,
	274, 269, 247, 246, 248, 209, 226, 267, 257, 342,
	341, 254, 250, 277, 62, 63, 64, 65, 66, 291,
	203, 238, 263, 261, 97, 93, 290, 92, 271, 95,
	94, 96, 280, 281, 253, 294, 207, 61, 221, 195,
	285, 192, 301, 300, 293, 286, 260, 202, 201, 186,
	298, 363, 304, 269, 116, 233, 98, 343, 226, 267,
	58, 303, 268, 124, 241, 306, 53, 230, 126, 308,
	58, 34, 297, 125, 270, 315, 247, 228, 114, 287,
	1, 332, 38, 57, 334, 11, 194, 340, 133, 49,
	56, 302, 312, 59, 54, 316, 278, 311, 55, 344,
	313, 138, 140, 59, 245, 127, 276, 119, 172, 272,
	273, 264, 265, 239, 236, 358, 359, 360, 357, 346,
	27, 26, 348, 349, 218, 366, 191, 30, 365, 368,
	28, 183, 0, 0, 0, 356, 81, 82, 83, 84,
	85, 86, 87, 88, 89, 90, 91, 80, 350, 79,
	78, 77, 76, 75, 73, 74, 69, 70, 71, 72,
	67, 68, 62, 63, 64, 65, 66, 0, 0, 0,
	0, 0, 97, 93, 0, 92, 0, 95, 94, 96,
	81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
	91, 80, 0, 79, 78, 77, 76, 75, 73, 74,
	69, 70, 71, 72, 67, 68, 62, 63, 64, 65,
	66, 0, 0, 0, 0, 0, 97, 93, 307, 92,
	0, 95, 94, 96, 81, 82, 83, 84, 85, 86,
	87, 88, 89, 90, 91, 80, 0, 79, 78, 77,
	76, 75, 73, 74, 69, 70, 71, 72, 67, 68,
	62, 63, 64, 65, 66, 0, 0, 0, 0, 0,
	97, 93, 0, 92, 275, 95, 94, 96, 215, 81,
	82, 83, 84, 85, 86, 87, 88, 89, 90, 91,
	80, 0, 79, 78, 77, 76, 75, 73, 74, 69,
	70, 71, 72, 67, 68, 62, 63, 64, 65, 66,
	0, 0, 0, 0, 0, 97, 93, 0, 92, 0,
	95, 94, 96, 81, 82, 83, 84, 85, 86, 87,
	88, 89, 90, 91, 80, 0, 79, 78, 77, 76,
	75, 73, 74, 69, 70, 71, 72, 67, 68, 62,
	63, 64, 65, 66, 0, 0, 0, 0, 0, 97,
	93, 0, 92, 0, 95, 94, 96, 53, 0, 0,
	40, 58, 0, 0, 0, 0, 47, 39, 0, 114,
	46, 0, 0, 0, 57, 42, 10, 43, 8, 9,
	21, 56, 0, 41, 44, 54, 51, 0, 37, 55,
	52, 45, 23, 48, 59, 80, 24, 79, 78, 77,
	76, 75, 73, 74, 69, 70, 71, 72, 67, 68,
	62, 63, 64, 65, 66, 0, 0, 0, 0, 13,
	97, 93, 0, 92, 0, 95, 94, 96, 14, 15,
	12, 0, 0, 0, 16, 17, 20, 0, 0, 0,
	0, 22, 0, 19, 18, 317, 326, 0, 0, 318,
	327, 319, 0, 0, 0, 0, 0, 0, 320, 321,
	322, 0, 0, 10, 0, 328, 9, 21, 53, 323,
	0, 40, 58, 0, 324, 0, 0, 47, 39, 23,
	29, 46, 325, 24, 0, 57, 42, 260, 43, 0,
	0, 0, 56, 0, 41, 44, 54, 51, 0, 37,
	55, 52, 45, 0, 48, 59, 13, 0, 32, 0,
	0, 0, 0, 0, 0, 14, 15, 12, 0, 0,
	0, 16, 17, 20, 0, 0, 0, 0, 22, 0,
	19, 18, 0, 0, 0, 53, 0, 314, 40, 58,
	0, 0, 0, 0, 47, 39, 0, 114, 46, 0,
	0, 0, 57, 42, 0, 43, 0, 0, 0, 56,
	251, 41, 44, 54, 51, 0, 37, 55, 52, 45,
	53, 48, 59, 40, 58, 0, 0, 0, 0, 47,
	39, 0, 114, 46, 0, 0, 0, 57, 42, 0,
	43, 0, 0, 0, 56, 0, 41, 44, 54, 51,
	0, 37, 55, 52, 45, 53, 48, 59, 40, 58,
	0, 0, 0, 0, 47, 39, 0, 29, 46, 0,
	0, 0, 57, 42, 0, 43, 0, 310, 0, 56,
	0, 41, 44, 54, 51, 0, 37, 55, 52, 45,
	0, 48, 59, 0, 0, 32, 77, 76, 75, 73,
	74, 69, 70, 71, 72, 67, 68, 62, 63, 64,
	65, 66, 262, 0, 0, 0, 0, 97, 93, 0,
	92, 0, 95, 94, 96, 0, 0, 76, 75, 73,
	74, 69, 70, 71, 72, 67, 68, 62, 63, 64,
	65, 66, 0, 0, 0, 0, 25, 97, 93, 0,
	92, 0, 95, 94, 96, 75, 73, 74, 69, 70,
	71, 72, 67, 68, 62, 63, 64, 65, 66, 0,
	0, 0, 0, 0, 97, 93, 0, 92, 0, 95,
	94, 96, 73, 74, 69, 70, 71, 72, 67, 68,
	62, 63, 64, 65, 66, 10, 0, 8, 9, 21,
	97, 93, 0, 92, 0, 95, 94, 96, 0, 0,
	0, 23, 0, 0, 0, 24, 0, 0, 0, 210,
	0, 0, 74, 69, 70, 71, 72, 67, 68, 62,
	63, 64, 65, 66, 0, 0, 0, 0, 13, 97,
	93, 0, 92, 0, 95, 94, 96, 14, 15, 12,
	0, 0, 0, 16, 17, 20, 0, 280, 281, 0,
	22, 0, 19, 18, 69, 70, 71, 72, 67, 68,
	62, 63, 64, 65, 66, 10, 0, 8, 9, 21,
	97, 93, 0, 92, 0, 95, 94, 96, 0, 0,
	0, 23, 0, 0, 10, 24, 8, 9, 21, 210,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	23, 0, 0, 0, 24, 0, 0, 0, 13, 0,
	0, 0, 0, 0, 0, 0, 0, 14, 15, 12,
	0, 0, 0, 16, 17, 20, 0, 13, 0, 0,
	22, 0, 19, 18, 0, 0, 14, 15, 12, 0,
	0, 0, 16, 17, 20, 0, 0, 0, 0, 22,
	0, 19, 18, 10, 0, 8, 9, 21, 0, 0,
	0, 0, 0, 0, 10, 0, 8, 9, 21, 23,
	0, 0, 0, 24, 0, 0, 0, 0, 0, 0,
	23, 0, 0, 0, 24, 0, 0, 0, 210, 0,
	0, 0, 0, 0, 0, 0, 13, 0, 0, 0,
	0, 0, 0, 0, 0, 14, 15, 12, 0, 0,
	0, 16, 17, 20, 0, 0, 0, 0, 108, 0,
	19, 18, 16, 17, 20, 0, 0, 0, 0, 22,
	53, 19, 18, 40, 58, 0, 0, 0, 222, 47,
	39, 0, 114, 46, 0, 0, 0, 57, 42, 0,
	43, 220, 0, 0, 56, 0, 41, 44, 54, 51,
	0, 37, 55, 52, 45, 53, 48, 59, 40, 58,
	0, 0, 0, 0, 47, 39, 0, 114, 46, 0,
	0, 0, 57, 42, 0, 43, 0, 0, 0, 56,
	0, 41, 44, 54, 51, 0, 37, 55, 52, 45,
	53, 48, 59, 40, 58, 0, 0, 0, 0, 47,
	0, 0, 114, 46, 0, 0, 0, 57, 42, 0,
	43, 0, 0, 0, 56, 0, 41, 44, 54, 0,
	0, 0, 55, 0, 45, 0, 48, 59,
}
var yyPact = [...]int{

	44, -1000, -1000, 1061, 841, 19, 229, 504, -1000, -1000,
	-1000, 253, 1061, 1061, 1061, 1061, 1061, 1061, 1061, 1061,
	1130, 140, 593, 134, -1000, -1000, -1000, -1000, -1000, 251,
	-1000, 147, 283, 1261, 302, 1296, -1000, -1000, 278, 278,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 1061, 1061, 1061, 1061, 1061, 1061, 1061, 1061, 1061,
	1061, 1061, 1061, 1061, 1061, 1061, 1061, 1061, 1061, 1061,
	1061, 1061, 1061, 1061, 1061, 1061, 1061, 1061, 1061, 1061,
	1061, 1061, 1061, 1061, -1000, -1000, 278, 278, -1000, 43,
	43, 43, 43, 43, 43, 43, 43, 43, 593, 1261,
	120, 117, 110, -1000, -1000, 1061, 242, 5, 137, 233,
	-1000, 292, 147, -1000, -1000, -1000, 1061, 302, 1296, -1000,
	-1000, 302, -1000, 1296, -1000, -1000, -1000, -1000, 241, -1000,
	240, 504, -34, -34, 43, 43, 43, 179, 179, 33,
	33, 33, 33, 944, 985, 905, 879, 852, 822, 85,
	199, 504, 504, 504, 504, 504, 504, 504, 504, 504,
	504, 504, 114, 229, 146, -1000, -1000, 90, 228, 1042,
	-1000, 148, 292, 131, 110, 460, -1000, -1000, 1226, 1061,
	1042, 1261, 147, 147, 292, -1000, 95, 504, -1000, -1000,
	-1000, 1261, 279, 1061, -1000, -1000, 1141, 1061, 43, -1000,
	17, 1061, 110, 1226, 54, 1261, 704, 89, 226, -1000,
	-1000, 160, -1000, 145, 504, -1000, 504, -1000, 239, -1000,
	147, -1000, 137, 111, -1000, -1000, 806, -1000, 147, 220,
	-1000, 191, 575, 415, -1000, 962, 144, 148, 76, -1000,
	58, -1000, -1000, 1226, 148, 111, 292, 160, -1000, -1000,
	-1000, 1, -1000, -1000, 0, 218, -1000, 111, 198, -1000,
	9, 279, -1000, -1000, 1061, -1000, 72, -1000, 187, -1000,
	278, 1061, -1000, -1000, -1000, -1000, 160, -1000, -1000, -1000,
	147, 1061, -1000, -1000, 504, -1000, 7, 1042, -1000, -1000,
	-1000, 371, 771, -1000, 504, -1000, -1000, -1000, -1000, -1000,
	-1000, 680, -1000, -1000, -1000, -1000, -16, -36, -42, -1000,
	128, 278, 126, 1061, 122, 112, 1061, 189, 188, -1000,
	-1000, -1000, 255, 1061, -44, 1061, -47, -1000, 1061, 1061,
	327, -1000, -1000, 86, -48, -1000, 52, -1000, 50, 41,
	-1000, 1061, 1061, -1000, -1000, -1000, 37, -49, 245, -1000,
	-1000, -50, 1061, -1000, -1000, 26, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 29, 371, 26, 370, 12, 367, 6, 366, 364,
	30, 33, 361, 360, 25, 354, 353, 8, 10, 352,
	351, 0, 31, 23, 5, 350, 349, 122, 348, 34,
	347, 32, 7, 346, 38, 344, 340, 337, 11, 336,
	332, 3, 1, 331, 329, 22, 59, 60, 35, 2,
	302, 101, 39, 328, 37, 326, 28, 325, 4, 322,
	36, 24, 311, 320, 319, 317, 314, 312,
}
var yyR1 = [...]int{

	0, 63, 63, 11, 11, 23, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 43,
	43, 43, 64, 41, 36, 36, 36, 42, 40, 40,
	40, 40, 40, 40, 40, 40, 40, 40, 40, 40,
	40, 1, 1, 1, 2, 2, 2, 17, 17, 17,
	17, 17, 3, 3, 3, 3, 29, 29, 44, 44,
	44, 44, 44, 44, 45, 45, 46, 46, 46, 46,
	46, 46, 46, 46, 46, 47, 47, 48, 48, 62,
	58, 58, 58, 58, 58, 61, 60, 7, 13, 6,
	12, 12, 12, 12, 65, 4, 49, 49, 59, 59,
	18, 18, 14, 62, 62, 38, 21, 21, 62, 62,
	5, 25, 32, 32, 34, 34, 34, 35, 35, 33,
	33, 38, 67, 67, 66, 66, 39, 39, 50, 50,
	24, 24, 22, 22, 27, 27, 28, 28, 8, 8,
	37, 37, 9, 9, 10, 10, 30, 30, 31, 31,
	55, 55, 56, 56, 51, 51, 52, 52, 53, 53,
	54, 54, 19, 19, 20, 20, 15, 15, 26, 26,
	16, 16, 57, 57,
}
var yyR2 = [...]int{

	0, 3, 3, 0, 2, 1, 1, 1, 1, 1,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 5, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 4,
	6, 4, 4, 3, 4, 4, 2, 2, 6, 0,
	2, 2, 0, 4, 3, 2, 2, 2, 1, 1,
	2, 2, 2, 7, 9, 3, 5, 7, 3, 5,
	5, 0, 3, 1, 4, 4, 3, 1, 3, 3,
	4, 4, 1, 2, 2, 1, 1, 3, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 3, 3, 2, 2, 1, 2, 3, 3, 3,
	1, 1, 5, 1, 0, 5, 1, 1, 1, 1,
	1, 3, 3, 2, 5, 2, 3, 3, 2, 6,
	2, 2, 1, 1, 2, 4, 5, 0, 3, 1,
	3, 3, 0, 1, 0, 1, 1, 2, 0, 1,
	0, 1, 0, 1, 1, 3, 0, 1, 0, 2,
	0, 2, 1, 3, 0, 1, 1, 3, 0, 1,
	1, 2, 0, 1, 1, 2, 0, 1, 1, 2,
	0, 1, 1, 3, 0, 1, 1, 2, 0, 1,
	1, 3, 1, 2,
}
var yyChk = [...]int{

	-1000, -63, 94, 93, -11, -23, -27, -21, 25, 26,
	23, -57, 77, 66, 75, 76, 81, 82, 91, 90,
	83, 27, 88, 39, 43, 95, -12, -13, -4, 16,
	-6, -58, 44, -51, -62, -46, -47, 35, -59, 14,
	7, 30, 22, 24, 31, 38, 17, 13, 40, -44,
	-45, 33, 37, 4, 32, 36, 28, 21, 8, 41,
	95, 48, 75, 76, 77, 78, 79, 73, 74, 69,
	70, 71, 72, 67, 68, 66, 65, 64, 63, 62,
	60, 49, 50, 51, 52, 53, 54, 55, 56, 57,
	58, 59, 88, 86, 91, 90, 92, 85, 43, -21,
	-21, -21, -21, -21, -21, -21, -21, -21, 88, 88,
	-60, -23, -61, -58, 16, 88, 43, -31, -17, -30,
	-49, 77, 88, -29, 25, 35, 25, -62, -46, -47,
	-52, -51, -54, -53, -48, -47, -46, -49, -50, -49,
	-50, -21, -21, -21, -21, -21, -21, -21, -21, -21,
	-21, -21, -21, -21, -21, -21, -21, -21, -21, -21,
	-23, -21, -21, -21, -21, -21, -21, -21, -21, -21,
	-21, -21, -28, -27, -23, -49, -49, -60, -60, 89,
	89, -1, 77, -2, 88, -21, 47, 97, 88, 86,
	49, -8, 48, -56, -55, -45, -17, -21, -52, -54,
	-48, 47, 47, 61, 89, 87, 89, 48, -21, -34,
	47, 86, -56, 88, -1, 48, -11, -10, -9, -3,
	25, -61, 12, -22, -21, -32, -21, -34, -65, -7,
	-58, -29, -17, -17, -45, 89, -15, -14, -61, -16,
	-5, 25, -21, -21, 96, -35, -22, -1, -10, 89,
	-60, 96, 89, 48, -1, -17, 77, 88, 87, -41,
	47, -31, 96, -14, -20, -19, -18, -17, -50, -49,
	-66, 48, -26, -25, 49, 89, -33, -32, -39, -38,
	85, 86, 87, 89, 89, -3, -56, -64, 97, 97,
	48, 61, 96, -5, -21, 96, 48, -67, -38, 49,
	-49, -21, -43, -18, -21, 96, -32, 87, -7, -42,
	96, -37, -40, -36, 97, -41, -23, 5, 9, 11,
	18, 19, 20, 29, 34, 42, 6, 10, 25, 97,
	97, 97, -42, 88, -49, 88, -24, -23, 88, 88,
	-21, 61, 61, 42, -24, 97, -23, 97, -23, -23,
	61, 88, 97, 89, 89, 89, -23, -24, -42, -42,
	-42, 89, 97, 46, 97, -24, -42, 89, -42,
}
var yyDef = [...]int{

	0, -2, 3, 0, 0, 0, 5, 174, 6, 7,
	8, 9, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 212, 1, 4, 130, 131, 100,
	133, 188, 0, 120, 196, 200, 194, 119, 168, 168,
	106, 107, 108, 109, 110, 111, 112, 113, 114, 115,
	116, 138, 139, 98, 99, 101, 102, 103, 104, 105,
	2, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 176, 0, 56, 57, 0, 0, 213, 40,
	41, 42, 43, 44, 45, 46, 47, 48, 0, 0,
	0, 0, 81, 125, 100, 0, 0, 0, -2, 189,
	87, 192, 0, 186, 136, 137, 0, 196, 200, 195,
	123, 197, 124, 201, 198, 117, 118, -2, 0, -2,
	0, 175, 10, 11, 12, 13, 14, 15, 16, 17,
	18, 19, 20, 21, 22, 23, 24, 25, 26, 27,
	0, 29, 30, 31, 32, 33, 34, 35, 36, 37,
	38, 39, 0, 177, 0, 146, 147, 0, 0, 0,
	53, 126, 192, 83, 81, 0, 3, 128, 184, 172,
	0, 134, 0, 0, 193, 190, 0, 129, 121, 122,
	199, 0, 0, 0, 54, 55, 49, 0, 51, 52,
	157, 172, 81, 184, 0, 0, 0, 0, 185, 182,
	92, 81, 95, 0, 173, 97, 152, 153, 0, 179,
	188, 187, 96, 88, 191, 89, 0, 206, -2, 164,
	210, 208, 28, 0, 154, 0, 0, 82, 0, 86,
	0, 132, 90, 0, 93, 94, 192, 81, 91, 135,
	62, 0, 144, 207, 0, 205, 202, 140, 0, -2,
	0, 165, 150, 209, 0, 50, 0, 159, 162, 166,
	0, 0, 85, 84, 58, 183, 81, 59, 127, 142,
	168, 0, 149, 211, 151, 155, 158, 0, 167, 163,
	145, 0, 180, 203, 141, 156, 160, 161, 60, 61,
	63, 0, 67, 181, 68, 69, 0, 0, 0, 180,
	0, 0, 0, 170, 0, 0, 0, 0, 6, 70,
	71, 72, 0, 170, 0, 0, 0, 171, 0, 0,
	0, 65, 66, 0, 0, 75, 0, 78, 0, 0,
	64, 0, 170, 180, 180, 180, 0, 0, 76, 79,
	80, 0, 170, 180, 73, 0, 77, 180, 74,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 81, 3, 3, 3, 79, 66, 3,
	88, 89, 77, 75, 48, 76, 85, 78, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 61, 97,
	69, 49, 70, 60, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 86, 3, 87, 65, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 47, 64, 96, 82,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 50, 51, 52, 53, 54,
	55, 56, 57, 58, 59, 62, 63, 67, 68, 71,
	72, 73, 74, 80, 83, 84, 90, 91, 92, 93,
	94, 95,
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
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:180
		{
			yylex.(*lexer).prog = &Prog{Decls: yyDollar[2].decls}
			return 0
		}
	case 2:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:185
		{
			yylex.(*lexer).expr = yyDollar[2].expr
			return 0
		}
	case 3:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:191
		{
			yyVAL.span = Span{}
			yyVAL.decls = nil
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:196
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.decls = append(yyDollar[1].decls, yyDollar[2].decls...)
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:203
		{
			yyVAL.span = yyDollar[1].span
			if len(yyDollar[1].exprs) == 1 {
				yyVAL.expr = yyDollar[1].exprs[0]
				break
			}
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Comma, List: yyDollar[1].exprs}
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:214
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Name, Text: yyDollar[1].str, XDecl: yyDollar[1].decl}
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:219
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Number, Text: yyDollar[1].str}
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:224
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Number, Text: yyDollar[1].str}
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:229
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: String, Texts: yyDollar[1].strs}
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:234
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Add, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:239
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Sub, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:244
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Mul, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:249
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Div, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:254
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Mod, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:259
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Lsh, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:264
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Rsh, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:269
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Lt, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:274
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Gt, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:279
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: LtEq, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:284
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: GtEq, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:289
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: EqEq, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:294
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: NotEq, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:299
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: And, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:304
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Xor, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:309
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Or, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:314
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: AndAnd, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:319
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: OrOr, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 28:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line cc.y:324
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[5].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Cond, List: []*Expr{yyDollar[1].expr, yyDollar[3].expr, yyDollar[5].expr}}
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:329
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Eq, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:334
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: AddEq, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:339
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: SubEq, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:344
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: MulEq, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:349
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: DivEq, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:354
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: ModEq, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:359
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: LshEq, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:364
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: RshEq, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:369
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: AndEq, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:374
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: XorEq, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:379
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: OrEq, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 40:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:384
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Indir, Left: yyDollar[2].expr}
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:389
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Addr, Left: yyDollar[2].expr}
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:394
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Plus, Left: yyDollar[2].expr}
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:399
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Minus, Left: yyDollar[2].expr}
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:404
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Not, Left: yyDollar[2].expr}
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:409
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Twid, Left: yyDollar[2].expr}
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:414
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: PreInc, Left: yyDollar[2].expr}
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:419
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: PreDec, Left: yyDollar[2].expr}
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:424
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: SizeofExpr, Left: yyDollar[2].expr}
		}
	case 49:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line cc.y:429
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[4].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: SizeofType, Type: yyDollar[3].typ}
		}
	case 50:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line cc.y:434
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[6].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Offsetof, Type: yyDollar[3].typ, Left: yyDollar[5].expr}
		}
	case 51:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line cc.y:439
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[4].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Cast, Type: yyDollar[2].typ, Left: yyDollar[4].expr}
		}
	case 52:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line cc.y:444
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[4].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: CastInit, Type: yyDollar[2].typ, Init: &Init{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Braced: yyDollar[4].inits}}
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:449
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Paren, Left: yyDollar[2].expr}
		}
	case 54:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line cc.y:454
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[4].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Call, Left: yyDollar[1].expr, List: yyDollar[3].exprs}
		}
	case 55:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line cc.y:459
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[4].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Index, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:464
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: PostInc, Left: yyDollar[1].expr}
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:469
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: PostDec, Left: yyDollar[1].expr}
		}
	case 58:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line cc.y:474
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[6].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: VaArg, Left: yyDollar[3].expr, Type: yyDollar[5].typ}
		}
	case 59:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:480
		{
			yyVAL.span = Span{}
			yyVAL.stmts = nil
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:485
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.stmts = yyDollar[1].stmts
			for _, d := range yyDollar[2].decls {
				yyVAL.stmts = append(yyVAL.stmts, &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: StmtDecl, Decl: d})
			}
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:493
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[2].stmt)
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:500
		{
			yylex.(*lexer).pushScope()
		}
	case 63:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line cc.y:504
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[4].span)
			yylex.(*lexer).popScope()
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Block, Block: yyDollar[3].stmts}
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:512
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.label = &Label{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Case, Expr: yyDollar[2].expr}
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:517
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.label = &Label{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Default}
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:522
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.label = &Label{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: LabelName, Name: yyDollar[1].str}
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:529
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.stmt = yyDollar[2].stmt
			yyVAL.stmt.Labels = yyDollar[1].labels
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:537
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Empty}
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:542
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:547
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: StmtExpr, Expr: yyDollar[1].expr}
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:552
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Break}
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:557
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Continue}
		}
	case 73:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line cc.y:562
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[7].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Do, Body: yyDollar[2].stmt, Expr: yyDollar[5].expr}
		}
	case 74:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line cc.y:567
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[9].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Op:   For,
				Pre:  yyDollar[3].expr,
				Expr: yyDollar[5].expr,
				Post: yyDollar[7].expr,
				Body: yyDollar[9].stmt,
			}
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:578
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Goto, Text: yyDollar[2].str}
		}
	case 76:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line cc.y:583
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[5].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: If, Expr: yyDollar[3].expr, Body: yyDollar[5].stmt}
		}
	case 77:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line cc.y:588
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[7].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: If, Expr: yyDollar[3].expr, Body: yyDollar[5].stmt, Else: yyDollar[7].stmt}
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:593
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Return, Expr: yyDollar[2].expr}
		}
	case 79:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line cc.y:598
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[5].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Switch, Expr: yyDollar[3].expr, Body: yyDollar[5].stmt}
		}
	case 80:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line cc.y:603
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[5].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: While, Expr: yyDollar[3].expr, Body: yyDollar[5].stmt}
		}
	case 81:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:610
		{
			yyVAL.span = Span{}
			yyVAL.abdecor = func(t *Type) *Type { return t }
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:615
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			_, q, _ := splitTypeWords(yyDollar[2].strs)
			abdecor := yyDollar[3].abdecor
			yyVAL.abdecor = func(t *Type) *Type {
				return abdecor(&Type{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Kind: Ptr, Base: t, Qual: q})
			}
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:624
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.abdecor = yyDollar[1].abdecor
		}
	case 84:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line cc.y:631
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[4].span)
			abdecor := yyDollar[1].abdecor
			decls := yyDollar[3].decls
			span := yyVAL.span
			for _, decl := range decls {
				t := decl.Type
				if t != nil {
					if t.Kind == TypedefType && t.Base != nil {
						t = t.Base
					}
					if t.Kind == Array {
						if t.Width == nil {
							t = t.Base
						}
						decl.Type = &Type{Kind: Ptr, Base: t}
					}
				}
			}
			yyVAL.abdecor = func(t *Type) *Type {
				return abdecor(&Type{SyntaxInfo: SyntaxInfo{Span: span}, Kind: Func, Base: t, Decls: decls})
			}
		}
	case 85:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line cc.y:655
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[4].span)
			abdecor := yyDollar[1].abdecor
			span := yyVAL.span
			expr := yyDollar[3].expr
			yyVAL.abdecor = func(t *Type) *Type {
				return abdecor(&Type{SyntaxInfo: SyntaxInfo{Span: span}, Kind: Array, Base: t, Width: expr})
			}

		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:666
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.abdecor = yyDollar[2].abdecor
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:674
		{
			yyVAL.span = yyDollar[1].span
			name := yyDollar[1].str
			yyVAL.decor = func(t *Type) (*Type, string) { return t, name }
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:680
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			_, q, _ := splitTypeWords(yyDollar[2].strs)
			decor := yyDollar[3].decor
			span := yyVAL.span
			yyVAL.decor = func(t *Type) (*Type, string) {
				return decor(&Type{SyntaxInfo: SyntaxInfo{Span: span}, Kind: Ptr, Base: t, Qual: q})
			}
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:690
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.decor = yyDollar[2].decor
		}
	case 90:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line cc.y:695
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[4].span)
			decor := yyDollar[1].decor
			decls := yyDollar[3].decls
			span := yyVAL.span
			yyVAL.decor = func(t *Type) (*Type, string) {
				return decor(&Type{SyntaxInfo: SyntaxInfo{Span: span}, Kind: Func, Base: t, Decls: decls})
			}
		}
	case 91:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line cc.y:705
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[4].span)
			decor := yyDollar[1].decor
			span := yyVAL.span
			expr := yyDollar[3].expr
			yyVAL.decor = func(t *Type) (*Type, string) {
				return decor(&Type{SyntaxInfo: SyntaxInfo{Span: span}, Kind: Array, Base: t, Width: expr})
			}
		}
	case 92:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:718
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.decl = &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Name: yyDollar[1].str}
		}
	case 93:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:723
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.decl = &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Type: yyDollar[2].abdecor(yyDollar[1].typ)}
		}
	case 94:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:728
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			typ, name := yyDollar[2].decor(yyDollar[1].typ)
			yyVAL.decl = &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Name: name, Type: typ}
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:734
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.decl = &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Name: "..."}
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:742
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.idec = idecor{yyDollar[1].decor, nil}
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:747
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.idec = idecor{yyDollar[1].decor, yyDollar[3].init}
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:755
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:760
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:765
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:770
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:775
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:780
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:788
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:793
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:801
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:806
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:811
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:816
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:821
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 111:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:826
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:831
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:836
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:841
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:848
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:853
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:860
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:865
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:873
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.typ = yyDollar[1].typ
			if yyVAL.typ == nil {
				yyVAL.typ = &Type{Kind: TypedefType, Name: yyDollar[1].str}
			}
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:889
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.tc.c, yyVAL.tc.q, yyVAL.tc.t = splitTypeWords(append(yyDollar[1].strs, "int"))
		}
	case 121:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:894
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.tc.c, yyVAL.tc.q, _ = splitTypeWords(append(yyDollar[1].strs, yyDollar[3].strs...))
			yyVAL.tc.t = yyDollar[2].typ
		}
	case 122:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:900
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyDollar[1].strs = append(yyDollar[1].strs, yyDollar[2].str)
			yyDollar[1].strs = append(yyDollar[1].strs, yyDollar[3].strs...)
			yyVAL.tc.c, yyVAL.tc.q, yyVAL.tc.t = splitTypeWords(yyDollar[1].strs)
		}
	case 123:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:907
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.tc.c, yyVAL.tc.q, _ = splitTypeWords(yyDollar[2].strs)
			yyVAL.tc.t = yyDollar[1].typ
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:913
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			var ts []string
			ts = append(ts, yyDollar[1].str)
			ts = append(ts, yyDollar[2].strs...)
			yyVAL.tc.c, yyVAL.tc.q, yyVAL.tc.t = splitTypeWords(ts)
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:924
		{
			yyVAL.span = yyDollar[1].span
			if yyDollar[1].tc.c != 0 {
				yylex.(*lexer).Errorf("%v not allowed here", yyDollar[1].tc.c)
			}
			if yyDollar[1].tc.q != 0 && yyDollar[1].tc.q != Const && yyDollar[1].tc.q != Volatile {
				yylex.(*lexer).Errorf("%v ignored here (TODO)?", yyDollar[1].tc.q)
			}
			yyVAL.typ = yyDollar[1].tc.t
		}
	case 126:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:937
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.typ = yyDollar[2].abdecor(yyDollar[1].typ)
		}
	case 127:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:945
		{
			lx := yylex.(*lexer)
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			// TODO: use $1.q
			yyVAL.decls = nil
			for _, idec := range yyDollar[2].idecs {
				typ, name := idec.d(yyDollar[1].tc.t)
				d := &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Name: name, Type: typ, Storage: yyDollar[1].tc.c, Init: idec.i}
				lx.pushDecl(d)
				yyVAL.decls = append(yyVAL.decls, d)
			}
			if yyDollar[2].idecs == nil {
				d := &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Name: "", Type: yyDollar[1].tc.t, Storage: yyDollar[1].tc.c}
				lx.pushDecl(d)
				yyVAL.decls = append(yyVAL.decls, d)
			}
		}
	case 128:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:965
		{
			lx := yylex.(*lexer)
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			// TODO: use $1.q
			yyVAL.decls = nil
			for _, idec := range yyDollar[2].idecs {
				typ, name := idec.d(yyDollar[1].tc.t)
				d := lx.lookupDecl(name)
				if d == nil {
					d = &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Name: name, Type: typ, Storage: yyDollar[1].tc.c, Init: idec.i}
					lx.pushDecl(d)
				} else {
					d.Span = yyVAL.span
					if idec.i != nil {
						d.Init = idec.i
					}
				}
				yyVAL.decls = append(yyVAL.decls, d)
			}
			if yyDollar[2].idecs == nil {
				d := &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Name: "", Type: yyDollar[1].tc.t, Storage: yyDollar[1].tc.c}
				lx.pushDecl(d)
				yyVAL.decls = append(yyVAL.decls, d)
			}
		}
	case 129:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:993
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.decl = &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Name: yyDollar[2].str, Init: &Init{SyntaxInfo: SyntaxInfo{Span: yyDollar[3].span}, Expr: yyDollar[3].expr}}
			yylex.(*lexer).pushDecl(yyVAL.decl)
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1001
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.decls = yyDollar[1].decls
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1006
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.decls = []*Decl{yyDollar[1].decl}
		}
	case 132:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line cc.y:1011
		{
			yyVAL.decls = yyDollar[4].decls
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1015
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.decls = []*Decl{yyDollar[1].decl}
		}
	case 134:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1022
		{
			lx := yylex.(*lexer)
			typ, name := yyDollar[2].decor(yyDollar[1].tc.t)
			if typ.Kind != Func {
				yylex.(*lexer).Errorf("invalid function definition")
				return 0
			}
			d := lx.lookupDecl(name)
			if d == nil {
				d = &Decl{Name: name, Type: typ, Storage: yyDollar[1].tc.c}
				lx.pushDecl(d)
			} else {
				d.Type = typ
			}
			yyVAL.decl = d
			lx.pushScope()
			for _, decl := range typ.Decls {
				lx.pushDecl(decl)
			}
		}
	case 135:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line cc.y:1043
		{
			yylex.(*lexer).popScope()
			yyVAL.span = span(yyDollar[1].span, yyDollar[5].span)
			yyVAL.decl = yyDollar[4].decl
			yyVAL.decl.Span = yyVAL.span
			if yyDollar[3].decls != nil {
				yylex.(*lexer).Errorf("cannot use pre-prototype definitions")
			}
			yyVAL.decl.Body = yyDollar[5].stmt
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1056
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1061
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1069
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.tk = Struct
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1074
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.tk = Union
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1081
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.decor = yyDollar[1].decor
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1086
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			name := yyDollar[1].str
			expr := yyDollar[3].expr
			yyVAL.decor = func(t *Type) (*Type, string) {
				t.Width = expr
				return t, name
			}
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1098
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.decls = nil
			for _, decor := range yyDollar[2].decors {
				typ, name := decor(yyDollar[1].typ)
				yyVAL.decls = append(yyVAL.decls, &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Name: name, Type: typ})
			}
			if yyDollar[2].decors == nil {
				yyVAL.decls = append(yyVAL.decls, &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Type: yyDollar[1].typ})
			}
		}
	case 143:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1112
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.typ = yylex.(*lexer).pushType(&Type{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Kind: yyDollar[1].tk, Tag: yyDollar[2].str})
		}
	case 144:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line cc.y:1117
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[5].span)
			yyVAL.typ = yylex.(*lexer).pushType(&Type{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Kind: yyDollar[1].tk, Tag: yyDollar[2].str, Decls: yyDollar[4].decls})
		}
	case 145:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1124
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.prefix = &Prefix{Span: yyVAL.span, Dot: yyDollar[2].str}
		}
	case 146:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1131
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Arrow, Left: yyDollar[1].expr, Text: yyDollar[3].str}
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1136
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Dot, Left: yyDollar[1].expr, Text: yyDollar[3].str}
		}
	case 148:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1144
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.typ = yylex.(*lexer).pushType(&Type{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Kind: Enum, Tag: yyDollar[2].str})
		}
	case 149:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line cc.y:1149
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[6].span)
			yyVAL.typ = yylex.(*lexer).pushType(&Type{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Kind: Enum, Tag: yyDollar[2].str, Decls: yyDollar[4].decls})
		}
	case 150:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1156
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			var x *Init
			if yyDollar[2].expr != nil {
				x = &Init{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Expr: yyDollar[2].expr}
			}
			yyVAL.decl = &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Name: yyDollar[1].str, Init: x}
			yylex.(*lexer).pushDecl(yyVAL.decl)
		}
	case 151:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1168
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = yyDollar[2].expr
		}
	case 152:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1176
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.init = &Init{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Expr: yyDollar[1].expr}
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1181
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.init = &Init{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Braced: yyDollar[1].inits}
		}
	case 154:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1188
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.inits = []*Init{}
		}
	case 155:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line cc.y:1193
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[4].span)
			yyVAL.inits = append(yyDollar[2].inits, yyDollar[3].init)
		}
	case 156:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line cc.y:1198
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[5].span)
			yyVAL.inits = append(yyDollar[2].inits, yyDollar[3].init)
		}
	case 157:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1204
		{
			yyVAL.span = Span{}
			yyVAL.inits = nil
		}
	case 158:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1209
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.inits = append(yyDollar[1].inits, yyDollar[2].init)
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1216
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.init = yyDollar[1].init
		}
	case 160:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1221
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.init = yyDollar[3].init
			yyVAL.init.Prefix = yyDollar[1].prefixes
		}
	case 161:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1229
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.prefix = &Prefix{Span: yyVAL.span, Index: yyDollar[2].expr}
		}
	case 162:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1235
		{
			yyVAL.span = Span{}
		}
	case 163:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1239
		{
			yyVAL.span = yyDollar[1].span
		}
	case 164:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1244
		{
			yyVAL.span = Span{}
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1248
		{
			yyVAL.span = yyDollar[1].span
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1257
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.prefixes = []*Prefix{yyDollar[1].prefix}
		}
	case 167:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1262
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.prefixes = append(yyDollar[1].prefixes, yyDollar[2].prefix)
		}
	case 168:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1268
		{
			yyVAL.span = Span{}
			yyVAL.str = ""
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1273
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 170:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1279
		{
			yyVAL.span = Span{}
			yyVAL.expr = nil
		}
	case 171:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1284
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.expr = yyDollar[1].expr
		}
	case 172:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1290
		{
			yyVAL.span = Span{}
			yyVAL.expr = nil
		}
	case 173:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1295
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.expr = yyDollar[1].expr
		}
	case 174:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1302
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.exprs = []*Expr{yyDollar[1].expr}
		}
	case 175:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1307
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[3].expr)
		}
	case 176:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1313
		{
			yyVAL.span = Span{}
			yyVAL.exprs = nil
		}
	case 177:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1318
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.exprs = yyDollar[1].exprs
		}
	case 178:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1324
		{
			yyVAL.span = Span{}
			yyVAL.decls = nil
		}
	case 179:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1329
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.decls = append(yyDollar[1].decls, yyDollar[2].decls...)
		}
	case 180:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1335
		{
			yyVAL.span = Span{}
			yyVAL.labels = nil
		}
	case 181:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1340
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.labels = append(yyDollar[1].labels, yyDollar[2].label)
		}
	case 182:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1347
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.decls = []*Decl{yyDollar[1].decl}
		}
	case 183:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1352
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.decls = append(yyDollar[1].decls, yyDollar[3].decl)
		}
	case 184:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1358
		{
			yyVAL.span = Span{}
			yyVAL.decls = nil
		}
	case 185:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1363
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.decls = yyDollar[1].decls
		}
	case 186:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1370
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.idecs = []idecor{yyDollar[1].idec}
		}
	case 187:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1375
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.idecs = append(yyDollar[1].idecs, yyDollar[3].idec)
		}
	case 188:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1381
		{
			yyVAL.span = Span{}
			yyVAL.idecs = nil
		}
	case 189:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1386
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.idecs = yyDollar[1].idecs
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1393
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.strs = []string{yyDollar[1].str}
		}
	case 191:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1398
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.strs = append(yyDollar[1].strs, yyDollar[2].str)
		}
	case 192:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1404
		{
			yyVAL.span = Span{}
			yyVAL.strs = nil
		}
	case 193:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1409
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.strs = yyDollar[1].strs
		}
	case 194:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1416
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.strs = []string{yyDollar[1].str}
		}
	case 195:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1421
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.strs = append(yyDollar[1].strs, yyDollar[2].str)
		}
	case 196:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1427
		{
			yyVAL.span = Span{}
			yyVAL.strs = nil
		}
	case 197:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1432
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.strs = yyDollar[1].strs
		}
	case 198:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1439
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.strs = []string{yyDollar[1].str}
		}
	case 199:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1444
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.strs = append(yyDollar[1].strs, yyDollar[2].str)
		}
	case 200:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1450
		{
			yyVAL.span = Span{}
			yyVAL.strs = nil
		}
	case 201:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1455
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.strs = yyDollar[1].strs
		}
	case 202:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1462
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.decors = nil
			yyVAL.decors = append(yyVAL.decors, yyDollar[1].decor)
		}
	case 203:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1468
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.decors = append(yyDollar[1].decors, yyDollar[3].decor)
		}
	case 204:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1474
		{
			yyVAL.span = Span{}
			yyVAL.decors = nil
		}
	case 205:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1479
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.decors = yyDollar[1].decors
		}
	case 206:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1486
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.decls = yyDollar[1].decls
		}
	case 207:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1491
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.decls = append(yyDollar[1].decls, yyDollar[2].decls...)
		}
	case 208:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1497
		{
			yyVAL.span = Span{}
			yyVAL.expr = nil
		}
	case 209:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1502
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.expr = yyDollar[1].expr
		}
	case 210:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1509
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.decls = []*Decl{yyDollar[1].decl}
		}
	case 211:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1514
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.decls = append(yyDollar[1].decls, yyDollar[3].decl)
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1521
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.strs = []string{yyDollar[1].str}
		}
	case 213:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1526
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.strs = append(yyDollar[1].strs, yyDollar[2].str)
		}
	}
	goto yystack /* stack new state and value */
}
