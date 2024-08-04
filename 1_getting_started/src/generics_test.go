package common

import (
	"fmt"
	"testing"
)

func Reverse[T any](arr []T) {
	i, j := 0, len(arr)-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

func TestReverse(t *testing.T) {
	intArr := []int{1, 2, 3, 4, 5, 6}
	strArr := []string{"red", "green", "blue"}

	Reverse(intArr)
	Reverse(strArr)

	fmt.Println(intArr)
	fmt.Println(strArr)
}
