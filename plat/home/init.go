package home

import (
	"github.com/liangran2018/lived/materiel"
	"github.com/liangran2018/lived/plat"
)

type buildingNature struct {
	maxdur int
	lvlupNeed map[materiel.Product]int
	timeConsume plat.Tc
}

type ownBuild struct {
	Lvl int	`json:"lvl"`
	Dur int `json:"dur"`
}

type OwnBuilding struct {
	Own []ownBuild `json:"own"`
}

type building struct {
	maxlvl int
	b []buildingNature
}

var obl *OwnBuilding

//不可变， 设施属性， 第一个参数是种类（room/cook等）， 第二个是等级，从1开始
var homeBuilding []building

//每个建筑的详情
var doing []func() *outputBuild

var actionDoing []func(i action) int

var convert []string

// 大背包
var BigBag bool

const (
	_ = iota
	bed
	fire
	water
	drug
	tool
	field
	fence
	fish
)

func init() {
	HomeBuildingInit()
	convert = []string{"", "床", "火堆", "净水器", "药盒", "工具台", "田地", "栅栏", "钓鱼台"}
	obl = &OwnBuilding{}
	obl.Own = make([]ownBuild, 9)

	doing = make([]func() *outputBuild, 9)
	doing = []func() *outputBuild {nil, sleep, cook, clean, medicine, equip, grow, rail, fishing}
	actionDoing = make([]func(i action) int, 9)
	actionDoing = []func(i action) int {nil, sleepAction, commonAction, cleanAction, commonAction, equipAction,
		growAction, railAction, fishAction}

	actionNature = make(map[action]actionLimit, actionEnd)
	actionNature[sleep1H] = actionLimit{lvl:1, t:bed}
	actionNature[sleep4H] = actionLimit{lvl:1, t:bed}
	actionNature[sleep8H] = actionLimit{lvl:1, t:bed}

	actionNature[cookBbq] = actionLimit{lvl:1, t:fire, time:plat.Tc{Mi:45},
		m:map[materiel.Product]int {materiel.Wood:1, materiel.Meat:1},
		get:map[materiel.Product]int {materiel.BBQ:1}}
	actionNature[cookRoastPotato] = actionLimit{lvl:1, t:fire, time:plat.Tc{Mi:30},
		m:map[materiel.Product]int {materiel.Wood:1, materiel.Potato:1},
		get:map[materiel.Product]int {materiel.RoastPotato:1}}
	actionNature[cookCongee] = actionLimit{lvl:1, t:fire, time:plat.Tc{Mi:20},
		m:map[materiel.Product]int {materiel.Water:1, materiel.Wood:1, materiel.Rise:1},
		get:map[materiel.Product]int {materiel.Congee:1}}
	actionNature[cookBroth] = actionLimit{lvl:2, t:fire, time:plat.Tc{H:1},
		m:map[materiel.Product]int {materiel.Water:1, materiel.Wood:1, materiel.Meat:2},
		get:map[materiel.Product]int {materiel.Broth:1}}
	actionNature[cookMashedPotato] = actionLimit{lvl:2, t:fire, time:plat.Tc{Mi:30},
		m:map[materiel.Product]int {materiel.Water:1, materiel.Wood:2, materiel.Potato:2},
		get:map[materiel.Product]int {materiel.MashedPotato:2}}
	actionNature[cookStew] = actionLimit{lvl:2, t:fire, time:plat.Tc{Mi:30},
		m:map[materiel.Product]int {materiel.Water:2, materiel.Wood:2, materiel.Meat:1,
		materiel.Mint:1, materiel.Honey:2, materiel.Potato:2},
		get:map[materiel.Product]int {materiel.Stew:2}}
	actionNature[cookDriedFish] = actionLimit{lvl:3, t:fire, time:plat.Tc{H:1},
		m:map[materiel.Product]int {materiel.Salt:2, materiel.Fish:1},
		get:map[materiel.Product]int {materiel.DriedFish:1}}
	actionNature[cookSmokedMeet] = actionLimit{lvl:3, t:fire, time:plat.Tc{H:1},
		m:map[materiel.Product]int {materiel.Salt:3, materiel.Meat:2},
		get:map[materiel.Product]int {materiel.SmokedMeet:2}}
	actionNature[cookAnimalBlood] = actionLimit{lvl:3, t:fire, time:plat.Tc{Mi:30},
		m:map[materiel.Product]int {materiel.Water:1, materiel.Blood:1, materiel.Meat:2, materiel.Potato:2},
		get:map[materiel.Product]int {materiel.AnimalBlood:1}}

	actionNature[distill] = actionLimit{lvl:1, t:water, time:plat.Tc{H:1},
		m:map[materiel.Product]int {materiel.Sea:3, materiel.Wood:1, materiel.Leaf:1},
		get:map[materiel.Product]int {materiel.Water:2, materiel.Salt:1}, delay:48}
	actionNature[filterWater] = actionLimit{lvl:1, t:water, time:plat.Tc{H:1},
		m:map[materiel.Product]int {materiel.Sewage:2, materiel.Liana:1, materiel.Leaf:1},
		get:map[materiel.Product]int {materiel.Water:2}, delay:24}
	actionNature[makeWine] = actionLimit{lvl:1, t:water, time:plat.Tc{H:1},
		m:map[materiel.Product]int {materiel.Water:2, materiel.Wood:1, materiel.Grape:2},
		get:map[materiel.Product]int {materiel.Alcohol:2}, delay:48}

	actionNature[makeBandage] = actionLimit{lvl:1, t:drug, time:plat.Tc{H:1},
		m:map[materiel.Product]int {materiel.Liana:1, materiel.Herb:1},
		get:map[materiel.Product]int {materiel.Bandage:1}}
	actionNature[makeDecoction] = actionLimit{lvl:1, t:drug, time:plat.Tc{H:1},
		m:map[materiel.Product]int {materiel.Water:1, materiel.Herb:1},
		get:map[materiel.Product]int {materiel.Decoction:1}}
	actionNature[makeMintTea] = actionLimit{lvl:1, t:drug, time:plat.Tc{H:1},
		m:map[materiel.Product]int {materiel.Water:1, materiel.Mint:1},
		get:map[materiel.Product]int {materiel.MintTea:1}}
	actionNature[makeMediWine] = actionLimit{lvl:2, t:drug, time:plat.Tc{H:1},
		m:map[materiel.Product]int {materiel.Herb:2, materiel.Alcohol:1},
		get:map[materiel.Product]int {materiel.MediWine:1}}
	actionNature[makeFirstAid] = actionLimit{lvl:2, t:drug, time:plat.Tc{H:1},
		m:map[materiel.Product]int {materiel.Herb:2, materiel.Alcohol:1, materiel.Mint:1, materiel.Ginseng:1},
		get:map[materiel.Product]int {materiel.FirstAid:1}}
	actionNature[makeGrassPaste] = actionLimit{lvl:2, t:drug, time:plat.Tc{H:1},
		m:map[materiel.Product]int {materiel.Mint:2, materiel.Water:1, materiel.Resin:1},
		get:map[materiel.Product]int {materiel.GrassPaste:1}}
	actionNature[makePlaster] = actionLimit{lvl:3, t:drug, time:plat.Tc{H:1},
		m:map[materiel.Product]int {materiel.Leaf:1, materiel.Liana:2, materiel.Herb:1, materiel.Resin:1},
		get:map[materiel.Product]int {materiel.Plaster:1}}
	actionNature[makeTonifyPill] = actionLimit{lvl:3, t:drug, time:plat.Tc{H:1},
		m:map[materiel.Product]int {materiel.Water:1, materiel.Honey:1, materiel.Herb:2, materiel.Resin:1},
		get:map[materiel.Product]int {materiel.Plaster:1}}

	actionNature[makeStoneAxe] = actionLimit{lvl:1, t:tool, time:plat.Tc{H:1},
		m:map[materiel.Product]int {materiel.Wood:1, materiel.Stone:1},
		get:map[materiel.Product]int {materiel.StoneAxe:1}}
	actionNature[makeTorch] = actionLimit{lvl:1, t:tool, time:plat.Tc{H:1},
		m:map[materiel.Product]int {materiel.Wood:1, materiel.Liana:1, materiel.Resin:1},
		get:map[materiel.Product]int {materiel.Torch:1}}
	actionNature[makeKnife] = actionLimit{lvl:1, t:tool, time:plat.Tc{H:1},
		m:map[materiel.Product]int {materiel.Stone:1, materiel.Pebble:1},
		get:map[materiel.Product]int {materiel.Knife:1}}
	actionNature[makeShortBow] = actionLimit{lvl:1, t:tool, time:plat.Tc{H:1},
		m:map[materiel.Product]int {materiel.Liana:1, materiel.Bamboo:2, materiel.Tendons:1},
		get:map[materiel.Product]int {materiel.ShortBow:1}}
	actionNature[make4Arrow] = actionLimit{lvl:1, t:tool, time:plat.Tc{Mi:10},
		m:map[materiel.Product]int {materiel.Pebble:2, materiel.Bamboo:1},
		get:map[materiel.Product]int {materiel.Arrow:4}}
	actionNature[makeBambooGun] = actionLimit{lvl:2, t:tool, time:plat.Tc{H:1},
		m:map[materiel.Product]int {materiel.Pebble:1, materiel.Bamboo:2},
		get:map[materiel.Product]int {materiel.BambooGun:1}}
	actionNature[makeClothArmor] = actionLimit{lvl:2, t:tool, time:plat.Tc{H:1},
		m:map[materiel.Product]int {materiel.Liana:2, materiel.Hide:1},
		get:map[materiel.Product]int {materiel.ClothArmor:1}}
	actionNature[makeTwoEdgedAxe] = actionLimit{lvl:2, t:tool, time:plat.Tc{H:1},
		m:map[materiel.Product]int {materiel.Stone:1, materiel.Pebble:2, materiel.Wood:2},
		get:map[materiel.Product]int {materiel.TwoEdgedAxe:1}}
	actionNature[makeHardBambooBow] = actionLimit{lvl:2, t:tool, time:plat.Tc{H:1},
		m:map[materiel.Product]int {materiel.Liana:4, materiel.Bamboo:8, materiel.Tendons:2},
		get:map[materiel.Product]int {materiel.HardBambooBow:1}}
	actionNature[makeBigBag] = actionLimit{lvl:2, t:tool, time:plat.Tc{H:1},
		m:map[materiel.Product]int {materiel.Liana:8, materiel.Hide:8},
		get:map[materiel.Product]int {materiel.BigBag:1}}
	actionNature[makeHardBambooCrossBow] = actionLimit{lvl:3, t:tool, time:plat.Tc{H:1},
		m:map[materiel.Product]int {materiel.Liana:2, materiel.Bamboo:4, materiel.Tendons:2, materiel.Wood:2},
		get:map[materiel.Product]int {materiel.HardBambooCrossBow:1}}
	actionNature[makeSharpTwoAxe] = actionLimit{lvl:3, t:tool, time:plat.Tc{H:1},
		m:map[materiel.Product]int {materiel.Stone:1, materiel.Pebble:2, materiel.Metal:1, materiel.Wood:2, materiel.Gold:1},
		get:map[materiel.Product]int {materiel.SharpTwoAxe:1}}
	actionNature[makeSharpBamboo] = actionLimit{lvl:3, t:tool, time:plat.Tc{H:1},
		m:map[materiel.Product]int {materiel.Stone:1, materiel.Pebble:2, materiel.Metal:1, materiel.Wood:2, materiel.Gold:1},
		get:map[materiel.Product]int {materiel.SharpBamboo:1}}
	actionNature[makePoisonAxe] = actionLimit{lvl:3, t:tool, time:plat.Tc{H:1},
		m:map[materiel.Product]int {materiel.Stone:1, materiel.Pebble:2, materiel.Metal:2, materiel.Wood:2, materiel.Venom:3},
		get:map[materiel.Product]int {materiel.PoisonAxe:1}}
	actionNature[makePoisonBamboo] = actionLimit{lvl:3, t:tool, time:plat.Tc{H:1},
		m:map[materiel.Product]int {materiel.Stone:1, materiel.Pebble:2, materiel.Metal:2, materiel.Wood:2, materiel.Venom:3},
		get:map[materiel.Product]int {materiel.PoisonAxe:1}}
	actionNature[makeRattanArmor] = actionLimit{lvl:3, t:tool, time:plat.Tc{H:1},
		m:map[materiel.Product]int {materiel.Liana:2, materiel.Hide:1, materiel.Bamboo:6, materiel.Gold:2, materiel.Metal:6},
		get:map[materiel.Product]int {materiel.PoisonAxe:1}}

	actionNature[growRise] = actionLimit{lvl:1, t:field, time:plat.Tc{H:1},
		m:map[materiel.Product]int {materiel.Water:3, materiel.Rise:3},
		get:map[materiel.Product]int {materiel.Rise:6}, delay:48}
	actionNature[growPotato] = actionLimit{lvl:1, t:field, time:plat.Tc{H:1},
		m:map[materiel.Product]int {materiel.Water:1, materiel.Potato:2},
		get:map[materiel.Product]int {materiel.Potato:4}, delay:36}
	actionNature[growMint] = actionLimit{lvl:2, t:field, time:plat.Tc{H:1},
		m:map[materiel.Product]int {materiel.Water:4, materiel.Mint:3},
		get:map[materiel.Product]int {materiel.Mint:6}, delay:36}

	actionNature[goFishing] = actionLimit{lvl:1, t:fish}
}

