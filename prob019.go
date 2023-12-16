package main
/*
  <title>Counting Sundays</title>
  <p>You are given the following information, but you may prefer to do some research for yourself.</p>
  <ul><li>1 Jan 1900 was a Monday.</li>
  <li>Thirty days has September,<br />
  April, June and November.<br />
  All the rest have thirty-one,<br />
  Saving February alone,<br />
  Which has twenty-eight, rain or shine.<br />
  And on leap years, twenty-nine.</li>
  <li>A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.</li>
  </ul><p>How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?</p>
*/

import (
  "fmt"
  "time"
)

const (
  _ int = iota
  JAN
  FEB
  MAR
  APR
  MAY
  JUN
  JUL
  AUG
  SEP
  OCT
  NOV
  DEC
)

const (
  SUN = iota
  MON
  TUE
  WED
  THU
  FRI
  SAT
)

type date struct {
  year, month, day, w int
}

func isLeap(y int) bool {
  if y % 400 == 0 { return true }
  if y % 100 == 0 { return false }
  if y % 4 == 0 { return true }
  return false
}

func (d date) next() date {
  n := date{d.year, d.month, d.day+1, (d.w + 1)%7}
  if n.month == FEB {
    if isLeap(n.year) && n.day > 29 {
      n.day = 1
      n.month += 1
    } else if !isLeap(n.year) && n.day > 28 {
      n.day = 1
      n.month += 1
    }
  }
  if (n.month == SEP || n.month == APR || n.month == JUN || n.month == NOV) && n.day > 30 {
    n.day = 1
    n.month += 1
  }
  if n.day > 31 {
    n.day = 1
    n.month += 1
  }
  if n.month > DEC {
    n.month = JAN
    n.year += 1
  }
  return n
}

// had to use the built-in datetime lib to debug my function

func main() {
  suns := 0
  d := date{1900, JAN, 1, MON}
  dg := time.Date(1900, time.January, 1, 9, 0, 0, 0, time.UTC)
  for d.year < 2001 {
    if d.w == SUN && d.day == 1 && d.year >= 1901 {
      suns += 1
    }
    fmt.Println(d, dg, dg.Weekday())
    if (d.year != dg.Year() || d.month != int(dg.Month()) || d.day != dg.Day()) { break }
    d = d.next()
    dg = dg.Add(24 * time.Hour)
  }
  fmt.Println(suns)
}
