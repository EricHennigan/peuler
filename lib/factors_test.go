package lib

import (
  "testing"
  "reflect"
)

func TestFactors(t *testing.T) {
  expected := []int{1, 2, 4, 7, 14, 28}
  factors := Factors(28)
  if len(expected) != len(factors) {
    t.Fatalf("incorrect factors for 28: %v\n", factors)
  }
  for i, e := range expected {
    if e != factors[i] {
      t.Fatalf("incorrect factors for 28: %v\n", factors)
    }
  }
}

func TestDivisors(t *testing.T) {
  data := map[int][]int {
    4: {1, 2},
    12: {1, 2, 3, 4, 6},
    28: {1, 2, 4, 7, 14},
    36: {1, 2, 3, 4, 6, 9, 12, 18},
  }
  for n, divs := range data {
    actual := Divisors(n)
    if !reflect.DeepEqual(divs, actual) {
      t.Fatalf("incorrect divisors for %v: got=%v want=%v\n", n, actual, divs)
    }
  }
}
