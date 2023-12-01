package main
/*
  <title>Sum Square Difference</title>
  <p>The sum of the squares of the first ten natural numbers is,</p>
  $$1^2 + 2^2 + ... + 10^2 = 385.$$
  <p>The square of the sum of the first ten natural numbers is,</p>
  $$(1 + 2 + ... + 10)^2 = 55^2 = 3025.$$
  <p>Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is $3025 - 385 = 2640$.</p>
  <p>Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.</p>
*/

import (
  "fmt"
  "math/big"
)

func main() {
  N := big.NewInt(100)
  sum, sum_sq := big.NewInt(0), big.NewInt(0)
  for i := big.NewInt(1); i.Cmp(N) != 1; i.Add(i, big.NewInt(1)) {
    var sq big.Int
    sq.Exp(i, big.NewInt(2), nil)
    sum_sq.Add(sum_sq, &sq)
    sum.Add(sum, i)
  }
  sum.Exp(sum, big.NewInt(2), nil)
  var ans big.Int
  ans.Sub(sum_sq, sum)
  fmt.Println(ans)
}
