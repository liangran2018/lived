package home

import (
	"fmt"
	"strconv"

	"github.com/liangran2018/lived/base"
	"github.com/liangran2018/lived/env"
)

var choose []func(ob *OwnBuilding) error

func init() {
	choose = make([]func(ob *OwnBuilding) error, 10)
	choose[1] = roomOpera
	choose[2] = cookOpera
	choose[3] = waterOpera
	choose[4] = drugOpera
	choose[5] = toolOpera
}

func ChooseBuilding() {
	//拥有的建筑
	ob := GetOwnBuilding()
	//展示
	fmt.Println(ob.Show())
	fmt.Println("选择要使用的建筑：1.床  2.火堆  3.净水器  4.药盒  5.工具台  6.取消")
	input, err := base.Input()
	if err != nil || input == "6" {
		return
	}

	i, err := strconv.Atoi(input)
	if err != nil || i < 1 || i > 5 {
		fmt.Println("输入有误")
		return
	}

	//如果选择的是尚未拥有，则走建造流程
	if ob.own[i] == 0 {
		ob.Build(i)
		return
	}

	err = choose[i](ob)
	base.CheckErr(err)
	return
}

func roomOpera(ob *OwnBuilding) error {
	fmt.Println("1.睡一小时  2.睡八小时  3.升级  4.取消")
	input, err := base.Input()
	if err != nil {
		return err
	}

	switch input {
	case "1":
		env.GetTime().Add(1, 0)
	case "2":
		env.GetTime().Add(8, 0)
	case "3":
		err = ob.Build(1)
	default:
	}

	return err
}

func cookOpera(ob *OwnBuilding) error {
	fmt.Println("1.烹饪  2.吃东西  3.升级  4.取消")
	input, err := base.Input()
	if err != nil {
		return err
	}

	switch input {
	case "1":

	case "2":

	case "3":
		err = ob.Build(2)
	default:
	}

	return err
}

func waterOpera(ob *OwnBuilding) error {
	fmt.Println("1.制造净水  2.制造酒精  3.喝水  4.喝酒  5.取消")
	input, err := base.Input()
	if err != nil {
		return err
	}

	switch input {
	case "1":

	case "2":

	case "3":

	case "4":
		err = ob.Build(3)
	default:
	}

	return err
}

func drugOpera(ob *OwnBuilding) error {
	fmt.Println("1.制造  2.服药  3.升级  4.取消")
	input, err := base.Input()
	if err != nil {
		return err
	}

	switch input {
	case "1":

	case "2":

	case "3":
		err = ob.Build(4)
	default:
	}

	return err
}

func toolOpera(ob *OwnBuilding) error {
	fmt.Println("1.制造  2.升级  3.取消")
	input, err := base.Input()
	if err != nil {
		return err
	}

	switch input {
	case "1":

	case "2":
		err = ob.Build(5)
	default:
	}

	return err
}

func (this *OwnBuilding) Build(i int) error {
	fmt.Println(this.lvlupNotice(i))
	fmt.Println("1.升级  2.取消")
	input, err := base.Input()
	if err != nil {
		return err
	}

	if input != "1" {
		return base.NoOper
	}

	//检查物料是否足够升级
	ok := this.Check(i)
	if !ok {
		return base.NotEnough
	}

	//升级，物料使用
	this.Update(i)
	//耗费时间
	env.GetTime().Add(this.TimeNeed(i))
	this.own[i]++
	return nil
}
