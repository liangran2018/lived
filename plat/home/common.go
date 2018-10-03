package home

import (
	"github.com/liangran2018/lived/base"
)

func NewOwnBuilding() {
	obl.Own[bed] = ownBuild{Lvl:1, Dur:100}
}

func LoadOwnBuilding(b []base.OB) {
	for k, v := range b {
		obl.Own[k] = ownBuild{Lvl:v.Lvl, Dur:v.Dur}
	}
}

func GetOwnBuilding() *OwnBuilding {
	return obl
}

func GetOwn() []bool {
	b := make([]bool, 9)
	for k, v := range obl.Own {
		if v.Lvl != 0 {
			b[k] = true
		}
	}

	return b
}
