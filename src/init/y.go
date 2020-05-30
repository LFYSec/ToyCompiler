// Code generated by goyacc -o src/init/y.go src/init/grammer.y. DO NOT EDIT.

//line src/init/grammer.y:2

package init

import __yyfmt__ "fmt"

//line src/init/grammer.y:3

import (
	"Compiler/src/ast"
	"Compiler/src/ast/expression/lvalue"
	"Compiler/src/ast/expression/lvalue/reference"
	"Compiler/src/ast/expression/rvalue"
	"Compiler/src/ast/expression/rvalue/binOperateResult"
	"Compiler/src/ast/expression/rvalue/literal"
	stmt "Compiler/src/ast/statement"
	"Compiler/src/ast/statement/function"
	"Compiler/src/global"
	"Compiler/src/symbolTable"
	"fmt"
)

var result *stmt.CompoundStmt

//line src/init/grammer.y:24
type yySymType struct {
	yys          int
	int_value    int
	double_value float64
	string_value string
	char_value   uint8
	node         ast.Node
	symbolType   global.SymbolType
}

const TYPE = 57346
const IDENTIFY = 57347
const INT_LITERAL = 57348
const DOUBLE_LITERAL = 57349
const CHAR_LITERAL = 57350
const ARRAY = 57351
const STRING = 57352
const INT = 57353
const DOUBLE = 57354
const CHAR = 57355
const BOOL = 57356
const VOID = 57357
const IF = 57358
const ELSE = 57359
const WHILE = 57360
const FOR = 57361
const PRINT = 57362
const EOL = 57363
const RETURN = 57364
const EOS = 57365
const SCAN = 57366
const EQ = 57367
const NE = 57368
const LT = 57369
const LE = 57370
const GT = 57371
const GE = 57372
const ASSIGN = 57373
const COMMA = 57374
const ADD = 57375
const SUB = 57376
const MUL = 57377
const DIV = 57378
const NOT = 57379
const LBRACKET = 57380
const RBRACKET = 57381
const LPARENTHESIS = 57382
const RPARENTHESIS = 57383
const LBRACE = 57384
const RBRACE = 57385

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"TYPE",
	"IDENTIFY",
	"INT_LITERAL",
	"DOUBLE_LITERAL",
	"CHAR_LITERAL",
	"ARRAY",
	"STRING",
	"INT",
	"DOUBLE",
	"CHAR",
	"BOOL",
	"VOID",
	"IF",
	"ELSE",
	"WHILE",
	"FOR",
	"PRINT",
	"EOL",
	"RETURN",
	"EOS",
	"SCAN",
	"EQ",
	"NE",
	"LT",
	"LE",
	"GT",
	"GE",
	"ASSIGN",
	"COMMA",
	"ADD",
	"SUB",
	"MUL",
	"DIV",
	"NOT",
	"LBRACKET",
	"RBRACKET",
	"LPARENTHESIS",
	"RPARENTHESIS",
	"LBRACE",
	"RBRACE",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line src/init/grammer.y:376

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 130

