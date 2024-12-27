package main

import (
	"fmt"
	"sync"
)

type TaskQueue struct {
	jobs  chan Job   // Buffered channel to hold jobs
	mu    sync.Mutex // Mutex to ensure thread-safe operations
	jobID int        // Auto-incrementing job ID
}

func NewTaskQueue(bufferSize int) *TaskQueue {
	return &TaskQueue{
		jobs:  make(chan Job, bufferSize),
		jobID: 0,
	}
}

// Add a new job to the queue
func (q *TaskQueue) AddJob(name string, description string) Job {
	q.mu.Lock()
	q.jobID++
	job := Job{
		ID:          q.jobID,
		Name:        name,
		Description: description,
	}
	q.jobs <- job
	q.mu.Unlock()
	fmt.Printf("Added job: %v\n", job)
	return job
}

// Fetch a job from the queue
func (q *TaskQueue) GetJob() Job {
	job := <-q.jobs
	fmt.Printf("Fetched job: %v\n", job)
	return job
}
