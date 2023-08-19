package lox

type LoxInstance struct {
	class  *LoxClass
	fields map[string]interface{}
}

func NewLoxInstance(class *LoxClass) *LoxInstance {
	l := &LoxInstance{
		class:  class,
		fields: make(map[string]interface{}),
	}
	return l
}

func (l *LoxInstance) String() string {
	return l.class.name + " instance"
}

func (l *LoxInstance) Get(name *Token) interface{} {
	value, ok := l.fields[name.lexeme]
	if ok {
		return value
	}

	method := l.class.FindMethod(name.lexeme)
	if method != nil {
		return method.Bind(l)
	}

	panic(NewRuntimeError(name, "Undefined property '"+name.lexeme+"'."))
}

func (l *LoxInstance) Set(name *Token, value interface{}) {
	l.fields[name.lexeme] = value
}
