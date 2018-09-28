package base

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Data struct {
	StartTime   string `json:"startTime"`
	LastTime	string `json:"lastTime"`
	Name        string `json:"name"`
	GameTime    Time   `json:"gameTime"`
	Human       Human  `json:"human"`
	OwnBuild    []int  `json:"ownBuild"`
	OwnProduct  map[int]int  `json:"ownMateriel"`
	PlatLastTime map[int]int `json:"platLastTime"`
	PlatProduct  map[int]map[int]int `json:"platProduct"`
	PlatAnimal   map[int]map[int]int `json:"platAnimal"`
}

type Time struct {
	Time 	int `json:"time"`
	Overday int `json:"overday"`
}

type Human struct {
	Hurt   int `json:"hurt"`
	Ill    int `json:"ill"`
	Hungry int `json:"hungry"`
	Thirst int `json:"thirst"`
	Blood  int `json:"blood"`
	Mood   int `json:"mood"`
	Wake   int `json:"wake"`
	Lvl    int `json:"lvl"`
	Exp    int `json:"exp"`
}

func NewFile() *os.File {
	i := 1
	for {
		if ok := Exists(fmt.Sprintf("game%d.json", i)); !ok {
			break
		}

		i++
	}

	file, err := os.Create(fmt.Sprintf("game%d.json", i))
	if err != nil {
		panic(err)
	}

	return file
}

func LoadFile(file string) (*Data, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf := make([]byte, 1024)
	n, err := f.Read(buf)
	if err != nil {
		return nil, err
	}

	var data Data
	err = json.Unmarshal(buf[:n], &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func DeleteFile(file string) error {
	if ok := Exists(file); !ok {
		return nil
	}

	if err := os.Remove(file); err != nil {
		return err
	}

	return nil
}

// 判断所给路径文件/文件夹是否存在
func Exists(file string) bool {
	_, err := os.Stat(file)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func ChooseKeep() []string {
	s := make([]string, 0)
	files, _ := ioutil.ReadDir("./")
	for _, file := range files {
		if file.IsDir() || !strings.Contains(file.Name(), ".json") {
			continue
		} else {
			f := file.Name()
			d, err := getFileInfo(f)
			if err == nil {
				s = append(s, f[:len(f)-5] + "-" + d)
			}
		}
	}
	return s
}

func getFileInfo(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	buf := make([]byte, 2048)
	n, err := f.Read(buf)
	if err != nil {
		return "", err
	}

	var data Data
	err = json.Unmarshal(buf[:n], &data)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s 第%d天 新游戏时间%s, 上次游戏时间%s",
		data.Name, data.GameTime.Overday, data.StartTime, data.LastTime), nil
}
