package main

/*
  <title>Digit Fifth Powers</title>
  <p>Surprisingly there are only three numbers that can be written as the sum of fourth powers of their digits:
  \begin{align}
  1634 &amp;= 1^4 + 6^4 + 3^4 + 4^4\\
  8208 &amp;= 8^4 + 2^4 + 0^4 + 8^4\\
  9474 &amp;= 9^4 + 4^4 + 7^4 + 4^4
  \end{align}
  </p><p class="smaller">As $1 = 1^4$ is not a sum it is not included.</p>
  <p>The sum of these numbers is $1634 + 8208 + 9474 = 19316$.</p>
  <p>Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.</p>
*/

import (
  "fmt"
  "peuler/lib"
)

func main() {
  powers := make([]int, 10)
  for i := 0; i < 10; i++ {
    powers[i] = i*i*i*i*i
  }

  var nums []int
  for n := 2; n <= 1_000_000; n++ {
    digits := lib.RDigits(n)
    for i, d := range digits {
      digits[i] = powers[d]
    }
    if lib.Sum(digits) == n {
      nums = append(nums, n)
      fmt.Println("found", n)
    }
  }
  fmt.Println(lib.Sum(nums))
}