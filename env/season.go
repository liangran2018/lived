package env

import (
	"github.com/liangran2018/lived/log"
)

type season int

const (
	spring season = iota
	summer
	autumn
	winter
)

func (this season) Name() string {
	return seasonName[int(this)]
}

func GetSeason() season {
	switch GetTime().Month {
	case 2, 3, 4, 5:
		return spring
	case 6, 7, 8:
		return summer
	case 9, 10:
		return autumn
	default:
		return winter
	}

	log.GetLogger().Log(log.Wrong, "GetSeason err", GetTime().Month)
	return spring
}
