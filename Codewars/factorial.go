package kata

func Factorial(n int) int {
	if n > 1 {
		return n * Factorial(n-1)
	} else {
		return 1 // 1! == 1
	}
}