package while

import (
	"Compiler/src/ast/expression/rvalue"
	stmt "Compiler/src/ast/statement"
	"Compiler/src/ast/statement/compound"
	"fmt"
)

var NextWhileId = 0

type WhileStmt struct {
	stmt.Stmt
	Condition 	rvalue.RValue
	CStmts 	*compound.CompoundStmt
}

func (w WhileStmt) GeneCode() {
	NextWhileId++
	fmt.Printf("br label %%while_condition_%d\n", NextWhileId)
	fmt.Printf("while_condition_%d:\n", NextWhileId)
	w.Condition.GeneRVCode()
	CRVIR := w.Condition.RvalueIR()
	fmt.Printf("br i1 %s, label %%while_body_%d,label %%while_end_%d\n",
			CRVIR,
			NextWhileId,
			NextWhileId,
		)
	fmt.Printf("while_body_%d:\n", NextWhileId)
	w.CStmts.GeneCode()
	fmt.Printf("br label %%while_condition_%d\n", NextWhileId)
	fmt.Printf("while_end_%d:\n", NextWhileId)
}

func CreateWhileStmt(c rvalue.RValue, s *compound.CompoundStmt) WhileStmt {
	var whileStmt WhileStmt
	whileStmt.Condition = c
	whileStmt.CStmts = s
	return whileStmt
}