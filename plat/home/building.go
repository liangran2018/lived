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
	Action     map[action]int  `json:"action"`
}

const (
	ok = iota
	lvlNotEnough
	busy //当事者
	other //旁观者
	only
)

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

	if b == 3 || b == 6 {
		base.Output(c, 0, code)
		return
	}

	base.Output(c, code, nil)
	return
}

func ActionCheck(c *gin.Context) {
	b, ok := getBuilding(c)
	if !ok {
		base.Output(c, base.ParaInvalid, nil)
		return
	}

	if b != 3 || b != 6 {
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

	if b == 3 {
		if CleanTime.b {
			base.Output(c, base.ParaInvalid, nil)
			return
		}

		if i != CleanTime.a {
			base.Output(c, base.ParaInvalid, nil)
			return
		}

		//过去的时间
		time := env.Int2Time(env.GetTimeInt().Time() - CleanTime.t)
		if time.Year != 0 || time.Month != 0 {
			base.Output(c, base.AlreadyGrowed, nil)
			return
		}

		t := time.Day * 24 + time.Hour
		if t >= actionNature[i].delay {
			base.Output(c, base.AlreadyGrowed, nil)
			return
		} else {
			base.Output(c, base.NotGrowed, struct {
				H int `json:"h"`
				Mi int `json:"mi"`
			}{H:t, Mi:time.Minute})
			return
		}
	}

	if b == 6 {
		if GrowTime.b {
			base.Output(c, base.ParaInvalid, nil)
			return
		}

		if i != GrowTime.a {
			base.Output(c, base.ParaInvalid, nil)
			return
		}

		//过去的时间
		time := env.Int2Time(env.GetTimeInt().Time() - GrowTime.t)
		if time.Year != 0 || time.Month != 0 {
			base.Output(c, base.AlreadyGrowed, nil)
			return
		}

		t := time.Day * 24 + time.Hour
		if t >= actionNature[i].delay {
			base.Output(c, base.AlreadyGrowed, nil)
			return
		} else {
			base.Output(c, base.NotGrowed, struct {
				H int `json:"h"`
				Mi int `json:"mi"`
			}{H:t, Mi:time.Minute})
			return
		}
	}

	base.Output(c, base.ParaInvalid, nil)
	return
}

func Harvest(c *gin.Context) {
	b, ok := getBuilding(c)
	if !ok {
		base.Output(c, base.ParaInvalid, nil)
		return
	}

	if b != 3 || b != 6 {
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

	if b == 3 {
		if CleanTime.b {
			base.Output(c, base.ParaInvalid, nil)
			return
		}

		if i != CleanTime.a {
			base.Output(c, base.ParaInvalid, nil)
			return
		}

		//过去的时间
		time := env.Int2Time(env.GetTimeInt().Time() - CleanTime.t)
		t := time.Day * 24 + time.Hour

		if time.Year != 0 || time.Month != 0 || t >= actionNature[i].delay {
			for k, v := range actionNature[i].get {
				materiel.GetOwnThings().AddProduct(k, v)
			}

			CleanTime.b = false
			base.Output(c, 0, nil)
			return
		} else {
			base.Output(c, base.NotGrowed, struct {
				H int `json:"h"`
				Mi int `json:"mi"`
			}{H:t, Mi:time.Minute})
			return
		}
	}

	if b == 6 {
		if GrowTime.b {
			base.Output(c, base.ParaInvalid, nil)
			return
		}

		if i != GrowTime.a {
			base.Output(c, base.ParaInvalid, nil)
			return
		}

		//过去的时间
		time := env.Int2Time(env.GetTimeInt().Time() - GrowTime.t)
		t := time.Day * 24 + time.Hour

		if time.Year != 0 || time.Month != 0 || t >= actionNature[i].delay {
			for k, v := range actionNature[i].get {
				materiel.GetOwnThings().AddProduct(k, v)
			}

			GrowTime.b = false
			base.Output(c, 0, nil)
			return
		} else {
			base.Output(c, base.NotGrowed, struct {
				H int `json:"h"`
				Mi int `json:"mi"`
			}{H:t, Mi:time.Minute})
			return
		}
	}

	base.Output(c, base.ParaInvalid, nil)
	return
}
