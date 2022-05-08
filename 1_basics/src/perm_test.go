package basics

import (
	"testing"
)

func TestPermutation(t *testing.T) {
	var (
		a [100]int /* сэлгэх утгуудыг агуулах массив */
		N int      /* сэлгэх элементийн тоо */
	)

	N = 10
	display := func() {
		for i := 0; i < N; i++ {
			print(a[i], " ")
		}
		println()
	}

	for i := 0; i < N; i++ {
		a[i] = i + 1
	}

	p := make([]int, N) /* сэлгэмлийг удирдах массив */
	var (
		j    int
		temp int
	)

	/* эхний хувилбарыг хэвлэх */
	display()

	for i := 1; i < N; {
		if p[i] < i {
			if i%2 > 0 {
				j = p[i]
			} else {
				j = 0
			}
			temp = a[j]
			a[j] = a[i]
			a[i] = temp

			/* сэлгэмлийн нэг хувилбарыг хэвлэх */
			display()

			p[i]++
			i = 1
		} else {
			p[i] = 0
			i++
		}
	}
}
