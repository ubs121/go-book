package sort

import (
	"fmt"
	"sort"
	"testing"
)

func TestSortSlice(t *testing.T) {
	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}

	// sort 'people' in ascending order by name
	sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println("By name:", people)

	// sort 'people' in ascending order by age
	sort.Slice(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	fmt.Println("By age:", people)
}

type Person struct {
	Name string
	Age  int
}

type ByName []Person

func (this ByName) Len() int {
	return len(this)
}
func (this ByName) Less(i, j int) bool {
	return this[i].Name < this[j].Name
}
func (this ByName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func TestSortPerson(t *testing.T) {
	kids := []Person{
		{"Бат", 9},
		{"Болд", 9},
		{"Амар", 10},
	}
	sort.Sort(ByName(kids))
	fmt.Println(kids)
}

func (p *Person) Talk() {
	fmt.Println("Сайна уу, Миний нэр ", p.Name)
}

// QuickSort - range [low..high], O(nlogn), worst O(n^2)
func QuickSort(arr []int, low int, high int) {
	i := low
	j := high
	pivot := arr[(low+high)/2]
	var temp int

	for i <= j {
		for arr[i] < pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}

		if i <= j {
			temp = arr[i]
			arr[i] = arr[j]
			arr[j] = temp
			i++
			j--
		}
	}

	if low < j {
		QuickSort(arr, low, j)
	} /* зүүн хэсгийг эрэмбэлэх */

	if i < high {
		QuickSort(arr, i, high)
	} /* баруун хэсгийг эрэмбэлэх */
}
