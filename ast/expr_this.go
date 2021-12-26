package ast

import "LoxScanner/token"

type This struct {
	Keyword *token.Token
}

func (e *This) Accept(v VisitorExpr) interface{} {
	return v.VisitThisExpr(e)
}
