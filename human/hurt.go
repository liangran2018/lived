package human

func (this *Human) IsHurt() bool {
	return this.Hurt <= 60
}

func (this *Human) HurtStatus() {
	this.Blood -= 4
	this.Mood -= 3
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

func (this *Human) hurtShow() string {
	if this.IsHurt() {
		return "重伤"
	}

	if this.Hurt < 80 {
		return "轻伤"
	}

	return "健康"
}
