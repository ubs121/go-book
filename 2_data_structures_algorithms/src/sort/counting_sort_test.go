package sort

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

// Counting sort is not a comparison sort algorithm, O(n)
// assumes each element is in range [1..maxVal)
func CountingSort(arr []int, maxVal int) {
	n := len(arr) // number of elements

	// count each elements
	countArr := make([]int, maxVal)
	for i := 0; i < n; i++ {
		countArr[arr[i]]++
	}

	// do prefix sum for each element
	for i := 1; i < maxVal; i++ {
		countArr[i] += countArr[i-1]
	}

	// place each elements in order
	sortedArr := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		sortedArr[countArr[arr[i]]-1] = arr[i]
		countArr[arr[i]]--
	}

	// copy back to 'arr'
	copy(arr, sortedArr)
}

func TestCountingSort(t *testing.T) {

	// test input
	arr := []int{4, 5, 1, 2, 4, 1, 3}

	// expected result
	arrExp := make([]int, len(arr))
	copy(arrExp, arr)
	sort.Ints(arrExp) // sort using standard func

	/* массивыг эрэмбэлэх */
	maxVal := arrExp[len(arrExp)-1] + 1
	CountingSort(arr, maxVal)

	if !reflect.DeepEqual(arr, arrExp) {
		t.Error("not equal")
	}

	/* эрэмбэлэгдсэн  массивыг хэвлэж харуулах */
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
}
