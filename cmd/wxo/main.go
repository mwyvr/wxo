package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/solutionroute/wxo/internal"
	"github.com/solutionroute/wxo/owm"
)

var (
	version   string = "0.0.1"
	units     string = "metric"
	lattitude float64
	longitude float64
	apiKey    string
)

const apiKeyVarName = "WXO_APIKEY"

func init() {
	apiKey = os.Getenv(apiKeyVarName) // 12-factor(ish), eh?
	flag.StringVar(&units, "units", units, "Units preference (metric, imperial, kelvin)")
	flag.Float64Var(&lattitude, "lat", lattitude, "*Lattitude of desired weather site")
	flag.Float64Var(&longitude, "long", longitude, "*Longitude of desired weather site")
	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "wxo (v%s) ", version)
		if apiKey == "" {
			fmt.Fprintf(os.Stdout, "Fatal: %s missing from environment.", apiKeyVarName)
		}
		fmt.Fprintf(os.Stdout, "\nMinimal correct use:\n\n  %s=yoursecretkey wxo -lat 49.123 -long -123.78\n\n", apiKeyVarName)
		flag.PrintDefaults()
	}
	// ensure the cache path exists
	err := os.MkdirAll(internal.GetCachePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.Parse()
	units = strings.ToLower(units)
	if units == "metric" || units == "imperial" || units == "kelvin" {
		apikey := os.Getenv(apiKeyVarName)
		if apikey == "" || lattitude == 0 || longitude == 0 {
			flag.Usage()
		}
	} else {
		flag.Usage()
	}

	client := owm.NewWeatherClient(apiKey, lattitude, longitude, units)
	// client := owm.NewWeatherClient("a2dbe40e6a2c2db83ad56f59d2f2b62c", 49.26062471888604, -123.11459123540007)
	wx, err := client.Fetch()
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		flag.Usage()
		os.Exit(1)
	}
	wx.Print()
}
