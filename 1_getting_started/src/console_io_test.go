package common

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
	"unicode"
)

func TestSscanf(t *testing.T) {
	inputStr := "10 20 30"
	var a, b, c int

	fmt.Sscanf(inputStr, "%d %d %d", &a, &b, &c)
	println(a, b, c)
}

func printAt(s string, x int, y int) {
	// курсорыг зөөх
	print("\033[" + strconv.Itoa(y) + ";" + strconv.Itoa(x) + "H")
	print("\033[92m")
	print(s + "\r")
}

func clearScreen() {
	print("\033[2J")
}

func matrix(c chan<- string) {
	for {
		// санамсаргүй тэмдэгт зохиож илгээх
		r := rune(rand.Intn(65_536))
		if unicode.IsPrint(r) {
			c <- strconv.QuoteRune(r)
		}
	}
}

func printer(c <-chan string) {
	column := 1
	i := 1
	for {
		msg := <-c
		printAt(msg, column, i)

		// багануудад хуваах
		if column < 80 {
			column = column + 10
		} else {
			column = 1
			i = i + 1
			println()
		}

		time.Sleep(time.Millisecond * 15)
	}
}

func TestMatrix(t *testing.T) {
	var c chan string = make(chan string)

	clearScreen() // дэлгэц цэвэрлэх

	go matrix(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
