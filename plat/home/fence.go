package home

func rail() *outputBuild {
	opb := &outputBuild{IsUpdate:true}

	this := obl.Own[fence]
	fence := homeBuilding[fence]
	if this.Lvl == fence.maxlvl {
		opb.IsUpdate = false
	}

	opb.DurPercent = this.Dur/fence.b[this.Lvl].maxdur * 100
	opb.Action = make(map[action]bool, 0)

	return opb
}
