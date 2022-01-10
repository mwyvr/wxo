# wxo
Command line weather grabber

**January 9, 2022: In release early, release often mode.**

`wxo` provides a command line utility that retrieves current weather conditions for a given geography,
formatting the results in a manner suitable for use in window manager status bars like
`dwm`. Example output:

    Clear Sky 2.5C ↖NNW 3.2km/

Coming soon: user configurable templates; you can use the output for other purposes.

## Usage

You need a free account and API key from
[OpenWeathermap.org](https://openweathermap.org/). It may take 5 - 30 minutes
for the key to become active. You *must* provide the WXO_APIKEY environment
variable wither on the command line or as part of your permanent environment.
At that point:

    $ WXO_APIKEY=yoursecretkey wxo -lat 49.123 -long -123.78

Delivers (at this writing):

    Clear Sky 2.5C ↖NNW 3.2km/

Required additional parameters include where you live or what to see weather
conditions for:

    -lat float
        *Lattitude of desired weather site
    -long float
        *Longitude of desired weather site

Optional:

    -units string
        Units preference (metric, imperial, kelvin) (default "metric")

## Current functionality

* Data sources:
    * [OpenWeathermap.org](https://openweathermap.org/) JSON REST API; you'll need a free API key.
    * Others down the road (Environment Canada written not yet integrated).
* Caching of results set by default to a 5 minute expiry

## Motivation

* [wttr.in]() often used for for status bar updates,  shows incorrect data for my area
* Once an hour updates from Environment Canada were not cutting it, but I do appreciate their alerts & warnings

## Available WX APIs

See: https://www.sigmdel.ca/michel/program/fpl/yweather/weather_api_en.html

