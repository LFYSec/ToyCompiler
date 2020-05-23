package stmt

import (
	"Compiler/src/symbolTable"
	"fmt"
	"reflect"
)

type DeclareStmt struct {
	Stmt
	Initial  	AssignStmt
	Variable 	*symbolTable.Symbol
}

func (d DeclareStmt) GeneCode() {
	fmt.Printf("%%%s_%d = alloca %s\n", d.Variable.Name, d.Variable.NamespaceId, symbolTable.TypeName(d.Variable))
	if !reflect.DeepEqual(d.Initial, AssignStmt{}) {
		d.Initial.GeneCode()
	}
}

func CreateDeclareStmt(s *symbolTable.Symbol, i AssignStmt) DeclareStmt {
	var declareStmt DeclareStmt
	declareStmt.Variable = s
	declareStmt.Initial = i
	return declareStmt
}