package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/kylemadkins/gotraining/basics"
	"github.com/kylemadkins/gotraining/funcs"
	"github.com/kylemadkins/gotraining/oop"
)

func ContentTypes(urls []string) {
	for _, url := range urls {
		ctype, cterr := funcs.ContentType(url)
		if cterr != nil {
			fmt.Println(cterr)
		}
		fmt.Println(ctype)
	}
}

func ContentTypesConcurrent(urls []string) {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			ctype, cterr := funcs.ContentType(url)
			if cterr != nil {
				fmt.Println(cterr)
			}
			fmt.Println(ctype)
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func ContentTypesChannel(urls []string) {
	ch := make(chan string)
	for _, url := range urls {
		go func(url string) {
			ctype, cterr := funcs.ContentType(url)
			if cterr != nil {
				fmt.Println(cterr)
			}
			ch <- ctype
		}(url)
	}

	for range urls {
		ctype := <-ch
		fmt.Println(ctype)
	}
}

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

	ctype, cterr := funcs.ContentType("https://pokeapi.co/api/v2/pokemon/ditto")
	if cterr != nil {
		fmt.Println(cterr)
	}
	fmt.Println(ctype)

	sq, sqerr := oop.NewSquare(1, 1, 10)
	if sqerr != nil {
		fmt.Println(sqerr)
	}
	sq.Move(2, 3)
	fmt.Printf("%+v\n", sq)
	fmt.Println(sq.Area())

	c := &oop.Capper{Wtr: os.Stdout}
	fmt.Fprintln(c, "Hello there")

	intv, interr := oop.Min([]int{-2, -9, 2, 4, 5})
	if interr != nil {
		fmt.Println(interr)
	}
	fmt.Println(intv)
	flv, flerr := oop.Min([]float32{19.2, 8.34, 1.77, 21.0, 0.95})
	if flerr != nil {
		fmt.Println(interr)
	}
	fmt.Println(flv)

	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/ip",
	}
	start := time.Now()
	ContentTypes(urls)
	fmt.Println(time.Since(start))

	start = time.Now()
	ContentTypesConcurrent(urls)
	fmt.Println(time.Since(start))

	start = time.Now()
	ContentTypesChannel(urls)
	fmt.Println(time.Since(start))
}
