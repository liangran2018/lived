package home

import (
	"github.com/liangran2018/lived/env"
	"github.com/liangran2018/lived/materiel"
)

func fishing() *outputBuild {
	opb := &outputBuild{}

	this := obl.Own[field]
	field := homeBuilding[field]

	opb.Action = make(map[action]int, 1)

	if this.Lvl == 0 {
		opb.IsUpdate = true
		opb.Action[goFishing] = lvlNotEnough
	} else {
		opb.DurPercent = this.Dur/field.b[this.Lvl].maxdur * 100
		opb.Action[goFishing] = ok
	}

	return opb
}

func fishAction(i action) int {
	materiel.GetOwnThings().PlusProduct(materiel.Meat, 1)

	materiel.GetOwnThings().AddProduct(materiel.Fish, 1)

	env.GetTime().Add(0, 30)

	return 0
}
