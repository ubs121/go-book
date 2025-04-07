package dsa

/*
A cafeteria table consists of a row of N seats, numbered from  1 to N from left to right. Social distancing guidelines require that every diner be seated such that
K seats to their left and K seats to their right (or all the remaining seats to that side if there are fewer than K) remain empty.
There are currently M diners seated at the table, the i-th of whom is in seat S[i]. No two diners are sitting in the same seat, and the social distancing guidelines are satisfied.
Determine the maximum number of additional diners who can potentially sit at the table without social distancing guidelines being violated for any new or existing diners,
assuming that the existing diners cannot move and that the additional diners will cooperate to maximize how many of them can sit down.

Sample Explanation

N = 10
K = 1
M = 2
S = [2, 6]

In this case, the cafeteria table has N=10 seats, with two diners currently at seats 2 and 6 respectively.
The table initially looks as follows, with brackets covering the K=1 seat to the left and right of each existing diner that may not be taken.
  1 2 3 4 5 6 7 8 9 10
  [   ]   [   ]
Three additional diners may sit at seats
4, 8, and 10 without violating the social distancing guidelines.
*/

import (
	"sort"
	"testing"
)

func getMaxAdditionalDinersCount(N int64, K int64, M int32, S []int64) int64 {
	if len(S) == 0 {
		return getMaxCountInOpen(1, N+1, K)
	}

	// sort occupied seats
	sort.Slice(S, func(i, j int) bool { return S[i] < S[j] })

	ret := int64(0)
	off := int64(1)
	for i := 0; i < len(S); i++ {
		ret += getMaxCountInOpen(off, S[i]-K, K)
		off = S[i] + K + 1
	}

	// the remaining seats after S[M-1]
	ret += getMaxCountInOpen(S[M-1]+K+1, N+1, K)

	return ret
}

// calculates the open seats in interval [lo:hi[,  excluding 'hi'
func getMaxCountInOpen(lo int64, hi int64, K int64) int64 {
	n := hi - lo
	ret := n / (K + 1)
	if n%(K+1) > 0 {
		return ret + 1
	}
	return ret
}

func TestGetMaxAdditionalDinersCount(t *testing.T) {
	testCases := []struct {
		N, K, M int64
		S       []int64
		exp     int64
	}{
		{10, 1, 2, []int64{2, 6}, 3},
		{15, 2, 3, []int64{11, 6, 14}, 1},
	}
	for i, tc := range testCases {
		got := getMaxAdditionalDinersCount(tc.N, tc.K, int32(tc.M), tc.S)
		if tc.exp != got {
			t.Errorf("tc %d: exp %d, got %d", i, tc.exp, got)
		}
	}
}
