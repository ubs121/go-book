## Эрэмбэлэлт

`sort` пакет нь төрөл бүрийн өгөгдлийн олонлогийг эрэмбэлэхэд зориулагдсан байдаг. Тоон эсвэл текст олонлогийг эрэмбэлэхэд `sort.Ints()` эсвэл `sort.Strings()` функцыг шууд ашиглаж болно.

TODO: sort.Slice\(\) функцийн тухай нэмэлт тайлбар оруулах

```go
sort.Slice(s, func(i, j int) bool {
    if s[i].Foo != s[j].Foo {
        return s[i].Foo < s[j].Foo
    }
    return s[i].Bar < s[j].Bar
})
```

Шинэ эрэмбэлэгч төрөл үүсгэх замаар мөн өгөгдлийг эрэмбэлж болно. Жишээ болгон `Person` төрлийн олонлогыг эрэмбэлэх `ByName` нэртэй эрэмбэлэгч төрөл үүсгэе.

```go
type Person struct {
    Name string
    Age int
}
```

Эрэмбэлэгч төрөлд хэрхэн харьцуулах, эрэмбэлэх тухай `Len`, `Less`, `Swap` гэсэн методуудад тодорхойлж өгдөг. Дараах `ByName` эрэмбэлэгч нь `Person` төрлийн утгуудыг нэрээр (`Name` талбараар) нь харьцуулж байна.

```go
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
```

Ингээд `sort.Sort()` функцийг `ByName` эрэмбэлэгчтэй хослуулан дараах байдлаар ашиглаж болно.

```go
func main() {
    kids := []Person{
        {"Бат",9},
        {"Болд",10},
        {"Амар",9},
    }
    sort.Sort(ByName(kids)) // ByName эрэмбэлэгчийг ашиглаж байна
    fmt.Println(kids)
}
```

## Хоёртын хайлт

```go
import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 55, 77}
	x := 10 // хайх утга

	// x утгатай хамгийн ойр байрлалыг олох
	i := sort.Search(len(arr), func(i int) bool { return arr[i] >= x })
	if i < len(arr) && arr[i] == x {
		fmt.Printf("%d нь %d байрлалд олдлоо", x, i)
	} else {
		fmt.Printf("%d олдсонгүй, гэхдээ %d байрлалд оруулж болно", x, i)
	}
}
```


## Мод бүтэц

Өгөгдлийг мод бүтцээр зохион байгуулах нь хайлтад маш тохиромжтой, хурдан байдаг. Жишээ нь үгсийг цагаан толгойн дарааллаар эрэмбэлээд, түүн дотроос үг хайх хэрэгтэй байна гэж бодоё. Үүнд жагсаалтыг ашиглаж болох боловч жагсаалтаас элемент хайхад эхнээс нь дараалан харьцуулалт хийж шалгадаг учир удаан байдаг. Мод бүтцийн тусламжтайгаар харьцуулах үйлдлийг их хэмжээгээр бууруулж болно.


TODO: Модны дүрслэлүүд - UnionFind, Binary indexed tree, Heap

## Heap

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