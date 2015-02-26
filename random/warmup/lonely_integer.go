package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// lonely finds the lonely (unique) int in a.
// all but one element in a appear in pairs.
func lonely(a []int) int {
	// use xor trick over hash table
	x := 0
	for _, v := range a {
		x = x ^ v
	}
	return x
}

// read reads stdin, solves lonely integer.
func read() {
	var N int
	a := make([]int, 0)

	fmt.Scanf("%d", &N)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for i := 0; i < N; i++ {
		scanner.Scan()
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		a = append(a, x)
	}

	fmt.Printf("%d\n", lonely(a))
}

func main() {
	// l := lonely([]int{1, 1, 2, 3, 4, 3, 4}) // 2
	// fmt.Println(l)

	read()
}
