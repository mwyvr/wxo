# wxo
Command line weather grabber

`wxo` provides a command line utility for retrieving "real-time-ish" current
weather conditions for a given geography, formatting the results in a manner
suitable for use in minimalist window manager status bars like `dwm`. All output
is text, there are no icons or colours to make your minimalist life cluttered.

Results are cached and expire in 5 minutes; this to avoid overwhelming weather
data providers and to remain within usage limits.

Example output for Vancouver (BC, Canada) today:

    Overcast Clouds 2.5C ↖NNW 3.2km/h

If a weather alert exists, such as this spot in Ontario, Canada:

    !Extreme Cold! Clear Sky -27.9C ↖WNW 29.6km/h

Or flooding in that other Vancouver, in WA:

    !Flood Advisory! Overcast Clouds 40.3F ←W 3.6mph

At any given point a report will contain zero or more of these weather alerts;
the tool concatenates and truncates them if necessary. Here's a truncated
example for the Landes region in France at -lat 44.050505 -long -0.893669:

    !Moderate Rain-Flood Warning/Moderate Coa...! Overcast Clouds 10.5C ↗ENE 5.3km/h

To the extent the provider offers multilingual responses, access your preferred
language with the `-lang xx` flag.

    # output example for Vielha, Spain
    !Moderate Wind Warning/Moderate Avalanche...! Overcast Clouds 4.7C ↑N 12.3km/h
    #  wxo -lat 42.701287 -long 0.793591 -lang es
    es: !Moderate Wind Warning/Moderate Avalanche...! Nubes 4.7C ↑N 12.3km

As you can see, alert text is not internationalized by the provider.

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

Incidentally, I have a hot-key defined to issue a shell command `kill -36
$(pidof goblocks)`, causing `goblocks` to run my weather "block" and update my
weather status line, subject to the caching built (by default 5 minutes) into
`wxo`.

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
