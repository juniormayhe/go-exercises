package main

import (
	"fmt"
	"math"

	"github.com/juniormayhe/currentFilename"
)

func compute(myfunc func(float64, float64) float64) float64 {
	return myfunc(3, 4)
}

func main() {
	fmt.Println(currentFilename.GetCurrentFileName())

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println("hypot calculates sqrt(5*5 + 12*12)=", hypot(5, 12))

	fmt.Println("compute(hypot) passes 3, 4 to hypot which calculates sqrt (3*3 + 4*4)=", compute(hypot))

	fmt.Println("compute(math.Pow) uses default values 3^4=", compute(math.Pow))

}
