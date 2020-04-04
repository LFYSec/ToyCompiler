package symbol

import (
	"Compiler/src/global"
	"os"
)

type Symbol struct {
	Type 		global.SymbolType
	Mutable 	bool
	Name		string
	NamespaceId	int64
}

func TypeName(symbol *Symbol) string {
	switch symbol.Type {
	case global.INT_LITERAL:
		return "i32"
	case global.DOUBLE_LITERAL:
		return "double"
	case global.ARRAY:
		// TODO result := fmt.Sprintf("[%d x %s]",)
		return "array"
	default:
		os.Exit(1)
		return ""
	}
}

func CreateSymbol(mutable bool, t global.SymbolType, name string) *Symbol {
	var result Symbol
	result.Mutable = mutable
	result.Name = name
	if t == global.INT {
		t = global.INT_LITERAL
	} else if t == global.DOUBLE {
		t = global.DOUBLE_LITERAL
	}
	result.Type = t
	result.NamespaceId = 2^64 - 1
	return &result
}

