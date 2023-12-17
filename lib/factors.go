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
  hasDup := 0
  for i := 1; i <= stop; i++ {
    q, r := n/i, n%i
    if r == 0 {
      p := pair{i, q}
      if i == q { hasDup = 1 }
      pairs = append(pairs, p)
    }
  }
  s := len(pairs)*2 - hasDup
  factors := make([]int, s)
  for i, p := range pairs {
    factors[i] = p.a
    factors[s-1-i] = p.b
  }
  return factors
}

func Divisors(n int) []int {
  f := Factors(n)
  return f[0:len(f)-1]
}

