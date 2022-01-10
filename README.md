# wxo
Command line weather grabber

`wxo` provides a command line utility for retrieving "real-time-ish" current
weather conditions for a given geography, formatting the results in a manner
suitable for use in minimalist window manager status bars like `dwm`. Example
output for Vancouver today:

    Overcast Clouds 2.5C ↖NNW 3.2km/h

If a weather alert exists, such as this spot in Ontario:

    !Extreme Cold! Clear Sky -27.9C ↖WNW 29.6km/h

There can be zero or more weather alerts; currently the tool simply
concatenates them. Here's an example from today for the Landes region in France
at -lat 44.050505 -long -0.893669:

    !Moderate Coastalevent Warning/Extreme Flooding Warning! Overcast Clouds 12.5C ↖NW 15.8km/h

Results are cached and expire in 5 minutes; this to avoid overwhelming weather
data providers and to remain within usage limits.

## Installation & Usage

Get the latest:

    go install github.com/solutionroute/wxo/cmd/wxo@latest

Using `wxo` requires a free account and API key from
[OpenWeathermap.org](https://openweathermap.org/) (OWM).  You *must* provide the
WXO_APIKEY environment variable, either on the command line, or as part of your
permanent environment. Example:

    $ WXO_APIKEY=yoursecretkey wxo -lat 49.123 -long -123.78

_Note: OWM informs new subscribers there may be delay (measured in minutes)
until the new account's API keys are usable._

**Coming soon**: 

* user configurable templates; you can use the output for other purposes.
* override cache timeout (5 minutes); -force option
* support for output in other languages.

## Current Data Sources 

It's my intent to add more options for weather data; an Environment Canada
module will be provided soon. Currently `wxo` supports:

* [OpenWeathermap.org](https://openweathermap.org/) via "One Call" API; you'll
  need a free API key. The free tier provides 1,000 "One Call" requests a day.

## Motivation

As a runner and outdoorsy person in general, weather is important to me! I like
seeing current temp, winds, and alerts on my `dwm` status bar; here's a relevant
excerpt from my [goblocks](https://github.com/Stargarth/Goblocks) config file:

    "actions":
    [
        {
            "prefix": "",
            "updateSignal": "36",
            "command": "WXO_APIKEY=myOWMkey /home/mw/go/bin/wxo -lat 49.2605 -long -123.1133 -units metric",
            "suffix": "",
            "timer": "10m"
        },
        ...
    ]

Also looked at:

* [wttr.in](https://wttr.in/) often used for for status bar updates via `curl`,
  but I found it often returned incorrect or wildly out of date information for
  my area, in addition to lacking weather alert data.
* Once an hour updates from Environment Canada were not cutting it, but I do
  appreciate their alerts & warnings. Writing the client for Open Weather Map
  (JSON data) was easy in comparison with dealing with EC's XML feed.

## Further Reading

While looking for weather APIs I stumbled across this detailed evaluation:

See: https://www.sigmdel.ca/michel/program/fpl/yweather/weather_api_en.html

