package explore

import (
	"os"
	"fmt"
	"time"
	"math/rand"

	"github.com/liangran2018/lived/plat"
	"github.com/liangran2018/lived/human"
	"github.com/liangran2018/lived/base"
	"github.com/liangran2018/lived/materiel"
	"github.com/liangran2018/lived/env"
	"github.com/liangran2018/lived/log"
)

type equipChoose struct {
	name materiel.Product
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

	Goods()

	if GetEquip().e[0] == materiel.ShortBow && GetBag().Count(materiel.Arrow) == 0 {
		fmt.Println("没有携带箭噢")
		NewHeroHot(false)
	} else {
		NewHeroHot(true)
	}


	env.GetTime().Add(explorePlace.NeedTime())

	fmt.Println("出发")

	if t := explorePlace.SetLasttime(env.GetTimeInt().Time()); t != 0 {
		explorePlace.PublicIncrease(t)
	}
	ChooseNext(explorePlace)
}

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
			GetBag().Show()
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
			ok := fight(a)
			if !ok {
				log.GetLogger().Log(log.Warning, "hero dead", a.Name(), a.Hot())
				os.Exit(0)
			}

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

func fight(i materiel.Animal) bool {
	hero := GetHeroHot()
	heroBlood := human.GetHuman().Blood
	animal := i.Hot()
	animalBlood := animal.Blood

	for {
		fmt.Printf("你有%d血量，%s有%d血量\n", heroBlood, i.Name(), animalBlood)

		fmt.Println("你发起进攻")
		blood, die := attack(hero.Attack, hero.Critical, animalBlood, animal.Defend, animal.Dodge)
		if die {
			fmt.Println("捕猎成功")
			return true
		}

		animalBlood = blood
		fmt.Printf("你有%d血量，%s有%d血量\n", heroBlood, i.Name(), animalBlood)

		fmt.Println(i.Name() + "发起反击")
		blood, die = attack(animal.Attack, animal.Critical, heroBlood, hero.Defend, hero.Dodge)
		if die {
			fmt.Println("死亡。。。")
			return false
		}

		heroBlood = blood
	}

}

func attack(attackerAtt, attackerCri, defenderBlo, defenderDef, defenderDod int) (int, bool) {
	if defenderBlo <= 0 {
		return 0, true
	}

	if defenderDod > 0 {
		rand.Seed(time.Now().UnixNano())
		if rand.Intn(10) < defenderDod {
			log.GetLogger().Log(log.Info, "fight dodge")
			return defenderBlo, false
		}
	}

	a := attackerAtt - defenderDef
	if a <= 0 {
		a = 1
	}
	defenderBlo -= a
	if defenderBlo <= 0 {
		return 0, true
	}

	if attackerCri > 0 {
		rand.Seed(time.Now().UnixNano())
		if rand.Intn(10) < attackerCri {
			log.GetLogger().Log(log.Info, "fight critical")
			defenderBlo -= a
		}
	}

	if defenderBlo <= 0 {
		return 0, true
	}

	return defenderBlo, false
}