func HomeBuildingInit() {
	homeBuilding = make([]building, 9)
	bedInit()
	fireInit()
	waterInit()
	drugInit()
	toolInit()
	fieldInit()
	fenceInit()
	fishInit()
}

func bedInit() {
	homeBuilding[bed].maxlvl = 3
	homeBuilding[bed].b = make([]buildingNature, 4)
	homeBuilding[bed].b[1] = buildingNature{maxdur:100, timeConsume:plat.Tc{Mi:30},
		lvlupNeed:map[materiel.Product]int{materiel.Wood:4, materiel.Pebble:5, materiel.Liana:2}}
	homeBuilding[bed].b[2] = buildingNature{maxdur:200, timeConsume:plat.Tc{H:1},
		lvlupNeed:map[materiel.Product]int{materiel.Wood:8, materiel.Stone:5, materiel.Pebble:10, materiel.Liana:5}}
	homeBuilding[bed].b[3] = buildingNature{maxdur:300, timeConsume:plat.Tc{H:2},
		lvlupNeed:map[materiel.Product]int{materiel.Wood:15, materiel.Stone:8, materiel.Pebble:15, materiel.Liana:10, materiel.Metal:5}}
}

func fireInit() {
	homeBuilding[fire].maxlvl = 3
	homeBuilding[fire].b = make([]buildingNature, 4)
	homeBuilding[fire].b[1] = buildingNature{maxdur:100, timeConsume:plat.Tc{Mi:40},
		lvlupNeed:map[materiel.Product]int{materiel.Wood:3, materiel.Grass:3, materiel.Pebble:6}}
	homeBuilding[fire].b[2] = buildingNature{maxdur:200, timeConsume:plat.Tc{H:1, Mi:20},
		lvlupNeed:map[materiel.Product]int{materiel.Wood:10, materiel.Grass:6, materiel.Leaf:7, materiel.Stone:4}}
	homeBuilding[fire].b[3] = buildingNature{maxdur:300, timeConsume:plat.Tc{H:2},
		lvlupNeed:map[materiel.Product]int{materiel.Wood:15, materiel.Grass:10, materiel.Leaf:10, materiel.Stone:8}}
}

