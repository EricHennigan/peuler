package lib

import (
  "testing"
  "reflect"
)

func TestRotations(t *testing.T) {
  input := []int{1, 2, 3}
  expected := [][]int{
    {1, 2, 3},
    {2, 3, 1},
    {3, 1, 2},
  }
  var actual [][]int
  for r := range Rotations(input) {
    actual = append(actual, r)
  }
  if !reflect.DeepEqual(expected, actual) {
    t.Fatalf("incorrect rotations %v", actual)
  }
}

func TestPermutations(t *testing.T) {
  input := []int{1, 2, 3}
  expected := [][]int{
    {1, 2, 3},
    {2, 1, 3},
    {3, 1, 2},
    {1, 3, 2},
    {2, 3, 1},
    {3, 2, 1},
  }
  var actual [][]int
  for p := range Permutations(input) {
    actual = append(actual, p)
  }
  if !reflect.DeepEqual(expected, actual) {
    t.Fatalf("incorrect permutations %v", actual)
  }
}
