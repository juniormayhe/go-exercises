package main

import (
	"fmt"
	"reflect"

	"github.com/juniormayhe/currentFilename"
)

type IString interface{} // empty interface

func main() {
	fmt.Println(currentFilename.GetCurrentFileName())

	var i IString = "hello" // interface value accepts only type string from now on

	//type assertion, where we can get the interface value by setting the type T (string)
	text := i.(string)
	fmt.Printf("text=%v\n", text)

	text, ok := i.(string)
	fmt.Printf("text=%v ok=%v\n", text, ok)

	fmt.Printf("i typeof == string = %v\n", reflect.TypeOf(i).String() == "string")

	number, ok := i.(float64)
	fmt.Printf("number=%v ok=%v\n\n", number, ok)

	//number = i.(float64) // panic: interface conversion: main.IString is string, not float64
	//fmt.Println(number)

	do(21)
	do("hello")
	do(true)
}

func do(anyValue interface{}) {
	switch value := anyValue.(type) {
	case int:
		fmt.Printf("do(%v) where type of interface value is number=%T\n", anyValue, value)
	case string:
		fmt.Printf("do(%v) where type of interface value is text=%T\n", anyValue, value)
	default:
		fmt.Printf("I don't know about type %T!\n", value)
	}

}
