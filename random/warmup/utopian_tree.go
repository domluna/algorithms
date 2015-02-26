package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// utopian returns the utopian height of
// the tree after the amount of cycles.
func utopian(cycles int) int {
	h := 1
	i := 0
	for {
		if i >= cycles {
			break
		}

		// part 1: double height
		h *= 2
		i++

		if i >= cycles {
			break
		}
		// part 2: add 1 meter to height
		h += 1
		i++
	}
	return h
}

func read() {
	var N int
	results := make([]int, 0)

	fmt.Scan(&N)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for i := 0; i < N; i++ {
		scanner.Scan()
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		results = append(results, x)
	}

	for _, v := range results {
		fmt.Printf("%v\n", utopian(v))
	}
}

func main() {
	read()
}
