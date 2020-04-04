package declare

import (
	stmt "Compiler/src/ast/statement"
	"Compiler/src/ast/statement/assign"
	"Compiler/src/symbolTable/symbol"
	"fmt"
	"reflect"
)

type DeclareStmt struct {
	stmt.Stmt
	Initial 	assign.AssignStmt
	Variable	*symbol.Symbol
}

func (d DeclareStmt) GeneCode() {
	fmt.Printf("%%%s_%d = alloca %s\n", d.Variable.Name, d.Variable.NamespaceId, symbol.TypeName(d.Variable))
	if !reflect.DeepEqual(d.Initial, assign.AssignStmt{}) {
		d.Initial.GeneCode()
	}
}

func CreateDeclareStmt(s *symbol.Symbol, i assign.AssignStmt) DeclareStmt {
	var declareStmt DeclareStmt
	declareStmt.Variable = s
	declareStmt.Initial = i
	return declareStmt
}