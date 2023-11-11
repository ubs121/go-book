package sort

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

func permuteRecursiveBacktrack(depth int, perm []byte, inPerm []bool, original []byte) {
	if depth == len(original) {
		println(string(perm))
	} else {
		for i := 0; i < len(original); i++ {
			// pick next un-used letter
			if !inPerm[i] {
				inPerm[i] = true
				perm[depth] = original[i]
				permuteRecursiveBacktrack(depth+1, perm, inPerm, original)
				inPerm[i] = false
			}
		}
	}
}

func TestPermuteRecurse(t *testing.T) {
	a := []byte{'a', 'b', 'c'} /* сэлгэх утгууд */

	// initial state
	perm := make([]byte, len(a))
	used := make([]bool, len(a))
	copy(perm, a)

	permuteRecursiveBacktrack(0, perm, used, a)
}

// recursive version with channel
func permuteRecurseChannel(k int, nums []int, out chan []int) {
	if k == 1 {
		// output
		cmb := make([]int, len(nums))
		copy(cmb, nums)
		out <- cmb
	} else {
		permuteRecurseChannel(k-1, nums, out)

		for i := 0; i < k-1; i++ {
			if k%2 == 0 {
				nums[i], nums[k-1] = nums[k-1], nums[i]
			} else {
				nums[0], nums[k-1] = nums[k-1], nums[0]
			}

			permuteRecurseChannel(k-1, nums, out)
		}
	}
}

func TestPermuteRecurseChannel(t *testing.T) {
	a := []int{1, 2, 3}

	out := make(chan []int)
	go func() {
		defer close(out)
		permuteRecurseChannel(len(a), a, out)
	}()

	for v := range out {
		fmt.Printf("%v\n", v)
	}
}

// N ширхэг тоогоор зохиох бүх боломжит сэлгэмэл
func PermuteStack(nums []int) [][]int {
	n := len(nums)

	var ret [][]int

	// output initial state
	cmb := make([]int, n)
	copy(cmb, nums)
	ret = append(ret, cmb)

	// an encoding of the stack state.
	p := make([]int, n)

	i := 1
	for i < n {
		if p[i] < i {
			j := 0
			if i%2 > 0 {
				j = p[i]
			}

			// swap
			nums[j], nums[i] = nums[i], nums[j]

			// output
			cmb := make([]int, n)
			copy(cmb, nums)
			ret = append(ret, cmb)

			p[i]++
			i = 1
		} else {
			p[i] = 0
			i++
		}
	}
	return ret
}

func TestPermuteStack(t *testing.T) {
	a := []int{1, 2, 3} /* сэлгэх утгуудыг агуулах массив */
	ret := PermuteStack(a)
	fmt.Printf("%v", ret)
}

func TestPermutation1(t *testing.T) {
	var (
		a [100]int /* сэлгэх утгуудыг агуулах массив */
		N int      /* сэлгэх элементийн тоо */
	)

	N = 3
	display := func() {
		for i := 0; i < N; i++ {
			print(a[i], " ")
		}
		println()
	}

	for i := 0; i < N; i++ {
		a[i] = i + 1
	}

	p := make([]int, N) /* сэлгэмлийг удирдах массив */
	var (
		j    int
		temp int
	)

	/* эхний хувилбарыг хэвлэх */
	display()

	for i := 1; i < N; {
		if p[i] < i {
			if i%2 > 0 {
				j = p[i]
			} else {
				j = 0
			}
			temp = a[j]
			a[j] = a[i]
			a[i] = temp

			/* сэлгэмлийн нэг хувилбарыг хэвлэх */
			display()

			p[i]++
			i = 1
		} else {
			p[i] = 0
			i++
		}
	}
}

// Given a list of integers, find the next permutation in lexicographic order.
func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2

	// Find the first pair where nums[i] < nums[i+1]
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		j := n - 1
		// Find the first element greater than nums[i]
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		// Swap the elements at indices i and j
		nums[i], nums[j] = nums[j], nums[i]
	}

	// Reverse the sublist starting from index i+1
	// This ensures that the next permutation is as small as possible while still being greater than the current permutation
	slices.Reverse(nums[i+1:])
}

func TestNextPermutation(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{1, 2, 3},
			expected: []int{1, 3, 2},
		},
		{
			nums:     []int{3, 2, 1},
			expected: []int{1, 2, 3},
		},
		{
			nums:     []int{1, 1, 5},
			expected: []int{1, 5, 1},
		},
		{
			nums:     []int{1},
			expected: []int{1},
		},
		{
			nums:     []int{1, 3, 2},
			expected: []int{2, 1, 3},
		},
	}

	for _, tc := range testCases {
		nextPermutation(tc.nums)
		if !reflect.DeepEqual(tc.nums, tc.expected) {
			t.Errorf("Expected %v, but got %v", tc.expected, tc.nums)
		}
	}
}
