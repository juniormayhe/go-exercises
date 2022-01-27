package main

import (
	"fmt"
	"strings"

	"github.com/juniormayhe/currentFilename"
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
	fmt.Println(currentFilename.GetCurrentFileName())
	wc.Test(WordCount)
}