func waterInit() {
	homeBuilding[water].maxlvl = 1
	homeBuilding[water].b = make([]buildingNature, 2)
	homeBuilding[water].b[1] = buildingNature{maxdur:100, timeConsume:plat.Tc{H:1},
		lvlupNeed:map[materiel.Product]int{materiel.Wood:5, materiel.Stone:5, materiel.Grass:5, materiel.Liana:3, materiel.Leaf:3}}
}

func drugInit() {
	homeBuilding[drug].maxlvl = 3
	homeBuilding[drug].b = make([]buildingNature, 4)
	homeBuilding[drug].b[1] = buildingNature{maxdur:100, timeConsume:plat.Tc{Mi:30},
		lvlupNeed:map[materiel.Product]int{materiel.Wood:3, materiel.Pebble:2, materiel.Liana:2}}
	homeBuilding[drug].b[2] = buildingNature{maxdur:200, timeConsume:plat.Tc{H:1, Mi:15},
		lvlupNeed:map[materiel.Product]int{materiel.Wood:5, materiel.Stone:3, materiel.Liana:4, materiel.Leaf:5}}
	homeBuilding[drug].b[3] = buildingNature{maxdur:300, timeConsume:plat.Tc{H:2},
		lvlupNeed:map[materiel.Product]int{materiel.Wood:8, materiel.Stone:5, materiel.Liana:7, materiel.Leaf:8}}
}

