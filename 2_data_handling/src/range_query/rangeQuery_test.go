package search

import (
	"math"
	"testing"
)

// Query a sparse table in range [l,r)
func minQuery(sparseTable [][]int, l, r int) int {
	p := int(math.Log(float64(r - l + 1)))
	return min(sparseTable[p][l], sparseTable[p][r-(1<<p)])
}

// Builds a sparse table for Range Minimum Query
func buildSparseTable(arr []int) [][]int {

	n := len(arr)                       // # of columns
	h := int(math.Log2(float64(n))) + 1 // # of rows (height)
	st := make([][]int, h)              // sparse table, st[k][j]=min(range(j:j+2^k))

	// base case: 2^0
	st[0] = make([]int, n)
	copy(st[0], arr)

	// iterative dynamic programming approach
	for k := 1; k < h; k++ {
		st[k] = make([]int, n-(1<<k)+1) // cut to actual length
		for j := 0; j+(1<<k)-1 < n; j++ {
			st[k][j] = min(st[k-1][j], st[k-1][j+(1<<(k-1))])
		}
	}
	return st
}

func TestSparseTableRMQ(t *testing.T) {
	testCases := []struct {
		arr []int
		rng []int
		exp int
	}{
		{[]int{3, 1, 5, 3, 4, 7, 6, 1}, []int{3, 5}, 3},
		{[]int{3, 1, 5, 3, 4, 7, 6, 1}, []int{2, 8}, 1},
		{[]int{1}, []int{0, 1}, 1},
	}

	for i, tc := range testCases {
		// build a sparse table
		st := buildSparseTable(tc.arr)

		// querying the range and getting the minimum value using the sparse table
		got := minQuery(st, tc.rng[0], tc.rng[1])
		if tc.exp != got {
			t.Errorf("tc: %d exp: %d got %d", i, tc.exp, got)
		}
	}
}
