package concurrency

import (
	"strconv"
	"testing"
	"time"
)

type Request struct{ Name string }
type Response struct{ Data string }

func doJob(request *Request, ch chan *Response) {
	// 200 ms-н ажил
	time.Sleep(time.Millisecond * 200)

	// хариу илгээх
	ch <- &Response{Data: request.Name + " дууслаа"}
}

func main() {
	jobChan := make(chan *Response)

	// ажлын жагсаалт
	n := 60
	var requests []Request
	for i := 0; i < n; i++ {
		requests = append(requests, Request{"task" + strconv.Itoa(i)})
	}

	// параллел давталт
	for _, req := range requests {
		req1 := req
		go doJob(&req1, jobChan)
	}

	// бүх функцээс өгөгдөл хүлээх
	for i := 0; i < len(requests); i++ {
		resp := <-jobChan
		println(resp.Data)
	}
}

func TestMulticore(t *testing.T) {
	main()
}
