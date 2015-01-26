// QuickSort
package main

import "fmt"

// quickSort runs the Quicksort algorithm.
func quickSort(a []int) {
	if len(a) < 2 {
		return
	}

	pivValue := a[0]
	storeIdx := 0
	pivIndex := 0
	l := len(a) - 1

	a[pivIndex], a[l] = a[l], a[pivIndex]

	for i, v := range a {
		if v < pivValue {
			a[storeIdx], a[i] = a[i], a[storeIdx]
			storeIdx += 1
		}

	}

	a[storeIdx], a[l] = a[l], a[storeIdx]
	quickSort(a[:storeIdx])
	quickSort(a[storeIdx+1:])
}

func main() {
	a := []int{10, 5, 3, 2, 55, 22, 6, 10}
	quickSort(a)
	fmt.Println(a)
}
