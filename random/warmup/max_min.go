package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Int64Slice []int64

func (is Int64Slice) Len() int {
	return len(is)
}

func (is Int64Slice) Less(i, j int) bool {
	return is[i] < is[j]
}

func (is Int64Slice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

var N int64
var K int

// maxMin
// fairness is defined as max(x1,x2,…,xk) − min(x1,x2,…,xk)
// Solution sort the numbers then do a K element scan through
// the sorted array.
//
// Time is dominated by sorting.
// Numbers are 10 digits at most so radix sorting would give us
// close to linear.
func maxMin(a []int64) int64 {

	is := Int64Slice(a)
	sort.Sort(is)

	return scanK(is)
}

func scanK(a []int64) int64 {
	var unfairness int64
	unfairness = (1 << 63) - 1
	for i := 0; i < len(a)-K+1; i++ {
		min, max := a[i], a[i+K-1]
		if max-min < unfairness {
			unfairness = max - min
		}
	}
	return unfairness
}

func read() {
	results := make([]int64, 0)

	fmt.Scan(&N)
	fmt.Scan(&K)

	s := bufio.NewScanner(os.Stdin)

	for i := int64(0); i < N; i++ {
		s.Scan()
		x, err := strconv.ParseInt(s.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		results = append(results, x)
	}

	fmt.Printf("%d\n", maxMin(results))
}

func main() {
	read()
}
