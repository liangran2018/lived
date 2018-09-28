package human

func (this *human) Thirst() int {
	return this.thirst
}

func (this *human) IsThirst() bool {
	return this.thirst <= 40
}

func (this *human) ThirstStatus() {
	this.ill--
	this.mood -= 3
}

func (this *human) ThristChangePerHour() {
	this.thirst -= 6
	if this.IsThirst() {
		this.ThirstStatus()
	}
}

func (this *human) thristShow() string {
	if this.IsThirst() {
		return "干渴"
	}

	return "正常"
}