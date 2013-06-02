package stack

import (
	"fmt"
)

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value interface{}
	next  *Element
}

func (s *Stack) String() string {
	return fmt.Sprintf("%T size=%v, top=[%v]", s, s.Len(), s.top)
}

func (e *Element) String() string {
	return fmt.Sprintf("%v", e.value)
}

func (s *Stack) Len() int {
	return s.size
}

func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}

func (s *Stack) Pop() (value interface{}) {
	if s.Len() > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}

	return nil
}
