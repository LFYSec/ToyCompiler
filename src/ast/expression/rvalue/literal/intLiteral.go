package literal

import (
	"Compiler/src/ast/expression/rvalue"
	"Compiler/src/global"
	"fmt"
)

type IntLiteral struct {
	rvalue.RValue
	Type 	global.SymbolType
	Value 	int
}

func (i IntLiteral) GeneRVCode() {

}

func (i IntLiteral) RvalueIR() string {
	result := fmt.Sprintf("%d", i.Value)
	return result
}

func (i IntLiteral) GetType() global.SymbolType {
	return i.Type
}

func CreateIntLiteral(value int) IntLiteral {
	var intLiteral IntLiteral
	intLiteral.Value = value
	intLiteral.Type = global.INT_LITERAL
	return intLiteral
}