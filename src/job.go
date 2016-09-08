package main

//import "errors"
import "time"

type Job struct {
	Payload Payload
}

type WorkRequest struct {
	Name  string
	Delay time.Duration
}
