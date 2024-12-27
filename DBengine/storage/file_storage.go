_package storage

import (
	"encoding/gob"
	"errors"
	"os"
	"sync"
)

// NodeData represents a persistable B-tree node.
type NodeData struct {
	Keys     []int
	Children []int // Child node IDs.
	Leaf     bool
}

// FileStorage manages disk-backed storage for nodes.
type FileStorage struct {
	basePath string
	mu       sync.Mutex
}

// NewFileStorage initializes file storage at the given directory.
func NewFileStorage(basePath string) *FileStorage {
	_ = os.MkdirAll(basePath, 0755) // Create directory if it doesn't exist.
	return &FileStorage{basePath: basePath}
}

// SaveNode saves a node to disk.
func (fs *FileStorage) SaveNode(nodeID int, data *NodeData) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	file, err := os.Create(fs.nodeFilePath(nodeID))
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	return encoder.Encode(data)
}

// LoadNode loads a node from disk.
func (fs *FileStorage) LoadNode(nodeID int) (*NodeData, error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	file, err := os.Open(fs.nodeFilePath(nodeID))
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil // Node does not exist yet.
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data NodeData
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&data)
	return &data, err
}

// nodeFilePath constructs the file path for a given node ID.
func (fs *FileStorage) nodeFilePath(nodeID int) string {
	return fmt.Sprintf("%s/node_%d.gob", fs.basePath, nodeID)
}
