package debug

import (
	"fmt"
	"testing"
)

func TestDivByZero(t *testing.T) {
	i := 1
	j := 0
	fmt.Println("Хуваахын өмнө...")
	i = i / j /* 0-д хуваах алдаа */
	fmt.Println("Дараа нь")
}
