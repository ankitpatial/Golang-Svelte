/*
 * Copyright (c) 2022. SimSaw Software Private Limited.
 * Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
 * Author: Ankit Patial
 * Dated:  08/08/22, 6:28 pm
 */

package util

import (
	"roofix/pkg/util/log"
	"time"
)

// ParseISODate ISO date format to utc
func ParseISODate(dt string) time.Time {
	d, err := time.Parse(time.RFC3339, dt)
	if err != nil {
		log.Error(err)
	}

	return d.UTC()
}

func CountWorkingDays(from, end time.Time) int {
	workingDays := 0
	start := from //from.AddDate(0, 0, 1)
	for d := start; d.Before(end) || d.Equal(end); d = d.AddDate(0, 0, 1) {
		weekday := d.Weekday()
		if weekday != time.Saturday && weekday != time.Sunday {
			workingDays++
		}
	}

	return workingDays
}

func StartOfDay(date time.Time) time.Time {
	year, month, day := date.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, date.Location())
}

func EndOfDay(date time.Time) time.Time {
	year, month, day := date.Date()
	return time.Date(year, month, day, 23, 59, 59, 0, date.Location())
}

func StartOfMonth(date time.Time) time.Time {
	year, month, _ := date.Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, date.Location())
}

func EndOfMonth(date time.Time) time.Time {
	return EndOfDay(date.AddDate(0, 1, -date.Day()))
}

func StartOfWeek(date time.Time) time.Time {
	return StartOfDay(date.AddDate(0, 0, int(time.Sunday)-int(date.Weekday())))
}

func EndOfWeek(date time.Time) time.Time {
	return EndOfDay(date.AddDate(0, 0, int(time.Saturday-date.Weekday())))
}

func StartOfYear(date time.Time) time.Time {
	year, _, _ := date.Date()
	return time.Date(year, 1, 1, 0, 0, 0, 0, date.Location())
}

func EndOfYear(date time.Time) time.Time {
	return EndOfDay(date.AddDate(0, int(time.December-date.Month())+1, -date.Day()))
}

func LastWeek(date time.Time) time.Time {
	return date.AddDate(0, 0, -7)
}

func LastMonth(date time.Time) time.Time {
	return date.AddDate(0, -1, 0)
}

func SubtractMonths(date time.Time, m int) time.Time {
	return date.AddDate(0, -m, 0)
}

func SubtractDays(date time.Time, d int) time.Time {
	return date.AddDate(0, 0, -d)
}

func SubtractAYear(date time.Time) time.Time {
	return date.AddDate(-1, 0, 0)
}