var yyAct = [...]int{

	7, 10, 25, 18, 17, 20, 92, 17, 20, 82,
	80, 38, 41, 39, 34, 94, 21, 87, 19, 21,
	16, 19, 14, 16, 22, 14, 38, 22, 33, 100,
	31, 26, 27, 30, 102, 56, 2, 99, 65, 95,
	93, 68, 18, 63, 55, 18, 70, 71, 72, 73,
	74, 75, 76, 77, 78, 79, 83, 81, 54, 23,
	42, 91, 103, 85, 37, 29, 40, 28, 84, 15,
	89, 101, 58, 64, 53, 31, 26, 27, 30, 57,
	61, 108, 32, 66, 105, 35, 69, 60, 106, 62,
	104, 86, 98, 97, 20, 26, 27, 30, 12, 11,
	9, 3, 90, 13, 15, 107, 8, 6, 5, 4,
	29, 88, 52, 51, 48, 49, 47, 50, 59, 24,
	43, 44, 45, 46, 67, 96, 36, 1, 0, 29,
}
var yyPact = [...]int{

	-1000, 3, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 25, -3, -26, 80, -1000, 25,
	-27, 25, -28, 37, 87, -1000, -1000, -1000, -1000, 25,
	-1000, -12, 35, 89, 25, 49, 0, -39, 25, -1000,
	-39, 25, -1000, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, -31, -1000, -1000, 34, -32, -1000, 33,
	62, 89, -1000, -1000, -1000, -1000, -22, 70, 44, -35,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 17, -1000, -24, 16, -1000, -1000, -1000, -1000,
	25, -39, 14, -1000, 6, -1000, 30, -1000, -1000, -1000,
	-1000, -1000, 79, 84, -39, -1000, 76, -1000, -1000,
}
var yyPgo = [...]int{

	0, 44, 36, 127, 82, 0, 126, 125, 124, 2,
	119, 67, 109, 108, 107, 106, 103, 101, 100, 1,
	99, 98, 91, 90,
}
var yyR1 = [...]int{

	0, 3, 3, 14, 14, 14, 4, 12, 13, 16,
	21, 21, 15, 17, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 9, 9, 9, 9, 9,
	11, 11, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 1, 5, 6, 6, 19, 20, 20,
	22, 23, 18, 7, 7, 7, 8, 8, 8,
}
var yyR2 = [...]int{

	0, 0, 2, 3, 4, 6, 2, 3, 5, 5,
	3, 5, 3, 3, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 3, 1,
	1, 4, 1, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 1, 3, 2, 0, 4, 4, 5,
	0, 0, 8, 0, 3, 4, 0, 2, 3,
}
var yyChk = [...]int{

	-1000, -3, -2, -17, -12, -13, -14, -5, -15, -18,
	-19, -20, -21, -16, 22, -11, 20, 4, 42, 18,
	5, 16, 24, -1, -10, -9, 6, 7, -11, 40,
	8, 5, -4, 31, 40, 5, -6, -1, 38, 40,
	-1, 40, 23, 33, 34, 35, 36, 29, 27, 28,
	30, 26, 25, -1, 23, -1, -19, -1, 23, -4,
	38, 31, 40, 43, -2, -5, -1, -8, -5, -1,
	-9, -9, -9, -9, -9, -9, -9, -9, -9, -9,
	41, 23, 41, 23, 6, -19, -22, 39, 41, -9,
	32, 17, 41, 23, 39, 23, -7, -9, -5, 23,
	23, 41, 4, 32, -23, 5, 4, -5, 5,
}
var yyDef = [...]int{

	1, -2, 2, 14, 15, 16, 17, 18, 19, 20,
	21, 22, 23, 24, 0, 0, 0, 0, 46, 0,
	30, 0, 0, 0, 43, 32, 25, 26, 27, 0,
	29, 30, 0, 0, 0, 0, 0, 0, 0, 56,
	0, 0, 13, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 7, 6, 0, 0, 3, 0,
	0, 0, 50, 44, 45, 12, 0, 0, 10, 0,
	33, 34, 35, 36, 37, 38, 39, 40, 41, 42,
	28, 48, 0, 4, 0, 0, 53, 31, 47, 57,
	0, 0, 0, 8, 0, 49, 0, 58, 11, 9,
	5, 51, 0, 0, 0, 54, 0, 52, 55,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43,
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
		yyDollar = yyS[yypt-0 : yypt+1]
//line src/init/grammer.y:62
		{
			result = stmt.CreateCompoundStmt()
			yyVAL.node = result
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
//line src/init/grammer.y:66
		{
			v1 := (yyDollar[1].node).(*stmt.CompoundStmt)
			v2 := (yyDollar[2].node).(stmt.Stmt)
			stmt.AddStmt(v1, v2)
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
//line src/init/grammer.y:74
		{
			s := symbolTable.CreateSymbol(true, yyDollar[1].symbolType, yyDollar[2].string_value, 0)
			symbolTable.AddSymbol(s)
			yyVAL.node = stmt.CreateDeclareStmt(s, stmt.AssignStmt{})
		}
	case 4:
		yyDollar = yyS[yypt-4 : yypt+1]
//line src/init/grammer.y:79
		{
			s := symbolTable.CreateSymbol(true, yyDollar[1].symbolType, yyDollar[2].string_value, 0)
			symbolTable.AddSymbol(s)
			ref := reference.CreateVariableReference(s)
			v := (yyDollar[3].node).(rvalue.RValue)
			a := stmt.CreateAssignStmt(ref, v)
			yyVAL.node = stmt.CreateDeclareStmt(s, a)
		}
	case 5:
		yyDollar = yyS[yypt-6 : yypt+1]
//line src/init/grammer.y:88
		{
			s := symbolTable.CreateSymbol(true, yyDollar[1].symbolType, yyDollar[2].string_value, yyDollar[4].int_value)
			symbolTable.AddSymbol(s)
			yyVAL.node = stmt.CreateDeclareStmt(s, stmt.AssignStmt{})
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
//line src/init/grammer.y:96
		{
			yyVAL.node = yyDollar[2].node
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
//line src/init/grammer.y:102
		{
			v1 := (yyDollar[1].node).(lvalue.LValue)
			v2 := (yyDollar[2].node).(rvalue.RValue)
			yyVAL.node = stmt.CreateAssignStmt(v1, v2)
		}
	case 8:
		yyDollar = yyS[yypt-5 : yypt+1]
//line src/init/grammer.y:110
		{
			v3 := (yyDollar[3].node).(rvalue.RValue)
			yyVAL.node = stmt.CreatePrintStmt(v3)
		}
	case 9:
		yyDollar = yyS[yypt-5 : yypt+1]
//line src/init/grammer.y:117
		{
			v3 := (yyDollar[3].node).(rvalue.RValue)
			yyVAL.node = stmt.CreateScanStmt(v3)
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
//line src/init/grammer.y:124
		{
			symbolTable.PushFrame()
			v2 := (yyDollar[2].node).(rvalue.RValue)
			v3 := (yyDollar[3].node).(*stmt.CompoundStmt)
			yyVAL.node = stmt.CreateIfStmt(v2, v3, nil)
			symbolTable.PopFrame()
		}
	case 11:
		yyDollar = yyS[yypt-5 : yypt+1]
//line src/init/grammer.y:131
		{
			symbolTable.PushFrame()
			v2 := (yyDollar[2].node).(rvalue.RValue)
			v3 := (yyDollar[3].node).(*stmt.CompoundStmt)
			v5 := (yyDollar[5].node).(*stmt.CompoundStmt)
			yyVAL.node = stmt.CreateIfStmt(v2, v3, v5)
			symbolTable.PopFrame()
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
//line src/init/grammer.y:141
		{
			symbolTable.PushFrame()
			v2 := (yyDollar[2].node).(rvalue.RValue)
			v3 := (yyDollar[3].node).(*stmt.CompoundStmt)
			yyVAL.node = stmt.CreateWhileStmt(v2, v3)
			symbolTable.PopFrame()
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
//line src/init/grammer.y:151
		{
			v2 := (yyDollar[2].node).(rvalue.RValue)
			yyVAL.node = stmt.CreateRetStmt(v2)
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
//line src/init/grammer.y:173
		{
			yyVAL.node = literal.CreateIntLiteral(yyDollar[1].int_value)
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
//line src/init/grammer.y:176
		{
			yyVAL.node = literal.CreateDoubleLiteral(yyDollar[1].double_value)
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
//line src/init/grammer.y:179
		{
			yyVAL.node = yyDollar[1].node
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
//line src/init/grammer.y:182
		{
			yyVAL.node = yyDollar[2].node
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
//line src/init/grammer.y:185
		{
			yyVAL.node = literal.CreateCharLiteral(yyDollar[1].char_value)
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
//line src/init/grammer.y:191
		{
			sb := symbolTable.GetVarSymbol(yyDollar[1].string_value)
			if sb == nil {
				panic("[*] Error: Undefined variable: " + string(yyDollar[1].string_value))
			}
			yyVAL.node = reference.CreateVariableReference(sb)
		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
//line src/init/grammer.y:198
		{
			sb := symbolTable.GetVarSymbol(yyDollar[1].string_value)
			if sb == nil {
				panic("[*] Error: Undefined array: " + string(yyDollar[1].string_value))
			}
			v3 := (yyDollar[3].node).(rvalue.RValue)
			yyVAL.node = reference.CreateArrayReference(sb, v3)
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
//line src/init/grammer.y:209
		{
			yyVAL.node = yyDollar[1].node
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
//line src/init/grammer.y:212
		{
			v1 := (yyDollar[1].node).(rvalue.RValue)
			v3 := (yyDollar[3].node).(rvalue.RValue)
			yyVAL.node = binOperateResult.CreateBinOperateResult(ADD, v1, v3)
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
//line src/init/grammer.y:217
		{
			v1 := (yyDollar[1].node).(rvalue.RValue)
			v3 := (yyDollar[3].node).(rvalue.RValue)
			yyVAL.node = binOperateResult.CreateBinOperateResult(SUB, v1, v3)
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
//line src/init/grammer.y:222
		{
			v1 := (yyDollar[1].node).(rvalue.RValue)
			v3 := (yyDollar[3].node).(rvalue.RValue)
			yyVAL.node = binOperateResult.CreateBinOperateResult(MUL, v1, v3)
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
//line src/init/grammer.y:227
		{
			v1 := (yyDollar[1].node).(rvalue.RValue)
			v3 := (yyDollar[3].node).(rvalue.RValue)
			yyVAL.node = binOperateResult.CreateBinOperateResult(DIV, v1, v3)
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
//line src/init/grammer.y:232
		{
			v1 := (yyDollar[1].node).(rvalue.RValue)
			v3 := (yyDollar[3].node).(rvalue.RValue)
			yyVAL.node = binOperateResult.CreateBinOperateResult(GT, v1, v3)
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
//line src/init/grammer.y:237
		{
			v1 := (yyDollar[1].node).(rvalue.RValue)
			v3 := (yyDollar[3].node).(rvalue.RValue)
			yyVAL.node = binOperateResult.CreateBinOperateResult(LT, v1, v3)
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
//line src/init/grammer.y:242
		{
			v1 := (yyDollar[1].node).(rvalue.RValue)
			v3 := (yyDollar[3].node).(rvalue.RValue)
			yyVAL.node = binOperateResult.CreateBinOperateResult(LE, v1, v3)
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
//line src/init/grammer.y:247
		{
			v1 := (yyDollar[1].node).(rvalue.RValue)
			v3 := (yyDollar[3].node).(rvalue.RValue)
			yyVAL.node = binOperateResult.CreateBinOperateResult(GE, v1, v3)
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
//line src/init/grammer.y:252
		{
			v1 := (yyDollar[1].node).(rvalue.RValue)
			v3 := (yyDollar[3].node).(rvalue.RValue)
			yyVAL.node = binOperateResult.CreateBinOperateResult(NE, v1, v3)
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
//line src/init/grammer.y:257
		{
			v1 := (yyDollar[1].node).(rvalue.RValue)
			v3 := (yyDollar[3].node).(rvalue.RValue)
			yyVAL.node = binOperateResult.CreateBinOperateResult(EQ, v1, v3)
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
//line src/init/grammer.y:265
		{
			yyVAL.node = yyDollar[1].node
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
//line src/init/grammer.y:271
		{
			yyVAL.node = yyDollar[2].node
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
//line src/init/grammer.y:277
		{
			v1 := (yyDollar[1].node).(*stmt.CompoundStmt)
			v2 := (yyDollar[2].node).(stmt.Stmt)
			stmt.AddStmt(v1, v2)
		}
	case 46:
		yyDollar = yyS[yypt-0 : yypt+1]
//line src/init/grammer.y:282
		{
			yyVAL.node = stmt.CreateCompoundStmt()
		}
	case 47:
		yyDollar = yyS[yypt-4 : yypt+1]
//line src/init/grammer.y:290
		{
			v1 := yyDollar[1].string_value
			if (yyDollar[3].node) != nil {
				v3 := (yyDollar[3].node).([]*rvalue.RValue)
				yyVAL.node = function.CreateFuncCallStmt(v1, v3)
			} else {
				yyVAL.node = function.CreateFuncCallStmt(v1, nil)
			}
		}
	case 48:
		yyDollar = yyS[yypt-4 : yypt+1]
//line src/init/grammer.y:302
		{
			v1 := (yyDollar[1].node).(lvalue.LValue)
			v3 := (yyDollar[3].node).(stmt.Stmt)
			yyVAL.node = function.CreateAssignFunCallStmt(v1, v3)
		}
	case 49:
		yyDollar = yyS[yypt-5 : yypt+1]
//line src/init/grammer.y:308
		{
			s := symbolTable.CreateSymbol(true, yyDollar[1].symbolType, yyDollar[2].string_value, 0)
			symbolTable.AddSymbol(s)
			ref := reference.CreateVariableReference(s)
			v4 := (yyDollar[4].node).(stmt.Stmt)
			yyVAL.node = function.CreateAssignFunCallStmt(ref, v4)
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
//line src/init/grammer.y:318
		{
			symbolTable.PushFrame()
		}
	case 51:
		yyDollar = yyS[yypt-6 : yypt+1]
//line src/init/grammer.y:318
		{
			if (yyDollar[5].node) == nil {
				function.AddFunc(yyDollar[1].symbolType, yyDollar[2].string_value, nil)
			} else {
				v5 := (yyDollar[5].node).([]*stmt.DeclareStmt)
				function.AddFunc(yyDollar[1].symbolType, yyDollar[2].string_value, v5)
			}
		}
	case 52:
		yyDollar = yyS[yypt-8 : yypt+1]
//line src/init/grammer.y:325
		{
			/*
						function name, add to symbolTable
			                	(作用域？全局符号表第几级？)
			                	update：目前仍全局符号栈里了
			*/
			v2 := yyDollar[2].string_value
			v8 := (yyDollar[8].node).(*stmt.CompoundStmt)
			yyVAL.node = function.CreateFuncDefine(v2, v8)
			symbolTable.PopFrame()
		}
	case 53:
		yyDollar = yyS[yypt-0 : yypt+1]
//line src/init/grammer.y:339
		{
			yyVAL.node = nil
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
//line src/init/grammer.y:342
		{
			s := symbolTable.CreateSymbol(true, yyDollar[2].symbolType, yyDollar[3].string_value, 0)
			symbolTable.AddSymbol(s)
			v1 := stmt.CreateDeclareStmt(s, stmt.AssignStmt{})
			var declareStmts []*stmt.DeclareStmt
			yyVAL.node = append(declareStmts, &v1)
		}
	case 55:
		yyDollar = yyS[yypt-4 : yypt+1]
//line src/init/grammer.y:349
		{
			s := symbolTable.CreateSymbol(true, yyDollar[3].symbolType, yyDollar[4].string_value, 0)
			symbolTable.AddSymbol(s)
			v3 := stmt.CreateDeclareStmt(s, stmt.AssignStmt{})

			v1 := (yyDollar[1].node).([]*stmt.DeclareStmt)
			yyVAL.node = append(v1, &v3)
		}
	case 56:
		yyDollar = yyS[yypt-0 : yypt+1]
//line src/init/grammer.y:360
		{
			yyVAL.node = nil
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
//line src/init/grammer.y:363
		{
			var rv []*rvalue.RValue
			v2 := (yyDollar[2].node).(rvalue.RValue)
			yyVAL.node = append(rv, &v2)
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
//line src/init/grammer.y:368
		{
			v1 := (yyDollar[1].node).([]*rvalue.RValue)
			v3 := (yyDollar[3].node).(rvalue.RValue)
			yyVAL.node = append(v1, &v3)
			fmt.Println("Another callargs")
		}
	}
	goto yystack /* stack new state and value */
}
