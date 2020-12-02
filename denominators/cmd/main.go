package main

import (
	"fmt"

	"github.com/jlevitt/katas/denominators"
)

func main() {
	result := denominators.Lcm(6, 8)

	fmt.Printf("Mul:%v\n", result)

	primes := denominators.PrimeFactors(18)
	grouped := denominators.GroupFactors(primes)
	fmt.Printf("Primes: %v Grouped: %v\n", primes, grouped)
}
