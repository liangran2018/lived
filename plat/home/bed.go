package home

import (
	"github.com/liangran2018/lived/base"
	"github.com/liangran2018/lived/env"
	"github.com/liangran2018/lived/human"
	"github.com/liangran2018/lived/log"

	"github.com/gin-gonic/gin"
)

func sleep() *outputBuild {
	opb := &outputBuild{IsUpdate:true}

	this := obl.Own[bed]
	bed := homeBuilding[bed]
	if this.Lvl == bed.maxlvl {
		opb.IsUpdate = false
	}

	opb.DurPercent = this.Dur/bed.b[this.Lvl].maxdur * 100
	opb.Action = make(map[action]bool, sleep8H - sleep1H + 1)

	for i:= sleep1H; i<= sleep8H; i++ {
		opb.Action[i] = false

		if i.Lvl() >= this.Lvl {
			opb.Action[i] = true
		}
	}

	return opb
}

func BedChoose(c *gin.Context) {
	a := c.Query("action")
	if base.Empty(a) {
		base.Output(c, base.ParaInvalid, nil)
		return
	}

	switch a {
	case "0": //sleep1H
		human.GetHuman().SleepingChangePerHour()
		env.GetTime().Add(1, 0)
		log.GetLogger().Log(log.Info, "Sleep1H")
	case "1": //sleep4H
		for i:=0; i<4; i++ {
			human.GetHuman().SleepingChangePerHour()
		}
		env.GetTime().Add(4, 0)
		log.GetLogger().Log(log.Info, "Sleep4H")
	case "2": //sleep8H
		for i := 0; i < 8; i++ {
			human.GetHuman().SleepingChangePerHour()
		}
		env.GetTime().Add(8, 0)
		log.GetLogger().Log(log.Info, "Sleep8H")
	default:
		log.GetLogger().Log(log.Wrong, "bedChooseErr", a)
		base.Output(c, base.ParaInvalid, nil)
		return
	}
	
	base.Output(c, 0, fillPara())
	return
}

