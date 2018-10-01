package materiel

import (
	"time"
	"math/rand"
)

type Animal int

type AnimalNature struct {
	meat0 int
	meat1 int
	meat2 int
	Blood int
	Attack int
	Defend int
	Dodge int
	Critical int
	exp   int
}

var animalName []string
var animal map[Animal]AnimalNature

func (this Animal) Name() string {
	return animalName[int(this)]
}

func (this Animal) Hot() AnimalNature {
	return animal[this]
}

func (this Animal) Exp() int {
	return animal[this].exp
}

func (this Animal) Meat() int {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(10)
	if i <= this.Hot().meat0 {
		return 0
	} else if i > this.Hot().meat2 {
		return 2
	} else {
		return 1
	}
}

func animalInit() {
	animalName = make([]string, end)
	animalName = []string{"狼", "瘦弱的狼", "凶猛的狼", "熊", "瘦弱的熊", "凶猛的熊", "蛇"}

	animal = make(map[Animal]AnimalNature, 20)
	animal[Wolf] = AnimalNature{meat0: 2, meat1: 7, meat2: 9, Blood: 60, Attack: 8, Defend: 1, Dodge: 1, Critical: 1, exp: 10}
	animal[WeakWolf] = AnimalNature{meat0: 5, meat1: 8, meat2: 9, Blood: 40, Attack: 6, Defend: 0, Dodge: 0, Critical: 0, exp: 8}
	animal[FierceWolf] = AnimalNature{meat0: 1, meat1: 6, meat2: 9, Blood: 80, Attack: 10, Defend: 1, Dodge: 1, Critical: 1, exp: 12}
	animal[Bear] = AnimalNature{meat0: 1, meat1: 6, meat2: 9, Blood: 90, Attack: 12, Defend: 3, Dodge: 0, Critical: 1, exp: 16}
	animal[WeakBear] = AnimalNature{meat0: 4, meat1: 7, meat2: 9, Blood: 75, Attack: 10, Defend: 1, Dodge: 0, Critical: 0, exp: 13}
	animal[FierceBear] = AnimalNature{meat0: 0, meat1: 5, meat2: 9, Blood: 100, Attack: 15, Defend: 3, Dodge: 1, Critical: 1, exp: 19}
	animal[Snake] = AnimalNature{meat0: 5, meat1: 8, meat2: 9, Blood: 40, Attack: 5, Defend: 0, Dodge: 0, Critical: 0, exp: 8}
}
