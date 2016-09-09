package main

import "fmt"

// Payload type for Sensors jobs like ask for temperature, humidity, external light level, etc
type JobSensors struct {
    address      uint16
    registryType int
    data         []byte
    
    status Status //maybe its a long processed payload. than status represents current status
    error  error
}

func (js JobSensors) process() {
    fmt.Println("processing JobSensors job")
    // TODO: add some sensor functionality
    return
}

func (js JobSensors) shutdown() {
    return
}
