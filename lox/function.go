package lox

type LoxFunction struct {
	declaration   *FunctionStmt
	closure       *Environment
	isInitializer bool
}

func NewLoxFunction(declaration *FunctionStmt, closure *Environment, isInitializer bool) *LoxFunction {
	f := &LoxFunction{
		declaration:   declaration,
		closure:       closure,
		isInitializer: isInitializer,
	}
	return f
}

func (l *LoxFunction) Arity() int {
	return len(l.declaration.params)
}

func (l *LoxFunction) Call(interpreter *Interpreter, arguments []interface{}) (returnValue interface{}) {
	environment := NewEnvironment(l.closure)
	for i, p := range l.declaration.params {
		environment.define(p.lexeme, arguments[i])
	}

	defer func() {
		if r := recover(); r != nil {
			if ret, ok := r.(*Return); ok {
				if l.isInitializer {
					returnValue = l.closure.getAt(0, "this")
				} else {
					returnValue = ret.value
				}
			} else {
				panic(r)
			}
		}
	}()

	interpreter.executeBlock(l.declaration.body, environment)
	returnValue = nil
	if l.isInitializer {
		returnValue = l.closure.getAt(0, "this")
	}

	return
}
func (l *LoxFunction) String() string {
	return "<fn " + l.declaration.name.lexeme + ">"
}

func (l *LoxFunction) Bind(instance *LoxInstance) *LoxFunction {
	environment := NewEnvironment(l.closure)
	environment.define("this", instance)
	return NewLoxFunction(l.declaration, environment, l.isInitializer)
}
