package main

import (
	"fmt"

	"github.com/juniormayhe/currentFilename"
)

func main() {
	fmt.Println(currentFilename.GetCurrentFileName())

	s := []int{1, 2, 3}
	printSlice("s	| ", s)

	// Slice the array to give it zero length.
	// Elements are not dropped. Still have the same memory address to underlying array
	s = s[:0]
	printSlice("s[:0]	| ", s)

	// extend array to original size,
	// all elements are still there in the same memory address of underlying array
	s = s[:3]
	printSlice("s[:3]	| ", s)

	// when skipping first two values by setting low limit to 2,
	// a new array is created in memory with a different len and capacity

	s = s[2:] // the underlying array is replaced with a new address
	printSlice("s[2:]	| ", s)

}

func printSlice(m string, s []int) {
	fmt.Printf("%v len=%d cap=%d address=%p value=%v\n", m, len(s), cap(s), s, s)
}
