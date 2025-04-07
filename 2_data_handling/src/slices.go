package dsa

import (
	"sort"
	"strconv"
	"strings"
)

// Returns sum
func Sum(nums []int) int {
	s := 0
	for _, num := range nums {
		s += num
	}
	return s
}

// IndexMax returns max value indexes
func IndexMax(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}

	var maxIxs []int
	maxVal := arr[0] // default max

	for i := 0; i < len(arr); i++ {
		if maxVal < arr[i] {
			maxVal = arr[i] // new max
			maxIxs = nil    // reset because another max found
		} else if maxVal == arr[i] {
			// collect max positions
			maxIxs = append(maxIxs, i)
		}
	}

	return maxIxs
}

// Parses text into integer array
func ParseIntArray(strArr string) []int {
	items := strings.Split(strArr, ",")
	var arr []int
	for _, s := range items {
		if len(s) > 0 {
			n, _ := strconv.Atoi(s)
			arr = append(arr, n)
		}
	}
	return arr
}

func CmpUnorderedStringArray(a, b []string) bool {
	if len(a) != len(b) {
		return false // not equal length
	}

	sort.Strings(a)
	sort.Strings(b)
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
