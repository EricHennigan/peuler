package main

/*
  <title>Largest Prime Factor</title>
  <p>The prime factors of $13195$ are $5, 7, 13$ and $29$.</p>
  <p>What is the largest prime factor of the number $600851475143$?</p>
*/

import (
  "log"
  "math"
  "peuler/lib"
)

func main() {
  stop := int(math.Sqrt(600851475143))
  for p := range lib.Primes() {
    if 600851475143 % p == 0 {
      log.Println(p)
    }
    if p > stop {
      return
    }
  }
}

