package human

func (this *human) Hurt() int {
	return this.hurt
}

func (this *human) IsHurt() bool {
	return this.hurt <= 60
}

func (this *human) HurtStatus() {
	this.blood -= 4
	this.mood -= 3
}

func (this *human) HurtChangePerHour() {
	if this.IsHurt() {
		this.HurtStatus()
	} else {
		this.hurt++
		if this.hurt > 100 {
			this.hurt = 100
		}
	}
}

func (this *human) hurtShow() string {
	if this.IsHurt() {
		return "重伤"
	}

	if this.hurt < 80 {
		return "轻伤"
	}

	return "健康"
}
