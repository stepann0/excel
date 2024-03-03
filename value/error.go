package value

import "fmt"

type Error struct {
	Msg error
}

func (e Error) Type() ValueType { return ErrorType }

func (e Error) ToType(fn string, which ValueType) Value {
	if which == ErrorType {
		return e
	}
	TypeError()
	return nil
}

func (e Error) String() string {
	if e.Msg == nil {
		return "#ERR:nil"
	}
	return fmt.Sprintf("#ERR: %s", e.Msg.Error())
}

func errorf(format string, args ...interface{}) {
	panic(fmt.Errorf(format, args...))
}

func TypeError() {
	errorf("Type error")
}

func ArgCountError() {
	errorf("Wrong number of arguments")
}

func NotImplementedError() {
	errorf("Not implemented")
}
