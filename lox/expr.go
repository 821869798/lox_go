package lox

type Expr interface {
}

type Assign struct{
	name *Token
	value Expr
}

func NewAssign(name *Token, value Expr)*Assign{
	a := &Assign{
		name: name,
		value: value,
	}
	return a
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

type Variable struct{
	name *Token
}

func NewVariable(name *Token)*Variable{
	v := &Variable{
		name: name,
	}
	return v
}

type ExprVisitor interface{
	VisitAssignExpr(assign *Assign)
	VisitBinaryExpr(binary *Binary)
	VisitGroupingExpr(grouping *Grouping)
	VisitLiteralExpr(literal *Literal)
	VisitUnaryExpr(unary *Unary)
	VisitVariableExpr(variable *Variable)
}

func VisitorExpr(v ExprVisitor,e Expr){
	switch e.(type){
	case *Assign:
		v.VisitAssignExpr(e.(*Assign))
	case *Binary:
		v.VisitBinaryExpr(e.(*Binary))
	case *Grouping:
		v.VisitGroupingExpr(e.(*Grouping))
	case *Literal:
		v.VisitLiteralExpr(e.(*Literal))
	case *Unary:
		v.VisitUnaryExpr(e.(*Unary))
	case *Variable:
		v.VisitVariableExpr(e.(*Variable))
	}
}

type ExprVisitorWithVal[T any] interface{
	VisitAssignExpr(assign *Assign) T
	VisitBinaryExpr(binary *Binary) T
	VisitGroupingExpr(grouping *Grouping) T
	VisitLiteralExpr(literal *Literal) T
	VisitUnaryExpr(unary *Unary) T
	VisitVariableExpr(variable *Variable) T
}

func VisitorExprWithVal[T any](v ExprVisitorWithVal[T],e Expr) T{
	switch e.(type){
	case *Assign:
		return v.VisitAssignExpr(e.(*Assign))
	case *Binary:
		return v.VisitBinaryExpr(e.(*Binary))
	case *Grouping:
		return v.VisitGroupingExpr(e.(*Grouping))
	case *Literal:
		return v.VisitLiteralExpr(e.(*Literal))
	case *Unary:
		return v.VisitUnaryExpr(e.(*Unary))
	case *Variable:
		return v.VisitVariableExpr(e.(*Variable))
	default:
		panic("can't find Expr")
	}
}

