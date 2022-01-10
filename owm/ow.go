package owm

// Package owm provides a client for openweathermap.org's "one call" API

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/solutionroute/wxo"
	"github.com/solutionroute/wxo/internal"
)

const baseURL = "https://api.openweathermap.org/data/2.5/onecall"

// static check ensures we implement the interface, simple as it is
var _ wxo.WeatherClient = (*OpenWeatherMapClient)(nil)

// OpenWeatherMapClient represents a client to the OWM API.
// Fields, in particular appID and uri, are not exported to avoid accidental
// disclosure of secrets.
type OpenWeatherMapClient struct {
	uri   string
	units string
}

// NewWeatherClient sets up an OpenWeatherMap.org api client
func NewWeatherClient(appID string, lattitude float64, longitude float64, units string) *OpenWeatherMapClient {

	v := url.Values{}
	v.Set("exclude", "minutely,hourly,daily") // we do nothing with that data
	lat := fmt.Sprintf("%f", lattitude)
	lon := fmt.Sprintf("%f", longitude)
	v.Set("units", units)
	v.Set("lat", lat) // lat, lon, appid are required parameters of the OWM API
	v.Set("lon", lon)
	v.Set("appid", appID)
	return &OpenWeatherMapClient{
		uri:   baseURL + "?" + v.Encode(),
		units: units,
	}
}

// Fetch returns a completed as can be SiteData object
func (c *OpenWeatherMapClient) Fetch() (*wxo.SiteData, error) {
	// GetData will return []byte from cache or a new http.Get
	data, err := internal.GetData(c.uri)
	if err != nil {
		return nil, err
	}
	r := &oneCallResponse{}
	err = json.Unmarshal(data, &r)
	if err != nil {
		return nil, err
	}
	return c.makeSiteData(*r)
}

// makeSiteData converts the weather provider's data into a common vocabulary in
// wxo.SiteData. Copy this pattern for new weather provider implementations.
func (c *OpenWeatherMapClient) makeSiteData(r oneCallResponse) (*wxo.SiteData, error) {
	sd := &wxo.SiteData{}

	sd.Timezone = r.Timezone
	sd.TimezoneOffset = r.TimezoneOffset
	tz, err := time.LoadLocation(sd.Timezone)
	if err != nil {
		tz = time.FixedZone("Local", sd.TimezoneOffset)
	}
	// Observation timestamp
	sd.Timestamp = time.Unix(int64(r.Current.Dt), 0).In(tz)

	// Conditions are text summaries "Windy", "Light Rain"
	conditions := []string{}
	for _, v := range r.Current.Weather {
		conditions = append(conditions, strings.Title(v.Description))
	}
	sd.Condition = strings.Join(conditions, "/")

	// Temperature
	sd.Temp = r.Current.Temp
	sd.TempUnits = wxo.GetUnits(c.units).Temp
	sd.TempFeelsLike = r.Current.FeelsLike

	// Wind
	// TODO deal with conversion of meters/sec to km/h or mi/h
	sd.WindSpeed = r.Current.WindSpeed * 3.6
	sd.WindSpeedUnits = wxo.GetUnits(c.units).Speed
	sd.WindGust = r.Current.WindGust * 3.6 // source is m/sec
	sd.WindDegree = int(r.Current.WindDeg)
	sd.WindDirection = wxo.DirectionFromDegree(sd.WindDegree, false)
	sd.WindVane = wxo.ArrowFromOrdinal(sd.WindDirection)

	// Location related
	sd.Country = "" // not available in this api
	sd.Latitude = r.Lat
	sd.Longitude = r.Lon
	sd.Sunrise = time.Unix(int64(r.Current.Sunrise), 0).In(tz)
	sd.Sunset = time.Unix(int64(r.Current.Sunset), 0).In(tz)

	return sd, nil
}

// oneCallResponse provides a structure to unmarshall a JSON response from
// the OneCall API documented at https://openweathermap.org/api/one-call-api.
//
// Note: Their API will not populate fields with nil values, i.e. no Rain value
// when no rain, or no Rain value when it's snowing. This code was
// auto-generated from a representative ow response, with minor tuning by hand.
//
// Useful tool: https://mholt.github.io/json-to-go/
type oneCallResponse struct {
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	Timezone       string  `json:"timezone"`
	TimezoneOffset int     `json:"timezone_offset"`
	Current        struct {
		Dt         int     `json:"dt"`
		Sunrise    int     `json:"sunrise"`
		Sunset     int     `json:"sunset"`
		Temp       float64 `json:"temp"`
		FeelsLike  float64 `json:"feels_like"`
		Pressure   int     `json:"pressure"`
		Humidity   int     `json:"humidity"`
		DewPoint   float64 `json:"dew_point"`
		Uvi        float64 `json:"uvi"`
		Clouds     int     `json:"clouds"`
		Visibility int     `json:"visibility"`
		WindSpeed  float64 `json:"wind_speed"`
		WindDeg    int     `json:"wind_deg"`
		WindGust   float64 `json:"wind_gust"`
		Weather    []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Rain struct {
			OneH float64 `json:"1h"`
		} `json:"rain"`
		Snow struct {
			OneH float64 `json:"1h"`
		} `json:"snow"`
	} `json:"current"`
	Alerts []struct {
		SenderName  string   `json:"sender_name"`
		Event       string   `json:"event"`
		Start       int      `json:"start"`
		End         int      `json:"end"`
		Description string   `json:"description"`
		Tags        []string `json:"tags"`
	} `json:"alerts"`
}
