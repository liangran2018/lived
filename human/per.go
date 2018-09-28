package human

import (
	"fmt"
	"github.com/liangran2018/lived/food"
	"github.com/liangran2018/lived/base"
)

var summer *human

func NewHuman(name string) {
	summer = &human{name:name, hurt:100, ill:100, hungry:80, thirst:80, mood:90, blood:100, wakeful:100,
		lvl:1, exp:0, fight:hero[0]}
}

func LoadHuman(h base.Human, name string) {
	summer = &human{name:name, hurt:h.Hurt, ill:h.Ill, hungry:h.Hungry, thirst:h.Thirst,
		blood:h.Blood, mood:h.Mood, wakeful:h.Wake, lvl:h.Lvl, exp:h.Exp, fight:hero[0]}
}

func GetHuman() *human {
	return summer
}

func (this *human) Name() string {
	return this.name
}

func (this *human) Show() string {
	return fmt.Sprintf("外伤:%s, 内伤:%s, 饥饿:%s, 干渴:%s, 血量:%d, 心情:%s, 失眠度:%d", this.hurtShow(),
		this.illShow(), this.hungryShow(), this.thristShow(), this.blood, this.moodShow(), this.wakeful)
}

func (this *human) Detail() string {
	s := fmt.Sprintf("外伤值:%d, 内伤值:%d, 饥饿值:%d, 干渴值:%d, 血量:%d, 心情值:%d, 失眠值:%d", this.hurt,
		this.ill, this.hungry, this.thirst, this.blood, this.mood, this.wakeful)
	if this.IsPoisoned() {
		s += fmt.Sprintf("中毒: 第%d天, 共%d天, 每天失去血量:%d, 心情值:%d, 失眠值:%d, 内伤值:%d", this.poison.day,
			this.poison.limit, this.poison.Change.Blood, this.poison.Change.Mood, this.poison.Change.Wake, this.poison.Change.Blood/2)
	}

	s += "\n"
	return s
}

func (this *human) Lvl() int {
	return this.lvl
}

func (this *human) Exp() int {
	return this.exp
}

func (this *human) ExpAdd(i int) {
	this.exp += i
	if this.exp >= this.lvl * 100 {
		this.exp -= this.lvl * 100
		this.lvl++
		fmt.Printf("等级提升：lvl.%d\n", this.lvl)
	}
}

func (this *human) Eat(f food.Food) {
	h := f.Heat()
	this.blood += h.Blood
	this.mood += h.Mood
	this.wakeful += h.Wake

	if f.Poisoned() {
		this.poison.isPoison = true
		this.poison.Change, this.poison.limit = f.Lost()
		this.poison.day = 1
		this.Poisoned()
	}
}

func (this *human) IsPoisoned() bool {
	return this.poison.isPoison
}

func (this *poisonStatus) clearPoison() {
	this.isPoison = false
	this.day = 0
	this.limit = 0
	this.Change = food.Change{0,0,0}
	return
}

func (this *human) lost() {
	this.blood -= this.poison.Change.Blood
	this.mood -= this.poison.Change.Mood
	this.wakeful -= this.poison.Change.Wake
	this.ill -= this.poison.Change.Blood/2
}

func (this *human) Poisoned() {
	if this.poison.day > this.poison.limit {
		this.poison.clearPoison()
		return
	}

	this.lost()
}

func (this *human) Blood() int {
	return this.blood
}

func (this *human) Dead() bool {
	return this.blood <= 0
}

func (this *human) ChangePerHour() {
	this.HungryChangePerHour()
	this.ThristChangePerHour()
	this.HurtChangePerHour()
	this.IllChangePerHour()
}

func (this *human) ChangePerDay() {
	//this.MoodChangePerDay()
	this.Poisoned()
}