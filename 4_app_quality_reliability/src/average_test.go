package debug

import "testing"

func Average(nums []float64) float64 {
	sum := 0.0
	for _, n := range nums {
		sum += n
	}
	return sum / float64(len(nums))
}

func TestAverage(t *testing.T) {
	// тестийн оролт болон гаралтын утгууд
	var testCases = []struct {
		values   []float64
		expected float64
	}{
		{values: []float64{1, 2}, expected: 1.5},
		{values: []float64{1, 1, 1, 1, 1, 1}, expected: 1},
		{values: []float64{-1, 1}, expected: 0},
	}

	for _, tc := range testCases {
		got := Average(tc.values) // функцийг дуудах

		if got != tc.expected {
			t.Error(
				"Тест өгөгдөл", tc.values,
				"Хүлээх үр дүн", tc.expected,
				"Гарсан үр дүн", got,
			)
		}
	}
}
