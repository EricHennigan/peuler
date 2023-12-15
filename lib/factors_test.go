package lib

import (
  "testing"
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
