package sort

import (
	"sort"
)

func isort(A sort.Interface, a, b int) {
	for j := a + 1; j < b; j++ {
		for i := j; i > a && A.Less(i, i-1); i-- {
			A.Swap(i, i-1)
		}
	}
}

func partition(A sort.Interface, a, b int) int {
	b--
	i := a-1

	for j := a; j < b; j++ {
		if A.Less(j, b) {
			i++
			A.Swap(i, j)
		}
	}

	i++
	A.Swap(i, b)
	return i
}

func qsort(A sort.Interface, a, b int) {
	if a < b-1 {
		p := partition(A, a, b)
		qsort(A, a, p)
		qsort(A, p+1, b)
	}
}
