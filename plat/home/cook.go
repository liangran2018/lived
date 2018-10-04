package home

func cook() *outputBuild {
	opb := &outputBuild{IsUpdate:true}

	this := obl.Own[fire]
	fire := homeBuilding[fire]
	if this.Lvl == fire.maxlvl {
		opb.IsUpdate = false
	}

	opb.DurPercent = this.Dur/fire.b[this.Lvl].maxdur * 100
	opb.Action = make(map[action]bool, cookAnimalBlood - cookBbq + 1)

	for i:= cookBbq; i<= cookAnimalBlood; i++ {
		opb.Action[i] = false

		if i.Lvl() >= this.Lvl {
			opb.Action[i] = true
		}
	}

	return opb
}
