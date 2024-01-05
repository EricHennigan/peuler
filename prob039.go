package main
/*
  <title>Integer Right Triangles</title>
  <p>If $p$ is the perimeter of a right angle triangle with integral length sides, $\{a, b, c\}$, there are exactly three solutions for $p = 120$.</p>
  <p>$\{20,48,52\}$, $\{24,45,51\}$, $\{30,40,50\}$</p>
  <p>For which value of $p \le 1000$, is the number of solutions maximised?</p>
*/

import (
  "fmt"
)

func main() {
  triangles := make(map[int][][]int)
  for a := 1; a <= 1000; a++ {
    for b := 1; b <= 1000; b++ {
      for c := 1; c <= 1000; c++ {
	p := a + b + c
	if p > 1000 { continue }
	if c*c != a*a + b*b { continue }
	triangles[p] = append(triangles[p], []int{a, b, c})
      }
    }
  }
  perim, count := 0, 0
  for p, c := range triangles {
    if len(c) > count {
      perim, count = p, len(c)
    }
  }
  fmt.Println(perim)
}
