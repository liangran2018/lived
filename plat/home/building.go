package home

import (
	"github.com/liangran2018/lived/base"
	"github.com/liangran2018/lived/env"
	"github.com/liangran2018/lived/materiel"
	"github.com/liangran2018/lived/plat"

	"github.com/gin-gonic/gin"
)

type outputBuild struct {
	IsUpdate   bool            `json:"isUpdate"`
	DurPercent int             `json:"durPercent"`
	Action     map[action]bool `json:"action"`
}

func Show(c *gin.Context) {
	i, ok := getBuilding(c)
	if !ok {
		base.Output(c, base.ParaInvalid, nil)
		return
	}

	opb := doing[i]()
	base.Output(c, 0, opb)
	return
}

func Notice(c *gin.Context) {
	i, ok := getBuilding(c)
	if !ok {
		base.Output(c, base.ParaInvalid, nil)
		return
	}

	this := GetOwnBuilding()
	lvl := this.Own[i].Lvl
	if lvl == homeBuilding[i].maxlvl {
		base.Output(c, base.AlreadyMaxed, nil)
		return
	}

	nature := homeBuilding[i].b[lvl+1]

	om := materiel.GetOwnThings().OwnProduct()

	for k, v := range nature.lvlupNeed {
		if om[k] < v {
			base.Output(c, base.MaterialNotEnough, struct {
				NeedTime     plat.Tc                  `json:"needTime"`
				NeedMateriel map[materiel.Product]int `json:"needMateriel"`
			}{NeedTime: nature.timeConsume, NeedMateriel: nature.lvlupNeed})
			return
		}
	}

	base.Output(c, 0, struct {
		NeedTime     plat.Tc                  `json:"needTime"`
		NeedMateriel map[materiel.Product]int `json:"needMateriel"`
	}{NeedTime: nature.timeConsume, NeedMateriel: nature.lvlupNeed})
	return
}

func Update(c *gin.Context) {
	i, ok := getBuilding(c)
	if !ok {
		base.Output(c, base.ParaInvalid, nil)
		return
	}

	timeNeed := homeBuilding[i].b[obl.Own[i].Lvl+1].timeConsume
	lvlupNeed := homeBuilding[i].b[obl.Own[i].Lvl+1].lvlupNeed

	for k, v := range lvlupNeed {
		materiel.GetOwnThings().PlusProduct(k, v)
	}

	env.GetTime().Add(timeNeed.H, timeNeed.Mi)
	obl.Own[i].Lvl++
	base.Output(c, 0, fillPara())
	return
}

func getBuilding(c *gin.Context) (int, bool) {
	b := c.Query("building")
	if base.Empty(b) {
		return 0, false
	}

	i, err := base.IntVal(b)
	if err != nil || i < 1 || i > len(obl.Own) {
		return 0, false
	}

	return i, true
}
