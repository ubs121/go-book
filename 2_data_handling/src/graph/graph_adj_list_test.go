package graph

import (
	"fmt"
	"testing"
)

// Iterative Depth-First Search (DFS)
// g is a graph represented in adjacent list
func DFS_AdjList(nodes []any, g [][]int, source int) {
	visited := make([]bool, len(g))
	stack := []int{source} // stack

	for len(stack) > 0 {
		u := stack[len(stack)-1]     // top element
		stack = stack[:len(stack)-1] // cut the stack

		if visited[u] {
			continue
		}

		fmt.Printf("at %v\n", nodes[u])
		visited[u] = true

		// add neighbors of 'u'
		stack = append(stack, g[u]...)
	}
}

func BFS_AdjList(nodes []any, g [][]int, source int) {
	visited := make([]bool, len(g))
	q := []int{source} // queue

	for len(q) > 0 {
		u := q[0] // get first
		q = q[1:] // cut the queue

		fmt.Printf("at %v\n", nodes[u])
		visited[u] = true

		// add neighbors of 'u'
		for _, v := range g[u] {
			if !visited[v] {
				q = append(q, v)
			}
		}
	}
}

func TestDFS_AdjList(t *testing.T) {
	nodes := []any{"A", "B", "C", "G", "H"}              // nodes
	g := [][]int{{1, 2, 3, 4}, {0}, {0, 3}, {0, 2}, {0}} // relations graph

	DFS_AdjList(nodes, g, 0)
}

func TestBFS_AdjList(t *testing.T) {
	nodes := []any{"n5", "n4", "n2", "n3", "n1", "n7"}     // nodes
	g := [][]int{{2, 1}, {4, 3, 0}, {5, 0}, {1}, {1}, {2}} // relations

	BFS_AdjList(nodes, g, 0)
}
