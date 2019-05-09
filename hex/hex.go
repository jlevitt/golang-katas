package hex

import (
	"math"
	"strings"
)

var hexToIntMap = map[rune]int {
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'a': 10,
	'b': 11,
	'c': 12,
	'd': 13,
	'e': 14,
	'f': 15,
}

var intToHexMap = map[int]string{}

func init() {
	for k, v := range hexToIntMap {
		intToHexMap[v] = string(k)
	}
}

func baseToInt(n string, base int) int {
	n = strings.ToLower(n)
	baseF := float64(base)

	placeValue := math.Pow(baseF, float64(len(n) - 1))
	var result int
	for _, c := range n {
		result += hexToIntMap[c] * int(placeValue)
		placeValue /= baseF
	}

	return result
}

func intToBase(n int, base int) string {
	baseF := float64(base)
	place := math.Floor(math.Log(float64(n)) / math.Log(baseF))
	placeValue := int(math.Pow(baseF, place))

	var result string
	for placeValue >= 1 {
		quotient := n / placeValue
		n %= placeValue
		result += intToHexMap[quotient]
		placeValue /= base
	}

	return result
}

func HexToBin(hex string) string {
	return intToBase(baseToInt(hex, 16), 2)
}

func BinToHex(bin string) string {
	return intToBase(baseToInt(bin, 2), 16)
}
