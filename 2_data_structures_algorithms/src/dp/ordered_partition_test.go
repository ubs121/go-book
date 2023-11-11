package dsa

import (
	"math"
	"testing"
)

/*
Problem: Integer Partition without Rearrangement
Input: An arrangement S of non-negative numbers s1, . . . , sn and an integer k.
Output: Partition S into k or fewer ranges, to minimize the maximum sum over all the ranges, without reordering any of the numbers.

Examples: k=3
100 100 100 | 100 100 100 | 100 100 100
100 200 300 400 500 | 600 700 | 800 900
*/

func partition(s []int, k int) int {
	if k == 1 {
		return sum(s)
	}

	if len(s) == 1 {
		return s[0]
	}

	// TODO: use DP
	totalMin := math.MaxInt
	for i := 1; i < len(s); i++ {
		// go into smaller problem with (s[:i], k-1)
		newMax := max(partition(s[:i], k-1), sum(s[i:]))
		if newMax < totalMin {
			totalMin = newMax
		}
	}

	return totalMin
}

func sum(a []int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	return s
}

func TestPartition(t *testing.T) {
	testCases := []struct {
		arr []int
		k   int
		exp int
	}{
		{[]int{100, 100, 100, 100, 100, 100, 100, 100, 100}, 2, 500},
		{[]int{100, 100, 100, 100, 100, 100, 100, 100, 100}, 3, 300},
		{[]int{100, 100, 100, 100, 100, 100, 100, 100, 100}, 4, 300},

		{[]int{100, 200, 300, 400, 500, 600, 700, 800, 900}, 2, 2400},
		{[]int{100, 200, 300, 400, 500, 600, 700, 800, 900}, 3, 1700},
		{[]int{100, 200, 300, 400, 500, 600, 700, 800, 900}, 4, 1500},
		{[]int{100, 200, 300, 400, 500, 600, 700, 800, 900}, 5, 1100},
		{[]int{100, 200, 300, 400, 500, 600, 700, 800, 900}, 6, 900},
		{[]int{100, 200, 300, 400, 500, 600, 700, 800, 900}, 7, 900},
		{[]int{100, 200, 300, 400, 500, 600, 700, 800, 900}, 8, 900},
	}
	for i, tc := range testCases {
		got := partition(tc.arr, tc.k)
		if got != tc.exp {
			t.Errorf("%d: exp %d, got %d", i, tc.exp, got)
		}
	}
}
