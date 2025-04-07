package common

import (
	"testing"
)

// https://leetcode.com/problems/container-with-most-water/
func maxArea(height []int) int {
	// Area = length of shorter vertical line * distance between lines
	l := 0
	r := len(height)
	max := 0
	lh := 0
	rh := 0
	area := 0

	for l < r {
		lh = height[l]
		rh = height[r-1]

		if lh < rh {
			area = lh * (r - l - 1)
			l++
		} else {
			area = rh * (r - l - 1)
			r--
		}

		if max < area {
			max = area
		}
	}
	return max
}

func TestMaxArea(t *testing.T) {
	testCases := []struct {
		arr []int
		exp int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{4, 3, 2, 1, 4}, 16},
		{[]int{1, 2, 1}, 2},
		{[]int{7, 1, 3, 20, 1, 4}, 21},
		{[]int{1, 3, 2, 5, 25, 24, 5}, 24},
	}

	for i, tc := range testCases {
		got := maxArea(tc.arr)
		if got != tc.exp {
			t.Errorf("tc %d: exp: %d, got %d", i, tc.exp, got)
		}
	}
}
