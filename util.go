package wxo

import (
	"strings"
	"unicode/utf8"
)

// Open Weather related

// https://home.openweathermap.org/api_keys
// API docs: https://openweathermap.org/current

type units struct {
	Distance string
	Speed    string
	Temp     string
}

var unitsMap = map[string]units{
	"metric":   {"km", "km/h", "C"},
	"imperial": {"mi", "mph", "F"},
	"kelvin":   {"m", "m/s", "K"},
}

type point struct {
	ordinal string
	arrow   string
}

var directions = []point{
	{"N", "↑"}, {"NNE", "↗"}, {"NE", "↗"}, {"ENE", "↗"},
	{"E", "→"}, {"ESE", "↘"}, {"SE", "↘"}, {"SSE", "↘"},
	{"S", "↓"}, {"SSW", "↙"}, {"SW", "↙"}, {"WSW", "↙"},
	{"W", "←"}, {"WNW", "↖"}, {"NW", "↖"}, {"NNW", "↖"},
}

// Returns ordinal and cardinal values (N, S, WNW, ESE, etc) from
// compass degrees
func DirectionFromDegree(v int, asArrow bool) string {
	val := float32(v)
	i := int((val + 11.25) / 22.5)
	if asArrow {
		return directions[i%16].arrow
	}
	return directions[i%16].ordinal
}

// Returns a direction arrow mapping on to N, E, S, W, NNE, etc
// or "" if an unsupported value is provided.
func ArrowFromOrdinal(ordinal string) string {
	ordinal = strings.ToUpper(ordinal)
	for _, v := range directions {
		if v.ordinal == ordinal {
			return v.arrow
		}
	}
	return ""
}

// GetUnits returns the units of measure
func GetUnits(uom string) units {
	v, ok := unitsMap[strings.ToLower(uom)]
	if !ok {
		v = unitsMap["error"]
	}
	return v
}

// TruncateWebString returns string with length specificed, truncated with
// multibyte runes taken into account.
func TruncateWebString(s string, length int) string {
	count := 0
	result := ""
	if len(s) <= length {
		return s
	}
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		s = s[size:]
		count = count + 1
		result = result + string(r)
		if count >= length {
			break
		}
	}
	return result
}

// Currently not used

func KelvinToC(k float64) float64 {
	return k - 2713.15
}

func KelvinToF(k float64) float64 {
	return (k-273.15)*9/5 + 32
}

func MPSToKMH(metersPerSecond float64) float64 {
	return metersPerSecond * 3.6
}

func MPSToMPH(metersPerSecond float64) float64 {
	return metersPerSecond * 2.23694
}
