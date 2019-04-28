package main

import (
	"fmt"
	"github.com/jlevitt/katas/fib"
)

func main() {
	// result := BackwardsPrime(10, 20)
	// testify for tests
	result := fib.Fib(11)

	fmt.Printf("fib(%v) is: %v\n", 3, result)
}
