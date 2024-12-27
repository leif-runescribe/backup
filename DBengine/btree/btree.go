package btree

import (
	"embedded-db/storage"
	"errors"
)

const MaxKeys = 3 // Maximum keys per node for simplicity.

type Node struct {
	keys     []int
	children []*Node
	leaf     bool
}

// PersistentBTree stores the B-tree and interacts with storage.
type PersistentBTree struct {
	root     *Node
	storage  *storage.FileStorage
	nextNode int
}

// NewPersistentBTree initializes a B-tree with persistent storage.
func NewPersistentBTree(storage *storage.FileStorage) *PersistentBTree {
	root := &Node{leaf: true}
	return &PersistentBTree{
		root:     root,
		storage:  storage,
		nextNode: 1,
	}
}

// SaveNode persists a node to storage.
func (bt *PersistentBTree) SaveNode(node *Node, nodeID int) error {
	data := &storage.NodeData{
		Keys:     node.keys,
		Leaf:     node.leaf,
		Children: []int{},
	}
	for _, child := range node.children {
		childID := bt.nextNode
		bt.nextNode++
		bt.SaveNode(child, childID)
		data.Children = append(data.Children, childID)
	}
	return bt.storage.SaveNode(nodeID, data)
}

// LoadNode loads a node from storage.
func (bt *PersistentBTree) LoadNode(nodeID int) (*Node, error) {
	data, err := bt.storage.LoadNode(nodeID)
	if err != nil || data == nil {
		return nil, err
	}
	node := &Node{
		keys: data.Keys,
		leaf: data.Leaf,
	}
	for _, childID := range data.Children {
		child, err := bt.LoadNode(childID)
		if err != nil {
			return nil, err
		}
		node.children = append(node.children, child)
	}
	return node, nil
}

// Insert inserts a key into the persistent B-tree.
func (bt *PersistentBTree) Insert(key int) error {
	root := bt.root
	if len(root.keys) == MaxKeys { // root is full
		// Create a new root and split the old root.
		newRoot := &Node{}
		newRoot.children = append(newRoot.children, root)
		bt.splitChild(newRoot, 0)
		bt.root = newRoot
		bt.insertNonFull(newRoot, key)
	} else {
		bt.insertNonFull(root, key)
	}

	return bt.SaveNode(bt.root, 0) // Save the updated root
}

// insertNonFull inserts a key into a node that is not full.
func (bt *PersistentBTree) insertNonFull(node *Node, key int) {
	i := len(node.keys) - 1
	if node.leaf {
		// Insert the key into the correct position.
		node.keys = append(node.keys, 0)
		for i >= 0 && key < node.keys[i] {
			node.keys[i+1] = node.keys[i]
			i--
		}
		node.keys[i+1] = key
	} else {
		// Find the child to recurse into.
		for i >= 0 && key < node.keys[i] {
			i--
		}
		i++

		// If the child is full, split it first.
		if len(node.children[i].keys) == MaxKeys {
			bt.splitChild(node, i)
			// After splitting, the key may have been inserted into the new child, so we choose the correct child.
			if key > node.keys[i] {
				i++
			}
		}
		bt.insertNonFull(node.children[i], key)
	}
}

// splitChild splits the child of a node that is full.
func (bt *PersistentBTree) splitChild(parent *Node, index int) {
	child := parent.children[index]
	mid := MaxKeys / 2
	newChild := &Node{leaf: child.leaf}
	newChild.keys = append(newChild.keys, child.keys[mid+1:]...)
	child.keys = child.keys[:mid]

	// Split the children if not leaf
	if !child.leaf {
		newChild.children = append(newChild.children, child.children[mid+1:]...)
		child.children = child.children[:mid+1]
	}

	// Move the middle key up to the parent
	parent.keys = append(parent.keys[:index], append([]int{child.keys[mid]}, parent.keys[index:]...)...)
	parent.children = append(parent.children[:index+1], append([]*Node{newChild}, parent.children[index+1:]...)...)
}

// Search searches for a key in the persistent B-tree.
func (bt *PersistentBTree) Search(key int) (*Node, int, error) {
	return bt.search(bt.root, key)
}

// search recursively searches for a key in a node.
func (bt *PersistentBTree) search(node *Node, key int) (*Node, int, error) {
	i := 0
	for i < len(node.keys) && key > node.keys[i] {
		i++
	}
	if i < len(node.keys) && key == node.keys[i] {
		return node, i, nil
	}
	if node.leaf {
		return nil, -1, errors.New("key not found")
	}
	return bt.search(node.children[i], key)
}
