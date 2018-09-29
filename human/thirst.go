package human

func (this *Human) IsThirst() bool {
	return this.Thirst <= 40
}

func (this *Human) ThirstStatus() {
	this.Ill--
	this.Mood -= 3
}

func (this *Human) ThristChangePerHour() {
	this.Thirst -= 6
	if this.IsThirst() {
		this.ThirstStatus()
	}
}

func (this *Human) thristShow() string {
	if this.IsThirst() {
		return "口渴"
	}

	return "正常"
}