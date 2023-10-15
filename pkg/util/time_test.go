package util

import (
	"testing"
	"time"
)

func TestCountWorkingDays(t *testing.T) {
	s, _ := time.Parse(time.DateOnly, "2023-03-31")
	e, _ := time.Parse(time.DateOnly, "2023-03-31")
	c := CountWorkingDays(s, e)
	if c != 1 {
		t.Error("expect it to be", 0, "but got as", c)
	}
}
