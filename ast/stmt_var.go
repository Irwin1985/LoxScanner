package ast

import "LoxScanner/token"

type Var struct {
	Name        *token.Token
	Initializer Expr
}

func (s *Var) Accept(v VisitorStmt) interface{} {
	return v.VisitVarStmt(s)
}
