package lox

type Environment struct {
	values map[string]interface{}
}

func NewEnvironment() *Environment {
	e := &Environment{
		values: make(map[string]interface{}),
	}
	return e
}

func (e *Environment) define(name string, value interface{}) {
	e.values[name] = value
}

func (e *Environment) get(name *Token) interface{} {
	value, ok := e.values[name.lexeme]
	if ok {
		return value
	}
	panic(NewRuntimeError(name, "Undefined variable '"+name.lexeme+"'."))
}

func (e *Environment) assign(name *Token, value interface{}) {
	value, ok := e.values[name.lexeme]
	if ok {
		e.values[name.lexeme] = value
		return
	}
	panic(NewRuntimeError(name, "Undefined variable '"+name.lexeme+"'."))
}
