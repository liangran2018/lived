package materiel

import (
	"fmt"
	"strconv"

	"github.com/liangran2018/lived/log"
	"github.com/liangran2018/lived/base"
)

type OwnThings struct {
	product  map[Product]int
}

var oth *OwnThings

func init() {
	animalInit()
	productInit()

	oth = &OwnThings{}
	oth.product = make(map[Product]int, Undefined)
}

func NewOwnThings() {
	oth.AddProduct(Wood, 2)
	oth.AddProduct(BBQ, 2)
	oth.AddProduct(Herb, 2)
}

func LoadOwnThings(m map[int]int) {
	for k, v := range m {
		oth.AddProduct(Product(k), v)
	}
}

func GetOwnThings() *OwnThings {
	return oth
}

func (this *OwnThings) OwnProduct() map[Product]int {
	return this.product
}

func (this *OwnThings) OwnEquip() map[Product]int {
	m := make(map[Product]int, 0)
	for k, v := range this.product {
		if k.Type() == Equip {
			m[k] = v
		}
	}

	return m
}

func (this *OwnThings) ChooseProduct() []Product {
	s := make([]Product, 0)
	for k, v := range this.product {
		if v != 0 {
			s = append(s, k)
		}
	}

	return s
}

func (this *OwnThings) Count(p Product) int {
	return this.product[p]
}

func (this *OwnThings) AddProduct(k Product, v int) {
	this.product[k] += v
}

func (this *OwnThings) PlusProduct(k Product, v int) {
	this.product[k] -= v
	if this.product[k] < 0 {
		log.GetLogger().Log(log.Wrong, "PlusProduct err", this.product[k], int(k), v)
		this.product[k] = 0
	}
}

func (this *OwnThings) Nothing() bool {
	if len(this.product) == 0 {
		return true
	}
	return false
}

func (this *OwnThings) Show() {
	if this.Nothing() {
		fmt.Println("你一无所有")
		return
	}

	var (
		m = "材料："
		f = "食物："
		d = "药品："
		e = "装备："
		s string
	)

	for k, v := range this.product {
		s = k.Name() + "x" + strconv.Itoa(v) + " "
		switch k.Type() {
		case Materiel:
			m += s
		case Food:
			f += s
		case Drug:
			f += s
		case Equip:
			e += s
		default:
			log.GetLogger().Log(log.Wrong, "OwnShow", k.Type())
		}
	}

	fmt.Println(m + "\n" + f + "\n" + d + "\n" + e + "\n")
	return
}

func (this *OwnThings) Use() {
	for {
		fmt.Println("使用: 1.食物  2.药品  3.取消")
		input, err := base.Input()
		if err != nil {
			log.GetLogger().Log(log.Warning, "ownthingUse input err", err.Error())
			fmt.Println("输入失败")
			continue
		}

		switch input {
		case "1":

		case "2":
		default:
			return
		}
	}
}