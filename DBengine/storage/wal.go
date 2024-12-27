package storage

import (
	"bufio"
	"os"
	"sync"
)

type WAL struct {
	file *os.File
	mu   sync.Mutex
}

// OpenWAL initializes or opens a WAL file.
func OpenWAL(filename string) (*WAL, error) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return &WAL{file: file}, nil
}

// WriteLog appends an operation to the WAL.
func (wal *WAL) WriteLog(operation string) error {
	wal.mu.Lock()
	defer wal.mu.Unlock()
	_, err := wal.file.WriteString(operation + "\n")
	if err != nil {
		return err
	}
	return wal.file.Sync()
}

// ReadLogs reads all operations from the WAL.
func (wal *WAL) ReadLogs() ([]string, error) {
	wal.mu.Lock()
	defer wal.mu.Unlock()

	_, err := wal.file.Seek(0, 0) // Reset to beginning of file.
	if err != nil {
		return nil, err
	}

	var logs []string
	scanner := bufio.NewScanner(wal.file)
	for scanner.Scan() {
		logs = append(logs, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return logs, nil
}

// Close closes the WAL file.
func (wal *WAL) Close() error {
	return wal.file.Close()
}
