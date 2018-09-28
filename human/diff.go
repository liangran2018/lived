package human

type humanFight struct {
	attack int
	defend int
	dodge int
	critical int
}

var hero []humanFight

func init() {
	hero = make([]humanFight, 1)
	hero[0] = humanFight{attack:30, defend:10, dodge:3, critical:3}
}

func (this *human) Attack() int {
	return this.fight.attack
}

func (this *human) Defend() int {
	return this.fight.defend
}

func (this *human) Dodge() int {
	return this.fight.dodge
}

func (this *human) Critical() int {
	return this.fight.critical
}
