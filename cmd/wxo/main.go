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
	units     string = "metric"
	lattitude float64
	longitude float64
	apiKey    string
)

const apiKeyVarName = "WXO_APIKEY"

func init() {
	apiKey = os.Getenv(apiKeyVarName) // 12-factor(ish), eh?
	flag.StringVar(&units, "units", units, "Units preference (metric, imperial, kelvin, scientific)")
	flag.Float64Var(&lattitude, "lat", lattitude, "*Latitude of desired weather site")
	flag.Float64Var(&longitude, "long", longitude, "*Longitude of desired weather site")
	flag.Usage = func() {
		fmt.Fprintf(os.Stdout,
			"\nMinimal correct use:\n\n  %s=yoursecretkey wxo -lat 49.123 -long -123.78\n\n", apiKeyVarName)
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
	valid_units := []string{"metric", "imperial", "kelvin", "scientific"}
	units = strings.ToLower(units)
	if units == "scientific" { // an alias
		units = "kelvin"
	}

	// api key must be provided in the environment; there is no flag for secrets
	apikey := os.Getenv(apiKeyVarName)
	if apikey == "" {
		fmt.Fprintf(os.Stdout, "missing required API key environment variable: %s\n", apiKeyVarName)
		flag.Usage()
		os.Exit(2)
	}

	// lat and long are required
	required := []string{"lat", "long"}
	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
	for _, req := range required {
		if !seen[req] {
			fmt.Fprintf(os.Stdout, "missing required flag: -%s\n", req)
			flag.Usage()
			os.Exit(2)
		}
	}

	// check optional units parameter is sane
	if !func(list []string, s string) bool {
		for _, v := range list {
			if v == s {
				return true
			}
		}
		return false
	}(valid_units, units) {
		fmt.Fprintf(os.Stdout, "invalid -units %s supplied; valid: %s\n", units, strings.Join(valid_units, ", "))
		os.Exit(2)
	}

	client := owm.NewWeatherClient(apiKey, lattitude, longitude, units)
	wx, err := client.Fetch()
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
	// TODO add user-supplied template option
	wx.Print()
}
