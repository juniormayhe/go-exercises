package main

import (
	"fmt"

	"github.com/juniormayhe/currentFilename"
)

type ErrInvalidAge float32

func (err ErrInvalidAge) Error() string {
	return fmt.Sprintf("ERR: age cannot be negative: %.2f", err)
}

func DisplayMessage(age float32) (string, error) {
	if age < 0 {
		return "<invalid age>", ErrInvalidAge(age)
	}
	return fmt.Sprintf("%v", age), nil
}

func main() {
	fmt.Println(currentFilename.GetCurrentFileName())

	fmt.Println(DisplayMessage(40))

	if m, err := DisplayMessage(-2); err != nil {
		fmt.Printf("message=%s\n", m)
		fmt.Printf("error=%s\n", err)
	}

}
