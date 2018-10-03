package frame

import (
	"github.com/liangran2018/lived/base"
)

var data base.Data

func init() {
	data = base.Data{}
	data.OwnBuild = make([]base.OB, 6)
	data.OwnProduct = make(map[int]int)
	data.PlatLastTime = make(map[int]int)
	data.PlatProduct = make(map[int]map[int]int)
	data.PlatAnimal = make(map[int]map[int]int)
}