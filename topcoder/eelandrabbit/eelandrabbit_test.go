package eelandrabbit

import (
	"testing"
)

func TestCatchOne(tst *testing.T) {
	l := []int{2, 4, 3, 2, 2, 1, 10}
	t := []int{2, 6, 3, 7, 0, 2, 0}

	m := CatchOne(l, t)
	if m != 4 {
		tst.Errorf("got %d, expected %d", m, 4)
	}
}

func TestCatchTwo(tst *testing.T) {
	l := []int{2, 4, 3, 2, 2, 1, 10}
	t := []int{2, 6, 3, 7, 0, 2, 0}

	m := CatchTwo(l, t)
	if m != 6 {
		tst.Errorf("got %d, expected %d", m, 6)
	}
}

func BenchmarkCatchOne(b *testing.B) {
	l := []int{2, 4, 3, 2, 2, 1, 10}
	t := []int{2, 6, 3, 7, 0, 2, 0}

	for i := 0; i < b.N; i++ {
		CatchOne(l, t)
	}
}

func BenchmarkCatchTwo(b *testing.B) {
	l := []int{2, 4, 3, 2, 2, 1, 10}
	t := []int{2, 6, 3, 7, 0, 2, 0}

	for i := 0; i < b.N; i++ {
		CatchTwo(l, t)
	}
}
