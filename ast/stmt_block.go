package ast

type Block struct {
	Statements []Stmt
}

func (s *Block) Accept(v VisitorStmt) interface{} {
	return v.VisitBlockStmt(s)
}
