package home

import (
	"github.com/liangran2018/lived/env"
	"github.com/liangran2018/lived/materiel"
)

func rail() *outputBuild {
	opb := &outputBuild{}

	this := obl.Own[fence]
	fence := homeBuilding[fence]

	if this.Lvl < fence.maxlvl {
		opb.IsUpdate = true
	}

	if this.Lvl != 0 {
		opb.DurPercent = this.Dur/fence.b[this.Lvl].maxdur * 100
	}

	opb.Action = make(map[action]bool, 0)

	return opb
}

func railAction(i action) int {
	time := actionNature[i].time

	for k, v := range actionNature[i].m {
		materiel.GetOwnThings().PlusProduct(k, v)
	}

	env.GetTime().Add(time.H, time.Mi)

	return 0
}
