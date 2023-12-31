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
	Reverse(intArr)
	fmt.Println(intArr)

	strArr := []string{"red", "green", "blue"}
	Reverse(strArr)
	fmt.Println(strArr)
}
