package lox

type LoxClass struct {
	name       string
	methods    map[string]*LoxFunction
	superclass *LoxClass
}

func NewLoxClass(name string, superclass *LoxClass, methods map[string]*LoxFunction) *LoxClass {
	l := &LoxClass{
		name:       name,
		methods:    methods,
		superclass: superclass,
	}
	return l
}

func (l *LoxClass) String() string {
	return l.name
}

func (l *LoxClass) Arity() int {
	initializer := l.FindMethod("init")
	if initializer != nil {
		return initializer.Arity()
	}
	return 0
}

func (l *LoxClass) Call(interpreter *Interpreter, arguments []interface{}) interface{} {
	instance := NewLoxInstance(l)
	initializer := l.FindMethod("init")
	if initializer != nil {
		initializer.Bind(instance).Call(interpreter, arguments)
	}
	return instance
}

func (l *LoxClass) FindMethod(name string) *LoxFunction {
	method, ok := l.methods[name]
	if ok {
		return method
	}
	if l.superclass != nil {
		return l.superclass.FindMethod(name)
	}
	return nil
}
