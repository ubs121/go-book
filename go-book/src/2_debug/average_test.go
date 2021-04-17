package math

import "testing"

func TestAverage(t *testing.T) {
	type testpair struct {
		values  []float64
		average float64
	}

	// тестийн оролт болон гаралтын утгууд
	var tests = []testpair{
		{[]float64{1, 2}, 1.5},
		{[]float64{1, 1, 1, 1, 1, 1}, 1},
		{[]float64{-1, 1}, 0},
	}

	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"Оролт & Гаралт", pair.values,
				"Хүлээх үр дүн", pair.average,
				"Бодит үр дүн", v,
			)
		}
	}
}
