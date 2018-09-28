package plat

import (
	"github.com/liangran2018/lived/materiel"
)

func riprapInit() {
	public[lsd] = &Nature{timeConsume:Tc{H:0, Mi:40}, lvl:1}
	public[lsd].product = map[materiel.Product]*count {
		materiel.Stone:&count{maxNum:10, increase:6*60},
		materiel.Pebble:&count{maxNum:18, increase:2*60},
	}
	public[lsd].animal = map[materiel.Animal]*count {
		materiel.WeakWolf:&count{maxNum:3, increase:15*60},
		materiel.Snake:&count{maxNum:9, increase:11*60},
	}
}

func riprapNew() {
	public[lsd].product[materiel.Stone].num = 10
	public[lsd].product[materiel.Pebble].num = 15
	public[lsd].animal[materiel.Wolf].num = 2
	public[lsd].animal[materiel.Snake].num = 5
}
