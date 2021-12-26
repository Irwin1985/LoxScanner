package ast

type VisitorStmt interface {
	VisitBlockStmt(stmt *Block) interface{}
	VisitClassStmt(stmt *Class) interface{}
	VisitExpressionStmt(stmt *Expression) interface{}
	VisitFunctionStmt(stmt *Function) interface{}
	VisitIfStmt(stmt *If) interface{}
	VisitPrintStmt(stmt *Print) interface{}
	VisitReturnStmt(stmt *Return) interface{}
	VisitVarStmt(stmt *Var) interface{}
	VisitWhileStmt(stmt *While) interface{}
}

type Stmt interface {
	Accept(v VisitorStmt) interface{}
}
