package lib

import "math/big"

func DigitSum(n *big.Int) int {
  s := 0
  for _, c := range n.String() {
    s += int(c - '0')
  }
  return s
}

// produces the reverse of the digits
// The number 0 yields []
func RDigits(n int) []int {
  var d []int
  for n > 0 {
    d = append(d, n % 10)
    n /= 10
  }
  return d
}

func Sum(nums []int) int {
  s := 0
  for _, n := range nums {
    s += n
  }
  return s
}
