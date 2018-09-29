package base

import (
	"os"
	"fmt"
)

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
