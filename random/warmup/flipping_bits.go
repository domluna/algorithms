package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// flip flips the bits of the passed uint32
func flip(x uint32) uint32 {
	return x ^ uint32((1<<32)-1)
}

func read() {
	var N int
	results := make([]uint32, 0)

	fmt.Scan(&N)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for i := 0; i < N; i++ {
		scanner.Scan()
		x, err := strconv.ParseUint(scanner.Text(), 10, 32)
		if err != nil {
			panic(err)
		}
		results = append(results, uint32(x))
	}

	for _, v := range results {
		fmt.Printf("%v\n", flip(v))
	}
}

func main() {
	read()
}
