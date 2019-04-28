package main

import (
	"fmt"
	"github.com/jlevitt/katas/pi"
)

func main() {
	n := 1000000000
	result := pi.Pi(n)

	fmt.Printf("Pi(%v) is: %v\n", n, result)
	//3.141592653589793238462643383279502884197169399375105820974944592307816406286
	//3.1415926525880504
}
