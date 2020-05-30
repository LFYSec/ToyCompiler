package doubleLiteral

import (
	"Compiler/src/ast/expression/rvalue"
	"Compiler/src/global"
	"fmt"
)

type DoubleLiteral struct {
	rvalue.RValue
	Type 	global.SymbolType
	Value 	float64
}

func (d DoubleLiteral) GeneRVCode() {

}

func (d DoubleLiteral) RvalueIR() string {
	result := fmt.Sprintf("%64.32f", d.Value)
	return result
}

func (d DoubleLiteral) GetType() global.SymbolType {
	return d.Type
}


func CreateDoubleLiteral(value float64) DoubleLiteral {
	var doubleLiteral DoubleLiteral
	doubleLiteral.Value = value
	doubleLiteral.Type = global.DOUBLE_LITERAL
	return doubleLiteral
}
