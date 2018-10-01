package plat

import (
	"fmt"

	"github.com/liangran2018/lived/materiel"
)

type Tc struct {
	H int
	Mi int
}

type animalNature struct {
	meat0 int
	meat1 int
	meat2 int
	blood int
	attack int
	defend int
	dodge int
	critical int
}

type Nature struct {
	timeConsume Tc
	lastTime int
	product map[materiel.Product]*count
	animal  map[materiel.Animal]*count
	lvl     int
}

type count struct {
	num int
	maxNum int
	increase int
}

type Place int

const (
	lsd Place = iota
	gmc

	undefined
)

var placeName  []string
var public map[Place]*Nature

func init() {
	publicInit()
}

func NewPublic() {
	riprapNew()
	bushNew()
}

func LoadPublic(m map[int]int) {
	for k, v := range m {
		public[Place(k)].lastTime = v
	}
}

func LoadProduct(m map[int]map[int]int) {
	for k, v := range m {
		for a, b := range v {
			public[Place(k)].product[materiel.Product(a)].num = b
		}
	}
}

func LoadAnimal(m map[int]map[int]int) {
	for k, v := range m {
		for a, b := range v {
			public[Place(k)].animal[materiel.Animal(a)].num = b
		}
	}
}

func GetPublic() map[Place]*Nature {
	return public
}

func (this *Nature) PublicIncrease(i int) {
	for _, v := range this.product {
		if v.num >= v.maxNum {
			continue
		}

		v.num += i/v.increase
		if v.num > v.maxNum {
			v.num = v.maxNum
		}
	}

	for _, v := range this.animal {
		if v.num >= v.maxNum {
			continue
		}

		v.num += i/v.increase
		if v.num > v.maxNum {
			v.num = v.maxNum
		}
	}
}

func (this *Nature) Lvl() int {
	return this.lvl
}

func (this *Nature) Product() map[materiel.Product]int {
	p := make(map[materiel.Product]int, len(this.product))
	for k, v := range this.product {
		p[k] = v.num
	}
	return p
}

func (this *Nature) Count(p materiel.Product) int {
	return this.product[p].num
}

func (this *Nature) LostProduct(p materiel.Product, i int) {
	this.product[p].num -= i
}

func (this *Nature) Animal() map[materiel.Animal]int {
	p := make(map[materiel.Animal]int, len(this.animal))
	for k, v := range this.animal {
		p[k] = v.num
	}
	return p
}

func (this *Nature) LostAnimal(a materiel.Animal) {
	this.animal[a].num--
}

func (this *Nature) Show() string {
	h, mi := this.NeedTime()
	str := fmt.Sprintf("距离约%d小时%d分钟路程 可采集:", h, mi)
	for p, q := range this.Product() {
		str += fmt.Sprintf("%s%d个 ", p.Name(), q)
	}
	str += "动物:"
	for l, q := range this.Animal() {
		str += fmt.Sprintf("%s%d只 ", l.Name(), q)
	}

	return str
}

func (this Place) Name() string {
	return placeName[int(this)]
}

func publicInit() {
	placeName = make([]string, undefined)
	placeName = []string{"乱石堆", "灌木丛"}

	public = make(map[Place]*Nature, undefined)
	riprapInit()
	bushInit()
}
