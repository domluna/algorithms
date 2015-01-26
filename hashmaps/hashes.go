// Experiments with non-crypto hash functions.
package main

import (
	"fmt"
	"hash/adler32"
)

func main() {
	h := adler32.New()
	h.Write([]byte("hello world"))
	fmt.Printf("%v\n", h.Sum32())
}
