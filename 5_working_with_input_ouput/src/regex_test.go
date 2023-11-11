package inout

import (
	"fmt"
	"regexp"
	"testing"
)

func TestEmail(t *testing.T) {
	txt := "Та ub@gmail.com, ub@hotmail.com хаягуудаар холбоо барьж болно."

	re, _ := regexp.Compile(`\w+@\w+\.\w+`)
	all := re.FindAllString(txt, -1)

	for _, m := range all {
		fmt.Printf("%s\n", m)
	}
}

func TestNumbers(t *testing.T) {
	str := "98.0 +12.1 hello 11"
	re, _ := regexp.Compile("[+-]?\\d*\\.?\\d*")

	all := re.FindAllString(str, -1)
	fmt.Println("FindAllString = ", all)
}

func TestColors(t *testing.T) {
	colorText := "улаан,ногоон шар;цэнхэр"
	re := regexp.MustCompile("[,;\\s]")
	fmt.Printf("%q\n", re.Split(colorText, -1))
}

func TestReplace(t *testing.T) {
	re := regexp.MustCompile("i")
	fmt.Println(re.ReplaceAllString("sift rise", "o"))

	re1 := regexp.MustCompile(`(\w+),\s*(\w+)`)
	fmt.Println(re1.ReplaceAllString("one, two", "$2, $1"))
}

func TestSubmatch(t *testing.T) {
	re := regexp.MustCompile("a(x*)b(y|z)c")
	fmt.Printf("%q\n", re.FindStringSubmatch("-axxxbyc-"))
	fmt.Printf("%q\n", re.FindStringSubmatch("-abzc-"))
}

func TestFindStringSubmatch(t *testing.T) {
	re := regexp.MustCompile("a(x*)b(y|z)c")
	fmt.Printf("%q\n", re.FindStringSubmatch("-axxxbyc-"))
	fmt.Printf("%q\n", re.FindStringSubmatch("-abzc-"))
}

func TestFindAllStringSubmatch(t *testing.T) {
	txt := "Даваа Дээд:26 Доод:21 Мягмар Дээд:23 Доод:20"

	re, _ := regexp.Compile(`([а-яА-Я]+)\s*Дээд:(\d+)\s*Доод:(\d+)`)
	all := re.FindAllStringSubmatch(txt, -1)

	for _, m := range all {
		fmt.Printf("%s: %s-%s\n", m[1], m[2], m[3])
	}
}

func TestSplit(t *testing.T) {
	colorText := "улаан,ногоон шар;цэнхэр"

	re := regexp.MustCompile("[,;\\s]")
	fmt.Printf("%q\n", re.Split(colorText, -1))
}
