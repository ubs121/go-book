package basics

import (
	"fmt"
	"testing"
)

func TestSpeller(t *testing.T) {
	var (
		digitNames = []string{"тэг", "нэг", "хоёр", "гурав", "дөрөв", "тав", "зургаа", "долоо", "найм", "ес"}
	)

	var number int /* 3 оронтой эерэг бүхэл тоо */
	fmt.Printf("3 оронтой тоо оруулна уу: ")
	fmt.Scanf("%d", &number)

	d1 := number / 100       /* 100-тын орныг тооцоолох */
	d2 := (number / 10) % 10 /* 10-тын орныг тооцоолох */
	d3 := number % 10        /* нэгжийн орныг тооцоолох */

	fmt.Printf("%d -> %s %s %s\n", number, digitNames[d1],
		digitNames[d2], digitNames[d3])
}
