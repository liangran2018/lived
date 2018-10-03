package human

import (
	"github.com/liangran2018/lived/base"
)

func (this *Human) IsHurt() bool {
	return this.Hurt <= 60
}

func (this *Human) HurtStatus() {
	this.Blood -= 4
	this.Mood -= 3

	if this.Blood < 0 {
		panic(base.DEAD{Reason:"外伤不致"})
	}

	if this.Mood < 0 {
		this.Mood = 0
	}
}

func (this *Human) HurtChangePerHour() {
	if this.IsHurt() {
		this.HurtStatus()
	} else {
		this.Hurt++
		if this.Hurt > 100 {
			this.Hurt = 100
		}
	}
}
