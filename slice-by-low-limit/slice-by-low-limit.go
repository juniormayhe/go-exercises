package main

import (
	"fmt"

	"github.com/juniormayhe/currentFilename"
)

func main() {
	fmt.Println(currentFilename.GetCurrentFileName())

	arr := []int{1, 2, 3}
	printSlice("arr	| ", arr)

	// Slice the array to give it zero length.
	// The slice points to the memory address of the underlying array.
	s := arr[:0] // high limit is 0
	fmt.Printf("\t|  note: arr address=%p == s address==%p\n", arr, s)
	printSlice("s[:0]	| ", s)

	// extend slice to original length and capacity of the underlying array,
	s = s[:3] // high limit is 3 (we get all arr elements)
	printSlice("s[:3]	| ", s)

	// when skipping first two values by setting low limit to 2,
	// the pointer moves forward to a different memory address and we get a smaller length and capacity compared to underlying array
	s = s[2:] // the pointer to initial position moves to a new memory address
	printSlice("s[2:]	| ", s)
}

func printSlice(m string, s []int) {
	fmt.Printf("%v len=%d cap=%d address=%p value=%v\n", m, len(s), cap(s), s, s)
}
