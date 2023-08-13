package lox

import (
	"lox_go/util"
	"strings"
)

type AstPrinter struct {
}

func NewAstPrinter() *AstPrinter {
	a := &AstPrinter{}
	return a
}

func (a *AstPrinter) Print(expr Expr) string {
	return VisitorExprWithVal[string](a, expr)
}

func (a *AstPrinter) VisitBinaryExpr(expr *BinaryExpr) string {
	return a.parenthesize(expr.operator.lexeme, expr.left, expr.right)
}
func (a *AstPrinter) VisitCallExpr(expr *CallExpr) string {
	return a.parenthesize("function", expr.callee)
}

func (a *AstPrinter) VisitGroupingExpr(grouping *GroupingExpr) string {
	return a.parenthesize("group", grouping.expression)
}

func (a *AstPrinter) VisitLiteralExpr(literal *LiteralExpr) string {
	if literal.value == nil {
		return "nil"
	}
	return util.GetInterfaceToString(literal.value)
}

func (a *AstPrinter) VisitLogicalExpr(expr *LogicalExpr) string {
	return a.parenthesize(expr.operator.lexeme, expr.left, expr.right)
}

func (a *AstPrinter) VisitUnaryExpr(unary *UnaryExpr) string {
	return a.parenthesize(unary.operator.lexeme, unary.right)
}

func (a *AstPrinter) VisitVariableExpr(variable *VariableExpr) string {
	return variable.name.lexeme
}

func (a *AstPrinter) VisitAssignExpr(assign *AssignExpr) string {
	return ""
}

func (a *AstPrinter) parenthesize(name string, exprs ...Expr) string {
	var b strings.Builder
	b.WriteString("(")
	b.WriteString(name)
	for _, expr := range exprs {
		b.WriteString(" ")
		b.WriteString(VisitorExprWithVal[string](a, expr))
	}
	b.WriteString(")")
	return b.String()
}
