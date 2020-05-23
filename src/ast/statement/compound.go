package stmt

var NextNamespaceId = 0

type CompoundStmt struct {
	Stmt
	ChildCount 	int
	Stmts 		[]Stmt
	NamespaceId	int
}

func (c CompoundStmt) GeneCode() {
	for _, s := range c.Stmts {
		s.GeneCode()
	}
}

func AddStmt(compound *CompoundStmt, s Stmt) {
	compound.Stmts = append(compound.Stmts, s)
	compound.ChildCount++
}

func CreateCompoundStmt() *CompoundStmt {
	var compoundStmt CompoundStmt
	compoundStmt.ChildCount = 0
	NextNamespaceId++
	compoundStmt.NamespaceId = NextNamespaceId
	return &compoundStmt
}