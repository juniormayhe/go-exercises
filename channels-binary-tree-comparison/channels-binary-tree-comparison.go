package main

import (
	"fmt"

	"github.com/juniormayhe/currentFilename"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {

		return
	}

	Walk(t.Left, ch)
	ch <- t.Value

	Walk(t.Right, ch)

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {

	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < 5; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}

	return true

}

func main() {
	fmt.Println(currentFilename.GetCurrentFileName())

	fmt.Printf("\nNew(1) and New(1): %v", Same(tree.New(1), tree.New(1)))
	fmt.Printf("\nNew(1) and New(2): %v", Same(tree.New(1), tree.New(2)))

}
