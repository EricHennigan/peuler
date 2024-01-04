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
func RDigits(n int) []int {
  if n == 0 { return []int{0} }
  var d []int
  for n > 0 {
    d = append(d, n % 10)
    n /= 10
  }
  return d
}

// turn a number into a digit list
func AsDigits(n int) []int {
  d := RDigits(n)
  for i, j := 0, len(d)-1; i < j; i, j = i+1, j-1 {
   d[i], d[j] = d[j], d[i]
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

// Concatenates digits into a number
func DigitsAsInt(digits []int) int {
  num, t := 0, 1
  for i := len(digits)-1; i >= 0; i-- {
    num += digits[i] * t
    t *= 10
  }
  return num
}
