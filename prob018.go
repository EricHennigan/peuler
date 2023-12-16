package main
/*
  <title>Maximum Path Sum I</title>
  <p>By starting at the top of the triangle below and moving to adjacent numbers on the row below, the maximum total from top to bottom is $23$.</p>
  <p class="monospace center"><span class="red"><b>3</b></span><br><span class="red"><b>7</b></span> 4<br>
  2 <span class="red"><b>4</b></span> 6<br>
  8 5 <span class="red"><b>9</b></span> 3</p>
  <p>That is, $3 + 7 + 4 + 9 = 23$.</p>
  <p>Find the maximum total from top to bottom of the triangle below:</p>
  <p class="monospace center">75<br>
  95 64<br>
  17 47 82<br>
  18 35 87 10<br>
  20 04 82 47 65<br>
  19 01 23 75 03 34<br>
  88 02 77 73 07 63 67<br>
  99 65 04 28 06 16 70 92<br>
  41 41 26 56 83 40 80 70 33<br>
  41 48 72 33 47 32 37 16 94 29<br>
  53 71 44 65 25 43 91 52 97 51 14<br>
  70 11 33 28 77 73 17 78 39 68 17 57<br>
  91 71 52 38 17 14 91 43 58 50 27 29 48<br>
  63 66 04 68 89 53 67 30 73 16 69 87 40 31<br>
  04 62 98 27 23 09 70 98 73 93 38 53 60 04 23</p>
  <p class="note"><b>NOTE:</b> As there are only $16384$ routes, it is possible to solve this problem by trying every route. However, <a href="problem=67">Problem 67</a>, is the same challenge with a triangle containing one-hundred rows; it cannot be solved by brute force, and requires a clever method! ;o)</p>
*/

import (
  "fmt"
  "bufio"
  "strings"
  // "container/list"
  "strconv"
)

const input = `
75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23
`

/*
const input = `
3
7 4
2 4 6
8 5 9 3
`
/*
    3                    23>22
   7 4                17<20  17<19
  2 4 6            10>7, 9<13, 15<9
 8 5 9 3
*/

func main() {
  var grid [15][16]int
  var path [15][16]int

  scanner := bufio.NewScanner(strings.NewReader(input))
  scanner.Scan() // empty first line
  for i := 0; i < 15; i++ {
    scanner.Scan()
    line := strings.Split(scanner.Text(), " ")
    for j := 0; j < len(line); j++ {
      v, _ := strconv.Atoi(line[j])
      grid[i][j] = v
    }
  }

  path[14] = grid[14]

  for i := 13; i >= 0; i-- {
    for j := 0; j <= i; j++ {
      path[i][j] += grid[i][j] + max(path[i+1][j], path[i+1][j+1])
    }
    fmt.Println(grid[i])
    fmt.Println()
    fmt.Println(path[i])
  }
  fmt.Println(path[0][0])
}



// Don't bother making a tree

/*
type tnode struct {
  parent *tnode
  left, right *tnode
  val int
}

func main() {
  scanner := bufio.NewScanner(strings.NewReader(input))
  in := list.New()

  scanner.Scan() // empty first line
  for scanner.Scan() {
    nums := strings.Split(scanner.Text(), " ")
    for _, n := range nums {
      v, _ := strconv.Atoi(n)
      t := &tnode{val: v}
      in.PushBack(t)
    }
  }

  q := list.New()
  q.PushFront(in.Front().Value)
  root := q.Front().Value.(*tnode)
  in.Remove(in.Front())
  fmt.Printf("\tlen=%v\n", in.Len())

  for t := q.Front(); in.Len() > 0; t = t.Next() {
    tn := t.Value.(*tnode)
    fmt.Printf("process %v\n", tn)
    tn.left = in.Front().Value.(*tnode)
    tn.left.parent = tn
    q.PushBack(tn.left)
    in.Remove(in.Front())
    tn.right = in.Front().Value.(*tnode)
    tn.right.parent = tn
    q.PushBack(tn.right)
    in.Remove(in.Front())
    fmt.Printf("process %v\n", tn)
    fmt.Printf("\tlen=%v\n", in.Len())
  }
    
  for e := in.Front(); e != nil; e = e.Next() {
    fmt.Println(e.Value)
  }
  fmt.Println(root)
}
*/
