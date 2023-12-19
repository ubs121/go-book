package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	time.Sleep(3 * time.Second)
}

func TestGoroutine(t *testing.T) {
	go f(10)

	var input string
	fmt.Scanln(&input)
}
