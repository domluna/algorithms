// Making Change
//
// Given an amount of money and
// a list of coin denominations,
// figure out all ways to make the money.
//
package main

import "fmt"

var cache = make(map[int][][]int)

func coin(amount int, denoms []int) [][]int {

	ways := make([][]int, 0)

	if amount == 0 {
		ways = append(ways, []int{})
		return ways
	}

	if res, ok := cache[amount]; ok {
		return res
	}

	for _, d := range denoms {
		diff := amount - d
		if diff < 0 {
			continue
		}

		res := coin(diff, denoms)
		for _, r := range res {
			// add d to the results
			ways = append(ways, append(r, d))
		}

	}

	cache[amount] = ways
	return ways
}

func coinBotUp(amount int, denoms []int) [][]int {

	// Figure out how many ways there are to 1..amount
	// coins with the denoms
	c := make(map[int][][]int)

	for i := 1; i <= amount; i++ {
		c[i] = make([][]int, 0)

		for _, d := range denoms {
			prev := c[i-d]
			for _, p := range prev {
				c[i] = append(c[i], append(p, d))
			}

			if i == d {
				c[i] = append(c[i], []int{d})
			}
		}
	}

	return c[amount]
}

var tests = []struct {
	amount int
	denoms []int
	want   [][]int
}{
	{
		amount: 4,
		denoms: []int{1, 2, 3},
		want: [][]int{
			[]int{1, 1, 1, 1},
			[]int{1, 1, 2},
			[]int{1, 3},
			[]int{2, 2},
		},
	},
}

func main() {
	ways := coinBotUp(4, []int{3, 2, 1})
	fmt.Println(ways)
}
