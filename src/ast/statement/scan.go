package stmt

import (
	"Compiler/src/ast/expression/lvalue/reference"
	"Compiler/src/ast/expression/rvalue"
	"Compiler/src/global"
	"fmt"
	"strconv"
)

type ScanStmt struct {
	Stmt
	Expr rvalue.RValue
}

func (s ScanStmt) GeneCode() {
	x := s.Expr.(reference.VariableReference)
	ir := "%" + x.Variable.Name + "_" + strconv.Itoa(x.Variable.NamespaceId)
	switch s.Expr.GetType() {
	case global.INT_LITERAL:

		fmt.Printf("call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([3 x i8], [3 x i8]* @.scan.int, i32 0, i32 0), i32* %s)\n", ir)
		break
	case global.DOUBLE_LITERAL:
		fmt.Printf("call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.scan.double, i32 0, i32 0), double* %s)\n", ir)
		break
	case global.CHAR_LITERAL:
		//p := ir + "_c"
		fmt.Printf("call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([3 x i8], [3 x i8]* @.scan, i32 0, i32 0), i8* %s)\n", ir)
		break
	default:
		break
	}
}

func CreateScanStmt(rhs rvalue.RValue) ScanStmt {
	var s ScanStmt
	s.Expr = rhs
	return s
}