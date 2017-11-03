//line parse.y:13
package build

import __yyfmt__ "fmt"

//line parse.y:13
//line parse.y:18
type yySymType struct {
	yys int
	// input tokens
	tok    string   // raw input syntax
	str    string   // decoding of quoted string
	pos    Position // position of token
	triple bool     // was string triple quoted?

	// partial syntax trees
	expr    Expr
	exprs   []Expr
	forc    *ForClause
	fors    []*ForClause
	ifs     []*IfClause
	string  *StringExpr
	strings []*StringExpr

	// supporting information
	comma    Position // position of trailing comma in list, if present
	lastRule Expr     // most recent rule, to attach line comments to
}

const _ADDEQ = 57346
const _AND = 57347
const _COMMENT = 57348
const _EOF = 57349
const _EQ = 57350
const _FOR = 57351
const _GE = 57352
const _IDENT = 57353
const _IF = 57354
const _ELSE = 57355
const _IN = 57356
const _IS = 57357
const _LAMBDA = 57358
const _LE = 57359
const _NE = 57360
const _NOT = 57361
const _OR = 57362
const _PYTHON = 57363
const _STRING = 57364
const ShiftInstead = 57365
const _ASSERT = 57366
const _UNARY = 57367

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"'%'",
	"'('",
	"')'",
	"'*'",
	"'+'",
	"','",
	"'-'",
	"'.'",
	"'/'",
	"':'",
	"'<'",
	"'='",
	"'>'",
	"'['",
	"']'",
	"'{'",
	"'}'",
	"_ADDEQ",
	"_AND",
	"_COMMENT",
	"_EOF",
	"_EQ",
	"_FOR",
	"_GE",
	"_IDENT",
	"_IF",
	"_ELSE",
	"_IN",
	"_IS",
	"_LAMBDA",
	"_LE",
	"_NE",
	"_NOT",
	"_OR",
	"_PYTHON",
	"_STRING",
	"ShiftInstead",
	"'\\n'",
	"_ASSERT",
	"_UNARY",
	"';'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parse.y:564

// Go helper code.

// unary returns a unary expression with the given
// position, operator, and subexpression.
func unary(pos Position, op string, x Expr) Expr {
	return &UnaryExpr{
		OpStart: pos,
		Op:      op,
		X:       x,
	}
}

// binary returns a binary expression with the given
// operands, position, and operator.
func binary(x Expr, pos Position, op string, y Expr) Expr {
	_, xend := x.Span()
	ystart, _ := y.Span()
	return &BinaryExpr{
		X:         x,
		OpStart:   pos,
		Op:        op,
		LineBreak: xend.Line < ystart.Line,
		Y:         y,
	}
}

// forceCompact returns the setting for the ForceCompact field for a call or tuple.
//
// NOTE 1: The field is called ForceCompact, not ForceSingleLine,
// because it only affects the formatting associated with the call or tuple syntax,
// not the formatting of the arguments. For example:
//
//	call([
//		1,
//		2,
//		3,
//	])
//
// is still a compact call even though it runs on multiple lines.
//
// In contrast the multiline form puts a linebreak after the (.
//
//	call(
//		[
//			1,
//			2,
//			3,
//		],
//	)
//
// NOTE 2: Because of NOTE 1, we cannot use start and end on the
// same line as a signal for compact mode: the formatting of an
// embedded list might move the end to a different line, which would
// then look different on rereading and cause buildifier not to be
// idempotent. Instead, we have to look at properties guaranteed
// to be preserved by the reformatting, namely that the opening
// paren and the first expression are on the same line and that
// each subsequent expression begins on the same line as the last
// one ended (no line breaks after comma).
func forceCompact(start Position, list []Expr, end Position) bool {
	if len(list) <= 1 {
		// The call or tuple will probably be compact anyway; don't force it.
		return false
	}

	// If there are any named arguments or non-string, non-literal
	// arguments, cannot force compact mode.
	line := start.Line
	for _, x := range list {
		start, end := x.Span()
		if start.Line != line {
			return false
		}
		line = end.Line
		switch x.(type) {
		case *LiteralExpr, *StringExpr:
			// ok
		default:
			return false
		}
	}
	return end.Line == line
}

