package main

// job payload. every type of payload able to run Process method.
// TODO: think about

type (
	Payload interface {
		Process()
		Shutdown()
	}
	/*
		inputData  interface{}
		outputData interface{}
		Status     int

		struct {
			id     int
			input  inputData
			result outputData

			status Status  //maybe its a long processed payload. than status represents current status
			error  error
		}
	*/
)

func Process() {
}
