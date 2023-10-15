/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package postal

import (
	"strings"
)

type State struct {
	ID     string
	Name   string
	Region Region
	Cities []*City
}

type City struct {
	Name string
	Zip  string
}

var (
	Alabama = &State{
		ID:     "AL",
		Name:   "Alabama",
		Region: RegionSouth,
		Cities: []*City{
			{Name: "Birmingham", Zip: "35203"},
			{Name: "Gulf Shores", Zip: "36542"},
		},
	}
	Alaska = &State{
		ID:     "AK",
		Name:   "Alaska",
		Region: RegionWest,
		Cities: nil,
	}
	Arizona = &State{
		ID:     "AZ",
		Name:   "Arizona",
		Region: RegionWest,
		Cities: []*City{
			{Name: "Phoenix", Zip: "85003"},
			{Name: "Tucson", Zip: "85701"},
		},
	}
	Arkansas = &State{
		ID:     "AR",
		Name:   "Arkansas",
		Region: RegionSouth,
		Cities: []*City{
			{Name: "Little Rock", Zip: "72201"},
		},
	}
	California = &State{
		ID:     "CA",
		Name:   "California",
		Region: RegionWest,
		Cities: []*City{
			{Name: "Sacramento", Zip: "95814"},
			{Name: "Los Angeles", Zip: "90014"},
			{Name: "Santa Ana", Zip: "92701"},
			{Name: "San Diego", Zip: "92101"},
		},
	}
	Colorado = &State{
		ID:     "CO",
		Name:   "Colorado",
		Region: RegionWest,
		Cities: []*City{
			{Name: "Denver", Zip: "80202"},
			{Name: "Pueblo", Zip: "81003"},
		},
	}
	Connecticut = &State{
		ID:     "CT",
		Name:   "Connecticut",
		Region: RegionNortheast,
		Cities: []*City{
			{Name: "Hartford", Zip: "06103"},
			{Name: "Greenwich", Zip: "06830"},
		},
	}
	Delaware = &State{
		ID:     "DE",
		Name:   "Delaware",
		Region: RegionSouth,
		Cities: []*City{
			{Name: "Wilmington", Zip: "19803"},
			{Name: "Dover", Zip: "19904"},
		},
	}
	DistrictOfColumbia = &State{
		ID:     "DC",
		Name:   "District of Columbia",
		Region: RegionSouthEast,
		Cities: []*City{
			{Name: "Wilmington", Zip: "19803"},
			{Name: "Dover", Zip: "19904"},
		},
	}
	Florida = &State{
		ID:     "FL",
		Name:   "Florida",
		Region: RegionSouth,
		Cities: []*City{
			{Name: "Orlando", Zip: "32801"},
			{Name: "Naples", Zip: "34102"},
			{Name: "Miami", Zip: "33312"},
			{Name: "Ft. Lauderdale", Zip: "33301"},
			{Name: "Ft. Myers", Zip: "33901"},
		},
	}
	Georgia = &State{
		ID:     "GA",
		Name:   "Georgia",
		Region: RegionSouth,
		Cities: []*City{
			{Name: "Atlanta", Zip: "30303"},
			{Name: "Athena", Zip: "30601"},
		},
	}
	Hawaii = &State{
		ID:     "HI",
		Name:   "Hawaii",
		Region: RegionWest,
		Cities: nil,
	}
	Idaho = &State{
		ID:     "ID",
		Name:   "Idaho",
		Region: RegionWest,
		Cities: []*City{
			{Name: "Boise", Zip: "83702"},
			{Name: "Coeur d'Alene", Zip: "83814"},
		},
	}
	Illinois = &State{
		ID:     "IL",
		Name:   "Illinois",
		Region: RegionMidwest,
		Cities: []*City{
			{Name: "Chicago", Zip: "60610"},
			{Name: "Joliet", Zip: "60432"},
		},
	}
	Indiana = &State{
		ID:     "IN",
		Name:   "Indiana",
		Region: RegionMidwest,
		Cities: []*City{
			{Name: "Indianapolis", Zip: "46202"},
		},
	}
	Iowa = &State{
		ID:     "IA",
		Name:   "Iowa",
		Region: RegionMidwest,
		Cities: []*City{
			{Name: "Des Moines", Zip: "50309"},
		},
	}
	Kansas = &State{
		ID:     "KS",
		Name:   "Kansas",
		Region: RegionMidwest,
		Cities: []*City{
			{Name: "Kansas City", Zip: "66102"},
			{Name: "Topeka", Zip: "66609"},
		},
	}
	Kentucky = &State{
		ID:     "KY",
		Name:   "Kentucky",
		Region: RegionSouth,
		Cities: []*City{
			{Name: "Louisville", Zip: "40202"},
		},
	}
	Louisiana = &State{
		ID:     "LA",
		Name:   "Louisiana",
		Region: RegionSouth,
		Cities: []*City{
			{Name: "Shreveport", Zip: "71101"},
			{Name: "Lake Charles", Zip: "70601"},
		},
	}
	Maine = &State{
		ID:     "ME",
		Name:   "Maine",
		Region: RegionNortheast,
		Cities: []*City{
			{Name: "Portland", Zip: "04101"},
		},
	}
	Maryland = &State{
		ID:     "MD",
		Name:   "Maryland",
		Region: RegionSouth,
		Cities: []*City{
			{Name: "Baltimore", Zip: "21201"},
			{Name: "Columbia", Zip: "21044"},
		},
	}
	Massachusetts = &State{
		ID:     "MA",
		Name:   "Massachusetts",
		Region: RegionNortheast,
		Cities: []*City{
			{Name: "Boston", Zip: "02110"},
			{Name: "Springfield", Zip: "01104"},
		},
	}
	Michigan = &State{
		ID:     "MI",
		Name:   "Michigan",
		Region: RegionMidwest, Cities: []*City{
			{Name: "Detroit", Zip: "48201"},
			{Name: "Grand Rapids", Zip: "49503"},
		},
	}
	Minnesota = &State{
		ID:     "MN",
		Name:   "Minnesota",
		Region: RegionMidwest,
		Cities: []*City{
			{Name: "Minneapolis", Zip: "55401"},
			{Name: "Rochester", Zip: "55904"},
		},
	}
	Mississippi = &State{
		ID:     "MS",
		Name:   "Mississippi",
		Region: RegionSouth,
		Cities: []*City{
			{Name: "Jackson", Zip: "39201"},
			{Name: "Gulfport", Zip: "39503"},
		},
	}
	Missouri = &State{
		ID:     "MO",
		Name:   "Missouri",
		Region: RegionMidwest,
		Cities: []*City{
			{Name: "Kansas City", Zip: "64106"},
			{Name: "Columbia", Zip: "65201"},
		},
	}
	Montana = &State{
		ID:     "MT",
		Name:   "Montana",
		Region: RegionWest,
		Cities: []*City{
			{Name: "Helena", Zip: "59602"},
		},
	}
	Nebraska = &State{
		ID:     "NE",
		Name:   "Nebraska",
		Region: RegionMidwest,
		Cities: []*City{
			{Name: "Omaha", Zip: "68102"},
		},
	}
	Nevada = &State{
		ID:     "NV",
		Name:   "Nevada",
		Region: RegionWest,
		Cities: []*City{
			{Name: "Renon", Zip: "89501"},
		},
	}
	NewHampshire = &State{
		ID:     "NH",
		Name:   "New Hampshire",
		Region: RegionNortheast,
		Cities: []*City{
			{Name: "Manchester", Zip: "03104"},
		},
	}
	NewJersey = &State{
		ID:     "NJ",
		Name:   "New Jersey",
		Region: RegionNortheast,
		Cities: []*City{
			{Name: "Newark", Zip: "07102"},
		},
	}
	NewMexico = &State{
		ID:     "NM",
		Name:   "New Mexico",
		Region: RegionWest,
		Cities: []*City{
			{Name: "Santa Fe", Zip: "87501"},
			{Name: "Albuquerque", Zip: "87102"},
		},
	}
	NewYork = &State{
		ID:     "NY",
		Name:   "New York",
		Region: RegionNortheast,
		Cities: []*City{
			{Name: "New York", Zip: "10002"},
			{Name: "Buffalo", Zip: "14201"},
		},
	}
	NorthCarolina = &State{
		ID:     "NC",
		Name:   "North Carolina",
		Region: RegionSouth,
		Cities: []*City{
			{Name: "Charlotte", Zip: "28202"},
			{Name: "Carolina Beach", Zip: "28428"},
		},
	}
	NorthDakota = &State{
		ID:     "ND",
		Name:   "North Dakota",
		Region: RegionMidwest,
		Cities: []*City{
			{Name: "Fargo", Zip: "58102"},
		},
	}
	Ohio = &State{
		ID:     "OH",
		Name:   "Ohio",
		Region: RegionMidwest,
		Cities: []*City{
			{Name: "Columbus", Zip: "43215"},
			{Name: "Cleveland", Zip: "44114"},
		},
	}
	Oklahoma = &State{
		ID:     "OK",
		Name:   "Oklahoma",
		Region: RegionSouth,
		Cities: []*City{
			{Name: "Tulsa", Zip: "741008"},
		},
	}
	Oregon = &State{
		ID:     "OR",
		Name:   "Oregon",
		Region: RegionWest,
		Cities: []*City{
			{Name: "Portland", Zip: "97201"},
			{Name: "Eugene", Zip: "97401"},
		},
	}
	Pennsylvania = &State{
		ID:     "PA",
		Name:   "Pennsylvania",
		Region: RegionNortheast,
		Cities: []*City{
			{Name: "Philadelphia", Zip: "19102"},
			{Name: "Pittsburgh", Zip: "15222"},
			{Name: "Erie", Zip: "16505"},
		}}
	RhodeIsland = &State{
		ID:     "RI",
		Name:   "Rhode Island",
		Region: RegionNortheast,
		Cities: []*City{
			{Name: "Providence", Zip: "02903"},
		},
	}
	SouthCarolina = &State{
		ID:     "SC",
		Name:   "South Carolina",
		Region: RegionSouth,
		Cities: []*City{
			{Name: "Charleston", Zip: "29401"},
			{Name: "Myrtle Beach", Zip: "29577"},
		},
	}
	SouthDakota = &State{
		ID:     "SD",
		Name:   "South Dakota",
		Region: RegionMidwest,
		Cities: []*City{
			{Name: "Sioux Falls", Zip: "57104"},
		},
	}
	Tennessee = &State{
		ID:     "TN",
		Name:   "Tennessee",
		Region: RegionSouth,
		Cities: []*City{
			{Name: "Nashville", Zip: "37201"},
			{Name: "Knoxville", Zip: "37902"},
		},
	}
	Texas = &State{
		ID:     "TX",
		Name:   "Texas",
		Region: RegionSouth,
		Cities: []*City{
			{Name: "Houston", Zip: "77002"},
			{Name: "Galveston", Zip: "77551"},
			{Name: "Lubbock", Zip: "79401"},
		},
	}
	Utah = &State{
		ID:     "UT",
		Name:   "Utah",
		Region: RegionWest,
		Cities: []*City{
			{Name: "Salt Lake City", Zip: "84101"},
			{Name: "RegionWest Jordan", Zip: "84088"},
		},
	}
	Vermont = &State{
		ID:     "VT",
		Name:   "Vermont",
		Region: RegionNortheast,
		Cities: []*City{
			{Name: "Montpelier", Zip: "05602"},
		},
	}
	Virginia = &State{
		ID:     "VA",
		Name:   "Virginia",
		Region: RegionSouth,
		Cities: []*City{
			{Name: "Virginia Beach", Zip: "23454"},
			{Name: "Alexandria", Zip: "22302"},
		},
	}
	Washington = &State{
		ID:     "WA",
		Name:   "Washington",
		Region: RegionWest,
		Cities: []*City{
			{Name: "Seattle", Zip: "98121"},
			{Name: "Spokane", Zip: "99201"},
		},
	}
	WestVirginia = &State{
		ID:     "WV",
		Name:   "West Virginia",
		Region: RegionSouth,
		Cities: []*City{
			{Name: "Charleston", Zip: "25301"},
		},
	}
	Wisconsin = &State{
		ID:     "WI",
		Name:   "Wisconsin",
		Region: RegionMidwest,
		Cities: []*City{
			{Name: "Green Bay", Zip: "54301"},
		},
	}
	Wyoming = &State{
		ID:     "WY",
		Name:   "Wyoming",
		Region: RegionWest,
		Cities: []*City{
			{Name: "Cheyenne", Zip: "82009"},
		},
	}

	AllStates = []*State{
		Alabama, Alaska, Arizona, Arkansas, California, Colorado, Connecticut, Delaware, DistrictOfColumbia, Florida, Georgia, Hawaii, Idaho,
		Illinois, Indiana, Iowa, Kansas, Kentucky, Louisiana, Maine, Maryland, Massachusetts, Michigan, Minnesota,
		Mississippi, Missouri, Montana, Nebraska, Nevada, NewHampshire, NewJersey, NewMexico, NewYork, NorthCarolina,
		NorthDakota, Ohio, Oklahoma, Oregon, Pennsylvania, RhodeIsland, SouthCarolina, SouthDakota, Tennessee, Texas,
		Utah, Vermont, Virginia, Washington, WestVirginia, Wisconsin, Wyoming,
	}
)

func GetStateRegion(state string) uint8 {
	st := strings.Trim(strings.ToLower(state), " ")
	for _, s := range AllStates {
		if strings.ToLower(s.ID) == st || strings.ToLower(s.Name) == st {
			return s.Region.ToUInt8()
		}
	}

	return RegionUnknown.ToUInt8()
}

func Search(q string, take int) []*State {
	var result []*State
	if strings.TrimSpace(q) == "" {
		return result
	}

	q = strings.ToLower(q)
	var abr, name string
	for _, s := range AllStates {
		if len(result) == take {
			break
		}
		abr = strings.ToLower(s.ID)
		name = strings.ToLower(s.Name)
		if abr == q || strings.Contains(name, q) {
			result = append(result, s)
		}
	}

	return result
}

func IsValidState(state string) bool {
	st := strings.Trim(strings.ToLower(state), " ")
	for _, s := range AllStates {
		if strings.ToLower(s.ID) == st || strings.ToLower(s.Name) == st {
			return true
		}
	}

	return false
}
