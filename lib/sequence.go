package lib

// Triangle numbers formed by adding the natural numbers
// 7th: 1 + 2 + 3 + 4 + 5 + 6 + 7 = 28
func TriangleNumbers() chan int {
  c := make(chan int)
  i, t := 1, 0
  go func() {
    for {
      t += i
      c <- t
      i++
    }
    close(c)
  }()
  return c
}
      
