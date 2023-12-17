package main
/*
  <title>Largest Exponential</title>
  <p>Comparing two numbers written in index form like $2^{11}$ and $3^7$ is not difficult, as any calculator would confirm that $2^{11} = 2048 \lt 3^7 = 2187$.</p>
  <p>However, confirming that $632382^{518061} \gt 519432^{525806}$ would be much more difficult, as both numbers contain over three million digits.</p>
  <p>Using <a href="resources/documents/0099_base_exp.txt">base_exp.txt</a> (right click and 'Save Link/Target As...'), a 22K text file containing one thousand lines with a base/exponent pair on each line, determine which line number has the greatest numerical value.</p>
  <p class="smaller">NOTE: The first two lines in the file represent the numbers in the example given above.</p>
*/

import (
  "bufio"
  "math"
  "fmt"
  "log"
  "os"
  "strings"
  "strconv"
)

func main() {
  file, err := os.Open("./resources/documents/0099_base_exp.txt")
  if err != nil { log.Fatal("could not open file") }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  idx, ln := 1, 1
  nmax := 0.0
  for scanner.Scan() {
    nums := strings.Split(scanner.Text(), ",")
    base, err := strconv.Atoi(nums[0])
    if err != nil { log.Fatalf("could not convert %v", nums[0]) }
    exp, err := strconv.Atoi(nums[1])
    if err != nil { log.Fatalf("could not convert %v", nums[1]) }
    val := float64(exp) * math.Log(float64(base))
    if val > nmax {
      nmax = val
      idx = ln
      fmt.Printf("\t%v: %v, %v\n", ln, base, exp)
    }
    ln += 1
  }
  fmt.Println(idx)
}
