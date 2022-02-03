package main

import (
	"fmt"

	"github.com/juniormayhe/currentFilename"
)

func addTen() func() int {
	i := 0              // closure property
	return func() int { //closure inner / anonymous function
		i += 10
		return i
	}
}

func main() {
	fmt.Println(currentFilename.GetCurrentFileName())

	// each variable has its own state, where i has a unique value
	fmt.Println("no variable 10=", addTen()())
	next10 := addTen()
	anotherNext10 := addTen()

	fmt.Println("next 10=", next10())

	fmt.Println("another 10=", anotherNext10())

	next10()
	fmt.Println("next 10=", next10())
	fmt.Println("next 10=", next10())

}
