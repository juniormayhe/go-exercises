package main

import (
	"fmt"

	"github.com/juniormayhe/currentFilename"
)

type IReport interface {
	DisplayName()
}

type Person struct {
	Name string
}

func (p *Person) DisplayName() {
	if p == nil {
		fmt.Println("<no name provided>\n")
		return
	}
	fmt.Println(p.Name)
}

func main() {
	fmt.Println(currentFilename.GetCurrentFileName())

	var p *Person = nil
	var i IReport = p
	describe(i)
	i.DisplayName()

	i = &Person{"Julia"}
	describe(i)
	i.DisplayName()

}

func describe(i IReport) {
	fmt.Printf("(%+v, %T)\n", i, i)
}
