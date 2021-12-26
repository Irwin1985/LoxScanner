package ast

import "LoxScanner/token"

type Call struct {
	Callee    Expr
	Paren     *token.Token
	Arguments []Expr
}

func (e *Call) Accept(v VisitorExpr) interface{} {
	return v.VisitCallExpr(e)
}
