package main

import (
	"fmt"

	"github.com/juniormayhe/currentFilename"
)

func tenTimes(n int) int {

	return n * 10
}

func main() {
	fmt.Println(currentFilename.GetCurrentFileName())

	c := make(chan int, 10)
	for i := 0; i < cap(c); i++ {
		c <- tenTimes(i)
	}
	close(c)

	for channelItem := range c {

		fmt.Println("channel item:", channelItem)
	}

	_, isOpen := <-c
	fmt.Printf("isOpen: %v ", isOpen)

}
