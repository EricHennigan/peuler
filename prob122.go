package main
/*
  <title>Efficient Exponentiation</title>
  <p>The most naive way of computing $n^{15}$ requires fourteen multiplications:
  $$n \times n \times \cdots \times n = n^{15}.$$</p>
  <p>But using a "binary" method you can compute it in six multiplications:</p>
  \begin{align}
  n \times n &amp;= n^2\\
  n^2 \times n^2 &amp;= n^4\\
  n^4 \times n^4 &amp;= n^8\\
  n^8 \times n^4 &amp;= n^{12}\\
  n^{12} \times n^2 &amp;= n^{14}\\
  n^{14} \times n &amp;= n^{15}
  \end{align}
  <p>However it is yet possible to compute it in only five multiplications:</p>
  \begin{align}
  n \times n &amp;= n^2\\
  n^2 \times n &amp;= n^3\\
  n^3 \times n^3 &amp;= n^6\\
  n^6 \times n^6 &amp;= n^{12}\\
  n^{12} \times n^3 &amp;= n^{15}
  \end{align}
  <p>We shall define $m(k)$ to be the minimum number of multiplications to compute $n^k$; for example $m(15) = 5$.</p>
  <p>Find $\sum\limits_{k = 1}^{200} m(k)$.</p>
*/

import (
  "fmt"
)

/* Originally thought that I could do something like
   15 = 1 + max(12, 3)
   But that doesn't work, because we generate 3 on the way to 15
   So, better to explicitly create the paths

   Unlocked Knowledge: Knuth's Power Tree
*/

type path []int
type queue []path
const N = 200

func main() {
  var q queue
  q = append(q, path{1})
  nums := make([]path, N+1)
  nums[1] = q[0]

  for len(q) > 0 {
    var p path
    p, q = q[0], q[1:]
    //fmt.Printf("\t\t\tpop %v | q=%v\n", p, q)
    dist := len(p)
    n := p[dist-1]

    // generate new candidate paths, by extending the tail of the current path
    for _, i := range p {
      // don't generate farther than we need
      if n + i > N { continue }
      // don't generate long paths for results already in place
      if len(nums[n + i]) != 0 && len(nums[n + i]) <= len(p) { continue }

      pn := make(path, dist+1)
      copy(pn, p)
      pn[dist] = n + i
      // if n + i == 15 { fmt.Printf("%v len(q)=%v\n", pn, len(q)) }
      // fmt.Printf("enqueue %v\n", pn)
      q = append(q, pn)
      // first time we see a path, it's the shortest
      if len(nums[n + i]) == 0 { nums[n + i] = pn }
    }

    // fmt.Printf("\t\t\tq=%v\n", q)
  }

  s := -(len(nums[0]) - 1) // will cancel out
  for i, p := range nums {
    fmt.Println(i, len(p), p)
    s += len(p) - 1
  }
  fmt.Println(s)
}

/*
[1] -> [1, 2]
[1,2] -> [1, 2, 3] [1, 2, 4]
[1,2,3] -> [1, 2, 3, 4]
           [1, 2, 3, 5]
	   [1, 2, 3, 6]
[1,2,4] -> [1, 2, 4, 5]
           [1, 2, 4, 6]
	   [1, 2, 4, 8]

0 | 0
1 | 0
2 | 1   1
3 | 2   1 2
4 | 3   1 2 4
5 | 4   1 2 3 4
6 | 4   1 2 3 6  or  1 2 4 6
7 | 5   1 2 3 6 7  or 1 2 3 4 7
8 | 4   1 2 4 8
*/


/* Approach that doesn't work

type pair struct {
  a, b, s int
}

func BySum(i, j pair) int {
  if s := cmp.Compare(i.s, j.s); s != 0 {
    return s
  }
  return cmp.Compare(i.a, j.a)
}

func main() {
  N := 200
  var ps []pair

  for i := 1; i <= N; i++ {
    for j := i; j <= N; j++ {
      ps = append(ps, pair{i, j, i+j})
    }
  }
  slices.SortFunc(ps, BySum)

  lens := make([]int, N+1)
  lens[0] = 1
  lens[1] = 0
  for _, p := range ps {
    if p.s > N {continue }
    length := 1 + lens[p.a] + lens[p.b]
    if lens[p.s] == 0 {
      lens[p.s] = length
    }
    lens[p.s] = min(lens[p.s], length)
  }

  // compute the sum
  s := 0
  for i := 1; i <= N; i++ {
    s += lens[i]
  }
  for i := 1; i <= 15; i++ {
    fmt.Printf("[%v] = %v\n", i, lens[i])
  }
  fmt.Println(s)
  fmt.Printf("%v\n", lens)
  fmt.Printf("%v\n", lens[15])
  fmt.Printf("%v", lens[200])
}
*/
