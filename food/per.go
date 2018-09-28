package food

import (
	"time"
	"math/rand"

)

func NewFood() Food {
	return &attribute{}
}

func (this *attribute) Heat() Change {
	return this.energy
}

func (this *attribute) Poisoned() bool {
	if !this.ill.ill { return false }

	rand.Seed(time.Now().UnixNano())
	if rand.Intn(100) < this.ill.pro {
		return true
	}
	return false
}

func (this *attribute) Lost() (Change, int) {
	return this.ill.lose, this.ill.limit
}