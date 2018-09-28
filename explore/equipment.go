package explore

import (
	"fmt"
	"strconv"

	"github.com/liangran2018/lived/base"
	"github.com/liangran2018/lived/materiel"
	"github.com/liangran2018/lived/human"
)

type exEquip struct {
	e []materiel.Product
}

var exploreEquip *exEquip
var exploreHot materiel.Fight

func NewEquip() {
	exploreEquip = &exEquip{}
	exploreEquip.e = make([]materiel.Product, 3)
}

func GetEquip() *exEquip {
	return exploreEquip
}

func (this *exEquip) Set(t int, e materiel.Product) {
	this.e[t] = e
}

func (this *exEquip) Clear() {
	this.e = this.e[0:0]
	this.e = make([]materiel.Product, 3)
}

func (this *exEquip) Show() {
	flag := false
	if this.e[0] != materiel.Meat {
		fmt.Printf("武器: %s\n", this.e[0].Name())
		flag = true
	}

	if this.e[1] != materiel.Meat {
		fmt.Printf("护甲: %s\n", this.e[1].Name())
		flag = true
	}

	if this.e[2] != materiel.Meat {
		fmt.Printf("火把: %s\n", this.e[2].Name())
		flag = true
	}

	if !flag {
		fmt.Println("没有装备")
	}

	return
}

func EquipMent() {
	ownEquip := materiel.GetOwnThings().OwnEquip()
	if len(ownEquip) == 0 {
		fmt.Println("你没有装备")
		return
	}

	for {
		fmt.Println("携带装备: 1.武器  2.护甲  3.火把  4.准备好了")
		input, err := base.Input()
		if err != nil {
			fmt.Println("输入失败\n")
			continue
		}

		if input == "4" {
			return
		}

		i, err := strconv.Atoi(input)
		if err != nil || i < 1 || i > 3 {
			fmt.Println("输入错误")
			continue
		}

		ec := chooseEquip(ownEquip, i-1)
		for k, v := range ec {
			fmt.Println(strconv.Itoa(k) + ":" + v.Name() + " ")
		}

		input, err = base.Input()
		if err != nil {
			fmt.Println("输入失败\n")
		}

		j, err := strconv.Atoi(input)
		if err != nil || j < 0 || j >= len(ec) {
			fmt.Println("输入错误")
		}

		fmt.Println("装备成功")
		exploreEquip.Set(i-1, ec[j])
	}
}

func chooseEquip(ot map[materiel.Product]int, t int) []materiel.Product {
	s := make([]materiel.Product, 0)

	for k, v := range ot {
		if k.EquipType() == t && v != 0 {
			s = append(s, k)
		}
	}

	return s
}

func NewHeroHot(arrow bool) {
	exploreHot.Attack = human.GetHuman().Attack() + exploreEquip.e[2].EquipHot().Attack
	exploreHot.Defend = human.GetHuman().Defend() + exploreEquip.e[1].EquipHot().Defend
	exploreHot.Dodge = human.GetHuman().Dodge() + exploreEquip.e[1].EquipHot().Dodge
	exploreHot.Critical = human.GetHuman().Critical() + exploreEquip.e[2].EquipHot().Critical
	if arrow {
		exploreHot.Attack += exploreEquip.e[0].EquipHot().Attack
		exploreHot.Critical += exploreEquip.e[0].EquipHot().Critical
	}
}

func GetHeroHot() materiel.Fight {
	return exploreHot
}

func FightClear() {
	exploreHot = materiel.Fight{0,0,0,0}
}