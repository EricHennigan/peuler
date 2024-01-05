package main
/*
  <title>Pandigital Multiples</title>
  <p>Take the number $192$ and multiply it by each of $1$, $2$, and $3$:</p>
  \begin{align}
  192 \times 1 &amp;= 192\\
  192 \times 2 &amp;= 384\\
  192 \times 3 &amp;= 576
  \end{align}
  <p>By concatenating each product we get the $1$ to $9$ pandigital, $192384576$. We will call $192384576$ the concatenated product of $192$ and $(1,2,3)$.</p>
  <p>The same can be achieved by starting with $9$ and multiplying by $1$, $2$, $3$, $4$, and $5$, giving the pandigital, $918273645$, which is the concatenated product of $9$ and $(1,2,3,4,5)$.</p>
  <p>What is the largest $1$ to $9$ pandigital $9$-digit number that can be formed as the concatenated product of an integer with $(1,2, \dots, n)$ where $n \gt 1$?</p>
*/

import (
  "peuler/lib"
  "fmt"
)

func isPanDigital(n []int) bool {
  digits := make(map[int]bool)
  for _, d := range n {
    if digits[d] == true {
      return false
    } else {
      digits[d] = true
    }
  }
  return len(digits) == 9 && digits[0] == false
}

func main() {
  nums := make(map[int]int)
  for n := 1; n < 100_000; n++ {
    var digits []int
    for p := 1; p <= 9 && len(digits) < 9; p++ {
      for _, d := range lib.AsDigits(n*p) {
        digits = append(digits, d)
      }
      num := lib.DigitsAsInt(digits)
      if isPanDigital(digits) {
        nums[num] = n
      }
    }
  }
  mx := 0
  for k, _ := range nums {
    if k > mx {
      mx = k
    }
  }
  fmt.Println(mx)
}
