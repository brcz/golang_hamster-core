package main

//import "errors"

// job request placed into queue channel
type JobRequest struct {
	Payload Payload
}

// job payload. every type of payload able to run Process method.
// TODO: think about
type Payload interface {
	process()
	shutdown()
}
