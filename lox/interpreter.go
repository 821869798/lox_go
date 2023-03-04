package lox

import (
	"fmt"
	"lox_go/util"
)

type Interpreter struct {
	env *Environment
}

func NewInterpreter() *Interpreter {
	i := &Interpreter{
		env: NewEnvironment(),
	}
	return i
}

func (i *Interpreter) interpret(statements []Stmt) {

	for _, statement := range statements {
		i.execute(statement)
	}
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
}

func (i *Interpreter) execute(stmt Stmt) {
	VisitorStmt(i, stmt)
}

func (i *Interpreter) VisitExpressionStmt(stmt *Expression) {
	i.evaluate(stmt.expression)
}

func (i *Interpreter) VisitPrintStmt(stmt *Print) {
	value := i.evaluate(stmt.expression)
	fmt.Print(util.GetInterfaceToString(value))
}

func (i *Interpreter) VisitVarStmtStmt(stmt *VarStmt) {
	var value interface{} = nil
	if stmt.initializer != nil {
		value = i.evaluate(stmt)
	}
	i.env.define(stmt.name.lexeme, value)
}

func (i *Interpreter) VisitAssignExpr(expr *Assign) interface{} {
	value := i.evaluate(expr)
	i.env.assign(expr.name, value)
	return value
}

func (i *Interpreter) evaluate(expr Expr) interface{} {
	return VisitorExprWithVal[interface{}](i, expr)
}

func (i *Interpreter) VisitLiteralExpr(expr *Literal) interface{} {
	return expr.value
}

func (i *Interpreter) VisitGroupingExpr(expr *Grouping) interface{} {
	return i.evaluate(expr.expression)
}

func (i *Interpreter) VisitUnaryExpr(expr *Unary) interface{} {
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

func (i *Interpreter) VisitVariableExpr(expr *Variable) interface{} {
	return i.env.get(expr.name)
}

func (i *Interpreter) VisitBinaryExpr(expr *Binary) interface{} {
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
		panic(NewRuntimeError(expr.operator, "Operands must be two numbers or two strings."))
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
