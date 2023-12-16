package lib

import "math/big"

func DigitSum(n *big.Int) int {
  s := 0
  for _, c := range n.String() {
    s += int(c - '0')
  }
  return s
}

func Sum(nums []int) int {
  s := 0
  for _, n := range nums {
    s += n
  }
  return s
}
