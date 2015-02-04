// Largest Stack
package main

import "fmt"

type MaxStack struct {
	items []int

	// This stack contains the max items
	// items are only pushed unto this
	// stack is it is greater than the current max.
	maxItems []int
}

func New() *MaxStack {
	return &MaxStack{
		items:    make([]int, 0),
		maxItems: make([]int, 0),
	}
}

func (s *MaxStack) Push(x int) {
	s.items = append(s.items, x)
	max := s.Max()
	if x > max {
		s.maxItems = append(s.maxItems, x)
	}
}
func (s *MaxStack) Pop() int {
	l := len(s.items) - 1
	popped := s.items[l]
	max := s.Max()
	if popped == max {
		s.maxItems = s.maxItems[:len(s.maxItems)-1]
	}
	s.items = s.items[:l]
	return popped
}
func (s *MaxStack) Max() int {
	if len(s.maxItems) > 0 {
		l := len(s.maxItems) - 1
		return s.maxItems[l]
	}
	return 0
}

func main() {
	s := New()
	s.Push(5)
	s.Push(10)
	s.Push(7)
	s.Push(20)
	s.Push(25)

	max := s.Max()
	if max != 25 {
		fmt.Println("ERROR: got, %d, want %d", max, 25)
		return
	}

	s.Pop() // pop 25

	max = s.Max()
	if max != 20 {
		fmt.Println("ERROR: got, %d, want %d", max, 20)
		return
	}

	s.Pop() // pop 20

	max = s.Max()
	if max != 10 {
		fmt.Println("ERROR: got, %d, want %d", max, 10)
		return
	}

	s.Pop()

	max = s.Max()
	if max != 10 {
		fmt.Println("ERROR: got, %d, want %d", max, 10)
		return
	}

	s.Pop()

	max = s.Max()
	if max != 5 {
		fmt.Println("ERROR: got, %d, want %d", max, 5)
		return
	}
}
