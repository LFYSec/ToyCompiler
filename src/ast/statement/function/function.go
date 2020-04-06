package function

import (
	stmt "Compiler/src/ast/statement"
	"Compiler/src/ast/statement/compound"
	"Compiler/src/ast/statement/declare"
	"Compiler/src/global"
	"Compiler/src/symbolTable"
	"Compiler/src/symbolTable/symbol"
	"fmt"
)

type Function struct {
	stmt.Stmt
	FuncSB				symbol.FuncSymbol
	CompoundStmts 	*compound.CompoundStmt
}

func (f Function) GeneCode() {
	var args = ""
	for k, v := range f.FuncSB.ArgsType {
		args += fmt.Sprintf("%s %%%s,", symbolTable.TypeString(v), k)
	}
	if args != "" {
		args = args[0 : len(args)-1]
	}
	fmt.Printf("define %s @%s(%s) #0 {\n",
			symbolTable.TypeString(f.FuncSB.ReturnType),
			f.FuncSB.Name,
			args,
		)
	for k, v := range f.FuncSB.ArgsType {
		fmt.Printf("%%%s = alloca %s, align 4\n", k[4:], symbolTable.TypeString(v))
		fmt.Printf("store %s %%%s, %s* %%%s, align 4\n", symbolTable.TypeString(v), k, symbolTable.TypeString(v), k[4:])
	}
	f.CompoundStmts.GeneCode()
	fmt.Printf("ret void\n}\n") // For now, don't achieve return value. TODO
}

func CreateFuncDefine(t global.SymbolType, name string,
						d []*declare.DeclareStmt, c *compound.CompoundStmt) Function {

	//var funcSymbol symbol.FuncSymbol
	funcSymbol := symbol.FuncSymbol{
		Name:       "",
		ReturnType: 0,
		ArgsType:   make(map[string]global.SymbolType),
	}
	funcSymbol.Name = name
	funcSymbol.ReturnType = t
	for _, i := range d {
		k := fmt.Sprintf("arg_%s_%d",i.Variable.Name, i.Variable.NamespaceId)
		funcSymbol.ArgsType[k] = i.Variable.Type
	}
	symbolTable.AddFunc(&funcSymbol)

	var function Function
	function.FuncSB = funcSymbol
	function.CompoundStmts = c
	return function
}