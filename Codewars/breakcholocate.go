package kata

func BreakChocolate(n, m int) int {
	if (n<1) || (m<1) { return 0 }
	return (m*n)-1
}
