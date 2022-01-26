package main

import (
	"fmt"

	"github.com/juniormayhe/currentFilename"
)

func main() {
	fmt.Println(currentFilename.GetCurrentFileName())

	arr := []int{1, 2, 3, 4, 5}
	printSlice("arr	| ", arr)

	//max elements per chunk
	maxElements := 2
	for i := 0; i < len(arr); i += maxElements {
		max := min(i+maxElements, len(arr))
		chunk := arr[i:max]
		printSlice(fmt.Sprintf("arr[%v:%v]| ", i, max), chunk)
	}

	// since we didn't modify the array in the loop, the elements are the same
	printSlice("arr	| ", arr)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func printSlice(m string, s []int) {
	fmt.Printf("%v len=%d cap=%d address=%p value=%v\n", m, len(s), cap(s), s, s)
}
