package funcs

import (
	"fmt"
	"math"
)

func Sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0.0, fmt.Errorf("Square root of a negative value %f", n)
	}
	return math.Sqrt(n), nil
}
