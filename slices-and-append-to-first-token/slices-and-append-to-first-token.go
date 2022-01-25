package main

import (
	"fmt"

	"github.com/juniormayhe/currentFilename"
)

func main() {
	fmt.Println(currentFilename.GetCurrentFileName())

	s := []int{1, 2, 3}
	t := s

	// t and s point to the same array in memory
	s[0] = 100
	printSlice("s	| ", s)
	printSlice("t	| ", t)

	// append allocates a new array because currenct capacity is not enough
	s = append(s, 4)
	printSlice("s appen	| ", s) // a copy of s is created
	printSlice("t	| ", t)       // points to original array address

	// s and t do not share same data anymore
}

func printSlice(m string, s []int) {
	fmt.Printf("%v len=%d cap=%d value=%v\n", m, len(s), cap(s), s)
}
