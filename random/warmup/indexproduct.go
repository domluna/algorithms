package main

import "fmt"

type Pair struct {
	index int
	value int
}

type Storer struct {
	nums     []int
	leftMap  map[int]Pair
	rightMap map[int]Pair
}

func (s *Storer) IndexProduct() int {
	best := 0
	for i := range a {
		v := s.(i) * s.right(i)
		if v > best {
			best = v
		}
	}
	return best
}

func (s *Storer) right(i int) int {
	n := len(s.nums)

	// check if cached
	if v, ok := s.rightMap[i]; ok {
		return v.index
	}

	for ; i < n; i++ {
		iv := s.nums[i]

		if iv < s.sums[i+1] {
			s.rightMap[i] =
		}
		if v, ok := s.rightMap[i+1]; ok {
			if iv < v.value
		}
	}
}

func (s *Storer) left(i int) {
}

func read() {
	var c, r int
	fmt.Scan(&r, &c)

	fmt.Println(strangeGrid(r, c))
}

func main() {
	read()
}
