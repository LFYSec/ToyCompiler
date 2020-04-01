package doubleLiteral

import (
	"fmt"
	rvalue2 "github.com/LFYSec/Compiler/ast/expression/rvalue"
)

type DoubleLiteral struct {
	Base  rvalue2.RValue
	Value float64
}


func CreateDoubleLiteral() *DoubleLiteral {

	return nil
}

func (i *DoubleLiteral) geneRVCode() {

}

func (i *DoubleLiteral) rvalueIR() string {
	result := fmt.Sprintf("%64.32f", i.Value)
	return result
}