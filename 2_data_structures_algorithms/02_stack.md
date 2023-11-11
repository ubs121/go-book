## Стек

Стек нь “эхэлж орсон нь эхэлж гарах” зарчим дээр тулгуурласан жагсаалт төрлийн өгөгдлийн бүтэц юм. Стек бүтэц дээр дараах хоёр үйлдэл хийгдэнэ:

* Push – стекийн орой дээр элемент нэмэх
* Pop – стекийн оройноос элемент сугалах

Дараах жишээнд тоон стек хэрхэн үүсгэж, ашиглахыг харуулав.

```go
package main
import "fmt"

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

func main() {
	stack := new(Stack[int])

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	for stack.Len() > 0 {
		fmt.Printf("%d ", stack.Pop())
	}
}
```