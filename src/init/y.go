// Code generated by goyacc -o src/init/y.go src/init/grammer.y. DO NOT EDIT.

//line src/init/grammer.y:2

package init

import __yyfmt__ "fmt"

//line src/init/grammer.y:3
import (
	"Compiler/src/ast"
	"Compiler/src/ast/expression/lvalue"
	"Compiler/src/ast/expression/lvalue/variableReference"
	"Compiler/src/ast/expression/rvalue"
	"Compiler/src/ast/expression/rvalue/binOperateResult"
	"Compiler/src/ast/expression/rvalue/doubleLiteral"
	"Compiler/src/ast/expression/rvalue/intLiteral"
	stmt "Compiler/src/ast/statement"
	"Compiler/src/ast/statement/assign"
	"Compiler/src/ast/statement/compound"
	"Compiler/src/ast/statement/declare"
	"Compiler/src/ast/statement/print"
	"Compiler/src/global"
	"Compiler/src/symbolTable"
	"Compiler/src/symbolTable/symbol"
)

var result *compound.CompoundStmt

//line src/init/grammer.y:31
type yySymType struct {
	yys          int
	int_value    int
	double_value float64
	string_value string
	node         ast.Node
	symbolType   global.SymbolType
}

const TYPE = 57346
const IDENTIFY = 57347
const INT_LITERAL = 57348
const DOUBLE_LITERAL = 57349
const ARRAY = 57350
const STRING = 57351
const INT = 57352
const DOUBLE = 57353
const CHAR = 57354
const BOOL = 57355
const IF = 57356
const ELSE = 57357
const WHILE = 57358
const FOR = 57359
const PRINT = 57360
const EOL = 57361
const EQ = 57362
const NE = 57363
const LT = 57364
const LE = 57365
const GT = 57366
const GE = 57367
const ASSIGN = 57368
const ADD = 57369
const SUB = 57370
const MUL = 57371
const DIV = 57372
const NOT = 57373
const LBRACKET = 57374
const RBRACKET = 57375
const LPARENTHESIS = 57376
const RPARENTHESIS = 57377

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"TYPE",
	"IDENTIFY",
	"INT_LITERAL",
	"DOUBLE_LITERAL",
	"ARRAY",
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
	"LPARENTHESIS",
	"RPARENTHESIS",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line src/init/grammer.y:169

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 40

var yyAct = [...]int{

	9, 20, 21, 9, 20, 21, 34, 32, 13, 17,
	29, 12, 10, 15, 8, 9, 18, 14, 11, 22,
	19, 6, 24, 25, 5, 4, 26, 27, 7, 23,
	3, 16, 23, 28, 1, 2, 30, 31, 0, 33,
}
var yyPact = [...]int{

	-1000, 10, -7, -1000, -1000, -1000, -15, -26, 12, -1000,
	-1000, -1000, -5, -5, -15, -1000, -17, -1000, -1000, -2,
	-1000, -1000, -1000, -5, -1000, -1000, -1000, -28, -1000, -5,
	-1000, -29, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 13, 35, 34, 18, 16, 9, 31, 19, 30,
	25, 24, 20,
}
var yyR1 = [...]int{

	0, 3, 3, 4, 11, 11, 9, 10, 2, 2,
	2, 5, 5, 5, 5, 8, 12, 12, 12, 6,
	6, 7, 7, 1,
}
var yyR2 = [...]int{

	0, 0, 3, 2, 2, 3, 2, 4, 1, 1,
	1, 1, 1, 1, 3, 1, 1, 1, 1, 1,
	2, 1, 3, 1,
}
var yyChk = [...]int{

	-1000, -3, -2, -9, -10, -11, -8, 18, 4, 5,
	19, -4, 26, 34, 5, -1, -7, -6, -5, -12,
	6, 7, -8, 34, 27, 28, 31, -1, -4, 27,
	-5, -1, 35, -6, 35,
}
var yyDef = [...]int{

	1, -2, 0, 8, 9, 10, 0, 0, 0, 15,
	2, 6, 0, 0, 4, 3, 23, 21, 19, 0,
	11, 12, 13, 0, 16, 17, 18, 0, 5, 0,
	20, 0, 7, 22, 14,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35,
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
//line src/init/grammer.y:66
		{
			result = compound.CreateCompoundStmt()
			yyVAL.node = result
		}
	case 2:
		yyDollar = yyS[yypt-3 : yypt+1]
//line src/init/grammer.y:70
		{
			v1 := (yyDollar[1].node).(*compound.CompoundStmt)
			v2 := (yyDollar[2].node).(stmt.Stmt)
			compound.AddStmt(v1, v2)
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
//line src/init/grammer.y:78
		{
			yyVAL.node = yyDollar[2].node
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
//line src/init/grammer.y:84
		{
			s := symbol.CreateSymbol(true, yyDollar[1].symbolType, yyDollar[2].string_value)
			symbolTable.AddSymbol(s)
			yyVAL.node = declare.CreateDeclareStmt(s, assign.AssignStmt{})
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
//line src/init/grammer.y:89
		{
			s := symbol.CreateSymbol(true, yyDollar[1].symbolType, yyDollar[2].string_value)
			symbolTable.AddSymbol(s)
			ref := variableReference.CreateVariableReference(s)
			v := (yyDollar[3].node).(rvalue.RValue)
			a := assign.CreateAssignStmt(ref, v)
			yyVAL.node = declare.CreateDeclareStmt(s, a)
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
//line src/init/grammer.y:100
		{
			v1 := (yyDollar[1].node).(lvalue.LValue)
			v2 := (yyDollar[2].node).(rvalue.RValue)
			yyVAL.node = assign.CreateAssignStmt(v1, v2)
		}
	case 7:
		yyDollar = yyS[yypt-4 : yypt+1]
//line src/init/grammer.y:108
		{
			v3 := (yyDollar[3].node).(rvalue.RValue)
			yyVAL.node = print.CreatePrintStmt(v3)
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
//line src/init/grammer.y:122
		{
			yyVAL.node = intLiteral.CreateIntLiteral(yyDollar[1].int_value)
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line src/init/grammer.y:125
		{
			yyVAL.node = doubleLiteral.CreateDoubleLiteral(yyDollar[1].double_value)
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
//line src/init/grammer.y:128
		{
			yyVAL.node = yyDollar[1].node
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
//line src/init/grammer.y:131
		{
			yyVAL.node = yyDollar[2].node
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
//line src/init/grammer.y:137
		{
			sb := symbolTable.GetSymbol(yyDollar[1].string_value)
			yyVAL.node = variableReference.CreateVariableReference(sb)
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
//line src/init/grammer.y:149
		{
			yyVAL.node = yyDollar[1].node
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
//line src/init/grammer.y:150
		{
			yyVAL.node = yyDollar[2].node
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
//line src/init/grammer.y:154
		{
			yyVAL.node = yyDollar[1].node
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
//line src/init/grammer.y:155
		{
			v1 := (yyDollar[1].node).(rvalue.RValue)
			v3 := (yyDollar[3].node).(rvalue.RValue)
			yyVAL.node = binOperateResult.CreateBinOperateResult(ADD, v1, v3)
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
//line src/init/grammer.y:163
		{
			yyVAL.node = yyDollar[1].node
		}
	}
	goto yystack /* stack new state and value */
}