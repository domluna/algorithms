// Single Rifle
package main

import "fmt"

// A deck is a 52 element array, 1..52

// type Riffler struct {
// 	deck []int
// }

// shuffleDeck returns true if the deck
// is a single riffle from half1 and half2.
func Riffled(deck, half1, half2 []int) bool {

	// loop through the deck
	// at each card the current card
	// matches a card at the front
	// of either halves.
	//
	// move the pointer of the front to the next
	// card if it matches

	i, j := 0, 0
	ll, lr := len(half1), len(half2)

	for _, card := range deck {

		if i < ll && card == half1[i] {
			i++
		} else if j < lr && card == half2[j] {
			j++
		} else {
			return false
		}
	}

	return true
}

func main() {
	var r bool
	r = Riffled([]int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3}, []int{4, 5, 6})
	fmt.Println(r) // true
	r = Riffled([]int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3}, []int{6, 5, 4})
	fmt.Println(r) // false
	r = Riffled([]int{1, 2, 3, 4, 5, 6}, []int{1, 3, 5}, []int{2, 4, 6})
	fmt.Println(r) // true
	r = Riffled([]int{1, 2, 3, 4, 5, 6, 5}, []int{1, 3, 6}, []int{2, 4, 5, 5})
	fmt.Println(r) // true
}
