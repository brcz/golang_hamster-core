package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	NWorkers     = flag.Int("n", 4, "The number of workers to start")
	JobQueueSize = flag.Int("j", 50, "The size of job queue")
	//HTTPAddr = flag.String("http", "127.0.0.1:8000", "Address to listen for HTTP requests on")

	// A buffered channel that we can send work requests on.
	WorkQueue = make(chan WorkRequest, 100)

	// A buffered channel that we can send work requests on.
	JobQueue chan Job
)

func main() {
	// Parse the command-line flags.
	flag.Parse()

	JobQueue = make(chan Job, *JobQueueSize)

	// Start the dispatcher.
	fmt.Println("Starting the dispatcher")
	StartDispatcher(*NWorkers)

	// Register our collector as an HTTP handler function.
	fmt.Println("Registering the collector")
	http.HandleFunc("/work", Collector)

	// Start the HTTP server!
	fmt.Println("HTTP server listening on", *HTTPAddr)
	if err := http.ListenAndServe(*HTTPAddr, nil); err != nil {
		fmt.Println(err.Error())
	}
}
