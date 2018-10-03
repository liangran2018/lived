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

func (this *Human) HungryStatus() {
	this.Ill -= 3
	if this.Ill < 0 {
		this.Ill = 0
	}

	this.Mood -= 2
	if this.Mood < 0 {
		this.Blood = 0
	}
}

func (this *Human) HungryChangePerHour() {
	this.Hungry -= 5
	if this.IsFull() {
		this.FullStatus()
	}

	if this.IsHungry() {
		this.HungryStatus()
	}

	if this.Hungry < 0 {
		this.Hungry = 0
	}
}
