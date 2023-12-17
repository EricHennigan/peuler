package main

/*
  <title>Summation of Primes</title>
  <p>The sum of the primes below $10$ is $2 + 3 + 5 + 7 = 17$.</p>
  <p>Find the sum of all the primes below two million.</p>
*/

import (
  "fmt"
  "peuler/lib"
)

const STOP = 2_000_000

func main() {
  sum := 0
  for p := range lib.PrimesUpTo(STOP) {
    sum += p
  }
  fmt.Println(sum)
}

