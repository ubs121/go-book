package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	var c chan string = make(chan string)

	pinger := func(c chan string) {
		for i := 0; ; i++ {
			c <- "ping"
		}
	}
	printer := func(c chan string) {
		for {
			msg := <-c
			fmt.Println(msg)
			time.Sleep(time.Second * 1)
		}
	}

	ponger := func(c chan string) {
		for i := 0; ; i++ {
			c <- "pong"
		}
	}

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
