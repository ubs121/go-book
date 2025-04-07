## Эрэмбэлэлт

Эрэмбэлэлт нь компютерын шинжлэх ухаанд чухал алгоритмын бүлэг юм. Элементүүдийг тодорхой эрэмбэд оруулснаар олон асуудлыг хялбараар шийдэж болно.

Энгийн слайсыг эрэмбэлэхэд `sort.Ints()`, `sort.String()` эсвэл ерөнхий `slices.Sort()` функцыг ашиглаж болно.

```go
/* бүхэл тоон слайсыг эрэмбэлэх */
numbers:=[]int{5,2,6,3,1,4}

sort.Ints(numbers) // эсвэл slices.Sort(numbers)

/* тэмдэг мөрүүдийг эрэмбэлэх */
colors := []string{"улаан", "ногоон", "цэнхэр"}
slices.Sort(colors)
```

Бүтэц төрлийн слайс бол `sort.Slice()` функцийг хэрэглэж болно.

```go
type Person struct {
    Name string
    Age int
}

// ...
kids := []Person{
    {"Амар", 10},
    {"Болд", 9},
    {"Батаа", 9},
}

// kids слайсыг эрэмбэлж байна
sort.Slice(kids, func(i, j int) bool {
    // нас нь зөрүүтэй бол насаар нь эрэмбэлэх
    if kids[i].Age != kids[j].Age {
        return kids[i].Age < kids[j].Age
    }
    // эсрэг тохиолдод нэрээр нь эрэмбэлэх
    return kids[i].Name < kids[j].Name
})
// Гаралт: [{Батаа 9} {Болд 9} {Амар 10}]
```

Мөн шинэ эрэмбэлэгч төрөл үүсгэх замаар бүтэц өгөгдлийг эрэмбэлж болно. Эрэмбэлэгч төрөлд хэрхэн харьцуулах, эрэмбэлэх тухай `Len`, `Less`, `Swap` гэсэн методуудад тодорхойлж өгдөг.

Жишээ болгон `ByName` нэртэй эрэмбэлэгч төрөл үүсгэе. Дараах `ByName` эрэмбэлэгч нь `Person` төрлийн утгуудыг нэрээр (`Name` талбараар) нь харьцуулж эрэмбэлнэ.

```go
type ByName []Person

func (p ByName) Len() int {
    return len(p)
}
func (p ByName) Less(i, j int) bool {
    return p[i].Name < p[j].Name
}
func (p ByName) Swap(i, j int) {
    p[i], p[j] = p[j], p[i]
}
```

Ингээд `sort.Sort()` функцийг `ByName` эрэмбэлэгчтэй хослуулан дараах байдлаар ашиглаж болно.

```go
// ByName эрэмбэлэгчийг ашиглан эрэмбэлэх
sort.Sort(ByName(kids)) 
// Гаралт: [{Амар 10} {Батаа 9} {Болд 9}]
```

## Heap эрэмбэлэлт

```go
// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func TestIntHeap(t *testing.T) {
	// initial values
	h := &IntHeap{2, 1, 5}
	heap.Init(h)

	// push 3
	heap.Push(h, 3)

	fmt.Printf("minimum: %d\n", (*h)[0])

	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h)) // h.Pop() gives different result !!!
	}
}
```

## Тоолж эрэмбэлэх

Тоолж эрэмбэлэх арга нь харьцуулалт хийхгүйгээр эрэмбэлдэг. Үндсэн санаа нь элемент бүрийн өмнө хэдэн элемент байгааг тоолох юм, хэрэв энэ тоо тодорхой бол элементийн  байрлал тодорхой болно. Жишээлбэл `x` элементээс бага 10 элемент байгаа бол `x` нь 11 дүгээр байрлал дээр байна гэсэн үг. Эрэмбэлэлт зөв ажиллахын тулд `x` нь [0..maxVal] интервалд байх хэрэгтэй.

```go
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
```