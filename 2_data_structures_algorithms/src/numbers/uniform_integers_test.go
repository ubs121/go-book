package numbers

import (
	"strconv"
	"strings"
	"testing"
)

// A positive integer is considered uniform if all of its digits are equal. For example, 222 is uniform, while  223 is not.
// Given two positive integers A and B, determine the number of uniform integers between A and B, inclusive.
func getUniformIntegerCountInInterval(A int64, B int64) int32 {
	aStr := strconv.Itoa(int(A))
	bStr := strconv.Itoa(int(B))
	loDigits := len(aStr)
	hiDigits := len(bStr)

	total := (hiDigits - loDigits + 1) * 9 // total uniform integers

	// subtract integers less than A
	total = total - int(aStr[0]-'0')

	if aUni := strings.Repeat(string(aStr[0]), loDigits); aUni >= aStr {
		total++ // aUni is inclusive
	}

	// subtract integers greater than B
	total = total - (10 - int(bStr[0]-'0'))

	if bUni := strings.Repeat(string(bStr[0]), hiDigits); bUni <= bStr {
		total++ // bUni is inclusive
	}

	return int32(total)
}

func TestGetUniformIntegerCountInInterval(t *testing.T) {
	testCases := []struct {
		A, B int64
		exp  int32
	}{
		{75, 300, 5},
		{1, 9, 9},
		{999999999999, 999999999999, 1},
		{75, 3000, 14},
	}
	for i, tc := range testCases {
		got := getUniformIntegerCountInInterval(tc.A, tc.B)
		if tc.exp != got {
			t.Errorf("tc %d: exp %d, got %d", i, tc.exp, got)
		}
	}
}
