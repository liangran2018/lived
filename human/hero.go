package human

type hero int

var allHero map[hero]*Human

const (
	common hero = iota
)

func init() {
	allHero = make(map[hero]*Human, 1)

	allHero[common] = &Human{Hurt:100, Ill:100, Hungry:80, Thirst:80, Mood:90, Blood:100, Wake:100,
		Lvl:1, Exp:0, Attack:15, Defend:5, Dodge:1, Critical:1}
}