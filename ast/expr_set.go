package ast

import "LoxScanner/token"

type Set struct {
	Object Expr
	Name   *token.Token
	Value  Expr
}

func (e *Set) Accept(v VisitorExpr) interface{} {
	return v.VisitSetExpr(e)
}
