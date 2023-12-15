package main
/*
  <title>Factorial Digit Sum</title>
  <p>$n!$ means $n \times (n - 1) \times \cdots \times 3 \times 2 \times 1$.</p>
  <p>For example, $10! = 10 \times 9 \times \cdots \times 3 \times 2 \times 1 = 3628800$,<br>and the sum of the digits in the number $10!$ is $3 + 6 + 2 + 8 + 8 + 0 + 0 = 27$.</p>
  <p>Find the sum of the digits in the number $100!$.</p>
*/

import (
  "fmt"
  "math/big"
)

const N = 100

func main() {
  m := big.NewInt(1)
  m.MulRange(1, N)

  s := m.String()
  v := 0
  for _, a := range s {
    v += int(a - '0')
  }
  fmt.Println(v)
}
