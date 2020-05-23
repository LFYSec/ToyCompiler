package symbolTable

import (
	"Compiler/src/global"
	"fmt"
)

type Symbol struct {
	Type 			global.SymbolType
	Mutable 		bool
	Name			string
	NamespaceId		int64
	Size			int
}

type FuncSymbol struct {
	Name 			string
	ReturnType	 	global.SymbolType
	ArgsType 		map[string]global.SymbolType
}

func TypeName(symbol *Symbol) string { //  TODO distinguish TypeName and TypeString
	if symbol.Size > 0 {
		result := fmt.Sprintf("[%d x %s]", symbol.Size, TypeString(symbol.Type))
		return result
	}
	switch symbol.Type {
	case global.INT_LITERAL:
		return "i32"
	case global.DOUBLE_LITERAL:
		return "double"
	default:
		panic("[*]Error: Unknown type")
	}
}

func CreateSymbol(mutable bool, t global.SymbolType, name string, size int) *Symbol {
	var result Symbol
	result.Mutable = mutable
	result.Name = name
	if t == global.INT {
		t = global.INT_LITERAL
	} else if t == global.DOUBLE {
		t = global.DOUBLE_LITERAL
	} else if t == global.BOOL {
		t = global.BOOL
	}
	result.Type = t
	//result.NamespaceId = 2^64 - 1  // TODO Question?
	result.NamespaceId = 0  // TODO Question?
	result.Size = size
	return &result
}

