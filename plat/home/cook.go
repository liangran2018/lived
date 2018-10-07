package home

import (
	"github.com/liangran2018/lived/env"
	"github.com/liangran2018/lived/materiel"
)

func cook() *outputBuild {
	opb := &outputBuild{}

	this := obl.Own[fire]
	fire := homeBuilding[fire]

	if this.Lvl < fire.maxlvl {
		opb.IsUpdate = true
	}

	if this.Lvl != 0 {
		opb.DurPercent = this.Dur/fire.b[this.Lvl].maxdur * 100
	}

	opb.Action = make(map[action]int, cookAnimalBlood - cookBbq + 1)

	for i:= cookBbq; i<= cookAnimalBlood; i++ {
		opb.Action[i] = lvlNotEnough

		if i.Lvl() <= this.Lvl {
			opb.Action[i] = ok
		}
	}

	return opb
}

func commonAction(i action) int {
	time := actionNature[i].time

	for k, v := range actionNature[i].m {
		materiel.GetOwnThings().PlusProduct(k, v)
	}

	for k, v := range actionNature[i].get {
		materiel.GetOwnThings().AddProduct(k, v)
	}

	env.GetTime().Add(time.H, time.Mi)

	return 0
}
