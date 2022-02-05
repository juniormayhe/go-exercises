package main

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/juniormayhe/currentFilename"
)

type Person struct {
	Name   string
	Salary float64
}

func (p Person) FormatSalary() string {
	//return fmt.Sprintf("Salary=%.2f USD", p.Salary)
	return fmt.Sprintf("Salary=%s USD", humanize.FormatFloat("#,###.##", p.Salary))
}

func (pp *Person) ApplyRaise(percentage float64) {
	pp.Salary += pp.Salary * (percentage / 100)
}

func main() {
	fmt.Println(currentFilename.GetCurrentFileName())

	p := Person{"Junior", 90850.77}
	raisePercentage := 15.0
	fmt.Println(p.FormatSalary())
	fmt.Printf("Apply %.2f%% raise\n", raisePercentage)
	p.ApplyRaise(raisePercentage)
	fmt.Println(p.FormatSalary())
}
