# wxo
Command line weather grabber

`wxo` provides a command line utility for retrieving "real-time-ish" current
weather conditions for a given geography, formatting the results in a manner
suitable for use in minimalist window manager status bars like `dwm`. All output
is text, there are no icons or colours to make your minimalist life cluttered.

Results are cached and expire in 5 minutes; this to avoid overwhelming weather
data providers and to remain within usage limits.

## Installation & Usage

Get the latest:

    go install github.com/solutionroute/wxo/cmd/wxo@latest

Using `wxo` requires a free account and API key from
[OpenWeathermap.org](https://openweathermap.org/) (OWM); other weather providers
_may_ follow.  

Run `wxo` for a usage screen. Note: You *must* provide the WXO_APIKEY environment
variable, either on the command line, or as part of your permanent environment.
Example:

    $ WXO_APIKEY=yoursecretkey wxo -lat 49.123 -long -123.78

![in use on dwm with goblocks](https://raw.githubusercontent.com/solutionroute/wxo/main/doc/20220110-151745.png)
_In use on `dwm` with goblocks_

Example output without a weather alert:

    Overcast Clouds 2.5C ↖NNW 3.2km/h

Some weather alert types have more meaning:

    !Extreme Cold! Clear Sky -27.9C ↖WNW 29.6km/h
    !Flood Advisory! Overcast Clouds 40.3F ←W 3.6mph

Multiple alerts are concatenated and may be truncated (`...`):

    !Moderate Rain-Flood Warning/Moderate Coa...! Overcast Clouds 10.5C ↗ENE 5.3km/h

Some limited internationalization is available, weather data provider dependent:

    # Vielha, Spain
    #  wxo -lat 42.701287 -long 0.793591
    en: !Moderate Wind Warning/Moderate Avalanche...! Overcast Clouds 4.7C ↑N 12.3km/h
    #  wxo -lat 42.701287 -long 0.793591 -lang es
    es: !Moderate Wind Warning/Moderate Avalanche...! Nubes 4.7C ↑N 12.3km

**Coming soon**: 

* user configurable output templates
* override cache timeout (5 minutes); -force option
* another data provider, or two

As you can see, alert text is not internationalized by the provider.

## Current Data Sources 

It's my intent to add more options for weather data; an Environment Canada
module will be provided soon. Currently `wxo` supports:

* [OpenWeathermap.org](https://openweathermap.org/) via "One Call" API; you'll
  need a free API key. The free tier provides 1,000 "One Call" requests a day.

## Motivation

As a distance runner, weather is important to me! As a weather geek, I don't
need an excuse. I want to see seeing current temp, winds, and alerts on my `dwm`
status bar; here's a relevant excerpt from my
[goblocks](https://github.com/Stargarth/Goblocks) config file:

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

A hot-key defined to issue a shell command `kill -36 $(pidof goblocks)`, causes
`goblocks` to run the weather "block" and update the weather status line,
subject to the default cache expiry (5 minutes).

## Other Solutions

* [wttr.in](https://wttr.in/) is used by many for status bar updates via `curl`,
  but I found the data returned was often partially incorrect (wind info in
  particular) or quite out of date, in addition to lacking weather alert data.
