package home

import (
	"github.com/liangran2018/lived/env"
	"github.com/liangran2018/lived/materiel"
)

func clean() *outputBuild {
	opb := &outputBuild{}

	this := obl.Own[water]
	water := homeBuilding[water]

	if this.Lvl == 0 {
		opb.IsUpdate = true
	} else {
		opb.DurPercent = this.Dur/water.b[this.Lvl].maxdur * 100
	}

	opb.Action = make(map[action]bool, makeWine - distill + 1)

	for i:= distill; i<= makeWine; i++ {
		opb.Action[i] = false

		if i.Lvl() <= this.Lvl {
			opb.Action[i] = true
		}
	}

	return opb
}

func cleanAction(i action) int {
	time := actionNature[i].time

	for k, v := range actionNature[i].m {
		materiel.GetOwnThings().PlusProduct(k, v)
	}

	env.GetTime().Add(time.H, time.Mi)

	//for k, v := range actionNature[i].get {
	//	materiel.GetOwnThings().AddProduct(k, v)
	//}

	return actionNature[i].delay
}
