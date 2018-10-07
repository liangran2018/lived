package home

import (
	"github.com/liangran2018/lived/base"
	"github.com/liangran2018/lived/env"
	"github.com/liangran2018/lived/materiel"
	"github.com/liangran2018/lived/explore"
)

func equip() *outputBuild {
	opb := &outputBuild{}

	this := obl.Own[tool]
	tool := homeBuilding[tool]

	if this.Lvl < tool.maxlvl {
		opb.IsUpdate = true
	}

	if this.Lvl != 0 {
		opb.DurPercent = this.Dur/tool.b[this.Lvl].maxdur * 100
	}

	opb.Action = make(map[action]int, makeRattanArmor - cookBbq + 1)

	for i:= makeStoneAxe; i<= makeRattanArmor; i++ {
		opb.Action[i] = lvlNotEnough

		if i.Lvl() <= this.Lvl {
			opb.Action[i] = ok
		}
	}

	if BigBag {
		opb.Action[makeBigBag] = only
	}

	return opb
}

func equipAction(i action) int {
	if BigBag && i == makeBigBag {
		return base.ParaInvalid
	}

	time := actionNature[i].time

	for k, v := range actionNature[i].m {
		materiel.GetOwnThings().PlusProduct(k, v)
	}

	if i == makeBigBag {
		explore.LoadBag()
		env.GetTime().Add(time.H, time.Mi)
		return 0
	}

	for k, v := range actionNature[i].get {
		materiel.GetOwnThings().AddProduct(k, v)
	}

	env.GetTime().Add(time.H, time.Mi)

	return 0
}
