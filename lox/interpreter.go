package lox

import (
	"fmt"
	"lox_go/util"
	"strconv"
)

type Interpreter struct {
	env     *Environment
	globals *Environment
	locals  map[Expr]int
}

func NewInterpreter() *Interpreter {
	i := &Interpreter{}
	i.globals = NewEnvironment(nil)
	i.env = i.globals
	i.locals = make(map[Expr]int)

	i.globals.define("clock", NewCallableClock())
	return i
}

func (i *Interpreter) interpret(statements []Stmt) {
	defer func() {
		if err := recover(); err != nil {
			v, ok := err.(*RuntimeError)
			if ok {
				reportRuntimeError(v)
			} else {
				panic(err)
			}
		}
	}()

	for _, statement := range statements {
		i.execute(statement)
	}
}

func (i *Interpreter) execute(stmt Stmt) {
	VisitorStmt(i, stmt)
}

func (i *Interpreter) resolve(expr Expr, depth int) {
	i.locals[expr] = depth
}

func (i *Interpreter) executeBlock(statements []Stmt, environment *Environment) {
	previous := i.env
	defer func() {
		i.env = previous
	}()

	i.env = environment
	for _, statement := range statements {
		i.execute(statement)
	}
}

func (i *Interpreter) VisitBlockStmt(stmt *BlockStmt) {
	i.executeBlock(stmt.statements, NewEnvironment(i.env))
}

func (i *Interpreter) VisitExpressionStmt(stmt *ExpressionStmt) {
	i.evaluate(stmt.expression)
}

func (i *Interpreter) VisitFunctionStmt(stmt *FunctionStmt) {
	function := NewLoxFunction(stmt, i.env)
	i.env.define(stmt.name.lexeme, function)
}

func (i *Interpreter) VisitIfStmt(stmt *IfStmt) {
	if i.isTruthy(i.evaluate(stmt.condition)) {
		i.execute(stmt.thenBranch)
	} else if stmt.elseBranch != nil {
		i.execute(stmt.elseBranch)
	}
}

func (i *Interpreter) VisitPrintStmt(stmt *PrintStmt) {
	value := i.evaluate(stmt.expression)
	fmt.Print(util.GetInterfaceToString(value))
}

func (i *Interpreter) VisitReturnStmt(stmt *ReturnStmt) {
	var value interface{} = nil
	if stmt.value != nil {
		value = i.evaluate(stmt.value)
	}
	panic(NewReturn(value))
}

func (i *Interpreter) VisitVarStmt(stmt *VarStmt) {
	var value interface{} = nil
	if stmt.initializer != nil {
		value = i.evaluate(stmt.initializer)
	}
	i.env.define(stmt.name.lexeme, value)
}

func (i *Interpreter) VisitWhileStmt(stmt *WhileStmt) {
	for i.isTruthy(i.evaluate(stmt.condition)) {
		i.execute(stmt.body)
	}
}

func (i *Interpreter) VisitAssignExpr(expr *AssignExpr) interface{} {
	value := i.evaluate(expr.value)
	//i.env.assign(expr.name, value)
	distance, ok := i.locals[expr]
	if ok {
		i.env.assignAt(distance, expr.name, value)
	} else {
		i.globals.assign(expr.name, value)
	}
	return value
}

func (i *Interpreter) evaluate(expr Expr) interface{} {
	return VisitorExprWithVal[interface{}](i, expr)
}

func (i *Interpreter) VisitLiteralExpr(expr *LiteralExpr) interface{} {
	return expr.value
}

func (i *Interpreter) VisitLogicalExpr(expr *LogicalExpr) interface{} {
	left := i.evaluate(expr.left)

	if expr.operator.tokenType == TokenType_OR {
		if i.isTruthy(left) {
			return left
		}
	} else {
		if !i.isTruthy(left) {
			return left
		}
	}

	return i.evaluate(expr.right)
}

func (i *Interpreter) VisitGroupingExpr(expr *GroupingExpr) interface{} {
	return i.evaluate(expr.expression)
}

func (i *Interpreter) VisitUnaryExpr(expr *UnaryExpr) interface{} {
	right := i.evaluate(expr.right)

	switch expr.operator.tokenType {
	case TokenType_MINUS:
		i.checkNumberOperand(expr.operator, right)
		return -(right.(float64))
	case TokenType_BANG:
		return !i.isTruthy(right)
	}
	return nil
}

