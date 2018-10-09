package home

import (
	"github.com/liangran2018/lived/base"
	"github.com/liangran2018/lived/log"
	"github.com/liangran2018/lived/plat"
	"github.com/liangran2018/lived/materiel"

	"github.com/gin-gonic/gin"
)

type notice struct {
	code     int
	NeedTime plat.Tc `json:"needTime"`
	NeedMateriel map[materiel.Product]int `json:"needMateriel"`
}

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

func GetOwnBuild() []int {
	b := make([]int, 9)
	for k, v := range obl.Own {
		b[k] = v.Lvl
	}

	return b
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

func getAction(c *gin.Context) (action, bool) {
	b := c.Query("action")
	if base.Empty(b) {
		return action(0), false
	}

	i, err := base.IntVal(b)
	if err != nil || i < 0 || i >= int(actionEnd) {
		return action(0), false
	}

	return action(i), true
}

func actionNotice(i int, a action) *notice {
	n := &notice{code:0}
	this := actionNature[a]

	if this.t != i || this.lvl > obl.Own[i].Lvl || (BigBag && a == makeBigBag) {
		n.code = base.ParaInvalid
		log.GetLogger().Log(log.Wrong, "actionNotice para", i, a)
		return n
	}

	n.NeedTime = this.time
	n.NeedMateriel = this.m

	om := materiel.GetOwnThings().OwnProduct()
	for k, v := range n.NeedMateriel {
		if om[k] < v {
			n.code = base.MaterialNotEnough
			return n
		}
	}

	return n
}
