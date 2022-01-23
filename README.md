# wxo
`wxo` is primarily a command line utility for retrieving "real-time-ish" current
weather conditions for a given location.

The default output to `stdout` is formatted for use in minimalist window manager
status bars such as [dwmblocks][3] or [Goblocks][4] on [dwm][1]. All output is
text (UTF-8 text); there are no icons or colours to make your minimalist life
cluttered.

Results are cached and expire by default in 5-10 minutes; expiry depends on the
wx data provider's daily usage limits.

## Installation & Usage

Get the latest:

    go install github.com/solutionroute/wxo/cmd/wxo@latest

Using `wxo` with [OpenWeathermap.org][4] requires a free account and API key.

Run `wxo` for a usage screen. Note: You *must* provide the WXO_APIKEY environment
variable, either on the command line, or as part of your permanent environment.
Example:

    $ WXO_APIKEY=yoursecretkey wxo -lat 49.123 -long -123.78

![in use on dwm with Goblocks](https://raw.githubusercontent.com/solutionroute/wxo/main/doc/20220122-212612.png)
_[My dwm config][2] feeding wxo output to [Goblocks][3] status bar_

Example output without a weather alert:

    Overcast Clouds 2.5C ↖NNW 3.2km/h

Some weather alerts:

    !Extreme Cold! Clear Sky -27.9C ↖WNW 29.6km/h
    !Flood Advisory! Overcast Clouds 40.3F ←W 3.6mph

Multiple alerts are concatenated and may be truncated (`...`):

    !Moderate Rain-Flood Warning/Moderate Coa...! Overcast Clouds 10.5C ↗ENE 5.3km/h

Limited internationalization is available and is weather data provider dependent. Alert text may not be available in other languages.

    # Vielha, Spain
    #  wxo -lat 42.701287 -long 0.793591
    en: !Moderate Wind Warning/Moderate Avalanche...! Overcast Clouds 4.7C ↑N 12.3km/h
    #  wxo -lat 42.701287 -long 0.793591 -lang es
    es: !Moderate Wind Warning/Moderate Avalanche...! Nubes 4.7C ↑N 12.3km

### As a library

While it wasn't written for the purpose, the package could be used as a library
for accessing various weather data sources. The wxo cmd `main.go` provides an
example:

    client := owm.NewWeatherClient(apiKey, latitude, longitude, units, lang)
	wx, err := client.Fetch() // returns a SiteData object

### Weather Sources

Currently `wxo` supports:

* [OpenWeathermap.org][5] via their "One Call" API; you'll need a free API key.
  The free tier provides 1,000 "One Call" requests a day.

## Motivation

As a distance runner, weather is important to me! As a weather geek, I don't
need an excuse. I want to see seeing current temp, winds, and alerts on my `dwm`
status bar; here's a relevant excerpt from my [Goblocks][3] config file:

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

A hot-key issuing a shell command `kill -36 $(pidof goblocks)`, causes
`goblocks` to run the weather "block" and update the weather status line,
subject to the default cache expiry (5 minutes).

**Other Solutions**:

* [wttr.in](https://wttr.in/) is used by many for status bar updates via `curl`,
  but I found the data returned was often partially incorrect (wind info in
  particular) or quite out of date, in addition to lacking weather alert data.

## TODO

`wxo` was born on January 8th, 2022, and will continue to evolve with a few more
features--in rough order of priority:

* ~~Cache output~~
* Specify cache expiry, subject to weather data provider daily limits
* Custom output templates
* Other weather data providers


[1]: <https://dwm.suckless.org/>
[2]: <https://github.com/solutionroute/suckless> "My dwm config and other patches"
[3]: <https://github.com/Stargarth/Goblocks>
[4]: <https://github.com/torrinfail/dwmblocks>
[5]: <https://openweathermap.org/> 
