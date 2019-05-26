package main

import (
	"fmt"

	"github.com/jlevitt/katas/polish"
)

func main() {
	result := polish.Evaluate("5 1 2 + 4 * + 3 -")
	fmt.Printf("Result: %v\n", result)
}
