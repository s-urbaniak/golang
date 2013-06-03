package boggle

import (
	"testing"
)

func TestNew(t *testing.T) {
	b := New([]string{"abc", "def", "ghi"})
	if b == nil {
		t.Error("expecting non nil value")
	}

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
