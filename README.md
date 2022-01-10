# wxo
Command line weather grabber

**January 9, 2022: In release early, release often mode.**

`wxo` provides a command line utility to retrieve current weather conditions and
format them in a manner suitable for use in status bars for window managers like
`dwm`. 

A common weather vocabulary is defied by the `SiteData` object; custom fetch-ers
for each weather source implementation map the provided data on to a SiteData object.

## Installation & Usage

    wxo URI -force

    (-force overrides cache expiry time)

Results for Vancouver, Canada:

    # Using openweathermap.org api (Details: https://openweathermap.org/current)
    # Note: calling by city ID is recommended by OW.
    # City codes are available here: http://bulk.openweathermap.org/sample/

    wxo "api.openweathermap.org/data/2.5/weather?id=6173331&units=metric&appid=yourlongapikey"

    # Note the extra comma, 'state' is only supported for US locations
    wxo "api.openweathermap.org/data/2.5/weather?q=Vancouver,,CA&units=metric&appid=yourlongapikey"

    # Vancouver, WA, USA
    wxo "api.openweathermap.org/data/2.5/weather?q=Vancouver,OR,US&units=metric&appid=yourlongapikey"




## Current functionality

* Data sources:
    * [OpenWeathermap.org]() JSON REST API; you'll need a free API key.
    * Others down the road (Environment Canada written not yet integrated).
* Caching of results set by default to a 15 minute expiry - that's 96 weather inquiries a day, perhaps enough for you?

## Coming soon

* User configurable template for output
* Temperature format choice of C or F or K
* For Canadian locations, a means of blending EC alerts with OpenWeathermap data

## Motivation

* [wttr.in]() often used for this purpose, shows incorrect data for my area
* Once an hour updates from Evironment Canada were not cutting it, but I do appreciate their alerts & warnings

curl "api.openweathermap.org/data/2.5/weather?id=6173331&appid=a2dbe40e6a2c2db83ad56f59d2f2b62c"


## Available WX APIs

See: https://www.sigmdel.ca/michel/program/fpl/yweather/weather_api_en.html

XXX include EC links