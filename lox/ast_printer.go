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

func (a *AstPrinter) print(expr Expr) string {
	return VisitorExprWithVal[string](a, expr)
}

func (a *AstPrinter) VisitBinaryExpr(expr *Binary) string {
	return a.parenthesize(expr.operator.lexeme, expr.left, expr.right)
}

func (a *AstPrinter) VisitGroupingExpr(grouping *Grouping) string {
	return a.parenthesize("group", grouping.expression)
}

func (a *AstPrinter) VisitLiteralExpr(literal *Literal) string {
	if literal.value == nil {
		return "nil"
	}
	return util.GetInterfaceToString(literal.value)
}

func (a *AstPrinter) VisitUnaryExpr(unary *Unary) string {
	return a.parenthesize(unary.operator.lexeme, unary.right)
}

func (a *AstPrinter) VisitVariableExpr(variable *Variable) string {
	return variable.name.lexeme
}

func (a *AstPrinter) VisitAssignExpr(assign *Assign) string {
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
