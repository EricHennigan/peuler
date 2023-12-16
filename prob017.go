package main
/*
  <title>Number Letter Counts</title>
  <p>If the numbers $1$ to $5$ are written out in words: one, two, three, four, five, then there are $3 + 3 + 5 + 4 + 4 = 19$ letters used in total.</p>
  <p>If all the numbers from $1$ to $1000$ (one thousand) inclusive were written out in words, how many letters would be used? </p>
  <br><p class="note"><b>NOTE:</b> Do not count spaces or hyphens. For example, $342$ (three hundred and forty-two) contains $23$ letters and $115$ (one hundred and fifteen) contains $20$ letters. The use of "and" when writing out numbers is in compliance with British usage.</p>
*/

import (
  "fmt"
  "strings"
)

func english(n int) string {
  words := map[int]string{
    0: "",
    1: "one",
    2: "two",
    3: "three",
    4: "four",
    5: "five",
    6: "six",
    7: "seven",
    8: "eight",
    9: "nine",
    10: "ten",
    11: "eleven",
    12: "twelve",
    13: "thirteen",
    14: "fourteen",
    15: "fifteen",
    16: "sixteen",
    17: "seventeen",
    18: "eighteen",
    19: "nineteen",
    20: "twenty",
    30: "thirty",
    40: "forty",
    50: "fifty",
    60: "sixty",
    70: "seventy",
    80: "eighty",
    90: "ninety",
    100: "hundred",
  }
  d0 := n % 10
  n /= 10
  d1 := n % 10
  n /= 10
  d2 := n % 10
  // fmt.Println(d2, d1, d0)

  s := ""
  if d2 > 0 {
    s += words[d2] + " hundred "
  }
  if d1 == 0 && d0 == 0 {
    return s
  }
  if d2 > 0 {
    s += "and "
  }
  if d1 >= 2 {
    s += words[d1 * 10]
    d1 = 0
  }
  s += words[d1 * 10 + d0]
  return s
}

func main() {
  count := 0
  for i := 1; i <= 999; i++ {
    s := english(i)
    c := len(s) - strings.Count(s, " ")
    fmt.Printf("%3v | %3v '%v'\n", i, c, s)
    count += c
  }
  s := "one thousand"
  c := len(s) - 1
  count += c
  fmt.Printf("%3v | %3v '%v'\n", 1000, c, s)
  fmt.Println(count)
}
