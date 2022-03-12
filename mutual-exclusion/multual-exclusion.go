package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/juniormayhe/currentFilename"
)

type SafeCounter struct {
	mutex  sync.Mutex
	number int
}

func (counter *SafeCounter) addOne() {
	counter.mutex.Lock()
	counter.number++
	fmt.Println("counter:", counter.number)
	counter.mutex.Unlock()
}

func (counter *SafeCounter) getNumber() int {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	return counter.number
}

func main() {
	fmt.Println(currentFilename.GetCurrentFileName())
	c := SafeCounter{number: 0}
	for i := 0; i < 10; i++ {
		go c.addOne()
	}

	time.Sleep(time.Second)
	fmt.Println("finished", c.getNumber())

}
