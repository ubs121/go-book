package common

import (
	"container/ring"
	"fmt"
	"testing"
)

func TestRing(t *testing.T) {
	// Create a new ring with 5 elements
	r := ring.New(5)

	// Populate the ring with values
	for i := 1; i <= 5; i++ {
		r.Value = i
		r = r.Next()
	}

	// Traverse the ring and verify the values
	r.Do(func(value interface{}) {
		fmt.Print(value, ",")
	})
	fmt.Println()

	// Move the ring pointer by 2 positions
	r = r.Move(2)

	// Traverse the ring and verify the values
	r.Do(func(value interface{}) {
		fmt.Print(value, ",")
	})
	fmt.Println()

	// Get the value at the current position
	value := r.Value.(int)
	expectedValue := 3

	// Check if the value matches the expected value
	if value != expectedValue {
		t.Errorf("Expected value %d, but got %d", expectedValue, value)
	}
}

/*
Josephus Circle: N people have decided to elect a leader by arranging themselves in a circle and eliminating every Mth person around the circle,
closing ranks as each person drops out. Find which person will be the last one remaining (with rank 1).
*/
func getJosephusPosition(n, m int) int {
	// Create a circular list of size N.
	r := make([]int, n)

	// Populate the ring with 1..n values
	for i := 0; i < n; i++ {
		r[i] = i + 1 // starts from 1
	}

	// repeat until one node is left
	m1 := 0
	for len(r) > 1 {
		// Find m-th node
		i := (m1 + m - 1) % len(r)
		// Remove the m-th node
		r = append(r[:i], r[i+1:]...)
		m1 = i // update m1
	}

	return r[0] // Last person left standing (Josephus Position)
}

func TestGetJosephusPosition(t *testing.T) {
	testCases := []struct {
		n, m int
		exp  int
	}{
		{
			4, 2, 1,
		},
		{
			5, 2, 3,
		},
		{
			14, 2, 13,
		},
		{
			14, 1, 14,
		},
	}
	for i, tc := range testCases {
		got := getJosephusPosition(tc.n, tc.m)
		if got != tc.exp {
			t.Errorf("%d: exp %d, got %d", i, tc.exp, got)
		}
	}
}
