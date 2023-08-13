package lox

import "fmt"

type LoxCallable interface {
	fmt.Stringer
	Arity() int
	Call(interpreter *Interpreter, arguments []interface{}) interface{}
}
