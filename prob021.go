package main
/*
  <title>Amicable Numbers</title>
  <p>Let $d(n)$ be defined as the sum of proper divisors of $n$ (numbers less than $n$ which divide evenly into $n$).<br>
  If $d(a) = b$ and $d(b) = a$, where $a \ne b$, then $a$ and $b$ are an amicable pair and each of $a$ and $b$ are called amicable numbers.</p>
  <p>For example, the proper divisors of $220$ are $1, 2, 4, 5, 10, 11, 20, 22, 44, 55$ and $110$; therefore $d(220) = 284$. The proper divisors of $284$ are $1, 2, 4, 71$ and $142$; so $d(284) = 220$.</p>
  <p>Evaluate the sum of all the amicable numbers under $10000$.</p>
*/

import (
  "fmt"
  "peuler/lib"
)

func main() {
  d := make(map[int]int)

  for i := 1; i < 10_000; i++ {
    d[i] = lib.Sum(lib.Divisors(i))
  }
  //fmt.Println(d)
  s := 0
  for k, v := range d {
    if d[v] == k && v != k {
      fmt.Println(k, v)
      s += k
    }
  }
  fmt.Println(s)
}
