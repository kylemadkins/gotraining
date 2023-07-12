package basics

import (
	"log"
	"regexp"
)

func RemovePunctuation(text string) string {
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	return re.ReplaceAllString(text, " ")
}
