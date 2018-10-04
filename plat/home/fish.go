package home

func fishing() *outputBuild {
	opb := &outputBuild{IsUpdate:true}

	this := obl.Own[field]
	field := homeBuilding[field]
	if this.Lvl == field.maxlvl {
		opb.IsUpdate = false
	}

	opb.DurPercent = this.Dur/field.b[this.Lvl].maxdur * 100
	opb.Action = make(map[action]bool, 1)
	opb.Action[goFishing] = false

	if this.Lvl == 1 {
		opb.Action[goFishing] = true
	}
	
	return opb
}