// forceMultiLine returns the setting for the ForceMultiLine field.
func forceMultiLine(start Position, list []Expr, end Position) bool {
	if len(list) > 1 {
		// The call will be multiline anyway, because it has multiple elements. Don't force it.
		return false
	}

	if len(list) == 0 {
		// Empty list: use position of brackets.
		return start.Line != end.Line
	}

	// Single-element list.
	// Check whether opening bracket is on different line than beginning of
	// element, or closing bracket is on different line than end of element.
	elemStart, elemEnd := list[0].Span()
	return start.Line != elemStart.Line || end.Line != elemEnd.Line
}

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 71
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 542

var yyAct = [...]int{

	53, 109, 9, 7, 65, 51, 86, 100, 21, 20,
	135, 80, 47, 49, 87, 56, 57, 58, 59, 124,
	18, 61, 107, 129, 127, 63, 64, 66, 67, 68,
	69, 70, 71, 72, 73, 74, 75, 76, 77, 78,
	79, 125, 81, 82, 83, 84, 123, 123, 12, 110,
	17, 88, 128, 15, 122, 46, 91, 90, 93, 94,
	11, 123, 13, 97, 130, 123, 85, 104, 50, 48,
	102, 18, 18, 96, 99, 89, 14, 24, 98, 16,
	62, 105, 20, 23, 55, 27, 24, 22, 26, 25,
	112, 111, 23, 28, 134, 101, 115, 124, 25, 117,
	112, 108, 116, 92, 19, 120, 108, 121, 108, 119,
	60, 1, 126, 111, 113, 45, 114, 108, 10, 52,
	54, 4, 2, 0, 131, 118, 133, 132, 0, 0,
	0, 27, 24, 0, 26, 29, 136, 30, 23, 28,
	0, 31, 37, 32, 25, 0, 0, 0, 38, 42,
	0, 0, 33, 0, 36, 0, 44, 106, 39, 43,
	0, 34, 35, 40, 41, 27, 24, 0, 26, 29,
	0, 30, 23, 28, 0, 31, 37, 32, 25, 103,
	0, 0, 38, 42, 0, 0, 33, 0, 36, 0,
	44, 0, 39, 43, 0, 34, 35, 40, 41, 27,
	24, 0, 26, 29, 0, 30, 23, 28, 0, 31,
	37, 32, 25, 0, 0, 0, 38, 42, 0, 0,
	33, 88, 36, 0, 44, 0, 39, 43, 0, 34,
	35, 40, 41, 27, 24, 0, 26, 29, 0, 30,
	23, 28, 95, 31, 37, 32, 25, 0, 0, 0,
	38, 42, 0, 0, 33, 0, 36, 0, 44, 0,
	39, 43, 0, 34, 35, 40, 41, 27, 24, 0,
	26, 29, 0, 30, 23, 28, 0, 31, 37, 32,
	25, 0, 0, 0, 38, 42, 0, 0, 33, 0,
	36, 0, 44, 0, 39, 43, 0, 34, 35, 40,
	41, 27, 24, 0, 26, 29, 0, 30, 23, 28,
	0, 31, 37, 32, 25, 0, 0, 0, 38, 42,
	0, 0, 33, 0, 36, 0, 0, 0, 39, 43,
	0, 34, 35, 40, 41, 27, 24, 0, 26, 29,
	0, 30, 23, 28, 0, 31, 0, 32, 25, 0,
	0, 0, 0, 42, 0, 0, 33, 0, 36, 0,
	44, 0, 39, 43, 0, 34, 35, 40, 41, 27,
	24, 0, 26, 29, 0, 30, 23, 28, 0, 31,
	0, 32, 25, 0, 0, 0, 0, 42, 0, 0,
	33, 0, 36, 12, 0, 17, 39, 43, 15, 34,
	35, 40, 41, 0, 0, 11, 0, 13, 0, 0,
	0, 6, 3, 0, 0, 0, 18, 0, 0, 0,
	0, 14, 0, 0, 16, 0, 8, 20, 0, 5,
	27, 24, 0, 26, 29, 0, 30, 23, 28, 0,
	31, 0, 32, 25, 0, 0, 0, 0, 42, 0,
	0, 33, 0, 36, 0, 0, 0, 0, 0, 0,
	34, 35, 0, 41, 27, 24, 0, 26, 29, 0,
	30, 23, 28, 0, 31, 0, 32, 25, 0, 0,
	0, 0, 42, 0, 0, 33, 0, 36, 0, 0,
	0, 0, 0, 0, 34, 35, 27, 24, 0, 26,
	29, 0, 30, 23, 28, 0, 31, 0, 32, 25,
	0, 0, 0, 0, 0, 0, 0, 33, 0, 36,
	0, 0, 0, 0, 0, 0, 34, 35, 27, 24,
	0, 26, 29, 0, 30, 23, 28, 0, 0, 0,
	0, 25,
}
var yyPact = [...]int{

	-1000, -1000, 388, -1000, 78, -1000, -1000, 263, -1000, -1000,
	-30, 43, 43, 43, 43, 43, 43, 43, -1000, -1000,
	-1000, -1000, -1000, -7, 43, 43, 43, 43, 43, 43,
	43, 43, 43, 43, 43, 43, 43, 43, 43, 43,
	-20, 43, 43, 43, 43, -1000, 48, 195, 66, 195,
	97, 25, 39, 229, 64, 65, 263, -1000, -1000, -1000,
	-37, -1000, 89, 195, 161, 54, 72, 72, 72, 81,
	81, 524, 524, 524, 524, 524, 524, 331, 331, 426,
	43, 460, 492, 426, 127, -1000, 25, -1000, 44, 43,
	-1000, 25, -1000, 25, -1000, 43, 43, -1000, 43, 43,
	-1000, -1000, 25, -1000, 43, 426, 43, 36, -1000, 10,
	-8, -1000, 263, 18, 32, 263, -1000, 365, 17, 46,
	263, 365, -1000, 43, -8, 43, 88, -1000, -1000, -1000,
	-1000, 297, -1000, 297, -21, 43, 297,
}
var yyPgo = [...]int{

	0, 8, 0, 4, 69, 55, 14, 6, 2, 1,
	22, 122, 121, 5, 120, 119, 104, 118, 111, 110,
}
var yyR1 = [...]int{

	0, 18, 11, 11, 11, 11, 12, 12, 19, 19,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 3, 3, 1, 1,
	13, 14, 14, 15, 15, 4, 4, 5, 5, 16,
	17, 17, 8, 9, 9, 6, 6, 7, 7, 10,
	10,
}
var yyR2 = [...]int{

	0, 2, 0, 4, 2, 2, 1, 1, 0, 2,
	1, 1, 3, 5, 5, 5, 3, 3, 3, 4,
	6, 4, 6, 4, 2, 2, 2, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 4, 3, 3, 3, 5, 0, 1, 0, 1,
	3, 1, 3, 0, 2, 1, 3, 0, 2, 1,
	1, 2, 1, 1, 3, 4, 6, 1, 2, 0,
	3,
}
var yyChk = [...]int{

	-1000, -18, -11, 24, -12, 41, 23, -2, 38, -8,
	-17, 17, 5, 19, 33, 10, 36, 7, 28, -16,
	39, -1, 9, 11, 5, 17, 7, 4, 12, 8,
	10, 14, 16, 25, 34, 35, 27, 15, 21, 31,
	36, 37, 22, 32, 29, -16, -5, -2, -4, -2,
	-5, -13, -15, -2, -14, -4, -2, -2, -2, -2,
	-19, 28, -5, -2, -2, -3, -2, -2, -2, -2,
	-2, -2, -2, -2, -2, -2, -2, -2, -2, -2,
	31, -2, -2, -2, -2, 18, -7, -6, 26, 9,
	-1, -7, 6, -7, 20, 13, 9, -1, 13, 9,
	44, 6, -7, 18, 13, -2, 30, -10, -6, -9,
	5, -8, -2, -10, -10, -2, -13, -2, -10, -3,
	-2, -2, 18, 29, 9, 31, -9, 6, 20, 6,
	18, -2, -8, -2, 6, 31, -2,
}
var yyDef = [...]int{

	2, -2, 0, 1, 48, 4, 5, 6, 7, 10,
	11, 57, 57, 53, 0, 0, 0, 0, 62, 60,
	59, 8, 49, 0, 57, 46, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 61, 0, 55, 48, 55,
	0, 51, 0, 0, 48, 0, 55, 24, 25, 26,
	3, 18, 0, 55, 47, 0, 27, 28, 29, 30,
	31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
	0, 42, 43, 44, 0, 12, 69, 67, 0, 49,
	58, 69, 17, 69, 16, 0, 49, 54, 0, 0,
	9, 19, 69, 21, 46, 41, 0, 0, 68, 0,
	0, 63, 56, 0, 0, 50, 52, 23, 0, 0,
	47, 45, 13, 0, 0, 0, 0, 14, 15, 20,
	22, 70, 64, 65, 0, 0, 66,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	41, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 4, 3, 3,
	5, 6, 7, 8, 9, 10, 11, 12, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 13, 44,
	14, 15, 16, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 17, 3, 18, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 19, 3, 20,
}
var yyTok2 = [...]int{

	2, 3, 21, 22, 23, 24, 25, 26, 27, 28,
	29, 30, 31, 32, 33, 34, 35, 36, 37, 38,
	39, 40, 42, 43,
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
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:151
		{
			yylex.(*input).file = &File{Stmt: yyDollar[1].exprs}
			return 0
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:157
		{
			yyVAL.exprs = nil
			yyVAL.lastRule = nil
		}
	case 3:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:162
		{
			// If this statement follows a comment block,
			// attach the comments to the statement.
			if cb, ok := yyDollar[1].lastRule.(*CommentBlock); ok {
				yyVAL.exprs = yyDollar[1].exprs
				yyVAL.exprs[len(yyDollar[1].exprs)-1] = yyDollar[2].expr
				yyDollar[2].expr.Comment().Before = cb.After
				yyVAL.lastRule = yyDollar[2].expr
				break
			}

			// Otherwise add to list.
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[2].expr)
			yyVAL.lastRule = yyDollar[2].expr

			// Consider this input:
			//
			//	foo()
			//	# bar
			//	baz()
			//
			// If we've just parsed baz(), the # bar is attached to
			// foo() as an After comment. Make it a Before comment
			// for baz() instead.
			if x := yyDollar[1].lastRule; x != nil {
				com := x.Comment()
				yyDollar[2].expr.Comment().Before = com.After
				com.After = nil
			}
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:193
		{
			// Blank line; sever last rule from future comments.
			yyVAL.exprs = yyDollar[1].exprs
			yyVAL.lastRule = nil
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:199
		{
			yyVAL.exprs = yyDollar[1].exprs
			yyVAL.lastRule = yyDollar[1].lastRule
			if yyVAL.lastRule == nil {
				cb := &CommentBlock{Start: yyDollar[2].pos}
				yyVAL.exprs = append(yyVAL.exprs, cb)
				yyVAL.lastRule = cb
			}
			com := yyVAL.lastRule.Comment()
			com.After = append(com.After, Comment{Start: yyDollar[2].pos, Token: yyDollar[2].tok})
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:214
		{
			yyVAL.expr = &PythonBlock{Start: yyDollar[1].pos, Token: yyDollar[1].tok}
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:224
		{
			if len(yyDollar[1].strings) == 1 {
				yyVAL.expr = yyDollar[1].strings[0]
				break
			}

			yyVAL.expr = yyDollar[1].strings[0]
			for _, x := range yyDollar[1].strings[1:] {
				_, end := yyVAL.expr.Span()
				yyVAL.expr = binary(yyVAL.expr, end, "+", x)
			}
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:237
		{
			yyVAL.expr = &ListExpr{
				Start:          yyDollar[1].pos,
				List:           yyDollar[2].exprs,
				Comma:          yyDollar[2].comma,
				End:            End{Pos: yyDollar[3].pos},
				ForceMultiLine: forceMultiLine(yyDollar[1].pos, yyDollar[2].exprs, yyDollar[3].pos),
			}
		}
	case 13:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parse.y:247
		{
			exprStart, _ := yyDollar[2].expr.Span()
			yyVAL.expr = &ListForExpr{
				Brack:          "[]",
				Start:          yyDollar[1].pos,
				X:              yyDollar[2].expr,
				For:            yyDollar[3].fors,
				If:             yyDollar[4].ifs,
				End:            End{Pos: yyDollar[5].pos},
				ForceMultiLine: yyDollar[1].pos.Line != exprStart.Line,
			}
		}
	case 14:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parse.y:260
		{
			exprStart, _ := yyDollar[2].expr.Span()
			yyVAL.expr = &ListForExpr{
				Brack:          "()",
				Start:          yyDollar[1].pos,
				X:              yyDollar[2].expr,
				For:            yyDollar[3].fors,
				If:             yyDollar[4].ifs,
				End:            End{Pos: yyDollar[5].pos},
				ForceMultiLine: yyDollar[1].pos.Line != exprStart.Line,
			}
		}
	case 15:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parse.y:273
		{
			exprStart, _ := yyDollar[2].expr.Span()
			yyVAL.expr = &ListForExpr{
				Brack:          "{}",
				Start:          yyDollar[1].pos,
				X:              yyDollar[2].expr,
				For:            yyDollar[3].fors,
				If:             yyDollar[4].ifs,
				End:            End{Pos: yyDollar[5].pos},
				ForceMultiLine: yyDollar[1].pos.Line != exprStart.Line,
			}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:286
		{
			yyVAL.expr = &DictExpr{
				Start:          yyDollar[1].pos,
				List:           yyDollar[2].exprs,
				Comma:          yyDollar[2].comma,
				End:            End{Pos: yyDollar[3].pos},
				ForceMultiLine: forceMultiLine(yyDollar[1].pos, yyDollar[2].exprs, yyDollar[3].pos),
			}
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:296
		{
			if len(yyDollar[2].exprs) == 1 && yyDollar[2].comma.Line == 0 {
				// Just a parenthesized expression, not a tuple.
				yyVAL.expr = &ParenExpr{
					Start:          yyDollar[1].pos,
					X:              yyDollar[2].exprs[0],
					End:            End{Pos: yyDollar[3].pos},
					ForceMultiLine: forceMultiLine(yyDollar[1].pos, yyDollar[2].exprs, yyDollar[3].pos),
				}
			} else {
				yyVAL.expr = &TupleExpr{
					Start:          yyDollar[1].pos,
					List:           yyDollar[2].exprs,
					Comma:          yyDollar[2].comma,
					End:            End{Pos: yyDollar[3].pos},
					ForceCompact:   forceCompact(yyDollar[1].pos, yyDollar[2].exprs, yyDollar[3].pos),
					ForceMultiLine: forceMultiLine(yyDollar[1].pos, yyDollar[2].exprs, yyDollar[3].pos),
				}
			}
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:317
		{
			yyVAL.expr = &DotExpr{
				X:       yyDollar[1].expr,
				Dot:     yyDollar[2].pos,
				NamePos: yyDollar[3].pos,
				Name:    yyDollar[3].tok,
			}
		}
	case 19:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:326
		{
			yyVAL.expr = &CallExpr{
				X:              yyDollar[1].expr,
				ListStart:      yyDollar[2].pos,
				List:           yyDollar[3].exprs,
				End:            End{Pos: yyDollar[4].pos},
				ForceCompact:   forceCompact(yyDollar[2].pos, yyDollar[3].exprs, yyDollar[4].pos),
				ForceMultiLine: forceMultiLine(yyDollar[2].pos, yyDollar[3].exprs, yyDollar[4].pos),
			}
		}
	case 20:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parse.y:337
		{
			yyVAL.expr = &CallExpr{
				X:         yyDollar[1].expr,
				ListStart: yyDollar[2].pos,
				List: []Expr{
					&ListForExpr{
						Brack: "",
						Start: yyDollar[2].pos,
						X:     yyDollar[3].expr,
						For:   yyDollar[4].fors,
						If:    yyDollar[5].ifs,
						End:   End{Pos: yyDollar[6].pos},
					},
				},
				End: End{Pos: yyDollar[6].pos},
			}
		}
	case 21:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:355
		{
			yyVAL.expr = &IndexExpr{
				X:          yyDollar[1].expr,
				IndexStart: yyDollar[2].pos,
				Y:          yyDollar[3].expr,
				End:        yyDollar[4].pos,
			}
		}
	case 22:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parse.y:364
		{
			yyVAL.expr = &SliceExpr{
				X:          yyDollar[1].expr,
				SliceStart: yyDollar[2].pos,
				Y:          yyDollar[3].expr,
				Colon:      yyDollar[4].pos,
				Z:          yyDollar[5].expr,
				End:        yyDollar[6].pos,
			}
		}
	case 23:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:375
		{
			yyVAL.expr = &LambdaExpr{
				Lambda: yyDollar[1].pos,
				Var:    yyDollar[2].exprs,
				Colon:  yyDollar[3].pos,
				Expr:   yyDollar[4].expr,
			}
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:383
		{
			yyVAL.expr = unary(yyDollar[1].pos, yyDollar[1].tok, yyDollar[2].expr)
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:384
		{
			yyVAL.expr = unary(yyDollar[1].pos, yyDollar[1].tok, yyDollar[2].expr)
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:385
		{
			yyVAL.expr = unary(yyDollar[1].pos, yyDollar[1].tok, yyDollar[2].expr)
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:386
		{
			yyVAL.expr = binary(yyDollar[1].expr, yyDollar[2].pos, yyDollar[2].tok, yyDollar[3].expr)
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:387
		{
			yyVAL.expr = binary(yyDollar[1].expr, yyDollar[2].pos, yyDollar[2].tok, yyDollar[3].expr)
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:388
		{
			yyVAL.expr = binary(yyDollar[1].expr, yyDollar[2].pos, yyDollar[2].tok, yyDollar[3].expr)
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:389
		{
			yyVAL.expr = binary(yyDollar[1].expr, yyDollar[2].pos, yyDollar[2].tok, yyDollar[3].expr)
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:390
		{
			yyVAL.expr = binary(yyDollar[1].expr, yyDollar[2].pos, yyDollar[2].tok, yyDollar[3].expr)
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:391
		{
			yyVAL.expr = binary(yyDollar[1].expr, yyDollar[2].pos, yyDollar[2].tok, yyDollar[3].expr)
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:392
		{
			yyVAL.expr = binary(yyDollar[1].expr, yyDollar[2].pos, yyDollar[2].tok, yyDollar[3].expr)
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:393
		{
			yyVAL.expr = binary(yyDollar[1].expr, yyDollar[2].pos, yyDollar[2].tok, yyDollar[3].expr)
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:394
		{
			yyVAL.expr = binary(yyDollar[1].expr, yyDollar[2].pos, yyDollar[2].tok, yyDollar[3].expr)
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:395
		{
			yyVAL.expr = binary(yyDollar[1].expr, yyDollar[2].pos, yyDollar[2].tok, yyDollar[3].expr)
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:396
		{
			yyVAL.expr = binary(yyDollar[1].expr, yyDollar[2].pos, yyDollar[2].tok, yyDollar[3].expr)
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:397
		{
			yyVAL.expr = binary(yyDollar[1].expr, yyDollar[2].pos, yyDollar[2].tok, yyDollar[3].expr)
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:398
		{
			yyVAL.expr = binary(yyDollar[1].expr, yyDollar[2].pos, yyDollar[2].tok, yyDollar[3].expr)
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:399
		{
			yyVAL.expr = binary(yyDollar[1].expr, yyDollar[2].pos, yyDollar[2].tok, yyDollar[3].expr)
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:400
		{
			yyVAL.expr = binary(yyDollar[1].expr, yyDollar[2].pos, "not in", yyDollar[4].expr)
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:401
		{
			yyVAL.expr = binary(yyDollar[1].expr, yyDollar[2].pos, yyDollar[2].tok, yyDollar[3].expr)
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:402
		{
			yyVAL.expr = binary(yyDollar[1].expr, yyDollar[2].pos, yyDollar[2].tok, yyDollar[3].expr)
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:404
		{
			if b, ok := yyDollar[3].expr.(*UnaryExpr); ok && b.Op == "not" {
				yyVAL.expr = binary(yyDollar[1].expr, yyDollar[2].pos, "is not", b.X)
			} else {
				yyVAL.expr = binary(yyDollar[1].expr, yyDollar[2].pos, yyDollar[2].tok, yyDollar[3].expr)
			}
		}
	case 45:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parse.y:412
		{
			yyVAL.expr = &ConditionalExpr{
				Then:      yyDollar[1].expr,
				IfStart:   yyDollar[2].pos,
				Test:      yyDollar[3].expr,
				ElseStart: yyDollar[4].pos,
				Else:      yyDollar[5].expr,
			}
		}
	case 46:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:423
		{
			yyVAL.expr = nil
		}
	case 48:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:433
		{
			yyVAL.pos = Position{}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:439
		{
			yyVAL.expr = &KeyValueExpr{
				Key:   yyDollar[1].expr,
				Colon: yyDollar[2].pos,
				Value: yyDollar[3].expr,
			}
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:449
		{
			yyVAL.exprs = []Expr{yyDollar[1].expr}
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:453
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[3].expr)
		}
	case 53:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:458
		{
			yyVAL.exprs, yyVAL.comma = nil, Position{}
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:462
		{
			yyVAL.exprs, yyVAL.comma = yyDollar[1].exprs, yyDollar[2].pos
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:468
		{
			yyVAL.exprs = []Expr{yyDollar[1].expr}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:472
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[3].expr)
		}
	case 57:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:477
		{
			yyVAL.exprs, yyVAL.comma = nil, Position{}
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:481
		{
			yyVAL.exprs, yyVAL.comma = yyDollar[1].exprs, yyDollar[2].pos
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:487
		{
			yyVAL.string = &StringExpr{
				Start:       yyDollar[1].pos,
				Value:       yyDollar[1].str,
				TripleQuote: yyDollar[1].triple,
				End:         yyDollar[1].pos.add(yyDollar[1].tok),
				Token:       yyDollar[1].tok,
			}
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:499
		{
			yyVAL.strings = []*StringExpr{yyDollar[1].string}
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:503
		{
			yyVAL.strings = append(yyDollar[1].strings, yyDollar[2].string)
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:509
		{
			yyVAL.expr = &LiteralExpr{Start: yyDollar[1].pos, Token: yyDollar[1].tok}
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:515
		{
			yyVAL.exprs = []Expr{yyDollar[1].expr}
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:519
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[3].expr)
		}
	case 65:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parse.y:525
		{
			yyVAL.forc = &ForClause{
				For:  yyDollar[1].pos,
				Var:  yyDollar[2].exprs,
				In:   yyDollar[3].pos,
				Expr: yyDollar[4].expr,
			}
		}
	case 66:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parse.y:534
		{
			yyVAL.forc = &ForClause{
				For:  yyDollar[1].pos,
				Var:  yyDollar[3].exprs,
				In:   yyDollar[5].pos,
				Expr: yyDollar[6].expr,
			}
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parse.y:545
		{
			yyVAL.fors = []*ForClause{yyDollar[1].forc}
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parse.y:548
		{
			yyVAL.fors = append(yyDollar[1].fors, yyDollar[2].forc)
		}
	case 69:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parse.y:553
		{
			yyVAL.ifs = nil
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parse.y:557
		{
			yyVAL.ifs = append(yyDollar[1].ifs, &IfClause{
				If:   yyDollar[2].pos,
				Cond: yyDollar[3].expr,
			})
		}
	}
	goto yystack /* stack new state and value */
}
