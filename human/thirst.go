package human

func (this *Human) IsThirst() bool {
	return this.Thirst <= 40
}

func (this *Human) ThirstStatus() {
	this.Ill--
	this.Mood -= 3
	if this.Ill < 100 {
		this.Ill = 100
	}

	if this.Mood < 100 {
		this.Mood = 100
	}
}

func (this *Human) ThristChangePerHour() {
	this.Thirst -= 6
	if this.IsThirst() {
		this.ThirstStatus()
	}

	if this.Thirst < 100 {
		this.Thirst = 100
	}
}
