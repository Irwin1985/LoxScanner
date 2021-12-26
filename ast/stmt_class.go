package ast

import "LoxScanner/token"

type Class struct {
	Name       *token.Token
	SuperClass Variable
	Methods    []Function
}

func (s *Class) Accept(v VisitorStmt) interface{} {
	return v.VisitClassStmt(s)
}
