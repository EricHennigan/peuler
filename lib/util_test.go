package lib

import (
  "testing"
  "reflect"
)

func TestAsDigits(t *testing.T) {
  tests := map[int][]int {
    0: {0},
    1: {1},
    10: {1, 0},
    100: {1, 0, 0},
    999: {9, 9, 9},
  }
  for in, ex := range tests {
    actual := AsDigits(in)
    if !reflect.DeepEqual(ex, actual) {
      t.Fatalf("incorrect conversion to digit list want=%v got=%v", ex, actual)
    }
  }
}

func TestDigitsAsInt(t *testing.T) {
  tests := map[int][]int {
    0: {0},
    1: {1},
    10: {1, 0},
    100: {1, 0, 0},
    999: {9, 9, 9},
    1234: {1, 2, 3, 4},
  }
  for ex, in := range tests {
    actual := DigitsAsInt(in)
    if actual != ex {
      t.Fatalf("did not convert digit list to int, want=%v got=%v", ex, actual)
    }
  }
}
