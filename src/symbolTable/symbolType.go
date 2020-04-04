package symbolTable

import (
	"Compiler/src/global"
)

func TypeString(t global.SymbolType) string {
	switch t {
	case global.BOOL:
		return "i1"
	case global.INT_LITERAL:
		return "i32"
	case global.DOUBLE_LITERAL:
		return "double"
	default:
		return ""
	}
}
