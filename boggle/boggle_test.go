package boggle

import (
	"testing"
)

func AssertPos(t *testing.T, p Pos, x, y int) {
	if p.x != x || p.y != y {
		t.Error(p, " != ", x, y)
	}
}

func AssertClosed(t *testing.T, p chan Pos) {
	_, ok := <-p

	if ok {
		t.Error("channel is not closed")
	}
}

func TestNew(t *testing.T) {
	b := New([]string{"abc", "def", "ghi"})
	t.Log(b)
	if b == nil {
		t.Error("expecting non nil value")
	}

	i := b.IterNeighbours(Pos{0, 0})
	AssertPos(t, <-i, 1, 0)
	AssertPos(t, <-i, 0, 1)
	AssertPos(t, <-i, 1, 1)
	AssertClosed(t, i)

	i = b.IterNeighbours(Pos{2, 2})
	AssertPos(t, <-i, 1, 1)
	AssertPos(t, <-i, 2, 1)
	AssertPos(t, <-i, 1, 2)
	AssertClosed(t, i)

	i = b.IterNeighbours(Pos{3, 2})
	AssertClosed(t, i)

	b = New([]string{"abcd"})
	if b == nil {
		t.Error("expecting non nil value")
	}

	b = New([]string{"abcd", "def", "ghi"})
	if b != nil {
		t.Error(b, " != nil")
	}

	b = New([]string{})
	if b != nil {
		t.Error(b, " != nil")
	}
}