func toolInit() {
	homeBuilding[tool].maxlvl = 3
	homeBuilding[tool].b = make([]buildingNature, 4)
	homeBuilding[tool].b[1] = buildingNature{maxdur:100, timeConsume:plat.Tc{H:1},
		lvlupNeed:map[materiel.Product]int{materiel.Wood:3, materiel.Stone:2, materiel.Liana:2}}
	homeBuilding[tool].b[2] = buildingNature{maxdur:200, timeConsume:plat.Tc{H:2},
		lvlupNeed:map[materiel.Product]int{materiel.Wood:5, materiel.Stone:3, materiel.Liana:4, materiel.Leaf:5}}
	homeBuilding[tool].b[3] = buildingNature{maxdur:300, timeConsume:plat.Tc{H:3},
		lvlupNeed:map[materiel.Product]int{materiel.Wood:8, materiel.Stone:5, materiel.Liana:7, materiel.Leaf:8, materiel.Metal:5}}
}

func fieldInit() {
	homeBuilding[field].maxlvl = 2
	homeBuilding[field].b = make([]buildingNature, 3)
	homeBuilding[field].b[1] = buildingNature{maxdur:100, timeConsume:plat.Tc{H:1},
		lvlupNeed:map[materiel.Product]int{materiel.Wood:3, materiel.Stone:2, materiel.Liana:2}}
	homeBuilding[field].b[2] = buildingNature{maxdur:200, timeConsume:plat.Tc{H:2},
		lvlupNeed:map[materiel.Product]int{materiel.Wood:5, materiel.Stone:3, materiel.Liana:4, materiel.Leaf:5}}
}

