package stmt

import (
	"Compiler/src/ast/expression/rvalue"
	"Compiler/src/symbolTable"
	"fmt"
)

type RetStmt struct {
	Stmt
	Expr rvalue.RValue
}

func (p RetStmt) GeneCode() {
	p.Expr.GeneRVCode()
	ir := p.Expr.RvalueIR()
	fmt.Printf("\tret %s %s\n", symbolTable.TypeString(p.Expr.GetType()), ir)
}

func CreateRetStmt(rhs rvalue.RValue) RetStmt {
	var p RetStmt
	p.Expr = rhs
	return p
}
