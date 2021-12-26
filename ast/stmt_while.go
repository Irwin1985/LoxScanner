package ast

type While struct {
	Condition Expr
	Body      Stmt
}

func (s *While) Accept(v VisitorStmt) interface{} {
	return v.VisitWhileStmt(s)
}
