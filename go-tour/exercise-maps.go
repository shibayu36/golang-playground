package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordToCount := make(map[string]int)
	for _, word := range words {
		wordToCount[word]++
	}
	return wordToCount
}

func main() {
	wc.Test(WordCount)
}
