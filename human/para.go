package human

import (
	"github.com/liangran2018/lived/food"

)

type Human interface {
	Name() string
	Show() string
	Detail() string
	Lvl() int
	ChangePerHour()
	ChangePerDay()

	Eat(f food.Food)
	IsPoisoned() bool
	Poisoned()

	Hurt() int
	IsHurt() bool
	HurtStatus()
	//HurtChangePerHour()

	Ill() int
	IsIll() bool
	IllStatus()
	//IllChangePerHour()

	Hungry() int
	IsHungry() bool
	IsFull() bool
	FullStatus()
	//HungryChangePerHour()

	Thirst() int
	IsThirst() bool
	ThirstStatus()
	//ThristChangePerHour()

	Mood() int
	MoodChange(i int)
	MoodStatus() moodStatus
	//MoodChangePerDay()

	Wake() int
}

type human struct {
	name 	string
	hurt    int
	ill     int
	hungry  int
	thirst  int
	mood    int
	blood   int
	wakeful int

	poison  poisonStatus
	lvl 	int
	exp     int
	fight   humanFight
}

type poisonStatus struct {
	isPoison bool
	day 	 int
	Change   food.Change
	limit    int
}

