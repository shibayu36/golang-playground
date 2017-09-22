package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for diff := 1.0; math.Abs(diff) > 1e-10; {
		diff = ((z * z) - x) / (2 * z)
		z = z - diff
	}
	return z
}

func main() {
	fmt.Println(Sqrt(3))
	fmt.Println(math.Sqrt(3))
}
