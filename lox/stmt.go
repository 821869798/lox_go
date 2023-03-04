package lox

type Stmt interface {
}

type Expression struct{
	expression Expr
}

func NewExpression(expression Expr)*Expression{
	e := &Expression{
		expression: expression,
	}
	return e
}

type Print struct{
	expression Expr
}

func NewPrint(expression Expr)*Print{
	p := &Print{
		expression: expression,
	}
	return p
}

type VarStmt struct{
	name *Token
	initializer Expr
}

func NewVarStmt(name *Token, initializer Expr)*VarStmt{
	v := &VarStmt{
		name: name,
		initializer: initializer,
	}
	return v
}

type StmtVisitor interface{
	VisitExpressionStmt(expression *Expression)
	VisitPrintStmt(print *Print)
	VisitVarStmtStmt(varstmt *VarStmt)
}

func VisitorStmt(v StmtVisitor,s Stmt){
	switch s.(type){
	case *Expression:
		v.VisitExpressionStmt(s.(*Expression))
	case *Print:
		v.VisitPrintStmt(s.(*Print))
	case *VarStmt:
		v.VisitVarStmtStmt(s.(*VarStmt))
	}
}

type StmtVisitorWithVal[T any] interface{
	VisitExpressionStmt(expression *Expression) T
	VisitPrintStmt(print *Print) T
	VisitVarStmtStmt(varstmt *VarStmt) T
}

func VisitorStmtWithVal[T any](v StmtVisitorWithVal[T],s Stmt) T{
	switch s.(type){
	case *Expression:
		return v.VisitExpressionStmt(s.(*Expression))
	case *Print:
		return v.VisitPrintStmt(s.(*Print))
	case *VarStmt:
		return v.VisitVarStmtStmt(s.(*VarStmt))
	default:
		panic("can't find Stmt")
	}
}

