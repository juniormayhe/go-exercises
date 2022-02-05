package main

import (
	"fmt"

	"github.com/juniormayhe/currentFilename"
)

type Number int

type Person struct {
	Name string
	Age  int
}

// we can add methods to non-struct types
func (n Number) FormatNumber() string {
	return fmt.Sprintf("My number=%d", n)
}

// we can add methods to struct types
func (p Person) FormatAge() string {
	return fmt.Sprintf("Age=%v years", p.Age)
}

func main() {
	fmt.Println(currentFilename.GetCurrentFileName())

	i := Number(42)
	fmt.Println(i.FormatNumber())

	p := Person{Name: "Junior", Age: 30}
	fmt.Println(p)
	fmt.Println(p.FormatAge())
}
