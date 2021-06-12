package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z, s := 1.0, 0.0
	for math.Abs(z-s) >= 1e-15 {
		s = z
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
