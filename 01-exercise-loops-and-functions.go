// https://tour.golang.org/flowcontrol/8

package main

import (
	"fmt"
	"math"
)

// TODO: remove second argument
func Sqrt(x float64, i int) float64 {
	z := 1.0
	// TODO: loop until close enough (< 0.00000000000001)
	for n := 1; n < i; n++ {
		z = z - (z*z-x)/(2*z)
	}
	return z
}

func PrintResults(x float64, i int) {
	fmt.Println(Sqrt(x, i), math.Sqrt(x), Sqrt(x, i)-math.Sqrt(x))
}

func main() {
	PrintResults(4001, 20)
}
