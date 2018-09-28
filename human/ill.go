package human

func (this *human) Ill() int {
	return this.ill
}

func (this *human) IsIll() bool {
	return this.ill <= 75
}

func (this *human) IllStatus() {
	this.blood -= 3
	this.mood -= 4
}

func (this *human) IllChangePerHour() {
	if this.IsIll() {
		this.IllStatus()
	} else {
		this.ill++
		if this.ill > 100 {
			this.ill = 100
		}
	}
}

func (this *human) illShow() string {
	if this.IsIll() {
		return "内伤"
	}

	return "健康"
}