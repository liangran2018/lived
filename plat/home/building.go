package home

import (
	"fmt"
	"github.com/liangran2018/lived/materiel"
	"github.com/liangran2018/lived/plat"
)

type buildingNature struct {
	dur int
	maxdur int
	lvl int
	maxlvl int
	lvlupNeed map[materiel.Product]int
	timeConsume plat.Tc
}

type OwnBuilding struct {
	own []int
}

var obl *OwnBuilding

//不可变， 设施属性， 第一个参数是种类（room/cook等）， 第二个是等级，从1开始
var homeBuilding [][]buildingNature

//玩家拥有的建筑，key是种类， value等级
//var OwnBuilding []int

var convert []string

const (
	_ = iota
	room
	cook
	water
	drug
	tool
)

func init() {
	NewHomeBuilding()
	convert = []string{"", "床", "火堆", "净水器", "药盒", "工具台"}
	obl = &OwnBuilding{}
	obl.own = make([]int, 6)
}

func NewOwnBuilding() {
	obl.own[room] = 1
}

func LoadOwnBuilding(b []int) {
	for k, v := range b {
		obl.own[k] = v
	}
}

func GetOwnBuilding() *OwnBuilding {
	return obl
}

func (this *OwnBuilding) Get() []int {
	return this.own
}

func NewHomeBuilding() {
	homeBuilding = make([][]buildingNature, 10)
	roomInit()
	cookInit()
	waterInit()
	drugInit()
	toolInit()
}

func roomInit() {
	homeBuilding[room] = make([]buildingNature, 6)
	for i:=1; i<=5; i++ {
		homeBuilding[room][i] = buildingNature{dur:100+i*50, maxdur:100+i*50, lvl:i, maxlvl:5}
		homeBuilding[room][i].timeConsume.H = i/2
		if i%2 == 0 {
			homeBuilding[room][i].timeConsume.Mi = 0
		} else {
			homeBuilding[room][i].timeConsume.Mi = 30
		}
	}

	homeBuilding[room][1].lvlupNeed = map[materiel.Product]int{materiel.Wood:2, materiel.Pebble:1}
	homeBuilding[room][2].lvlupNeed = map[materiel.Product]int{materiel.Wood:3, materiel.Pebble:3}
	homeBuilding[room][3].lvlupNeed = map[materiel.Product]int{materiel.Wood:5, materiel.Stone:5}
	homeBuilding[room][4].lvlupNeed = map[materiel.Product]int{materiel.Wood:8, materiel.Cloth:3, materiel.Liana:5}
	homeBuilding[room][5].lvlupNeed = map[materiel.Product]int{materiel.Wood:5, materiel.Cloth:5, materiel.Liana:5, materiel.Grass:15}
}

func cookInit() {
	homeBuilding[cook] = make([]buildingNature, 6)
	for i:=1; i<=5; i++ {
		homeBuilding[cook][i] = buildingNature{dur:50+i*20, maxdur:50+i*20, lvl:i, maxlvl:4}
		homeBuilding[cook][i].timeConsume.H = i/2
		if i%2 == 0 {
			homeBuilding[cook][i].timeConsume.Mi = 0
		} else {
			homeBuilding[cook][i].timeConsume.Mi = 30
		}
	}

	homeBuilding[cook][1].lvlupNeed = map[materiel.Product]int{materiel.Wood:1, materiel.Pebble:2}
	homeBuilding[cook][2].lvlupNeed = map[materiel.Product]int{materiel.Wood:3, materiel.Stone:2}
	homeBuilding[cook][3].lvlupNeed = map[materiel.Product]int{materiel.Wood:5, materiel.Stone:3, materiel.Cloth:3}
	homeBuilding[cook][4].lvlupNeed = map[materiel.Product]int{materiel.Wood:10, materiel.Stone:5, materiel.Liana:6, materiel.Grass:8}
}

func waterInit() {
	homeBuilding[water] = make([]buildingNature, 2)
	homeBuilding[water][1] = buildingNature{dur:100, maxdur:100, lvl:1, maxlvl:1, timeConsume:plat.Tc{H:1, Mi:0},
		lvlupNeed:map[materiel.Product]int{materiel.Wood:3, materiel.Stone:3, materiel.Cloth:5, materiel.Liana:2}}
}

