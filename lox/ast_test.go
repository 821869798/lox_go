package lox

import (
	"testing"
)

func TestAstPrint(t *testing.T) {
	b := NewBinaryExpr(
		NewUnaryExpr(NewToken(TokenType_MINUS, "-", nil, 1), NewLiteral(123)),
		NewToken(TokenType_STAR, "*", nil, 1),
		NewGroupingExpr(NewLiteralExpr(45.67)),
	)
	t.Log(NewAstPrinter().print(b))
}
