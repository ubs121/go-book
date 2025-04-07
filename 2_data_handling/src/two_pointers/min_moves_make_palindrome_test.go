package main

import (
	"testing"
)

func minMovesToMakePalindrome(s string) int {
	n := len(s)
	runes := []rune(s)
	moves := 0

	for i, j := 0, n-1; i < j; i++ {
		k := j
		for k > i {
			if runes[i] == runes[k] {
				for k < j {
					runes[k], runes[k+1] = runes[k+1], runes[k]
					moves++
					k++
				}
				j--
				break
			}
			k--
		}
		if k == i {
			moves += n/2 - i
		}
	}
	return moves
}

func TestMinMovesToMakePalindrome(t *testing.T) {
	in := "arcacer"
	got := minMovesToMakePalindrome(in)
	exp := 4
	if got != exp {
		t.Errorf("exp %v, got %v", exp, got)
	}
}
