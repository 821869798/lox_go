package lox

import "fmt"

type RuntimeError struct {
	Token   *Token
	Message string
}

func NewRuntimeError(token *Token, message string) *RuntimeError {
	e := &RuntimeError{
		Token:   token,
		Message: message,
	}
	return e
}

func (e *RuntimeError) Error() string {
	return fmt.Sprintf("%s\n[line%d]", e.Message, e.Token.line)
}
