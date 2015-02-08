// Product of All Other Numbers
package main

import "fmt"

type Direction bool

const (
	Left  Direction = false
	Right Direction = true
)

type Key struct {
	index int
	dir   Direction
}

// Table represents the lookup table.
type Table map[Key]int

// genTable, here we generate the lookup table
// for the products.
// Might have to handle an edge case for len(a) < 3
// or something like that but not a dealbreaker.
// O(n) time
//
// Note: slightly better solution use an array
// instead of table. Essentially we can do the same
// thing here in gen table but we right the results
// to the array instead of a table.
//
// Either way it's O(n) time/space.
func genTable(a []int) Table {

	table := make(Table)

	// from left
	table[Key{0, Left}] = 1
	table[Key{1, Left}] = a[0]
	for i := 2; i < len(a); i++ {
		k := Key{i, Left}
		table[k] = table[Key{i - 1, Left}] * a[i-1]
	}

	// from right
	l := len(a) - 1
	table[Key{l, Right}] = 1
	table[Key{l - 1, Right}] = a[l]
	for i := l - 2; i > -1; i-- {
		k := Key{i, Right}
		table[k] = table[Key{i + 1, Right}] * a[i+1]
	}

	return table
}

// Loop through the array looking up the precomputed products.
// O(n) time
func allProducts(a []int) []int {
	table := genTable(a)
	results := make([]int, 0)
	for i := range a {
		k1 := Key{i, Left}
		k2 := Key{i, Right}
		results = append(results, table[k1]*table[k2])
	}
	return results
}

func main() {
	a := []int{1, 7, 3, 4, 0}
	results := allProducts(a)
	fmt.Println(results)
}
