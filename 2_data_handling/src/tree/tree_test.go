package search

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Value int       // зангилаан дээрх утга
	Left  *TreeNode // зүүн зангилаа
	Right *TreeNode // баруун зангилаа
}

func TestTreeCreate(t *testing.T) {
	root := &TreeNode{
		Value: 5,
		Left: &TreeNode{
			Value: 4,
			Left:  &TreeNode{Value: 3},
			Right: &TreeNode{Value: 3},
		},
		Right: &TreeNode{Value: 2,
			Left:  &TreeNode{Value: 7},
			Right: nil,
		},
	}
	fmt.Println(root)
}

// модны нэг зангилаа
type Node struct {
	Left, Right *Node
	// зүүн, баруун зангилаа
	Value string // зангилаан дээрх утга
}

// мод бүтэц
type BinaryTree struct {
	root *Node
	// оройн элементийг заана
}

func addNode(val string) *Node {
	return &Node{nil, nil, val}
}

func (b *BinaryTree) Insert(val string) (n *Node) {
	if b.root == nil {
		n = addNode(val)
		b.root = n
	} else {
		n = b.insert(b.root, val)
	}
	return
}

func (b *BinaryTree) insert(root *Node, val string) *Node {
	switch {
	case root == nil:
		return addNode(val)
	case val <= root.Value:
		root.Left = b.insert(root.Left, val)
	case val > root.Value:
		root.Right = b.insert(root.Right, val)
	}
	return root
}

func (b *BinaryTree) Print() {
	printTree(b.root)
}

func printTree(n *Node) {
	if n == nil {
		return
	}
	printTree(n.Left)
	fmt.Printf("%s\n", n.Value)
	printTree(n.Right)
}

func find(n *Node, val string) {
	if n == nil {
		return
	}
	if n.Value == val {
		fmt.Printf("%s үг модонд олдлоо!\n", val)
	}
	if val <= n.Value {
		find(n.Left, val)
	} else {
		find(n.Right, val)
	}
}

func (b *BinaryTree) Find(val string) {
	find(b.root, val)
}

func TestTree(t *testing.T) {
	b := new(BinaryTree)
	b.Insert("lemon")
	b.Insert("apple")
	b.Insert("grape")
	b.Insert("orange")
	b.Insert("plum")
	b.Insert("pear")

	b.Print()

	b.Find("pear")
}
