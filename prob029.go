package main

/*
  <title>Distinct Powers</title>
  <p>Consider all integer combinations of $a^b$ for $2 \le a \le 5$ and $2 \le b \le 5$:</p>
  \begin{matrix}
  2^2=4, &amp;2^3=8, &amp;2^4=16, &amp;2^5=32\\
  3^2=9, &amp;3^3=27, &amp;3^4=81, &amp;3^5=243\\
  4^2=16, &amp;4^3=64, &amp;4^4=256, &amp;4^5=1024\\
  5^2=25, &amp;5^3=125, &amp;5^4=625, &amp;5^5=3125
  \end{matrix}
  <p>If they are then placed in numerical order, with any repeats removed, we get the following sequence of $15$ distinct terms:
  $$4, 8, 9, 16, 25, 27, 32, 64, 81, 125, 243, 256, 625, 1024, 3125.$$</p>
  <p>How many distinct terms are in the sequence generated by $a^b$ for $2 \le a \le 100$ and $2 \le b \le 100$?</p>
*/

import (
  "fmt"
  "math/big"
)

/* Figuring out collisions in floating point space doesn't work
func main() {
  nums := make(map[float64]bool)
  for a := 2.; a <= 5.; a++ {
    la := math.Log(a)
    for b := 2.; b <= 5.; b++ {
      nums[b * la] = true
    }
  }
  fmt.Println(len(nums))
}
*/

// Next, use big math
func main() {
  a := big.NewInt(2)
  b := big.NewInt(4)
  z1 := big.NewInt(0)
  z1.Exp(a, b, nil)
  z2 := big.NewInt(0)
  z2.Exp(b, a, nil)

  nums := make(map[string]bool)
  for a := 2; a <= 100; a++ {
    ba := big.NewInt(int64(a))
    for b := 2; b <= 100; b++ {
      bb := big.NewInt(int64(b))
      z := big.NewInt(0)
      z.Exp(ba, bb, nil)
      nums[z.String()] = true
    }
  }
  fmt.Println(len(nums))
}
