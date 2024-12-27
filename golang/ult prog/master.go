package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// Master represents the central server
type Master struct {
	taskQueue *TaskQueue
	workers   []chan Job
	mu        sync.Mutex
}

// NewMaster creates a new master with a task queue
func NewMaster(queueSize int) *Master {
	return &Master{
		taskQueue: NewTaskQueue(queueSize),
		workers:   make([]chan Job, 0),
	}
}

// RegisterWorker allows a worker to register with the master
func (m *Master) RegisterWorker(worker chan Job) {
	m.mu.Lock()
	m.workers = append(m.workers, worker)
	m.mu.Unlock()
	fmt.Println("Worker registered.")
}

// DispatchJobs continuously dispatches jobs to available workers
func (m *Master) DispatchJobs() {
	go func() {
		for {
			job := m.taskQueue.GetJob()
			for _, worker := range m.workers {
				worker <- job
				break
			}
		}
	}()
}

// SubmitJob handles incoming job submissions from HTTP requests
func (m *Master) SubmitJob(w http.ResponseWriter, r *http.Request) {
	var jobData Job
	err := json.NewDecoder(r.Body).Decode(&jobData)
	if err != nil {
		http.Error(w, "Invalid job data", http.StatusBadRequest)
		return
	}
	job := m.taskQueue.AddJob(jobData.Name, jobData.Description)
	json.NewEncoder(w).Encode(job)
}
