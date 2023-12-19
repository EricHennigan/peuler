package main
/*
  <title>Lexicographic Permutations</title>
  <p>A permutation is an ordered arrangement of objects. For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4. If all of the permutations are listed numerically or alphabetically, we call it lexicographic order. The lexicographic permutations of 0, 1 and 2 are:</p>
  <p class="center">012   021   102   120   201   210</p>
  <p>What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?</p>
*/

import (
  "peuler/lib"
  "fmt"
  "sort"
)

func main() {
  var strs []string
  //arr := []byte{'0', '1' , '2', '3', '4'}
  arr := []byte{'0', '1' , '2', '3', '4', '5', '6', '7', '8', '9'}
  for p := range lib.Permutations(arr) {
    strs = append(strs, string(p))
  }
  sort.Strings(strs)
  fmt.Println(strs[1_000_000-1])
}