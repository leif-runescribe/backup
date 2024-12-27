package tests

import (
	"embedded-db/storage"
	"testing"
)

func TestWAL(t *testing.T) {
	wal, err := storage.OpenWAL("test.wal")
	if err != nil {
		t.Fatal(err)
	}
	defer wal.Close()

	err = wal.WriteLog("INSERT 10")
	if err != nil {
		t.Fatal(err)
	}

	logs, err := wal.ReadLogs()
	if err != nil {
		t.Fatal(err)
	}

	if len(logs) != 1 || logs[0] != "INSERT 10" {
		t.Fatalf("Unexpected logs: %+v", logs)
	}
}

func TestFileStorage(t *testing.T) {
	fs := storage.NewFileStorage("test_data")
	node := &storage.NodeData{
		Keys: []int{10, 20},
		Leaf: true,
	}

	err := fs.SaveNode(1, node)
	if err != nil {
		t.Fatal(err)
	}

	loadedNode, err := fs.LoadNode(1)
	if err != nil || loadedNode == nil || len(loadedNode.Keys) != 2 {
		t.Fatalf("Unexpected loaded node: %+v", loadedNode)
	}
}
