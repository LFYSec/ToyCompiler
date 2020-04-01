// Code generated by goyacc grammer.y. DO NOT EDIT.

//line grammer.y:2

package main

import __yyfmt__ "fmt"

//line grammer.y:3

import (
	"fmt"
	"strconv"
)

type SymbolType int

//line grammer.y:14
type yySymType struct {
	yys          int
	int_value    int
	double_value float64
	string_value string
	//node 		*Node;
	symbolType SymbolType
}

const TYPE = 57346
const IDENTIFY = 57347
const INT_LITERAL = 57348
const DOUBLE_LITERAL = 57349
const STRING = 57350
const INT = 57351
const DOUBLE = 57352
const CHAR = 57353
const BOOL = 57354
const IF = 57355
const ELSE = 57356
const WHILE = 57357
const FOR = 57358
const PRINT = 57359
const EOL = 57360
const EQ = 57361
const NE = 57362
const LT = 57363
const LE = 57364
const GT = 57365
const GE = 57366
const ASSIGN = 57367
const ADD = 57368
const SUB = 57369
const MUL = 57370
const DIV = 57371
const NOT = 57372
const LBRACKET = 57373
const RBRACKET = 57374

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"TYPE",
	"IDENTIFY",
	"INT_LITERAL",
	"DOUBLE_LITERAL",
	"STRING",
	"INT",
	"DOUBLE",
	"CHAR",
	"BOOL",
	"IF",
	"ELSE",
	"WHILE",
	"FOR",
	"PRINT",
	"EOL",
	"EQ",
	"NE",
	"LT",
	"LE",
	"GT",
	"GE",
	"ASSIGN",
	"ADD",
	"SUB",
	"MUL",
	"DIV",
	"NOT",
	"LBRACKET",
	"RBRACKET",
	"'('",
	"')'",
	"'+'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line grammer.y:125

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 40

var yyAct = [...]int{

	21, 19, 20, 21, 19, 20, 28, 33, 31, 12,
	16, 11, 9, 14, 8, 6, 17, 13, 10, 1,
	18, 23, 24, 5, 4, 25, 26, 7, 22, 3,
	15, 22, 27, 2, 0, 29, 30, 0, 0, 32,
}
var yyPact = [...]int{

	-1000, 10, -6, -1000, -1000, -1000, -14, -24, 12, -1000,
	-1000, -5, -5, -14, -1000, -29, -1000, -1000, -2, -1000,
	-1000, -1000, -5, -1000, -1000, -1000, -26, -1000, -5, -1000,
	-27, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 13, 33, 16, 10, 30, 29, 24, 23, 20,
	18, 19,
}
var yyR1 = [...]int{

	0, 11, 11, 10, 8, 8, 6, 7, 2, 2,
	2, 3, 3, 3, 3, 9, 9, 9, 4, 4,
	5, 5, 1,
}
var yyR2 = [...]int{

	0, 3, 0, 2, 2, 3, 2, 4, 1, 1,
	1, 1, 1, 1, 3, 1, 1, 1, 1, 2,
	1, 3, 1,
}
var yyChk = [...]int{

	-1000, -11, -2, -6, -7, -8, 5, 17, 4, 18,
	-10, 25, 33, 5, -1, -5, -4, -3, -9, 6,
	7, 5, 33, 26, 27, 30, -1, -10, 35, -3,
	-1, 34, -4, 34,
}
var yyDef = [...]int{

	2, -2, 0, 8, 9, 10, 0, 0, 0, 1,
	6, 0, 0, 4, 3, 22, 20, 18, 0, 11,
	12, 13, 0, 15, 16, 17, 0, 5, 0, 19,
	0, 7, 21, 14,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	33, 34, 3, 35,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32,
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
//line grammer.y:48
		{
			fmt.Println("statement: ", yyDollar[2].string_value)
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
//line grammer.y:49
		{
			fmt.Println("start")
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammer.y:53
		{
			yyVAL.string_value = yyDollar[2].string_value
			fmt.Println("assign: ", yyDollar[2].string_value)
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammer.y:60
		{
			yyVAL.string_value = strconv.Itoa(int(yyDollar[1].symbolType))
			fmt.Println("TYPE: ", yyDollar[1].symbolType)
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammer.y:64
		{
			fmt.Println("TYPE: ", yyDollar[1].symbolType)
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammer.y:70
		{
			yyVAL.string_value = yyDollar[2].string_value
		}
	case 7:
		yyDollar = yyS[yypt-4 : yypt+1]
//line grammer.y:76
		{
			fmt.Println("PRINT: ", yyDollar[3].string_value)
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammer.y:87
		{
			yyVAL.string_value = strconv.Itoa(yyDollar[1].int_value)
			fmt.Println("INT_LITERAL: ", yyDollar[1].int_value)
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammer.y:91
		{
			yyVAL.string_value = strconv.FormatFloat(yyDollar[1].double_value, 'E', -1, 64)
			fmt.Println("DOU_LITERAL: ", yyDollar[1].double_value)
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammer.y:95
		{
			yyVAL.string_value = yyDollar[1].string_value
			fmt.Println("IDENTIFY: ", yyDollar[1].string_value)
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammer.y:99
		{
			yyVAL.string_value = yyDollar[2].string_value
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammer.y:109
		{
			yyVAL.string_value = yyDollar[1].string_value
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
//line grammer.y:110
		{
			yyVAL.string_value = yyDollar[2].string_value
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammer.y:114
		{
			yyVAL.string_value = yyDollar[1].string_value
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
//line grammer.y:115
		{
			yyVAL.string_value = yyDollar[1].string_value + yyDollar[3].string_value
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
//line grammer.y:119
		{
			yyVAL.string_value = yyDollar[1].string_value
		}
	}
	goto yystack /* stack new state and value */
}
