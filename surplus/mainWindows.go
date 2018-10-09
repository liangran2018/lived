package surplus

import (
	"github.com/liangran2018/lived/base"
	"github.com/liangran2018/lived/human"
	"github.com/liangran2018/lived/env"
	"github.com/liangran2018/lived/plat/home"

	"github.com/gin-gonic/gin"
)

type mainPagePara struct {
	User       *human.Human  `json:"user"`
	Time       *env.GameTime `json:"time"`
	Season     string        `json:"season"`
	Weather    string        `json:"weather"`
	Temprature int           `json:"temprature"`
	Building   []int 		 `json:"building"`
}

func fillPara() *mainPagePara {
	p := &mainPagePara{}
	p.User = human.GetHuman()
	p.Time = env.GetTime()
	p.Season = env.GetSeason().Name()
	p.Weather = env.GetWeather().Name()
	p.Temprature = env.GetTemp()
	p.Building = home.GetOwnBuild()

	return p
}

func MainWin(c *gin.Context) {
	base.Output(c, 0, fillPara())
	return
}