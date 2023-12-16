package main
/*
  <title>Names Scores</title>
  <p>Using <a href="resources/documents/0022_names.txt">names.txt</a> (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.</p>
  <p>For example, when the list is sorted into alphabetical order, COLIN, which is worth $3 + 15 + 12 + 9 + 14 = 53$, is the $938$th name in the list. So, COLIN would obtain a score of $938 \times 53 = 49714$.</p>
  <p>What is the total of all the name scores in the file?</p>
*/

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strings"
  "sort"
)

func main() {
  file, err := os.Open("./resources/documents/0022_names.txt")
  if err != nil { log.Fatal("could not open file") }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  scanner.Scan()
  names := strings.Split(scanner.Text(), ",")
  sort.Strings(names)
  s := 0
  for i, n := range names {
    name := n[1:len(n)-1]
    value := 0
    for _, c := range name {
      value += int(c - 'A' + 1)
    }
    score := (i + 1) * value
    s += score
  }
  fmt.Println(s)
}
