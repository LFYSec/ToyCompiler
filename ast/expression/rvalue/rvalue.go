package rvalue

import (
	"github.com/LFYSec/Compiler/ast/expression"
)

type RValue struct {
	Base expression.Expression
	Type int		// TODO
}

func (rv *RValue) geneRVCode() {

}

func (rv *RValue) rvalueIR() {

}
