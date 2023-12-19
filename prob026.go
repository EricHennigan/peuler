package main
/*
  <title>Reciprocal Cycles</title>
  <p>A unit fraction contains $1$ in the numerator. The decimal representation of the unit fractions with denominators $2$ to $10$ are given:</p>
  \begin{align}
  1/2 &amp;= 0.5\\
  1/3 &amp;=0.(3)\\
  1/4 &amp;=0.25\\
  1/5 &amp;= 0.2\\
  1/6 &amp;= 0.1(6)\\
  1/7 &amp;= 0.(142857)\\
  1/8 &amp;= 0.125\\
  1/9 &amp;= 0.(1)\\
  1/10 &amp;= 0.1
  \end{align}
  <p>Where $0.1(6)$ means $0.166666\cdots$, and has a $1$-digit recurring cycle. It can be seen that $1/7$ has a $6$-digit recurring cycle.</p>
  <p>Find the value of $d \lt 1000$ for which $1/d$ contains the longest recurring cycle in its decimal fraction part.</p>
*/

/*
  Was looking at a remainder calculator online. Recall that when you do long division the repeating cycle will show up in the remainders.

  Let's try a prime
  1        / 7 =          1
  10       / 7 = 1        3
  100      / 7 = 14       2
  1000     / 7 = 142      6
  10000    / 7 = 1428     4
  100000   / 7 = 14285    5
  1000000  / 7 = 142857   1
  10000000 / 7 = 1428571  3

  Would prime 37 have 36-digit long cycle? No
  1       / 37 =          1
  10      / 37 =         10
  100     / 37 = 2       26
  1000    / 37 = 27       1
  10000   / 37 = 270     10
  100000  / 37 = 2702    26
  1000000 / 37 = 27,027   1

  This number does not cycle at 1, but 10. Interestingly, has the same sequence as 1/7, shifted
  1         / 14 =          1
  10        / 14 =         10
  100       / 14 = 7        2
  1000      / 14 = 71       6
  10000     / 14 = 714      4
  100000    / 14 = 7142    12
  1000000   / 14 = 71428    8
  10000000  / 14 = 714285  10
  100000000 / 14 = 7142857  2

  Let's try one that doesn't repeat
  1         / 80 =         1
  10        / 80 =        10
  100       / 80 = 1      20
  1000      / 80 = 12     40
  10000     / 80 = 125     0
  100000    / 80 = 1250    0

  Or one with a prefix followed by a repeat
  1         / 48 =         1
  10        / 48 =        10
  100       / 48 = 2       4
  1000      / 48 = 20     40
  10000     / 48 = 208    16
  100000    / 48 = 2083   16
  1000000   / 48 = 20833  16

     _0_._1_
  7 / 1 . 0 0 0 0 0
      0 . 0
      -----
      1 . 0
        . 7
	-----
        . 3

So track the long division, until we see a remainder we've seen before
If r==0, then no cycle
Else the cycle is how many steps since that remainder was last seen
*/

import (
  "fmt"
  //"math/big"
)

func main() {
  /* print the decimal expansion of the first 500 numbers
  for i := 1; i < 100; i++ {
    r := big.NewRat(1, int64(i))
    fmt.Printf("%3v\t%v\n", i, r.FloatString(200))
  }
  */

  cycles := make(map[int]int)
  for denom := 1; denom < 1000; denom++ {
    //fmt.Println("DENOM =", denom)
    numer := 1
    remainders := make([]int, denom)
    i, d, r := 1, numer / denom, numer % denom
    for ; remainders[r] == 0 ; i++ {
      remainders[r] = i
      //fmt.Println("\t", d, r, remainders)
      numer = (numer - d * denom) * 10
      d, r = numer / denom, numer % denom
    }
    cycle := i - remainders[r]
    if r == 0 { cycle-- }
    //n := big.NewRat(1, int64(denom))
    //fmt.Printf("\tcycle=%v | i=%v, r=%v, remainders[i]=%v | %v\n", cycle, i, r, remainders[r], n.FloatString(50))
    cycles[denom] = cycle
  }
  max_cycle := -1
  for denom, cycle := range cycles {
    if cycle > max_cycle {
      max_cycle = cycle
      fmt.Println(denom)
    }
  }
}
