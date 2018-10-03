package env

import (
	"time"
	"math/rand"
)

var todayTemp int

func NewTempToday() {
	t := tempraturePro[int(GetSeason())][GetWeather()]
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(t.max - t.min)
	todayTemp = randNum + t.min
}

func LoadTemp(i int) {
	todayTemp = i
}

func GetBaseTemp() int {
	return todayTemp
}

func GetTemp() int {
	if gmTime.Hour >= 0 && gmTime.Hour < 3 {
		return todayTemp
	}

	if gmTime.Hour >= 3 && gmTime.Hour < 6 {
		return todayTemp+1
	}

	if gmTime.Hour >= 6 && gmTime.Hour < 9 {
		return todayTemp+3
	}

	if gmTime.Hour >= 9 && gmTime.Hour < 12 {
		return todayTemp+4
	}

	if gmTime.Hour >= 12 && gmTime.Hour < 15 {
		return todayTemp+6
	}

	if gmTime.Hour >= 15 && gmTime.Hour < 18 {
		return todayTemp+4
	}

	if gmTime.Hour >= 18 && gmTime.Hour < 21 {
		return todayTemp+3
	}

	return todayTemp+1
}