func drugInit() {
	homeBuilding[drug] = make([]buildingNature, 6)
	for i:=1; i<=5; i++ {
		homeBuilding[drug][i] = buildingNature{dur:50+i*30, maxdur:100, lvl:i, maxlvl:4}
		homeBuilding[drug][i].timeConsume.H = i
	}

	homeBuilding[drug][1].lvlupNeed = map[materiel.Product]int{materiel.Wood:1, materiel.Pebble:1, materiel.Liana:1}
	homeBuilding[drug][2].lvlupNeed = map[materiel.Product]int{materiel.Wood:1, materiel.Pebble:2, materiel.Stone:1, materiel.Liana:2}
	homeBuilding[drug][3].lvlupNeed = map[materiel.Product]int{materiel.Wood:2, materiel.Pebble:3, materiel.Stone:1, materiel.Liana:3}
	homeBuilding[drug][4].lvlupNeed = map[materiel.Product]int{materiel.Wood:2, materiel.Pebble:4, materiel.Stone:2, materiel.Liana:4}
}

func toolInit() {
	homeBuilding[tool] = make([]buildingNature, 6)
	for i:=1; i<=5; i++ {
		homeBuilding[tool][i] = buildingNature{dur:150+i*50, maxdur:150+i*50, lvl:i, maxlvl:5}
	}
	homeBuilding[tool][1].timeConsume = plat.Tc{H:1, Mi:0}
	homeBuilding[tool][2].timeConsume = plat.Tc{H:1, Mi:30}
	homeBuilding[tool][3].timeConsume = plat.Tc{H:2, Mi:30}
	homeBuilding[tool][4].timeConsume = plat.Tc{H:3, Mi:30}
	homeBuilding[tool][5].timeConsume = plat.Tc{H:5, Mi:0}

	homeBuilding[tool][1].lvlupNeed = map[materiel.Product]int{materiel.Wood:2, materiel.Stone:1, materiel.Liana:3, materiel.Metal:1}
	homeBuilding[tool][2].lvlupNeed = map[materiel.Product]int{materiel.Wood:2, materiel.Stone:2, materiel.Liana:5, materiel.Metal:2}
	homeBuilding[tool][3].lvlupNeed = map[materiel.Product]int{materiel.Wood:2, materiel.Stone:3, materiel.Liana:8, materiel.Metal:2, materiel.Gold:1}
	homeBuilding[tool][4].lvlupNeed = map[materiel.Product]int{materiel.Liana:10, materiel.Metal:3, materiel.Gold:3}
	homeBuilding[tool][5].lvlupNeed = map[materiel.Product]int{materiel.Liana:15, materiel.Metal:5, materiel.Gold:5}
}

func (this *OwnBuilding) Show() string {
	s := "你拥有："
	for k, v := range this.own {
		if v != 0 {
			s += fmt.Sprintf("%d级 %s ", v, convert[k])
		}
	}

	return s
}

func (this *OwnBuilding) lvlupNotice(i int) string {
	l := this.own[i]
	if l == homeBuilding[i][l].maxlvl {
		return convert[i] + "已满级"
	}

	s := fmt.Sprintf("需要时间：%d小时%d分钟，需要材料：", homeBuilding[i][l+1].timeConsume.H, homeBuilding[i][l+1].timeConsume.Mi)
	for k, v := range homeBuilding[i][l+1].lvlupNeed {
		s += fmt.Sprintf("%s %d个 ", k.Name(), v)
	}

	return s
}

func (this *OwnBuilding) Check(i int) bool {
	l := this.own[i]
	om := materiel.GetOwnThings().OwnProduct()

	for k, v := range homeBuilding[i][l+1].lvlupNeed {
		if om[k] < v {
			return false
		}
	}

	return true
}

func (this *OwnBuilding) Update(i int) {
	l := this.own[i]
	om := materiel.GetOwnThings().OwnProduct()

	for k, v := range homeBuilding[i][l+1].lvlupNeed {
		om[k] -= v
	}
}

func (this *OwnBuilding) TimeNeed(i int) (int, int) {
	l := this.own[i]
	tc := homeBuilding[i][l+1].timeConsume
	return tc.H, tc.Mi
}
