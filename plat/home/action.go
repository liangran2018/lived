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

	filterWater
	makeWine

	actionEnd
)

func (this action) Lvl() int {
	return actionNature[this].lvl
}

func (this action) Type() int {
	return actionNature[this].t
}
