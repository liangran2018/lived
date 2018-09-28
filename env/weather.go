package env

import (
	"time"
	"math/rand"

	"github.com/liangran2018/lived/human"
)

type wl struct {
	min int
	max int
}

var weather []map[int]wl
var swMood []map[int]int
var todayWeather int

const (
	sunny = iota
	cloudy
	foggy
	overcast
	windy
	lightRainy
	heavyRainy
	lightSnowy
	heavySnowy
	storm
)

func init() {
	weather = make([]map[int]wl, 4)
	weather[spring] = map[int]wl {
		sunny:wl{min:0, max:49},
		cloudy:wl{min:50, max:59},
		foggy:wl{min:60, max:65},
		overcast:wl{min:66, max:73},
		windy:wl{min:74, max:81},
		lightRainy:wl{min:82, max:89},
		heavyRainy:wl{min:90, max:91},
		lightSnowy:wl{min:92, max:100},
	}
	weather[summer] = map[int]wl {
		sunny:wl{min:0, max:54},
		cloudy:wl{min:55, max:69},
		overcast:wl{min:70, max:74},
		windy:wl{min:75, max:78},
		lightRainy:wl{min:79, max:86},
		heavyRainy:wl{min:87, max:94},
		storm:wl{min:95, max:100},
	}
	weather[autumn] = map[int]wl {
		sunny:wl{min:0, max:64},
		cloudy:wl{min:65, max:74},
		overcast:wl{min:75, max:82},
		windy:wl{min:83, max:86},
		lightRainy:wl{min:87, max:96},
		heavyRainy:wl{min:97, max:100},
	}
	weather[winter] = map[int]wl {
		sunny:wl{min:0, max:42},
		cloudy:wl{min:43, max:62},
		foggy:wl{min:63, max:66},
		overcast:wl{min:67, max:70},
		windy:wl{min:71, max:76},
		lightSnowy:wl{min:77, max:86},
		heavySnowy:wl{min:87, max:94},
		storm:wl{min:95, max:100},
	}

	swMood = make([]map[int]int, 4)
	swMood[spring] = map[int]int {
		sunny:8,
		cloudy:5,
		foggy:0,
		overcast:0,
		windy:-1,
		lightRainy:-2,
		heavyRainy:-8,
		lightSnowy:-1,
	}
	swMood[summer] = map[int]int {
		sunny:-3,
		cloudy:0,
		overcast:5,
		windy:4,
		lightRainy:8,
		heavyRainy:-3,
		storm:-5,
	}
	swMood[autumn] = map[int]int {
		sunny:10,
		cloudy:6,
		overcast:0,
		windy:2,
		lightRainy:-3,
		heavyRainy:-8,
	}
	swMood[winter] = map[int]int {
		sunny:8,
		cloudy:0,
		foggy:0,
		overcast:-1,
		windy:-3,
		lightSnowy:-3,
		heavySnowy:-6,
		storm:-10,
	}
}

func NewWeather() {
	m := weather[GetSeason()]
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(100)

	for k, v := range m {
		if v.min <= i && i <= v.max {
			todayWeather = k
		}
	}
}

func GetWeather() int {
	return todayWeather
}

func MoodChangePerDay() {
	human.GetHuman().MoodChange(swMood[GetSeason()][GetWeather()])
}

func ShowWeather() string {
	switch GetWeather() {
	case sunny:
		return "晴天"
	case cloudy:
		return "多云"
	case foggy:
		return "雾"
	case overcast:
		return "阴天"
	case windy:
		return "起风"
	case lightRainy:
		return "小雨"
	case heavyRainy:
		return "大雨"
	case lightSnowy:
		return "小雪"
	case heavySnowy:
		return "大雪"
	default:
		return "暴风雨"
	}

	panic(GetWeather())
}
