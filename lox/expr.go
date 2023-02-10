package lox

type Expr interface {
}

type Binary struct{
	left Expr
	operator *Token
	right Expr
}

func NewBinary(left Expr, operator *Token, right Expr)*Binary{
	b := &Binary{
		left: left,
		operator: operator,
		right: right,
	}
	return b
}

type Grouping struct{
	expression Expr
}

func NewGrouping(expression Expr)*Grouping{
	g := &Grouping{
		expression: expression,
	}
	return g
}

type Literal struct{
	value interface{}
}

func NewLiteral(value interface{})*Literal{
	l := &Literal{
		value: value,
	}
	return l
}

type Unary struct{
	operator *Token
	right Expr
}

func NewUnary(operator *Token, right Expr)*Unary{
	u := &Unary{
		operator: operator,
		right: right,
	}
	return u
}

type Visitor interface{
	VisitBinaryExpr(binary *Binary)
	VisitGroupingExpr(grouping *Grouping)
	VisitLiteralExpr(literal *Literal)
	VisitUnaryExpr(unary *Unary)
}

func VisitorExpr(v Visitor,e Expr){
	switch e.(type){
	case *Binary:
		v.VisitBinaryExpr(e.(*Binary))
	case *Grouping:
		v.VisitGroupingExpr(e.(*Grouping))
	case *Literal:
		v.VisitLiteralExpr(e.(*Literal))
	case *Unary:
		v.VisitUnaryExpr(e.(*Unary))
	}
}

type VisitorWithVal[T any] interface{
	VisitBinaryExpr(binary *Binary) T
	VisitGroupingExpr(grouping *Grouping) T
	VisitLiteralExpr(literal *Literal) T
	VisitUnaryExpr(unary *Unary) T
}

func VisitorExprWithVal[T any](v VisitorWithVal[T],e Expr) T{
	switch e.(type){
	case *Binary:
		return v.VisitBinaryExpr(e.(*Binary))
	case *Grouping:
		return v.VisitGroupingExpr(e.(*Grouping))
	case *Literal:
		return v.VisitLiteralExpr(e.(*Literal))
	case *Unary:
		return v.VisitUnaryExpr(e.(*Unary))
	default:
		panic("can't find Expr")
	}
}

