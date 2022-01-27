package main

import (
	"fmt"

	"github.com/juniormayhe/currentFilename"
)

type State struct {
	Name       string
	Population int
}

func main() {
	fmt.Println(currentFilename.GetCurrentFileName())

	// make a map with string keys and struct values
	states := make(map[string]State)
	id := "RJ"
	states[id] = State{Name: "Rio de Janeiro", Population: 13544462}
	fmt.Printf("map=%v\n", states)

	// inline declaration
	states = map[string]State{
		"RJ": {Name: "Rio de Janeiro", Population: 13544462},
		"SP": {Name: "São Paulo", Population: 22237472},
	}

	fmt.Printf("\nmap=%v\n", states)
	fmt.Printf("key=%q, value=%v\n", id, states[id])

	delete(states, id)
	item, isPresent := states[id]
	fmt.Printf("The value: %v %T, Present? %v", item, item, isPresent)

}
