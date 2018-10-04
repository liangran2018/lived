package home

var actionNature map[action]actionLimit

type action int

type actionLimit struct {
	lvl int
	t int
}

const (
	sleep1H action = iota
	sleep4H
	sleep8H

	cookBbq
	cookRoastPotato
	cookCongee
	cookBroth
	cookMashedPotato
	cookStew
	cookDriedFish
	cookSmokedMeet
	cookAnimalBlood

	distill
	filterWater
	makeWine

	makeBandage
	makeDecoction
	makeMintTea
	makeMediWine
	makeFirstAid
	makeGrassPaste
	makePlaster
	makeTonifyPill

	makeStoneAxe
	makeTorch
	makeKnife
	makeShortBow
	make20Arrow
	makeBambooGun
	makeClothArmor
	makeTwoEdgedAxe
	makeHardBambooBow
	makeBigBag
	makeHardBambooCrossBow
	makeSharpTwoAxe
	makeSharpBamboo
	makePoisonAxe
	makePoisonBamboo
	makeRattanArmor

	growFood
	growPotato
	growMint

	goFishing

	actionEnd
)

func (this action) Lvl() int {
	return actionNature[this].lvl
}

func (this action) Type() int {
	return actionNature[this].t
}
