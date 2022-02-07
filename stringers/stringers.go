package main

import (
	"fmt"

	"github.com/juniormayhe/currentFilename"
)

type IPAddr [4]byte

// fmt and other packages uses a "String() string" to print values.
func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	fmt.Println(currentFilename.GetCurrentFileName())

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
