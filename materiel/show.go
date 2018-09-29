package materiel

import (
	"github.com/liangran2018/lived/base"
	"github.com/liangran2018/lived/log"

	"github.com/gin-gonic/gin"
)

type ownThing struct {
	Materiel  []one `json:"materiel"`
	Food      []one `json:"food"`
	Drug      []one `json:"drug"`
	Equipment []one `json:"equipment"`
}

type one struct {
	Name string `json:"name"`
	Num  int    `json:"num"`
}

func Show(c *gin.Context) {
	this := GetOwnThings()

	if this.Nothing() {
		base.Output(c, base.NoOwnThing, nil)
		return
	}

	ot := &ownThing{}
	ot.Materiel = make([]one, 0)
	ot.Food = make([]one, 0)
	ot.Drug = make([]one, 0)
	ot.Equipment = make([]one, 0)
	
	for k, v := range this.product {
		if v != 0 {
			o := one{Name:k.Name(), Num:v}
			switch k.Type() {
			case Materiel:
				ot.Materiel = append(ot.Materiel, o)
			case Food:
				ot.Food = append(ot.Food, o)
			case Drug:
				ot.Drug = append(ot.Drug, o)
			case Equip:
				ot.Equipment = append(ot.Equipment, o)
			default:
				log.GetLogger().Log(log.Wrong, "OwnShow", k.Type(), k.Name(), k)
			}
		}
	}

	base.Output(c, 0, ot)
	return
}
