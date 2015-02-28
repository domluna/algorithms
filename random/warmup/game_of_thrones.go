package main

import "fmt"

func anagramOfPalindrome(s string) bool {
	// do something to unscramble the string
	// check if it is a palindrome

	// a string is a palindrome if
	// 1. # of chars are all even (even length)
	// 2. # of chars are all even except (odd length)
	// one char whose number is odd

	m := make(map[rune]int)

	for _, v := range s {
		if _, ok := m[v]; !ok {
			m[v] = 1
		} else {
			m[v]++
		}
	}

	if len(s)%2 == 0 { // even
		for _, v := range m {
			if v%2 != 0 {
				return false
			}
		}
	} else { // odd
		c := 0
		for _, v := range m {
			if v%2 != 0 {
				c++
			}
		}

		if c > 1 {
			return false
		}
	}

	return true
}

func read() {
	var s string
	fmt.Scan(&s)

	if anagramOfPalindrome(s) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func main() {
	read()
}
