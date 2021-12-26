package ast

import "LoxScanner/token"

type Function struct {
	Name   *token.Token
	Params []Expr
	Body   []Stmt
}

func (s *Function) Accept(v VisitorStmt) interface{} {
	return v.VisitFunctionStmt(s)
}
