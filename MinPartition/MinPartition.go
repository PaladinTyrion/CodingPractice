// 一个整数可以由若干个整数的平方和表示，求最小划分。

package main

import (
	"fmt"
	"math"
)

func MatrixR(x uint64) (matrix []uint64) {
	if x == 1 {
		return []uint64{x}
	}

	y := math.Sqrt(float64(x))
	z := uint64(y)
	matrix = append(matrix, z)

	x = x - (z * z)
	if x >= 1 {
		matrix = append(matrix, MatrixR(x)[:]...)
	}

	return matrix
}

func main() {
	var x uint64 = 523452345
	a := MatrixR(x)
	fmt.Println(a)
}
