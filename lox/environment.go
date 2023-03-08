package lox

type Environment struct {
	values    map[string]interface{}
	enclosing *Environment
}

func NewEnvironment(enclosing *Environment) *Environment {
	e := &Environment{
		values:    make(map[string]interface{}),
		enclosing: enclosing,
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
	if e.enclosing != nil {
		//如果当前环境中没有找到变量，就在外围环境中尝试
		return e.enclosing.get(name)
	}

	panic(NewRuntimeError(name, "Undefined variable '"+name.lexeme+"'."))
}

func (e *Environment) assign(name *Token, value interface{}) {
	value, ok := e.values[name.lexeme]
	if ok {
		e.values[name.lexeme] = value
		return
	}

	if e.enclosing != nil {
		e.enclosing.assign(name, value)
		return
	}
	panic(NewRuntimeError(name, "Undefined variable '"+name.lexeme+"'."))
}
