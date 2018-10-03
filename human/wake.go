package human

func (this *Human) IsSleepy() bool {
	return this.Wake <= 30
}

func (this *Human) IsAwake() bool {
	return this.Wake > 80
}

func (this *Human) WakeChangePerHour() {
	this.Wake -= 4
	if this.Wake < 0 {
		this.Wake = 0
	}
}

func (this *Human) SleepingChangePerHour() {
	if this.IsAwake() {
		this.Blood += 8
		this.Mood += 6
	} else if this.IsSleepy() {
		this.Blood += 4
		this.Mood += 2
	} else {
		this.Blood += 6
		this.Mood += 4
	}

	if this.Blood > 100 {
		this.Blood = 100
	}

	if this.Mood > 100 {
		this.Mood = 100
	}
}
