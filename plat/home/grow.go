package home

func grow() *outputBuild {
	opb := &outputBuild{IsUpdate:true}

	this := obl.Own[field]
	field := homeBuilding[field]
	if this.Lvl == field.maxlvl {
		opb.IsUpdate = false
	}

	opb.DurPercent = this.Dur/field.b[this.Lvl].maxdur * 100
	opb.Action = make(map[action]bool, growMint - growFood + 1)

	for i:= growFood; i<= growMint; i++ {
		opb.Action[i] = false

		if i.Lvl() >= this.Lvl {
			opb.Action[i] = true
		}
	}

	return opb
}
