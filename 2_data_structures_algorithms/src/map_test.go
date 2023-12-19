package dsa

import (
	"fmt"
	"maps"
	"testing"
)

func TestMapClone(t *testing.T) {
	colors := map[string]int{
		"red":   0xFF0000,
		"green": 0x00FF00,
		"blue":  0x0000FF,
	}
	colors1 := maps.Clone(colors)
	colors1["red"] = 0

	fmt.Println(colors, colors1)
}
