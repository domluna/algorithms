// Cooler solution would be to use a channel and stream the results
// since the number of combinations strings grows exponentially.
package main

import (
	"fmt"
	"strconv"
)

// combos returns the combinations of s.
// In this context a combination is the string
// in order but one or more letters can be replaced
// with a number represented the amount of letters.
//
// Ex. s = home
//	...
//	3e
//	2m1
//	home
//	hom1
//	...
//
// So "3e" is valid because there are 3 letters in home before "e".
// "2m1", 2 lettes before "m" and 1 after.
//
// All in all there will be 2^n combinations where n is len(s) .
func combos(s string) []string {

	if len(s) < 2 {
		return []string{s, "1"}
	}

	results := make([]string, 0)

	head := string(s[0])

	tail := combos(s[1:])

	for _, c := range tail {
		if v, err := strconv.Atoi(string(c[0])); err == nil {
			results = append(results, strconv.Itoa(v+1)+c[1:])
		} else {
			results = append(results, "1"+c)
		}
		results = append(results, head+c)
	}

	return results
}

func read() {
	var s string
	fmt.Scan(&s)

	cs := combos(s)

	fmt.Printf("Number of combinations: %d\n", len(cs))
	for _, c := range cs {
		fmt.Println(c)
	}
}

func main() {
	read()
}
