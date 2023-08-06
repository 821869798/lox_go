package test

import (
	"lox_go/lox"
	"testing"
)

func TestAstPrint(t *testing.T) {
	b := lox.NewBinaryExpr(
		lox.NewUnaryExpr(lox.NewToken(lox.TokenType_MINUS, "-", nil, 1), lox.NewLiteralExpr(123)),
		lox.NewToken(lox.TokenType_STAR, "*", nil, 1),
		lox.NewGroupingExpr(lox.NewLiteralExpr(45.67)),
	)
	t.Log(lox.NewAstPrinter().Print(b))
}
