package basics

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestArraySort(t *testing.T) {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')

	splits := strings.Split(str, " ") // сул зайгаар хуваах
	arr := make([]int, len(splits))   // тоон массив үүсгэх

	for i, s := range splits {
		arr[i], _ = strconv.Atoi(s) //элемент бүрийг тоо болгон хувиргах
	}

	/* массивыг эрэмбэлэх */
	sort.Ints(arr)

	/* эрэмбэлэгдсэн  массивыг хэвлэж харуулах */
	fmt.Printf("%v", arr)
}

func TestArrayAvg(t *testing.T) {
	data := []float32{34.0, 27.0, 45.0, 82.0, 22.0}

	total := data[0] + data[1] + data[2] + data[3] + data[4]
	average := total / 5.0
	fmt.Printf("Нийт %f Дундаж %f\n", total, average)
}
