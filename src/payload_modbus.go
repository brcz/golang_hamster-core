package main

import "fmt"

// Payload type for Modbus jobs
type JobModbus struct {
	address      uint16
	registryType int
	data         []byte

	status Status //maybe its a long processed payload. than status represents current status
	error  error
}

func (jm JobModbus) process() {
	fmt.Println("processing Modbus job")
    // TODO: add some modbus communication functionality
	return
}

func (jm JobModbus) shutdown() {
	return
}
