package debug

import (
	"fmt"
	"testing"
)

// 2 бүхэл тоон хувьсагчийн утгыг хооронд нь солино.
func swap(p1 *int, p2 *int) {
	tmp := p1
	p1 = p2
	p2 = tmp
}

func TestSwap(t *testing.T) {
	a := 10
	b := 20

	fmt.Printf("Анхны утгууд: a = %d; b = %d\n", a, b)

	swap(&a, &b)

	fmt.Printf("Шинэ утгууд: a = %d; b = %d\n", a, b)
}
