package basics

import (
	"fmt"
	"testing"
)

func TestSscanf(t *testing.T) {
	s := "10 20 30"
	a := 0
	b := 0
	c := 0
	fmt.Sscanf(s, "%d", &a, &b, &c)
	println(a, b, c)
}
