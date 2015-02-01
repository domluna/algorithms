// Apple Stocks
package main

import "fmt"

func maxProfit(a []float32) float32 {
	buy := a[0]
	sell := a[0]

	var profit float32

	for i := 1; i < len(a); i++ {

		// is the price is lower we change
		// when we buy the stock
		if a[i] < buy {
			newProfit := sell - buy
			if newProfit > profit {
				profit = newProfit
			}

			buy = a[i]
			sell = a[i]
		}

		if a[i] > sell {
			sell = a[i]
		}
	}

	newProfit := sell - buy
	if newProfit > profit {
		profit = newProfit
	}
	return profit
}

var tests = []struct {
	input []float32
	want  float32
}{
	{
		[]float32{1.0, 2.0, 3.0, 4.0, 5.0},
		4.0,
	},
	{
		[]float32{10.0, 12.0, 13.0, 1.0, 14.0, 15.0},
		14.0,
	},
	{
		[]float32{10.0, 12.0, 13.0, 1.0, 14.0, 15.0, 10.0, 0.0},
		14.0,
	},
}

func main() {
	for _, t := range tests {
		got := maxProfit(t.input)
		if got != t.want {
			fmt.Printf("ERROR input %v: got %v, want %v\n",
				t.input, got, t.want)
		}
	}
}
