package main

import (
	"fmt"
	"time"
)

// Worker represents a worker that executes jobs
type Worker struct {
	ID    int
	JobCh chan Job
}

// NewWorker creates a new worker
func NewWorker(id int) *Worker {
	return &Worker{
		ID:    id,
		JobCh: make(chan Job),
	}
}

// StartWorker starts the worker to listen for jobs
func (w *Worker) StartWorker(master *Master) {
	master.RegisterWorker(w.JobCh)
	go func() {
		for job := range w.JobCh {
			fmt.Printf("Worker %d processing job %d: %s\n", w.ID, job.ID, job.Name)
			time.Sleep(2 * time.Second) // Simulate job processing time
			fmt.Printf("Worker %d completed job %d\n", w.ID, job.ID)
		}
	}()
}
