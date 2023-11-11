package numbers

import "testing"

var oldNumber, currentNumber, nextNumber int = 1, 1, 2

/* Фибоначийн тоон цуваа
* 	   f(n) = f(n-1)+f(n-2)
*  Энд n>1, f(1)=1, f(2)=1 байна.
 */
func TestFib(t *testing.T) {
	print("1 ") /* дарааллын эхний тоог хэвлэх */

	for currentNumber < 100 {
		print(currentNumber, " ")
		nextNumber = currentNumber + oldNumber

		oldNumber = currentNumber
		currentNumber = nextNumber
	}
}
