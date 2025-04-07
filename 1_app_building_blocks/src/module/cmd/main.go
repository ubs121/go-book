package main

import (
	"fmt"

	"example.com/mymodule/greetings"
)

func main() {
	message := greetings.Hello("World!")
	fmt.Println(message)
}
