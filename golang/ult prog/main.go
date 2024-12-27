package main

import (
	"log"
	"net/http"
)

func main() {
	master := NewMaster(10) // Master with a queue size of 10
	go master.DispatchJobs()

	// Start a few workers
	for i := 1; i <= 3; i++ {
		worker := NewWorker(i)
		worker.StartWorker(master)
	}

	// Set up the HTTP server for job submission
	http.HandleFunc("/submit", master.SubmitJob)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
