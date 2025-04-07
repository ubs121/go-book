package dsa

import (
	"fmt"
	"index/suffixarray"
	"slices"
	"testing"
)

// index/suffixarray
func TestSuffixArray(t *testing.T) {
	// create an index for data
	index := suffixarray.New([]byte("banana"))

	// lookup
	offsets := index.Lookup([]byte("ana"), -1)
	for _, off := range offsets {
		fmt.Println(off)
	}
}

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 55, 77}

	x := 4
	i, found := slices.BinarySearch(arr, x)
	fmt.Printf("%d found=%v, loc~ %d\n", x, found, i)

	x = 10
	i, found = slices.BinarySearch(arr, x)
	fmt.Printf("%d found=%v, loc~ %d\n", x, found, i)
}
