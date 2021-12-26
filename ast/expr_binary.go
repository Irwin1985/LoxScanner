package ast

import "LoxScanner/token"

type Binary struct {
	Left     Expr
	Operator *token.Token
	Right    Expr
}

func (e *Binary) Accept(v VisitorExpr) interface{} {
	return v.VisitBinaryExpr(e)
}
