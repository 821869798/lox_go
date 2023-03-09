package lox

type Stmt interface {
}

type BlockStmt struct{
	statements []Stmt
}

func NewBlockStmt(statements []Stmt)*BlockStmt{
	b := &BlockStmt{
		statements: statements,
	}
	return b
}

type ExpressionStmt struct{
	expression Expr
}

func NewExpressionStmt(expression Expr)*ExpressionStmt{
	e := &ExpressionStmt{
		expression: expression,
	}
	return e
}

type IfStmt struct{
	condition Expr
	thenBranch Stmt
	elseBranch Stmt
}

func NewIfStmt(condition Expr, thenBranch Stmt, elseBranch Stmt)*IfStmt{
	i := &IfStmt{
		condition: condition,
		thenBranch: thenBranch,
		elseBranch: elseBranch,
	}
	return i
}

type PrintStmt struct{
	expression Expr
}

func NewPrintStmt(expression Expr)*PrintStmt{
	p := &PrintStmt{
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

type WhileStmt struct{
	condition Expr
	body Stmt
}

func NewWhileStmt(condition Expr, body Stmt)*WhileStmt{
	w := &WhileStmt{
		condition: condition,
		body: body,
	}
	return w
}

type StmtVisitor interface{
	VisitBlockStmt(blockstmt *BlockStmt)
	VisitExpressionStmt(expressionstmt *ExpressionStmt)
	VisitIfStmt(ifstmt *IfStmt)
	VisitPrintStmt(printstmt *PrintStmt)
	VisitVarStmt(varstmt *VarStmt)
	VisitWhileStmt(whilestmt *WhileStmt)
}

func VisitorStmt(v StmtVisitor,s Stmt){
	switch s.(type){
	case *BlockStmt:
		v.VisitBlockStmt(s.(*BlockStmt))
	case *ExpressionStmt:
		v.VisitExpressionStmt(s.(*ExpressionStmt))
	case *IfStmt:
		v.VisitIfStmt(s.(*IfStmt))
	case *PrintStmt:
		v.VisitPrintStmt(s.(*PrintStmt))
	case *VarStmt:
		v.VisitVarStmt(s.(*VarStmt))
	case *WhileStmt:
		v.VisitWhileStmt(s.(*WhileStmt))
	}
}

type StmtVisitorWithVal[T any] interface{
	VisitBlockStmt(blockstmt *BlockStmt) T
	VisitExpressionStmt(expressionstmt *ExpressionStmt) T
	VisitIfStmt(ifstmt *IfStmt) T
	VisitPrintStmt(printstmt *PrintStmt) T
	VisitVarStmt(varstmt *VarStmt) T
	VisitWhileStmt(whilestmt *WhileStmt) T
}

func VisitorStmtWithVal[T any](v StmtVisitorWithVal[T],s Stmt) T{
	switch s.(type){
	case *BlockStmt:
		return v.VisitBlockStmt(s.(*BlockStmt))
	case *ExpressionStmt:
		return v.VisitExpressionStmt(s.(*ExpressionStmt))
	case *IfStmt:
		return v.VisitIfStmt(s.(*IfStmt))
	case *PrintStmt:
		return v.VisitPrintStmt(s.(*PrintStmt))
	case *VarStmt:
		return v.VisitVarStmt(s.(*VarStmt))
	case *WhileStmt:
		return v.VisitWhileStmt(s.(*WhileStmt))
	default:
		panic("can't find Stmt")
	}
}

