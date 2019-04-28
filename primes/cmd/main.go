package main

import (
	"fmt"

	"github.com/jlevitt/katas/primes"
)

func main() {
	// result := BackwardsPrime(10, 20)
	// testify for tests
	result := primes.BackwardsPrime(1, 100000)

	fmt.Printf("Primes: %v\n", result)
}
