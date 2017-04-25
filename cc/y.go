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

const tokARGBEGIN = 57346
const tokARGEND = 57347
const tokAUTOLIB = 57348
const tokSET = 57349
const tokUSED = 57350
const tokAuto = 57351
const tokBreak = 57352
const tokCase = 57353
const tokChar = 57354
const tokConst = 57355
const tokContinue = 57356
const tokDefault = 57357
const tokDo = 57358
const tokDotDotDot = 57359
const tokDouble = 57360
const tokEnum = 57361
const tokError = 57362
const tokExtern = 57363
const tokFloat = 57364
const tokFor = 57365
const tokGoto = 57366
const tokIf = 57367
const tokInline = 57368
const tokInt = 57369
const tokLitChar = 57370
const tokLong = 57371
const tokName = 57372
const tokNumber = 57373
const tokOffsetof = 57374
const tokRegister = 57375
const tokReturn = 57376
const tokShort = 57377
const tokSigned = 57378
const tokStatic = 57379
const tokStruct = 57380
const tokSwitch = 57381
const tokTypeName = 57382
const tokTypedef = 57383
const tokUnion = 57384
const tokUnsigned = 57385
const tokVaArg = 57386
const tokVoid = 57387
const tokVolatile = 57388
const tokWhile = 57389
const tokString = 57390
const tokDefine = 57391
const tokShift = 57392
const tokElse = 57393
const tokAddEq = 57394
const tokSubEq = 57395
const tokMulEq = 57396
const tokDivEq = 57397
const tokModEq = 57398
const tokLshEq = 57399
const tokRshEq = 57400
const tokAndEq = 57401
const tokXorEq = 57402
const tokOrEq = 57403
const tokOrOr = 57404
const tokAndAnd = 57405
const tokEqEq = 57406
const tokNotEq = 57407
const tokLtEq = 57408
const tokGtEq = 57409
const tokLsh = 57410
const tokRsh = 57411
const tokCast = 57412
const tokSizeof = 57413
const tokUnary = 57414
const tokDec = 57415
const tokInc = 57416
const tokArrow = 57417
const startExpr = 57418
const startProg = 57419
const tokEOF = 57420

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"tokARGBEGIN",
	"tokARGEND",
	"tokAUTOLIB",
	"tokSET",
	"tokUSED",
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
	-1, 120,
	53, 100,
	102, 100,
	-2, 182,
	-1, 139,
	52, 173,
	-2, 147,
	-1, 141,
	52, 173,
	-2, 152,
	-1, 242,
	102, 208,
	-2, 172,
	-1, 273,
	66, 173,
	-2, 91,
}

const yyNprod = 218
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1424

