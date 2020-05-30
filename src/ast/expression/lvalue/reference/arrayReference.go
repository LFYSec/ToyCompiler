package reference

import (
	"Compiler/src/ast/expression/lvalue"
	"Compiler/src/ast/expression/rvalue"
	"Compiler/src/global"
	"Compiler/src/symbolTable"
	"fmt"
)

type ArrayReference struct {
	lvalue.LValue
	Type 		global.SymbolType
	Index 		rvalue.RValue
	Array		*symbolTable.Symbol
	RvRegId 	int
	LvRegId 	int
}

func (a ArrayReference) GetType() global.SymbolType {
	return a.Type
}

func (a ArrayReference) GeneLVCode()  {
	LVIR := a.LvalueIR()
	a.Index.GeneRVCode()

	IndexRVIR := a.Index.RvalueIR()
	fmt.Printf("%s = getelementptr inbounds [%d x %s], [%d x %s]* %%%s_%d, i64 0, i32 %s\n",
		LVIR,
		a.Array.Size,
		symbolTable.TypeString(a.Array.Type),
		a.Array.Size,
		symbolTable.TypeString(a.Array.Type),
		a.Array.Name,
		a.Array.NamespaceId,
		IndexRVIR,
	)
}

func (a ArrayReference) GeneRVCode()  {
	RVIR := a.RvalueIR()
	LVIR := a.LvalueIR()
	a.GeneLVCode()
	fmt.Printf("%s = load %s, %s* %s\n",
		RVIR,
		symbolTable.TypeString(a.Array.Type),
		symbolTable.TypeString(a.Array.Type),
		LVIR,
	)
}

func (a ArrayReference) LvalueIR() string {
	result := fmt.Sprintf("%%%s_element_%d", a.Array.Name, a.LvRegId)
	return result
}

func (a ArrayReference) RvalueIR() string {
	result := fmt.Sprintf("%%temp_%d", a.RvRegId)
	return result
}

func CreateArrayReference(array *symbolTable.Symbol, index rvalue.RValue) ArrayReference {
	if array == nil {
		panic("[*] Error: CreateArrayReference has nil symbol args")
	}
	var result ArrayReference
	result.Array = array
	result.Index = index

	rvalue.NextTmpReg++
	result.RvRegId = rvalue.NextTmpReg

	rvalue.NextTmpReg++
	result.LvRegId = rvalue.NextTmpReg
	result.Type = array.Type
	return result
}