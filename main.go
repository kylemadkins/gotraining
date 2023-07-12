package main

import (
	"fmt"

	"github.com/kylemadkins/gotraining/basics"
	"github.com/kylemadkins/gotraining/funcs"
)

func main() {
	basics.FizzBuzz(1, 20)

	fmt.Println(basics.IsEvenEnded(42))
	fmt.Println(basics.IsEvenEnded(101))

	fmt.Println(basics.Max([]int{1, 2, 3, 9, 1}))

	fmt.Println(basics.WordCount("Failure to observe what is in the mind of another has seldom made a man unhappy; but those who do not observe the movements of their own minds must of necessity be unhappy."))

	fmt.Println(funcs.Divmod(9, 4))

	nv, nerr := funcs.Sqrt(-2)
	if nerr != nil {
		fmt.Println(nerr)
	}
	fmt.Println(nv)
	v, err := funcs.Sqrt(25)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)

	funcs.Worker()
}
