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
		return 0, ErrNegativeSqrt(-2)
	}
		
	z := x / 2
	for i := 0; true; i++ {
		if math.Abs(z*z-x) < 0.0000001 {
			return z, nil
		}

		z -= (z*z - x) / (2 * z)
		//fmt.Printf("i: %v, z: %v\n", i, z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
