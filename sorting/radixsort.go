// Radix sort
package main

import "fmt"

func radixSort(a []int) {

	maxVal := 0
	for _, v := range a {
		if v > maxVal {
			maxVal = v
		}
	}

	i := 0
	// assume decimal
	const base = 10

	// init the buckets
	var buckets [base][]int
	for i := range buckets {
		buckets[i] = []int{}
	}
	for pow(base, i) < maxVal {

		// Fill the buckets
		for _, v := range a {

			digit := (v / pow(base, i)) % base
			buckets[digit] = append(buckets[digit], v)
		}

		// Empty buckets, write back to array.
		// An improvement here would be to use queues.
		j := 0
		for b := range buckets {
			for _, v := range buckets[b] {
				a[j] = v
				j++
			}
			buckets[b] = []int{}
		}
		i++
	}
}

func pow(base, power int) int {
	result := 1
	for i := 0; i < power; i++ {
		result *= base
	}
	return result
}

func main() {
	a := []int{10, 5, 3, 2, 55, 22, 6, 10}
	radixSort(a)
	fmt.Println(a)
}
