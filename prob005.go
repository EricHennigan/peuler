package main

/*
  <title>Smallest Multiple</title>
  <p>$2520$ is the smallest number that can be divided by each of the numbers from $1$ to $10$ without any remainder.</p>
  <p>What is the smallest positive number that is <strong class="tooltip">evenly divisible<span class="tooltiptext">divisible with no remainder</span></strong> by all of the numbers from $1$ to $20$?</p>
*/

import (
	"fmt"
)

func main() {
NUMS:
	for n := 10; true; n++ {
		for i := 2; i < 20; i++ {
			if n%i != 0 {
				continue NUMS
			}
		}
		fmt.Println(n)
		return
	}
}
