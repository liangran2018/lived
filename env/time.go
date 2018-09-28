package env

import (
	"fmt"
	"log"
	"github.com/liangran2018/lived/human"
	"github.com/liangran2018/lived/base"
)

var gmTimeInt *gameTimeInt
var gmTime *gameTime

type gameTime struct {
	y int
	m int
	d int
	h int
	mi int
	over int
}

type gameTimeInt struct {
	t int
	over int
}

var MonthDay = map[int]int{
	1: 31,
	2: 28,
	3: 31,
	4: 30,
	5: 31,
	6: 30,
	7: 31,
	8: 31,
	9: 30,
	10: 31,
	11: 30,
	12: 31,
}

func NewTime() {
	gmTimeInt = &gameTimeInt{201803010800, 1}
	gmTime = &gameTime{y:2018, m:3, d:1, h:8, mi:0, over:1}
}

func LoadTime(t base.Time) {
	gmTimeInt = &gameTimeInt{t:t.Time, over:t.Overday}
	gmTime = gmTimeInt.Int2Time()
}

func GetTimeInt() *gameTimeInt {
	return gmTimeInt
}

func GetTime() *gameTime {
	return gmTime
}

func (this *gameTimeInt) Int2Time() *gameTime {
	t := this.t
	gt := &gameTime{}
	gt.mi = t%100
	gt.h = (t%10000)/100
	gt.d = (t/10000)%100
	gt.m = (t/1000000)%100
	gt.y = t/100000000
	gt.over = this.over
	return gt
}

func Int2Time(t int) *gameTime {
	gt := &gameTime{}
	gt.mi = t%100
	gt.h = (t%10000)/100
	gt.d = (t/10000)%100
	gt.m = (t/1000000)%100
	gt.y = t/100000000
	return gt
}

func (this *gameTime) Time2Int() *gameTimeInt {
	gt := &gameTimeInt{}
	gt.t = this.mi + this.h * 100 + this.d * 10000 + this.m * 1000000 + this.y * 100000000
	gt.over = this.over
	return gt
}

func (this *gameTime) Show() string {
	return fmt.Sprintf("%d-%d-%d %d:%d, 第%d天", this.y, this.m, this.d, this.h, this.mi, this.over)
}

func (this *gameTime) Add(h, mi int) {
	for i:=0; i<h; i++ {
		human.GetHuman().ChangePerHour()
	}

	this.mi += mi
	if this.mi >= 60 {
		this.h++
		this.mi -= 60
		human.GetHuman().ChangePerHour()
	}

	this.h += h
	if this.h >= 24 {
		this.d++
		this.h -= 24
		this.over++
		log.Printf("第%d天\n", this.over)
		NewWeather()
		human.GetHuman().ChangePerDay()
		MoodChangePerDay()
	}

	d, _ := MonthDay[this.m]
	if this.d > d {
		this.m++
		this.d = 1
	}

	if this.m > 12 {
		this.y++
		this.m = 1
	}

	gmTimeInt = this.Time2Int()
}

func (this *gameTimeInt) Time() int {
	return this.t
}

func (this *gameTimeInt) Over() int {
	return this.over
}

func (this *gameTime) Month() int {
	return this.m
}

func (this *gameTime) Year() int {
	return this.y
}

func (this *gameTime) Day() int {
	return this.d
}

func (this *gameTime) Hour() int {
	return this.h
}

func (this *gameTime) Minute() int {
	return this.mi
}
