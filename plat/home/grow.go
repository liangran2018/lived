package home

import (
	"github.com/liangran2018/lived/materiel"
	"github.com/liangran2018/lived/env"
)

func grow() *outputBuild {
	opb := &outputBuild{}

	this := obl.Own[field]
	field := homeBuilding[field]

	if this.Lvl < field.maxlvl {
		opb.IsUpdate = true
	}

	if this.Lvl != 0 {
		opb.DurPercent = this.Dur/field.b[this.Lvl].maxdur * 100
	}

	opb.Action = make(map[action]bool, growMint - growRise + 1)

	for i:= growRise; i<= growMint; i++ {
		opb.Action[i] = false

		if i.Lvl() <= this.Lvl {
			opb.Action[i] = true
		}
	}

	return opb
}

func growAction(i action) int {
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
