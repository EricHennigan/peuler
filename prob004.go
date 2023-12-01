package main

/*
  <title>Largest Palindrome Product</title>
  <p>A palindromic number reads the same both ways. The largest palindrome made from the product of two $2$-digit numbers is $9009 = 91 \times 99$.</p>
  <p>Find the largest palindrome made from the product of two $3$-digit numbers.</p>
*/

import (
  "fmt"
  "log"
)

func isPalindrome(s string) bool {
  end := len(s) -1
  for i, j := 0, end; i <= j; {
    if s[i] != s[j] { return false }
    i++
    j--
  }
  return true
}

func main() {
  N := 0
  for i := 0; i < 1000; i++ {
    for j := 0; j < 1000; j++ {
      n := i * j
      s := fmt.Sprintf("%d", n)
      if !isPalindrome(s) { continue }
      if (n > N) { N = n }
    }
  }
  log.Printf("%d", N)
}

