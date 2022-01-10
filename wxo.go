package wxo

import (
	"fmt"
	"os"
	"text/template"
	"time"
)

const DefaultCacheExpiry = 15 * 60 // 15 minutes or 96 remote requests per day
const defaultFormat = "{{.Condition}} {{printf \"%.1f\" .Temp}}{{.TempUnits}} {{.WindVane}}{{.WindDirection}} {{printf \"%.1f\" .WindSpeed}}{{.WindSpeedUnits}}"

// TODO NOT USING THIS YET
type WxoConfig struct {
	CacheExpiry int
	APIKey      string
	Provider    WeatherClient
}

// TODO not used yet
type WeatherClient interface {
	Fetch() (*SiteData, error)
}

// SiteData provides a minimal, cross-provider view of weather data
type SiteData struct {
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
	WindDirection  string // ordinal N, S, E, W, NE, etc
	WindVane       string // arrow pointing the general direction
	Latitude       float64
	Longitude      float64
	Country        string
	Sunrise        time.Time
	Sunset         time.Time
}

// Execute allows callers to supply their own format (Go text template) string
func (s *SiteData) Execute(format string) {
	t, err := template.New("wxo").Parse(format)
	if err != nil {
		fmt.Fprintf(os.Stdout, "%v", err)
	}
	err = t.Execute(os.Stdout, s)
	if err != nil {
		fmt.Fprintf(os.Stdout, "%v", err)
	}
}

// Print the gathered site data with a default output format
func (s *SiteData) Print() {
	s.Execute(defaultFormat)
}
