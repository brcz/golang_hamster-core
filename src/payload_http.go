package main

import (
	"fmt"
	"time"
)

// Payload type for HTTP jobs
type JobHTTP struct {
	Name  string
	Delay time.Duration

	status Status //maybe its a long processed payload. than status represents current status
	error  error
}

func (jh JobHTTP) process() {
	fmt.Printf("processing job, delaying for %f seconds\n", jh.Delay.Seconds())
	time.Sleep(jh.Delay)
	fmt.Printf("Hello, %s!\n", jh.Name)
	return
}

func (jh JobHTTP) shutdown() {
	return
}
