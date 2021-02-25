// https://go-tour-jp.appspot.com/methods/20

package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0
	Z := 0.0
	for i := 0; i < 10; i++ {
		Z = z - (z*z-x)/(2*z)

		if math.Abs(Z-z) < 0.0000001 {
			break
		}

		z = Z
	}
	return z, nil
}

func main() {
	x := -3.0
	z, err := Sqrt(x)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(z)
}

// >result
// cannot Sqrt negative number: -3
