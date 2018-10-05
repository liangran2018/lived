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

	lvl := obl.Own[i].Lvl
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

	lvl := obl.Own[i].Lvl
	if lvl == homeBuilding[i].maxlvl {
		base.Output(c, base.AlreadyMaxed, nil)
		return
	}

	timeNeed := homeBuilding[i].b[lvl+1].timeConsume
	lvlupNeed := homeBuilding[i].b[lvl+1].lvlupNeed

	for k, v := range lvlupNeed {
		materiel.GetOwnThings().PlusProduct(k, v)
	}

	env.GetTime().Add(timeNeed.H, timeNeed.Mi)
	obl.Own[i].Lvl++
	base.Output(c, 0, fillPara())
	return
}

func ActionNotice(c *gin.Context) {
	b, ok := getBuilding(c)
	if !ok {
		base.Output(c, base.ParaInvalid, nil)
		return
	}

	if b == 1 {
		base.Output(c, 0, nil)
		return
	}

	a, ok := getAction(c)
	if !ok {
		base.Output(c, base.ParaInvalid, nil)
		return
	}

	n := actionNotice(b, a)
	base.Output(c, n.code, n)
	return
}

func ActionChoose(c *gin.Context) {
	b, ok := getBuilding(c)
	if !ok {
		base.Output(c, base.ParaInvalid, nil)
		return
	}

	i, ok := getAction(c)
	if !ok {
		base.Output(c, base.ParaInvalid, nil)
		return
	}

	if i.Type() != b {
		base.Output(c, base.ParaInvalid, nil)
		return
	}

	if BigBag && i == makeBigBag {
		base.Output(c, base.ParaInvalid, nil)
		return
	}

	code := actionDoing[b](i)

	if code != 0 {
		base.Output(c, code, nil)
		return
	}

	base.Output(c, 0, fillPara())
	return
}
