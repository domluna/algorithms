package main

import (
	"fmt"
	"os"
)

// palindromeIndex returns the index whose deletion
// would cause in s becoming a palindrome.
// -1 is returned if s is already a palindrome.
func palindromeIndex(s string) int {
	n := len(s)
	for i := 0; i < n/2; i++ {
		// problem, find which index
		// is at fault.
		if s[i] != s[n-i-1] {
			if palindrome(s[:i] + s[i+1:]) {
				return i
			} else {
				return n - i - 1
			}
		}
	}
	return -1
}

func palindrome(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		// problem, find which index
		// is at fault.
		if s[i] != s[n-i-1] {
			return false
		}
	}
	return true
}

// read input
func read() {
	var T int
	var s string

	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&s)
		fmt.Fprintf(os.Stdout, "%d\n", palindromeIndex(s))
	}
}

func main() {
	read()
}
