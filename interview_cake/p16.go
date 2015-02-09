// Cake Thief
package main

import "fmt"

type cakeThief struct {
	// capacity the thief can steal
	cap int
}

type cake struct {
	weight int
	value  int
}

// steal steals the max value of cake w.r.t to
// cakeThief's capacity.
func (ct *cakeThief) steal(cakes []cake) int {
	memo := make([]int, ct.cap+1)

	for w := 1; w <= ct.cap; w++ {
		max := memo[w-1]
		// memo[i] = max(memo[i-1],
		// max(v_i * x_i + m[w - x_i * w_i]))
		// x_i * w_i <= w

		for _, c := range cakes {
			xi := 1

			// capacity we are adding must be less
			// than our current

			// for all intensive purposes this is
			// silly, nothing practical will have
			// weight 0.
			//
			// Since we can an unlimited number of each item
			// a 0 weight item would give us inf profit.
			//
			// So we do the sane thing here and just ignore
			// the cake where weight == 0.
			//
			// We could return inf but I'm not sure how that
			// works that works when converting it to an int.
			if c.weight == 0 {
				continue
			}

			for xi*c.weight <= w {
				val := c.value*xi + memo[w-xi*c.weight]
				if val > max {
					max = val
				}
				xi++
			}
		}
		memo[w] = max

	}

	fmt.Println(memo)

	return memo[ct.cap]
}

func main() {
	ct := cakeThief{}
	ct.cap = 20
	profit := ct.steal([]cake{
		{7, 160},
		{3, 90},
		{2, 15},
		{0, 10},
	})

	fmt.Println(profit) // 555
}
