package function

import (
	stmt "Compiler/src/ast/statement"
	"Compiler/src/global"
	st "Compiler/src/symbolTable"
	"fmt"
)

type Function struct {
	stmt.Stmt
	FuncSymb		*st.FuncSymbol
	CompoundStmts 	*stmt.CompoundStmt
}

func (f Function) GeneCode() {
	var args = ""
	for k, v := range f.FuncSymb.ArgsType {
		args += fmt.Sprintf("%s %%%s,", st.TypeString(v), k)
	}
	if args != "" {
		args = args[0 : len(args)-1]
	}
	fmt.Printf("define %s @%s(%s) #0 {\n",
			st.TypeString(f.FuncSymb.ReturnType),
			f.FuncSymb.Name,
			args,
		)
	for k, v := range f.FuncSymb.ArgsType {
		fmt.Printf("%%%s = alloca %s, align 4\n", k[4:], st.TypeString(v))
		fmt.Printf("store %s %%%s, %s* %%%s, align 4\n", st.TypeString(v), k, st.TypeString(v), k[4:])
	}
	f.CompoundStmts.GeneCode()
	fmt.Printf("}\n")
}

func AddFunc(t global.SymbolType, name string, d []*stmt.DeclareStmt) {
	funcSymbol := st.FuncSymbol{
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
	st.AddFunc(&funcSymbol)
}

func CreateFuncDefine(name string, c *stmt.CompoundStmt) Function {
	funcSymbol := st.GetFuncSymbol(name)

	var function Function
	function.FuncSymb = funcSymbol
	function.CompoundStmts = c
	return function
}