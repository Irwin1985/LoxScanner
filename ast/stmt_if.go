package ast

type If struct {
	Condition  Expr
	ThenBranch Stmt
	ElseBranch Stmt
}

func (s *If) Accept(v VisitorStmt) interface{} {
	return v.VisitIfStmt(s)
}
