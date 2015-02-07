// Merge sorted arrays
package main

import (
	"fmt"
	"reflect"
)

func merge(left, right []int) []int {
	ll, lr := len(left), len(right)
	i, j := 0, 0

	merged := make([]int, 0)

	// loop until we exhause one of the arrays
	for i < ll && j < lr {
		if left[i] < right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}

	// add the remaining elements
	// of whichever we didn't exhaust
	for i < ll {
		merged = append(merged, left[i])
		i++
	}

	for j < lr {
		merged = append(merged, right[j])
		j++
	}

	return merged
}

var tests = []struct {
	left, right []int
	want        []int
}{
	{
		left:  []int{3, 4, 6, 10, 11, 15},
		right: []int{1, 5, 8, 12, 14, 19},
		want:  []int{1, 3, 4, 5, 6, 8, 10, 11, 12, 14, 15, 19},
	},
}

func main() {
	for _, t := range tests {
		got := merge(t.left, t.right)
		if !reflect.DeepEqual(got, t.want) {
			fmt.Printf("got %v, want %v\n", got, t.want)
		}
	}

}
