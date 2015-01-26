// Given a list of tuples of meetings
// (startTime, endTime)
//
// Returns a list of condensed meeting times.
package main

import (
	"fmt"
	"reflect"
	"sort"
)

type Meeting struct {
	start int
	end   int
}

type Meetings []Meeting

func (m Meetings) Len() int {
	return len(m)
}

func (m Meetings) Less(i, j int) bool {
	return m[i].start < m[j].start
}

func (m Meetings) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

// Cases
// 1. cm.end < sm.start, append cm, cm = sm
// 2. cm.end >= sm.start
//  a. cm.end >= sm.end, leave cm unchanged, increment
//  b. cm.end < sm.end, cm.end = sm.end, increment
func condenseMeetings(m Meetings) Meetings {
	s := 1
	currMeet := m[0]
	newMeets := Meetings{}

	for s < len(m) {
		secondMeet := m[s]
		if currMeet.end >= secondMeet.start {
			if currMeet.end < secondMeet.end {
				currMeet.end = secondMeet.end
			}
		} else {
			newMeets = append(newMeets, currMeet)
			currMeet = secondMeet
		}
		s++
	}
	newMeets = append(newMeets, currMeet)
	return newMeets

}

func computeMeetings(m Meetings) Meetings {
	sort.Sort(m)
	return condenseMeetings(m)
}

var tests = []struct {
	in  Meetings
	out Meetings
}{
	{
		Meetings{
			Meeting{0, 1},
			Meeting{3, 4},
			Meeting{4, 8},
			Meeting{10, 12},
			Meeting{9, 10},
		},
		Meetings{
			Meeting{0, 1},
			Meeting{3, 8},
			Meeting{9, 12},
		},
	},
	{
		Meetings{
			Meeting{1, 2},
			Meeting{2, 3},
		},
		Meetings{
			Meeting{1, 3},
		},
	},
	{
		Meetings{
			Meeting{1, 10},
			Meeting{2, 3},
			Meeting{5, 7},
			Meeting{3, 10},
			Meeting{8, 9},
		},
		Meetings{
			Meeting{1, 10},
		},
	},
	{

		Meetings{
			Meeting{1, 5},
			Meeting{2, 3},
		},
		Meetings{
			Meeting{1, 5},
		},
	},
	{

		Meetings{
			Meeting{1, 2},
			Meeting{3, 4},
		},
		Meetings{
			Meeting{1, 2},
			Meeting{3, 4},
		},
	},
}

func main() {
	for _, tt := range tests {
		result := computeMeetings(tt.in)
		if !reflect.DeepEqual(result, tt.out) {
			fmt.Printf("want %v, got %v\n", tt.out, result)
		}
	}
}
