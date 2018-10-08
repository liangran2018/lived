package explore

import (
	"fmt"

	"github.com/liangran2018/lived/base"
	"github.com/liangran2018/lived/env"
	"github.com/liangran2018/lived/human"
	"github.com/liangran2018/lived/materiel"
	"github.com/liangran2018/lived/plat"
)

type equipChoose struct {
	name   materiel.Product
	detail []int
}

func init() {
	NewEquip()
	NewBag()
}

func Go() {
	place := plat.GetPublic()

	fmt.Println("请选择：")
	for k, v := range place {
		fmt.Println(base.StrVal(int(k)) + ":" + k.Name() + " " + v.Show())
	}

	input, err := base.Input()
	if err != nil {
		fmt.Println("输入失败")
		return
	}

	i, err := base.IntVal(input)
	if err != nil || i < 0 || i >= len(place) {
		fmt.Println("输入错误")
		return
	}

	explorePlace := place[plat.Place(i)]

	if explorePlace.Lvl() > human.GetHuman().Lvl {
		fmt.Println("你当前等级较低噢")
	}

	//EquipMent()

	if GetEquip().e[0] == materiel.ShortBow {
		fmt.Println("需要携带箭")
	}

	//Goods()

	if GetEquip().e[0] == materiel.ShortBow && GetBag().Count(materiel.Arrow) == 0 {
		fmt.Println("没有携带箭噢")
	}

	env.GetTime().Add(explorePlace.NeedTime())

	fmt.Println("出发")

	if t := explorePlace.SetLasttime(env.GetTimeInt().Time()); t != 0 {
		explorePlace.PublicIncrease(t)
	}

}
/*
func ChooseNext(place *plat.Nature) {
	for {
		fmt.Println("1.查看人物详细状态\n2.查看装备\n3.查看背包\n4.采集\n5.捕猎\n6.回家")
	input:
		input, err := base.Input()
		if err != nil {
			fmt.Println("输入失败")
		}

		if input == "" {
			goto input
		}

		switch input {
		case "1":
			//fmt.Println(human.GetHuman().Detail())
		case "2":
			//GetEquip().Show()
		case "3":
			//GetBag().Show()
		case "4":
			fmt.Println("采集:")
			s := collect(place)
			for k, v := range s {
				fmt.Println(base.StrVal(k) + ":" + v.Name())
			}

			input, err = base.Input()
			if err != nil {
				fmt.Println("输入失败")
				continue
			}

			i, err := base.IntVal(input)
			if err != nil || i < 0 || i > len(s) {
				fmt.Println("输入错误")
				continue
			}

			fmt.Println("输入数量:")
			input, err = base.Input()
			if err != nil {
				fmt.Println("输入失败")
				continue
			}

			j, err := base.IntVal(input)
			if err != nil || j < 0 || j > place.Count(s[i]) {
				fmt.Println("输入错误")
				continue
			}

			ok := GetBag().CanPut(j * s[i].Weight())
			if !ok {
				fmt.Println("超重")
				continue
			}

			place.LostProduct(s[i], j)
			GetBag().Set(s[i], j)
			fmt.Println("采集成功")
		case "5":
			a := hunt(place)
			if a == -1 {
				fmt.Println("无动物")
				continue
			}

			fmt.Println("前方出现一只" + a.Name() + ", 进入战斗")

			log.GetLogger().Log(log.Info, "hunt", a.Name())
			human.GetHuman().ExpAdd(a.Exp())

			place.LostAnimal(a)
			m := a.Meat()
			if m == 0 {
				fmt.Println("运气不好")
				continue
			} else {
				fmt.Printf("获得了%d块肉\n", m)
				ok = GetBag().Set(materiel.Meat, m)
				if !ok {
					if m == 1 {
						fmt.Println("背包满了，装不下了，回家吧")
						continue
					} else {
						ok = GetBag().Set(materiel.Meat, 1)
						if !ok {
							fmt.Println("背包满了，装不下了，回家吧")
							continue
						} else {
							fmt.Println("背包满了，只装下一块，回家吧")
							continue
						}
					}
				}
				fmt.Println("装包成功")
				continue
			}
		case "6":
			env.GetTime().Add(place.NeedTime())
			GetEquip().Clear()
			if len(GetBag().product) == 0 {
				GetBag().Clear()
				return
			}

			ot := materiel.GetOwnThings()
			for k, v := range GetBag().product {
				ot.AddProduct(k, v)
			}
			return
		}
	}
}

func collect(place *plat.Nature) []materiel.Product {
	s := make([]materiel.Product, 0)
	for k, v := range place.Product() {
		if v != 0 {
			s = append(s, k)
		}
	}

	return s
}

func hunt(place *plat.Nature) materiel.Animal {
	m := make([]materiel.Animal, 0)
	for k, v := range place.Animal() {
		if v != 0 {
			m = append(m, k)
		}
	}

	if len(m) == 0 {
		return -1
	}

	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(m))
	return m[i]
}
*/