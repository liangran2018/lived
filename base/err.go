package base

import (
	"fmt"
	"errors"
)

var NoOper = errors.New("no operate")
var NotEnough = errors.New("not enough")

func CheckErr(err error) {
	switch err {
	case NoOper, nil:
		return
	case NotEnough:
		fmt.Println("物料不足，升级失败")
		return
	default:
		fmt.Println("操作失败")
		return
	}
}