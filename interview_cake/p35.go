// In-Place Shuffle
//
// Fisher-Yates Shuffle
package main

import (
	"fmt"
	"math/rand"
)

// Idea: randomly choose an element from
// the array (i, n), n = len(a)-1.
//
// i start at 0 and increments by 1.
// so we choose an element a(i, n) then
// swap that element with a[i], increment i
// and keep going until we're done.
//
// Assuming constant RNG and swapping this
// runs in O(n).
func shuffle(a []int) {
	swaps := rand.Perm(len(a))
	for i := range a {
		j := swaps[i]
		a[i], a[j] = a[j], a[i]
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// shuffle should be different everytime

	shuffle(a)
	fmt.Println(a)

	shuffle(a)
	fmt.Println(a)

	shuffle(a)
	fmt.Println(a)

	shuffle(a)
	fmt.Println(a)

	shuffle(a)
	fmt.Println(a)

	shuffle(a)
	fmt.Println(a)
}