func fenceInit() {
	homeBuilding[fence].maxlvl = 3
	homeBuilding[fence].b = make([]buildingNature, 4)
	homeBuilding[fence].b[1] = buildingNature{maxdur:150, timeConsume:plat.Tc{H:1},
		lvlupNeed:map[materiel.Product]int{materiel.Wood:6, materiel.Pebble:4, materiel.Liana:4}}
	homeBuilding[fence].b[2] = buildingNature{maxdur:300, timeConsume:plat.Tc{H:1, Mi:30},
		lvlupNeed:map[materiel.Product]int{materiel.Wood:10, materiel.Stone:5, materiel.Liana:7, materiel.Leaf:5}}
	homeBuilding[fence].b[3] = buildingNature{maxdur:450, timeConsume:plat.Tc{H:2},
		lvlupNeed:map[materiel.Product]int{materiel.Wood:15, materiel.Stone:10, materiel.Liana:10, materiel.Leaf:10, materiel.Metal:3}}
}

func fishInit() {
	homeBuilding[fish].maxlvl = 1
	homeBuilding[fish].b = make([]buildingNature, 2)
	homeBuilding[fish].b[1] = buildingNature{maxdur:100, timeConsume:plat.Tc{H:1},
		lvlupNeed:map[materiel.Product]int{materiel.Wood:5, materiel.Stone:5, materiel.Liana:5}}
}
