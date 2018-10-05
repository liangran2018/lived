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

	opb.Action = make(map[action]bool, makeTonifyPill - makeBandage + 1)

	for i:= makeBandage; i<= makeTonifyPill; i++ {
		opb.Action[i] = false

		if i.Lvl() <= this.Lvl {
			opb.Action[i] = true
		}
	}

	return opb
}
