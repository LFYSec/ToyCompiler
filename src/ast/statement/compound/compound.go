package compound

import (
	stmt "Compiler/src/ast/statement"
)

var NextNamespaceId = 0

type CompoundStmt struct {
	stmt.Stmt
	ChildCount 	int
	Stmts 		[]stmt.Stmt
	NamespaceId	int
}

func (c CompoundStmt) GeneCode() {
	for _, s := range c.Stmts {
		s.GeneCode()
	}
}

func AddStmt(compound *CompoundStmt, s stmt.Stmt) {
	compound.Stmts = append(compound.Stmts, s)
	compound.ChildCount++
}

func CreateCompoundStmt() *CompoundStmt {
	var compoundStmt CompoundStmt
	compoundStmt.ChildCount = 0
	NextNamespaceId++
	compoundStmt.NamespaceId = NextNamespaceId
	return &compoundStmt
}