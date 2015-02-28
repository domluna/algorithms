package main

import (
	"fmt"
	"os"
)

func funnyString(s string) bool {
	n := len(s)
	for i := 1; i < n; i++ {
		// fmt.Printf("%c %c %c %c\n", s[i], s[i-1], s[n-i-1], s[n-i])
		// fmt.Println(s[i], s[i-1], s[n-i-1], s[n-i])
		// fmt.Println(s[i]-s[i-1], s[n-i-1]-s[n-i])
		if absSub(s[i], s[i-1]) != absSub(s[n-i-1], s[n-i]) {
			return false
		}
	}
	return true
}

func absSub(x, y uint8) uint8 {
	if x > y {
		return x - y
	}
	return y - x
}

func read() {
	var T int
	var s string
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&s)
		if funnyString(s) {
			fmt.Fprintln(os.Stdout, "Funny")
		} else {
			fmt.Fprintln(os.Stdout, "Not Funny")
		}
	}
}

func main() {
	read()
}
