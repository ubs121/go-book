package search

import (
	"fmt"
	"strings"
	"testing"
)

// BST represents the Binary Search Tree.
type BST struct {
	Root *Node
}

// NewBST creates a new empty Binary Search Tree.
func NewBST() *BST {
	return &BST{}
}

// Insert inserts a new key into the BST.
func (bst *BST) Insert(key string) {
	bst.Root = bst.insertRec(bst.Root, key)
}

// insertRec is a helper function to recursively insert a key into the BST.
func (bst *BST) insertRec(root *Node, key string) *Node {
	if root == nil {
		return &Node{Value: key}
	}

	if key < root.Value {
		root.Left = bst.insertRec(root.Left, key)
	} else if key > root.Value {
		root.Right = bst.insertRec(root.Right, key)
	}

	return root
}

// Search searches for a key in the BST.
func (bst *BST) Search(key string) bool {
	return bst.searchRec(bst.Root, key)
}

// searchRec is a helper function to recursively search for a key in the BST.
func (bst *BST) searchRec(root *Node, key string) bool {
	if root == nil {
		return false
	}

	if key == root.Value {
		return true
	} else if key < root.Value {
		return bst.searchRec(root.Left, key)
	} else {
		return bst.searchRec(root.Right, key)
	}
}

// InOrderTraversal performs an in-order traversal of the BST.
func (bst *BST) InOrderTraversal() []string {
	var result []string
	bst.inOrderRec(bst.Root, &result)
	return result
}

// inOrderRec is a helper function for in-order traversal.
func (bst *BST) inOrderRec(node *Node, result *[]string) {
	if node != nil {
		bst.inOrderRec(node.Left, result)
		*result = append(*result, node.Value)
		bst.inOrderRec(node.Right, result)
	}
}

func TestBST(t *testing.T) {
	bst := NewBST()

	// Insert elements into the BST
	bst.Insert("50")
	bst.Insert("30")
	bst.Insert("70")
	bst.Insert("20")
	bst.Insert("40")
	bst.Insert("60")
	bst.Insert("80")

	// Search for a key
	keyToFind := "40"
	if bst.Search(keyToFind) {
		fmt.Printf("%s found in the BST\n", keyToFind)
	} else {
		fmt.Printf("%s not found in the BST\n", keyToFind)
	}

	// In-order traversal of the BST
	fmt.Println("In-order traversal:", bst.InOrderTraversal())
}

type Color int

const (
	RED Color = iota
	BLACK
)

// RedBlackTree node
type RbtNode struct {
	Key         string   // key
	Data        string   // data
	Left, Right *RbtNode // left, right subtrees
	Color       Color    // color of parent link
	Size        int      // subtree count
}

type RedBlackBST struct {
	root *RbtNode
}

// Inserts a key-value pair into the tree
func (rbt *RedBlackBST) Put(key string, data string) {
	rbt.root = rbt.put(rbt.root, key, data)
	rbt.root.Color = BLACK // enforces the root node to be black
}

func (rbt *RedBlackBST) put(n *RbtNode, key string, data string) *RbtNode {
	if n == nil {
		return &RbtNode{Key: key, Data: data, Color: RED, Size: 1}
	}

	cmp := strings.Compare(key, n.Key)
	if cmp < 0 {
		n.Left = rbt.put(n.Left, key, data)
	} else if cmp > 0 {
		n.Right = rbt.put(n.Right, key, data)
	} else {
		n.Data = data
	}

	if rbt.IsRed(n.Right) && !rbt.IsRed(n.Left) {
		n = rbt.RotateLeft(n)
	}
	if rbt.IsRed(n.Left) && rbt.IsRed(n.Left.Left) {
		n = rbt.RotateRight(n)
	}
	if rbt.IsRed(n.Left) && rbt.IsRed(n.Right) {
		rbt.FlipColors(n)
	}

	n.Size = rbt.Size(n.Left) + rbt.Size(n.Right) + 1
	return n
}

