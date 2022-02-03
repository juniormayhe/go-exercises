package main

import (
	"fmt"

	"github.com/juniormayhe/currentFilename"
)

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	// start values
	previousTerm, nextTerm := 0, 1

	return func() int {
		aux := previousTerm

		previousTerm += nextTerm
		nextTerm = aux

		fmt.Printf("previousTerm=%d nextTerm=%d => ", previousTerm, nextTerm)

		return aux
	}

}

func main() {
	fmt.Println(currentFilename.GetCurrentFileName())

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

}
