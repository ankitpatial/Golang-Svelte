package num

import (
	"bytes"
	"fmt"
	"math"
	"roofix/config"
	"roofix/pkg/util/log"
	"strconv"
	"strings"
)

func ReadIntPtr(val *int) int {
	if val != nil {
		return *val
	}

	return 0
}

func ReadUIntPtr(val *uint) uint {
	if val != nil {
		return *val
	}

	return 0
}

func ReadUInt8Ptr(val *uint8) uint8 {
	if val != nil {
		return *val
	}

	return 0
}

func ReadInt8Ptr(val *int8) int8 {
	if val != nil {
		return *val
	}

	return 0
}

func ReadFloatPtr(val *float64) float64 {
	if val != nil {
		return *val
	}

	return 0
}

func ParseFloat32(inp string) float32 {
	if inp == "" {
		return 0
	}

	if strings.Contains(inp, ",") {
		inp = strings.ReplaceAll(inp, ",", "")
	}

	r, err := strconv.ParseFloat(inp, 32)
	if err != nil {
		log.Error(err)
		return 0
	}

	return float32(r)
}

func ParseFloat(inp string) float64 {
	if inp == "" {
		return 0
	}

	if strings.Contains(inp, ",") {
		inp = strings.ReplaceAll(inp, ",", "")
	}

	r, err := strconv.ParseFloat(inp, 32)
	if err != nil {
		log.Error(err)
		return 0
	}

	return r
}

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

// FormatMoney with a thousand comma and decimal place
func FormatMoney(value float64) string {
	if value == 0 {
		return fmt.Sprintf("%s0", config.MoneySymbol)
	}

	var v string
	if math.Trunc(value) == value {
		v = formatNumberString(fmt.Sprintf("%.0f", value), ",", ".")
	} else {
		v = formatNumberString(fmt.Sprintf("%.2f", value), ",", ".")
	}
	return fmt.Sprintf("%s%s", config.MoneySymbol, v)
}

// RoundAndFormatFloat round to 2 decimal places and format with a thousand comma
func RoundAndFormatFloat(value float64) string {
	if value == 0 {
		return "0"
	}

	return formatNumberString(fmt.Sprintf("%.2f", value), ",", ".")
}

func formatNumberString(x string, thousand string, decimalStr string) string {
	idx := strings.Index(x, ".") - 1
	if idx < 0 {
		idx = len(x) - 1
	}

	var buffer []byte
	var strBuffer bytes.Buffer

	j := 0
	for i := idx; i >= 0; i-- {
		j++
		buffer = append(buffer, x[i])

		if j == 3 && i > 0 && !(i == 1 && x[0] == '-') {
			buffer = append(buffer, ',')
			j = 0
		}
	}

	for i := len(buffer) - 1; i >= 0; i-- {
		strBuffer.WriteByte(buffer[i])
	}
	result := strBuffer.String()

	if thousand != "," {
		result = strings.Replace(result, ",", thousand, -1)
	}

	extra := x[idx+1:]
	if decimalStr != "." {
		extra = strings.Replace(extra, ".", decimalStr, 1)
	}

	return result + extra
}
