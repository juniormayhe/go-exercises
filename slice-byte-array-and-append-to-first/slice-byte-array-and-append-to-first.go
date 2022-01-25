package main

import (
	"bytes"
	"fmt"

	"github.com/juniormayhe/currentFilename"
)

func main() {
	fmt.Println(currentFilename.GetCurrentFileName())

	b := []byte("10,20")

	sp := bytes.Split(b, []byte(","))

	printSlice("b	| ", b)
	printSlice("sp[0]	| ", sp[0])
	// change slice value does not override original byte array items as in Go < 1.10
	sp[0] = append(sp[0], '!', '!')

	printSlice("sp[0] up| ", sp[0]) // creates a slice in a new memory address
	printSlice("b up	| ", b)        // b items are not overwriten
}

func printSlice(m string, b []byte) {
	fmt.Printf("%v len=%d cap=%d address=%p value=%s\n", m, len(b), cap(b), b, b)
}
