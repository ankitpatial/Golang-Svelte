package eagleview

import "testing"

func TestMeasurement(t *testing.T) {
	r := &ReportDetail{
		Measurement: &Measurement{
			AreaValue:  4379.0,
			PitchValue: 12.0,
			PitchTable: []*AreaWithPitch{
				{
					Area:           "508.5 sq. ft",
					Pitch:          "4/12",
					AreaPercentage: "11.6%",
				},
				{
					Area:           "1096.4 sq. ft",
					Pitch:          "6/12",
					AreaPercentage: "25%",
				},
				{
					Area:           "337.5 sq. ft",
					Pitch:          "10/12",
					AreaPercentage: "7.7%",
				},
				{
					Area:           "2436.6 sq. ft",
					Pitch:          "12/12",
					AreaPercentage: "55.6%",
				},
			},
		},
	}

	result := measurements(r)
	if len(result) == 0 {
		t.Error("must have values")
	}
}
