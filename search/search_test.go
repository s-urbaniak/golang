package search

import (
	"sort"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	A := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	f := func(i int) bool {
		return A[i] >= 2
	}

	i := sort.Search(len(A), f)
	t.Log("Result:", i)
}
