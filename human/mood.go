package human

func (this *Human) MoodChange(i int) {
	this.Mood += i
	if this.Mood > 100 {
		this.Mood = 100
	}

	if this.Mood < 0 {
		this.Mood = 0
	}
}
