package env

import (
	"github.com/liangran2018/lived/base"
	_ "github.com/liangran2018/lived/human"
	"log"
)

var gmTimeInt *gameTimeInt
var gmTime *GameTime

type GameTime struct {
	Year    int `json:"year"`
	Month   int `json:"month"`
	Day     int `json:"day"`
	Hour    int `json:"hour"`
	Minute  int `json:"minute"`
	Overday int `json:"overday"`
}

type gameTimeInt struct {
	t    int
	over int
}

var MonthDay = map[int]int{
	1:  31,
	2:  28,
	3:  31,
	4:  30,
	5:  31,
	6:  30,
	7:  31,
	8:  31,
	9:  30,
	10: 31,
	11: 30,
	12: 31,
}

func NewTime() {
	gmTimeInt = &gameTimeInt{1803010800, 1}
	gmTime = &GameTime{Year: 2018, Month: 3, Day: 1, Hour: 8, Minute: 0, Overday: 1}
}

func LoadTime(t base.Time) {
	gmTimeInt = &gameTimeInt{t: t.Time, over: t.Overday}
	gmTime = gmTimeInt.Int2Time()
}

func GetTimeInt() *gameTimeInt {
	return gmTimeInt
}

func GetTime() *GameTime {
	return gmTime
}

func (this *gameTimeInt) Int2Time() *GameTime {
	t := this.t
	gt := &GameTime{}
	gt.Minute = t % 100
	gt.Hour = (t % 10000) / 100
	gt.Day = (t / 10000) % 100
	gt.Month = (t / 1000000) % 100
	gt.Year = t/100000000 + 2000
	gt.Overday = this.over
	return gt
}

func Int2Time(t int) *GameTime {
	gt := &GameTime{}
	gt.Minute = t % 100
	gt.Hour = (t % 10000) / 100
	gt.Day = (t / 10000) % 100
	gt.Month = (t / 1000000) % 100
	gt.Year = t/100000000 + 2000
	return gt
}

func (this *GameTime) Time2Int() *gameTimeInt {
	gt := &gameTimeInt{}
	gt.t = this.Minute + this.Hour*100 + this.Day*10000 + this.Month*1000000 + (this.Year-2000)*100000000
	gt.over = this.Overday
	return gt
}

func (this *GameTime) Add(h, mi int) {
	for i := 0; i < h; i++ {
		//human.GetHuman().ChangePerHour()
	}

	this.Minute += mi
	if this.Minute >= 60 {
		this.Hour++
		this.Minute -= 60
		//human.GetHuman().ChangePerHour()
	}

	this.Hour += h
	if this.Hour >= 24 {
		this.Day++
		this.Hour -= 24
		this.Overday++
		log.Printf("第%d天\n", this.Overday)
		NewWeather()
		//	human.GetHuman().ChangePerDay()
		MoodChangePerDay()
	}

	d, _ := MonthDay[this.Month]
	if this.Day > d {
		this.Month++
		this.Day = 1
	}

	if this.Month > 12 {
		this.Year++
		this.Month = 1
	}

	gmTimeInt = this.Time2Int()
}

func (this *gameTimeInt) Time() int {
	return this.t
}
