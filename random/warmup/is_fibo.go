package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	// good enough for a cap
	maxFib int64 = 10 << 32 // 10^10 is the max 10 << 32 > 10^10
)

type fibMap map[int64]struct{}

// utopian returns the utopian height of
// the tree after the amount of cycles.
func isFibo(m fibMap, n int64) bool {
	_, ok := m[n]
	return ok
}

// fib computes all fibonacci numbers
// [1, n]. Returns a fibMap containing those
// numbers.
func fib(n int64) fibMap {
	var f, s int64
	f, s = 0, 1

	m := make(fibMap)

	// add the first number

	for s < n {
		// add numbers to map.
		m[s] = struct{}{}

		// compute next numbers
		temp := f + s
		f = s
		s = temp
	}

	return m
}

func read() {
	var N int
	results := make([]int64, 0)

	fmt.Scan(&N)

	// compute fibs
	m := fib(maxFib)

	// fmt.Println(m)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for i := 0; i < N; i++ {
		scanner.Scan()
		x, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		results = append(results, x)
	}

	for _, v := range results {
		if isFibo(m, v) {
			fmt.Println("IsFibo")
		} else {
			fmt.Println("IsNotFibo")
		}
	}
}

func main() {
	read()
}
