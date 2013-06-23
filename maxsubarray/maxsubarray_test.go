package maxsubarray

import (
	"testing"
)

func TestMaxsubarray(t *testing.T) {
	A := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	t.Log(maxsubarray(A, 0, len(A)))
}
