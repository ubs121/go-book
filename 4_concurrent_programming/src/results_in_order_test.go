package concurrency

import (
	"fmt"
	"testing"
)

type ResultType struct {
	Order int
	Data  any
}

func TestOrderedResultsFromWorkers(t *testing.T) {
	n := 5 // number of workers

	// a work simulation
	DoWork := func(p int) ResultType {
		return ResultType{Order: p}
	}

	// create an array of channels
	resultChans := make([]chan ResultType, n)

	for i := 0; i < n; i++ {
		// create a channel for worker 'i'
		resultChans[i] = make(chan ResultType)

		// start a worker 'i'
		go func(order int) {
			result := DoWork(order)
			resultChans[order] <- result
		}(i)
	}

	// Collect and process results in order
	for i := 0; i < n; i++ {
		result := <-resultChans[i]
		// Process the result
		fmt.Println("result ", result)
	}
}
