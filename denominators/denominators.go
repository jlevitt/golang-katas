package denominators

func ConvertFracts(a [][]int) string {
	// your code
	return ""
}

func Lcm(x, y int) int {
	xMultiple := x
	yMultiple := y

	for xMultiple != yMultiple {
		if xMultiple > yMultiple {
			yMultiple += y
		} else {
			xMultiple += x
		}
	}

	return xMultiple
}

func BetterLcm(nums ...int) {
	
}

type PrimeFactor struct {
	factor, power int
}

func PrimeFactors(n int) []int {
	half := n / 2

	var factors []int

	for denom := 2; denom < half; denom++ {
		if n%denom == 0 {
			factors = append(factors, PrimeFactors(denom)...)
			factors = append(factors, PrimeFactors(n/denom)...)
			break
		}
	}

	if len(factors) == 0 {
		factors = append(factors, n)
	}

	return factors
}

func GroupFactors(factors []int) *[]PrimeFactor {
	factorMap := make(map[int]*PrimeFactor)

	for _, factor := range factors {
		if primeFactor, ok := factorMap[factor]; ok {
			primeFactor.power++
		} else {
			factorMap[factor] = &PrimeFactor{
				factor: factor,
				power:  1,
			}
		}
	}

	var result []PrimeFactor
	for _, v := range factorMap {
		result = append(result, *v)
	}

	return &result
}
