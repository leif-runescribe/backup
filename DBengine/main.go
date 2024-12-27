package main

import (
	"fmt"
)

// BTreeNode represents a single node in a BTree
type BTreeNode struct {
	keys     []int        // Keys in the node
	children []*BTreeNode // Child pointers
	isLeaf   bool         // True if leaf node
	t        int          // Minimum degree
}

// BTree represents the overall BTree
type BTree struct {
	root *BTreeNode
	t    int // Minimum degree
}

// NewBTree creates a new BTree with a given minimum degree
func NewBTree(t int) *BTree {
	return &BTree{
		root: nil,
		t:    t,
	}
}

// traverse prints all the keys in the subtree rooted at this node
func (node *BTreeNode) traverse() {
	for i := 0; i < len(node.keys); i++ {
		// If this is not a leaf, print the subtree rooted at the child before this key
		if !node.isLeaf {
			node.children[i].traverse()
		}
		// Print the key
		fmt.Printf("%d ", node.keys[i])
	}

	// Print the subtree rooted at the last child
	if !node.isLeaf {
		node.children[len(node.keys)].traverse()
	}
}

// search finds the key in the subtree rooted at this node
func (node *BTreeNode) search(k int) *BTreeNode {
	i := 0
	for i < len(node.keys) && k > node.keys[i] {
		i++
	}

	// If the key is found, return this node
	if i < len(node.keys) && node.keys[i] == k {
		return node
	}

	// If the key is not found and this is a leaf, the key doesn't exist
	if node.isLeaf {
		return nil
	}

	// Search in the appropriate child
	return node.children[i].search(k)
}

// insert inserts a new key into the BTree
func (tree *BTree) insert(k int) {
	if tree.root == nil {
		// Tree is empty; allocate a new root
		tree.root = &BTreeNode{
			keys:     []int{k},
			children: nil,
			isLeaf:   true,
			t:        tree.t,
		}
	} else {
		// If the root is full, grow the tree in height
		if len(tree.root.keys) == 2*tree.t-1 {
			newRoot := &BTreeNode{
				keys:     []int{},
				children: []*BTreeNode{tree.root},
				isLeaf:   false,
				t:        tree.t,
			}
			newRoot.splitChild(0, tree.root)
			tree.root = newRoot
		}
		// Insert non-full
		tree.root.insertNonFull(k)
	}
}

// insertNonFull inserts a key into a node that is guaranteed not to be full
func (node *BTreeNode) insertNonFull(k int) {
	i := len(node.keys) - 1

	if node.isLeaf {
		// Find the location to insert the new key and shift others
		node.keys = append(node.keys, 0)
		for i >= 0 && k < node.keys[i] {
			node.keys[i+1] = node.keys[i]
			i--
		}
		node.keys[i+1] = k
	} else {
		// Find the child to recurse into
		for i >= 0 && k < node.keys[i] {
			i--
		}
		i++

		// If the child is full, split it
		if len(node.children[i].keys) == 2*node.t-1 {
			node.splitChild(i, node.children[i])
			// If the key is greater than the middle key, go to the next child
			if k > node.keys[i] {
				i++
			}
		}
		node.children[i].insertNonFull(k)
	}
}

// splitChild splits the child y at index i
func (node *BTreeNode) splitChild(i int, y *BTreeNode) {
	t := y.t
	z := &BTreeNode{
		keys:     append([]int{}, y.keys[t:]...),
		children: nil,
		isLeaf:   y.isLeaf,
		t:        t,
	}
	if !y.isLeaf {
		z.children = append([]*BTreeNode{}, y.children[t:]...)
	}
	y.keys = y.keys[:t-1]
	y.children = y.children[:t]

	// Insert the median key into the parent node
	node.keys = append(node.keys[:i], append([]int{y.keys[t-1]}, node.keys[i:]...)...)
	node.children = append(node.children[:i+1], append([]*BTreeNode{z}, node.children[i+1:]...)...)
}

// Search searches for a key in the BTree
func (tree *BTree) Search(k int) *BTreeNode {
	if tree.root == nil {
		return nil
	}
	return tree.root.search(k)
}

// Traverse traverses the entire tree and prints the keys
func (tree *BTree) Traverse() {
	if tree.root != nil {
		tree.root.traverse()
		fmt.Println()
	}
}

func main() {
	// Test the BTree
	t := 3 // Minimum degree
	tree := NewBTree(t)

	keys := []int{10, 20, 5, 6, 12, 30, 7, 17}
	for _, k := range keys {
		tree.insert(k)
	}

	fmt.Println("Traverse the constructed tree:")
	tree.Traverse()

	searchKey := 12
	fmt.Printf("Search for key %d: ", searchKey)
	if tree.Search(searchKey) != nil {
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found")
	}
}
