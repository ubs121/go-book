## Хоёртын хайлт

```go
import (
	"fmt"
	"slices"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 55, 77}
	x := 4
	
	i, found := slices.BinarySearch(arr, 4)
	fmt.Printf("%d found=%v, loc=%d\n", x, found, i)
}
```
