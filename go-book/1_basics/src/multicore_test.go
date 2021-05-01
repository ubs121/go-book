package basics

import (
	"testing"
	"time"
)

func doJob(quit chan int) {
	// тодорхой ажил гүйцэтгэхийг 'time.Sleep'-ээр орлуулав
	time.Sleep(time.Millisecond * 200)
	quit <- 1
}

func TestMulticore(t *testing.T) {
	routineQuit := make(chan int)

	for i := 0; i < 50; i++ {
		go doJob(routineQuit)
	}

	// бүх функцээс дууссан дохио хүлээх
	for i := 0; i < 50; i++ {
		<-routineQuit
	}
}
