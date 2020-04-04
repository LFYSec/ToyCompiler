package print

import (
	"Compiler/src/ast/expression/rvalue"
	stmt "Compiler/src/ast/statement"
	"Compiler/src/global"
	"fmt"
)

type PrintStmt struct {
	stmt.Stmt
	Expr rvalue.RValue
}

func (p PrintStmt) GeneCode() {
	p.Expr.GeneRVCode()
	ir := p.Expr.RvalueIR()
	switch p.Expr.GetType() {
	case global.INT_LITERAL:
		fmt.Printf("call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @int_fmt_str, i32 0, i32 0), i32 %s)\n", ir)
		break
	case global.DOUBLE_LITERAL:
		fmt.Printf("call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @double_fmt_str, i32 0, i32 0), double %s)\n", ir)
		break
	default:
		break
	}
}

func CreatePrintStmt(rhs rvalue.RValue) PrintStmt {
	var p PrintStmt
	p.Expr = rhs
	return p
}