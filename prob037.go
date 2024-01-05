package main
/*
  <title>Truncatable Primes</title>
  <p>The number $3797$ has an interesting property. Being prime itself, it is possible to continuously remove digits from left to right, and remain prime at each stage: $3797$, $797$, $97$, and $7$. Similarly we can work from right to left: $3797$, $379$, $37$, and $3$.</p>
  <p>Find the sum of the only eleven primes that are both truncatable from left to right and right to left.</p>
  <p class="smaller">NOTE: $2$, $3$, $5$, and $7$ are not considered to be truncatable primes.</p>
*/

import (
  "fmt"
  "peuler/lib"
)

var primes map[int]bool

func isPrime(n int) bool {
  if n == 0 { return true }
  _, ok := primes[n]
  //fmt.Println("prime? ", n, ok)
  return ok
}

func isTruncatable(n int) bool {
  digits := lib.AsDigits(n)
  s := len(digits)
  for i := 0; i < s; i++ {
    if !isPrime(lib.DigitsAsInt(digits[i:])) {
      //fmt.Println(digits[i:], "is not prime")
      return false
    }
    if !isPrime(lib.DigitsAsInt(digits[:s-i-1])) {
      //fmt.Println(digits[:s-i-1], "is not prime")
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
  fmt.Println(isTruncatable(3797))
  fmt.Println()

  sum := 0
  for p, _ := range primes {
    if isTruncatable(p) {
      fmt.Println(p)
      sum += p
    }
  }
  fmt.Println()
  // remove special cases
  sum -= 2
  sum -= 3
  sum -= 5
  sum -= 7
  fmt.Println(sum)
}
