// Given a list of ints
// calculate the product of all
// the numbers except the current index.
//
// We return a list of products.
//
// We're going to the same work multiple
// times we use memoization.
package main

type Key struct {
	missing int // missing index
	len     int
}

// Table represents the lookup table.
var Table map[Key]int

func productWithoutIndex(a []int, index int) int {
	left, right := a[:index], a[index:]
}

func recurseProduct(a []int, index int) int {
	l = len(a)
	if l < 2 {
		Table[Key{i, l}] = 1
		return 1
	}

	k := Key{i, l}
	if v, ok := Table[k]; ok {
		return v
	}

	mid := l / 2

	v = 
}

func main() {

}
