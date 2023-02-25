package lox

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
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
	return GetInterfaceToString(literal.value)
}

func (a *AstPrinter) VisitUnaryExpr(unary *Unary) string {
	return a.parenthesize(unary.operator.lexeme, unary.right)
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

func GetInterfaceToString(value interface{}) string {
	// interface 转 string
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case time.Time:
		t, _ := value.(time.Time)
		key = t.String()
		// 2022-11-23 11:29:07 +0800 CST  这类格式把尾巴去掉
		key = strings.Replace(key, " +0800 CST", "", 1)
		key = strings.Replace(key, " +0000 UTC", "", 1)
	case []byte:
		key = string(value.([]byte))
	case nil:
		key = "nil"
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}
