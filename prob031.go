package main
/*
  <title>Coin Sums</title>
  <p>In the United Kingdom the currency is made up of pound (£) and pence (p). There are eight coins in general circulation:</p>
  <blockquote>1p, 2p, 5p, 10p, 20p, 50p, £1 (100p), and £2 (200p).</blockquote>
  <p>It is possible to make £2 in the following way:</p>
  <blockquote>1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p</blockquote>
  <p>How many different ways can £2 be made using any number of coins?</p>
*/

import "fmt"

var COINS = []int{1, 2, 5, 10, 20, 50, 100, 200}

// Use DP approach

// 3:  1 1 1
//     1 2
// 
// 4:  1 1 1 1
//     1 1 2
//     2 2
//
// 6:  1 1 1 1 1 1
//     1 1 1 1 2
//     1 1 2 2
//     2 2 2
//
// 5:  1 1 1 1 1
//     1 1 1 2
//     1 2 2
//     5
//
// 6:  1 1 1 1 1 1
//     1 1 1 1 2
//     1 1 2 2
//     2 2 2
//     1 5
//
// 7:  1 1 1 1 1 1 1
//     1 1 1 1 1 2
//     1 1 1 2 2
//     1 2 2 2
//     5 2
//     5 1 1
//
// 8:  1 1 1 1 1 1 1 1
//     1 1 1 1 1 1 2
//     1 1 1 1 2 2
//     1 1 2 2 2
//     2 2 2 2
//     5 1 1 1
//     5 1 2
// 
// 9:  1 1 1 1 1 1 1 1 1    1
//     1 1 1 1 1 1 1 2      2
//     1 1 1 1 1 2 2        3
//     1 1 1 2 2 2          4
//     1 2 2 2 2            5
//     5 1 1 1 1            6
//     5 2 1 1              7
//     5 2 2                8
//
// 10: 1 1 1 1 1 1 1 1 1 1  1
//     1 1 1 1 1 1 1 1 2    2
//     1 1 1 1 1 1 2 2      3
//     1 1 1 1 2 2 2        4
//     1 1 2 2 2 2          5
//     2 2 2 2 2            6
//     5 5                  7
//     5 1 1 1 1 1          8
//     5 1 1 1 2            9
//     5 1 2 2              10
//
//       1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20
// 1: 0  1 1 1 1 1 1 1 1 1  1  1  1  1  1  1  1  1  1  1  1
// 2: 0  1 2 2 3 3 4 4 5 5  6  6  7  7  8  8  9  9 10 10 11
// 5: 0  1 2 2 3 4 5 6 6 8 10 

// So seed 1 way when v == c
// Then  #(previous coin at value)       = dp[c-1][v]
//     + #(current coin at value - c)    = dp[c][v-1]
//     + seed                            = dp[c][v]

type pair struct {
  c, v int
}

type grid struct {
  g map[pair]int
}

func (g grid) Get(p pair) int {
  if p.c < 0 || p.v <= 0 { return 0 }
  return g.g[p]
}

func (g grid) Put(p pair, n int) {
  g.g[p] = n
}

func main() {
  g := grid{make(map[pair]int)}
  // Each coin can make itself
  for ci, c := range COINS {
    g.Put(pair{ci, c}, 1)
  }

  for ci, c := range COINS {
    for v := 0; v <= 200; v++ {
      val := g.Get(pair{ci-1, v}) + g.Get(pair{ci, v-c})
      val += g.Get(pair{ci, v})
      g.Put(pair{ci,v}, val)
    }
  }
  for ci, c := range COINS {
    fmt.Println()
    fmt.Printf("%v, %v:\t", ci, c)
    for v := 0; v <= 10; v++ {
      fmt.Printf("%v ", g.Get(pair{ci, v}))
    }
  }
  fmt.Println()
  fmt.Println(g.Get(pair{len(COINS)-1, 200}))
}
