/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2023.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package config

import (
	"time"
)

// location of business/app time zone
func location() *time.Location {
	loc, _ := time.LoadLocation("US/Central") // TimeZone is US/Central, America/Chicago
	return loc
}

// TimeNowAppTZ is time.now as per business time location
func TimeNowAppTZ() time.Time {
	now := time.Now()
	return now.In(location())
}

// TimeInAppTZ as per the business time location
func TimeInAppTZ(t time.Time) time.Time {
	return t.In(location())
}

// TimeStrAppTZ as per the business time location
func TimeStrAppTZ(t time.Time) string {
	return t.In(location()).Format("2006-01-02 15:04 MST")
}

// TimeParseInAppTZ as per the business time location
func TimeParseInAppTZ(layout, dt string) (time.Time, error) {
	return time.ParseInLocation(layout, dt, location())
}
