package env

const (
	spring = iota
	summer
	autumn
	winter
)

func GetSeason() int {
	switch GetTime().Month() {
	case 2, 3, 4, 5:
		return spring
	case 6, 7, 8:
		return summer
	case 9, 10:
		return autumn
	default:
		return winter
	}

	panic(GetTime().Month())
}

func ShowSeason() string {
	switch GetSeason() {
	case spring:
		return "春"
	case summer:
		return "夏"
	case autumn:
		return "秋"
	default:
		return "冬"
	}

	panic(GetSeason())
}
