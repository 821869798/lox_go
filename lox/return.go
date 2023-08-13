package lox

type Return struct {
	value interface{}
}

func NewReturn(value interface{}) *Return {
	r := &Return{
		value: value,
	}
	return r
}
