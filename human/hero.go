package human

type hero int

var allHero map[hero]*Human
var allHeroLvlup map[hero]func(this *Human)

const (
	common hero = iota
)

func init() {
	allHero = make(map[hero]*Human, 1)
	allHeroLvlup = make(map[hero]func(this *Human), 1)

	allHero[common] = &Human{Hero:common, Hurt:100, Ill:100, Hungry:80, Thirst:80, Mood:90, Blood:100, Wake:100,
		Lvl:1, Exp:0, Attack:25, Defend:10, Dodge:1, Critical:1}

	allHeroLvlup[common] = Lvlup
}

func Lvlup(this *Human) {
	this.Attack += 3
	this.Defend++
	if this.Lvl % 5 == 0 {
		if this.Critical < 5 {
			this.Critical++
		}

		if this.Dodge < 5 {
			this.Dodge++
		}
	}
	this.Blood += 10
	if this.Blood > 100 {
		this.Blood = 100
	}
}