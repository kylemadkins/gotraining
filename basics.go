package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func fizzBuzz(start int, end int) {
	for i := start; i <= end; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizz buzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func isEvenEnded(n int) bool {
	str := fmt.Sprint(n)
	return str[0] == str[len(str)-1]
}

func max(list []int) int {
	max := list[0]
	for _, v := range list[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func wordCount(text string) map[string]int {
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	text = re.ReplaceAllString(text, " ")
	words := map[string]int{}
	for _, word := range strings.Fields(text) {
		words[strings.ToLower(word)] += 1
	}
	return words
}

func main() {
	fizzBuzz(1, 20)
	fmt.Println(isEvenEnded(42))
	fmt.Println(isEvenEnded(101))
	fmt.Println(max([]int{1, 2, 3, 9, 1}))
	fmt.Println(wordCount("Failure to observe what is in the mind of another has seldom made a man unhappy; but those who do not observe the movements of their own minds must of necessity be unhappy."))
}
