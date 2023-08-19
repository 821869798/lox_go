package lox

import (
	"lox_go/generic/stack"
)

type FunctionType int

const (
	FunctionType_None FunctionType = iota
	FunctionType_Function
	FunctionType_Initializer
	FunctionType_Method
)

type ClassType int

const (
	ClassType_None ClassType = iota
	ClassType_Class
)

type Resolver struct {
	interpreter     *Interpreter
	scopes          *stack.Stack[map[string]bool]
	currentFunction FunctionType
	currentClass    ClassType
}

func NewResolver(interpreter *Interpreter) *Resolver {
	r := &Resolver{
		interpreter:     interpreter,
		scopes:          stack.New[map[string]bool](),
		currentFunction: FunctionType_None,
		currentClass:    ClassType_None,
	}
	return r
}

func (r *Resolver) resolveStmt(statements []Stmt) {
	for _, statement := range statements {
		r.resolveStmtOne(statement)
	}
}

func (r *Resolver) resolveStmtOne(stmt Stmt) {
	VisitorStmt(r, stmt)
}

func (r *Resolver) resolveExpr(expr Expr) {
	VisitorExpr(r, expr)
}

func (r *Resolver) beginScope() {
	r.scopes.Push(make(map[string]bool))
}

func (r *Resolver) endScope() {
	r.scopes.Pop()
}

func (r *Resolver) declare(name *Token) {
	if r.scopes.Size() <= 0 {
		return
	}

	scope := r.scopes.Peek()
	if _, ok := scope[name.lexeme]; ok {
		reportErrorToken(name, "Already a variable with this name in this scope.")
	}
	scope[name.lexeme] = false
}

func (r *Resolver) define(name *Token) {
	if r.scopes.Size() <= 0 {
		return
	}
	r.scopes.Peek()[name.lexeme] = true
}

func (r *Resolver) VisitBlockStmt(blockstmt *BlockStmt) {
	r.beginScope()
	r.resolveStmt(blockstmt.statements)
	r.endScope()
}

func (r *Resolver) VisitClassStmt(stmt *ClassStmt) {
	enclosinngClass := r.currentClass
	r.currentClass = ClassType_Class
	r.declare(stmt.name)
	r.define(stmt.name)
	r.beginScope()
	r.scopes.Peek()["this"] = true
	for _, method := range stmt.methods {
		declaration := FunctionType_Method
		if method.name.lexeme == "init" {
			declaration = FunctionType_Initializer
		}
		r.resolveFunction(method, declaration)
	}
	r.endScope()
	r.currentClass = enclosinngClass
}

func (r *Resolver) VisitVarStmt(varstmt *VarStmt) {
	r.declare(varstmt.name)
	if varstmt.initializer != nil {
		r.resolveExpr(varstmt.initializer)
	}
	r.define(varstmt.name)
}

func (r *Resolver) VisitVariableExpr(variableexpr *VariableExpr) {
	if r.scopes.Size() > 0 {
		scope := r.scopes.Peek()
		if value, ok := scope[variableexpr.name.lexeme]; value == false && ok {
			reportErrorToken(variableexpr.name, "Can't read local variable in its own initializer.")
		}
	}
	r.resolveLocal(variableexpr, variableexpr.name)

}

func (r *Resolver) resolveLocal(expr Expr, name *Token) {
	for i := r.scopes.Size() - 1; i >= 0; i-- {
		scope := r.scopes.Get(i)
		if _, ok := scope[name.lexeme]; ok {
			r.interpreter.resolve(expr, r.scopes.Size()-1-i)
			return
		}
	}
}

func (r *Resolver) VisitAssignExpr(assignexpr *AssignExpr) {
	r.resolveExpr(assignexpr.value)
	r.resolveLocal(assignexpr, assignexpr.name)
}

func (r *Resolver) VisitFunctionStmt(functionstmt *FunctionStmt) {
	r.declare(functionstmt.name)
	r.define(functionstmt.name)
	r.resolveFunction(functionstmt, FunctionType_Function)
}

func (r *Resolver) resolveFunction(functionstmt *FunctionStmt, functionType FunctionType) {
	enclosingFunction := r.currentFunction
	r.currentFunction = functionType
	r.beginScope()
	for _, param := range functionstmt.params {
		r.declare(param)
		r.define(param)
	}
	r.resolveStmt(functionstmt.body)
	r.endScope()
	r.currentFunction = enclosingFunction
}

func (r *Resolver) VisitExpressionStmt(expressionstmt *ExpressionStmt) {
	r.resolveExpr(expressionstmt.expression)
}

func (r *Resolver) VisitIfStmt(ifstmt *IfStmt) {
	r.resolveExpr(ifstmt.condition)
	r.resolveStmtOne(ifstmt.thenBranch)
	if ifstmt.elseBranch != nil {
		r.resolveStmtOne(ifstmt.elseBranch)
	}
}

func (r *Resolver) VisitPrintStmt(printstmt *PrintStmt) {
	r.resolveExpr(printstmt.expression)
}

func (r *Resolver) VisitReturnStmt(returnstmt *ReturnStmt) {
	if r.currentFunction == FunctionType_None {
		reportErrorToken(returnstmt.keyword, "Can't return from top-level code.")
	}
	if returnstmt.value != nil {
		if r.currentFunction == FunctionType_Initializer {
			reportErrorToken(returnstmt.keyword, "Can't return a value from an initializer.")
		}
		r.resolveExpr(returnstmt.value)
	}
}

func (r *Resolver) VisitWhileStmt(whilestmt *WhileStmt) {
	r.resolveExpr(whilestmt.condition)
	r.resolveStmtOne(whilestmt.body)
}

func (r *Resolver) VisitBinaryExpr(binaryexpr *BinaryExpr) {
	r.resolveExpr(binaryexpr.left)
	r.resolveExpr(binaryexpr.right)
}

func (r *Resolver) VisitCallExpr(callexpr *CallExpr) {
	r.resolveExpr(callexpr.callee)
	for _, argument := range callexpr.arguments {
		r.resolveExpr(argument)
	}
}
func (r *Resolver) VisitGetExpr(getexpr *GetExpr) {
	r.resolveExpr(getexpr.object)
}

func (r *Resolver) VisitGroupingExpr(groupingexpr *GroupingExpr) {
	r.resolveExpr(groupingexpr.expression)
}

func (r *Resolver) VisitLiteralExpr(literalexpr *LiteralExpr) {

}

func (r *Resolver) VisitLogicalExpr(logicalexpr *LogicalExpr) {
	r.resolveExpr(logicalexpr.left)
	r.resolveExpr(logicalexpr.right)
}

func (r *Resolver) VisitSetExpr(setexpr *SetExpr) {
	r.resolveExpr(setexpr.value)
	r.resolveExpr(setexpr.object)
}

func (r *Resolver) VisitThisExpr(thisexpr *ThisExpr) {
	if r.currentClass == ClassType_None {
		reportErrorToken(thisexpr.keyword, "Can't use 'this' outside of a class.")
	}
	r.resolveLocal(thisexpr, thisexpr.keyword)
}

func (r *Resolver) VisitUnaryExpr(unaryexpr *UnaryExpr) {
	r.resolveExpr(unaryexpr.right)
}
