package main

/*
  <title>10001st Prime</title>
  <p>By listing the first six prime numbers: $2, 3, 5, 7, 11$, and $13$, we can see that the $6$th prime is $13$.</p>
  <p>What is the $10\,001$st prime number?</p>
*/

import (
  "peuler/lib"
  "fmt"
)

func main() {
  stop := 10_001
  i := 1
  for p := range lib.Primes() {
    if i == stop {
      fmt.Println(p)
      return
    }
    i++
  }
}

