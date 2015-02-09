// 5-Sided Dice
package main

import (
	"fmt"
	"math/rand"
)

// rand7 generates a random number from 1..7.
// Write rand7 in terms of rand5
//
// Trick here use the combination of rand5 to generate
// uniformly numbers from 1..25
//
// Taking the mod of 7 to any number here will give
// a number 0-6. This only works if the highest outcome
// is a multiple of 7. Here we choose 21 but 14 or 7 work
// as well. In that case we will loop more times though.
//
// The solution I came up with is a bit different.
// We use rand5 to simulate rand2 we then we numbers
// 1..10 uniformly we can do this via 2 initial calls to rand5,
// one of them being rand5() + 5.
//
// We then view range 1..5 as bucket1 and range 6..10 as bucket2
// We use rand2 to decide which bucket to draw the number from.
// The probably of selecting a number from bucket1 is 1/2 * 5/5 (if
// we enter bucket1 any number in the bucket is in the valid range).
//
// the probability of selecting an item from bucket2 is 1/2 * 2/5
// or 2/10. If we pick 8..10 we repeat the process.
//
// So we have prob 1/2 of selecting an item from bucket1 and
// 2/10 prob from bucket2. At first this looks like it doesn't
// get us anywhere.
//
// We'll that's our new probability space so we add the 2 numbers together
// and we get 7/10. It's not 1. So we normalize.
//
// Key: Divide of probs by 7/10, so 5/10 / 7/10 and 2/10 / 7/10. This
// results in 5/7 and 2/7!
//
// So we proved that we have a 5/7 chance to draw from bucket1 and a 2/7
// chance to draw from bucket2.
//
// We have a 1/5 to pick any element in bucket1 and a 1/2 chance to pick any elem// in bucket2 so the chance of any number 1..7 being picked is 1/7.
func rand7() int {
	r1 := rand5()
	r2 := rand5()
	outcome := (r1-1)*5 + r2

	for outcome > 21 {
		r1 = rand5()
		r2 = rand5()
		outcome = (r1-1)*5 + r2
	}

	return (outcome % 7) + 1

}

// rand5 generates a random number from 1..5.
func rand5() int {
	return rand.Intn(5) + 1
}

func main() {
	buckets := make([]int, 7)
	n := 1000000
	for i := 0; i < n; i++ {
		r := rand7()
		buckets[r-1]++
	}

	fmt.Printf("%d numbers generated\n", n)
	for i := range buckets {
		percentage := float32(buckets[i]) / float32(n)
		fmt.Printf("Bucket %d: %f\n", i+1, percentage)
	}
}
