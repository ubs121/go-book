package concurrency

import (
	"fmt"
	"strconv"
	"testing"
)

func orderedResults(n int) {
	// сувгийн жагсаалт
	results := make([]chan *Response, n)

	for i := 0; i < n; i++ {
		//  'i' ажилд зориулж суваг үүсгэх
		results[i] = make(chan *Response)

		// шинэ ажил эхлүүлэх
		go func(order int) {
			reply := doJob(&Request{Name: "task" + strconv.Itoa(order)})
			results[order] <- reply
		}(i)
	}

	// үр дүнг дарааллаар нь цуглуулах
	for i := 0; i < n; i++ {
		result := <-results[i]

		fmt.Println("result ", result)
	}
}

func TestOrderedResultsFromWorkers(t *testing.T) {
	n := 5 // number of workers
	orderedResults(n)
}
