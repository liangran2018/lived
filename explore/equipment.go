package explore

import (
	"io/ioutil"
	"encoding/json"

	"github.com/liangran2018/lived/base"
	"github.com/liangran2018/lived/materiel"
	"github.com/liangran2018/lived/log"

	"github.com/gin-gonic/gin"
)

type exEquip struct {
	e []materiel.Product
}

var exploreEquip *exEquip

func NewEquip() {
	exploreEquip = &exEquip{}
	exploreEquip.e = make([]materiel.Product, 4)
}

func GetEquip() *exEquip {
	return exploreEquip
}

func (this *exEquip) Set(t int, e materiel.Product) {
	this.e[t] = e
}

func (this *exEquip) Clear() {
	this.e = this.e[0:0]
	this.e = make([]materiel.Product, 4)
}

func EquipNotice(c *gin.Context) {
	t := c.Query("type")
	if base.Empty(t) {
		log.GetLogger().Log(log.Wrong, "EquipChoose", t)
		base.Output(c, base.ParaInvalid, nil)
		return
	}

	i, err := base.IntVal(t)
	if err != nil || i < 0 || i >= materiel.Unknown {
		log.GetLogger().Log(log.Wrong, "EquipChoose", i)
		base.Output(c, base.ParaInvalid, nil)
		return
	}

	s := make([]materiel.Product, 0)

	oe := materiel.GetOwnThings().OwnEquip()
	for k, v := range oe {
		if k.EquipType() == i && v != 0 {
			s = append(s, k)
		}
	}

	base.Output(c, 0, s)
	return
}

func Equip(c *gin.Context) {
	c.Request.ParseForm()

	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		base.Output(c, base.ParaInvalid, err.Error())
		log.GetLogger().Log(log.Wrong, "Equip err", err.Error())
		return
	}
	defer c.Request.Body.Close()

	var data map[int]int
	err = json.Unmarshal(b, &data)
	if err != nil {
		base.Output(c, base.ParaInvalid, err.Error())
		log.GetLogger().Log(log.Wrong, "Equip err", err.Error())
		return
	}

	if len(data) > 4 {
		base.Output(c, base.ParaInvalid, nil)
		log.GetLogger().Log(log.Wrong, "Equip err", data)
		return
	}

	for k, v := range data {
		if k != materiel.Product(v).EquipType() {
			continue
		}

		exploreEquip.Set(k, materiel.Product(v))
	}

	base.Output(c, 0, nil)
	return
}

func Show(c *gin.Context) {
	base.Output(c, 0, exploreEquip.e)
	return
}
