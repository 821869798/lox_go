package lox

type LoxFunction struct {
	declaration *FunctionStmt
	closure     *Environment
}

func NewLoxFunction(declaration *FunctionStmt, closure *Environment) LoxCallable {
	f := &LoxFunction{
		declaration: declaration,
		closure:     closure,
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
				returnValue = ret.value
			} else {
				panic(r)
			}
		}
	}()

	interpreter.executeBlock(l.declaration.body, environment)

	returnValue = nil
	return
}
func (l *LoxFunction) String() string {
	return "<fn " + l.declaration.name.lexeme + ">"
}
