package dsa

import (
	"fmt"
	"slices"
	"testing"
)

func TestParseArray(t *testing.T) {
	a := ParseIntArray("")
	if len(a) > 0 {
		t.Errorf("non-empty array: %v", a)
	}

	a = []int{1}
	a = a[0:0]
	fmt.Printf("%v", a)
}

func TestInitSlice(t *testing.T) {
	var arr = []int{1, 2, 3}

	fmt.Println(arr)
}

func TestSliceClone(t *testing.T) {
	arr := []int{1, 2, 1}
	arr1 := slices.Clone(arr)
	arr1[2] = 3
	fmt.Println(arr, arr1)
}

func TestSlicesContains(t *testing.T) {
	arr := []int{1, 2, 1}
	fmt.Println(slices.Contains(arr, 3))
}
