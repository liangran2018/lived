package env

//概率
var weatherPro []map[weather]wl
//对心情影响
var swMood []map[weather]int
var seasonName []string
var weatherName []string

func init() {
	weatherPro = make([]map[weather]wl, 4)
	weatherPro[spring] = map[weather]wl {
		sunny:wl{min:0, max:49},
		cloudy:wl{min:50, max:59},
		foggy:wl{min:60, max:65},
		overcast:wl{min:66, max:73},
		windy:wl{min:74, max:81},
		lightRainy:wl{min:82, max:89},
		heavyRainy:wl{min:90, max:91},
		lightSnowy:wl{min:92, max:100},
	}
	weatherPro[summer] = map[weather]wl {
		sunny:wl{min:0, max:54},
		cloudy:wl{min:55, max:69},
		overcast:wl{min:70, max:74},
		windy:wl{min:75, max:78},
		lightRainy:wl{min:79, max:86},
		heavyRainy:wl{min:87, max:94},
		storm:wl{min:95, max:100},
	}
	weatherPro[autumn] = map[weather]wl {
		sunny:wl{min:0, max:64},
		cloudy:wl{min:65, max:74},
		overcast:wl{min:75, max:82},
		windy:wl{min:83, max:86},
		lightRainy:wl{min:87, max:96},
		heavyRainy:wl{min:97, max:100},
	}
	weatherPro[winter] = map[weather]wl {
		sunny:wl{min:0, max:42},
		cloudy:wl{min:43, max:62},
		foggy:wl{min:63, max:66},
		overcast:wl{min:67, max:70},
		windy:wl{min:71, max:76},
		lightSnowy:wl{min:77, max:86},
		heavySnowy:wl{min:87, max:94},
		storm:wl{min:95, max:100},
	}

	swMood = make([]map[weather]int, 4)
	swMood[spring] = map[weather]int {
		sunny:8,
		cloudy:5,
		foggy:0,
		overcast:0,
		windy:-1,
		lightRainy:-2,
		heavyRainy:-8,
		lightSnowy:-1,
	}
	swMood[summer] = map[weather]int {
		sunny:-3,
		cloudy:0,
		overcast:5,
		windy:4,
		lightRainy:8,
		heavyRainy:-3,
		storm:-5,
	}
	swMood[autumn] = map[weather]int {
		sunny:10,
		cloudy:6,
		overcast:0,
		windy:2,
		lightRainy:-3,
		heavyRainy:-8,
	}
	swMood[winter] = map[weather]int {
		sunny:8,
		cloudy:0,
		foggy:0,
		overcast:-1,
		windy:-3,
		lightSnowy:-3,
		heavySnowy:-6,
		storm:-10,
	}

	seasonName = make([]string, 4)
	seasonName = []string{"春", "夏", "秋", "冬"}

	weatherName = make([]string, 10)
	weatherName = []string{"晴天", "多云", "雾", "阴天", "起风", "小雨", "大雨", "小雪", "大雪", "暴风雨"}
}