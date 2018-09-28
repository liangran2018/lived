package food

type Food interface {
	Heat() Change
	Poisoned() bool
	Lost() (Change, int)
}

type attribute struct {
	energy Change
	ill illPara
}

type Change struct {
	Blood int
	Mood  int
	Wake  int
}

type illPara struct {
	ill bool
	pro int
	lose Change
	limit int
}
