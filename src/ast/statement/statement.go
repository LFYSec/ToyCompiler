package stmt

import (
	"Compiler/src/ast"
)

type Stmt interface {
	ast.Node
	GeneCode()
}