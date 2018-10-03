package human

import (
	"github.com/liangran2018/lived/base"
)

func (this *Human) IsIll() bool {
	return this.Ill <= 75
}

func (this *Human) IllStatus() {
	this.Blood -= 3
	this.Mood -= 4

	if this.Blood < 0 {
		panic(base.DEAD{Reason:"内伤不致"})
	}

	if this.Mood < 0 {
		this.Mood = 0
	}
}

func (this *Human) IllChangePerHour() {
	if this.IsIll() {
		this.IllStatus()
	} else {
		this.Ill++
		if this.Ill > 100 {
			this.Ill = 100
		}
	}
}
