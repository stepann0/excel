package value

import "fmt"

// Printing all types;
// this will be printed on cells
func (n Number[T]) String() string {
	return fmt.Sprint(n.Val)
}

func (s String) String() string {
	return s.Val
}

func (b Boolean) String() string {
	if b.Val {
		return TRUE_LITERAL
	}
	return FALSE_LITERAL
}

func (a Area) String() string {
	return fmt.Sprintf("area: %v", a.Val)
}

func (e Error) String() string {
	if e.Msg == nil {
		return "#ERR:nil"
	}
	return fmt.Sprintf("#ERR: %s", e.Msg.Error())
}
