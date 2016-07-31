package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/damog/mastodon"
)

type StatusSource func(*mastodon.Config) *mastodon.StatusInfo
type JsonAny interface{}

type ClickMessage struct {
	Name     string `json:"name,omitempty"`
	Instance string `json:"instance,omitempty"`
	Button   int    `json:"button"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
}

var Modules = map[string]StatusSource{
	"battery":  mastodon.Battery,
	"clock":    mastodon.Clock,
	"cpu":      mastodon.CPU,
	"disk":     mastodon.Disk,
	"hostname": mastodon.Hostname,
	"ip":       mastodon.IPAddress,
	"loadavg":  mastodon.LoadAvg,
	"memory":   mastodon.Memory,
	"uptime":   mastodon.Uptime,
	"weather":  mastodon.Weather,
}

func PrintHeader() {
	fmt.Println("{\"version\":1, \"click_events\": true}")
	fmt.Println("[")
}

func LoadConfig() *mastodon.Config {
	config := mastodon.NewConfig()
	config.ApplyXresources()
	config.ReadConfig()
	config.ParseTemplates()
	return config
}

func main() {
	config := LoadConfig()
	duration := config.ReadInterval()

	module_names := strings.Split(config.Data["order"], ",")
	for _, module_name := range module_names {
		if _, ok := config.Data[module_name]; !ok {
			config.Data[module_name] = "color_normal"
		}
	}

	jsonArray := make([]map[string]JsonAny, len(module_names))

	PrintHeader()

	// start new goroutine to handle Stdin
	go func() {
		f, _ := os.Create("/home/thomas/barresult")
		f.WriteString("...\n")
		for {
			bio := bufio.NewReader(os.Stdin)
			line, _, _ := bio.ReadLine()

			var m ClickMessage

			if line[0] == ',' {
				line = line[1:]
			}

			err := json.Unmarshal(line, &m)
			if err != nil {
				f.WriteString(err.Error())
			} else {
				if r, ok := config.Data["onclick_"+m.Name]; ok {
					split := strings.Split(r, " ")
					err := exec.Command(split[0], split[1:]...).Start()
					if err != nil {
						f.WriteString(err.Error())
					}
				}

			}

			time.Sleep(1)
		}
	}()

	for {
		for idx, module_name := range module_names {
			si := Modules[module_name](config)
			color := config.Data[module_name]
			if si.IsBad() {
				color = config.Data["color_bad"]
			}
			if _, ok := config.Data[color]; ok {
				color = config.Data[color]
			}
			border_top := 0
			border_left := 0
			border_bottom := 0
			border_right := 0
			border_color := "#ff0000"

			if b, ok := config.Data["border_"+module_name]; ok {
				s := strings.Split(b, " ")
				var err error
				border_top, err = strconv.Atoi(s[0])
				if err != nil {
					fmt.Println("border_" + module_name + "  are not numbers")
					return
				}
				border_right, err = strconv.Atoi(s[1])
				if err != nil {
					fmt.Println("border_" + module_name + "  are not numbers")
					return
				}
				border_bottom, err = strconv.Atoi(s[2])
				if err != nil {
					fmt.Println("border_" + module_name + "  are not numbers")
					return
				}
				border_left, err = strconv.Atoi(s[3])
				if err != nil {
					fmt.Println("border_" + module_name + "  are not numbers")
					return
				}
				border_color = s[4]
			}

			jsonArray[idx] = map[string]JsonAny{
				"full_text":     si.FullText,
				"color":         color,
				"border":        border_color,
				"border_top":    border_top,
				"border_right":  border_right,
				"border_bottom": border_bottom,
				"border_left":   border_left,
				"name":          module_name,
			}
		}

		jsonData, _ := json.Marshal(jsonArray)
		fmt.Print(string(jsonData))
		fmt.Printf(",\n")
		time.Sleep(duration)
	}
}
