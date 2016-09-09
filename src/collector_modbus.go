package main
// job collector that handle incoming modbus jobs

import "fmt"

func CollectorModbus() {

    // Set job initial parameters and send it to processing queue
    work := JobRequest{Payload: JobModbus{}}
    
    // Push the work onto the queue.
    JobQueue <- work
    fmt.Println("Modbus job request queued", work)

}