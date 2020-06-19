package kata

func PositiveSum(numbers []int) int {
  sum := 0
  for _, s := range numbers {
    if s > 0 {
      sum += s
    }
  }
  return sum
}
