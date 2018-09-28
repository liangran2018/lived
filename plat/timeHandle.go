package plat

import (
	"github.com/liangran2018/lived/env"
)

func (this *Nature) NeedTime() (int, int) {
	return this.timeConsume.H, this.timeConsume.Mi
}

func (this *Nature) LastTime() int {
	return this.lastTime
}

func (this *Nature) SetLasttime(i int) int {
	if this.lastTime == 0 {
		this.lastTime = i
		return 0
	}

	newTime := env.Int2Time(i)
	oldTime := env.Int2Time(this.lastTime)
	this.lastTime = i

	plus := 0
	if newTime.Year() > oldTime.Year() {
		plus += 365
		if newTime.Month() < oldTime.Month() {
			for j:=newTime.Month(); j<oldTime.Month(); j++ {
				plus -= env.MonthDay[j]
			}
		}
	}

	if newTime.Month() > oldTime.Month() {
		for j:=oldTime.Month(); j<newTime.Month(); j++ {
			plus += env.MonthDay[j]
		}
	}

	if newTime.Day() != oldTime.Day() {
		if newTime.Day() < oldTime.Day() {
			plus -= oldTime.Day() - newTime.Day()
		} else {
			plus += newTime.Day() - oldTime.Day()
		}
	}

	plus = plus * 24
	if newTime.Hour() != oldTime.Hour() {
		if newTime.Hour() < oldTime.Hour() {
			plus -= oldTime.Hour() - newTime.Hour()
		} else {
			plus += newTime.Hour() - oldTime.Hour()
		}
	}

	plus = plus * 60
	if newTime.Minute() != oldTime.Minute() {
		if newTime.Minute() < oldTime.Minute() {
			plus -= oldTime.Minute() - newTime.Minute()
		} else {
			plus += newTime.Minute() - oldTime.Minute()
		}
	}
	return plus
}
