// Recursive String Permutations
package main

import "fmt"

var tests = []struct {
	s    string
	want []string
}{
	{
		s:    "abc",
		want: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
	},
	{
		s:    "c",
		want: []string{"c"},
	},
	{
		s:    "",
		want: nil,
	},
}

// TODO: complete function
// use memoization
// IDEA: say we have the permutations for the n-1 letters computed.
// Then to get the permutations for n letters we just append another letter
// to the end.
//
// We can't just append any letter it has to be a letter that hasn't be used
// yet, but that's the gist.
//
func computePermutations(s string) []string {
	if len(s) < 1 {
		return nil
	}

	table := make(map[int][]string)
	return compute(s, len(s)-1, table)
}

func compute(s string, length int, table map[int][]string) []string {
	if length < 1 {
		perms := make([]string, 0)
		perms = append(perms, s)

		table[length] = perms
		return perms
	}

	var perms []string
	var ok bool

	// precomputed
	if perms, ok = table[length]; !ok {
		perms = compute(s[:length], length-1, table)
	}

	newPerms := make([]string, 0)
	char := string(s[length])

	// For each permutation add on current char
	// to every position of the permutation.
	for _, p := range perms {
		for i := range p {
			newPerms = append(newPerms, p[:i]+char+p[i:])
		}
		// add on the end
		newPerms = append(newPerms, p+char)
	}

	table[length] = newPerms
	return newPerms
}

// check if the the two permutations are equal
func equalPerms(p1, p2 []string) bool {
	if len(p1) != len(p2) {
		return false
	}

	for _, s1 := range p1 {
		var in bool
		for _, s2 := range p2 {
			if s1 == s2 {
				in = true
			}
		}
		if !in {
			return false
		}
	}
	return true
}

func main() {
	for _, t := range tests {
		got := computePermutations(t.s)
		if !equalPerms(got, t.want) {
			fmt.Printf("ERROR input %q: got %v, want %v\n",
				t.s, got, t.want)
		}
	}
}
