// https://go-tour-jp.appspot.com/moretypes/23

package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wordCount := make(map[string]int)

	words := strings.Split(s, " ")
	fmt.Println(words)

	for _, w := range words {
		wordCount[w]++
	}
	return wordCount
}

func main() {
	wc.Test(WordCount)
}
