package lox

import (
	"testing"
)

func TestAstPrint(t *testing.T) {
	b := NewBinary(
		NewUnary(NewToken(TokenType_MINUS, "-", nil, 1), NewLiteral(123)),
		NewToken(TokenType_STAR, "*", nil, 1),
		NewGrouping(NewLiteral(45.67)),
	)
	t.Log(NewAstPrinter().print(b))
}
