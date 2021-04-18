package basics

import "testing"

func wordCount(s string) int {
	c := 0

	for i := 0; i < len(s); i++ {
		for i < len(s) && s[i] == ' ' { // сул зайн тэмдэгт
			i++
		}

		if i < len(s) && s[i] != ' ' {
			c++
		}

		for i < len(s) && s[i] != ' ' { // бусад тэмдэгт
			i++
		}

	}

	return c
}

func TestWordCount(t *testing.T) {
	str := "Hello World!  "

	println(str)
	println(wordCount(str), "үг")
}
