// Recursive String Permutations
package main

import "fmt"

var tests = struct {
	s    string
	want []string
}{}

// TODO: complete function
func computePermutations(s string) []string {
	return nil
}

func main() {
	for _, t := range tests {
		got := computePermutations(t.s)
		if got != t.want {
			fmt.Errorf("input %s: got %v, want %v",
				t.s, got, t.want)
		}
	}
}
