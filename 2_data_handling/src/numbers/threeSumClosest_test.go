package numbers

import (
	"math"
	"sort"
	"testing"
)

// https://leetcode.com/problems/3sum-closest/
func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)

	sum := nums[n-1] + nums[n-2] + nums[n-3] // answer
	dist3 := math.MaxInt64                   // distance of 3 sum

	for i := 0; i < n-2; i++ {
		x := nums[i]
		target2 := target - x

		// find two closest sum to 'target2'
		sum2 := twoSumClosest(nums[i+1:], target2)
		dist2 := dist(target, x+sum2)
		if dist2 < dist3 {
			dist3 = dist2
			sum = x + sum2
			if dist3 == 0 {
				break
			}
		}

	}

	return sum
}

func twoSumClosest(nums []int, target int) int {
	n := len(nums)     // must be > 2
	l := 0             // left index
	r := n - 1         // right index
	sum := 0           // initial sum
	d := math.MaxInt64 // distance between two sum and target

	for l < r {
		sum2 := nums[l] + nums[r]
		dist2 := dist(sum2, target)

		if dist2 < d {
			d = dist2
			sum = sum2
		}

		if sum2 == target {
			return sum2
		} else if sum2 > target {
			r--
		} else {
			l++
		}
	}

	return sum
}

func dist(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func Test3SumClosest(t *testing.T) {
	testCases := []struct {
		arr    []int
		target int
		exp    int
	}{
		{[]int{-1, 2, 1, -4}, 1, 2},
		{[]int{1, 2, 2, 3, 3, 4}, 2, 5},
		{[]int{0, 0, 0}, 1, 0},
		{[]int{0, 1, 2}, 0, 3},
		{[]int{1, 1, 1, 1}, 100, 3},
		{[]int{1, 1, 1, 0}, 100, 3},
		{[]int{-3, -2, -5, 3, -4}, -1, -2},
		{[]int{1, 2, 4, 8, 16, 32, 64, 128}, 82, 82},
		{[]int{0, 2, 1, -3}, 1, 0},
		{[]int{-1, 0, 1, 1, 55}, 3, 2},
		{[]int{1, 1, 1, 0}, -100, 2},
		{[]int{4, 0, 5, -5, 3, 3, 0, -4, -5}, -2, -2},
	}

	for i, tc := range testCases {
		got := threeSumClosest(tc.arr, tc.target)
		if got != tc.exp {
			t.Errorf("tc %d: exp %d, got %d", i, tc.exp, got)
		}
	}
}
