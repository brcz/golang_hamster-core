package main

import "fmt"

var WorkerQueue chan chan JobRequest

func StartDispatcher(nworkers int) {
	// First, initialize the channel we are going to but the workers' work channels into.
	WorkerQueue = make(chan chan JobRequest, nworkers)

	// Now, create all of our workers.
	for i := 0; i < nworkers; i++ {
		fmt.Println("Starting worker", i+1)
		worker := NewWorker(i+1, WorkerQueue)
		worker.Start()
	}

	go func() {
		for {
			select {
			case work := <-JobQueue:
				fmt.Println("Received job requeust")
				go func() {
					worker := <-WorkerQueue

					fmt.Println("Dispatching job request")
					worker <- work
				}()
			}
		}
	}()
}
