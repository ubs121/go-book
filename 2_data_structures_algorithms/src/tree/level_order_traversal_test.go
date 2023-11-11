package search

import (
	"fmt"
	"testing"
)

func LevelOrderTraversal(root *Node) [][]string { // Data from each level is being returned as a separate list
	if root == nil {
		return nil
	}

	var result [][]string
	queue := []*Node{root}
	for len(queue) > 0 {
		qlen := len(queue)

		var levelValues []string
		for i := 0; i < qlen; i++ {
			node := queue[0]
			levelValues = append(levelValues, node.Value)
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, levelValues)
	}
	return result
}

func TestLevelOrderTraversal(t *testing.T) {
	root := &Node{
		Value: "1",
		Left: &Node{
			Value: "2",
			Left:  &Node{Value: "4"},
			Right: &Node{Value: "5"},
		},
		Right: &Node{
			Value: "3",
			Left:  &Node{Value: "6"},
			Right: &Node{Value: "7"},
		},
	}

	fmt.Println(LevelOrderTraversal(root))
}

type msg struct {
	elem int
	lvl  int
}

// tree traversal by levels
func concurrentTraverseByLevels(root []any) {
	out := make(chan *msg)

	go func() {
		defer close(out)
		_traverse(root, 0, out)
	}()

	lvls := map[int][]int{}

	for {
		m := <-out
		if m == nil {
			break
		}

		lvls[m.lvl] = append(lvls[m.lvl], m.elem)
	}

	fmt.Printf("Answer %v", lvls)
}

func _traverse(root []any, lvl int, out chan *msg) {

	for _, e := range root {
		switch v := e.(type) {
		case int:
			out <- &msg{v, lvl}
		case []any:
			_traverse(v, lvl+1, out)
		default:
			// interface, so go deeper???
		}
	}
}

func TestTraverseLevels(t *testing.T) {
	root := []any{5, 12, []any{1, []any{8, 10, 2}, 1, 100}, []any{15}, 7}
	concurrentTraverseByLevels(root)
}
