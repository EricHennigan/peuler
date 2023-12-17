package main
/*
  <title>Non-Abundant Sums</title>
  <p>A perfect number is a number for which the sum of its proper divisors is exactly equal to the number. For example, the sum of the proper divisors of $28$ would be $1 + 2 + 4 + 7 + 14 = 28$, which means that $28$ is a perfect number.</p>
  <p>A number $n$ is called deficient if the sum of its proper divisors is less than $n$ and it is called abundant if this sum exceeds $n$.</p>

  <p>As $12$ is the smallest abundant number, $1 + 2 + 3 + 4 + 6 = 16$, the smallest number that can be written as the sum of two abundant numbers is $24$. By mathematical analysis, it can be shown that all integers greater than $28123$ can be written as the sum of two abundant numbers. However, this upper limit cannot be reduced any further by analysis even though it is known that the greatest number that cannot be expressed as the sum of two abundant numbers is less than this limit.</p>
  <p>Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.</p>
*/

import (
  "fmt"
  "peuler/lib"
)

func isAbundant(n int) bool {
  return n < lib.Sum(lib.Divisors(n))
}

//const N = 50
const N = 28123

func main() {
  var abundant []int
  for n := 1; n <= N; n++ {
    if isAbundant(n) { abundant = append(abundant, n) }
  }
  // fmt.Println(abundant)

  nums := make(map[int]bool)
  for i := 0; i < len(abundant); i++ {
    for j := i; j < len(abundant); j++ {
      nums[abundant[i] + abundant[j]] = true
    }
  }
  //fmt.Println(nums)

  sum := 0
  for n := 1; n < N; n++ {
    if _, ok := nums[n]; !ok {
      sum += n
    }
  }
  fmt.Println(sum)
}
