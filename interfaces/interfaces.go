package main

import (
	"fmt"
	"strings"

	"github.com/juniormayhe/currentFilename"
)

type ICareerEmployee interface {
	PromoteEmployee(percentage float64)
	ToUpper() string
}

type Employee struct {
	Name   string
	Salary float64
}

func (pe *Employee) PromoteEmployee(percentage float64) {
	pe.Salary += pe.Salary * (percentage / 100)

}

func (e Employee) ToUpper() string {
	return strings.ToUpper(e.Name)
}

func main() {
	fmt.Println(currentFilename.GetCurrentFileName())
	e := Employee{"Junior", 90850.77}
	fmt.Printf("employee before promotion=%+v\n", e)

	var ce ICareerEmployee = &e
	raisePercentage := 15.0
	ce.PromoteEmployee(raisePercentage)

	fmt.Printf("employee after promotion=%+v\n", e)

	// our func has a receiver of type Employee
	fmt.Printf("interface e.ToUpper()=%s\n", e.ToUpper())

	// our func also works with the interface that has a type Employee
	fmt.Printf("interface ce.ToUpper()=%s\n", ce.ToUpper())
}
