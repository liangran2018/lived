package home

import (
	"github.com/liangran2018/lived/human"
	"github.com/liangran2018/lived/env"
	"github.com/liangran2018/lived/log"
	"github.com/liangran2018/lived/base"
)

func sleep() *outputBuild {
	opb := &outputBuild{}

	this := obl.Own[bed]
	bed := homeBuilding[bed]

	if this.Lvl < bed.maxlvl {
		opb.IsUpdate = true
	}

	if this.Lvl != 0 {
		opb.DurPercent = this.Dur/bed.b[this.Lvl].maxdur * 100
		opb.Action = map[action]bool {sleep1H:true, sleep4H:true, sleep8H:true}
	} else {
		opb.Action = map[action]bool {sleep1H:false, sleep4H:false, sleep8H:false}
	}

	return opb
}

func sleepAction(i action) int {
	switch i {
	case sleep1H:
		human.GetHuman().SleepingChangePerHour()
		env.GetTime().Add(1, 0)
		log.GetLogger().Log(log.Info, "Sleep1H")
	case sleep4H:
		for i:=0; i<4; i++ {
			human.GetHuman().SleepingChangePerHour()
		}
		env.GetTime().Add(4, 0)
		log.GetLogger().Log(log.Info, "Sleep4H")
	case sleep8H:
		for i := 0; i < 8; i++ {
			human.GetHuman().SleepingChangePerHour()
		}
		env.GetTime().Add(8, 0)
		log.GetLogger().Log(log.Info, "Sleep8H")
	default:
		log.GetLogger().Log(log.Wrong, "bedChooseErr", i)
		return base.ParaInvalid
	}

	return 0
}