var yyAct = [...]int{

	313, 7, 114, 122, 346, 306, 263, 32, 229, 199,
	270, 217, 283, 244, 100, 101, 102, 103, 104, 105,
	106, 107, 108, 113, 51, 196, 223, 347, 241, 111,
	119, 5, 221, 125, 4, 227, 312, 231, 136, 134,
	381, 132, 120, 139, 141, 379, 372, 34, 371, 367,
	112, 360, 358, 341, 340, 338, 300, 293, 292, 190,
	309, 296, 37, 248, 143, 144, 145, 146, 147, 148,
	149, 150, 151, 152, 153, 154, 155, 156, 157, 158,
	159, 160, 161, 133, 163, 164, 165, 166, 167, 168,
	169, 170, 171, 172, 173, 61, 193, 131, 384, 137,
	214, 177, 178, 253, 299, 98, 94, 36, 93, 162,
	96, 95, 97, 378, 3, 2, 370, 192, 187, 191,
	239, 126, 176, 369, 368, 183, 365, 364, 288, 287,
	200, 127, 192, 192, 191, 191, 126, 112, 256, 179,
	180, 219, 130, 209, 138, 207, 127, 182, 198, 181,
	81, 6, 80, 79, 78, 77, 76, 74, 75, 70,
	71, 72, 73, 68, 69, 63, 64, 65, 66, 67,
	202, 201, 366, 123, 203, 98, 94, 133, 93, 184,
	96, 95, 97, 211, 124, 349, 348, 345, 260, 343,
	186, 337, 336, 137, 228, 230, 131, 234, 137, 261,
	216, 117, 116, 110, 286, 236, 237, 262, 246, 198,
	215, 211, 247, 208, 214, 225, 228, 303, 352, 212,
	351, 295, 238, 32, 220, 206, 278, 251, 242, 235,
	264, 233, 294, 275, 257, 259, 210, 258, 138, 62,
	225, 195, 205, 138, 236, 175, 273, 212, 254, 252,
	250, 230, 271, 284, 285, 204, 189, 380, 281, 65,
	66, 67, 118, 99, 242, 265, 356, 98, 94, 267,
	93, 272, 96, 95, 97, 59, 245, 188, 35, 128,
	298, 225, 126, 301, 289, 198, 290, 305, 304, 297,
	274, 232, 127, 291, 1, 302, 39, 308, 273, 11,
	237, 197, 251, 230, 271, 307, 135, 50, 60, 234,
	310, 140, 142, 129, 316, 282, 315, 317, 249, 280,
	121, 174, 321, 276, 277, 268, 269, 342, 243, 339,
	240, 28, 344, 26, 222, 350, 63, 64, 65, 66,
	67, 194, 234, 322, 31, 29, 98, 94, 357, 93,
	185, 96, 95, 97, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 353, 354, 0, 0, 0, 375,
	376, 377, 374, 359, 0, 0, 361, 362, 0, 0,
	0, 383, 0, 0, 382, 385, 0, 0, 0, 0,
	0, 0, 0, 0, 373, 82, 83, 84, 85, 86,
	87, 88, 89, 90, 91, 92, 81, 363, 80, 79,
	78, 77, 76, 74, 75, 70, 71, 72, 73, 68,
	69, 63, 64, 65, 66, 67, 0, 0, 0, 0,
	0, 98, 94, 0, 93, 0, 96, 95, 97, 323,
	54, 0, 320, 319, 59, 324, 333, 0, 0, 325,
	334, 326, 115, 0, 0, 0, 0, 58, 327, 328,
	329, 0, 0, 10, 57, 335, 9, 21, 55, 330,
	0, 0, 56, 0, 331, 0, 0, 60, 0, 23,
	0, 0, 332, 24, 0, 0, 0, 264, 76, 74,
	75, 70, 71, 72, 73, 68, 69, 63, 64, 65,
	66, 67, 0, 0, 0, 0, 13, 98, 94, 0,
	93, 0, 96, 95, 97, 14, 15, 12, 0, 0,
	0, 16, 17, 20, 0, 0, 0, 0, 22, 0,
	19, 18, 0, 0, 0, 0, 0, 318, 82, 83,
	84, 85, 86, 87, 88, 89, 90, 91, 92, 81,
	0, 80, 79, 78, 77, 76, 74, 75, 70, 71,
	72, 73, 68, 69, 63, 64, 65, 66, 67, 0,
	0, 0, 0, 0, 98, 94, 311, 93, 0, 96,
	95, 97, 82, 83, 84, 85, 86, 87, 88, 89,
	90, 91, 92, 81, 0, 80, 79, 78, 77, 76,
	74, 75, 70, 71, 72, 73, 68, 69, 63, 64,
	65, 66, 67, 0, 0, 0, 0, 0, 98, 94,
	0, 93, 279, 96, 95, 97, 218, 82, 83, 84,
	85, 86, 87, 88, 89, 90, 91, 92, 81, 0,
	80, 79, 78, 77, 76, 74, 75, 70, 71, 72,
	73, 68, 69, 63, 64, 65, 66, 67, 0, 0,
	0, 0, 0, 98, 94, 0, 93, 0, 96, 95,
	97, 82, 83, 84, 85, 86, 87, 88, 89, 90,
	91, 92, 81, 0, 80, 79, 78, 77, 76, 74,
	75, 70, 71, 72, 73, 68, 69, 63, 64, 65,
	66, 67, 0, 0, 0, 0, 0, 98, 94, 0,
	93, 0, 96, 95, 97, 54, 0, 0, 41, 59,
	0, 0, 0, 0, 48, 40, 0, 115, 47, 0,
	0, 0, 58, 43, 10, 44, 8, 9, 21, 57,
	0, 42, 45, 55, 52, 0, 38, 56, 53, 46,
	23, 49, 60, 0, 24, 0, 79, 78, 77, 76,
	74, 75, 70, 71, 72, 73, 68, 69, 63, 64,
	65, 66, 67, 0, 0, 0, 0, 13, 98, 94,
	0, 93, 0, 96, 95, 97, 14, 15, 12, 0,
	0, 0, 16, 17, 20, 0, 0, 27, 0, 22,
	54, 19, 18, 41, 59, 0, 0, 0, 0, 48,
	40, 0, 30, 47, 0, 0, 0, 58, 43, 0,
	44, 0, 0, 0, 57, 0, 42, 45, 55, 52,
	0, 38, 56, 53, 46, 27, 49, 60, 54, 0,
	33, 41, 59, 0, 0, 0, 0, 48, 40, 0,
	30, 47, 0, 0, 0, 58, 43, 0, 44, 0,
	0, 0, 57, 0, 42, 45, 55, 52, 0, 38,
	56, 53, 46, 54, 49, 60, 41, 59, 33, 0,
	0, 0, 48, 40, 0, 115, 47, 0, 0, 0,
	58, 43, 255, 44, 0, 0, 0, 57, 0, 42,
	45, 55, 52, 0, 38, 56, 53, 46, 54, 49,
	60, 41, 59, 0, 0, 0, 0, 48, 40, 0,
	115, 47, 0, 0, 0, 58, 43, 0, 44, 25,
	0, 0, 57, 0, 42, 45, 55, 52, 0, 38,
	56, 53, 46, 0, 49, 60, 78, 77, 76, 74,
	75, 70, 71, 72, 73, 68, 69, 63, 64, 65,
	66, 67, 0, 0, 0, 314, 0, 98, 94, 0,
	93, 0, 96, 95, 97, 0, 77, 76, 74, 75,
	70, 71, 72, 73, 68, 69, 63, 64, 65, 66,
	67, 0, 0, 0, 0, 0, 98, 94, 0, 93,
	266, 96, 95, 97, 74, 75, 70, 71, 72, 73,
	68, 69, 63, 64, 65, 66, 67, 10, 0, 8,
	9, 21, 98, 94, 0, 93, 0, 96, 95, 97,
	0, 0, 0, 23, 0, 0, 0, 24, 0, 0,
	0, 213, 0, 0, 75, 70, 71, 72, 73, 68,
	69, 63, 64, 65, 66, 67, 0, 0, 0, 0,
	13, 98, 94, 0, 93, 0, 96, 95, 97, 14,
	15, 12, 0, 0, 0, 16, 17, 20, 0, 284,
	285, 0, 22, 0, 19, 18, 70, 71, 72, 73,
	68, 69, 63, 64, 65, 66, 67, 10, 0, 8,
	9, 21, 98, 94, 0, 93, 0, 96, 95, 97,
	0, 0, 0, 23, 0, 0, 10, 24, 8, 9,
	21, 213, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 23, 0, 0, 0, 24, 0, 0, 0,
	13, 0, 0, 0, 0, 0, 0, 0, 0, 14,
	15, 12, 0, 0, 0, 16, 17, 20, 0, 13,
	0, 0, 22, 0, 19, 18, 0, 0, 14, 15,
	12, 0, 0, 0, 16, 17, 20, 0, 0, 0,
	0, 22, 0, 19, 18, 10, 0, 8, 9, 21,
	0, 68, 69, 63, 64, 65, 66, 67, 0, 0,
	0, 23, 0, 98, 94, 24, 93, 0, 96, 95,
	97, 10, 0, 8, 9, 21, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 23, 13, 0,
	0, 24, 0, 0, 0, 213, 0, 14, 15, 12,
	0, 0, 0, 16, 17, 20, 0, 0, 0, 0,
	109, 0, 19, 18, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 16,
	17, 20, 0, 0, 0, 0, 22, 54, 19, 18,
	41, 59, 0, 0, 0, 226, 48, 40, 0, 115,
	47, 0, 0, 0, 58, 43, 0, 44, 224, 0,
	0, 57, 0, 42, 45, 55, 52, 0, 38, 56,
	53, 46, 355, 49, 60, 0, 54, 0, 0, 41,
	59, 0, 0, 0, 0, 48, 40, 0, 115, 47,
	0, 0, 0, 58, 43, 0, 44, 0, 0, 0,
	57, 0, 42, 45, 55, 52, 0, 38, 56, 53,
	46, 54, 49, 60, 41, 59, 0, 0, 0, 0,
	48, 40, 0, 115, 47, 0, 0, 0, 58, 43,
	0, 44, 0, 0, 0, 57, 0, 42, 45, 55,
	52, 0, 38, 56, 53, 46, 54, 49, 60, 41,
	59, 0, 0, 0, 0, 48, 0, 0, 115, 47,
	0, 0, 0, 58, 43, 0, 44, 0, 0, 0,
	57, 0, 42, 45, 55, 0, 0, 0, 56, 0,
	46, 0, 49, 60,
}
var yyPact = [...]int{

	16, -1000, -1000, 1088, 829, -5, 186, 617, -1000, -1000,
	-1000, 215, 1088, 1088, 1088, 1088, 1088, 1088, 1088, 1088,
	1157, 110, 706, 109, -1000, -1000, -1000, 108, -1000, -1000,
	214, -1000, 91, 249, 1342, 431, 1377, -1000, -1000, 252,
	252, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 1088, 1088, 1088, 1088, 1088, 1088, 1088, 1088,
	1088, 1088, 1088, 1088, 1088, 1088, 1088, 1088, 1088, 1088,
	1088, 1088, 1088, 1088, 1088, 1088, 1088, 1088, 1088, 1088,
	1088, 1088, 1088, 1088, 1088, -1000, -1000, 252, 252, -1000,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 706,
	1342, 55, 53, 97, -1000, -1000, 1088, 247, 204, -43,
	42, 188, -1000, 262, 91, -1000, -1000, -1000, 1088, 431,
	1377, -1000, -1000, 431, -1000, 1377, -1000, -1000, -1000, -1000,
	203, -1000, 190, 617, 177, 177, 15, 15, 15, 256,
	256, 1113, 1113, 1113, 1113, 971, 1012, 932, 417, 906,
	877, 688, 159, 617, 617, 617, 617, 617, 617, 617,
	617, 617, 617, 617, 51, 186, 121, -1000, -1000, 49,
	183, 1069, -1000, 123, 262, 107, 97, 573, 47, -1000,
	-1000, 1268, 1088, 1069, 1342, 91, 91, 262, -1000, 26,
	617, -1000, -1000, -1000, 1342, 246, 1088, -1000, -1000, 1183,
	1088, 15, -1000, -38, 1088, 97, 1268, 9, 1342, -1000,
	791, 44, 181, -1000, -1000, 106, -1000, 115, 617, -1000,
	617, -1000, 178, -1000, 91, -1000, 42, 41, -1000, -1000,
	899, -1000, 91, 180, -1000, 172, 85, 528, -1000, 989,
	112, 123, 35, -1000, 34, -1000, -1000, 1268, 123, 41,
	262, 106, -1000, -1000, -1000, -44, -1000, -1000, -45, 179,
	-1000, 41, 155, -1000, -40, 246, -1000, -1000, 1088, -1000,
	3, -1000, 163, -1000, 252, 1088, -1000, -1000, -1000, -1000,
	106, -1000, -1000, -1000, 91, 1088, -1000, -1000, 617, -1000,
	-41, 1069, -1000, -1000, -1000, 484, 864, -1000, 617, -1000,
	-1000, -1000, -1000, -1000, -1000, 435, -1000, -1000, -1000, 99,
	98, -1000, -47, -1000, -48, -49, -1000, 96, 252, 94,
	1088, 93, 92, 1088, 154, 152, 1088, 1088, -1000, 1307,
	-1000, -1000, 219, 1088, -50, 1088, -51, -1000, 1088, 1088,
	341, -1000, -1000, 33, 32, -1000, 79, -53, -1000, 30,
	-1000, 29, 22, -1000, -54, -56, 1088, 1088, -1000, -1000,
	-1000, -1000, -1000, 19, -57, 206, -1000, -1000, -62, 1088,
	-1000, -1000, 4, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 11, 350, 26, 345, 13, 344, 36, 341, 334,
	32, 34, 333, 331, 28, 330, 328, 9, 10, 326,
	325, 1, 35, 27, 4, 324, 323, 151, 321, 33,
	320, 30, 8, 319, 37, 318, 317, 316, 12, 315,
	314, 6, 0, 5, 307, 24, 107, 62, 38, 3,
	271, 47, 41, 306, 39, 301, 25, 299, 2, 296,
	29, 23, 278, 294, 293, 291, 290, 283,
}
var yyR1 = [...]int{

	0, 63, 63, 11, 11, 11, 23, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	43, 43, 43, 64, 41, 36, 36, 36, 42, 40,
	40, 40, 40, 40, 40, 40, 40, 40, 40, 40,
	40, 40, 40, 40, 40, 1, 1, 1, 2, 2,
	2, 17, 17, 17, 17, 17, 3, 3, 3, 3,
	29, 29, 44, 44, 44, 44, 44, 44, 45, 45,
	46, 46, 46, 46, 46, 46, 46, 46, 46, 47,
	47, 48, 48, 62, 58, 58, 58, 58, 58, 61,
	60, 7, 13, 6, 12, 12, 12, 12, 65, 4,
	49, 49, 59, 59, 18, 18, 14, 62, 62, 38,
	21, 21, 62, 62, 5, 25, 32, 32, 34, 34,
	34, 35, 35, 33, 33, 38, 67, 67, 66, 66,
	39, 39, 50, 50, 24, 24, 22, 22, 27, 27,
	28, 28, 8, 8, 37, 37, 9, 9, 10, 10,
	30, 30, 31, 31, 55, 55, 56, 56, 51, 51,
	52, 52, 53, 53, 54, 54, 19, 19, 20, 20,
	15, 15, 26, 26, 16, 16, 57, 57,
}
var yyR2 = [...]int{

	0, 3, 3, 0, 2, 5, 1, 1, 1, 1,
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 5,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	4, 6, 4, 4, 3, 4, 4, 2, 2, 6,
	0, 2, 2, 0, 4, 3, 2, 2, 2, 1,
	5, 5, 1, 2, 3, 2, 2, 7, 9, 3,
	5, 7, 3, 5, 5, 0, 3, 1, 4, 4,
	3, 1, 3, 3, 4, 4, 1, 2, 2, 1,
	1, 3, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 3, 3, 2, 2, 1,
	2, 3, 3, 3, 1, 1, 5, 1, 0, 5,
	1, 1, 1, 1, 1, 3, 3, 2, 5, 2,
	3, 3, 2, 6, 2, 2, 1, 1, 2, 4,
	5, 0, 3, 1, 3, 3, 0, 1, 0, 1,
	1, 2, 0, 1, 0, 1, 0, 1, 1, 3,
	0, 1, 0, 2, 0, 2, 1, 3, 0, 1,
	1, 3, 0, 1, 1, 2, 0, 1, 1, 2,
	0, 1, 1, 2, 0, 1, 1, 3, 0, 1,
	1, 2, 0, 1, 1, 3, 1, 2,
}
var yyChk = [...]int{

	-1000, -63, 99, 98, -11, -23, -27, -21, 30, 31,
	28, -57, 82, 71, 80, 81, 86, 87, 96, 95,
	88, 32, 93, 44, 48, 100, -12, 6, -13, -4,
	21, -6, -58, 49, -51, -62, -46, -47, 40, -59,
	19, 12, 35, 27, 29, 36, 43, 22, 18, 45,
	-44, -45, 38, 42, 9, 37, 41, 33, 26, 13,
	46, 100, 53, 80, 81, 82, 83, 84, 78, 79,
	74, 75, 76, 77, 72, 73, 71, 70, 69, 68,
	67, 65, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 93, 91, 96, 95, 97, 90, 48,
	-21, -21, -21, -21, -21, -21, -21, -21, -21, 93,
	93, -60, -23, -61, -58, 21, 93, 93, 48, -31,
	-17, -30, -49, 82, 93, -29, 30, 40, 30, -62,
	-46, -47, -52, -51, -54, -53, -48, -47, -46, -49,
	-50, -49, -50, -21, -21, -21, -21, -21, -21, -21,
	-21, -21, -21, -21, -21, -21, -21, -21, -21, -21,
	-21, -21, -23, -21, -21, -21, -21, -21, -21, -21,
	-21, -21, -21, -21, -28, -27, -23, -49, -49, -60,
	-60, 94, 94, -1, 82, -2, 93, -21, 30, 52,
	102, 93, 91, 54, -8, 53, -56, -55, -45, -17,
	-21, -52, -54, -48, 52, 52, 66, 94, 92, 94,
	53, -21, -34, 52, 91, -56, 93, -1, 53, 94,
	-11, -10, -9, -3, 30, -61, 17, -22, -21, -32,
	-21, -34, -65, -7, -58, -29, -17, -17, -45, 94,
	-15, -14, -61, -16, -5, 30, -21, -21, 101, -35,
	-22, -1, -10, 94, -60, 101, 94, 53, -1, -17,
	82, 93, 92, -41, 52, -31, 101, -14, -20, -19,
	-18, -17, -50, -49, -66, 53, -26, -25, 54, 94,
	-33, -32, -39, -38, 90, 91, 92, 94, 94, -3,
	-56, -64, 102, 102, 53, 66, 101, -5, -21, 101,
	53, -67, -38, 54, -49, -21, -43, -18, -21, 101,
	-32, 92, -7, -42, 101, -37, -40, -36, 102, 8,
	7, -41, -23, 4, 10, 14, 16, 23, 24, 25,
	34, 39, 47, 11, 15, 30, 93, 93, 102, -43,
	102, 102, -42, 93, -49, 93, -24, -23, 93, 93,
	-21, 66, 66, -23, -23, 5, 47, -24, 102, -23,
	102, -23, -23, 66, 94, 94, 93, 102, 94, 94,
	94, 102, 102, -23, -24, -42, -42, -42, 94, 102,
	51, 102, -24, -42, 94, -42,
}
var yyDef = [...]int{

	0, -2, 3, 0, 0, 0, 6, 178, 7, 8,
	9, 10, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 216, 1, 4, 0, 134, 135,
	104, 137, 192, 0, 124, 200, 204, 198, 123, 172,
	172, 110, 111, 112, 113, 114, 115, 116, 117, 118,
	119, 120, 142, 143, 102, 103, 105, 106, 107, 108,
	109, 2, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 180, 0, 57, 58, 0, 0, 217,
	41, 42, 43, 44, 45, 46, 47, 48, 49, 0,
	0, 0, 0, 85, 129, 104, 0, 0, 0, 0,
	-2, 193, 91, 196, 0, 190, 140, 141, 0, 200,
	204, 199, 127, 201, 128, 205, 202, 121, 122, -2,
	0, -2, 0, 179, 11, 12, 13, 14, 15, 16,
	17, 18, 19, 20, 21, 22, 23, 24, 25, 26,
	27, 28, 0, 30, 31, 32, 33, 34, 35, 36,
	37, 38, 39, 40, 0, 181, 0, 150, 151, 0,
	0, 0, 54, 130, 196, 87, 85, 0, 0, 3,
	132, 188, 176, 0, 138, 0, 0, 197, 194, 0,
	133, 125, 126, 203, 0, 0, 0, 55, 56, 50,
	0, 52, 53, 161, 176, 85, 188, 0, 0, 5,
	0, 0, 189, 186, 96, 85, 99, 0, 177, 101,
	156, 157, 0, 183, 192, 191, 100, 92, 195, 93,
	0, 210, -2, 168, 214, 212, 29, 0, 158, 0,
	0, 86, 0, 90, 0, 136, 94, 0, 97, 98,
	196, 85, 95, 139, 63, 0, 148, 211, 0, 209,
	206, 144, 0, -2, 0, 169, 154, 213, 0, 51,
	0, 163, 166, 170, 0, 0, 89, 88, 59, 187,
	85, 60, 131, 146, 172, 0, 153, 215, 155, 159,
	162, 0, 171, 167, 149, 0, 184, 207, 145, 160,
	164, 165, 61, 62, 64, 0, 68, 185, 69, 0,
	0, 72, 0, 60, 0, 0, 184, 0, 0, 0,
	174, 0, 0, 0, 0, 7, 0, 0, 73, 184,
	75, 76, 0, 174, 0, 0, 0, 175, 0, 0,
	0, 66, 67, 0, 0, 74, 0, 0, 79, 0,
	82, 0, 0, 65, 0, 0, 0, 174, 184, 184,
	184, 70, 71, 0, 0, 80, 83, 84, 0, 174,
	184, 77, 0, 81, 184, 78,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 86, 3, 3, 3, 84, 71, 3,
	93, 94, 82, 80, 53, 81, 90, 83, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 66, 102,
	74, 54, 75, 65, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 91, 3, 92, 70, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 52, 69, 101, 87,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	55, 56, 57, 58, 59, 60, 61, 62, 63, 64,
	67, 68, 72, 73, 76, 77, 78, 79, 85, 88,
	89, 95, 96, 97, 98, 99, 100,
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
		//line cc.y:186
		{
			yylex.(*lexer).prog = &Prog{Decls: yyDollar[2].decls}
			return 0
		}
	case 2:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:191
		{
			yylex.(*lexer).expr = yyDollar[2].expr
			return 0
		}
	case 3:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:197
		{
			yyVAL.span = Span{}
			yyVAL.decls = nil
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:202
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.decls = append(yyDollar[1].decls, yyDollar[2].decls...)
		}
	case 5:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line cc.y:207
		{
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:212
		{
			yyVAL.span = yyDollar[1].span
			if len(yyDollar[1].exprs) == 1 {
				yyVAL.expr = yyDollar[1].exprs[0]
				break
			}
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Comma, List: yyDollar[1].exprs}
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:223
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Name, Text: yyDollar[1].str, XDecl: yyDollar[1].decl}
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:228
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Number, Text: yyDollar[1].str}
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:233
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Number, Text: yyDollar[1].str}
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:238
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: String, Texts: yyDollar[1].strs}
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:243
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Add, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:248
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Sub, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:253
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Mul, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:258
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Div, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:263
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Mod, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:268
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Lsh, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:273
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Rsh, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:278
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Lt, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:283
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Gt, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:288
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: LtEq, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:293
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: GtEq, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:298
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: EqEq, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:303
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: NotEq, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:308
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: And, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:313
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Xor, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:318
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Or, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:323
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: AndAnd, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:328
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: OrOr, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 29:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line cc.y:333
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[5].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Cond, List: []*Expr{yyDollar[1].expr, yyDollar[3].expr, yyDollar[5].expr}}
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:338
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Eq, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:343
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: AddEq, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:348
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: SubEq, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:353
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: MulEq, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:358
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: DivEq, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:363
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: ModEq, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:368
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: LshEq, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:373
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: RshEq, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:378
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: AndEq, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:383
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: XorEq, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:388
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: OrEq, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:393
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Indir, Left: yyDollar[2].expr}
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:398
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Addr, Left: yyDollar[2].expr}
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:403
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Plus, Left: yyDollar[2].expr}
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:408
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Minus, Left: yyDollar[2].expr}
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:413
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Not, Left: yyDollar[2].expr}
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:418
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Twid, Left: yyDollar[2].expr}
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:423
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: PreInc, Left: yyDollar[2].expr}
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:428
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: PreDec, Left: yyDollar[2].expr}
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:433
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: SizeofExpr, Left: yyDollar[2].expr}
		}
	case 50:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line cc.y:438
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[4].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: SizeofType, Type: yyDollar[3].typ}
		}
	case 51:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line cc.y:443
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[6].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Offsetof, Type: yyDollar[3].typ, Left: yyDollar[5].expr}
		}
	case 52:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line cc.y:448
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[4].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Cast, Type: yyDollar[2].typ, Left: yyDollar[4].expr}
		}
	case 53:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line cc.y:453
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[4].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: CastInit, Type: yyDollar[2].typ, Init: &Init{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Braced: yyDollar[4].inits}}
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:458
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Paren, Left: yyDollar[2].expr}
		}
	case 55:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line cc.y:463
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[4].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Call, Left: yyDollar[1].expr, List: yyDollar[3].exprs}
		}
	case 56:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line cc.y:468
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[4].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Index, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:473
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: PostInc, Left: yyDollar[1].expr}
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:478
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: PostDec, Left: yyDollar[1].expr}
		}
	case 59:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line cc.y:483
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[6].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: VaArg, Left: yyDollar[3].expr, Type: yyDollar[5].typ}
		}
	case 60:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:489
		{
			yyVAL.span = Span{}
			yyVAL.stmts = nil
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:494
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.stmts = yyDollar[1].stmts
			for _, d := range yyDollar[2].decls {
				yyVAL.stmts = append(yyVAL.stmts, &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: StmtDecl, Decl: d})
			}
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:502
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[2].stmt)
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:509
		{
			yylex.(*lexer).pushScope()
		}
	case 64:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line cc.y:513
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[4].span)
			yylex.(*lexer).popScope()
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Block, Block: yyDollar[3].stmts}
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:521
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.label = &Label{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Case, Expr: yyDollar[2].expr}
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:526
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.label = &Label{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Default}
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:531
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.label = &Label{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: LabelName, Name: yyDollar[1].str}
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:538
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.stmt = yyDollar[2].stmt
			yyVAL.stmt.Labels = yyDollar[1].labels
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:546
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Empty}
		}
	case 70:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line cc.y:551
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Empty}
		}
	case 71:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line cc.y:556
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Empty}
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:561
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:566
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: StmtExpr, Expr: yyDollar[1].expr}
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:571
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: ARGBEGIN, Block: yyDollar[2].stmts}
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:576
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Break}
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:581
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Continue}
		}
	case 77:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line cc.y:586
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[7].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Do, Body: yyDollar[2].stmt, Expr: yyDollar[5].expr}
		}
	case 78:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line cc.y:591
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
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:602
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Goto, Text: yyDollar[2].str}
		}
	case 80:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line cc.y:607
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[5].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: If, Expr: yyDollar[3].expr, Body: yyDollar[5].stmt}
		}
	case 81:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line cc.y:612
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[7].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: If, Expr: yyDollar[3].expr, Body: yyDollar[5].stmt, Else: yyDollar[7].stmt}
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:617
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Return, Expr: yyDollar[2].expr}
		}
	case 83:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line cc.y:622
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[5].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Switch, Expr: yyDollar[3].expr, Body: yyDollar[5].stmt}
		}
	case 84:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line cc.y:627
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[5].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: While, Expr: yyDollar[3].expr, Body: yyDollar[5].stmt}
		}
	case 85:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:634
		{
			yyVAL.span = Span{}
			yyVAL.abdecor = func(t *Type) *Type { return t }
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:639
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			_, q, _ := splitTypeWords(yyDollar[2].strs)
			abdecor := yyDollar[3].abdecor
			yyVAL.abdecor = func(t *Type) *Type {
				return abdecor(&Type{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Kind: Ptr, Base: t, Qual: q})
			}
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:648
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.abdecor = yyDollar[1].abdecor
		}
	case 88:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line cc.y:655
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
	case 89:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line cc.y:679
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[4].span)
			abdecor := yyDollar[1].abdecor
			span := yyVAL.span
			expr := yyDollar[3].expr
			yyVAL.abdecor = func(t *Type) *Type {
				return abdecor(&Type{SyntaxInfo: SyntaxInfo{Span: span}, Kind: Array, Base: t, Width: expr})
			}

		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:690
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.abdecor = yyDollar[2].abdecor
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:698
		{
			yyVAL.span = yyDollar[1].span
			name := yyDollar[1].str
			yyVAL.decor = func(t *Type) (*Type, string) { return t, name }
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:704
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			_, q, _ := splitTypeWords(yyDollar[2].strs)
			decor := yyDollar[3].decor
			span := yyVAL.span
			yyVAL.decor = func(t *Type) (*Type, string) {
				return decor(&Type{SyntaxInfo: SyntaxInfo{Span: span}, Kind: Ptr, Base: t, Qual: q})
			}
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:714
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.decor = yyDollar[2].decor
		}
	case 94:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line cc.y:719
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[4].span)
			decor := yyDollar[1].decor
			decls := yyDollar[3].decls
			span := yyVAL.span
			yyVAL.decor = func(t *Type) (*Type, string) {
				return decor(&Type{SyntaxInfo: SyntaxInfo{Span: span}, Kind: Func, Base: t, Decls: decls})
			}
		}
	case 95:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line cc.y:729
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[4].span)
			decor := yyDollar[1].decor
			span := yyVAL.span
			expr := yyDollar[3].expr
			yyVAL.decor = func(t *Type) (*Type, string) {
				return decor(&Type{SyntaxInfo: SyntaxInfo{Span: span}, Kind: Array, Base: t, Width: expr})
			}
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:742
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.decl = &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Name: yyDollar[1].str}
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:747
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.decl = &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Type: yyDollar[2].abdecor(yyDollar[1].typ)}
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:752
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			typ, name := yyDollar[2].decor(yyDollar[1].typ)
			yyVAL.decl = &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Name: name, Type: typ}
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:758
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.decl = &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Name: "..."}
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:766
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.idec = idecor{yyDollar[1].decor, nil}
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:771
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.idec = idecor{yyDollar[1].decor, yyDollar[3].init}
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:779
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:784
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:789
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:794
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:799
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:804
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:812
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:817
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:825
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 111:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:830
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:835
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:840
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:845
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:850
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:855
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
		//line cc.y:872
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:877
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:884
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:889
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:897
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.typ = yyDollar[1].typ
			if yyVAL.typ == nil {
				yyVAL.typ = &Type{Kind: TypedefType, Name: yyDollar[1].str}
			}
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:913
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.tc.c, yyVAL.tc.q, yyVAL.tc.t = splitTypeWords(append(yyDollar[1].strs, "int"))
		}
	case 125:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:918
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.tc.c, yyVAL.tc.q, _ = splitTypeWords(append(yyDollar[1].strs, yyDollar[3].strs...))
			yyVAL.tc.t = yyDollar[2].typ
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:924
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyDollar[1].strs = append(yyDollar[1].strs, yyDollar[2].str)
			yyDollar[1].strs = append(yyDollar[1].strs, yyDollar[3].strs...)
			yyVAL.tc.c, yyVAL.tc.q, yyVAL.tc.t = splitTypeWords(yyDollar[1].strs)
		}
	case 127:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:931
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.tc.c, yyVAL.tc.q, _ = splitTypeWords(yyDollar[2].strs)
			yyVAL.tc.t = yyDollar[1].typ
		}
	case 128:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:937
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			var ts []string
			ts = append(ts, yyDollar[1].str)
			ts = append(ts, yyDollar[2].strs...)
			yyVAL.tc.c, yyVAL.tc.q, yyVAL.tc.t = splitTypeWords(ts)
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:948
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
	case 130:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:961
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.typ = yyDollar[2].abdecor(yyDollar[1].typ)
		}
	case 131:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:969
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
	case 132:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:989
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
	case 133:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1017
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.decl = &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Name: yyDollar[2].str, Init: &Init{SyntaxInfo: SyntaxInfo{Span: yyDollar[3].span}, Expr: yyDollar[3].expr}}
			yylex.(*lexer).pushDecl(yyVAL.decl)
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1025
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.decls = yyDollar[1].decls
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1030
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.decls = []*Decl{yyDollar[1].decl}
		}
	case 136:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line cc.y:1035
		{
			yyVAL.decls = yyDollar[4].decls
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1039
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.decls = []*Decl{yyDollar[1].decl}
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1046
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
	case 139:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line cc.y:1067
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
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1080
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1085
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 142:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1093
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.tk = Struct
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1098
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.tk = Union
		}
	case 144:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1105
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.decor = yyDollar[1].decor
		}
	case 145:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1110
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			name := yyDollar[1].str
			expr := yyDollar[3].expr
			yyVAL.decor = func(t *Type) (*Type, string) {
				t.Width = expr
				return t, name
			}
		}
	case 146:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1122
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
	case 147:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1136
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.typ = yylex.(*lexer).pushType(&Type{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Kind: yyDollar[1].tk, Tag: yyDollar[2].str})
		}
	case 148:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line cc.y:1141
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[5].span)
			yyVAL.typ = yylex.(*lexer).pushType(&Type{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Kind: yyDollar[1].tk, Tag: yyDollar[2].str, Decls: yyDollar[4].decls})
		}
	case 149:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1148
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.prefix = &Prefix{Span: yyVAL.span, Dot: yyDollar[2].str}
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1155
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Arrow, Left: yyDollar[1].expr, Text: yyDollar[3].str}
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1160
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Dot, Left: yyDollar[1].expr, Text: yyDollar[3].str}
		}
	case 152:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1168
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.typ = yylex.(*lexer).pushType(&Type{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Kind: Enum, Tag: yyDollar[2].str})
		}
	case 153:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line cc.y:1173
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[6].span)
			yyVAL.typ = yylex.(*lexer).pushType(&Type{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Kind: Enum, Tag: yyDollar[2].str, Decls: yyDollar[4].decls})
		}
	case 154:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1180
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			var x *Init
			if yyDollar[2].expr != nil {
				x = &Init{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Expr: yyDollar[2].expr}
			}
			yyVAL.decl = &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Name: yyDollar[1].str, Init: x}
			yylex.(*lexer).pushDecl(yyVAL.decl)
		}
	case 155:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1192
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = yyDollar[2].expr
		}
	case 156:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1200
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.init = &Init{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Expr: yyDollar[1].expr}
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1205
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.init = &Init{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Braced: yyDollar[1].inits}
		}
	case 158:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1212
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.inits = []*Init{}
		}
	case 159:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line cc.y:1217
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[4].span)
			yyVAL.inits = append(yyDollar[2].inits, yyDollar[3].init)
		}
	case 160:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line cc.y:1222
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[5].span)
			yyVAL.inits = append(yyDollar[2].inits, yyDollar[3].init)
		}
	case 161:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1228
		{
			yyVAL.span = Span{}
			yyVAL.inits = nil
		}
	case 162:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1233
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.inits = append(yyDollar[1].inits, yyDollar[2].init)
		}
	case 163:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1240
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.init = yyDollar[1].init
		}
	case 164:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1245
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.init = yyDollar[3].init
			yyVAL.init.Prefix = yyDollar[1].prefixes
		}
	case 165:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1253
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.prefix = &Prefix{Span: yyVAL.span, Index: yyDollar[2].expr}
		}
	case 166:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1259
		{
			yyVAL.span = Span{}
		}
	case 167:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1263
		{
			yyVAL.span = yyDollar[1].span
		}
	case 168:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1268
		{
			yyVAL.span = Span{}
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1272
		{
			yyVAL.span = yyDollar[1].span
		}
	case 170:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1281
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.prefixes = []*Prefix{yyDollar[1].prefix}
		}
	case 171:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1286
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.prefixes = append(yyDollar[1].prefixes, yyDollar[2].prefix)
		}
	case 172:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1292
		{
			yyVAL.span = Span{}
			yyVAL.str = ""
		}
	case 173:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1297
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.str = yyDollar[1].str
		}
	case 174:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1303
		{
			yyVAL.span = Span{}
			yyVAL.expr = nil
		}
	case 175:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1308
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.expr = yyDollar[1].expr
		}
	case 176:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1314
		{
			yyVAL.span = Span{}
			yyVAL.expr = nil
		}
	case 177:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1319
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.expr = yyDollar[1].expr
		}
	case 178:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1326
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.exprs = []*Expr{yyDollar[1].expr}
		}
	case 179:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1331
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[3].expr)
		}
	case 180:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1337
		{
			yyVAL.span = Span{}
			yyVAL.exprs = nil
		}
	case 181:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1342
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.exprs = yyDollar[1].exprs
		}
	case 182:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1348
		{
			yyVAL.span = Span{}
			yyVAL.decls = nil
		}
	case 183:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1353
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.decls = append(yyDollar[1].decls, yyDollar[2].decls...)
		}
	case 184:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1359
		{
			yyVAL.span = Span{}
			yyVAL.labels = nil
		}
	case 185:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1364
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.labels = append(yyDollar[1].labels, yyDollar[2].label)
		}
	case 186:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1371
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.decls = []*Decl{yyDollar[1].decl}
		}
	case 187:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1376
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.decls = append(yyDollar[1].decls, yyDollar[3].decl)
		}
	case 188:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1382
		{
			yyVAL.span = Span{}
			yyVAL.decls = nil
		}
	case 189:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1387
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.decls = yyDollar[1].decls
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1394
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.idecs = []idecor{yyDollar[1].idec}
		}
	case 191:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1399
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.idecs = append(yyDollar[1].idecs, yyDollar[3].idec)
		}
	case 192:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1405
		{
			yyVAL.span = Span{}
			yyVAL.idecs = nil
		}
	case 193:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1410
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.idecs = yyDollar[1].idecs
		}
	case 194:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1417
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.strs = []string{yyDollar[1].str}
		}
	case 195:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1422
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.strs = append(yyDollar[1].strs, yyDollar[2].str)
		}
	case 196:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1428
		{
			yyVAL.span = Span{}
			yyVAL.strs = nil
		}
	case 197:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1433
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.strs = yyDollar[1].strs
		}
	case 198:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1440
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.strs = []string{yyDollar[1].str}
		}
	case 199:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1445
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.strs = append(yyDollar[1].strs, yyDollar[2].str)
		}
	case 200:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1451
		{
			yyVAL.span = Span{}
			yyVAL.strs = nil
		}
	case 201:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1456
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.strs = yyDollar[1].strs
		}
	case 202:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1463
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.strs = []string{yyDollar[1].str}
		}
	case 203:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1468
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.strs = append(yyDollar[1].strs, yyDollar[2].str)
		}
	case 204:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1474
		{
			yyVAL.span = Span{}
			yyVAL.strs = nil
		}
	case 205:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1479
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.strs = yyDollar[1].strs
		}
	case 206:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1486
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.decors = nil
			yyVAL.decors = append(yyVAL.decors, yyDollar[1].decor)
		}
	case 207:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1492
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.decors = append(yyDollar[1].decors, yyDollar[3].decor)
		}
	case 208:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1498
		{
			yyVAL.span = Span{}
			yyVAL.decors = nil
		}
	case 209:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1503
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.decors = yyDollar[1].decors
		}
	case 210:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1510
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.decls = yyDollar[1].decls
		}
	case 211:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1515
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.decls = append(yyDollar[1].decls, yyDollar[2].decls...)
		}
	case 212:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1521
		{
			yyVAL.span = Span{}
			yyVAL.expr = nil
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1526
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.expr = yyDollar[1].expr
		}
	case 214:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1533
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.decls = []*Decl{yyDollar[1].decl}
		}
	case 215:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1538
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.decls = append(yyDollar[1].decls, yyDollar[3].decl)
		}
	case 216:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1545
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.strs = []string{yyDollar[1].str}
		}
	case 217:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1550
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.strs = append(yyDollar[1].strs, yyDollar[2].str)
		}
	}
	goto yystack /* stack new state and value */
}
