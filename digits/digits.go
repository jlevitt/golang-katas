package digits

import (
	"math"
)

func DigPow(n, p int) int {
	sum := getPowSum(getDigits(n), float64(p))

	if sum % n == 0 {
		return sum / n
	} else {
		return -1
	}

}

func getDigits(n int) []int {
	var result []int
	for n > 0 {
		mod := n % 10
		n -= mod
		n /= 10
		result = append([]int{mod}, result...)
	}

	return result
}

func getPowSum(nums []int, pow float64) int {
	var result float64

	for _, num := range nums {
		result += math.Pow(float64(num), pow)
		pow++
	}

	return int(result)
}
