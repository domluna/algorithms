package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// countReductions counts the number of reductions to convert
// the string to a palindrome.
// Letters > in alphabetical order can be reduced to letters
// earlier in order.
//
// Ex. d -> a, 3 reduction to get from d to a; a however, cannot
// be reduced to d.
//
// s is guaranteed to be lowercase.
func countReductions(s string) int {

	l := len(s) - 1
	reductions := 0

	for i := 0; i <= (len(s)-1)/2; i++ {
		start := s[i]
		end := s[l-i]

		is := strings.Index(alphabet, string(start))
		ie := strings.Index(alphabet, string(end))

		// fmt.Println(s, string(start), string(end), is, ie)

		if is < ie {
			reductions += ie - is
		} else {
			reductions += is - ie
		}
	}

	return reductions
}

func read() {
	var N int
	results := make([]string, 0)

	fmt.Scan(&N)

	r := bufio.NewReader(os.Stdin)

	for i := 0; i < N; i++ {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			results = append(results, line)
		} else {
			results = append(results, line[:len(line)-1]) // remove newline
		}
	}

	for _, v := range results {
		fmt.Fprintf(os.Stdout, "%d\n", countReductions(v))
	}
}

func main() {
	read()
}
