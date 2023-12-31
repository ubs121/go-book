package concurrency

import "time"

type Job interface {
	ID() string
	Do() error
}

type JobResult struct {
	JobID string
	Err   error
}

type Request struct{ Name string }
type Response struct{ Data string }

// ажил гүйцэтгэх
func doJob(request *Request) *Response {
	time.Sleep(200 * time.Millisecond) // 200 ms-н ажил

	// хариу илгээх
	return &Response{Data: request.Name + " дууслаа"}
}
