package ast

import "LoxScanner/token"

type Assign struct {
	Token *token.Token
	Value Expr
}

func (e *Assign) Accept(v VisitorExpr) interface{} {
	return v.VisitAssignExpr(e)
}
