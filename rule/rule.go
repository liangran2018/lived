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

/*
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
*/