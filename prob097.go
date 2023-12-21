package main

/*
  <title>Large Non-Mersenne Prime</title>
  <p>The first known prime found to exceed one million digits was discovered in 1999, and is a Mersenne prime of the form $2^{6972593} - 1$; it contains exactly $2\,098\,960$ digits. Subsequently other Mersenne primes, of the form $2^p - 1$, have been found which contain more digits.</p>
  <p>However, in 2004 there was found a massive non-Mersenne prime which contains $2\,357\,207$ digits: $28433 \times 2^{7830457} + 1$.</p>
  <p>Find the last ten digits of this prime number.</p>
*/



import (
  "fmt"
  "math/big"
)

func main() {
  x := big.NewInt(2)
  x.Exp(x, big.NewInt(7830457), nil)
  x.Mul(x, big.NewInt(28433))
  x.Add(x, big.NewInt(1))
  x.Rem(x, big.NewInt(10_000_000_000))
  fmt.Println(x)
}
