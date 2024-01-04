package main
/*
  <title>Double-base Palindromes</title>
  <p>The decimal number, $585 = 1001001001_2$ (binary), is palindromic in both bases.</p>
  <p>Find the sum of all numbers, less than one million, which are palindromic in base $10$ and base $2$.</p>
  <p class="smaller">(Please note that the palindromic number, in either base, may not include leading zeros.)</p>
*/

import (
  "fmt"
  "strconv"
  "peuler/lib"
)

func main() {
  sum := 0
  for n := 1; n < 1_000_000; n++ {
    digits := lib.AsDigits(n)
    if !lib.IsPalindrome(digits) { continue }
    binary := strconv.FormatInt(int64(n), 2)
    var bdigits []int
    bdigits = make([]int, len(binary))
    for i, b := range binary {
      bdigits[i] = int(b - '0')
    }
    if !lib.IsPalindrome(bdigits) { continue }
    fmt.Println(n, bdigits)
    sum += n
  }
  fmt.Println()
  fmt.Println(sum)
}
