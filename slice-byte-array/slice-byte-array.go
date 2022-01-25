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
	// a change slice value changes original byte array
	sp[0][1] = '!'
	printSlice("sp[0] up| ", sp[0])
	printSlice("b up	| ", b)
}

func printSlice(m string, b []byte) {
	fmt.Printf("%v len=%d cap=%d address=%p value=%s\n", m, len(b), cap(b), b, b)
}
