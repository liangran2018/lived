package human

import (
	"github.com/liangran2018/lived/base"
	"github.com/liangran2018/lived/log"
)

type Human struct {
	Name   string
	Hero   hero
	Hurt   int
	Ill    int
	Hungry int
	Thirst int
	Mood   int
	Blood  int
	Wake   int

	Lvl      int
	Exp      int
	Attack   int
	Defend   int
	Dodge    int
	Critical int
}

var summer *Human

func NewHuman(name string) {
	summer = allHero[common]
	summer.Name = name
	log.GetLogger().Log(log.Info, "NewHuman", *summer)
}

func LoadHuman(u base.User, name string) {
	summer = &Human{Name: name, Hero: hero(u.Hero), Hurt:u.Hurt, Ill: u.Ill, Hungry: u.Hungry, Thirst: u.Thirst,
		Blood: u.Blood, Mood: u.Mood, Wake: u.Wake, Lvl: u.Lvl, Exp: u.Exp}
	h := allHero[hero(u.Hero)]
	summer.Attack = h.Attack
	summer.Defend = h.Defend
	summer.Dodge = h.Dodge
	summer.Critical = h.Critical

	log.GetLogger().Log(log.Info, "LoadHuman", *summer)
}

func GetHuman() *Human {
	return summer
}

func (this *Human) ExpAdd(i int) {
	this.Exp += i
	log.GetLogger().Log(log.Info, "hero expadd" + base.StrVal(i), this.Exp, this.Lvl)
	if this.Exp >= this.Lvl * 100 {
		this.Exp -= this.Lvl * 100
		this.Lvl++
		allHeroLvlup[this.Hero](this)
		log.GetLogger().Log(log.Info, "hero lvlup", this.Exp, this.Lvl)
	}
}

func (this *Human) ChangePerHour() {
	this.HungryChangePerHour()
	this.ThristChangePerHour()
	this.HurtChangePerHour()
	this.IllChangePerHour()
	this.WakeChangePerHour()
}
