// Find Repeat, Space Edition
// elements are ints [1,n]
// there are n+1 elements
//
// Want O(1) space
package main

import "fmt"

// Use a hashmap to keep track of duplicates.
// We want to optimize for space so this probably is a no go.
func naiveSolution(a []int) []int {
	m := make(map[int]int)

	for _, v := range a {
		if _, ok := m[v]; !ok {
			m[v] = 1
		} else {
			m[v]++
		}
	}

	// check the ones that appear twice or more
	keys := make([]int, 0)
	for k, v := range m {
		if v > 1 {
			keys = append(keys, k)
		}
	}

	return keys
}

// use the array itself as the hashmap
func goodSolution(a []int) []int {
	duplicates := make([]int, 0)

	swaps := 0
	for i := 0; i < len(a); i++ {
		if a[i] != i && a[i] != 0 {
			a[0] = a[i]
			blank := a[0]

			for a[0] != a[blank] { // duplicate found
				a[0], a[blank] = a[blank], a[0]
				blank = a[0]
				swaps++
			}

			// reset 0th element to 0
			a[0] = 0
			duplicates = append(duplicates, blank)
		}
	}
	fmt.Printf("Swaps %d\n", swaps)

	return duplicates
}

var tests = []struct {
	input []int
	want  []int
}{
	{
		input: []int{4, 3, 2, 1, 1, 5, 5, 4},
		want:  []int{1, 5, 4},
	},
	{
		input: []int{1, 1, 2, 2},
		want:  []int{1, 2},
	},
	{
		input: []int{1, 1, 2, 2, 4},
		want:  []int{1, 2},
	},
	{
		input: []int{1, 1},
		want:  []int{1},
	},
	{
		input: []int{},
		want:  []int{},
	},
	{
		input: []int{1, 2, 3, 4, 3, 2, 1},
		want:  []int{3, 2, 1},
	},
	{
		input: []int{5, 5, 4, 4, 3, 3, 2, 2, 1, 1, 7, 6},
		want:  []int{5, 4, 3, 2, 1},
	},
}

func main() {
	for _, t := range tests {
		got := goodSolution(t.input)
		fmt.Printf("input %v: got %v, want %v\n", t.input, got, t.want)
	}
}
