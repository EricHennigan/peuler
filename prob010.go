package main

/*
  <title>Summation of Primes</title>
  <p>The sum of the primes below $10$ is $2 + 3 + 5 + 7 = 17$.</p>
  <p>Find the sum of all the primes below two million.</p>
*/

import (
  "fmt"
  "time"
  "peuler/lib"
)

const STOP = 2_000_000

func main() {
  start := time.Now()
  sum := 0
  for p := range lib.Primes() {
    if p >= STOP {
      fmt.Println(sum, time.Since(start))
      return
    }
    sum += p
  }
}

