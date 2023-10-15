/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package postal

import (
	"fmt"
	"io"
	"strconv"
)

type Detail struct {
	Country   string
	State     string
	StateAbr  string
	Province  string
	City      string
	Zip       string
	Latitude  float64
	Longitude float64
	Accuracy  uint8 // accuracy of lat/lng, 1=estimated, 4=geonameid, 6=centroid
}

// Region is US country region ==>
// sample ref: https://education.nationalgeographic.org/resource/united-states-regions
type Region string

const (
	CountryUS              = "US"
	RegionUnknown   Region = ""
	RegionNortheast Region = "Northeast"
	RegionSouth     Region = "South"
	RegionSouthEast Region = "SouthEast"
	RegionSouthWest Region = "SouthWest"
	RegionMidwest   Region = "Midwest"
	RegionWest      Region = "West"
)

func (e Region) ToUInt8() uint8 {
	switch e {
	case RegionNortheast:
		return 1
	case RegionSouth:
		return 2
	case RegionWest:
		return 3
	case RegionMidwest:
		return 4
	case RegionSouthEast:
		return 5

	default:
		return 0
	}
}

func (e *Region) FromUInt8(v uint8) {
	switch v {
	case 0:
		*e = RegionUnknown
	case 1:
		*e = RegionNortheast
	case 2:
		*e = RegionSouth
	case 3:
		*e = RegionWest
	case 4:
		*e = RegionMidwest
	case 5:
		*e = RegionSouthEast
	}
}

func (e Region) IsValid() bool {
	switch e {
	case RegionUnknown, RegionNortheast, RegionSouth, RegionWest, RegionMidwest:
		return true
	}
	return false
}

func (e Region) String() string {
	return string(e)
}

func (e *Region) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Region(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Region", str)
	}
	return nil
}

func (e Region) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// ~~
