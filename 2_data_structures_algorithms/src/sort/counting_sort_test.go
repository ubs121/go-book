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
	n := len(arr) // элементийн тоо

	// элемент бүрийн тоог гаргах
	countArr := make([]int, maxVal)
	for _, x := range arr {
		countArr[x]++
	}

	// өмнөх элементүүдийн тоог нэмэх
	for i := 1; i < maxVal; i++ {
		countArr[i] += countArr[i-1]
	}

	// элементүүдийг эрэмбээр нь байрлуулах
	sortedArr := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		sortedArr[countArr[arr[i]]-1] = arr[i]
		countArr[arr[i]]--
	}

	// буцаан хуулах
	copy(arr, sortedArr)
}

func TestCountingSort(t *testing.T) {

	// test input
	arr := []int{4, 5, 1, 2, 4, 1, 3}

	// expected result
	arrExp := make([]int, len(arr))
	copy(arrExp, arr)
	sort.Ints(arrExp) // sort using standard function

	/* слайсыг эрэмбэлэх */
	maxVal := arrExp[len(arrExp)-1] + 1
	CountingSort(arr, maxVal)

	if !reflect.DeepEqual(arr, arrExp) {
		t.Error("not equal")
	}

	/* эрэмбэлэгдсэн слайсыг хэвлэж харуулах */
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
}

// Sort Colors: Given an array with n objects colored red, white, or blue,
// sort them in-place so that objects of the same color are adjacent,
// with the colors in the order red, white, and blue.
func sortColors(nums []int) {
	count := make([]int, 3) // Count the occurrences of each color

	// Count the number of occurrences of each color
	for _, color := range nums {
		count[color]++
	}

	index := 0

	// Fill the nums array with the sorted colors
	for color := 0; color <= 2; color++ {
		for count[color] > 0 {
			nums[index] = color
			index++
			count[color]-- // descrease the counter
		}
	}
}

func TestSortColors(t *testing.T) {
	colors := []int{0, 2, 2, 1, 2, 0, 1}
	sortColors(colors)
	fmt.Printf("colors=%v", colors)
}
