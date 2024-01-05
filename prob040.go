package main
/*
  <title>Champernowne's Constant</title>
  <p>An irrational decimal fraction is created by concatenating the positive integers:
  $$0.12345678910{\color{red}\mathbf 1}112131415161718192021\cdots$$</p>
  <p>It can be seen that the $12$<sup>th</sup> digit of the fractional part is $1$.</p>
  <p>If $d_n$ represents the $n$<sup>th</sup> digit of the fractional part, find the value of the following expression.
  $$d_1 \times d_{10} \times d_{100} \times d_{1000} \times d_{10000} \times d_{100000} \times d_{1000000}$$</p>
*/

import (
  "fmt"
  "peuler/lib"
)

func champernowne() chan int {
  c := make(chan int)
  generate := func() {
    for n := 1; ; n++ {
      for _, d := range lib.AsDigits(n) {
        c <- d
      }
    }
  }
  go generate()
  return c
}

func main() {
  i := 1
  t := 1
  c := champernowne()
  val := 1
  for d := range c {
    //fmt.Println(i, d)
    if i > 1_000_000 { break }
    if i % t == 0 {
      t *= 10
      fmt.Println("\t", i, d)
      val *= d
    }
    i++
  }
  close(c)
  fmt.Println(val)
}
