package lib

import "math"

type pair struct {
  a, b int
}

// List the factors of N.
// Ex 28: 1, 2, 4, 7, 14, 28
func Factors(n int) []int {
  var pairs []pair
  stop := int(math.Sqrt(float64(n)))
  for i := 1; i <= stop; i++ {
    q, r := n/i, n%i
    if r == 0 {
      pairs = append(pairs, pair{i, q})
    }
  }
  s := len(pairs)*2
  factors := make([]int, s)
  for i, p := range pairs {
    factors[i] = p.a
    factors[s-1-i] = p.b
  }
  return factors
}
