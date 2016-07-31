## Mastodon

Status bar replacement for [i3status](http://i3wm.org/i3status/). Works with
`i3bar`.

The goals of this project are:

* More configurable than i3status.
* Use colors from Xresources by default.
* Learn golang.

## Config
### Default config
```
bar_size:          10,
color_bad:         #d00000,
color_good:        #00d000,
color_normal:      #cccccc,
date_format:       2006-01-02 15:04:05,
format_battery:    {{if .battery}}{{.prefix}} {{.bar}} ({{.remaining}} {{.wattage}}W){{else}}No battery{{end}},
format_clock:      {{.time}},
format_cpu:        C {{.bar}},
format_disk:       D {{.bar}},
format_hostname:   {{.hostname}},
format_ip:         {{.ip}},
format_loadavg:    {{.fifteen}} {{.five}} {{.one}},
format_memory:     R {{.bar}},
format_uptime:     {{.uptime}},
format_weather:    {{if .error}}{{.error}}{{ else }}{{.today}} {{.high}}/{{.low}} ({{.next}}){{ end }},
interval:          1,
network_interface: wlan0,
order:             weather,cpu,memory,disk,battery,ip,loadavg,clock,
bar_start:         [,
bar_end:           ],
bar_empty:          ,
bar_full:          #,
```

You can save a `mastodon.conf` file in `.config/` confaining your custom config options, like:

```
format_disk=  {{.bar}}
```


### Borders 
Following the following syntax:
```
border_<module>: <top> <right> <bottom> <left> <color>
```

Where module is the name of the module used in `order`, and top, right bottom and left are the width of the line (default to 0). Color is obviously the color.

Example: 
```
border_cpu=0 0 2 0 #ffffff
border_memory=0 0 2 0 #00ff00
```

### How does it look?

![](http://i.imgur.com/1QOCIRR.png)

The above screenshot shows some of the default modules:

* Weather
* CPU
* Memory usage
* HDD usage
* Battery remaining
* IP address
* Load average
* Clock

### Example Config

![](http://i.imgur.com/3oyWLiG.png)

The above screenshot shows another example, using the config options below:

```
order=cpu,memory,disk,ip,loadavg,uptime,clock

# Formats
format_cpu=  {{.bar}}
format_memory=  {{.bar}}
format_disk=  {{.bar}}
format_ip=  {{.ip}}
format_loadavg=  {{.one}} {{.five}} {{.fifteen}}
format_uptime=  {{.uptime}}
format_clock=  {{.time}}

# Borders
border_cpu=0 0 2 0 #ffffff
border_memory=0 0 2 0 #00ff00
border_disk=0 0 2 0 #ffff00
border_ip=0 0 2 0 #00ffff
border_uptime=0 0 2 0 #ffffff
border_loadavg=0 0 2 0 #ff0000
border_clock=0 0 2 0 #909090

bar_empty=□
bar_full=■
bar_start=
bar_end=

onclick_cpu=gnome-system-monitor -p
onclick_memory=gnome-system-monitor -r
onclick_disk=gnome-system-monitor -f

interval=2
```