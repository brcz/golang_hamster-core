package main

import (
	"fmt"
	"time"
)

// Payload type for HTTP jobs
type JobHTTP struct {
	Name  string
	Delay time.Duration
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

// Payload type for Modbus jobs
type JobModbus struct {
	address      uint16
	registryType int
	data         []byte
}

func (jm JobModbus) process() {
	fmt.Println("processing Modbus job")
	return
}

func (jm JobModbus) shutdown() {
	return
}
