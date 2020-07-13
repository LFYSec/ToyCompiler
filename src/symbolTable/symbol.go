package symbolTable

import (
	"Compiler/src/global"
	"fmt"
)

type Symbol struct {
	Type 			global.SymbolType
	Name			string
	NamespaceId		int
	Size			int
}

type FuncSymbol struct {
	Name 			string
	ReturnType	 	global.SymbolType
	ArgsType 		map[string]global.SymbolType
}

func TypeName(symbol *Symbol) string {
	if symbol.Size > 0 {
		result := fmt.Sprintf("[%d x %s]", symbol.Size, TypeString(symbol.Type))
		return result
	}
	switch symbol.Type {
	case global.INT_LITERAL:
		return "i32"
	case global.DOUBLE_LITERAL:
		return "double"
	case global.CHAR_LITERAL:
		return "i8"
	default:
		panic("[*]Error: Unknown type")
	}
}

func CreateSymbol(mutable bool, t global.SymbolType, name string, size int) *Symbol {
	var result Symbol
	result.Name = name
	if t == global.INT {
		t = global.INT_LITERAL
	} else if t == global.DOUBLE {
		t = global.DOUBLE_LITERAL
	} else if t == global.BOOL {
		t = global.BOOL
	} else if t == global.CHAR {
		t = global.CHAR_LITERAL
	}

	result.Type = t
	//result.NamespaceId = 2^64 - 1
	result.NamespaceId = 0
	result.Size = size
	return &result
}

