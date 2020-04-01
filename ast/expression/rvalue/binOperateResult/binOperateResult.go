package binOperateResult

import (
	rvalue2 "github.com/LFYSec/Compiler/ast/expression/rvalue"
)

type BinOperateResult struct {
	Base       rvalue2.RValue
	TmpRegId   int
	OperatorId int
	LHS        *rvalue2.RValue
	RHS        *rvalue2.RValue
}
