package ast

import "LoxScanner/token"

type Get struct {
	Object Expr
	Name   *token.Token
}

func (e *Get) Accept(v VisitorExpr) interface{} {
	return v.VisitGetExpr(e)
}
