package lox

import (
	"time"
)

type CallableClock struct {
}

func NewCallableClock() LoxCallable {
	return &CallableClock{}
}

func (c *CallableClock) Arity() int {
	return 0
}

func (c *CallableClock) Call(interpreter *Interpreter, arguments []interface{}) interface{} {
	currentTime := time.Now()
	secondsWithFractional := float64(currentTime.UnixNano()) / 1e9
	return secondsWithFractional
}

func (c *CallableClock) String() string {
	return "<native fn clock>"
}
