package human

func (this *Human) IsHungry() bool {
	return this.Hungry <= 25
}

func (this *Human) IsFull() bool {
	return this.Hungry >= 75
}

func (this *Human) FullStatus() {
	if !this.IsHurt() {
		this.Hurt++
		if this.Hurt > 100 {
			this.Hurt = 100
		}
	}

	if !this.IsIll() {
		this.Ill += 2
		if this.Ill > 100 {
			this.Ill = 100
		}
	}

	this.Blood += 3
	if this.Blood > 100 {
		this.Blood = 100
	}

	this.Mood += 4
	if this.Mood > 100 {
		this.Blood = 100
	}
}

func (this *Human) HungryChangePerHour() {
	this.Hungry -= 5
	if this.IsFull() {
		this.FullStatus()
	}
}

func (this *Human) hungryShow() string {
	if this.IsHungry() {
		return "饥饿"
	}

	if this.IsFull() {
		return "饱足"
	}

	return "正常"
}