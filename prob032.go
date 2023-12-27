package main
/*
  <title>Pandigital Products</title>
  <p>We shall say that an $n$-digit number is pandigital if it makes use of all the digits $1$ to $n$ exactly once; for example, the $5$-digit number, $15234$, is $1$ through $5$ pandigital.</p>

  <p>The product $7254$ is unusual, as the identity, $39 \times 186 = 7254$, containing multiplicand, multiplier, and product is $1$ through $9$ pandigital.</p>

  <p>Find the sum of all products whose multiplicand/multiplier/product identity can be written as a $1$ through $9$ pandigital.</p>

  <div class="note">HINT: Some products can be obtained in more than one way so be sure to only include it once in your sum.</div>
*/

import (
  "peuler/lib"
  "fmt"
)

func main() {
  digits := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
  res := make(map[int]bool)
  for arr := range lib.Permutations(digits) {
    //fmt.Println(arr)
    for i := 1; i < len(digits)-2; i++ {
      for j := i+1; j < len(digits)-2; j++ {
        mcand := lib.DigitsAsInt(arr[:i])
        mplier := lib.DigitsAsInt(arr[i:j])
        prod := lib.DigitsAsInt(arr[j:])

	if mcand * mplier != prod { continue }
	if _, ok := res[prod]; !ok {
	  fmt.Println(mcand, mplier, prod)
	  res[prod] = true
	}
      }
    }
  }
  s := 0
  for k, _ := range res {
    s += k
  }
  fmt.Println(s)
}
