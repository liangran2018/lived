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
	Name string `json:"name"`
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
	Materiel  []one `json:"materiel"`
	Food      []one `json:"food"`
	Drug      []one `json:"drug"`
	Equipment []one `json:"equipment"`
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

func (this *bag) UsedCap() int {
	return this.usedCap
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

func (this *bag) GetMateriel() map[materiel.Product]int {
	return this.product
}

func (this *bag) Count(p materiel.Product) int {
	return this.product[p]
}

func (this *bag) Clear() {
	this.usedCap = 0

	for p, _ := range this.product {
		delete(this.product, p)
	}
}

func BagNotice(c *gin.Context) {
	ot := &ownThing{}
	ot.Materiel = make([]one, 0)
	ot.Food = make([]one, 0)
	ot.Drug = make([]one, 0)
	ot.Equipment = make([]one, 0)
	ot.Bag = exploreBag.max

	for k, v := range materiel.GetOwnThings().OwnProduct() {
		if v != 0 {
			o := one{Name:k.Name(), Num:v, Weight:k.Weight()}
			switch k.Type() {
			case materiel.Materiel:
				ot.Materiel = append(ot.Materiel, o)
			case materiel.Food:
				ot.Food = append(ot.Food, o)
			case materiel.Drug:
				ot.Drug = append(ot.Drug, o)
			case materiel.Equip:
				ot.Equipment = append(ot.Equipment, o)
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
