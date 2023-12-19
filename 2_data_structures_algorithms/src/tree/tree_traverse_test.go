package search

import (
	"fmt"
	"testing"
)

// DLR: Pre-order traversal is to visit the root first. Then traverse the left subtree. Finally, traverse the right subtree
func PreorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	// values from pre-order traversal
	var ret []int
	visit := func(e *TreeNode) {
		ret = append(ret, e.Value)
	}

	var stack []*TreeNode       // processing stack
	stack = append(stack, root) // push root

	for len(stack) > 0 {
		// pop
		e := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		visit(e)

		// push right first so that left processed first
		if e.Right != nil {
			stack = append(stack, e.Right)
		}

		// push left
		if e.Left != nil {
			stack = append(stack, e.Left)
		}
	}

	return ret
}

// LDR: In-order traversal is to traverse the left subtree first. Then visit the root. Finally, traverse the right subtree.
func InorderTraversal(root *TreeNode) []int {
	// values from in-order traversal
	var ret []int
	visit := func(e *TreeNode) {
		ret = append(ret, e.Value)
	}

	var stack []*TreeNode // processing stack

	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1] // pop
			stack = stack[:len(stack)-1]

			visit(root)

			root = root.Right
		}
	}

	return ret
}

// LRD: Post-order traversal is to traverse the left subtree first. Then traverse the right subtree. Finally, visit the root.
func PostorderTraversal(root *TreeNode) []int {
	// values from post-order traversal
	var ret []int
	visit := func(e *TreeNode) {
		ret = append(ret, e.Value)
	}

	var stack []*TreeNode  // processing stack
	var lastNode *TreeNode // last node pointer

	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			topNode := stack[len(stack)-1] // peek

			if topNode.Right != nil && lastNode != topNode.Right {
				root = topNode.Right
			} else {
				visit(topNode)

				// pop
				lastNode = topNode
				stack = stack[:len(stack)-1]
			}
		}
	}

	return ret
}

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

	result := LevelOrderTraversal(root)
	for level, nodes := range result {
		fmt.Printf("Level %d, Nodes %v\n", level, nodes)
	}
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

// A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
func MaximumDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := MaximumDepth(root.Left)
	rightDepth := MaximumDepth(root.Right) // can be called concurrently if it's a large tree

	max := leftDepth
	if max < rightDepth {
		max = rightDepth
	}
	return max + 1
}

func hasPathSum(root *TreeNode, targetSum int) bool {

	return pathSum(root, 0, targetSum)
}

func pathSum(root *TreeNode, sum, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		if root.Value+sum == targetSum {
			return true
		}
	}

	return pathSum(root.Left, root.Value+sum, targetSum) ||
		pathSum(root.Right, root.Value+sum, targetSum)
}

func TestHasPathSum(t *testing.T) {
	root := &TreeNode{
		Value: 1,
		Left: &TreeNode{
			Value: -2,
			Left:  &TreeNode{Value: 1},
			Right: &TreeNode{Value: 3},
		},
		Right: &TreeNode{
			Value: -3,
			Left:  &TreeNode{Value: -1},
			Right: &TreeNode{Value: -2},
		},
	}

	ans := hasPathSum(root, 3)
	if ans != false {
		t.Errorf("exp: false, got : true")
	}
}

// https://leetcode.com/problems/even-odd-tree/
// A binary tree is named Even-Odd if it meets the following conditions:
//
// The root of the binary tree is at level index 0, its children are at level index 1, their children are at level index 2, etc.
// For every even-indexed level, all nodes at the level have odd integer values in strictly increasing order (from left to right).
// For every odd-indexed level, all nodes at the level have even integer values in strictly decreasing order (from left to right).
// Given the root of a binary tree, return true if the binary tree is Even-Odd, otherwise return false.
func isEvenOddTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var q []TreeNodeLvl
	q = append(q, TreeNodeLvl{*root, 0}) // push the root

	lastElemAtLvl := map[int]int{} // level => last value, can be an array

	for len(q) > 0 {
		node := q[0] // top element

		// validate the even/odd requirement
		if node.Level%2 == node.Value%2 {
			return false // failed
		}

		// validate the order requirement
		if lastVal, exists := lastElemAtLvl[node.Level]; exists {
			if node.Level%2 == 0 {
				if lastVal >= node.Value {
					return false // failed
				}
			} else {
				if lastVal <= node.Value {
					return false // failed
				}
			}
		}

		if node.Left != nil {
			q = append(q, TreeNodeLvl{*node.Left, node.Level + 1})
		}
		if node.Right != nil {
			q = append(q, TreeNodeLvl{*node.Right, node.Level + 1})
		}

		lastElemAtLvl[node.Level] = node.Value
		q = q[1:] // remove 'node'

	}
	return true
}

type TreeNodeLvl struct {
	TreeNode
	Level int
}

func TestIfValidEvenOdd(t *testing.T) {
	testCases := []struct {
		root *TreeNode
		exp  bool
	}{
		{
			&TreeNode{1,
				&TreeNode{10,
					&TreeNode{3,
						&TreeNode{12, nil, nil},
						&TreeNode{8, nil, nil},
					},
					nil,
				},
				&TreeNode{4,
					&TreeNode{7,
						&TreeNode{6, nil, nil},
						nil,
					},
					&TreeNode{9,
						nil,
						&TreeNode{2, nil, nil},
					},
				},
			},
			true,
		},
		{
			&TreeNode{5,
				&TreeNode{4,
					&TreeNode{3, nil, nil},
					&TreeNode{3, nil, nil}},
				&TreeNode{2,
					&TreeNode{7, nil, nil},
					nil,
				},
			},
			false,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := isEvenOddTree(tc.root)
			if got != tc.exp {
				t.Errorf("exp %v, got %v", tc.exp, got)
			}
		})
	}

}
