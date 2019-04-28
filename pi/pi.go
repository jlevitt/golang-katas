package pi

func Pi(n int) float64 {
	var	result float64
	var multiplier float64 = 1
	var denominator float64 = 1

	for i := 0; i < n; i++ {
		result += multiplier * 4 / denominator
		multiplier *= -1
		denominator += 2
	}

	return result
}
