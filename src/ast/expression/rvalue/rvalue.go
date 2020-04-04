package rvalue

import (
	"Compiler/src/ast/expression"
	"Compiler/src/global"
)

var NextTmpReg = 0

type RValue interface {
	expression.Expression
	GeneRVCode()
	RvalueIR() 		string

	GetType()		global.SymbolType
}