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

type ClassStmt struct{
	name *Token
	methods []*FunctionStmt
}

func NewClassStmt(name *Token, methods []*FunctionStmt)*ClassStmt{
	c := &ClassStmt{
		name: name,
		methods: methods,
	}
	return c
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

type FunctionStmt struct{
	name *Token
	params []*Token
	body []Stmt
}

func NewFunctionStmt(name *Token, params []*Token, body []Stmt)*FunctionStmt{
	f := &FunctionStmt{
		name: name,
		params: params,
		body: body,
	}
	return f
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

type ReturnStmt struct{
	keyword *Token
	value Expr
}

func NewReturnStmt(keyword *Token, value Expr)*ReturnStmt{
	r := &ReturnStmt{
		keyword: keyword,
		value: value,
	}
	return r
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
	VisitClassStmt(classstmt *ClassStmt)
	VisitExpressionStmt(expressionstmt *ExpressionStmt)
	VisitFunctionStmt(functionstmt *FunctionStmt)
	VisitIfStmt(ifstmt *IfStmt)
	VisitPrintStmt(printstmt *PrintStmt)
	VisitReturnStmt(returnstmt *ReturnStmt)
	VisitVarStmt(varstmt *VarStmt)
	VisitWhileStmt(whilestmt *WhileStmt)
}

func VisitorStmt(v StmtVisitor,s Stmt){
	switch s.(type){
	case *BlockStmt:
		v.VisitBlockStmt(s.(*BlockStmt))
	case *ClassStmt:
		v.VisitClassStmt(s.(*ClassStmt))
	case *ExpressionStmt:
		v.VisitExpressionStmt(s.(*ExpressionStmt))
	case *FunctionStmt:
		v.VisitFunctionStmt(s.(*FunctionStmt))
	case *IfStmt:
		v.VisitIfStmt(s.(*IfStmt))
	case *PrintStmt:
		v.VisitPrintStmt(s.(*PrintStmt))
	case *ReturnStmt:
		v.VisitReturnStmt(s.(*ReturnStmt))
	case *VarStmt:
		v.VisitVarStmt(s.(*VarStmt))
	case *WhileStmt:
		v.VisitWhileStmt(s.(*WhileStmt))
	}
}

type StmtVisitorWithVal[T any] interface{
	VisitBlockStmt(blockstmt *BlockStmt) T
	VisitClassStmt(classstmt *ClassStmt) T
	VisitExpressionStmt(expressionstmt *ExpressionStmt) T
	VisitFunctionStmt(functionstmt *FunctionStmt) T
	VisitIfStmt(ifstmt *IfStmt) T
	VisitPrintStmt(printstmt *PrintStmt) T
	VisitReturnStmt(returnstmt *ReturnStmt) T
	VisitVarStmt(varstmt *VarStmt) T
	VisitWhileStmt(whilestmt *WhileStmt) T
}

func VisitorStmtWithVal[T any](v StmtVisitorWithVal[T],s Stmt) T{
	switch s.(type){
	case *BlockStmt:
		return v.VisitBlockStmt(s.(*BlockStmt))
	case *ClassStmt:
		return v.VisitClassStmt(s.(*ClassStmt))
	case *ExpressionStmt:
		return v.VisitExpressionStmt(s.(*ExpressionStmt))
	case *FunctionStmt:
		return v.VisitFunctionStmt(s.(*FunctionStmt))
	case *IfStmt:
		return v.VisitIfStmt(s.(*IfStmt))
	case *PrintStmt:
		return v.VisitPrintStmt(s.(*PrintStmt))
	case *ReturnStmt:
		return v.VisitReturnStmt(s.(*ReturnStmt))
	case *VarStmt:
		return v.VisitVarStmt(s.(*VarStmt))
	case *WhileStmt:
		return v.VisitWhileStmt(s.(*WhileStmt))
	default:
		panic("can't find Stmt")
	}
}

