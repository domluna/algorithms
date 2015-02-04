package main

import "fmt"

// Given two 32bit numbers x, y
// Write y as a substring of x
// between bit positions i, j
// i < j
func p1(x, y *uint32, i, j uint32) {
	yy := *y
	*y = *y << i
	*x = *x ^ *y
}

func main() {
	x := uint32(1 << 7)
	y := uint32(20)
	fmt.Println(x, y)
	p1(&x, &y, 0, 3)
	fmt.Println(x, y)
}
