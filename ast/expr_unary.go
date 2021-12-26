package ast

import "LoxScanner/token"

type Unary struct {
	Operator *token.Token
	Right    Expr
}

func (e *Unary) Accept(v VisitorExpr) interface{} {
	return v.VisitUnaryExpr(e)
}
