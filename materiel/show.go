package materiel

import (
	"github.com/liangran2018/lived/base"
	"github.com/liangran2018/lived/log"

	"github.com/gin-gonic/gin"
)

type ownThing struct {
	Materiel  map[Product]int `json:"materiel"`
	Food      map[Product]int `json:"food"`
	Drug      map[Product]int `json:"drug"`
	Equipment map[Product]int `json:"equipment"`
}

func Show(c *gin.Context) {
	this := GetOwnThings()

	if this.Nothing() {
		base.Output(c, base.NoOwnThing, nil)
		return
	}

	ot := &ownThing{}
	ot.Materiel = make(map[Product]int)
	ot.Food = make(map[Product]int)
	ot.Drug = make(map[Product]int)
	ot.Equipment = make(map[Product]int)
	
	for k, v := range this.product {
		if v != 0 {
			switch k.Type() {
			case Materiel:
				ot.Materiel[k] = v
			case Food:
				ot.Food[k] = v
			case Drug:
				ot.Drug[k] = v
			case Equip:
				ot.Equipment[k] = v
			default:
				log.GetLogger().Log(log.Wrong, "OwnShow", k.Type(), k.Name(), k)
			}
		}
	}

	base.Output(c, 0, ot)
	return
}
