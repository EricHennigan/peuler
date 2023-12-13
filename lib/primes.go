package lib

import "math"

// Filter of Eratosthenes for generating prime numbers
func Primes() chan int {
  // TODO: cache primes to a file
  // TODO: prob010 revealed this technique too slow: avoid the divisions by marking a bit array
  var primes []int
  primes = append(primes, 2)

  c := make(chan int)
  go func() {
    c <- 2
    n := 3
    NUMS: for ; ; n += 2 {
      for _, p := range primes {
        if n % p == 0 { continue NUMS }
      }
      primes = append(primes, n)
      c <- n
    }
    close(c)
  }()
  return c
}

// Prime numbers less or equal to N
// The sieve also records a prime factorization
func PrimesUpTo(N int) chan int {
  var nums []int = make([]int, N+2)
  c := make(chan int)
  go func() {
    stop := int(math.Sqrt(float64(N))) + 1
    for i := 2; i <= stop; i++ {
      if nums[i] != 0 { continue }
      c <- i
      for m := i*i; m <= N; m += i {
        if nums[m] == 0 {
	  nums[m] = i
	}
      }
    }
    close(c)
  }()
 return c
}
