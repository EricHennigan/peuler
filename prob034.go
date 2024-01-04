package main
/*
  <title>Digit Factorials</title>
  <p>$145$ is a curious number, as $1! + 4! + 5! = 1 + 24 + 120 = 145$.</p>
  <p>Find the sum of all numbers which are equal to the sum of the factorial of their digits.</p>
  <p class="smaller">Note: As $1! = 1$ and $2! = 2$ are not sums they are not included.</p>
*/

import (
  "fmt"
  "peuler/lib"
)

func factorial(n int) int {
  if n < 0 { panic("cannot take factorial of a negative number") }
  f := 1
  for ;n >= 1; n-- {
    f *= n
  }
  return f
}

func main() {
  f := make([]int, 11)
  for i := 0; i < 11; i++ {
    f[i] = factorial(i)
  }
  fmt.Println(f)
  sum := 0
  for n := 0; n <= f[10]; n++ {
    digits := lib.AsDigits(n)
    v := 0
    for _, d := range digits {
      v += f[d]
    }
    if n == v {
      fmt.Println(n)
      sum += n
    }
  }
  fmt.Println()
  // remove special cases from question
  sum -= 1
  sum -= 2
  fmt.Println(sum)
}
