// https://go-tour-jp.appspot.com/flowcontrol/8

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	Z := 0.0
	for i := 0; i < 10; i++ {
		Z = z - (z*z-x)/(2*z)

		if math.Abs(Z-z) < 0.0000001 {
			break
		}

		z = Z
	}
	return z
}

func main() {
	x := 3.0
	fmt.Println(Sqrt(x))
	fmt.Println(math.Sqrt(x))
}
