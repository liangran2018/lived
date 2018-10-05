package home

import (
	"github.com/liangran2018/lived/plat"
	"github.com/liangran2018/lived/materiel"
)

var actionNature map[action]actionLimit

type action int

type actionLimit struct {
	lvl int
	t int
	time plat.Tc
	m map[materiel.Product]int
	get map[materiel.Product]int
	delay int
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
	make4Arrow
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

	growRise
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
