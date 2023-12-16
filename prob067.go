package main
/*
  <title>Maximum Path Sum II</title>
  <p>By starting at the top of the triangle below and moving to adjacent numbers on the row below, the maximum total from top to bottom is 23.</p>
  <p class="monospace center"><span class="red"><b>3</b></span><br><span class="red"><b>7</b></span> 4<br>
  2 <span class="red"><b>4</b></span> 6<br>
  8 5 <span class="red"><b>9</b></span> 3</p>
  <p>That is, 3 + 7 + 4 + 9 = 23.</p>
  <p>Find the maximum total from top to bottom in <a href="resources/documents/0067_triangle.txt">triangle.txt</a> (right click and 'Save Link/Target As...'), a 15K text file containing a triangle with one-hundred rows.</p>
  <p class="smaller"><b>NOTE:</b> This is a much more difficult version of <a href="problem=18">Problem 18</a>. It is not possible to try every route to solve this problem, as there are 2<sup>99</sup> altogether! If you could check one trillion (10<sup>12</sup>) routes every second it would take over twenty billion years to check them all. There is an efficient algorithm to solve it. ;o)</p>
*/

import (
  "os"
  "log"
  "fmt"
  "bufio"
  "strings"
  "strconv"
)

func main() {
  var grid [100][100]int
  var path [100][100]int

  file, err := os.Open("./resources/documents/0067_triangle.txt")
  if err != nil { log.Fatal("could not open file") }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for i := 0; i < 100; i++ {
    scanner.Scan()
    line := strings.Split(scanner.Text(), " ")
    for j := 0; j < len(line); j++ {
      v, _ := strconv.Atoi(line[j])
      grid[i][j] = v
    }
  }
  path[99] = grid[99]

  for i := 98; i >= 0; i-- {
    for j := 0; j <= i; j++ {
      path[i][j] += grid[i][j] + max(path[i+1][j], path[i+1][j+1])
    }
  }
  fmt.Println(path[0][0])
}