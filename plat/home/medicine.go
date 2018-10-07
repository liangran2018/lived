package home

func medicine() *outputBuild {
	opb := &outputBuild{}

	this := obl.Own[drug]
	drug := homeBuilding[drug]

	if this.Lvl < drug.maxlvl {
		opb.IsUpdate = true
	}

	if this.Lvl != 0 {
		opb.DurPercent = this.Dur/drug.b[this.Lvl].maxdur * 100
	}

	opb.Action = make(map[action]int, makeTonifyPill - makeBandage + 1)

	for i:= makeBandage; i<= makeTonifyPill; i++ {
		opb.Action[i] = lvlNotEnough

		if i.Lvl() <= this.Lvl {
			opb.Action[i] = ok
		}
	}

	return opb
}
