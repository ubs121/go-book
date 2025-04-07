package concurrency

import (
	"strconv"
	"testing"
)

func parallelFor() {
	jobChan := make(chan *Response)

	// ажлын жагсаалт
	n := 60
	var requests []Request
	for i := 0; i < n; i++ {
		requests = append(requests, Request{"task" + strconv.Itoa(i)})
	}

	// параллел давталт
	for _, req := range requests {
		go func(r *Request) {
			reply := doJob(r)
			jobChan <- reply
		}(&req)
	}

	// бүх функцээс өгөгдөл хүлээх
	for i := 0; i < len(requests); i++ {
		resp := <-jobChan
		println(resp.Data)
	}
}

func TestMulticore(t *testing.T) {
	parallelFor()
}
