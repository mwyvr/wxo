# wxo
Command line weather grabber

**January 9, 2022: In release early, release often mode.**

`wxo` provides a command line utility to retrieve current weather conditions and
format them in a manner suitable for use in status bars for window managers like
`dwm`. 

A common weather vocabulary is defined by the `SiteData` object; custom fetch-ers
for each weather source can map provided data on to a SiteData object.

Being a Canadian, Environment Canada's (sadly) hourly data will soon be supported.

## Usage


## Current functionality

* Data sources:
    * [OpenWeathermap.org]() JSON REST API; you'll need a free API key.
    * Others down the road (Environment Canada written not yet integrated).
* Caching of results set by default to a 5 minute expiry

## Coming soon

* User configurable template for output

## Motivation

* [wttr.in]() often used for this purpose, shows incorrect data for my area
* Once an hour updates from Environment Canada were not cutting it, but I do appreciate their alerts & warnings

curl "api.openweathermap.org/data/2.5/weather?id=6173331&appid=a2dbe40e6a2c2db83ad56f59d2f2b62c"


## Available WX APIs

See: https://www.sigmdel.ca/michel/program/fpl/yweather/weather_api_en.html

XXX include EC links
