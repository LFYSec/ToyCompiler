package lvalue

import "Compiler/src/ast/expression/rvalue"

type LValue interface {
	rvalue.RValue
	GeneLVCode()
	LvalueIR() 		string
}
