package frame

import (
	"time"

	"github.com/liangran2018/lived/base"
	"github.com/liangran2018/lived/env"
	"github.com/liangran2018/lived/log"

	"github.com/gin-gonic/gin"
)

var data base.Data

func init() {
	data = base.Data{}
	data.OwnBuild = make([]int, 6)
	data.OwnProduct = make(map[int]int)
	data.PlatLastTime = make(map[int]int)
	data.PlatProduct = make(map[int]map[int]int)
	data.PlatAnimal = make(map[int]map[int]int)
}

const defaultName = "一一"

func NewGame(c *gin.Context) {
	name := c.Query("name")
	if base.Empty(name) {
		name = defaultName
	}

	//开始时间
	data.StartTime = time.Now().Format("2006-01-02 15:04:05")
	//上次游戏时间
	data.LastTime = time.Now().Format("2006-01-02 15:04:05")
	//玩家昵称
	data.Name = name
	//重置游戏内时间
	env.NewTime()
	//获取天气
	env.NewWeather()
	//新建人物
	//human.NewHuman(name)
	//营地起始建筑，只有床
	//home.NewOwnBuilding()
	//起始拥有物品
	//materiel.NewOwnThings()
	//地图初始化
	//plat.NewPublic()
	//新建日志文件
	log.NewLogFile()
	//记录
	log.GetLogger().Log(log.Info, "newgame start", name)


	base.Output(c, 0, nil)
}
