package graph

import (
	"testing"
)

/*
You've found yourself in a grid of cells with R rows and C columns. The cell in the ith row from the top and jth column from the left contains one of the following (indicated by the character G[i,j]):

* If G[i,j]=".", the cell is empty.
* If G[i,j]="S", the cell contains your starting position. There is exactly one such cell.
* If G[i,j]="E", the cell contains an exit. There is at least one such cell.
* If G[i,j]="#", the cell contains a wall.
* Otherwise, if G[i,j] is a lowercase letter (between "a" and "z", inclusive), the cell contains a portal marked with that letter.

Your objective is to reach any exit from your starting position as quickly as possible. Each second, you may take either of the following actions:
* Walk to a cell adjacent to your current one (directly above, below, to the left, or to the right), as long as you remain within the grid and that cell does not contain a wall.
* If your current cell contains a portal, teleport to any other cell in the grid containing a portal marked with the same letter as your current cell's portal.

Determine the minimum number of seconds required to reach any exit, if it's possible to do so at all. If it's not possible, return −1 instead.
*/

func minimumSecondsToExit(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])

	// Find the starting position and collect portal locations.
	var start Point
	portals := make(map[byte][]Point)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 'S' {
				start = Point{i, j}
			} else if grid[i][j] >= 'a' && grid[i][j] <= 'z' {
				portals[grid[i][j]] = append(portals[grid[i][j]], Point{i, j})
			}
		}
	}

	// Create a queue for BFS.
	var queue []Point
	queue = append(queue, start)
	visited := make(map[Point]bool)
	visited[start] = true
	seconds := 0

	// Possible directions to move (up, down, left, right).
	directions := [4]Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// Perform BFS to find the exit. BFS explores nodes level by level.
	for len(queue) > 0 {
		size := len(queue)
		seconds++
		//fmt.Printf("queue=%v\n", queue)

		for i := 0; i < size; i++ {
			current := queue[0]
			queue = queue[1:]

			// If the current cell is an exit, return the number of seconds
			if grid[current.x][current.y] == 'E' {
				return seconds
			}

			// teleport to all possible portal destinations.
			cellVal := grid[current.x][current.y]
			if 'a' <= cellVal && cellVal <= 'z' {
				portalDestinations := portals[cellVal]
				for _, dest := range portalDestinations {
					if !visited[dest] {
						queue = append(queue, dest)
						visited[dest] = true
					}
				}
			} else {
				// check all direction from 'current'
				for _, d := range directions {
					x1, y1 := current.x+d.x, current.y+d.y

					if x1 >= 0 && x1 < rows && y1 >= 0 && y1 < cols && grid[x1][y1] != '#' {
						next := Point{x1, y1}
						if !visited[next] && !('a' <= grid[x1][y1] && grid[x1][y1] <= 'z') {
							queue = append(queue, next)
							visited[next] = true
						}
					}
				}
			}
		}

	}

	// If no exit is found, return -1.
	return -1
}

type Point struct {
	x, y int
}

func TestEscapeGrid(t *testing.T) {
	testCases := []struct {
		grid [][]byte
		exp  int
	}{
		{
			grid: [][]byte{
				[]byte(".E."),
				[]byte(".#E"),
				[]byte(".S#"),
			},
			exp: 4, // walk left once, then up twice, and then finally right once.
		},
		{
			grid: [][]byte{
				[]byte("a.Sa"),
				[]byte("####"),
				[]byte("Eb.b"),
			},
			exp: -1, // you can never reach the exit
		},
		{
			grid: [][]byte{
				[]byte("aS.b"),
				[]byte("####"),
				[]byte("Eb.a"),
			},
			exp: 4, // walk right twice, then teleport to the cell in the 3rd row and 2nd column, and finally walk left once.
		},
		{
			grid: [][]byte{
				[]byte("xS..x..Ex"),
			},
			exp: 3, // walk left once, teleport to the cell in the last column, and walk left once more
		},
	}
	for i, tc := range testCases {
		got := minimumSecondsToExit(tc.grid)
		if tc.exp != got {
			t.Errorf("tc %d: exp %d, got %d", i, tc.exp, got)
		}
	}
}
