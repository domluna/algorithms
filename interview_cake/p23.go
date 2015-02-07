// Does the Linked-List have a cycle?
package main

import "fmt"

type node struct {
	next *node
	data interface{}
}

func cycle(n *node) bool {
	r1 := n
	r2 := n

	for {
		if r1 == nil {
			return false
		}
		r1 = r1.next

		if r2 == nil || r2.next == nil {
			return false
		}

		r2 = r2.next.next

		if r1 == r2 {
			return true
		}

	}
}

func main() {
	var a, b, c, d *node
	var want, got bool

	a = &node{data: "a"}
	b = &node{data: "b"}
	c = &node{data: "c"}
	a.next = b
	b.next = c
	c.next = nil

	want = false
	got = cycle(a)

	if got != want {
		fmt.Printf("T1, no cycle: got %v, want %v\n", got, want)
	}

	a = &node{data: "a"}
	b = &node{data: "b"}
	c = &node{data: "c"}
	a.next = b
	b.next = c
	c.next = a

	want = true
	got = cycle(a)

	if got != want {
		fmt.Printf("T2, cycle: got %v, want %v\n", got, want)
	}

	a = &node{data: "a"}
	a.next = a

	want = true
	got = cycle(a)

	if got != want {
		fmt.Printf("T3 cycle: got %v, want %v\n", got, want)
	}

	a = &node{data: "a"}
	b = &node{data: "b"}
	c = &node{data: "c"}
	d = &node{data: "d"}
	a.next = b
	b.next = c
	c.next = d
	d.next = b

	want = true
	got = cycle(a)

	if got != want {
		fmt.Printf("T4, cycle: got %v, want %v\n", got, want)
	}
}
