package main
/*
  <title>$1000$-digit Fibonacci Number</title>
  <p>The Fibonacci sequence is defined by the recurrence relation:</p>
  <blockquote>$F_n = F_{n - 1} + F_{n - 2}$, where $F_1 = 1$ and $F_2 = 1$.</blockquote>
  <p>Hence the first $12$ terms will be:</p>
  \begin{align}
  F_1 &amp;= 1\\
  F_2 &amp;= 1\\
  F_3 &amp;= 2\\
  F_4 &amp;= 3\\
  F_5 &amp;= 5\\
  F_6 &amp;= 8\\
  F_7 &amp;= 13\\
  F_8 &amp;= 21\\
  F_9 &amp;= 34\\
  F_{10} &amp;= 55\\
  F_{11} &amp;= 89\\
  F_{12} &amp;= 144
  \end{align}
  <p>The $12$th term, $F_{12}$, is the first term to contain three digits.</p>
  <p>What is the index of the first term in the Fibonacci sequence to contain $1000$ digits?</p>
*/

import (
  "fmt"
  "math/big"
)

func fib() chan *big.Int {
  c := make(chan *big.Int)
  f0 := big.NewInt(1)
  f1 := big.NewInt(1)
  go func() {
    c <- f0
    c <- f1
    for {
      fn := big.NewInt(0)
      fn.Add(f0, f1)
      c <- fn
      f0, f1 = f1, fn
    }
    close(c)
  }()
  return c
}

func main() {
  idx := 1
  for f := range fib() {
    s := f.String()
    // fmt.Println(idx, f, s)
    if (len(s) >= 1000) { break }
    idx++
  }
  fmt.Println(idx)
}
