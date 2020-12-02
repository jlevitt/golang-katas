package fusc

func Fusc(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n % 2 == 0 {
		return Fusc(n / 2)
	} else {
		n = (n - 1) / 2

		return Fusc(n) + Fusc(n + 1)
	}
}
