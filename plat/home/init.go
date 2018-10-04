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

var doing []func() *outputBuild

var convert []string

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
	doing[bed] = sleep
	doing[fire] = cook
	doing[water] = clean
	doing[drug] = medicine
	doing[tool] = equip
	doing[field] = grow
	doing[fence] = rail
	doing[fish] = fishing

	actionNature = make(map[action]actionLimit, actionEnd)
	actionNature[sleep1H] = actionLimit{lvl:1, t:bed}
	actionNature[sleep4H] = actionLimit{lvl:1, t:bed}
	actionNature[sleep8H] = actionLimit{lvl:1, t:bed}

	actionNature[cookBbq] = actionLimit{lvl:1, t:fire}
	actionNature[cookRoastPotato] = actionLimit{lvl:1, t:fire}
	actionNature[cookCongee] = actionLimit{lvl:1, t:fire}
	actionNature[cookBroth] = actionLimit{lvl:2, t:fire}
	actionNature[cookMashedPotato] = actionLimit{lvl:2, t:fire}
	actionNature[cookStew] = actionLimit{lvl:2, t:fire}
	actionNature[cookDriedFish] = actionLimit{lvl:3, t:fire}
	actionNature[cookSmokedMeet] = actionLimit{lvl:3, t:fire}
	actionNature[cookAnimalBlood] = actionLimit{lvl:3, t:fire}

	actionNature[distill] = actionLimit{lvl:1, t:water}
	actionNature[filterWater] = actionLimit{lvl:1, t:water}
	actionNature[makeWine] = actionLimit{lvl:1, t:water}

	actionNature[makeBandage] = actionLimit{lvl:1, t:drug}
	actionNature[makeDecoction] = actionLimit{lvl:1, t:drug}
	actionNature[makeMintTea] = actionLimit{lvl:1, t:drug}
	actionNature[makeMediWine] = actionLimit{lvl:2, t:drug}
	actionNature[makeFirstAid] = actionLimit{lvl:2, t:drug}
	actionNature[makeGrassPaste] = actionLimit{lvl:2, t:drug}
	actionNature[makePlaster] = actionLimit{lvl:3, t:drug}
	actionNature[makeTonifyPill] = actionLimit{lvl:3, t:drug}

	actionNature[makeStoneAxe] = actionLimit{lvl:1, t:tool}
	actionNature[makeTorch] = actionLimit{lvl:1, t:tool}
	actionNature[makeKnife] = actionLimit{lvl:1, t:tool}
	actionNature[makeShortBow] = actionLimit{lvl:1, t:tool}
	actionNature[make20Arrow] = actionLimit{lvl:1, t:tool}
	actionNature[makeBambooGun] = actionLimit{lvl:2, t:tool}
	actionNature[makeClothArmor] = actionLimit{lvl:2, t:tool}
	actionNature[makeTwoEdgedAxe] = actionLimit{lvl:2, t:tool}
	actionNature[makeHardBambooBow] = actionLimit{lvl:2, t:tool}
	actionNature[makeBigBag] = actionLimit{lvl:2, t:tool}
	actionNature[makeHardBambooCrossBow] = actionLimit{lvl:3, t:tool}
	actionNature[makeSharpTwoAxe] = actionLimit{lvl:3, t:tool}
	actionNature[makeSharpBamboo] = actionLimit{lvl:3, t:tool}
	actionNature[makePoisonAxe] = actionLimit{lvl:3, t:tool}
	actionNature[makePoisonBamboo] = actionLimit{lvl:3, t:tool}
	actionNature[makeRattanArmor] = actionLimit{lvl:3, t:tool}

	actionNature[growFood] = actionLimit{lvl:1, t:field}
	actionNature[growPotato] = actionLimit{lvl:1, t:field}
	actionNature[growMint] = actionLimit{lvl:2, t:field}

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
