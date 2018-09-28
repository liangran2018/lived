package human

func (this *human) Hungry() int {
	return this.hungry
}

func (this *human) IsHungry() bool {
	return this.hungry <= 25
}

func (this *human) IsFull() bool {
	return this.hungry >= 75
}

func (this *human) FullStatus() {
	if !this.IsHurt() {
		this.hurt++
		if this.hurt > 100 {
			this.hurt = 100
		}
	}

	if !this.IsIll() {
		this.ill += 2
		if this.ill > 100 {
			this.ill = 100
		}
	}

	this.blood += 3
	if this.blood > 100 {
		this.blood = 100
	}
	this.mood += 4
	if this.mood > 100 {
		this.blood = 100
	}
}

func (this *human) HungryChangePerHour() {
	this.hungry -= 5
	if this.IsFull() {
		this.FullStatus()
	}
}

func (this *human) hungryShow() string {
	if this.IsHungry() {
		return "饥饿"
	}

	if this.IsFull() {
		return "饱足"
	}

	return "正常"
}