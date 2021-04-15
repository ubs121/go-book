package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
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
