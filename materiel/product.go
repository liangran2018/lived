package materiel

import (
	"github.com/liangran2018/lived/log"
)

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
	Thirst int
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
}

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

	if this >= Wood && this < BigBag {
		return Materiel
	}

	if this >= ShortBow && this < Undefined {
		return Equip
	}

	log.GetLogger().Log(log.Wrong, "ProductType", this)
	return unknown
}

func (this Product) EquipType() int {
	if this.Type() != Equip {
		return Unknown
	}

	if this >= ShortBow && this < Knife {
		return Remote
	}

	if this >= Knife && this < ClothArmor {
		return Melee
	}

	if this >= ClothArmor && this < ClothArmor {
		return Armor
	}

	if this == Fire {
		return Fire
	}

	log.GetLogger().Log(log.Wrong, "EquipType", this.Type(), this)
	return Unknown
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
	productName = []string{"肉", "烤肉", "土豆", "烤土豆", "粥", "肉汤", "土豆泥", "秘制炖肉", "鱼", "鱼干",
		"烟熏肉", "兽血乱炖", "净水", "污水", "海水", "酒精", "椰子", "葡萄", "野枣", "薄荷",
		"蜂蜜", "人参", "粗粮", "盐", "兽血",
		"药草", "绷带", "汤药", "薄荷茶", "药酒", "急救药", "青草膏", "膏药", "大补丸",
		"木头", "石块", "石子", "藤蔓", "草", "沙子", "阔叶", "树脂", "竹子", "兽筋", "兽皮", "毒液",
		"金属", "黄金", "箭", "大背包",
		"短弓", "硬竹弓", "硬竹弩", "匕首", "石斧", "竹枪", "双刃斧", "锋利的双刃斧", "锋利的竹枪",
		"毒刃斧", "毒竹枪", "布衣", "藤甲", "火把"}

	productWeight = make([]int, BigBag)
	productWeight = []int{2, 2, 1, 1, 1, 3, 2, 3, 2, 1,
		2, 4, 1, 1, 1, 1, 2, 1, 1, 1,
		1, 1, 1, 1, 1,
		1, 1, 2, 2, 2, 1, 2, 1, 1,
		3, 3, 1, 1, 1, 1, 1, 1, 2, 1, 3, 1,
		3, 4, 2}

	equipHot = make(map[Product]Fight, Undefined - ShortBow)
	equipHot[ShortBow] = Fight{Attack:10}
	equipHot[HardBambooBow] = Fight{Attack:15, Critical:1}
	equipHot[HardBambooCrossBow] = Fight{Attack:20, Critical:2}

	equipHot[Knife] = Fight{Attack:12}
	equipHot[StoneAxe] = Fight{Attack:18, Critical:1}
	equipHot[BambooGun] = Fight{Attack:18}
	equipHot[TwoEdgedAxe] = Fight{Attack:24, Critical:1}
	equipHot[SharpTwoAxe] = Fight{Attack:30, Critical:2}
	equipHot[SharpBamboo] = Fight{Attack:24, Critical:1}
	equipHot[PoisonAxe] = Fight{Attack:36, Critical:3}
	equipHot[PoisonBamboo] = Fight{Attack:30, Critical:2}

	equipHot[ClothArmor] = Fight{Defend:18, Dodge:1}
	equipHot[RattanArmor] = Fight{Defend:26, Dodge:2}

	drug = make(map[Product]DrugEffect, Wood - Herb)
	drug[Herb] = DrugEffect{Blood:2, Wake:2}
	drug[Bandage] = DrugEffect{Blood:3, Hurt:15}
	drug[Decoction] = DrugEffect{Blood:3, Ill:15, Thirst:5}
	drug[MintTea] = DrugEffect{Wake:20, Thirst:10}
	drug[MediWine] = DrugEffect{Blood:5, Hurt:30, Thirst:5}
	drug[FirstAid] = DrugEffect{Blood:5, Ill:25, Wake:15}
	drug[GrassPaste] = DrugEffect{Wake:40}
	drug[Plaster] = DrugEffect{Hurt:45}
	drug[TonifyPill] = DrugEffect{Blood:10, Ill:40, Wake:25}

	food = make(map[Product]FoodEnergy, Herb)
	food[Meat] = FoodEnergy{Ill:-5, Hungry:20}
	food[BBQ] = FoodEnergy{Hungry:25, Thirst:-3}
	food[Potato] = FoodEnergy{Ill:-3, Hungry:15}
	food[RoastPotato] = FoodEnergy{Hungry:20, Thirst:-3}
	food[Congee] = FoodEnergy{Hungry:30, Thirst:5}
	food[Broth] = FoodEnergy{Hungry:40, Thirst:10}
	food[MashedPotato] = FoodEnergy{Hungry:35}
	food[Stew] = FoodEnergy{Hungry:50, Wake:30, Thirst:5}
	food[Fish] = FoodEnergy{Ill:5, Hungry:20}
	food[DriedFish] = FoodEnergy{Ill:15, Hungry:30, Thirst:-10}

	food[SmokedMeet] = FoodEnergy{Hungry:40, Thirst:-10}
	food[AnimalBlood] = FoodEnergy{Ill:25, Hungry:50, Wake:20, Thirst:10}
	food[Water] = FoodEnergy{Thirst:20}
	food[Sewage] = FoodEnergy{Ill:-10, Thirst:20}
	food[Sea] = FoodEnergy{Ill:-8, Thirst:10}
	food[Alcohol] = FoodEnergy{Ill:-10, Thirst:10}
	food[Coco] = FoodEnergy{Hungry:10, Thirst:5}
	food[Grape] = FoodEnergy{Hungry:4, Thirst:3}
	food[Jujube] = FoodEnergy{Hungry:5, Thirst:2}
	food[Mint] = FoodEnergy{Wake:5}
	food[Honey] = FoodEnergy{Hungry:5, Wake:10}
	food[Ginseng] = FoodEnergy{Ill:8, Hungry:5, Wake:8}
	food[Rise] = FoodEnergy{Ill:-1, Hungry:5, Thirst:-2}
	food[Salt] = FoodEnergy{Hungry:10, Thirst:-10}
	food[Blood] = FoodEnergy{Ill:2, Hungry:3, Thirst:1}
}

