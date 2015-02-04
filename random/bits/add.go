package main

import "fmt"

// add in binary.
func add(x, y int) int {
	carry := 0
	z := 0
	r := 0
	for i := uint(0); i < 32; i++ {
		b1 := x & 1
		b2 := y & 1
		r = b1 ^ b2 ^ carry
		carry = (b1 & b2) | ((b1 | b2) & carry)

		// change z
		z = z | (r << i)

		// move down
		x = x >> 1
		y = y >> 1

	}
	return z
}

var tests = []struct {
	x, y, z int
}{
	{
		x: 10,
		y: 4,
		z: 14,
	},
	{
		x: 1,
		y: 1,
		z: 2,
	},
	{
		x: 10,
		y: 10,
		z: 20,
	},
	{
		x: 10,
		y: 0,
		z: 10,
	},
}

func main() {
	for _, t := range tests {
		z := add(t.x, t.y)
		if z != t.z {
			fmt.Printf("ERROR: want %d, got %d\n", t.z, z)
		}
	}
}
