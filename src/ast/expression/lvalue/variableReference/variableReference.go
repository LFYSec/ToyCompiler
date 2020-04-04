package variableReference

import (
	"Compiler/src/ast/expression/lvalue"
	"Compiler/src/ast/expression/rvalue"
	"Compiler/src/global"
	"Compiler/src/symbolTable/symbol"
	"fmt"
)

type VariableReference struct {
	lvalue.LValue
	Type 		global.SymbolType
	TmpRegId 	int
	Variable	*symbol.Symbol
}

func (v VariableReference) GetType() global.SymbolType {
	return v.Type
}

func (v VariableReference) GeneLVCode()  {}

func (v VariableReference) GeneRVCode()  {
	RVIR := v.RvalueIR()
	LVIR := v.LvalueIR()
	fmt.Printf("%s = load %s, %s* %s\n",
		RVIR,
		symbol.TypeName(v.Variable),
		symbol.TypeName(v.Variable),
		LVIR,
	)
}

func (v VariableReference) LvalueIR() string {
	result := fmt.Sprintf("%%%s_%d", v.Variable.Name, v.Variable.NamespaceId)
	return result
}

func (v VariableReference) RvalueIR() string {
	result := fmt.Sprintf("%%temp_%d", v.TmpRegId)
	return result
}


func CreateVariableReference(variable *symbol.Symbol) VariableReference {
	if variable == nil {
		panic("[*] Error: CreateVariableReference has nil symbol args")
	}
	var result VariableReference
	result.Variable = variable
	rvalue.NextTmpReg++
	result.TmpRegId = rvalue.NextTmpReg
	result.Type = variable.Type
	return result
}