package ast

import "LoxScanner/token"

type Super struct {
	Keyword *token.Token
	Method  *token.Token
}

func (e *Super) Accept(v VisitorExpr) interface{} {
	return v.VisitSuperExpr(e)
}
