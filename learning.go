package learning

func Sum(input ...int) int {
  sum := 0
  for _, n := range input {
    sum = sum + n
  }

  return sum
}
