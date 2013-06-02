package eelandrabbit

import (
	"testing"
)

func NewEelAndRabbit7() *EelAndRabbit {
	l := []int{2, 4, 3, 2, 2, 1, 10}
	t := []int{2, 6, 3, 7, 0, 2, 0}

	return New(l, t)
}

func TestCatchOne(tst *testing.T) {
	m := NewEelAndRabbit7().CatchOne()
	if m != 4 {
		tst.Errorf("got %d, expected %d", m, 4)
	}
}

func TestCatchTwo(tst *testing.T) {
	m := NewEelAndRabbit7().CatchTwo()
	if m != 6 {
		tst.Errorf("got %d, expected %d", m, 6)
	}
}

func BenchmarkCatchOne(b *testing.B) {
	e := NewEelAndRabbit7()
	benchmarkCatch(b, e.CatchOne)
}

func BenchmarkCatchTwo(b *testing.B) {
	e := NewEelAndRabbit7()
	benchmarkCatch(b, e.CatchTwo)
}

func benchmarkCatch(b *testing.B, m EelAndRabbitAlgorithm) {
	for i := 0; i < b.N; i++ {
		m()
	}
}
