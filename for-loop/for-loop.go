package main

import (
	"fmt"

	"github.com/juniormayhe/currentFilename"
)

func main() {
	fmt.Println(currentFilename.GetCurrentFileName())

	// make a 5 positions array with 3 elements set to 0 by default
	s := []int{1, 2, 3}
	for i, item := range s {
		// item is an int
		fmt.Printf("index=%d value=%d \n", i, item)
	}
}
