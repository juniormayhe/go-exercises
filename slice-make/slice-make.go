package main

import (
	"fmt"

	"github.com/juniormayhe/currentFilename"
)

func main() {
	fmt.Println(currentFilename.GetCurrentFileName())

	// make a 5 positions array with 3 elements set to 0 by default
	s := make([]int, 3, 5)
	printSlice("s", s)

	// we can append more items then capacity
	t := append(s, 1, 2, 3, 4)

	printSlice("t", t) // [0 0 0 1 2 3 4]

	u := t[:5] // [0 0 0 1 2]
	printSlice("u", u)

	u = t[:3] // [0 0 0]
	printSlice("u", u)

	// u still points to underlying array t which has [0 0 0 1 2] so we can set low (3) and high limit (5).
	// a new array is created in a different memory address because we moved the low limit pointer
	v := u[3:5] //  [1 2]
	printSlice("v", v)
}

func printSlice(m string, s []int) {
	fmt.Printf("%v len=%d cap=%d address=%p value=%v\n", m, len(s), cap(s), s, s)
}
