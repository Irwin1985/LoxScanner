package ast

import "LoxScanner/token"

type Variable struct {
	Name *token.Token
}

func (e *Variable) Accept(v VisitorExpr) interface{} {
	return v.VisitVariableExpr(e)
}
