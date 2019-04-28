package primes

import (
	"fmt"
	"strconv"
)

//BackwardsPrimes
func BackwardsPrime(start int, stop int) []int {
	primes := []int{}
	backwardsPrimes := map[int]bool{}
	for i := start; i <= stop; i++ {
		if _, ok := backwardsPrimes[i]; ok {
			continue
		}

		backwards := reverse(i)

		if i != backwards && isPrime(i) && isPrime(backwards) {
			primes = append(primes, i)
			backwardsPrimes[i] = true
		}
	}

	if len(primes) == 0 {
		fmt.Println("Returning nil")

		return nil
	} else {
		return primes
	}
}

func isPrime(n int) bool {
	half := n / 2

	for i := 2; i <= half; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func reverse(i int) int {
	str := strconv.Itoa(i)
	runes := []rune(str)
	for start, end := 0, len(str)-1; start < end; start, end = start+1, end-1 {
		runes[start], runes[end] = runes[end], runes[start]
	}

	result, _ := strconv.Atoi(string(runes))

	return result
}
