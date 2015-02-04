package main

import "fmt"

// int is 32 bits
func count(x int) int {
	c := 0
	for i := 0; i < 32; i++ {
		r := x & 1
		if r == 1 {
			c++
		}
		x = x >> 1
	}
	return c
}

// diff returns the difference in bits between two ints.
func diff(x, y int) int {
	z := x ^ y
	return count(z)
}

func main() {
	fmt.Println("COUNT")
	fmt.Println(count(1 << 7))
	fmt.Println(count(31))
	fmt.Println(count(14))
	fmt.Println(count(1))
	fmt.Println(count(1 << 31)) // max

	fmt.Println("DIFF")
	fmt.Println(diff(31, 14))
}
