package lib

import "math/big"

func DigitSum(n *big.Int) int {
  s := 0
  for _, c := range n.String() {
    s += int(c - '0')
  }
  return s
}
