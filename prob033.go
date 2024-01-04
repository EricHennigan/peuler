package main
/*
  <title>Digit Cancelling Fractions</title>
  <p>The fraction $49/98$ is a curious fraction, as an inexperienced mathematician in attempting to simplify it may incorrectly believe that $49/98 = 4/8$, which is correct, is obtained by cancelling the $9$s.</p>
  <p>We shall consider fractions like, $30/50 = 3/5$, to be trivial examples.</p>
  <p>There are exactly four non-trivial examples of this type of fraction, less than one in value, and containing two digits in the numerator and denominator.</p>
  <p>If the product of these four fractions is given in its lowest common terms, find the value of the denominator.</p>
*/

import (
  "fmt"
  "math/big"
  "peuler/lib"
)

func removeCommonDigits(a, b []int) ([]int, []int) {
  count := make([]int, 10)
  for _, da := range a {
    count[da] |= 1
  }
  for _, db := range b {
    count[db] |= 2
  }

  // any digit in both now has value 3
  var na []int
  for _, d := range a {
    if count[d] != 3 {
      na = append(na, d)
    }
  }
  var nb []int
  for _, d := range b {
    if count[d] != 3 {
      nb = append(nb, d)
    }
  }
  return na, nb
}

func main() {
  var nums []*big.Rat
  // generate all the fractions
  for numer := 10; numer <= 99; numer++ {
    for denom := numer; denom <= 99; denom++ { // denom bigger than numer to ensure a fraction
      if (numer == denom) { continue } // trivial example
      r1 := big.NewRat(int64(numer), int64(denom))
      nd := lib.AsDigits(numer)
      dd := lib.AsDigits(denom)
      if (nd[1] == 0 && dd[1] == 0) { continue } // trivial example
      nd, dd = removeCommonDigits(nd, dd)
      if len(nd) != 1 || len(dd) != 1 { continue } // no cancellation
      nnum := int64(lib.DigitsAsInt(nd))
      dnum := int64(lib.DigitsAsInt(dd))
      if dnum == 0 { continue } // avoid div by 0
      r2 := big.NewRat(nnum, dnum)
      if r1.String() != r2.String() { continue } // not the same after cancellation
      fmt.Println(numer, denom, r1, nd, dd, r2)
      nums = append(nums, r1)
    }
  }
  fmt.Println(nums)
  z := big.NewRat(1, 1)
  for _, n := range nums {
    z = z.Mul(z, n)
  }
  fmt.Println(z)
  fmt.Println(z.Denom())
}
