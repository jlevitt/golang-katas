package numfilter

import (
	"math"
	"strconv"
)

func ShortenNumber(postfixes []string, base int) func (string) string {
	return func(n string) string {
		baseF := float64(base)
		num, err := strconv.Atoi(n)
		if err != nil {
			return n
		}

		place := math.Floor(math.Log(float64(num)) / math.Log(baseF))
		place = math.Min(place, float64(len(postfixes) - 1))

		if place < 0 {
			return n
		}

		num = int(float64(num) / math.Pow(baseF, place))

		result := strconv.Itoa(num)
		result += postfixes[int(place)]

		return result
	}
}

