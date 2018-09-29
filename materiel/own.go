package materiel

import (
	"fmt"

	"github.com/liangran2018/lived/log"
	"github.com/liangran2018/lived/base"
)

type OwnThings struct {
	product  map[Product]int
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
	for _, v := range this.product {
		if v != 0 {
			return false
		}
	}

	return true
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