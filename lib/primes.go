package lib

// Seive of Eratosthenes for generating prime numbers
func Primes() chan int {
  // TODO: cache primes to a file
  // TODO?: avoid the divisions by marking a bit array
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

