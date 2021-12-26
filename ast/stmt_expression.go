package ast

type Expression struct {
	Expression Expr
}

func (s *Expression) Accept(v VisitorStmt) interface{} {
	return v.VisitExpressionStmt(s)
}
