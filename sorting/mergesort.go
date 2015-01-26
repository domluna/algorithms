// Mergesort.
package main

import "fmt"

// mergeSort runs the Mergesort algorithm.
func mergeSort(a []int) {
	if len(a) < 2 {
		return
	}
	mid := len(a) / 2

	left := a[:mid]
	right := a[mid:]

	mergeSort(left)
	mergeSort(right)

	result := merge(left, right)
	for i, v := range result {
		a[i] = v
	}
}

// merge merges two sorted lists into one sorted list.
func merge(left []int, right []int) []int {
	var result []int
	ll := len(left)
	lr := len(right)
	i, j := 0, 0
	for i < ll && j < lr {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < ll {
		result = append(result, left[i])
		i++
	}

	for j < ll {
		result = append(result, right[j])
		j++
	}

	return result
}

func main() {
	a := []int{10, 5, 3, 2, 55, 22, 6, 10, -1}
	mergeSort(a)
	fmt.Println(a)
}
