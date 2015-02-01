// The Stolen Breakfast Drone
package main

import "fmt"

// x ^ x = 0
// The idea is that all duplicates will
// cancel eachother out.
//
// This leaves only the unique int.
func findUnique(a []int) int {
	id := 0
	for _, v := range a {
		id = v ^ id
	}
	return id
}

var tests = []struct {
	in   []int
	want int
}{
	{
		[]int{1, 1, 3, 2, 2, 10, 10},
		3,
	},
	{
		[]int{1},
		1,
	},
	{
		[]int{2, 1, 2},
		1,
	},
}

func main() {
	for _, t := range tests {
		got := findUnique(t.in)
		if got != t.want {
			fmt.Printf("ERROR input %v: got %v, want %v",
				t.in, got, t.want)
		}
	}
}
