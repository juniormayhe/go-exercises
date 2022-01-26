package main

import (
	"fmt"
	"strings"

	"github.com/juniormayhe/currentFilename"
)

func main() {
	fmt.Println(currentFilename.GetCurrentFileName())

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// [line, column]
	board[0][0] = "1"
	board[0][2] = "5"

	board[1][0] = "4"
	board[1][2] = "3"

	board[2][2] = "2"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// len(board) = 4 rows
	printSlice("board", board)

}

func printSlice(m string, s [][]string) {
	fmt.Printf("%v len=%d cap=%d address=%p value=%v\n", m, len(s), cap(s), s, s)
}
