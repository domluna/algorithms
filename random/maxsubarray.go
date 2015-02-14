package main

import "fmt"

type Result struct {
	c, nc int
}

// find the max subarray non-contiguous
func maxSubNC(a []int) int {
	s := 0
	for _, v := range a {
		if v > 0 {
			s += v
		}
	}
	return s
}

// find the max subarray contiguous
// candidates are split by negative numbers
func maxSubC(a []int) int {
	c := 0
	sum := 0

	for _, v := range a {
		val := c + v

		if val > 0 {
			c = val
		} else {
			c = 0
		}

		if c > sum {
			sum = c
		}
	}

	return sum
}

var tests = []struct {
	a      []int
	wantC  int // contiguous
	wantNC int // non-contiguous
}{
	{
		[]int{1, 2, 3, 4},
		10,
		10,
	},
	{
		[]int{2, -1, 2, 3, 4, -5, 15, 2},
		22,
		28,
	},
	{
		[]int{2, -1, 2, 3, 4, -5},
		10,
		11,
	},
	{
		[]int{2, 2, 2, -1, -2, -2, 10},
		11,
		16,
	},
}

func readIn() {
	var T int
	var N int
	var ai int
	var a []int

	results := make([]Result, 0)

	fmt.Scanf("%d", &T)

	for i := 0; i < T; i++ {
		fmt.Scanf("%d", &N)
		a = make([]int, 0)
		for j := 0; j < N; j++ {
			fmt.Scanf("%d", &ai)
			a = append(a, ai)
		}
		results = append(results, Result{maxSubC(a), maxSubNC(a)})
	}

	for _, r := range results {
		fmt.Printf("%d %d\n", r.c, r.nc)
	}
}

func runTests() {

	for i, t := range tests {
		gotC := maxSubC(t.a)
		gotNC := maxSubNC(t.a)
		fmt.Printf("Test %d\n", i)
		if gotC != t.wantC {
			fmt.Printf("\tERROR: got %v, want %v\n", gotC, t.wantC)
		}
		if gotNC != t.wantNC {
			fmt.Printf("\tERROR: got %v, want %v\n", gotNC, t.wantNC)
		}
	}
}

func main() {
	readIn()
}
