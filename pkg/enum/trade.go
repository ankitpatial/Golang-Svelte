package enum

import (
	"fmt"
	"io"
	"strconv"
)

type Trade string

const (
	TradeRoofing    Trade = "Roofing"
	TradeInsulation Trade = "Insulation"
	TradeMPUs       Trade = "MPUs"
	TradeSmartHome  Trade = "SmartHome"
	TradeHVAC       Trade = "HVAC"
	TradeWindows    Trade = "Windows"
	TradeTurf       Trade = "Turf"
	TradeBatteries  Trade = "Batteries"
)

var AllTrades = []Trade{
	TradeRoofing,
	TradeInsulation,
	TradeMPUs,
	TradeSmartHome,
	TradeHVAC,
	TradeWindows,
	TradeTurf,
	TradeBatteries,
}

func (t Trade) String() string {
	return string(t)
}

// Values provides list valid values for Enum.
func (Trade) Values() (kinds []string) {
	for _, s := range AllTrades {
		kinds = append(kinds, string(s))
	}
	return
}

func (t Trade) IsValid() bool {
	for _, t := range AllTrades {
		if t == t {
			return true
		}
	}
	return false
}

func (t *Trade) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*t = Trade(str)
	if !t.IsValid() {
		return fmt.Errorf("%s is not a valid TrainingType", str)
	}
	return nil
}

func (t Trade) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(t.String()))
}
