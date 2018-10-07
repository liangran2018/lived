package home

import (
	"github.com/liangran2018/lived/env"
	"github.com/liangran2018/lived/materiel"
)

var CleanTime struct{
	a action
	t int
	b bool
}

func clean() *outputBuild {
	opb := &outputBuild{}

	this := obl.Own[water]
	water := homeBuilding[water]

	if this.Lvl == 0 {
		opb.IsUpdate = true
	} else {
		opb.DurPercent = this.Dur/water.b[this.Lvl].maxdur * 100
	}

	opb.Action = make(map[action]int, makeWine - distill + 1)

	for i:= distill; i<= makeWine; i++ {
		if CleanTime.b {
			opb.Action[i] = other
		} else {
			opb.Action[i] = lvlNotEnough

			if i.Lvl() <= this.Lvl {
				opb.Action[i] = ok
			}
		}
	}

	if CleanTime.b {
		opb.Action[CleanTime.a] = busy
	}

	return opb
}

func cleanAction(i action) int {
	if CleanTime.b {
		return 0
	}

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
