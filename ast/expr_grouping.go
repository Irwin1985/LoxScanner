package ast

type Grouping struct {
	Expression Expr
}

func (e *Grouping) Accept(v VisitorExpr) interface{} {
	return v.VisitGroupingExpr(e)
}
