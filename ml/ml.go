package main

import (
	"fmt"
	"github.com/skelterjohn/go.matrix"
)

func main() {
	A, _ := matrix.ParseMatlab("[1 2 3]")
	B, _ := matrix.ParseMatlab("[4; 5; 6]")

	fmt.Println(matrix.ParallelProduct(A, B))
}
