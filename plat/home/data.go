package home

import (
	"github.com/liangran2018/lived/env"
	"github.com/liangran2018/lived/human"
)

type mainPagePara struct {
	User       *human.Human  `json:"user"`
	Time       *env.GameTime `json:"time"`
	Season     string        `json:"season"`
	Weather    string        `json:"weather"`
	Temprature int           `json:"temprature"`
}

func fillPara() *mainPagePara {
	p := &mainPagePara{}
	p.User = human.GetHuman()
	p.Time = env.GetTime()
	p.Season = env.GetSeason().Name()
	p.Weather = env.GetWeather().Name()
	p.Temprature = env.GetTemp()

	return p
}
