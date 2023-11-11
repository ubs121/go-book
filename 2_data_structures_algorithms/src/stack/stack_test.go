package common

import (
	"fmt"
	"reflect"
	"testing"
)

type Stack[T any] struct {
	elems []T
}

// стекийн урт
func (s *Stack[T]) Len() int {
	return len(s.elems)
}

// стекийн оройд элемент нэмэх
func (s *Stack[T]) Push(value T) {
	s.elems = append(s.elems, value)
}

// стекийн оройгоос элемент авах
func (s *Stack[T]) Pop() T {
	if len(s.elems) == 0 {
		panic("stack is empty")
	}

	topIdx := len(s.elems) - 1
	value := s.elems[topIdx]
	s.elems = s.elems[:topIdx]
	return value
}

func TestStack(t *testing.T) {
	stack := new(Stack[int])

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	for stack.Len() > 0 {
		fmt.Printf("%d ", stack.Pop())
	}
}

// Maximum number of consecutive elements preceding arr[i] such that a[j]<a[i], j<i
func FindSpans(arr []int) []int {
	var stack []int
	p := -1 // index of the closest greater element

	span := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {

		// find all lesser elements
		j := len(stack) - 1
		for len(stack) > 0 && arr[stack[j]] < arr[i] {
			j--
		}
		stack = stack[:j+1] // pop all lesser elements

		if len(stack) == 0 {
			p = -1
		} else {
			p = stack[len(stack)-1]
		}
		span[i] = i - p
		stack = append(stack, i)
	}
	return span
}

func TestFindSpans(t *testing.T) {
	testCases := []struct {
		arr []int
		exp []int
	}{
		{[]int{6, 3, 4, 5, 2}, []int{1, 1, 2, 3, 1}},
	}
	for i, tc := range testCases {
		got := FindSpans(tc.arr)
		if !reflect.DeepEqual(tc.exp, got) {
			t.Errorf("tc %d: exp %v, got %v", i, tc.exp, got)
		}
	}
}

// https://leetcode.com/problems/min-stack/
type MinStack struct {
	elems [][]int // (elem, min) pairs
	n     int     // number of elements
}

/** initialize your data structure here. */
func NewMinStack() *MinStack {
	return &MinStack{}
}

func (st *MinStack) Push(val int) {
	if st.n == 0 {
		st.elems = append(st.elems, []int{val, val})
	} else {
		min := st.GetMin()
		if min > val {
			min = val
		}
		st.elems = append(st.elems, []int{val, min})
	}
	st.n++
}

func (st *MinStack) Pop() {
	// always be called on non-empty stacks.
	st.n--
	st.elems = st.elems[:st.n]
}

func (st *MinStack) Top() int {
	// always be called on non-empty stacks.
	return st.elems[st.n-1][0]
}

func (st *MinStack) GetMin() int {
	// always be called on non-empty stacks.
	return st.elems[st.n-1][1]
}
