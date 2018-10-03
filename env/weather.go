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

type weather int

var todayWeather weather

const (
	sunny weather = iota
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

func NewWeather() {
	m := weatherPro[int(GetSeason())]
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(100)

	for k, v := range m {
		if v.min <= i && i <= v.max {
			todayWeather = k
		}
	}
}

func LoadWeather(w int) {
	todayWeather = weather(w)
}

func GetWeather() weather {
	return todayWeather
}

func MoodChangePerDay() {
	human.GetHuman().MoodChange(swMood[GetSeason()][GetWeather()])
}

func (this weather) Name() string {
	return weatherName[int(this)]
}
