package concurrency

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"testing"
	"time"
)

const queueSize = 3

const dataJson = `{"datetime":"2023-06-27 22:22:19.62710.192501","value":"1","partition":"p5"}
{"datetime":"2023-06-27 22:22:19.62710.193409","value":"2","partition":"p4"}
{"datetime":"2023-06-27 22:22:19.62710.193441","value":"3","partition":"p4"}
{"datetime":"2023-06-27 22:22:19.62710.193470","value":"4","partition":"p1"}
{"datetime":"2023-06-27 22:22:19.62710.193496","value":"5","partition":"p2"}`

type Message struct {
	Timestamp string `json:"datetime"`
	Value     string `json:"value"`
	Partition string `json:"partition"`
}

func readDataJson(f io.Reader, msgChan chan *Message) error {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		var rec Message
		if err := json.Unmarshal([]byte(line), &rec); err != nil {
			return err
		}

		// send it
		msgChan <- &rec
		//fmt.Printf("msg %v\n", rec)
	}

	return nil
}

func circQueue(size int, in chan *Message, reader chan *Message) {
	rb := make([]*Message, size)
	writeIndex := 0
	readIndex := 0

	for {
		select {
		case m := <-in: // try to read from 'in'
			if m != nil {
				// add into the queue
				rb[writeIndex%size] = m
				writeIndex++
			}
		default:
			// write/feed the reader forever
			m := rb[readIndex%size]
			if m != nil {
				reader <- m
				readIndex++

				// sleep
				time.Sleep(100 * time.Millisecond)
			}

		}
	}
}

func reader(feed chan *Message) {
	for m := range feed {
		fmt.Printf("read: %v\n", m)
	}
}

func TestReadDataJson(t *testing.T) {
	msgChan := make(chan *Message)
	outChan := make(chan *Message)

	defer close(msgChan)

	// read the data.json
	go func() {
		buf := bytes.NewBuffer([]byte(dataJson))
		readDataJson(buf, msgChan)
		// close(msgChan) - channel keep receiving nil
	}()

	// start the ring buffer service
	go func() {
		circQueue(queueSize, msgChan, outChan)
	}()

	// start consuming
	reader(outChan)
}