// IsRed checks if n is red; false if n is null ?
func (rbt *RedBlackBST) IsRed(n *RbtNode) bool {
	if n == nil {
		return false
	}
	return n.Color == RED
}

// Size returns a number of nodes in subtree rooted at n; 0 if n is null
func (rbt *RedBlackBST) Size(n *RbtNode) int {
	if n == nil {
		return 0
	}
	return rbt.Size(n.Left) + rbt.Size(n.Right) + 1
}

// Get finds a node with a given key
func (rbt *RedBlackBST) Get(key string) *RbtNode {
	x := rbt.root
	for x != nil {
		cmp := strings.Compare(key, x.Key)
		if cmp < 0 {
			x = x.Left
		} else if cmp > 0 {
			x = x.Right
		} else { // cmp == 0
			return x
		}
	}

	return nil
}

// RotateLeft orients a right-leaning red link to lean left
func (rbt *RedBlackBST) RotateLeft(n *RbtNode) *RbtNode {
	x := n.Right
	n.Right = x.Left
	x.Left = n
	x.Color = n.Color
	n.Color = RED
	x.Size = n.Size
	n.Size = 1 + rbt.Size(n.Left) + rbt.Size(n.Right)
	return x
}

// RotateRight orients a left-leaning red link to lean right
func (rbt *RedBlackBST) RotateRight(n *RbtNode) *RbtNode {
	x := n.Left
	n.Left = x.Right
	x.Right = n
	x.Color = n.Color
	n.Color = RED
	n.Size = x.Size
	x.Size = 1 + rbt.Size(x.Left) + rbt.Size(x.Right)
	return x
}

// FlipColors recolor to split 4-node (2 red links)
func (rbt *RedBlackBST) FlipColors(n *RbtNode) {
	n.Color = RED
	n.Left.Color = BLACK
	n.Right.Color = BLACK
}

func TestRedBlackBST(t *testing.T) {
	rbt := &RedBlackBST{}

	// Insert 5 nodes
	rbt.Put("E", "5")
	rbt.Put("A", "1")
	rbt.Put("S", "19")
	rbt.Put("Y", "25")
	rbt.Put("Q", "17")

	// Check node counts
	if rbt.Size(rbt.root) != 5 {
		t.Errorf("Expected size of 5, but got %d", rbt.Size(rbt.root))
	}

	// Check node values
	if rbt.Get("E").Data != "5" {
		t.Errorf("Expected node E to have data 5, but got %s", rbt.Get("E").Data)
	}
	if rbt.Get("A").Data != "1" {
		t.Errorf("Expected node A to have data 1, but got %s", rbt.Get("A").Data)
	}
	if rbt.Get("S").Data != "19" {
		t.Errorf("Expected node S to have data 19, but got %s", rbt.Get("S").Data)
	}
	if rbt.Get("Y").Data != "25" {
		t.Errorf("Expected node Y to have data 25, but got %s", rbt.Get("Y").Data)
	}
	if rbt.Get("Q").Data != "17" {
		t.Errorf("Expected node Q to have data 17, but got %s", rbt.Get("Q").Data)
	}
}

// SplayTree represents the Splay Tree.
type SplayTree struct {
	Root *Node
}

// NewSplayTree creates a new empty Splay Tree.
func NewSplayTree() *SplayTree {
	return &SplayTree{}
}

// Insert inserts a new key into the Splay Tree.
func (st *SplayTree) Insert(key string) {
	st.Root = st.insertRec(st.Root, key)
}

// insertRec is a helper function to recursively insert a key into the Splay Tree.
func (st *SplayTree) insertRec(root *Node, key string) *Node {
	if root == nil {
		return &Node{Value: key}
	}

	if key < root.Value {
		root.Left = st.insertRec(root.Left, key)
		return st.splay(root, key)
	}

	if key > root.Value {
		root.Right = st.insertRec(root.Right, key)
		return st.splay(root, key)
	}

	return root // Duplicate keys are not allowed
}

