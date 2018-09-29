package human

func (this *Human) IsIll() bool {
	return this.Ill <= 75
}

func (this *Human) IllStatus() {
	this.Blood -= 3
	this.Mood -= 4
}

func (this *Human) IllChangePerHour() {
	if this.IsIll() {
		this.IllStatus()
	} else {
		this.Ill++
		if this.Ill > 100 {
			this.Ill = 100
		}
	}
}

func (this *Human) illShow() string {
	if this.IsIll() {
		return "内伤"
	}

	return "健康"
}