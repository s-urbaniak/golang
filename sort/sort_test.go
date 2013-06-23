package sort

import (
	"fmt"
	"testing"
)

type TestData struct {
	data []int
}

func (d *TestData) Len() int {
	return len(d.data)
}

func (d *TestData) Less(i, j int) bool {
	return d.data[i] < d.data[j]
}

func (d *TestData) Swap(i, j int) {
	d.data[i], d.data[j] = d.data[j], d.data[i]
}

func TestInsertionsort(t *testing.T) {
	A := []int{10, 2, 9, 4, 1, 5, 90, 10}
	qsort(&TestData{A}, 0, len(A))
	fmt.Println(A)
}
