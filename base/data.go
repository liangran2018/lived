package base

type Data struct {
	StartTime    string              `json:"startTime"`
	LastTime     string              `json:"lastTime"`
	Name         string              `json:"name"`
	GameTime     Time                `json:"gameTime"`
	Weather      int				 `json:"weather"`
	Temprature   int 				 `json:"temprature"`
	User         User                `json:"user"`
	OwnBuild     []OB                `json:"ownBuild"`
	OwnProduct   map[int]int         `json:"ownMateriel"`
	BigBag       bool				 `json:"bigbag"`
	PlatLastTime map[int]int         `json:"platLastTime"`
	PlatProduct  map[int]map[int]int `json:"platProduct"`
	PlatAnimal   map[int]map[int]int `json:"platAnimal"`
}

type OB struct {
	Lvl int `json:"lvl"`
	Dur int `json:"dur"`
}

type Time struct {
	Time    int `json:"time"`
	Overday int `json:"overday"`
}

type User struct {
	Hurt   int `json:"hurt"`
	Ill    int `json:"ill"`
	Hungry int `json:"hungry"`
	Thirst int `json:"thirst"`
	Blood  int `json:"blood"`
	Mood   int `json:"mood"`
	Wake   int `json:"wake"`
	Lvl    int `json:"lvl"`
	Exp    int `json:"exp"`
	Hero   int `json:"hero"`
}
