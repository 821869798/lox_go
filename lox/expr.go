package lox

type Expr interface {
}

type AssignExpr struct{
	name *Token
	value Expr
}

func NewAssignExpr(name *Token, value Expr)*AssignExpr{
	a := &AssignExpr{
		name: name,
		value: value,
	}
	return a
}

type BinaryExpr struct{
	left Expr
	operator *Token
	right Expr
}

func NewBinaryExpr(left Expr, operator *Token, right Expr)*BinaryExpr{
	b := &BinaryExpr{
		left: left,
		operator: operator,
		right: right,
	}
	return b
}

type CallExpr struct{
	callee Expr
	paren *Token
	arguments []Expr
}

func NewCallExpr(callee Expr, paren *Token, arguments []Expr)*CallExpr{
	c := &CallExpr{
		callee: callee,
		paren: paren,
		arguments: arguments,
	}
	return c
}

type GetExpr struct{
	object Expr
	name *Token
}

func NewGetExpr(object Expr, name *Token)*GetExpr{
	g := &GetExpr{
		object: object,
		name: name,
	}
	return g
}

type GroupingExpr struct{
	expression Expr
}

func NewGroupingExpr(expression Expr)*GroupingExpr{
	g := &GroupingExpr{
		expression: expression,
	}
	return g
}

type LiteralExpr struct{
	value interface{}
}

func NewLiteralExpr(value interface{})*LiteralExpr{
	l := &LiteralExpr{
		value: value,
	}
	return l
}

type LogicalExpr struct{
	left Expr
	operator *Token
	right Expr
}

func NewLogicalExpr(left Expr, operator *Token, right Expr)*LogicalExpr{
	l := &LogicalExpr{
		left: left,
		operator: operator,
		right: right,
	}
	return l
}

type SetExpr struct{
	object Expr
	name *Token
	value Expr
}

func NewSetExpr(object Expr, name *Token, value Expr)*SetExpr{
	s := &SetExpr{
		object: object,
		name: name,
		value: value,
	}
	return s
}

type ThisExpr struct{
	keyword *Token
}

func NewThisExpr(keyword *Token)*ThisExpr{
	t := &ThisExpr{
		keyword: keyword,
	}
	return t
}

type UnaryExpr struct{
	operator *Token
	right Expr
}

func NewUnaryExpr(operator *Token, right Expr)*UnaryExpr{
	u := &UnaryExpr{
		operator: operator,
		right: right,
	}
	return u
}

type VariableExpr struct{
	name *Token
}

func NewVariableExpr(name *Token)*VariableExpr{
	v := &VariableExpr{
		name: name,
	}
	return v
}

type ExprVisitor interface{
	VisitAssignExpr(assignexpr *AssignExpr)
	VisitBinaryExpr(binaryexpr *BinaryExpr)
	VisitCallExpr(callexpr *CallExpr)
	VisitGetExpr(getexpr *GetExpr)
	VisitGroupingExpr(groupingexpr *GroupingExpr)
	VisitLiteralExpr(literalexpr *LiteralExpr)
	VisitLogicalExpr(logicalexpr *LogicalExpr)
	VisitSetExpr(setexpr *SetExpr)
	VisitThisExpr(thisexpr *ThisExpr)
	VisitUnaryExpr(unaryexpr *UnaryExpr)
	VisitVariableExpr(variableexpr *VariableExpr)
}

func VisitorExpr(v ExprVisitor,e Expr){
	switch e.(type){
	case *AssignExpr:
		v.VisitAssignExpr(e.(*AssignExpr))
	case *BinaryExpr:
		v.VisitBinaryExpr(e.(*BinaryExpr))
	case *CallExpr:
		v.VisitCallExpr(e.(*CallExpr))
	case *GetExpr:
		v.VisitGetExpr(e.(*GetExpr))
	case *GroupingExpr:
		v.VisitGroupingExpr(e.(*GroupingExpr))
	case *LiteralExpr:
		v.VisitLiteralExpr(e.(*LiteralExpr))
	case *LogicalExpr:
		v.VisitLogicalExpr(e.(*LogicalExpr))
	case *SetExpr:
		v.VisitSetExpr(e.(*SetExpr))
	case *ThisExpr:
		v.VisitThisExpr(e.(*ThisExpr))
	case *UnaryExpr:
		v.VisitUnaryExpr(e.(*UnaryExpr))
	case *VariableExpr:
		v.VisitVariableExpr(e.(*VariableExpr))
	}
}

type ExprVisitorWithVal[T any] interface{
	VisitAssignExpr(assignexpr *AssignExpr) T
	VisitBinaryExpr(binaryexpr *BinaryExpr) T
	VisitCallExpr(callexpr *CallExpr) T
	VisitGetExpr(getexpr *GetExpr) T
	VisitGroupingExpr(groupingexpr *GroupingExpr) T
	VisitLiteralExpr(literalexpr *LiteralExpr) T
	VisitLogicalExpr(logicalexpr *LogicalExpr) T
	VisitSetExpr(setexpr *SetExpr) T
	VisitThisExpr(thisexpr *ThisExpr) T
	VisitUnaryExpr(unaryexpr *UnaryExpr) T
	VisitVariableExpr(variableexpr *VariableExpr) T
}

func VisitorExprWithVal[T any](v ExprVisitorWithVal[T],e Expr) T{
	switch e.(type){
	case *AssignExpr:
		return v.VisitAssignExpr(e.(*AssignExpr))
	case *BinaryExpr:
		return v.VisitBinaryExpr(e.(*BinaryExpr))
	case *CallExpr:
		return v.VisitCallExpr(e.(*CallExpr))
	case *GetExpr:
		return v.VisitGetExpr(e.(*GetExpr))
	case *GroupingExpr:
		return v.VisitGroupingExpr(e.(*GroupingExpr))
	case *LiteralExpr:
		return v.VisitLiteralExpr(e.(*LiteralExpr))
	case *LogicalExpr:
		return v.VisitLogicalExpr(e.(*LogicalExpr))
	case *SetExpr:
		return v.VisitSetExpr(e.(*SetExpr))
	case *ThisExpr:
		return v.VisitThisExpr(e.(*ThisExpr))
	case *UnaryExpr:
		return v.VisitUnaryExpr(e.(*UnaryExpr))
	case *VariableExpr:
		return v.VisitVariableExpr(e.(*VariableExpr))
	default:
		panic("can't find Expr")
	}
}

