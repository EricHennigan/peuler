package answers

import (
  "testing"
  "bufio"
  "os"
  "os/exec"
  "strings"
)

func TestAnswers(t *testing.T) {
  file, err := os.Open("answers.txt")
  if err != nil { t.Fatal(err) }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    data := strings.Split(scanner.Text(), ": ")
    if len(data) != 2 { continue }
    data[0] = "prob" + data[0] + ".go"
    cmd := exec.Command("go", "run", data[0])
    cmd.Dir = "../"
    stdout, err := cmd.Output()
    if err != nil { t.Fatalf("testing %v: %v", data, err) }
    lines := strings.Split(string(stdout), "\n")
    answ := lines[len(lines)-2]
    if answ != data[1] {
      t.Fatalf("testing %v got %v wanted %v\n", data[0], answ, data[1])
    }
  }
}
