package stack

import (
	"testing"
)

func testStackState(s *Stack, t *testing.T, expectedSize int, expectedValue interface{}) {
	if s.Len() != expectedSize {
		t.Error(s.Len(), " != ", expectedSize)
	}

	if expectedSize > 0 {
		if s.top == nil {
			t.Error("s.top is nil")
		}

		if s.top.value != expectedValue {
			t.Error(s.top.value, " != ", expectedValue)
		}
	} else if s.top != nil {
		t.Error("s.top == ", s.top)
	}
}

func testPop(s *Stack, t *testing.T, expectedSize int, expectedValue interface{}) {
	v := s.Pop()

	if v != expectedValue {
		t.Error(v, " != ", expectedValue)
	} else {
		if expectedSize > 0 {
			testStackState(s, t, expectedSize, s.top.value)
		} else {
			testStackState(s, t, expectedSize, nil)
		}
	}
}

func testPush(s *Stack, t *testing.T, expectedSize int, expectedValue interface{}) {
	s.Push(expectedValue)
	testStackState(s, t, expectedSize, expectedValue)
}

func TestPush(t *testing.T) {
	s := new(Stack)
	s.Push("hello")
	testStackState(s, t, 1, "hello")
}

func TestPop(t *testing.T) {
	s := new(Stack)
	testPush(s, t, 1, "hello")
	testPush(s, t, 2, "hello2")
	t.Log(s)
	testPop(s, t, 1, "hello2")
	testPop(s, t, 0, "hello")
}
