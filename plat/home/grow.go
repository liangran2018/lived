package home

import (
	"github.com/liangran2018/lived/materiel"
	"github.com/liangran2018/lived/env"
)

var GrowTime struct{
	a action
	t int
	b bool
}

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

	opb.Action = make(map[action]int, growMint - growRise + 1)

	for i:= growRise; i<= growMint; i++ {
		if GrowTime.b {
			opb.Action[i] = other
		} else {
			opb.Action[i] = lvlNotEnough

			if i.Lvl() <= this.Lvl {
				opb.Action[i] = ok
			}
		}
	}

	if GrowTime.b {
		opb.Action[GrowTime.a] = busy
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

	CleanTime.t = env.GetTimeInt().Time()
	CleanTime.b = true
	CleanTime.a = i

	return actionNature[i].delay
}
