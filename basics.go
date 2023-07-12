package main

import (
	"fmt"
)

func main() {
	x, y := 1, 2

	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("y=%v, type of %T\n", y, y)

	mean := float32(x + y) / 2.0

	fmt.Printf("mean=%v, type of %T\n", mean, mean)
}