func (i *Interpreter) VisitVariableExpr(expr *VariableExpr) interface{} {
	//return i.env.get(expr.name)
	return i.lookUpVariable(expr.name, expr)
}

func (i *Interpreter) lookUpVariable(name *Token, expr Expr) interface{} {
	distance, ok := i.locals[expr]
	if ok {
		return i.env.getAt(distance, name.lexeme)
	} else {
		return i.globals.get(name)
	}
}

func (i *Interpreter) VisitBinaryExpr(expr *BinaryExpr) interface{} {
	left := i.evaluate(expr.left)
	right := i.evaluate(expr.right)

	switch expr.operator.tokenType {
	case TokenType_MINUS:
		i.checkNumberOperands(expr.operator, left, right)
		return left.(float64) - right.(float64)
	case TokenType_PLUS:
		// 加分特殊，不只是数值加法，还要考虑字符串连接
		f1, fok1 := left.(float64)
		f2, fok2 := right.(float64)
		if fok1 && fok2 {
			return f1 + f2
		}
		s1, sok1 := left.(string)
		s2, sok2 := right.(string)
		if sok1 && sok2 {
			return s1 + s2
		}
		if sok1 && fok2 {
			key2 := strconv.FormatFloat(f2, 'f', -1, 64)
			return s1 + key2
		}
		if sok2 && fok1 {
			key1 := strconv.FormatFloat(f1, 'f', -1, 64)
			return key1 + s2
		}
		panic(NewRuntimeError(expr.operator, "Operands must be two numbers or strings."))
	case TokenType_SLASH:
		i.checkNumberOperands(expr.operator, left, right)
		return left.(float64) / right.(float64)
	case TokenType_STAR:
		i.checkNumberOperands(expr.operator, left, right)
		return left.(float64) * right.(float64)
	case TokenType_GREATER:
		i.checkNumberOperands(expr.operator, left, right)
		return left.(float64) > right.(float64)
	case TokenType_GREATER_EQUAL:
		i.checkNumberOperands(expr.operator, left, right)
		return left.(float64) >= right.(float64)
	case TokenType_LESS:
		i.checkNumberOperands(expr.operator, left, right)
		return left.(float64) < right.(float64)
	case TokenType_LESS_EQUAL:
		i.checkNumberOperands(expr.operator, left, right)
		return left.(float64) <= right.(float64)
	case TokenType_BANG_EQUAL:
		return !i.isEqual(left, right)
	case TokenType_EQUAL_EQUAL:
		return i.isEqual(left, right)
	}
	return nil
}

func (i *Interpreter) VisitCallExpr(expr *CallExpr) interface{} {
	callee := i.evaluate(expr.callee)

	var arguments []interface{} = nil
	for _, argument := range expr.arguments {
		arguments = append(arguments, i.evaluate(argument))
	}

	function, ok := callee.(LoxCallable)
	if !ok {
		panic(NewRuntimeError(expr.paren, "Can only call functions and classes."))
	}
	// 新增部分开始
	if len(arguments) != function.Arity() {
		panic(NewRuntimeError(expr.paren, fmt.Sprintf("Expected %d arguments but got %d.", function.Arity(), len(arguments))))
	}

	return function.Call(i, arguments)
}

func (i *Interpreter) isTruthy(obj interface{}) bool {
	if obj == nil {
		return false
	}
	value, ok := obj.(bool)
	if ok {
		return value
	}
	return true
}

func (i *Interpreter) isEqual(a interface{}, b interface{}) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return false
	}
	return a == b
}

func (i *Interpreter) checkNumberOperand(operator *Token, operand interface{}) {
	_, ok := operand.(float64)
	if ok {
		return
	}
	panic(NewRuntimeError(operator, "Operand must be a number."))
}

func (i *Interpreter) checkNumberOperands(operator *Token, left interface{}, right interface{}) {
	_, ok1 := left.(float64)
	_, ok2 := right.(float64)
	if ok1 && ok2 {
		return
	}
	panic(NewRuntimeError(operator, "Operands must be a numbers."))
}
