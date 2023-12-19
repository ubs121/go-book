package search

import (
	"fmt"
	"slices"
	"testing"
)

// generate all the strings of n bits
func printBinaryDigits(n int) {

	var backtrack func(digits []byte, i int)
	backtrack = func(digits []byte, i int) {
		// бит дууссан учраас нэг хувилбар олдсон гэж үзээд хэвлэх
		if i < 1 {
			digitsRev := slices.Clone(digits)
			slices.Reverse(digitsRev)
			fmt.Printf("%s\n", string(digitsRev))
			return
		}

		// эхний сонголт: 'i-1' бит дээр '0' утгыг сонгоод цааш үргэлжлэх
		digits[i-1] = '0'
		backtrack(digits, i-1)

		// энэ цэгт буцаж байна

		// дараагийн сонголт: 'i-1' бит дээр '1'-г сонгоод цааш үргэлжлэх
		digits[i-1] = '1'
		backtrack(digits, i-1)

		// энэ цэг дээр өмнөх бит рүү буцна
	}

	// эхлүүлэх
	digits := make([]byte, n)
	backtrack(digits, n)
}

func TestPrintBinary(t *testing.T) {
	printBinaryDigits(3)
}
