package lib

import (
  "container/list"
  "testing"
  "time"
)

func TestPrimesSpeed(t *testing.T) {
  // Use the division algorithm
  var slow = func(N int) (time.Duration, *list.List) {
    start := time.Now()
    l := list.New()
    e := l.PushBack(1)
    for p := range Primes() {
      if p >= N { break }
      e = l.InsertAfter(p, e)
    }
    elapsed := time.Now().Sub(start)
    return elapsed, l
  }

  // Use the sieve
  var fast = func(N int) (time.Duration, *list.List) {
    start := time.Now()
    l := list.New()
    e := l.PushBack(1)
    for p := range PrimesUpTo(N) {
      e = l.InsertAfter(p, e)
    }
    elapsed := time.Now().Sub(start)
    return elapsed, l
  }

  for N := 10; N < 1_000_000_000; N *= 10 {
    fast_time, _ := fast(N)
    t.Logf("Primes up to % 7v fast=%v\n", N, fast_time)
  }

  for N := 10; N < 1000000; N *= 10 {
    slow_time, slow_list := slow(N)
    fast_time, fast_list := fast(N)

    // Lists must be the same
    e_slow := slow_list.Front()
    e_fast := fast_list.Front()
    for e_slow != nil && e_fast != nil {
      if e_slow.Value != e_fast.Value {
        t.Fatalf("primes slow=%v fast=%v are not the same", e_slow.Value, e_fast.Value)
      }
      e_slow = e_slow.Next()
      e_fast = e_fast.Next()
    }
    t.Logf("Primes up to % 7v fast=%v\tslow=%v\n", N, fast_time, slow_time)
  }
}
