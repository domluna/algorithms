package main

import "fmt"

// columns and rows start at 1 not 0
// 1 <= r <= 2 * 10^9
// 1 <= c <= 5
//
// Grid looks like this:
// ..............
//
// ..............
//
// 20 22 24 26 28
//
// 11 13 15 17 19
//
// 10 12 14 16 18
//
//  1  3  5  7  9
//
//  0  2  4  6  8

// stangeGrid returns the numbers corresponding to row
// and col in the strange grid pictured above.
func strangeGrid(row, col int) int {
	// every 2 rows the numbers jump by ten
	var v int

	if row%2 == 0 {
		v = (10 * ((row / 2) - 1))
		v += (col*2 - 1)
	} else {
		v = (10 * (row / 2))
		v += (col*2 - 2)
	}

	return v

}

func read() {
	var c, r int
	fmt.Scan(&r, &c)

	fmt.Println(strangeGrid(r, c))
}

func main() {
	read()
}
