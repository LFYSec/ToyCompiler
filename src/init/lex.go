package init

/*
	#include "lex.yy.h"
	#include "token.h"
*/
import "C"

import (
	"errors"
	"log"
	"strconv"
)

type lex struct {
	yytext  string
	lastErr error
}

var _ yyLexer = (*lex)(nil)

func newLexer(data []byte) *lex {
	p := new(lex)

	C.yy_scan_bytes(
		(*C.char)(C.CBytes(data)),
		C.yy_size_t(len(data)),
	)
	return p
}

func (l *lex) Lex(yylval *yySymType) int {
	l.lastErr = nil
	token := C.yylex() // 读取一个token

	l.yytext = C.GoString(C.yytext)

	//fmt.Println("yytext: ", l.yytext)

	switch token {
	case C.CHAR:
		yylval.symbolType = CHAR
		return TYPE
	case C.BOOL:
		yylval.symbolType = BOOL
		return TYPE
	case C.INT:
		yylval.symbolType = INT
		return TYPE
	case C.DOUBLE:
		yylval.symbolType = DOUBLE
		return TYPE
	case C.STRING:
		yylval.symbolType = STRING
		return TYPE
	case C.VOID:
		yylval.symbolType = VOID
		return TYPE
	case C.IF:
		return IF
	case C.ELSE:
		return ELSE
	case C.WHILE:
		return WHILE
	case C.RETURN:
		return RETURN
	case C.PRINT:
		return PRINT
	case C.SCAN:
		return SCAN
	case C.INT_LITERAL:
		yylval.int_value, _ = strconv.Atoi(l.yytext)
		return INT_LITERAL
	case C.DOUBLE_LITERAL:
		yylval.double_value, _ = strconv.ParseFloat(l.yytext, 64)
		return DOUBLE_LITERAL
	case C.CHAR_LITERAL:
		tmp := []rune(l.yytext[1:len(l.yytext)-1])
		yylval.char_value = uint8(tmp[0])
		return CHAR_LITERAL
	case C.IDENTIFY:
		yylval.string_value = l.yytext
		return IDENTIFY
	case C.EQ:
		yylval.string_value = l.yytext
		return EQ
	case C.NE:
		yylval.string_value = l.yytext
		return NE
	case C.LT:
		yylval.string_value = l.yytext
		return LT
	case C.LE:
		yylval.string_value = l.yytext
		return LE
	case C.GT:
		yylval.string_value = l.yytext
		return GT
	case C.GE:
		yylval.string_value = l.yytext
		return GE
	case C.ADD:
		yylval.string_value = l.yytext
		return ADD
	case C.SUB:
		yylval.string_value = l.yytext
		return SUB
	case C.MUL:
		yylval.string_value = l.yytext
		return MUL
	case C.DIV:
		yylval.string_value = l.yytext
		return DIV
	case C.NOT:
		yylval.string_value = l.yytext
		return NOT
	case C.LBRACKET:
		yylval.string_value = l.yytext
		return LBRACKET
	case C.RBRACKET:
		yylval.string_value = l.yytext
		return RBRACKET
	case C.LPARENTHESIS:
		yylval.string_value = l.yytext
		return LPARENTHESIS
	case C.RPARENTHESIS:
		yylval.string_value = l.yytext
		return RPARENTHESIS
	case C.LBRACE:
		yylval.string_value = l.yytext
		return LBRACE
	case C.RBRACE:
		yylval.string_value = l.yytext
		return RBRACE
	case C.COMMA:
		yylval.string_value = l.yytext
		return COMMA
	case C.EOL:
		return EOL
	case C.EOS:
		return EOS
	case C.ASSIGN:
		yylval.string_value = l.yytext
		return ASSIGN

	}
	return 0
}

func (l *lex) Error(s string) {
	l.lastErr = errors.New("yacc: " + s)
	if err := l.lastErr; err != nil {
		log.Println(err)
	}
}
