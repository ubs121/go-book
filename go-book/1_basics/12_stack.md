# Стек

Стек нь “эхэлж орсон нь эхэлж гарах” зарчим дээр тулгуурласан жагсаалт төрлийн өгөгдлийн бүтэц юм. Стек бүтэц дээр дараах хоёр үйлдэл хийгдэнэ:

* Push – стекийн орой дээр элемент нэмэх
* Pop – стекийн оройноос элемент сугалах

Дараах жишээнд тоон стек хэрхэн үүсгэж, ашиглахыг харуулав.

```go
package main
import "fmt"

type IntStack struct {
	elems []int
}

// стекийн урт
func (s *IntStack) Len() int {
	return len(s.elems)
}

// стекийн оройд элемент нэмэх
func (s *IntStack) Push(value int) {
	s.elems = append(s.elems, value)
}

// стекийн оройгоос элемент авах
func (s *IntStack) Pop() int {
	if len(s.elems) == 0 {
		panic("stack is empty")
	}

	top := len(s.elems) - 1
	value := s.elems[top]
	s.elems = s.elems[:top]
	return value
}

func main() {
	stack := new(IntStack)

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	for stack.Len() > 0 {
		fmt.Printf("%d ", stack.Pop())
	}
}
```
