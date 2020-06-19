func ExpressionMatter(a int, b int, c int) int {
  max := a + b + c
  plusTimes := (a + b) * c
  timesPlus := a * (b + c)
  timesTimes := a * b * c

  if plusTimes > max {
    max = plusTimes
  }

  if timesPlus > max {
    max = timesPlus
  }

  if timesTimes > max {
    max = timesTimes
  }

  return max
}
