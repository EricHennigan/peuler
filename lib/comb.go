package lib

func Rotations[T any](arr []T) chan []T {
  c := make(chan []T)
  s := len(arr)
  generate := func() {
    for i, _ := range arr {
      r := make([]T, s)
      copy(r[0:s-i], arr[i:])
      copy(r[s-i:], arr[:i])
      c <- r
    }
    close(c)
  }
  go generate()
  return c
}

// Efficient swapping implementation of permutations via Heap's Algorithm
func Permutations[T any](arr []T) chan []T {
  c := make(chan []T)

  var generate func(k int)
  generate = func(k int) {
    if k == 1 {
      cp := make([]T, len(arr))
      copy(cp, arr)
      c <- cp
      return
    }
    generate(k-1)
    for i := 0; i < k-1; i++ {
      if k & 1 == 0 {
        arr[i], arr[k-1] = arr[k-1], arr[i]
      } else {
        arr[0], arr[k-1] = arr[k-1], arr[0]
      }
      generate(k-1)
    }
  }
  go func() {
    generate(len(arr))
    close(c)
  }()
  return c
}
