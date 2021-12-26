package ast

import "LoxScanner/token"

type Literal struct {
	Value token.Object
}

func (e *Literal) Accept(v VisitorExpr) interface{} {
	return v.VisitLiteralExpr(e)
}
