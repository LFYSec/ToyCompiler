package stmt

import (
	"Compiler/src/ast/expression/rvalue"
	"Compiler/src/global"
	"fmt"
)

type PrintStmt struct {
	Stmt
	Expr rvalue.RValue
}

func (p PrintStmt) GeneCode() {
	p.Expr.GeneRVCode()
	ir := p.Expr.RvalueIR()
	switch p.Expr.GetType() {
	case global.INT_LITERAL:
		fmt.Printf("call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.str.int, i32 0, i32 0), i32 %s)\n", ir)
		break
	case global.DOUBLE_LITERAL:
		fmt.Printf("call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.str.double, i32 0, i32 0), double %s)\n", ir)
		break
	case global.CHAR_LITERAL:
		p := ir + "_c"
		fmt.Printf("%s = sext i8 %s to i32\n", p, ir)
		fmt.Printf("call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.str, i32 0, i32 0), i32 %s)\n", p)
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