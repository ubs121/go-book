package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	reader, writer := io.Pipe()

	go func() {
		defer writer.Close()
		fmt.Fprintln(writer, "Hello, Pipe!")
	}()

	_, err := io.Copy(os.Stdout, reader)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
