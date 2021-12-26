package ast

import "LoxScanner/token"

type Return struct {
	Keyword *token.Token
	Value   Expr
}

func (s *Return) Accept(v VisitorStmt) interface{} {
	return v.VisitReturnStmt(s)
}
