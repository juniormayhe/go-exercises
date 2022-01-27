package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	countByWord := make(map[string]int)

	words := strings.Fields(s)

	fmt.Printf("words=%q\n", words)

	for _, word := range words {

		fmt.Printf("key=%v\n", word)

		countByWord[word] += 1

	}

	return countByWord
}

func main() {
	wc.Test(WordCount)
}
