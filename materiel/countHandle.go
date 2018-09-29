package materiel

import (
	"io/ioutil"
	"encoding/json"

	"github.com/liangran2018/lived/base"
	"github.com/liangran2018/lived/log"

	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
	c.Request.ParseForm()

	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		base.Output(c, base.ParaInvalid, err.Error())
		log.GetLogger().Log(log.Wrong, "ownthingadd err", err.Error())
		return
	}
	defer c.Request.Body.Close()

	var data map[int]int
	err = json.Unmarshal(b, &data)
	if err != nil {
		base.Output(c, base.ParaInvalid, err.Error())
		log.GetLogger().Log(log.Wrong, "ownthingadd err", err.Error())
		return
	}

	ot := GetOwnThings()
	for k, v := range data {
		ot.AddProduct(Product(k), v)
	}

	base.Output(c, 0, nil)
	return
}

func Plus(c *gin.Context) {
	c.Request.ParseForm()

	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		base.Output(c, base.ParaInvalid, err.Error())
		log.GetLogger().Log(log.Wrong, "ownthingadd err", err.Error())
		return
	}
	defer c.Request.Body.Close()

	var data map[int]int
	err = json.Unmarshal(b, &data)
	if err != nil {
		base.Output(c, base.ParaInvalid, err.Error())
		log.GetLogger().Log(log.Wrong, "ownthingadd err", err.Error())
		return
	}

	ot := GetOwnThings()
	for k, v := range data {
		ot.PlusProduct(Product(k), v)
	}

	base.Output(c, 0, nil)
	return
}
