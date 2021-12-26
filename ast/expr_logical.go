package ast

import "LoxScanner/token"

type Logical struct {
	Left     Expr
	Operator *token.Token
	Right    Expr
}

func (e *Logical) Accept(v VisitorExpr) interface{} {
	return v.VisitLogicalExpr(e)
}
