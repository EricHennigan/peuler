
package main
/*
  <title>Power Digit Sum</title>
  <p>$2^{15} = 32768$ and the sum of its digits is $3 + 2 + 7 + 6 + 8 = 26$.</p>
  <p>What is the sum of the digits of the number $2^{1000}$?</p>
*/

import (
  "fmt"
  "math/big"
  "peuler/lib"
)

const N = 1000

func main() {
  n := big.NewInt(1)
  v := big.NewInt(2)
  exp := N
  fmt.Printf("exp=%32b\t%12v\t%12v\n", exp, v, n)
  for exp > 0 {
    if exp & 1 == 1 {
      n.Mul(n, v)
    }
    fmt.Printf("exp=%32b\t%12v\t%12v\n", exp, v, n)
    v.Mul(v, v)
    exp >>= 1
  }
  fmt.Println(n)
  fmt.Println(lib.DigitSum(n))
}
