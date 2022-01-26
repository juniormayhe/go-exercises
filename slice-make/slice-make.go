package main

import (
	"fmt"

	"github.com/juniormayhe/currentFilename"
)

func main() {
	fmt.Println(currentFilename.GetCurrentFileName())

	// make int array with 3 elements and max length of 5 elements
	// all 3 elements will be set to 0
	s := make([]int, 3, 5)
	printSlice("s", s)

	// we can append more items then capacity
	t := append(s, 1)
	t = append(t, 2)
	t = append(t, 3)
	t = append(t, 4)
	printSlice("t", t) // [0 0 0 1 2 3 4]

	u := s[:5] // [0 0 0 1 2]
	printSlice("u", u)

	u = s[:3] // [0 0 0]
	printSlice("u", u)

	// u still points to underlying array s which has [0 0 0 1 2] so we can set low (3) and high limit (5).
	// a new array is created in a different memory address because we moved the low limit pointer
	v := u[3:5] //  [1 2]
	printSlice("v", v)
}

func printSlice(m string, s []int) {
	fmt.Printf("%v len=%d cap=%d address=%p value=%v\n", m, len(s), cap(s), s, s)
}
