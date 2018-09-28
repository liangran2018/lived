package plat

import (
	"github.com/liangran2018/lived/materiel"
)

func bushInit() {
	public[gmc] = &Nature{timeConsume:Tc{H:1, Mi:0}, lvl:1}
	public[gmc].product = map[materiel.Product]*count {
		materiel.Wood:&count{maxNum:11, increase:10*60},
		materiel.Pebble:&count{maxNum:12, increase:3*60},
		materiel.Jujube:&count{maxNum:5, increase:15*60},
	}
	public[gmc].animal = map[materiel.Animal]*count {
		materiel.Snake:&count{maxNum:14, increase:9*60},
	}
}

func bushNew() {
	public[lsd].product[materiel.Wood].num = 11
	public[lsd].product[materiel.Pebble].num = 8
	public[lsd].animal[materiel.Jujube].num = 1
	public[lsd].animal[materiel.Snake].num = 12
}
