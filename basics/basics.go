package basics

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func FizzBuzz(start int, end int) {
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

func IsEvenEnded(n int) bool {
	str := fmt.Sprint(n)
	return str[0] == str[len(str)-1]
}

func Max(list []int) int {
	max := list[0]
	for _, v := range list[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func WordCount(text string) map[string]int {
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
