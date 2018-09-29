package human

type moodStatus int

const (
	normal moodStatus = iota
	lose	//失落
	hard	//难过
	depress //沮丧
	broken	//崩溃

	happy
	crazy
)

func (this *Human) MoodStatus() moodStatus {
	m := this.Mood
	if m >= 95 { return crazy }
	if m < 95 && m >= 80 { return happy }
	if m < 80 && m >= 55 { return normal }
	if m < 55 && m >= 40 { return lose }
	if m < 40 && m >= 25 { return hard }
	if m < 25 && m >= 10 { return depress }
	if m < 10 { return broken }

	return normal
}

func (this *Human) MoodChange(i int) {
	this.Mood += i
	if this.Mood > 100 {
		this.Mood = 100
	}

	if this.Mood < 0 {
		this.Mood = 0
	}
}

func (this *Human) moodShow() string {
	switch this.MoodStatus() {
	case crazy:
		return "欣喜"
	case happy:
		return "开心"
	case lose:
		return "失落"
	case hard:
		return "难过"
	case depress:
		return "沮丧"
	case broken:
		return "崩溃"
	default:
		return "正常"
	}

	panic(this.MoodStatus())
}