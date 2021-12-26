package ast

type Print struct {
	Expression Expr
}

func (s *Print) Accept(v VisitorStmt) interface{} {
	return v.VisitPrintStmt(s)
}
