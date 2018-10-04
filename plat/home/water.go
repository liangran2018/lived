package home

func clean() *outputBuild {
	opb := &outputBuild{IsUpdate:true}

	this := obl.Own[water]
	water := homeBuilding[water]
	if this.Lvl == water.maxlvl {
		opb.IsUpdate = false
	}

	opb.DurPercent = this.Dur/water.b[this.Lvl].maxdur * 100
	opb.Action = make(map[action]bool, makeWine - distill + 1)

	for i:= distill; i<= makeWine; i++ {
		opb.Action[i] = false

		if i.Lvl() >= this.Lvl {
			opb.Action[i] = true
		}
	}

	return opb
}
