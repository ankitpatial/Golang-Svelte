package num

import (
	"testing"
)

func TestFormatMoney(t *testing.T) {
	data := []struct {
		amt      float64
		expected string
	}{
		{amt: 0, expected: "$0"},
		{amt: 1, expected: "$1.00"},
		{amt: 10, expected: "$10.00"},
		{amt: 100, expected: "$100.00"},
		{amt: 1000, expected: "$1,000.00"},
		{amt: 10000, expected: "$10,000.00"},
		{amt: 100000, expected: "$100,000.00"},
		{amt: 1000000, expected: "$1,000,000.00"},
		{amt: 10000000, expected: "$10,000,000.00"},
		{amt: 100000000, expected: "$100,000,000.00"},
		{amt: 1000000000, expected: "$1,000,000,000.00"},
		{amt: 10000000000, expected: "$10,000,000,000.00"},
		{amt: 100000000000, expected: "$100,000,000,000.00"},
	}

	for _, v := range data {
		exp := "$1.00"
		if got := FormatMoney(v.amt); got != v.expected {
			t.Error("expected", exp, "but got", got)
		}
	}

}
