package intLiteral

import (
	"fmt"
	rvalue2 "github.com/LFYSec/Compiler/ast/expression/rvalue"
)

type IntLiteral struct {
	Base  rvalue2.RValue
	Value int
}


func CreateIntLiteral(value int) *IntLiteral {
	var intLiteral IntLiteral
	intLiteral.Value = value
	return &intLiteral
}

func (i *IntLiteral) geneRVCode() {

}

func (i *IntLiteral) rvalueIR() string {
	result := fmt.Sprintf("%d", i.Value)
	return result
}