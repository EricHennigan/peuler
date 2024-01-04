package main
/*
  <title>Circular Primes</title>
  <p>The number, $197$, is called a circular prime because all rotations of the digits: $197$, $971$, and $719$, are themselves prime.</p>
  <p>There are thirteen such primes below $100$: $2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79$, and $97$.</p>
  <p>How many circular primes are there below one million?</p>
*/

import (
  "fmt"
  "peuler/lib"
)

var primes map[int]bool

func isCircular(n int) bool {
  d := lib.AsDigits(n)
  for r := range lib.Rotations(d) {
    dr := lib.DigitsAsInt(r)
    if _, ok := primes[dr]; !ok {
      return false
    }
  }
  return true
}

func main() {
  primes = make(map[int]bool)
  for p := range lib.PrimesUpTo(1_000_000) {
    primes[p] = true
  }

  count := 0
  for p, _ := range primes {
    if isCircular(p) {
      fmt.Println(p)
      count++
    }
  }
  fmt.Println()
  fmt.Println(count)
}
