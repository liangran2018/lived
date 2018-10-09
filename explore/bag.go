package explore

import (
	"io/ioutil"
	"encoding/json"

	"github.com/liangran2018/lived/base"
	"github.com/liangran2018/lived/materiel"
	"github.com/liangran2018/lived/log"

	"github.com/gin-gonic/gin"
)

type one struct {
	Num  int    `json:"num"`
	Weight int  `json:"weight"`
}

var exploreBag *bag

type bag struct {
	usedCap  int
	max      int
	product map[materiel.Product]int
}

type ownThing struct {
	Materiel  map[materiel.Product]one `json:"materiel"`
	Food      map[materiel.Product]one `json:"food"`
	Drug      map[materiel.Product]one `json:"drug"`
	Equipment map[materiel.Product]one `json:"equipment"`
	Bag       int   `json:"bag"`
}

func NewBag() {
	exploreBag = &bag{}
	exploreBag.max = 40
	exploreBag.product = make(map[materiel.Product]int)
}

func LoadBag() {
	exploreBag.max = 60
}

func GetBag() *bag {
	return exploreBag
}

//剩余
func (this *bag) RemaindCap() int {
	return this.max - this.usedCap
}

func (this *bag) CanPut(weight int) bool {
	if weight <= this.RemaindCap() {
		return true
	}
	return false
}

func (this *bag) Set(p materiel.Product, i int) bool {
	weight := p.Weight() * i

	if this.CanPut(weight) {
		this.product[p] += i
		this.usedCap += weight
		return true
	}

	return false
}

func (this *bag) Clear() {
	this.usedCap = 0

	for p, _ := range this.product {
		delete(this.product, p)
	}
}

func BagNotice(c *gin.Context) {
	ot := &ownThing{}
	ot.Materiel = make(map[materiel.Product]one)
	ot.Food = make(map[materiel.Product]one)
	ot.Drug = make(map[materiel.Product]one)
	ot.Equipment = make(map[materiel.Product]one)
	ot.Bag = exploreBag.max

	for k, v := range materiel.GetOwnThings().OwnProduct() {
		if v != 0 {
			o := one{Num:v, Weight:k.Weight()}
			switch k.Type() {
			case materiel.Materiel:
				ot.Materiel[k] = o
			case materiel.Food:
				ot.Food[k] = o
			case materiel.Drug:
				ot.Drug[k] = o
			case materiel.Equip:
				ot.Equipment[k] = o
			default:
				log.GetLogger().Log(log.Wrong, "OwnShow", k.Type(), k.Name(), k)
			}
		}
	}

	base.Output(c, 0, ot)
	return
}

func BagChoose(c *gin.Context) {
	c.Request.ParseForm()

	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		base.Output(c, base.PostBodyReadFail, err.Error())
		log.GetLogger().Log(log.Wrong, "bag err", err.Error())
		return
	}
	defer c.Request.Body.Close()

	var data map[int]int
	err = json.Unmarshal(b, &data)
	if err != nil {
		base.Output(c, base.JsonErr, err.Error())
		log.GetLogger().Log(log.Wrong, "bag err", err.Error())
		return
	}

	for k, v := range data {
		if ok := exploreBag.Set(materiel.Product(k), v); !ok {
			base.Output(c, base.BagNotEnough, nil)
			exploreBag.Clear()
			return
		}

		materiel.GetOwnThings().PlusProduct(materiel.Product(k), v)
	}

	base.Output(c, 0, nil)
	return
}

func BagShow(c *gin.Context) {
	s := make(map[materiel.Product]int, len(exploreBag.product))

	for k, v := range exploreBag.product {
		if v != 0 {
			s[k] = v
		}
	}

	base.Output(c, 0, s)
	return
}

func BagAdd(c *gin.Context) {
	product := c.Query("product")
	if base.Empty(product) {
		base.Output(c, base.ParaInvalid, nil)
		return
	}

	p, err := base.IntVal(product)
	if err != nil || p < 0 || p > int(materiel.Gold) {
		base.Output(c, base.ParaInvalid, nil)
		return
	}

	count := c.Query("count")
	if base.Empty(count) {
		base.Output(c, base.ParaInvalid, nil)
		return
	}

	i, err := base.IntVal(count)
	if err != nil || i <= 0 {
		base.Output(c, base.ParaInvalid, nil)
		return
	}

	if ok := exploreBag.Set(materiel.Product(p), i); !ok {
		base.Output(c, base.BagNotEnough, nil)
		return
	}

	base.Output(c, 0, nil)
	return
}

func BagPlus(c *gin.Context) {
	product := c.Query("product")
	if base.Empty(product) {
		base.Output(c, base.ParaInvalid, nil)
		return
	}

	p, err := base.IntVal(product)
	if err != nil || p < 0 || p > int(materiel.Gold) {
		base.Output(c, base.ParaInvalid, nil)
		return
	}

	count := c.Query("count")
	if base.Empty(count) {
		base.Output(c, base.ParaInvalid, nil)
		return
	}

	i, err := base.IntVal(count)
	if err != nil || i <= 0 || i > exploreBag.product[materiel.Product(p)] {
		base.Output(c, base.ParaInvalid, nil)
		return
	}

	exploreBag.Set(materiel.Product(p), -i)

	base.Output(c, 0, nil)
	return
}
