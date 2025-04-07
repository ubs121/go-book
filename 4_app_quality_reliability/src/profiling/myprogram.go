package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "CPU профайлыг файл руу бичих")
var memprofile = flag.String("memprofile", "", "Санах ойн профайлыг файл руу бичих")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
		return
	}

	// Run a CPU-intensive task
	for i := 0; i < 10; i++ {
		go cpuIntensiveTask()
	}

	// Run a memory-intensive task
	memoryIntensiveTask()

	// Sleep to allow profiling data to be collected
	time.Sleep(10 * time.Second)
}

func cpuIntensiveTask() {
	for i := 0; i < 100000000; i++ {
		_ = i * i
	}
}

func memoryIntensiveTask() {
	const size = 1000000
	data := make([]byte, size)
	for i := 0; i < size; i++ {
		data[i] = 1
	}
}
