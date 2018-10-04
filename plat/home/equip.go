package home

func equip() *outputBuild {
	opb := &outputBuild{IsUpdate:true}

	this := obl.Own[tool]
	tool := homeBuilding[tool]
	if this.Lvl == tool.maxlvl {
		opb.IsUpdate = false
	}

	opb.DurPercent = this.Dur/tool.b[this.Lvl].maxdur * 100
	opb.Action = make(map[action]bool, makeRattanArmor - cookBbq + 1)

	for i:= makeStoneAxe; i<= makeRattanArmor; i++ {
		opb.Action[i] = false

		if i.Lvl() >= this.Lvl {
			opb.Action[i] = true
		}
	}

	return opb
}
