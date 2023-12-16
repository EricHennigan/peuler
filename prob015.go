package main
/*
  <title>Lattice Paths</title>
  <p>Starting in the top left corner of a $2 \times 2$ grid, and only being able to move to the right and down, there are exactly $6$ routes to the bottom right corner.</p>
  <div class="center">
  <img src="resources/images/0015.png?1678992052" class="dark_img" alt=""></div>
  <p>How many such routes are there through a $20 \times 20$ grid?</p>
*/

import "fmt"

const N = 21
const M = 21

func main() {
  var grid [N][M]int
  for i := 0; i < N; i++ {
    grid[i][0] = 1
  }
  for j := 0; j < M; j++ {
    grid[0][j] = 1
  }
  for i := 1; i < N; i++ {
    for j := 1; j < M; j++ {
      grid[i][j] = grid[i-1][j] + grid[i][j-1]
    }
  }
  fmt.Println(grid[2][2])
  fmt.Println(grid[N-1][M-1])
}
