package literal

import (
	"Compiler/src/ast/expression/rvalue"
	"Compiler/src/global"
	"fmt"
)

type CharLiteral struct {
	rvalue.RValue
	Type 	global.SymbolType
	Value 	uint8
}


func (c CharLiteral) GeneRVCode() {

}

func (c CharLiteral) RvalueIR() string {
	result := fmt.Sprintf("%d", c.Value)
	return result
}

func (c CharLiteral) GetType() global.SymbolType {
	return c.Type
}

func CreateCharLiteral(value uint8) CharLiteral {
	var charLiteral CharLiteral
	charLiteral.Value = value
	charLiteral.Type = global.CHAR_LITERAL
	return charLiteral
}