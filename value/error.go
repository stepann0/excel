package value

import "fmt"

type Error struct {
	Msg error
}

func (e Error) Type() ValueType { return ErrorType }

func (e Error) String() string {
	if e.Msg == nil {
		return "#ERR:nil"
	}
	return fmt.Sprintf("#ERR: %s", e.Msg.Error())
}
