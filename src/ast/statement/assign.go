package stmt

import (
	"Compiler/src/ast/expression/lvalue"
	"Compiler/src/ast/expression/rvalue"
	"Compiler/src/symbolTable"
	"fmt"
)

type AssignStmt struct {
	Stmt
	LHS 		lvalue.LValue
	RHS 		rvalue.RValue
}

func (a AssignStmt) GeneCode() {
	a.LHS.GeneLVCode()
	a.RHS.GeneRVCode()
	RVIR := a.RHS.RvalueIR()
	LVIR := a.LHS.LvalueIR()
	fmt.Printf("store %s %s, %s* %s\n",
		symbolTable.TypeString(a.RHS.GetType()), // TODO fix array.GetType()
		RVIR,
		symbolTable.TypeString(a.RHS.GetType()),
		LVIR,
	)
}

func CreateAssignStmt(lhs lvalue.LValue, rhs rvalue.RValue) AssignStmt {
	var assignStmt AssignStmt
	assignStmt.LHS = lhs
	assignStmt.RHS = rhs
	return assignStmt
}