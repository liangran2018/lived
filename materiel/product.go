package materiel

import (
	"github.com/liangran2018/lived/log"
)

type Product int
type ProType int

type Fight struct {
	Attack int
	Defend int
	Dodge int
	Critical int
}

type DrugEffect struct {
	Blood int
	Hurt int
	Ill  int
	Mood int
	Wake int
}

type FoodEnergy struct {
	Blood int
	Ill int
	Hungry int
	Thirst int
	Mood int
	Wake int
	Poison bool
	PoisonPro int
	PoisonDay int
	PoisonLostBlood int
	PoisonLostIll int
	PoisonLostMood int
	PoisonLostWake int
}

const (
	Meat Product = iota
	BBQ
	Potato
	Fish
	Water
	Sewage
	Sea
	Alcohol
	Coco
	Grape
	Jujube

	Herb

	Wood
	Stone
	Pebble	//石子
	Liana	//藤蔓
	Cloth	//布
	Grass 	//草
	Sand	//沙子
	Leaf
	Metal	//金属
	Gold
	Arrow

	Bow
	Knife

	ClothArmor

	SmallFire

	Undefined
)

const (
	Materiel ProType = iota
	Food
	Drug
	Equip

	unknown
)

const (
	Weapon = iota
	Armor
	Fire

	other
)

var productName []string
var productWeight []int
var equipHot map[Product]Fight
var drug     map[Product]DrugEffect
var food     map[Product]FoodEnergy

func (this Product) Name() string {
	return productName[int(this)]
}

func (this Product) Weight() int {
	return productWeight[int(this)]
}

func (this Product) Type() ProType {
	if this >= Meat && this < Herb {
		return Food
	}

	if this >= Herb && this < Wood {
		return Drug
	}

	if this >= Wood && this < Bow {
		return Materiel
	}

	if this >= Bow && this < Undefined {
		return Equip
	}

	log.GetLogger().Log(log.Wrong, "ProductType", this.Name(), this)
	return unknown
}

func (this Product) EquipType() int {
	if this.Type() != Equip {
		return other
	}

	if this >= Bow && this < ClothArmor {
		return Weapon
	}

	if this >= ClothArmor && this < ClothArmor {
		return Armor
	}

	if this >= SmallFire && this < Undefined {
		return Fire
	}

	log.GetLogger().Log(log.Wrong, "EquipType", this.Type(), this.Name(), this)
	return other
}

func (this Product) DrugEff() DrugEffect {
	if this.Type() != Drug {
		return DrugEffect{}
	}

	return drug[this]
}

func (this Product) Food() FoodEnergy {
	if this.Type() != Food {
		return FoodEnergy{}
	}

	return food[this]
}

func (this Product) EquipHot() Fight {
	if this.Type() != Equip {
		return Fight{}
	}

	return equipHot[this]
}

func productInit() {
	productName = make([]string, Undefined)
	productName = []string{"肉", "烤肉", "土豆", "鱼", "净水", "污水", "海水", "酒精", "椰子", "葡萄", "野枣",
		"药草", "木头", "石块", "石子", "藤蔓", "布", "草", "沙子", "阔叶", "金属", "黄金", "箭", "弓", "匕首", "布衣",
		"小火把"}

	productWeight = make([]int, Undefined)
	productWeight = []int{2, 2, 1, 2, 1, 1, 1, 1, 2, 1, 1, 1, 3, 3, 1, 1, 1, 1, 1, 1, 3, 4, 2, 0, 0, 0, 0}

	equipHot = make(map[Product]Fight, Undefined - Bow)
	equipHot[Bow] = Fight{Attack:8, Critical:1}
	equipHot[Knife] = Fight{Attack:12, Critical:2}
	equipHot[ClothArmor] = Fight{Defend:10, Dodge:1}
	equipHot[SmallFire] = Fight{Critical:1, Dodge:1}

	drug = make(map[Product]DrugEffect, Wood - Herb)
	drug[Herb] = DrugEffect{Blood:2, Wake:2}

	food = make(map[Product]FoodEnergy, Herb)

}

