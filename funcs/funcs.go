package funcs

import "fmt"

func Divmod(x int, y int) (int, int) {
	return x / y, x % y
}

func Worker() {
	fmt.Println("Worker starting...")
	a, err := Acquire("a")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer Release(a)
	fmt.Println("Worker ended")
}

func Acquire(name string) (string, error) {
	return name, nil
}

func Release(name string) {
	fmt.Printf("Cleaning up %s\n", name)
}
