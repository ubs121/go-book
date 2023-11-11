package numbers

import (
	"testing"
)

// https://codingcompetitions.withgoogle.com/kickstart/round/00000000008f4a94/0000000000b5496b
func countValidParenthesis(l, r int) int {
	// odd brakets divides the whole string, so don't put in the middle
	// no nesting because it divides
	m := min(l, r)
	sol := m * (m + 1) / 2
	return sol
}

func TestCoundValidParenthesis(t *testing.T) {
	testCases := []struct {
		l, r int
		exp  int
	}{
		{1, 0, 0},
		{1, 1, 1},
		{2, 2, 3},
		{3, 2, 3},
		{5, 7, 15},
	}
	for i, tc := range testCases {
		got := countValidParenthesis(tc.l, tc.r)
		if got != tc.exp {
			t.Errorf("tc %d: exp %d, got %d", i, tc.exp, got)
		}
	}
}
