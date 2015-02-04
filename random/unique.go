package main

import "fmt"

// unique returns true if all characters (runes) in s
// are unique.
func unique(s string) bool {
	set := make(map[rune]struct{})
	for _, c := range s {
		// if it's in already, not unique
		if _, ok := set[c]; ok {
			return false
		}
		set[c] = struct{}{}
	}
	return true
}

var tests = map[string]bool{
	"abcd":    true,
	"aa":      false,
	"abdcefa": false,
	"":        true,
	"a":       true,
	"bb":      false,
}

func main() {
	for s, b := range tests {
		got := unique(s)
		if got != b {
			fmt.Println("unique(%s): got %v, want %v", s, got, b)
		}
	}
}
