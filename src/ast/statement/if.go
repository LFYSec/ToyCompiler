package stmt

import (
	"Compiler/src/ast/expression/rvalue"
	"fmt"
)

var NextIfId = 0

type IfElseStmt struct {
	Stmt
	Condition 	rvalue.RValue
	IfStmt 		*CompoundStmt
	ElseStmt	*CompoundStmt
}

func (i IfElseStmt) GeneCode() {
	i.Condition.GeneRVCode()
	CondRVIR := i.Condition.RvalueIR()
	NextIfId++
	fmt.Printf("br i1 %s, label %%if_lable_%d, label %%else_lable_%d\n",
		CondRVIR,
		NextIfId,
		NextIfId,
	)
	fmt.Printf("if_lable_%d:\n", NextIfId)
	i.IfStmt.GeneCode()
	fmt.Printf("br label %%if_end_lable_%d\n", NextIfId)
	fmt.Printf("else_lable_%d:\n", NextIfId)
	if i.ElseStmt != nil {
		i.ElseStmt.GeneCode()
	}
	fmt.Printf("br label %%if_end_lable_%d\n", NextIfId)
	fmt.Printf("if_end_lable_%d:\n", NextIfId)
}

func CreateIfStmt(c rvalue.RValue, s1 *CompoundStmt, s2 *CompoundStmt) IfElseStmt {
	var ifStmt IfElseStmt
	ifStmt.Condition = c
	ifStmt.IfStmt = s1
	ifStmt.ElseStmt = s2
	return ifStmt
}