// Search searches for a key in the Splay Tree.
func (st *SplayTree) Search(key string) *Node {
	st.Root = st.searchRec(st.Root, key)
	return st.Root
}

// searchRec is a helper function to recursively search for a key in the Splay Tree.
func (st *SplayTree) searchRec(root *Node, key string) *Node {
	if root == nil || root.Value == key {
		return root
	}

	if key < root.Value {
		if root.Left == nil {
			return root
		}
		root.Left.Left = st.searchRec(root.Left.Left, key)
		root = st.rotateRight(root)
	} else {
		if root.Right == nil {
			return root
		}
		root.Right.Right = st.searchRec(root.Right.Right, key)
		root = st.rotateLeft(root)
	}

	if root.Left == nil {
		return root
	}

	return st.rotateRight(root)
}

// rotateLeft performs a left rotation on the given node.
func (st *SplayTree) rotateLeft(x *Node) *Node {
	y := x.Right
	x.Right = y.Left
	y.Left = x
	return y
}

// rotateRight performs a right rotation on the given node.
func (st *SplayTree) rotateRight(x *Node) *Node {
	y := x.Left
	x.Left = y.Right
	y.Right = x
	return y
}

// splay splays the node with the given key to the root.
func (st *SplayTree) splay(root *Node, key string) *Node {
	if root == nil || root.Value == key {
		return root
	}

	if key < root.Value {
		if root.Left == nil {
			return root
		}

		if key < root.Left.Value {
			// Zig-Zig (Left-Left)
			root.Left.Left = st.splay(root.Left.Left, key)
			root = st.rotateRight(root)
		} else if key > root.Left.Value {
			// Zig-Zag (Left-Right)
			root.Left.Right = st.splay(root.Left.Right, key)
			if root.Left.Right != nil {
				root.Left = st.rotateLeft(root.Left)
			}
		}

		if root.Left == nil {
			return root
		}

		return st.rotateRight(root)
	} else {
		if root.Right == nil {
			return root
		}

		if key < root.Right.Value {
			// Zag-Zig (Right-Left)
			root.Right.Left = st.splay(root.Right.Left, key)
			if root.Right.Left != nil {
				root.Right = st.rotateRight(root.Right)
			}
		} else if key > root.Right.Value {
			// Zag-Zag (Right-Right)
			root.Right.Right = st.splay(root.Right.Right, key)
			root = st.rotateLeft(root)
		}

		if root.Right == nil {
			return root
		}

		return st.rotateLeft(root)
	}
}

// InOrderTraversal performs an in-order traversal of the Splay Tree.
func (st *SplayTree) InOrderTraversal() []string {
	var result []string
	st.inOrderRec(st.Root, &result)
	return result
}

// inOrderRec is a helper function for in-order traversal.
func (st *SplayTree) inOrderRec(node *Node, result *[]string) {
	if node != nil {
		st.inOrderRec(node.Left, result)
		*result = append(*result, node.Value)
		st.inOrderRec(node.Right, result)
	}
}

func TestSplayTree(t *testing.T) {
	splayTree := NewSplayTree()

	// Insert elements into the Splay Tree
	keys := []string{"apple", "banana", "cherry", "date", "fig", "grape", "kiwi"}
	for _, key := range keys {
		splayTree.Insert(key)
	}

	// Search for a key (and splay it to the root)
	keyToFind := "date"
	foundNode := splayTree.Search(keyToFind)
	if foundNode != nil {
		fmt.Printf("%s found in the Splay Tree and splayed to the root.\n", keyToFind)
	} else {
		fmt.Printf("%s not found in the Splay Tree.\n", keyToFind)
	}

	// Print the keys in in-order traversal (sorted order)
	fmt.Println("In-order traversal:", splayTree.InOrderTraversal())
}
