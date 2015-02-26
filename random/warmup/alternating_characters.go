// We want strings to be alternating characters.
// Ex. ABAB is good, ABBA is bad, BB does not alternate.
// In this case we would need 1 deletion to turn ABBA to
// ABA.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// countDeletions counts the minimum number of deletions.
func countDeletions(s string) int {
	deletions := 0
	f := s[0]

	for i := 1; i < len(s); i++ {
		o := s[i]
		if f == o {
			deletions++
		}
		f = o
	}

	return deletions
}

func read() {
	var N int
	results := make([]string, 0)

	fmt.Scan(&N)

	r := bufio.NewReader(os.Stdin)

	var err error
	var line string

	for err != io.EOF {
		line, err = r.ReadString('\n')
		results = append(results, line)
	}

	for _, v := range results {
		fmt.Printf("%d\n", countDeletions(v))
	}
}

func main() {
	read()
}
