package explore

import (
	"fmt"
	"strconv"
	"io/ioutil"
	"encoding/json"

	"github.com/liangran2018/lived/base"
	"github.com/liangran2018/lived/materiel"
	"github.com/liangran2018/lived/log"

	"github.com/gin-gonic/gin"
)

var exploreBag *bag

type bag struct {
	usedCap  int
	max      int
	product map[materiel.Product]int
}

func NewBag() {
	exploreBag = &bag{}
	exploreBag.max = 50
	exploreBag.product = make(map[materiel.Product]int)
}

func GetBag() *bag {
	return exploreBag
}

func (this *bag) UsedCap() int {
	return this.usedCap
}

//剩余
func (this *bag) RemaindCap() int {
	return this.max - this.usedCap
}

func (this *bag) CanPut(weight int) bool {
	if weight <= this.RemaindCap() {
		return true
	}
	return false
}

func (this *bag) Set(p materiel.Product, i int) bool {
	weight := p.Weight() * i

	if this.CanPut(weight) {
		this.product[p] += i
		this.usedCap += weight
		return true
	}

	return false
}

func (this *bag) GetMateriel() map[materiel.Product]int {
	return this.product
}

func (this *bag) Count(p materiel.Product) int {
	return this.product[p]
}

func (this *bag) Clear() {
	this.usedCap = 0

	for p, _ := range this.product {
		delete(this.product, p)
	}
}

func (this *bag) Show() {
	if len(this.product) == 0 {
		fmt.Println("背包为空")
		return
	}

	for k, v := range this.product {
		fmt.Println(k.Name() + "x" + strconv.Itoa(v) + " ")
	}
}

func Bag(c *gin.Context) {
	c.Request.ParseForm()

	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		base.Output(c, base.ParaInvalid, err.Error())
		log.GetLogger().Log(log.Wrong, "bag err", err.Error())
		return
	}
	defer c.Request.Body.Close()

	var data map[int]int
	err = json.Unmarshal(b, &data)
	if err != nil {
		base.Output(c, base.ParaInvalid, err.Error())
		log.GetLogger().Log(log.Wrong, "bag err", err.Error())
		return
	}


}

func Goods() {
	ownThing := materiel.GetOwnThings()
	if ownThing.Nothing() {
		fmt.Println("你一无所有")
		return
	}

	for {
		fmt.Println("1.选择物品  2.准备好了")
		input, err := base.Input()
		if err != nil {
			fmt.Println("输入失败\n")
		}

		if input != "1" {
			return
		}

		p := ownThing.ChooseProduct()
		for k, v := range p {
			fmt.Println(strconv.Itoa(k) + ":" + v.Name())
		}

		input, err = base.Input()
		if err != nil {
			fmt.Println("输入失败\n")
			continue
		}

		i, err := strconv.Atoi(input)
		if err != nil || i < 0 || i >= len(p) {
			fmt.Println("输入错误")
			continue
		}

		c := ownThing.Count(p[i])
		fmt.Printf("有%d个, 输入数量\n", c)
		input, err = base.Input()
		if err != nil {
			fmt.Println("输入失败\n")
			continue
		}

		j, err := strconv.Atoi(input)
		if err != nil || j > c || (j < 0 && GetBag().Count(p[i]) + j < 0) {
			fmt.Println("输入错误")
			continue
		}

		if ok := GetBag().Set(p[i], j); !ok {
			fmt.Println("过重")
			continue
		}

		fmt.Println("放入成功")
		materiel.GetOwnThings().PlusProduct(p[i], j)
	}
}
