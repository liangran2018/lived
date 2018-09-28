package rule

import (
	"fmt"
	"time"
	"encoding/json"

	"github.com/liangran2018/lived/base"
	"github.com/liangran2018/lived/env"
	"github.com/liangran2018/lived/human"
	"github.com/liangran2018/lived/plat/home"
	"github.com/liangran2018/lived/materiel"
	"github.com/liangran2018/lived/plat"
	"github.com/liangran2018/lived/log"
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

func NewGame(c *gin.Context) {
	//开始时间
	data.StartTime = time.Now().Format("2006-01-02 15:04:05")
	//上次游戏时间
	data.LastTime = time.Now().Format("2006-01-02 15:04:05")
	//英雄昵称
	data.Name = name
	//重置游戏内时间
	env.NewTime()
	//获取天气
	env.NewWeather()
	//新建人物
	human.NewHuman(name)
	//营地起始建筑，只有床
	home.NewOwnBuilding()
	//起始拥有物品
	materiel.NewOwnThings()
	//地图初始化
	plat.NewPublic()
	//新建日志文件
	log.NewLogFile()
	//记录
	log.GetLogger().Log(log.Info, "newgame start", name)
}

func Load(d *base.Data) {
	//加载游戏时间
	env.LoadTime(d.GameTime)
	//加载人物
	human.LoadHuman(d.Human, d.Name)
	//加载拥有建筑
	home.LoadOwnBuilding(d.OwnBuild)
	//加载拥有物品
	materiel.LoadOwnThings(d.OwnProduct)
	//更新各地点上次到访时间
	plat.LoadPublic(d.PlatLastTime)
	//更新各地点物品数量
	plat.LoadProduct(d.PlatProduct)
	//更新各地点动物数量
	plat.LoadAnimal(d.PlatAnimal)

	data.StartTime = d.StartTime
	//上次游戏时间更新
	data.LastTime = time.Now().Format("2006-01-02 15:04:05")
	data.Name = d.Name
}

func Save(file string) error {
	h := human.GetHuman()
	data.Human = base.Human{Hurt:h.Hurt(), Ill:h.Ill(), Hungry:h.Hungry(), Thirst:h.Thirst(),
		Blood:h.Blood(), Mood:h.Mood(), Wake:h.Wake(), Lvl:h.Lvl(), Exp:h.Exp()}
	data.GameTime = base.Time{Time:env.GetTimeInt().Time(), Overday:env.GetTimeInt().Over()}
	data.OwnBuild = home.GetOwnBuilding().Get()

	for k, v := range materiel.GetOwnThings().OwnProduct() {
		data.OwnProduct[int(k)] = v
	}

	pl := plat.GetPublic()
	for p, q := range pl {
		data.PlatLastTime[int(p)] = q.LastTime()
		data.PlatProduct[int(p)] = make(map[int]int)
		for a, b := range q.Product() {
			data.PlatProduct[int(p)][int(a)] = b
		}

		data.PlatAnimal[int(p)] = make(map[int]int, len(q.Animal()))
		for c, b := range q.Animal() {
			data.PlatAnimal[int(p)][int(c)] = b
		}
	}

	err := base.DeleteFile(file)
	if err != nil {
		log.GetLogger().Log(log.Wrong, "删除文件失败", err.Error())
		return err
	}

	f := base.NewFile()
	d, _  := json.Marshal(data)
	_, err = f.Write(d)
	if err != nil {
		log.GetLogger().Log(log.Wrong, "保存数据失败", err.Error())
		return err
	}
	return nil
}

func Show() {
	fmt.Printf("\n%s %s %s %s\n", human.GetHuman().Name(), env.GetTime().Show(), env.ShowSeason(), env.ShowWeather())
	fmt.Println(human.GetHuman().Show() + "\n")
}
