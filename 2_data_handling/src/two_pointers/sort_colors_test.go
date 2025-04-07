package main

import (
	"slices"
	"testing"
)

/*
0 - red
1 - white
2 - blue
*/
// sort colors using 3-way partitioning with 2 pointers
// partitions will look like [red|white|blue]
func sortColors(colors []int) []int {
	redPos := 0            // swap position for red
	bluePos := len(colors) // swap position for blue

	i := 0
	for i < bluePos {
		if colors[i] == 0 {
			//  move to red section
			colors[i], colors[redPos] = colors[redPos], colors[i]
			redPos++
			i++
		} else if colors[i] == 1 {
			// leave it
			i++
		} else { // colors[i] == 2
			//  move to blue section
			bluePos--
			colors[i], colors[bluePos] = colors[bluePos], colors[i]
		}

		//fmt.Println(i, colors[:redPos], colors[redPos:bluePos], colors[bluePos:])
	}

	return colors
}

func TestSortColors(t *testing.T) {
	input := []int{0, 1, 0, 2, 0, 0, 2, 1, 2}
	exp := []int{0, 0, 0, 0, 1, 1, 2, 2, 2}

	got := sortColors(input)
	if slices.Equal(exp, got) == false {
		t.Errorf("exp %v, got %v", exp, got)
	}
}
