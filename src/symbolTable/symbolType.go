package symbolTable

import (
	"Compiler/src/global"
)

func TypeString(t global.SymbolType) string {
	switch t {
	case global.VOID:
		return "void"
	case global.BOOL:
		return "i1"
	case global.INT:
		return "i32"
	case global.INT_LITERAL:
		return "i32"
	case global.DOUBLE:
		return "double"
	case global.DOUBLE_LITERAL:
		return "double"
	case global.CHAR_LITERAL:
		return "i8"
	default:
		return ""
	}
}
