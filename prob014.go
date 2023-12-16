package main
/*
  <title>Longest Collatz Sequence</title>
  <p>The following iterative sequence is defined for the set of positive integers:</p>
  <ul style="list-style-type:none;">
  <li>$n \to n/2$ ($n$ is even)</li>
  <li>$n \to 3n + 1$ ($n$ is odd)</li></ul>
  <p>Using the rule above and starting with $13$, we generate the following sequence:
  $$13 \to 40 \to 20 \to 10 \to 5 \to 16 \to 8 \to 4 \to 2 \to 1.$$</p>
  <p>It can be seen that this sequence (starting at $13$ and finishing at $1$) contains $10$ terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at $1$.</p>
  <p>Which starting number, under one million, produces the longest chain?</p>
  <p class="note"><b>NOTE:</b> Once the chain starts the terms are allowed to go above one million.</p>
*/

import (
  "container/list"
  "fmt"
)

const N = 1_000_000

func collatz(n int) chan int {
  c := make(chan int)
  go func() {
    for n != 1 {
      if n % 2 == 0 {
        n /= 2
      } else {
        n *= 3
	n += 1
      }
      c <- n
    }
    close(c)
  }()
  return c
}

func main() {
  dist := make(map[int]int)
  dist[1] = 1

  for n := 1; n <= N; n++ {
    //fmt.Println("--")
    //fmt.Println(n)
    seq := list.New()
    d := 0
    for c := range collatz(n) {
      seq.PushFront(c)
      if v, ok := dist[c]; ok { 
        d = v
	break
      }
    }
    for e := seq.Front(); e != nil; e = e.Next() {
      v := e.Value.(int)
      dist[v] = d
      d++
      //fmt.Printf("%v -> ", e.Value.(int))
    }
    dist[n] = d
  }
  m := 0
  for n := 1; n <= N; n++ {
    if dist[n] > m {
      m = dist[n]
      fmt.Printf("%v has length %v\n", n, m)
    }
  }
}


/*
// Best approach is to run in reverse until all first 1M numbers are filled in
// Nope, apparently not, the queue gets far too big

func reverse_collatz(n int) (int, int) {
  return 2 * n, (n - 1) / 3
}

func main() {
  nums := make(map[int]int)
  remaining := N
  q := list.New();
  q.PushFront(1)
  for e := q.Front(); remaining > 0 ; e = e.Next() {
    n := e.Value.(int)
    c1, c2 := reverse_collatz(n)
    if _, ok := nums[c1]; !ok {
      nums[c1] = nums[n] + 1
      if 1 <= c1 && c1 <= N { remaining-- }
      q.PushBack(c1)
    }
    if _, ok := nums[c2]; !ok {
      nums[c2] = nums[n] + 1
      if 1 <= c2 && c2 <= N { remaining-- }
      q.PushBack(c2)
    }
  }

  m := 0
  for i := 1; i <= N; i++ {
    if m < nums[i] {
      m = nums[i]
      fmt.Println(i, m)
    }
  }
}
*/

/*
type pair struct {
  next, dist int
}

func collatz(n int) int {
  if n % 2 == 0 { return n / 2 }
  return 3 * n + 1
}

func main() {
  nums := make(map[int]pair)
  nums[1] = pair{0, 1}
  var q []int
  for n := 2; n < N; n++ {
    q = append(q, n)
  }

  fmt.Printf("%v\n", q)
  for len(q) != 0 {
    var n int
    n, q = q[0], q[1:]
    fmt.Printf("%v\n", len(q))
    c := collatz(n)
    _, ok := nums[c]
    if ok {
      nums[n] = pair{c, nums[c].dist + 1}
      //fmt.Printf("update %v\n", nums)
    } else {
      q = append(q, c)
      q = append(q, n)
    }
  }

  fmt.Println(nums)
}
*/
