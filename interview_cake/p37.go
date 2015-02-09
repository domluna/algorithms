// 7-Sided Dice
package main

import (
	"fmt"
	"math/rand"
)

// rand7 generates a random number from 1..7.
func rand7() int {
	return rand.Intn(7) + 1
}

// rand5 generates a random number from 1..5.
// Write rand5 in terms of rand7
func rand5() int {
	n := rand7()
	for n > 5 {
		n = rand7()
	}
	return n
}

func main() {
	buckets := make([]int, 5)
	n := 10000
	for i := 0; i < n; i++ {
		r := rand5()
		buckets[r-1]++
	}

	fmt.Printf("%d numbers generated\n", n)
	for i := range buckets {
		percentage := float32(buckets[i]) / float32(n)
		fmt.Printf("Bucket %d: %f\n", i+1, percentage)
	}
}
