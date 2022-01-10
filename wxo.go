package wxo

import (
	"fmt"
	"os"
	"strings"
	"text/template"
	"time"
)

var DefaultCacheExpiry int = 5 * 60 // 5 minutes or 288 remote requests per day
const defaultFormat = "{{.Alerts}} {{.Condition}} {{printf \"%.1f\" .Temp}}{{.TempUnits}} {{.WindVane}}{{.WindDirection}} {{printf \"%.1f\" .WindSpeed}}{{.WindSpeedUnits}}"

// WeatherClient provides a means to fetch from a provider a populated SiteData object.
type WeatherClient interface {
	Fetch() (*SiteData, error)
}

// SiteData provides a minimal, cross-provider view of weather data
type SiteData struct {
	Alerts         string // a concatenation of alerts, possibly truncated
	Location       string // name
	Timestamp      time.Time
	Timezone       string
	TimezoneOffset int
	Condition      string // Sunny, Light Rain, Hailing Frogs
	Temp           float64
	TempUnits      string
	TempFeelsLike  float64
	WindSpeed      float64
	WindSpeedUnits string
	WindGust       float64
	WindDegree     int
	WindDirection  string // ordinal N, S, E, W, NE, NNW etc
	WindVane       string // arrow pointing the general direction
	Latitude       float64
	Longitude      float64
	Country        string
	Sunrise        time.Time
	Sunset         time.Time
}

// ExecuteTemplate allows callers to supply their own format (Go text template)
// string. If format string isn't the default format, fall back and provide at
// least some output using the defaultFormat.
func (s *SiteData) ExecuteTemplate(format string) {
	t, err := template.New("wxo").Parse(format)

	if err != nil {
		fmt.Fprintf(os.Stdout, "%v", err)
		// attempt to fall back
		if format != defaultFormat {
			t, err = template.New("wxo").Parse("Template Error! Fallback: " + defaultFormat)
			if err != nil {
				fmt.Fprintf(os.Stdout, "%v", err)
				return
			}
		} else {
			return
		}
	}
	buf := strings.Builder{}
	err = t.Execute(&buf, s)
	if err != nil {
		fmt.Fprintf(os.Stdout, "%v", err)
	}
	fmt.Fprint(os.Stdout, strings.TrimSpace(buf.String()))
}

// Print the gathered site data with a default output format
func (s *SiteData) Print() {
	s.ExecuteTemplate(defaultFormat)
}
